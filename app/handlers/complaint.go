package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// ComplaintHandler handles pengaduan masyarakat.
type ComplaintHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	complaintSvc   *services.ComplaintService
	querier        *queries.Querier
}

func NewComplaintHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	complaintSvc *services.ComplaintService,
	querier *queries.Querier,
) *ComplaintHandler {
	return &ComplaintHandler{
		store:          store,
		inertiaService: inertiaService,
		complaintSvc:   complaintSvc,
		querier:        querier,
	}
}

func (h *ComplaintHandler) getUser(c *fiber.Ctx) *fiber.Map {
	// Try to get user from session (works on public routes without AuthRequired middleware)
	sess, err := h.store.Get(c)
	if err != nil || sess.Get("user_id") == nil {
		return nil
	}
	userID := sess.Get("user_id").(int64)
	role := ""
	if r := sess.Get("role"); r != nil {
		role = r.(string)
	}

	u, err := h.querier.GetUserByID(c.Context(), userID)
	if err != nil {
		// Return basic info from session even if DB lookup fails
		return &fiber.Map{
			"id":   userID,
			"role": role,
		}
	}
	return &fiber.Map{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"role":  string(u.Role),
	}
}

// Index — show public form or admin list.
func (h *ComplaintHandler) Index(c *fiber.Ctx) error {
	user := h.getUser(c)
	isLoggedIn := user != nil
	isAdmin := false
	if isLoggedIn {
		if role, ok := (*user)["role"].(string); ok {
			isAdmin = role == "admin"
		}
	}

	if isAdmin {
		return h.adminIndex(c, user)
	}
	return h.publicIndex(c, user)
}

func (h *ComplaintHandler) publicIndex(c *fiber.Ctx, user *fiber.Map) error {
	return h.inertiaService.Render(c, "app/Pengaduan", fiber.Map{
		"user":      user,
		"isAdmin":   false,
		"submitted": c.Query("success") == "true",
	})
}

func (h *ComplaintHandler) adminIndex(c *fiber.Ctx, user *fiber.Map) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))

	result, err := h.complaintSvc.List(c.Context(), page, perPage)
	if err != nil {
		slog.Error("complaint list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load complaints: " + err.Error(),
		})
	}

	stats, _ := h.complaintSvc.GetStats(c.Context())

	// Get resolved complaints for report
	resolvedPage, _ := strconv.Atoi(c.Query("resolved_page", "1"))
	resolvedResult, _ := h.complaintSvc.ListResolved(c.Context(), resolvedPage, 20)

	return h.inertiaService.Render(c, "app/Pengaduan", fiber.Map{
		"user":       user,
		"isAdmin":    true,
		"complaints": result,
		"stats":      stats,
		"resolved":   resolvedResult,
	})
}

