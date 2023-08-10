package main

import (
	"github.com/GoTurkiye/cargo-service/cargo"
	"github.com/GoTurkiye/cargo-service/handler"
	"github.com/GoTurkiye/cargo-service/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// TODO: Get it from config, environment variables or yaml whatever
	dsn := "host=postgres user=cargo password=password dbname=cargo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	db.AutoMigrate(&cargo.Cargo{})

	repo := repo.NewPostgresRepo(db)
	app := fiber.New()

	app.Post("/cargo", handler.Create(repo))

	// TODO: Config...
	app.Listen(":8080")
}
