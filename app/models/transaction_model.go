package models

import (
	"github.com/google/uuid"
	"time"
)

// Transaction struct to describe transaction object.
type Transaction struct {
	Id              uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Amount          int       `db:"amount" json:"amount" validate:"min=0"`
	Status          string    `db:"status" json:"status" validate:"required,lte=255"`
	TransactionType string    `db:"transaction_type" json:"transaction_type" validate:"lte=255"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	ClientId        uuid.UUID `db:"client_id" json:"client_id" validate:"uuid"`
}
