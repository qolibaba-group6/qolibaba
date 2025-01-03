package http

import (
	"errors"
	"qolibaba/api/pb"
	"qolibaba/api/service"
	"qolibaba/pkg/context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SignUp(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.UserSignUpRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func SingIn(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SingIn(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, service.ErrInvalidUserPassword) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

type UpdateRoleRequest struct {
	Role string `json:"role"`
}

func UpdateRole(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateRoleRequest 
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		userId, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "missing or invalid user id")
		}

		err = svc.UpdateRole(c.UserContext(), userId, req.Role)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

func TestHandler(ctx *fiber.Ctx) error {
	logger := context.GetLogger(ctx.UserContext())

	logger.Info("from test handler", "time", time.Now().Format(time.DateTime))

	return nil
}
