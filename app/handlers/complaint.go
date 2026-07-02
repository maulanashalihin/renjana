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
		"user":    user,
		"isAdmin": false,
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

	return h.inertiaService.Render(c, "app/Pengaduan", fiber.Map{
		"user":       user,
		"isAdmin":    true,
		"complaints": result,
		"stats":      stats,
	})
}

// Store — public submission.
func (h *ComplaintHandler) Store(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	category := c.FormValue("category")
	message := c.FormValue("message")

	if name == "" || email == "" || message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama, email, dan pesan harus diisi",
		})
	}

	_, err := h.complaintSvc.Create(c.Context(), name, email, phone, category, message)
	if err != nil {
		slog.Error("complaint create error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengirim pengaduan: " + err.Error(),
		})
	}

	return c.Redirect("/pengaduan?success=true", fiber.StatusSeeOther)
}

// UpdateStatus — admin respond/resolve.
func (h *ComplaintHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	status := c.FormValue("status")
	response := c.FormValue("response")

	userID := c.Locals("user_id").(int64)

	_, err = h.complaintSvc.UpdateStatus(c.Context(), id, status, response, userID)
	if err != nil {
		slog.Error("complaint update error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memperbarui pengaduan: " + err.Error(),
		})
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
