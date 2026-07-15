package handlers

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

type UploadHandler struct {
	store       *session.Store
	userService *services.UserService
}

func NewUploadHandler(store *session.Store, userService *services.UserService) *UploadHandler {
	return &UploadHandler{
		store:       store,
		userService: userService,
	}
}

// Upload subdirectories by purpose.
const (
	UploadTypeAvatar   = "avatar"
	UploadTypeDocument = "document"
	UploadTypeMedia    = "media"
)

// per-type config: allowed MIME types + max size (bytes)
var uploadConfigs = map[string]struct {
	allowedTypes []string
	maxSize      int64
	folder       string
}{
	UploadTypeAvatar: {
		allowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
		maxSize:      5 * 1024 * 1024, // 5MB
		folder:       "avatars",
	},
	UploadTypeMedia: {
		allowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp", "video/mp4", "video/webm"},
		maxSize:      20 * 1024 * 1024, // 20MB
		folder:       "media",
	},
	UploadTypeDocument: {
		allowedTypes: []string{
			"application/pdf",
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // .docx
			"application/vnd.ms-excel", // .xls
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",         // .xlsx
			"application/vnd.ms-powerpoint",                                             // .ppt
			"application/vnd.openxmlformats-officedocument.presentationml.presentation", // .pptx
			"text/plain",
		},
		maxSize: 20 * 1024 * 1024, // 20MB
		folder:  "documents",
	},
}

// Upload is the backward-compatible avatar upload (used by Profile).
func (h *UploadHandler) Upload(c *fiber.Ctx) error {
	return h.handleUpload(c, UploadTypeAvatar)
}

// UploadByPurpose lets clients choose a destination folder via form field `purpose`.
// Used by Berita (cover image), Dokumen (file), Galeri (media).
func (h *UploadHandler) UploadByPurpose(c *fiber.Ctx) error {
	purpose := strings.ToLower(strings.TrimSpace(c.FormValue("purpose")))
	switch purpose {
	case UploadTypeDocument:
		return h.handleUpload(c, UploadTypeDocument)
	case UploadTypeMedia:
		return h.handleUpload(c, UploadTypeMedia)
	case UploadTypeAvatar, "":
		return h.handleUpload(c, UploadTypeAvatar)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unknown upload purpose: " + purpose,
		})
	}
}

func (h *UploadHandler) handleUpload(c *fiber.Ctx, purpose string) error {
	sess, _ := h.store.Get(c)
	userID := sess.Get("user_id")

	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authenticated",
		})
	}

	cfg, ok := uploadConfigs[purpose]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid purpose: " + purpose,
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form",
		})
	}

	files := form.File["file"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	file := files[0]
	contentType := file.Header.Get("Content-Type")

	// Validate type
	isAllowed := false
	for _, allowed := range cfg.allowedTypes {
		if contentType == allowed {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		slog.Warn("upload invalid file type", "purpose", purpose, "content_type", contentType)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid file type for %s. Allowed: %v", purpose, cfg.allowedTypes),
		})
	}

	if file.Size > cfg.maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("File too large. Max: %d KB", cfg.maxSize/1024),
		})
	}

	// Generate unique filename: {userID}_{timestamp}_{nanoseconds}{ext}
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%d%s", userID.(int64), time.Now().UnixNano(), ext)

	uploadPath := filepath.Join("storage", cfg.folder, filename)
	if err := c.SaveFile(file, uploadPath); err != nil {
		slog.Error("upload failed to save file", "purpose", purpose, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	url := fmt.Sprintf("/storage/%s/%s", cfg.folder, filename)

	// Avatar upload has special behavior: update user avatar record.
	if purpose == UploadTypeAvatar {
		if err := h.userService.UpdateAvatar(userID.(int64), url); err != nil {
			slog.Error("upload failed to update avatar in DB", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update avatar",
			})
		}

		// Sync session with new avatar URL
		sess.Set("avatar", url)
		sess.Save()
	}

	slog.Info("upload success", "purpose", purpose, "url", url, "size", file.Size)

	return c.JSON(fiber.Map{
		"success":  true,
		"url":      url,
		"purpose":  purpose,
		"filename": filename,
		"size":     file.Size,
		"message":  "File uploaded successfully",
	})
}
