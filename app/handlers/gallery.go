package handlers

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
// Accepts file_urls (comma-separated for batch) or single file_url (deprecated).
func (h *GalleryHandler) Store(c *fiber.Ctx) error {
	user := h.getUser(c)
	if user == nil {
		return c.Redirect("/login")
	}

	title := c.FormValue("title")
	caption := c.FormValue("caption")
	isPublished := c.FormValue("is_published") == "true"

	// Collect all file URLs
	var fileURLs []string
	if raw := c.FormValue("file_urls"); raw != "" {
		for _, u := range strings.Split(raw, ",") {
			u = strings.TrimSpace(u)
			if u != "" {
				fileURLs = append(fileURLs, u)
			}
		}
	} else if single := c.FormValue("file_url"); single != "" {
		fileURLs = append(fileURLs, single)
	}

	if title == "" || len(fileURLs) == 0 {
		return c.Redirect("/galeri?error=missing_fields")
	}

	role, _ := (*user)["role"].(string)
	uploadedBy := int64(0)
	if role == "admin" {
		sess, _ := h.store.Get(c)
		if uid := sess.Get("user_id"); uid != nil {
			uploadedBy = uid.(int64)
		}
	}

	// Generate album_id to group all files in this batch
	albumID := uuid.New().String()

	for _, fileURL := range fileURLs {
		mt := detectMediaType(fileURL)
		if _, err := h.staticSvc.CreateMedia(c.Context(), title, fileURL, mt, caption, uploadedBy, isPublished, albumID); err != nil {
			slog.Error("gallery create error", "file_url", fileURL, "err", err)
			return c.Redirect("/galeri?error=" + err.Error())
		}
	}

	return c.Redirect("/galeri?success=created")
}

// detectMediaType guesses media type from file URL extension.
func detectMediaType(url string) string {
	lower := strings.ToLower(url)
	if strings.HasSuffix(lower, ".mp4") || strings.HasSuffix(lower, ".webm") || strings.HasSuffix(lower, ".mov") || strings.HasSuffix(lower, ".avi") {
		return "video"
	}
	return "image"
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

// EditAlbum — render edit form for an entire album.
func (h *GalleryHandler) EditAlbum(c *fiber.Ctx) error {
	user := h.getUser(c)
	albumID := c.Params("id")

	items, err := h.staticSvc.GetMediaByAlbumID(c.Context(), albumID)
	if err != nil || len(items) == 0 {
		return c.Redirect("/galeri?error=not_found")
	}

	// Convert MediaItems to map for JSON serialization
	mediaList := make([]fiber.Map, len(items))
	for i, m := range items {
		mediaList[i] = fiber.Map{
			"id":           m.ID,
			"title":        m.Title,
			"file_url":     m.FileURL,
			"media_type":   m.MediaType,
			"caption":      m.Caption,
			"is_published": m.IsPublished,
		}
	}

	return h.inertiaService.Render(c, "app/GaleriEditor", fiber.Map{
		"user":     user,
		"edit":     true,
		"album_id": albumID,
		"media":    mediaList,
	})
}

// UpdateAlbum — handle PUT /galeri/album/:albumId — update album title and photos.
func (h *GalleryHandler) UpdateAlbum(c *fiber.Ctx) error {
	albumID := c.Params("albumId")
	if albumID == "" {
		return c.Redirect("/galeri?error=bad_id")
	}

	title := c.FormValue("title")
	caption := c.FormValue("caption")
	isPublished := c.FormValue("is_published") == "true"

	// Collect all file URLs
	var fileURLs []string
	if raw := c.FormValue("file_urls"); raw != "" {
		for _, u := range strings.Split(raw, ",") {
			u = strings.TrimSpace(u)
			if u != "" {
				fileURLs = append(fileURLs, u)
			}
		}
	}

	if title == "" || len(fileURLs) == 0 {
		return c.Redirect("/galeri?error=missing_fields")
	}

	// Get uploader
	user := h.getUser(c)
	uploadedBy := int64(0)
	if user != nil {
		if role, _ := (*user)["role"].(string); role == "admin" {
			sess, _ := h.store.Get(c)
			if uid := sess.Get("user_id"); uid != nil {
				uploadedBy = uid.(int64)
			}
		}
	}

	// Delete existing items and recreate
	if err := h.staticSvc.DeleteMediaByAlbumID(c.Context(), albumID); err != nil {
		slog.Error("album update delete error", "album_id", albumID, "err", err)
		return c.Redirect("/galeri?error=" + err.Error())
	}

	for _, fileURL := range fileURLs {
		mt := detectMediaType(fileURL)
		if _, err := h.staticSvc.CreateMedia(c.Context(), title, fileURL, mt, caption, uploadedBy, isPublished, albumID); err != nil {
			slog.Error("album update create error", "file_url", fileURL, "err", err)
			return c.Redirect("/galeri?error=" + err.Error())
		}
	}

	return c.Redirect("/galeri?success=updated")
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

// DestroyAlbum — handle DELETE /galeri/album/:albumId — deletes all media in an album.
func (h *GalleryHandler) DestroyAlbum(c *fiber.Ctx) error {
	albumID := c.Params("albumId")
	if albumID == "" {
		return c.Redirect("/galeri?error=bad_id")
	}

	if err := h.staticSvc.DeleteMediaByAlbumID(c.Context(), albumID); err != nil {
		slog.Error("gallery delete album error", "album_id", albumID, "err", err)
		return c.Redirect("/galeri?error=" + err.Error())
	}

	return c.Redirect("/galeri?deleted=ok")
}
