package controllers

import (
	dto "backend/internal/dtos"
	"backend/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateMatch(ctx *fiber.Ctx) error {
	var req dto.CreateMatchRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request data",
		})
	}
	hostID := ctx.Locals("player_id").(string)
	res, err := services.CreateMatch(req, hostID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func JoinMatch(ctx *fiber.Ctx) error {
	var req dto.JoinMatchRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request data",
		})
	}
	clientID := ctx.Locals("player_id").(string)
	matchID := req.MatchID

}
