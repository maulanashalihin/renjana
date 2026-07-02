package handlers

import (
	"fmt"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// DocumentHandler handles document CRUD for admin on /dokumen.
type DocumentHandler struct {
	store     *session.Store
	staticSvc *services.StaticService
	uploadCfg uploadConfig
}

// uploadConfig is a subset of UploadHandler's config.
type uploadConfig struct {
	allowedTypes []string
	maxSize      int64
	folder       string
}

func NewDocumentHandler(
	store *session.Store,
	staticSvc *services.StaticService,
) *DocumentHandler {
	return &DocumentHandler{
		store:     store,
		staticSvc: staticSvc,
		uploadCfg: uploadConfig{
			allowedTypes: []string{
				"application/pdf",
				"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
				"application/vnd.ms-excel",
				"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
				"application/vnd.ms-powerpoint",
				"application/vnd.openxmlformats-officedocument.presentationml.presentation",
				"text/plain",
			},
			maxSize: 20 * 1024 * 1024, // 20MB
			folder:  "documents",
		},
	}
}

// authUserID extracts the current user ID from the session.
func (h *DocumentHandler) authUserID(c *fiber.Ctx) (int64, error) {
	rawID := c.Locals("user_id")
	if rawID == nil {
		return 0, fiber.ErrUnauthorized
	}
	id, ok := rawID.(int64)
	if !ok {
		return 0, fiber.ErrUnauthorized
	}
	return id, nil
}

// saveFile saves the uploaded file to storage and returns the URL + size + original name.
func (h *DocumentHandler) saveFile(c *fiber.Ctx, file *multipart.FileHeader, userID int64) (string, int64, string, error) {
	contentType := file.Header.Get("Content-Type")
	isAllowed := false
	for _, allowed := range h.uploadCfg.allowedTypes {
		if contentType == allowed {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return "", 0, "", fmt.Errorf("invalid file type. Allowed: PDF, DOCX, XLS, XLSX, PPT, PPTX, TXT")
	}
	if file.Size > h.uploadCfg.maxSize {
		return "", 0, "", fmt.Errorf("file too large. Max: %d MB", h.uploadCfg.maxSize/1024/1024)
	}

	originalName := file.Filename
	ext := filepath.Ext(originalName)
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().UnixNano(), ext)
	uploadPath := filepath.Join("storage", h.uploadCfg.folder, filename)

	if err := c.SaveFile(file, uploadPath); err != nil {
		return "", 0, "", fmt.Errorf("gagal menyimpan file: %w", err)
	}

	url := fmt.Sprintf("/storage/%s/%s", h.uploadCfg.folder, filename)
	return url, file.Size, originalName, nil
}

// Create — POST /dokumen (admin only)
func (h *DocumentHandler) Create(c *fiber.Ctx) error {
	userID, err := h.authUserID(c)
	if err != nil {
		return c.Redirect("/login")
	}

	title := c.FormValue("title", "")
	category := c.FormValue("category", "SOP")
	description := c.FormValue("description", "")
	versionStr := c.FormValue("version", "1")
	version, _ := strconv.ParseInt(versionStr, 10, 64)

	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		h.store.Flash(c, "error", "File wajib diupload: "+err.Error())
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	fileURL, fileSize, originalName, err := h.saveFile(c, file, userID)
	if err != nil {
		slog.Error("document: upload failed", "err", err)
		h.store.Flash(c, "error", err.Error())
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	_, err = h.staticSvc.CreateDocument(c.Context(), services.CreateDocumentRequest{
		Title:        title,
		FileURL:      fileURL,
		Category:     category,
		Version:      version,
		FileSize:     fileSize,
		Description:  description,
		OriginalName: originalName,
		UploadedBy:   userID,
	})
	if err != nil {
		slog.Error("document: create failed", "err", err, "user_id", userID)
		h.store.Flash(c, "error", err.Error())
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Dokumen berhasil diupload.")
	return c.Redirect("/dokumen", fiber.StatusSeeOther)
}

// Update — PUT /dokumen/:id (admin only)
// Accepts both multipart/form-data (with optional file) and JSON (metadata only).
func (h *DocumentHandler) Update(c *fiber.Ctx) error {
	_, err := h.authUserID(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	// Read form values (works for both multipart and URL-encoded)
	title := c.FormValue("title", "")
	category := c.FormValue("category", "")
	description := c.FormValue("description", "")
	versionStr := c.FormValue("version", "1")
	version, _ := strconv.ParseInt(versionStr, 10, 64)

	if title == "" {
		h.store.Flash(c, "error", "Judul dokumen wajib diisi.")
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	// Get the existing document to preserve file_url if no new file
	existing, err := h.staticSvc.GetDocumentByID(c.Context(), id)
	if err != nil || existing == nil {
		h.store.Flash(c, "error", "Dokumen tidak ditemukan.")
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	fileURL := existing.FileURL
	fileSize := existing.FileSize
	originalName := existing.OriginalName

	// Check if a new file was uploaded
	file, fileErr := c.FormFile("file")
	if fileErr == nil && file != nil {
		var newFileURL string
		var newFileSize int64
		var newOriginalName string
		newFileURL, newFileSize, newOriginalName, err = h.saveFile(c, file, existing.UploadedBy)
		if err != nil {
			slog.Error("document: file upload failed on update", "err", err)
			h.store.Flash(c, "error", err.Error())
			return c.Redirect("/dokumen", fiber.StatusSeeOther)
		}
		fileURL = newFileURL
		fileSize = newFileSize
		originalName = newOriginalName

		// TODO: delete old file from storage
	}

	req := services.UpdateDocumentRequest{
		Title:        title,
		FileURL:      fileURL,
		Category:     category,
		Version:      version,
		FileSize:     fileSize,
		Description:  description,
		OriginalName: originalName,
	}

	if err := h.staticSvc.UpdateDocument(c.Context(), id, req); err != nil {
		slog.Error("document: update failed", "err", err, "id", id)
		h.store.Flash(c, "error", err.Error())
		return c.Redirect("/dokumen", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Dokumen berhasil diperbarui.")
	return c.Redirect("/dokumen", fiber.StatusSeeOther)
}

// Destroy — DELETE /dokumen/:id (admin only)
func (h *DocumentHandler) Destroy(c *fiber.Ctx) error {
	_, err := h.authUserID(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	if err := h.staticSvc.DeleteDocument(c.Context(), id); err != nil {
		return c.Redirect(fmt.Sprintf("/dokumen?error=%s", err.Error()), fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Dokumen berhasil dihapus.")
	return c.Redirect("/dokumen", fiber.StatusSeeOther)
}
