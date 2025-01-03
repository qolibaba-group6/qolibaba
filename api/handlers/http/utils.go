package http

import (
	"qolibaba/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*jwt2.Token).Claims.(*jwt.UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}