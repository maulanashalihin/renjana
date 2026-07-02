package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// StaticHandler serves read-only content pages (Edukasi, Galeri, Dokumen, Peta).
type StaticHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	staticSvc      *services.StaticService
	querier        *queries.Querier
}

func NewStaticHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	staticSvc *services.StaticService,
	querier *queries.Querier,
) *StaticHandler {
	return &StaticHandler{
		store:          store,
		inertiaService: inertiaService,
		staticSvc:      staticSvc,
		querier:        querier,
	}
}

// getUser returns the authenticated user or nil for public access.
// Uses session directly (works without AuthRequired middleware).
func (h *StaticHandler) getUser(c *fiber.Ctx) *models.User {
	sess, err := h.store.Get(c)
	if err != nil {
		return nil
	}
	uid := sess.Get("user_id")
	if uid == nil {
		return nil
	}
	id := uid.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return nil
	}
	return u
}

// Edukasi — list education articles.
func (h *StaticHandler) Edukasi(c *fiber.Ctx) error {
	user := h.getUser(c)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	category := c.Query("category", "")

	result, err := h.staticSvc.ListEducation(c.Context(), category, page, perPage)
	if err != nil {
		slog.Error("education list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load education: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Edukasi", fiber.Map{
		"user":             user,
		"articles":         result,
		"current_category": category,
	})
}

// Galeri — list gallery media.
func (h *StaticHandler) Galeri(c *fiber.Ctx) error {
	user := h.getUser(c)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))

	// Check if we're viewing an album detail
	if albumID := c.Params("id"); albumID != "" && albumID != "create" {
		items, err := h.staticSvc.GetMediaByAlbumID(c.Context(), albumID)
		if err != nil || len(items) == 0 {
			return c.Redirect("/galeri")
		}
		return h.inertiaService.Render(c, "app/GaleriAlbum", fiber.Map{
			"user":  user,
			"media": items,
			"album": fiber.Map{
				"title": items[0].Title,
				"count": len(items),
			},
		})
	}

	result, err := h.staticSvc.ListAlbums(c.Context(), page, perPage)
	if err != nil {
		slog.Error("album list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load albums: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Galeri", fiber.Map{
		"user":   user,
		"albums": result,
	})
}

// Dokumen — list documents.
func (h *StaticHandler) Dokumen(c *fiber.Ctx) error {
	user := h.getUser(c)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	category := c.Query("category", "")

	result, err := h.staticSvc.ListDocuments(c.Context(), category, page, perPage)
	if err != nil {
		slog.Error("documents list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load documents: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Dokumen", fiber.Map{
		"user":             user,
		"documents":        result,
		"current_category": category,
	})
}

// Peta — map of districts with per-district enrichment data.
func (h *StaticHandler) Peta(c *fiber.Ctx) error {
	user := h.getUser(c)

	districts, _ := h.querier.GetActiveDistricts(c.Context())
	distributionRows, err := h.querier.CountVolunteersByDistrict(c.Context())
	if err != nil {
		slog.Error("volunteer distribution error", "err", err)
		distributionRows = nil
	}

	// Convert rows to a snake_case JSON shape for the frontend.
	distribution := make([]fiber.Map, 0, len(distributionRows))
	totalVolunteers := int64(0)
	for _, d := range distributionRows {
		distribution = append(distribution, fiber.Map{
			"id":              d.DistrictID,
			"name":            d.DistrictName,
			"volunteer_count": d.VolunteerCount,
		})
		totalVolunteers += d.VolunteerCount
	}

	// --- Per-district enrichment data ---

	// Schools per district
	schoolRows, _ := h.querier.CountSchoolsByDistrict(c.Context())
	schoolMap := make(map[string]int64, len(schoolRows))
	for _, s := range schoolRows {
		schoolMap[s.DistrictName] = s.SchoolCount
	}

	// Activity counts per district
	activityRows, _ := h.querier.CountActivitiesByDistrict(c.Context())
	activityMap := make(map[string]int64, len(activityRows))
	for _, a := range activityRows {
		activityMap[a.DistrictName] = a.ActivityCount
	}

	// Activity type breakdown per district
	typeRows, _ := h.querier.CountActivityTypesByDistrict(c.Context())
	typeBreakdown := make(map[string][]fiber.Map) // district_name -> [{type_name, type_color, count}]
	for _, t := range typeRows {
		typeBreakdown[t.DistrictName] = append(typeBreakdown[t.DistrictName], fiber.Map{
			"type_name":  t.TypeName,
			"type_color": t.TypeColor,
			"count":      t.ActivityCount,
		})
	}

	// Volunteer status breakdown per district
	statusRows, _ := h.querier.CountVolunteerStatusByDistrict(c.Context())
	statusBreakdown := make(map[string]fiber.Map)
	for _, s := range statusRows {
		name := s.DistrictName
		entry, ok := statusBreakdown[name]
		if !ok {
			entry = fiber.Map{"aktif": int64(0), "nonaktif": int64(0)}
		}
		status := "nonaktif"
		if s.Status.Valid && s.Status.String == "aktif" {
			status = "aktif"
		}
		entry[status] = s.VolunteerCount
		statusBreakdown[name] = entry
	}

	// Build district_detail map keyed by district name
	districtDetail := make([]fiber.Map, 0, len(distribution))
	for _, d := range distribution {
		name := d["name"].(string)
		schools := schoolMap[name]
		activities := activityMap[name]
		types := typeBreakdown[name]
		if types == nil {
			types = []fiber.Map{}
		}
		statuses, _ := statusBreakdown[name]
		if statuses == nil {
			statuses = fiber.Map{"aktif": int64(0), "nonaktif": int64(0)}
		}

		districtDetail = append(districtDetail, fiber.Map{
			"id":                 d["id"],
			"name":               name,
			"volunteer_count":    d["volunteer_count"],
			"school_count":       schools,
			"activity_count":     activities,
			"activity_breakdown": types,
			"volunteer_status":   statuses,
		})
	}

	return h.inertiaService.Render(c, "app/Peta", fiber.Map{
		"user":             user,
		"districts":        districts,
		"distribution":     distribution,
		"total_volunteers": totalVolunteers,
		"district_detail":  districtDetail,
	})
}
