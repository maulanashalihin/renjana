package middlewares

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/session"
)

// AuthRequired is a middleware that checks if the user is authenticated
func AuthRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.Info("checking auth", "path", c.Path())

		sess, err := store.Get(c)
		if err != nil {
			slog.Error("auth session error", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get session",
			})
		}

		userID := sess.Get("user_id")
		slog.Info("auth user id", "user_id", userID)

		// Check both nil AND zero value (int64(0) is non-nil in Go)
		isAuthed := false
		if userID != nil {
			if id, ok := userID.(int64); ok && id != 0 {
				isAuthed = true
			}
		}

		if !isAuthed {
			slog.Warn("not authenticated, redirecting to login")
			// For Inertia requests, return 409 Conflict with X-Inertia-Location header
			// This tells the Inertia client to do a full page redirect to /login
			if c.Get("X-Inertia") == "true" {
				c.Set("X-Inertia-Location", "/login")
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": "Please login to continue",
				})
			}
			return c.Redirect("/login")
		}

		// Store user info in locals for handlers to use
		c.Locals("user_id", userID)
		c.Locals("email", sess.Get("email"))
		c.Locals("role", sess.Get("role"))
		slog.Info("auth successful", "user_id", userID)

		return c.Next()
	}
}

// AdminRequired is a middleware that checks if the user is an admin
func AdminRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get session",
			})
		}

		userID := sess.Get("user_id")
		if userID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authenticated",
			})
		}

		role := sess.Get("role")
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Admin access required",
			})
		}

		c.Locals("user_id", userID)
		c.Locals("email", sess.Get("email"))
		c.Locals("role", role)

		return c.Next()
	}
}

// Guest is a middleware that redirects authenticated users away from login/register pages
func Guest(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Next()
		}

		userID := sess.Get("user_id")
		slog.Info("guest check", "user_id", userID)

		// Check both nil AND zero value (int64(0) is non-nil in Go)
		if userID != nil {
			if id, ok := userID.(int64); ok && id != 0 {
				slog.Info("guest already authenticated, redirecting")
				return c.Redirect("/")
			}
		}

		return c.Next()
	}
}

// KoordinatorRequired allows koordinator or admin to pass.
// Relawan (regular volunteers) are blocked.
func KoordinatorRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get session",
			})
		}

		userID := sess.Get("user_id")
		if userID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authenticated",
			})
		}

		role := sess.Get("role")
		if role != "koordinator" && role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Koordinator or admin access required",
			})
		}

		c.Locals("user_id", userID)
		c.Locals("email", sess.Get("email"))
		c.Locals("role", role)
		c.Locals("district_id", sess.Get("district_id"))

		return c.Next()
	}
}

// RelawanRequired allows any authenticated user (relawan, koordinator, admin).
func RelawanRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get session",
			})
		}

		userID := sess.Get("user_id")
		if userID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authenticated",
			})
		}

		role := sess.Get("role")
		if role == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Authenticated user with valid role required",
			})
		}

		c.Locals("user_id", userID)
		c.Locals("email", sess.Get("email"))
		c.Locals("role", role)
		c.Locals("district_id", sess.Get("district_id"))
		c.Locals("volunteer_id", sess.Get("volunteer_id"))

		return c.Next()
	}
}

// ScopeDistrict ensures the user has a district scope (koordinator or admin).
// For koordinator, only allow access to resources in their assigned district.
func ScopeDistrict(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Next()
		}

		role := sess.Get("role")
		districtID := sess.Get("district_id")

		// Admin bypass district scope
		if role == "admin" {
			return c.Next()
		}

		// Koord without district_id is invalid
		if role == "koordinator" && districtID == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Koordinator must be assigned to a district",
			})
		}

		// Stash district_id for handlers to use
		if districtID != nil {
			c.Locals("scope_district_id", districtID)
		}

		return c.Next()
	}
}

// Logger creates a simple logger middleware
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Simple logging - in production, use a proper logger
		c.Next()
		return nil
	}
}
