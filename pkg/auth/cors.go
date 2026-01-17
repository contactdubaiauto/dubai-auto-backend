package auth

import "github.com/gofiber/fiber/v2"

func Cors(c *fiber.Ctx) error {
	// Get the origin from the request
	origin := c.Get("Origin")

	// Set CORS headers
	// Never allow wildcard with credentials - this is a security vulnerability
	if origin != "" {
		c.Set("Access-Control-Allow-Origin", origin)
		c.Set("Access-Control-Allow-Credentials", "true")
	} else {
		// No origin specified - don't allow credentials with wildcard
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "false")
	}
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	c.Set("Access-Control-Max-Age", "86400") // 24 hours

	// Handle preflight requests
	if c.Method() == fiber.MethodOptions {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Next()
}
