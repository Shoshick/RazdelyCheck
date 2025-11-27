package dto

import "github.com/google/uuid"

type Group struct {
	ID uuid.UUID `db:"id" json:"id"`
}
