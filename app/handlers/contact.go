package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// ContactHandler handles the "Kontak" menu — read-only directory view.
type ContactHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	contactService *services.ContactService
	querier        *queries.Querier
}

func NewContactHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	contactService *services.ContactService,
	querier *queries.Querier,
) *ContactHandler {
	return &ContactHandler{
		store:          store,
		inertiaService: inertiaService,
		contactService: contactService,
		querier:        querier,
	}
}

func (h *ContactHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
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

// Index — list all contacts (grouped by district). Public GET.
func (h *ContactHandler) Index(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user, _ = h.querier.GetUserByID(c.Context(), uid.(int64))
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "50"))
	search := c.Query("search", "")
	districtID, _ := strconv.ParseInt(c.Query("district_id", "0"), 10, 64)

	result, err := h.contactService.List(c.Context(), search, districtID, page, perPage)
	if err != nil {
		slog.Error("contact list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load contacts: " + err.Error(),
		})
	}

	districts, _ := h.querier.GetActiveDistricts(c.Context())

	return h.inertiaService.Render(c, "app/Kontak", fiber.Map{
		"user":             user,
		"contacts":         result,
		"districts":        districts,
		"current_search":   search,
		"current_district": districtID,
	})
}

// Create — render with create modal open.
func (h *ContactHandler) Create(c *fiber.Ctx) error {
	return c.Redirect("/kontak?action=create")
}

// Store — handle POST /kontak.
func (h *ContactHandler) Store(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	var req services.CreateContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/kontak?action=create&error="+err.Error(), fiber.StatusSeeOther)
	}

	_, err = h.contactService.Create(c.Context(), req)
	if err != nil {
		return c.Redirect("/kontak?action=create&error="+err.Error(), fiber.StatusSeeOther)
	}

	return c.Redirect("/kontak?success=created", fiber.StatusSeeOther)
}

// Edit — render with edit modal opened.
func (h *ContactHandler) Edit(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return c.Redirect(fmt.Sprintf("/kontak?action=edit&id=%d", id))
}

// Update — handle PUT /kontak/:id.
func (h *ContactHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	var req services.UpdateContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect(fmt.Sprintf("/kontak?action=edit&id=%d&error=invalid", id), fiber.StatusSeeOther)
	}

	if err := h.contactService.Update(c.Context(), id, req); err != nil {
		return c.Redirect(fmt.Sprintf("/kontak?action=edit&id=%d&error=%s", id, err.Error()), fiber.StatusSeeOther)
	}

	return c.Redirect("/kontak?success=updated", fiber.StatusSeeOther)
}

// Destroy — handle DELETE /kontak/:id.
func (h *ContactHandler) Destroy(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.contactService.Delete(c.Context(), id); err != nil {
		if errors.Is(err, services.ErrContactNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Kontak tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Redirect("/kontak?success=deleted", fiber.StatusSeeOther)
}
