package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/handlers"
	"github.com/maulanashalihin/laju-go/app/middlewares"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

type Handlers struct {
	Auth          *handlers.AuthHandler
	App           *handlers.AppHandler
	Upload        *handlers.UploadHandler
	PasswordReset *handlers.PasswordResetHandler
	Volunteer     *handlers.VolunteerHandler
	Activity      *handlers.ActivityHandler
	Announcement  *handlers.AnnouncementHandler
	Contact       *handlers.ContactHandler
	Organization  *handlers.OrganizationHandler
	Onboarding    *handlers.OnboardingHandler
	Static        *handlers.StaticHandler
	UserAdmin     *handlers.UserAdminHandler
	Complaint     *handlers.ComplaintHandler
	Survey        *handlers.SurveyHandler
	Education     *handlers.EducationHandler
}

func SetupRoutes(app *fiber.App, h Handlers, store *session.Store, mailerService *services.MailerService, csrfMiddleware *middlewares.CSRFMiddleware) {
	// Setup static file serving
	setupStaticRoutes(app)

	// Client-side error reporting (no auth required)
	app.Post("/api/errors", handlers.HandleClientError)

	// Setup auth routes
	setupAuthRoutes(app, h.Auth, h.PasswordReset, store, mailerService)

	// Setup public content routes — no auth required
	setupPublicRoutes(app, h.App, h.Activity, h.Announcement, h.Contact, h.Organization, h.Static, h.Volunteer)

	// Setup education LMS routes — course detail, quiz, certificate
	setupEducationRoutes(app, h.Education, store)

	// Setup public form routes — no auth required
	setupPublicFormRoutes(app, h.Complaint, h.Survey)

	// Setup app routes (protected) — semua di root path
	setupAppRoutes(app, h.App, h.Upload, h.Volunteer, h.Activity, h.Announcement, h.Contact, h.Organization, h.Onboarding, h.Static, h.UserAdmin, h.Complaint, h.Survey, h.Education, store, csrfMiddleware)
}

// setupRegistrationRoutes is deprecated — /daftar flow removed.
// Users now register via /register then complete onboarding at /onboarding.
func setupRegistrationRoutes(app *fiber.App) {
	// Intentionally left blank — /daftar flow removed.
}

func setupStaticRoutes(app *fiber.App) {
	// Static assets with aggressive caching — hashed filenames from Vite are immutable
	// Compress: true caches compressed (brotli/gzip) versions in memory, minimizing CPU reuse.
	app.Static("/dist", "./dist", fiber.Static{
		CacheDuration: 365 * 24 * time.Hour,
		MaxAge:        31536000, // 1 year in seconds
		Compress:      true,
	})
	app.Static("/assets", "./dist/assets", fiber.Static{
		CacheDuration: 365 * 24 * time.Hour,
		MaxAge:        31536000,
		Compress:      true,
	})
	// Public assets (non-hashed, short cache)
	app.Static("/public", "./public", fiber.Static{
		CacheDuration: 1 * time.Hour,
		MaxAge:        3600,
	})
	// Uploaded files (avatars etc. — moderate cache)
	app.Static("/storage", "./storage", fiber.Static{
		CacheDuration: 24 * time.Hour,
		MaxAge:        86400,
	})
}

func setupPublicFormRoutes(app *fiber.App, complaintHandler *handlers.ComplaintHandler, surveyHandler *handlers.SurveyHandler) {
	// Pengaduan — public submit, admin manage
	app.Get("/pengaduan", complaintHandler.Index)
	app.Post("/pengaduan", complaintHandler.Store)

	// Survey Pelayanan Publik — public submit, admin results
	app.Get("/survey", surveyHandler.Index)
	app.Post("/survey", surveyHandler.Store)
}

func setupPublicRoutes(app *fiber.App, appHandler *handlers.AppHandler, activityHandler *handlers.ActivityHandler, announcementHandler *handlers.AnnouncementHandler, contactHandler *handlers.ContactHandler, organizationHandler *handlers.OrganizationHandler, staticHandler *handlers.StaticHandler, volunteerHandler *handlers.VolunteerHandler) {
	// Dashboard — public, no auth required
	app.Get("/", appHandler.Dashboard)
	app.Get("/profile", appHandler.Profile)

	// Static content pages — public, no auth required
	app.Get("/peta", staticHandler.Peta)
	app.Get("/edukasi", staticHandler.Edukasi)
	app.Get("/galeri", staticHandler.Galeri)
	app.Get("/dokumen", staticHandler.Dokumen)

	// Organization profile — public, no auth required
	app.Get("/profil", organizationHandler.Index)

	// Kegiatan — read-only list & detail, public
	app.Get("/kegiatan", activityHandler.Index)
	app.Get("/kegiatan/:id", activityHandler.Show)

	// Relawan — read-only list & detail, public
	app.Get("/relawan", volunteerHandler.Index)
	app.Get("/relawan/:id", volunteerHandler.Show)

	// Read-only public listing
	app.Get("/berita", announcementHandler.Index)
	app.Get("/berita/:id", announcementHandler.Show)
	app.Get("/kontak", contactHandler.Index)
}

