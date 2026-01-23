package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserRole int

const (
	ADMIN_ROLE       = 100
	USER_ROLE        = 1
	DEALER_ROLE      = 2
	LOGIST_ROLE      = 3
	BROKER_ROLE      = 4
	CAR_SERVICE_ROLE = 5
	ROLE_COUNT       = 5
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

	role_id := claims["role_id"]
	id := claims["id"]

	if id == nil || role_id == nil {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!"})
	}

	// Safe type assertion with error handling
	idFloat, ok := id.(float64)
	if !ok {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid! Invalid ID type"})
	}

	roleFloat, ok := role_id.(float64)
	if !ok {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid! Invalid role_id type"})
	}

	c.Locals("id", int(idFloat))
	c.Locals("role_id", int(roleFloat))
	return c.Next()
}

// Language checker middleware: sets "lang" in fiber.Locals based on Accept-Language or X-Language header
func LanguageChecker(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language")

	switch lang {
	case "ru":
		c.Locals("lang", "name_ru")
	case "ae":
		c.Locals("lang", "name_ae")
	default:
		c.Locals("lang", "name")
	}
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

	// Safe type assertion with error handling
	id := claims["id"]
	role_id := claims["role_id"]

	if id == nil || role_id == nil {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	idFloat, ok := id.(float64)
	if !ok {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	roleFloat, ok := role_id.(float64)
	if !ok {
		c.Locals("id", 0)
		c.Locals("role_id", 0)
		return c.Next()
	}

	c.Locals("id", int(idFloat))
	c.Locals("role_id", int(roleFloat))
	return c.Next()
}

func AdminGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)

	if !ok || role != ADMIN_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!xrt"})
	}

	return c.Next()
}

func ThirdPartyGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)
	// 0 admin, 1 user, 2 dealer, 3 logistic, 4 broker, 5 car service
	if !ok || role == USER_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!duio"})
	}
	return c.Next()
}

func DealerGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)

	if !ok || role != DEALER_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!duio"})
	}
	return c.Next()
}

func LogistGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)

	if !ok || role != LOGIST_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!ppol"})
	}
	return c.Next()
}

func BrokerGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)

	if !ok || role != BROKER_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!mmdt"})
	}
	return c.Next()
}

func CarServiceGuard(c *fiber.Ctx) error {
	role, ok := c.Locals("role_id").(int)

	if !ok || role != CAR_SERVICE_ROLE {
		return c.Status(http.StatusForbidden).JSON(ErrorResponse{Message: "token is invalid!dxre"})
	}
	return c.Next()
}
