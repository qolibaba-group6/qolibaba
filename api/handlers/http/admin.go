package http

import (
	"qolibaba/api/handlers/grpc"
	"qolibaba/api/pb"
	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
)

func SayHello(cfg config.AdminServiceConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.AdminSayHelloRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		grpcClient := grpc.NewAdminGRPCClient(cfg)
		resp, err := grpcClient.SayHello(c.UserContext(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
