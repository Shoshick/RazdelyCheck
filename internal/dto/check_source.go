package dto

import "github.com/google/uuid"

type CheckSource struct {
	ID      uuid.UUID `db:"id" json:"id"`
	CheckID uuid.UUID `db:"check_id" json:"checkId"`
	Data    string    `db:"data" json:"data"`
}
