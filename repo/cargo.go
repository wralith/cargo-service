package repo

import (
	"github.com/GoTurkiye/cargo-service/cargo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo interface {
	Create(*cargo.Cargo) (uuid.UUID, error)
}

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) Create(c *cargo.Cargo) (uuid.UUID, error) {
	err := r.db.Create(&c).Error
	return c.ID, err
}
