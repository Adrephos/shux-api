package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/auth"
	"golang.org/x/crypto/bcrypt"
)

func getBearerToken(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	authHeader := strings.Split(headers["Authorization"], " ")
	if len(authHeader) < 2 {
		return "", fiber.ErrUnauthorized
	}
	if authHeader[0] != "Bearer" {
		return "", fiber.ErrUnauthorized
	}
	return authHeader[1], nil
}

func VerifyToken(c *fiber.Ctx) error {
	token, err := getBearerToken(c)
	if err != nil {
		return err
	}

	_, err = auth.VerifyToken(token, false)
	if err != nil {
		return err
	}

	return c.Next()
}

func RegisterKey(c *fiber.Ctx) error {
	reqKey, err := getBearerToken(c)
	key := os.Getenv("REGISTER_SECRET")
	if err != nil {
		return err
	}

	// Encrypt using HS256
	bcrypt.CompareHashAndPassword([]byte(reqKey), []byte(key))
	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