// Store — public submission.
func (h *ComplaintHandler) Store(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Category string `json:"category"`
		Message  string `json:"message"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	if input.Name == "" || input.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama dan pesan harus diisi",
		})
	}

	// Validate field lengths to prevent bomb payload
	if len(input.Name) > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama terlalu panjang (maks 100 karakter)",
		})
	}
	if len(input.Phone) > 15 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nomor telepon terlalu panjang (maks 15 digit)",
		})
	}
	if len(input.Message) > 2000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Pesan terlalu panjang (maks 2000 karakter)",
		})
	}

	// Generate unique token for ticket URL
	token, err := h.complaintSvc.GenerateToken()
	if err != nil {
		slog.Error("token generation error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat tiket pengaduan",
		})
	}

	// Create complaint with token
	complaint, err := h.complaintSvc.Create(c.Context(), input.Name, "", input.Phone, input.Category, input.Message, token)
	if err != nil {
		slog.Error("complaint create error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengirim pengaduan: " + err.Error(),
		})
	}

	// Add the initial complaint message as the first message in the conversation
	_, err = h.complaintSvc.AddMessage(c.Context(), complaint.ID, "user", input.Name, input.Message)
	if err != nil {
		slog.Error("failed to add initial message", "err", err)
	}

	// Redirect to the ticket page
	return c.Redirect("/pengaduan/tiket/"+token, fiber.StatusSeeOther)
}

// ShowTicket — display a complaint ticket with conversation.
func (h *ComplaintHandler) ShowTicket(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Token tidak ditemukan"})
	}

	complaint, err := h.complaintSvc.GetByToken(c.Context(), token)
	if err != nil {
		slog.Error("get complaint by token error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memuat pengaduan: " + err.Error(),
		})
	}
	if complaint == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pengaduan tidak ditemukan"})
	}

	// Get messages
	messages, err := h.complaintSvc.GetMessages(c.Context(), complaint.ID)
	if err != nil {
		slog.Error("get messages error", "err", err)
		messages = []services.ComplaintMessageItem{}
	}

	user := h.getUser(c)
	isAdmin := false
	if user != nil {
		if role, ok := (*user)["role"].(string); ok {
			isAdmin = role == "admin"
		}
	}

	return h.inertiaService.Render(c, "app/PengaduanTicket", fiber.Map{
		"user":      user,
		"isAdmin":   isAdmin,
		"complaint": complaint,
		"messages":  messages,
	})
}

// AddReply — add a message to the complaint conversation (via ticket).
func (h *ComplaintHandler) AddReply(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token diperlukan"})
	}

	complaint, err := h.complaintSvc.GetByToken(c.Context(), token)
	if err != nil || complaint == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pengaduan tidak ditemukan"})
	}

	var input struct {
		SenderName string `json:"sender_name"`
		Message    string `json:"message"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	if input.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Pesan harus diisi",
		})
	}

	// Validate message length to prevent bomb payload
	if len(input.Message) > 2000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Pesan terlalu panjang (maks 2000 karakter)",
		})
	}

	// Determine sender type and name
	senderType := "user"
	senderName := input.SenderName

	user := h.getUser(c)
	isAdmin := false
	if user != nil {
		if role, ok := (*user)["role"].(string); ok {
			isAdmin = role == "admin"
		}
		if isAdmin {
			senderType = "admin"
			if name, ok := (*user)["name"].(string); ok && name != "" {
				senderName = name
			} else {
				senderName = "Admin"
			}
		}
	}

	if senderName == "" {
		senderName = "Pengguna"
	}

	_, err = h.complaintSvc.AddMessage(c.Context(), complaint.ID, senderType, senderName, input.Message)
	if err != nil {
		slog.Error("add reply error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengirim balasan",
		})
	}

	return c.Redirect("/pengaduan/tiket/"+token, fiber.StatusSeeOther)
}

// PublicResolve — user marks complaint as resolved via ticket.
func (h *ComplaintHandler) PublicResolve(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token diperlukan"})
	}

	complaint, err := h.complaintSvc.GetByToken(c.Context(), token)
	if err != nil || complaint == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pengaduan tidak ditemukan"})
	}

	_, err = h.complaintSvc.ResolveByUser(c.Context(), complaint.ID)
	if err != nil {
		slog.Error("resolve by user error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyelesaikan pengaduan",
		})
	}

	return c.Redirect("/pengaduan/tiket/"+token, fiber.StatusSeeOther)
}

// UpdateStatus — admin respond/resolve.
func (h *ComplaintHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var input struct {
		Status   string `json:"status"`
		Response string `json:"response"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	userID := c.Locals("user_id").(int64)

	// Validate response length to prevent bomb payload
	if len(input.Response) > 2000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Respon terlalu panjang (maks 2000 karakter)",
		})
	}

	_, err = h.complaintSvc.UpdateStatus(c.Context(), id, input.Status, input.Response, userID)
	if err != nil {
		slog.Error("complaint update error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memperbarui pengaduan: " + err.Error(),
		})
	}

	// Also add the response as a message in the conversation thread
	if input.Response != "" {
		user, _ := h.querier.GetUserByID(c.Context(), userID)
		adminName := "Admin"
		if user != nil && user.Name != "" {
			adminName = user.Name
		}
		_, err = h.complaintSvc.AddMessage(c.Context(), id, "admin", adminName, input.Response)
		if err != nil {
			slog.Error("failed to add admin response as message", "err", err)
		}
	}

	return c.Redirect("/pengaduan", fiber.StatusSeeOther)
}

// Destroy — admin delete.
func (h *ComplaintHandler) Destroy(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.complaintSvc.Delete(c.Context(), id); err != nil {
		slog.Error("complaint delete error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menghapus pengaduan: " + err.Error(),
		})
	}

	return c.Redirect("/pengaduan", fiber.StatusSeeOther)
}
