package http

import (
	"qolibaba/api/handlers/grpc"
	"qolibaba/api/pb"
	// "qolibaba/api/service"

	"github.com/gofiber/fiber/v2"
)

func SayHello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.AdminSayHelloRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := grpc.SayHelloClient(c.UserContext(), &req)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

// func SayHello(svc *service.AdminService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var req pb.AdminSayHelloRequest
// 		if err := c.BodyParser(&req); err != nil {
// 			return fiber.ErrBadRequest
// 		}

// 		resp, err := svc.SayHello(c.UserContext(), &req)
// 		if err != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 		}

// 		return c.JSON(resp)
// 	}
// }