func setupEducationRoutes(app *fiber.App, educationHandler *handlers.EducationHandler, store *session.Store) {
	// Course detail — public, no auth required
	app.Get("/edukasi/course/:id", educationHandler.CourseShow)

	// Quiz page — requires auth
	app.Get("/edukasi/course/:id/quiz", middlewares.AuthRequired(store), educationHandler.QuizShow)
	app.Post("/edukasi/course/:id/quiz", middlewares.AuthRequired(store), educationHandler.QuizSubmit)

	// Certificate page — requires auth
	app.Get("/edukasi/course/:id/certificate", middlewares.AuthRequired(store), educationHandler.CertificateShow)

	// Public certificate lookup by code
	app.Get("/edukasi/sertifikat/:code", educationHandler.CertificatePublic)
}

func setupAuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler, passwordResetHandler *handlers.PasswordResetHandler, store *session.Store, mailerService *services.MailerService) {
	// Login routes (with Guest middleware)
	app.Get("/login", middlewares.Guest(store), authHandler.ShowLoginForm)
	app.Post("/login", middlewares.Guest(store), authHandler.Login, middlewares.AuthRateLimit.Limit())

	// Register routes (with Guest middleware)
	app.Get("/register", middlewares.Guest(store), authHandler.ShowRegisterForm)
	app.Post("/register", middlewares.Guest(store), authHandler.Register, middlewares.AuthRateLimit.Limit())

	// OAuth routes
	app.Get("/auth/google", authHandler.GoogleLogin)
	app.Get("/auth/google/callback", authHandler.GoogleCallback)

	// Logout (requires auth)
	app.Post("/logout", middlewares.AuthRequired(store), authHandler.Logout)

	// API: Get current user
	app.Get("/api/me", middlewares.AuthRequired(store), authHandler.Me)

	// API: Get user avatar (proxied from external URL)
	app.Get("/api/avatar/:id", authHandler.GetAvatar)

	// Password reset routes
	app.Get("/forgot-password", passwordResetHandler.ShowForgotPasswordForm)
	app.Post("/forgot-password", passwordResetHandler.SendResetLink, middlewares.PasswordResetRateLimit.Limit())
	app.Get("/reset-password/:token", passwordResetHandler.ShowResetPasswordForm)
	app.Post("/reset-password/:token", passwordResetHandler.ResetPassword)
}

