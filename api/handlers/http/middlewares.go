package http

import (
	"fmt"
	"qolibaba/pkg/context"
	"qolibaba/pkg/jwt"
	"qolibaba/pkg/logger"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userClaims := userClaims(ctx)
			if userClaims == nil {
				return fiber.ErrUnauthorized
			}

			logger := context.GetLogger(ctx.UserContext())
			context.SetLogger(ctx.UserContext(), logger.With("user_id", userClaims.UserID))

			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
		AuthScheme: "Bearer",
	})
}

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext(), context.WithLogger(logger.NewLogger())))
	return c.Next()
}


func rolesAccessMiddleware(requiredRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := userClaims(c)
		if claims == nil {
			return fiber.ErrUnauthorized
		}

		if !contains(requiredRoles, claims.Role) {
			return fiber.NewError(fiber.StatusForbidden, fmt.Sprintf(
				"Access denied: user role '%s' is not authorized. Required roles: %v",
				claims.Role,
				requiredRoles,
			))
		}

		return c.Next()
	}
}
