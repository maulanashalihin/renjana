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

// EducationHandler handles LMS operations for courses, quizzes, and certificates.
type EducationHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	educationSvc   *services.EducationService
	querier        *queries.Querier
}

func NewEducationHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	educationSvc *services.EducationService,
	querier *queries.Querier,
) *EducationHandler {
	return &EducationHandler{
		store:          store,
		inertiaService: inertiaService,
		educationSvc:   educationSvc,
		querier:        querier,
	}
}

// getUser returns the authenticated user or nil for public access.
func (h *EducationHandler) getUser(c *fiber.Ctx) *models.User {
	userID := c.Locals("user_id")
	if userID == nil {
		return nil
	}
	id := userID.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return nil
	}
	return u
}

// getUserID returns the authenticated user ID or 0.
func (h *EducationHandler) getUserID(c *fiber.Ctx) int64 {
	userID := c.Locals("user_id")
	if userID == nil {
		return 0
	}
	return userID.(int64)
}

// CourseShow displays a course with its modules.
// Route: GET /edukasi/course/:id
func (h *EducationHandler) CourseShow(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	user := h.getUser(c)
	userID := h.getUserID(c)

	detail, err := h.educationSvc.GetCourseDetail(c.Context(), id, userID)
	if err != nil {
		slog.Error("course detail error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load course",
		})
	}
	if detail == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Course not found",
		})
	}

	return h.inertiaService.Render(c, "app/EdukasiCourse", fiber.Map{
		"user":            user,
		"course":          detail.Course,
		"modules":         detail.Modules,
		"progress":        detail.Progress,
		"quiz_count":      detail.QuizCount,
		"has_certificate": detail.HasCertificate,
		"certificate_id":  detail.CertificateID,
	})
}

// QuizShow displays the quiz page for a course.
// Route: GET /edukasi/course/:id/quiz
func (h *EducationHandler) QuizShow(c *fiber.Ctx) error {
	userID := h.getUserID(c)
	if userID == 0 {
		return c.Redirect("/login")
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	user := h.getUser(c)

	// Check if user already has a certificate
	cert, _ := h.educationSvc.GetCertificate(c.Context(), userID, id)
	if cert != nil {
		return c.Redirect("/edukasi/course/" + strconv.FormatInt(id, 10) + "/certificate")
	}

	questions, err := h.educationSvc.GetQuizQuestions(c.Context(), id)
	if err != nil {
		slog.Error("quiz questions error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load quiz questions",
		})
	}

	if len(questions) == 0 {
		return c.Redirect("/edukasi/course/" + strconv.FormatInt(id, 10))
	}

	return h.inertiaService.Render(c, "app/EdukasiQuiz", fiber.Map{
		"user":      user,
		"course_id": id,
		"questions": questions,
	})
}

// QuizSubmit processes quiz answers and returns the result.
// Route: POST /edukasi/course/:id/quiz
func (h *EducationHandler) QuizSubmit(c *fiber.Ctx) error {
	userID := h.getUserID(c)
	if userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authenticated",
		})
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	user := h.getUser(c)

	// Parse submitted answers
	var req struct {
		Answers map[string]int64 `json:"answers"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert string keys to int64
	selectedAnswers := make(map[int64]int64)
	for k, v := range req.Answers {
		qID, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			continue
		}
		selectedAnswers[qID] = v
	}

	result, err := h.educationSvc.SubmitQuizAttempt(c.Context(), userID, id, selectedAnswers)
	if err != nil {
		slog.Error("quiz submit error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to submit quiz: " + err.Error(),
		})
	}

	// If passed, redirect to certificate
	if result.Passed {
		return c.Redirect("/edukasi/course/" + strconv.FormatInt(id, 10) + "/certificate")
	}

	// Show result page
	return h.inertiaService.Render(c, "app/EdukasiQuizResult", fiber.Map{
		"user":   user,
		"result": result,
	})
}

// CertificateShow displays a certificate for a course.
// Route: GET /edukasi/course/:id/certificate
func (h *EducationHandler) CertificateShow(c *fiber.Ctx) error {
	userID := h.getUserID(c)
	if userID == 0 {
		return c.Redirect("/login")
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course ID",
		})
	}

	user := h.getUser(c)

	cert, err := h.educationSvc.GetCertificate(c.Context(), userID, id)
	if err != nil {
		slog.Error("certificate error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load certificate",
		})
	}
	if cert == nil {
		return c.Redirect("/edukasi/course/" + strconv.FormatInt(id, 10) + "/quiz")
	}

	return h.inertiaService.Render(c, "app/EdukasiCertificate", fiber.Map{
		"user":        user,
		"certificate": cert,
	})
}

// CertificatePublic displays a certificate by code (public access).
// Route: GET /edukasi/sertifikat/:code
func (h *EducationHandler) CertificatePublic(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Certificate code is required",
		})
	}

	cert, err := h.educationSvc.GetCertificateByCode(c.Context(), code)
	if err != nil {
		slog.Error("certificate lookup error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load certificate",
		})
	}
	if cert == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Certificate not found",
		})
	}

	return h.inertiaService.Render(c, "app/EdukasiCertificate", fiber.Map{
		"certificate": cert,
	})
}

// MyCertificates lists all certificates for the current user.
// Route: GET /app/sertifikat-saya
func (h *EducationHandler) MyCertificates(c *fiber.Ctx) error {
	userID := h.getUserID(c)
	if userID == 0 {
		return c.Redirect("/login")
	}

	user := h.getUser(c)

	certs, err := h.educationSvc.ListUserCertificates(c.Context(), userID)
	if err != nil {
		slog.Error("list certificates error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load certificates",
		})
	}

	return h.inertiaService.Render(c, "app/SertifikatSaya", fiber.Map{
		"user":         user,
		"certificates": certs,
	})
}