func setupAppRoutes(app *fiber.App, appHandler *handlers.AppHandler, uploadHandler *handlers.UploadHandler, volunteerHandler *handlers.VolunteerHandler, activityHandler *handlers.ActivityHandler, announcementHandler *handlers.AnnouncementHandler, contactHandler *handlers.ContactHandler, organizationHandler *handlers.OrganizationHandler, onboardingHandler *handlers.OnboardingHandler, staticHandler *handlers.StaticHandler, userAdminHandler *handlers.UserAdminHandler, complaintHandler *handlers.ComplaintHandler, surveyHandler *handlers.SurveyHandler, educationHandler *handlers.EducationHandler, store *session.Store, csrfMiddleware *middlewares.CSRFMiddleware) {
	// Protected routes (semua di root path, bukan /app/* lagi)
	// Apply AuthRequired globally — all routes below require auth
	app.Use(middlewares.AuthRequired(store))

	// CSRF for state-changing methods (skip GET/HEAD/OPTIONS)
	app.Use(csrfMiddleware.Protect())

	// Onboarding — new users (relawan) without volunteer record
	app.Get("/onboarding", onboardingHandler.Show)
	app.Post("/onboarding", onboardingHandler.Store)

	// Profile — mutation only (GET is public)
	app.Put("/profile", appHandler.UpdateProfile)
	app.Put("/profile/password", appHandler.UpdatePassword)

	// Organization profile — edit routes only (GET is public)
	app.Put("/profil", middlewares.AdminRequired(store), organizationHandler.Update)
	app.Post("/profil", middlewares.AdminRequired(store), organizationHandler.Update)

	// Kegiatan — create/edit only (list/detail is public)
	app.Get("/kegiatan/create", activityHandler.Create)
	app.Post("/kegiatan", middlewares.AdminRequired(store), activityHandler.Store)
	app.Get("/kegiatan/:id/edit", activityHandler.Edit)
	app.Put("/kegiatan/:id", middlewares.AdminRequired(store), activityHandler.Update)
	app.Delete("/kegiatan/:id", middlewares.AdminRequired(store), activityHandler.Destroy)

	// Berita — CRUD for admin only (GET index & show are public)
	app.Get("/berita/create", announcementHandler.Create)
	app.Post("/berita", middlewares.AdminRequired(store), announcementHandler.Store)
	app.Get("/berita/:id/edit", announcementHandler.Edit)
	app.Put("/berita/:id", middlewares.AdminRequired(store), announcementHandler.Update)
	app.Delete("/berita/:id", middlewares.AdminRequired(store), announcementHandler.Destroy)

	// Kontak — CRUD for admin only (GET index is public)
	app.Get("/kontak/create", contactHandler.Create)
	app.Post("/kontak", middlewares.AdminRequired(store), contactHandler.Store)
	app.Get("/kontak/:id/edit", contactHandler.Edit)
	app.Put("/kontak/:id", middlewares.AdminRequired(store), contactHandler.Update)
	app.Delete("/kontak/:id", middlewares.AdminRequired(store), contactHandler.Destroy)

	// Relawan — create/edit only (list/detail is public)
	app.Get("/relawan/create", volunteerHandler.Create)
	app.Post("/relawan", middlewares.AdminRequired(store), volunteerHandler.Store)
	app.Get("/relawan/:id/edit", volunteerHandler.Edit)
	app.Put("/relawan/:id", middlewares.AdminRequired(store), volunteerHandler.Update)
	app.Delete("/relawan/:id", middlewares.AdminRequired(store), volunteerHandler.Destroy)

	// Pengaduan — admin manage (public submit is in setupPublicFormRoutes)	// Pengaduan — admin manage (public submit is in setupPublicFormRoutes)
	app.Put("/pengaduan/:id", middlewares.AdminRequired(store), complaintHandler.UpdateStatus)
	app.Delete("/pengaduan/:id", middlewares.AdminRequired(store), complaintHandler.Destroy)

	// Upload — admin only (file uploads are sensitive)
	app.Post("/upload", middlewares.AdminRequired(store), uploadHandler.Upload)

	// Education LMS — authenticated users
	app.Get("/sertifikat-saya", educationHandler.MyCertificates)

	// User management — admin only (Iterasi 4C)
	app.Get("/admin/users", middlewares.AdminRequired(store), userAdminHandler.Index)
	app.Get("/admin/users/create", middlewares.AdminRequired(store), userAdminHandler.Create)
	app.Post("/admin/users", middlewares.AdminRequired(store), userAdminHandler.Store)
	app.Get("/admin/users/:id/edit", middlewares.AdminRequired(store), userAdminHandler.Edit)
	app.Put("/admin/users/:id/role", middlewares.AdminRequired(store), userAdminHandler.UpdateRole)
	app.Post("/admin/users/:id/toggle-active", middlewares.AdminRequired(store), userAdminHandler.ToggleActive)
	app.Delete("/admin/users/:id", middlewares.AdminRequired(store), userAdminHandler.Destroy)

}

// SetupCSRFMiddleware sets up the CSRF middleware
func SetupCSRFMiddleware(store *session.Store, secret string) *middlewares.CSRFMiddleware {
	config := middlewares.DefaultCSRFConfig(secret)
	config.Secure = false // Set to true in production with HTTPS
	config.SameSite = "Lax"
	return middlewares.NewCSRFMiddleware(store, config)
}

// SetupMailerService sets up the mailer service
func SetupMailerService(smtpHost string, smtpPort int, smtpUser, smtpPass, fromEmail, fromName string) *services.MailerService {
	return services.NewMailerService(smtpHost, smtpPort, smtpUser, smtpPass, fromEmail, fromName)
}

// SetupPasswordResetHandler sets up the password reset handler
func SetupPasswordResetHandler(
	mailerService *services.MailerService,
	userService *services.UserService,
	store *session.Store,
	inertiaService *services.InertiaService,
	appURL string,
) *handlers.PasswordResetHandler {
	return handlers.NewPasswordResetHandler(
		mailerService,
		userService,
		store,
		inertiaService,
		appURL,
	)
}

// GetAppURL returns the application URL based on environment
func GetAppURL(appPort string, appEnv string) string {
	if appEnv == "production" {
		return "https://yourdomain.com"
	}
	return fmt.Sprintf("http://localhost:%s", appPort)
}
