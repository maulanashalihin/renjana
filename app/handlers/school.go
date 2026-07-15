package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// SchoolHandler handles school CRUD (admin) + public search API.
type SchoolHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	schoolService  *services.SchoolService
	querier        *queries.Querier
}

func NewSchoolHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	schoolService *services.SchoolService,
	querier *queries.Querier,
) *SchoolHandler {
	return &SchoolHandler{
		store:          store,
		inertiaService: inertiaService,
		schoolService:  schoolService,
		querier:        querier,
	}
}

func (h *SchoolHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return 0, nil, fiber.ErrUnauthorized
	}
	id := userID.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return 0, nil, err
	}
	return id, u, nil
}

// SearchSchoolsAPI — public JSON API for autocomplete.
// GET /api/schools/search?q=...
func (h *SchoolHandler) SearchSchoolsAPI(c *fiber.Ctx) error {
	query := c.Query("q", "")
	if query == "" {
		return c.JSON(fiber.Map{"data": []interface{}{}})
	}

	results, err := h.schoolService.Search(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mencari sekolah",
		})
	}

	if results == nil {
		results = []services.SchoolSearchResult{}
	}

	return c.JSON(fiber.Map{"data": results})
}

// Index — admin list page.
// GET /admin/schools
func (h *SchoolHandler) Index(c *fiber.Ctx) error {
	_, user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	allSchools, err := h.schoolService.ListAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal memuat data sekolah")
	}

	return h.inertiaService.Render(c, "app/Schools", fiber.Map{
		"user":    user,
		"schools": allSchools,
	})
}

// Store — handle POST /admin/schools (create).
func (h *SchoolHandler) Store(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	var input services.SchoolInput
	if err := c.BodyParser(&input); err != nil {
		h.store.Flash(c, "error", "Gagal membaca data: "+err.Error())
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	_, err = h.schoolService.Create(c.Context(), input)
	if err != nil {
		h.store.Flash(c, "error", "Gagal menambah sekolah: "+err.Error())
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Sekolah berhasil ditambahkan")
	return c.Redirect("/admin/schools", fiber.StatusSeeOther)
}

// Update — handle PUT /admin/schools/:id.
func (h *SchoolHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		h.store.Flash(c, "error", "ID sekolah tidak valid")
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	var input services.SchoolInput
	if err := c.BodyParser(&input); err != nil {
		h.store.Flash(c, "error", "Gagal membaca data: "+err.Error())
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	err = h.schoolService.Update(c.Context(), id, input)
	if err != nil {
		h.store.Flash(c, "error", "Gagal mengupdate sekolah: "+err.Error())
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Sekolah berhasil diupdate")
	return c.Redirect("/admin/schools", fiber.StatusSeeOther)
}

// Destroy — handle DELETE /admin/schools/:id.
func (h *SchoolHandler) Destroy(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		h.store.Flash(c, "error", "ID sekolah tidak valid")
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	err = h.schoolService.Delete(c.Context(), id)
	if err != nil {
		h.store.Flash(c, "error", "Gagal menghapus sekolah")
		return c.Redirect("/admin/schools", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Sekolah berhasil dihapus")
	return c.Redirect("/admin/schools", fiber.StatusSeeOther)
}
