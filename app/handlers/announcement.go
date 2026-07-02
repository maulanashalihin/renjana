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

// AnnouncementHandler handles CRUD pages for the "Berita" menu.
type AnnouncementHandler struct {
	store           *session.Store
	inertiaService  *services.InertiaService
	announcementSvc *services.AnnouncementService
	querier         *queries.Querier
}

func NewAnnouncementHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	announcementSvc *services.AnnouncementService,
	querier *queries.Querier,
) *AnnouncementHandler {
	return &AnnouncementHandler{
		store:           store,
		inertiaService:  inertiaService,
		announcementSvc: announcementSvc,
		querier:         querier,
	}
}

func (h *AnnouncementHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
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

// Index — list with search, filter, pagination. Public GET.
func (h *AnnouncementHandler) Index(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user, _ = h.querier.GetUserByID(c.Context(), uid.(int64))
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	search := c.Query("search", "")
	category := c.Query("category", "")
	isPublished := c.Query("is_published", "")

	result, err := h.announcementSvc.List(c.Context(), search, category, isPublished, page, perPage)
	if err != nil {
		slog.Error("announcement list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load announcements: " + err.Error(),
		})
	}

	return h.inertiaService.Render(c, "app/Berita", fiber.Map{
		"user":              user,
		"announcements":     result,
		"current_search":    search,
		"current_category":  category,
		"current_published": isPublished,
	})
}

// Create — render with create modal open.
func (h *AnnouncementHandler) Create(c *fiber.Ctx) error {
	// Detect user from session
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user, _ = h.querier.GetUserByID(c.Context(), uid.(int64))
		}
	}
	return h.inertiaService.Render(c, "app/BeritaEditor", fiber.Map{
		"user": user,
		"edit": false,
	})
}

// Store — handle POST /berita.
func (h *AnnouncementHandler) Store(c *fiber.Ctx) error {
	userID, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	req := services.CreateAnnouncementRequest{
		Title:       c.FormValue("title"),
		Content:     c.FormValue("content"),
		Category:    c.FormValue("category"),
		Body:        c.FormValue("body"),
		CoverURL:    c.FormValue("cover_url"),
		AuthorID:    userID,
		IsPublished: c.FormValue("is_published") == "true",
	}

	_, err = h.announcementSvc.Create(c.Context(), req)
	if err != nil {
		return c.Redirect("/berita?error=" + err.Error())
	}

	return c.Redirect("/berita?success=created")
}

// Edit — render with edit modal opened.
func (h *AnnouncementHandler) Edit(c *fiber.Ctx) error {
	// Detect user from session
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user, _ = h.querier.GetUserByID(c.Context(), uid.(int64))
		}
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	ann, _ := h.announcementSvc.Get(c.Context(), id)
	return h.inertiaService.Render(c, "app/BeritaEditor", fiber.Map{
		"user":         user,
		"edit":         true,
		"announcement": ann,
	})
}

// Update — handle PUT /berita/:id.
func (h *AnnouncementHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	var req services.UpdateAnnouncementRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect(fmt.Sprintf("/berita?action=edit&id=%d&error=invalid", id))
	}

	if err := h.announcementSvc.Update(c.Context(), id, req); err != nil {
		return c.Redirect(fmt.Sprintf("/berita?action=edit&id=%d&error=%s", id, err.Error()))
	}

	return c.Redirect("/berita?success=updated")
}

// Destroy — handle DELETE /berita/:id.
func (h *AnnouncementHandler) Destroy(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.announcementSvc.Delete(c.Context(), id); err != nil {
		if errors.Is(err, services.ErrAnnouncementNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Berita tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Redirect("/berita?success=deleted")
}

// Show — view detail.
func (h *AnnouncementHandler) Show(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user, _ = h.querier.GetUserByID(c.Context(), uid.(int64))
		}
	}

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	ann, err := h.announcementSvc.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrAnnouncementNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Berita tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return h.inertiaService.Render(c, "app/AnnouncementDetail", fiber.Map{
		"user":         user,
		"announcement": ann,
	})
}
