package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"qolibaba/internal/bank/port"
	"qolibaba/pkg/adapter/storage/types"
)

type BankHandler struct {
	bankService port.Service
}

func NewBankHandler(bankService port.Service) *BankHandler {
	return &BankHandler{
		bankService: bankService,
	}
}

func (h *BankHandler) CreateWallet(c *fiber.Ctx) error {
	var wallet types.Wallet
	if err := c.BodyParser(&wallet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("error parsing request body: %v", err),
		})
	}
	newWallet, err := h.bankService.CreateWallet(&wallet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("error creating wallet: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newWallet)
}
