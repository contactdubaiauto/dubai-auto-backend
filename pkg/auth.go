package pkg

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func TokenGuard(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "not found any token there!"})
	}

	bearer := strings.Split(authorization, "Bearer ")
	if len(bearer) < 2 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "not found any token there!"})
	}

	token := bearer[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(ENV.ACCESS_KEY), nil
		},
	)
	if err != nil {
		log.Println("Error:", err.Error())
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
	}

	c.Locals("id", int(claims["id"].(float64)))
	c.Locals("role_id", claims["role_id"].(float64))
	return c.Next()
}

func UserGuardOrDefault(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	bearer := strings.Split(authorization, "Bearer ")
	if len(bearer) < 2 {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	token := bearer[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(ENV.ACCESS_KEY), nil
		},
	)
	if err != nil {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	c.Locals("id", int(claims["id"].(float64)))
	c.Locals("role_id", claims["role_id"].(float64))
	return c.Next()
}

func AdminGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role").(string)
	if !ok || role != "admin" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}
