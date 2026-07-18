package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// PartnerHandler handles partner CRUD for the "Mitra & Kolaborasi" section.
type PartnerHandler struct {
	store          *session.Store
	partnerService *services.PartnerService
}

func NewPartnerHandler(
	store *session.Store,
	partnerService *services.PartnerService,
) *PartnerHandler {
	return &PartnerHandler{
		store:          store,
		partnerService: partnerService,
	}
}

// List — GET /api/partners
func (h *PartnerHandler) List(c *fiber.Ctx) error {
	partners, err := h.partnerService.List(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memuat daftar mitra",
		})
	}
	if partners == nil {
		partners = []services.PartnerItem{}
	}
	return c.JSON(fiber.Map{"data": partners})
}

// Create — POST /api/partners
func (h *PartnerHandler) Create(c *fiber.Ctx) error {
	var input services.CreatePartnerInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	if err := h.partnerService.Create(c.Context(), input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Mitra berhasil ditambahkan"})
}

// Update — PUT /api/partners/:id
func (h *PartnerHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID tidak valid",
		})
	}

	var input services.UpdatePartnerInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	if err := h.partnerService.Update(c.Context(), id, input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Mitra berhasil diupdate"})
}

// Delete — DELETE /api/partners/:id
func (h *PartnerHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID tidak valid",
		})
	}

	if err := h.partnerService.Delete(c.Context(), id); err != nil {
		if err == services.ErrPartnerNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mitra tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Mitra berhasil dihapus"})
}
