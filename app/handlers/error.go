package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// HandleClientError logs browser-side errors reported by the frontend.
// This is a simple POST endpoint used by the global error handler in
// error-handler.ts. No authentication needed — error reports are low-risk.
func HandleClientError(c *fiber.Ctx) error {
	var body map[string]any
	if err := c.BodyParser(&body); err != nil {
		slog.Warn("client error: failed to parse body", "err", err)
		return c.Status(400).JSON(fiber.Map{"ok": false})
	}

	slog.Warn("client error",
		"type", body["type"],
		"message", body["message"],
		"url", body["url"],
		"userAgent", body["userAgent"],
	)

	return c.JSON(fiber.Map{"ok": true})
}
