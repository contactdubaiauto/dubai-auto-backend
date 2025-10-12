package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TokenGuard(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if authorization == "" {
		return c.Status(http.StatusUnauthorized).JSON(ErrorResponse{Message: "not found any token there!"})
	}

	bearer := strings.Split(authorization, "Bearer ")

	if len(bearer) < 2 {
		return c.Status(http.StatusUnauthorized).JSON(ErrorResponse{Message: "not found any token there!"})
	}

	token := bearer[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		token, claims,
		func(t *jwt.Token) (any, error) {
			return []byte(ENV.ACCESS_KEY), nil
		},
	)

	if err != nil {
		log.Println("Error:", err.Error())
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: err.Error()})
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
		func(t *jwt.Token) (any, error) {
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
	role, ok := c.Locals("role_id").(float64)

	if !ok || role != 100 {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

func DealerGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(float64)
	if !ok || role != 2 {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

func LogistGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(float64)
	if !ok || role != 3 {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

func BrokerGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(float64)
	if !ok || role != 4 {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

func CarServiceGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(float64)
	if !ok || role != 5 {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}
