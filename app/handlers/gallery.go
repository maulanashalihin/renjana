package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// GalleryHandler handles CRUD pages for the "Galeri" menu.
type GalleryHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	staticSvc      *services.StaticService
	querier        *queries.Querier
}

func NewGalleryHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	staticSvc *services.StaticService,
	querier *queries.Querier,
) *GalleryHandler {
	return &GalleryHandler{
		store:          store,
		inertiaService: inertiaService,
		staticSvc:      staticSvc,
		querier:        querier,
	}
}

// getUser reads user from session (works without AuthRequired middleware)
func (h *GalleryHandler) getUser(c *fiber.Ctx) *fiber.Map {
	sess, err := h.store.Get(c)
	if err != nil || sess.Get("user_id") == nil {
		return nil
	}
	uid := sess.Get("user_id").(int64)
	u, err := h.querier.GetUserByID(c.Context(), uid)
	if err != nil {
		return nil
	}
	return &fiber.Map{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"role":  u.Role,
	}
}

// Create — render the upload page for new galeri item.
func (h *GalleryHandler) Create(c *fiber.Ctx) error {
	user := h.getUser(c)
	return h.inertiaService.Render(c, "app/GaleriEditor", fiber.Map{
		"user": user,
		"edit": false,
	})
}

// Store — handle POST /galeri (with form-urlencoded data).
func (h *GalleryHandler) Store(c *fiber.Ctx) error {
	user := h.getUser(c)
	if user == nil {
		return c.Redirect("/login")
	}

	title := c.FormValue("title")
	fileURL := c.FormValue("file_url")
	mediaType := c.FormValue("media_type")
	caption := c.FormValue("caption")
	isPublished := c.FormValue("is_published") == "true"

	if title == "" || fileURL == "" || mediaType == "" {
		return c.Redirect("/galeri?error=missing_fields")
	}

	role, _ := (*user)["role"].(string)
	uploadedBy := int64(0)
	if role == "admin" {
		// Get uploader ID from session
		sess, _ := h.store.Get(c)
		if uid := sess.Get("user_id"); uid != nil {
			uploadedBy = uid.(int64)
		}
	}

	if _, err := h.staticSvc.CreateMedia(c.Context(), title, fileURL, mediaType, caption, uploadedBy, isPublished); err != nil {
		slog.Error("gallery create error", "err", err)
		return c.Redirect("/galeri?error=" + err.Error())
	}

	return c.Redirect("/galeri?success=created")
}

// Edit — render edit form for existing galeri item.
func (h *GalleryHandler) Edit(c *fiber.Ctx) error {
	user := h.getUser(c)
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	// Get the media item via direct query
	row, err := h.querier.GetMediaByID(c.Context(), id)
	if err != nil {
		return c.Redirect("/galeri?error=not_found")
	}

	caption := ""
	if row.Caption.Valid {
		caption = row.Caption.String
	}

	return h.inertiaService.Render(c, "app/GaleriEditor", fiber.Map{
		"user": user,
		"edit": true,
		"media": fiber.Map{
			"id":           row.ID,
			"title":        row.Title,
			"file_url":     row.FileUrl,
			"media_type":   row.MediaType,
			"caption":      caption,
			"is_published": row.IsPublished,
		},
	})
}

// Update — handle PUT /galeri/:id.
func (h *GalleryHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Redirect("/galeri?error=bad_id")
	}

	title := c.FormValue("title")
	fileURL := c.FormValue("file_url")
	mediaType := c.FormValue("media_type")
	caption := c.FormValue("caption")
	isPublished := c.FormValue("is_published") == "true"

	if _, err := h.staticSvc.UpdateMedia(c.Context(), id, title, fileURL, mediaType, caption, isPublished); err != nil {
		slog.Error("gallery update error", "err", err)
		return c.Redirect("/galeri?error=" + err.Error())
	}

	return c.Redirect("/galeri?success=updated")
}

// Destroy — handle DELETE /galeri/:id via _method=DELETE.
func (h *GalleryHandler) Destroy(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Redirect("/galeri?error=bad_id")
	}

	if err := h.staticSvc.DeleteMedia(c.Context(), id); err != nil {
		slog.Error("gallery delete error", "err", err)
		return c.Redirect("/galeri?error=" + err.Error())
	}

	return c.Redirect("/galeri?success=deleted")
}
