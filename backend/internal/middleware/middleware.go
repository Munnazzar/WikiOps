package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ensures every user has a session ID before hitting the controller
func SessionGuard(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		sessionID = "guest_" + uuid.New().String()

		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			HTTPOnly: true, // Prevents JavaScript from reading the cookie (Security!)
			SameSite: "Lax",
		})
	}
	// storing session id for later use in control
	c.Locals("player_id", sessionID)
	return c.Next()
}
