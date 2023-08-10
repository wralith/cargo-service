package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/GoTurkiye/cargo-service/cargo"
	"github.com/GoTurkiye/cargo-service/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/assert"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateHandler(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)
	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)
	db, err := gorm.Open(pg.Open(connStr), &gorm.Config{})
	assert.NoError(t, err)
	// db has the required schema and data, and is ready to use
	db.AutoMigrate(&cargo.Cargo{})

	repo := repo.NewPostgresRepo(db)
	app := fiber.New()
	app.Post("/cargo", Create(repo))
	dto := CreateInput{
		CustomerID: uuid.New(),
		Weight:     100.0,
	}

	mdto, err := json.Marshal(dto)
	assert.NoError(t, err)
	req := httptest.NewRequest("POST", "/cargo", bytes.NewReader(mdto))
	req.Header.Add("content-type", "application/json")
	response, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 201, response.StatusCode)
}
