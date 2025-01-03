package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"qolibaba/internal/bank/port"
	"qolibaba/pkg/adapter/storage/types"
	"strconv"
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

// ChargeWalletHandler handles the HTTP request to charge a wallet
func (h *BankHandler) ChargeWalletHandler(c *fiber.Ctx) error {

	var request struct {
		WalletID uint    `json:"wallet_id" validate:"required"`
		Amount   float64 `json:"amount" validate:"required,gt=0"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	updatedWallet, transaction, err := h.bankService.ChargeWallet(request.WalletID, request.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error charging wallet: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"wallet":      updatedWallet,
		"transaction": transaction,
	})
}

func (h *BankHandler) ProcessUnconfirmedClaim(c *fiber.Ctx) error {
	var claimRequest types.Claim
	if err := c.BodyParser(&claimRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}
	processedClaim, err := h.bankService.ProcessUnconfirmedClaim(&claimRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Claim processed successfully",
		"claim":   processedClaim,
	})
}

func (h *BankHandler) ProcessConfirmedClaimHandler(c *fiber.Ctx) error {
	claimID := c.Params("claim_id")
	if claimID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Claim ID is required",
		})
	}

	parsedClaimID, err := strconv.Atoi(claimID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid claim ID format",
		})
	}

	claim, err := h.bankService.ProcessConfirmedClaim(uint(parsedClaimID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error processing claim: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(claim)
}
