package handler

import (
	"github.com/GoTurkiye/cargo-service/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func FindByCustomerID(repo repo.PostgresRepo) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("cargo_id")
		uid, err := uuid.Parse(id)
		if err != nil {
			// TO DO error handling
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}
		page := ctx.QueryInt("page", 0)
		limit := ctx.QueryInt("limit", 10)
		cargo, err := repo.FindByCustomerID(uid, page, limit)
		if err != nil {
			// TO DO error handling
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}
		return ctx.JSON(cargo)
	}
}
