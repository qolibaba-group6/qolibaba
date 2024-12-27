package http

import (
	"qolibaba/api/handlers/grpc"
	"qolibaba/api/pb"
	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
)

func CreateTerminal(cfg config.RoutemapServiceConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.TerminalCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		grpcClient := grpc.NewRoutemapGRPCClient(cfg)
		resp, err := grpcClient.CreateTerminal(c.UserContext(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
