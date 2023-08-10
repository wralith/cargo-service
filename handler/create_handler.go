package handler

import (
	"github.com/GoTurkiye/cargo-service/cargo"
	"github.com/GoTurkiye/cargo-service/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateInput struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Weight     float64   `json:"weight"`
}

func Create(repo repo.Repo) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input CreateInput
		err := c.BodyParser(&input)
		// TODO: Add validator
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		crg := cargo.NewCargo(input.CustomerID, input.Weight)
		id, err := repo.Create(crg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": id,
		})
	}
}
