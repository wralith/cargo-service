package cargo

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	OrderArrived Status = "OrderArrived"
	Preparing    Status = "Preparing"
	OnTheWay     Status = "OnTheWay"
	Delivered    Status = "Delivered"
)

type Cargo struct {
	ID         uuid.UUID
	Status     Status
	CustomerID uuid.UUID
	Weight     float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewCargo(CustomerID uuid.UUID, Weight float64) *Cargo {
	return &Cargo{
		ID:         uuid.New(),
		Status:     OrderArrived,
		CustomerID: CustomerID,
		Weight:     Weight,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
