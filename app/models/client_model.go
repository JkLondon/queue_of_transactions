package models

import (
	"github.com/google/uuid"
)

type Client struct {
	Id      uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Balance int       `db:"balance" json:"balance" validate:"min=0"`
}
