package dto

import (
	"github.com/google/uuid"
	"time"
)

type Check struct {
	ID        uuid.UUID `db:"id" json:"id"`
	UserID    uuid.UUID `db:"user_id" json:"userId"`
	TotalSum  float64   `db:"total_sum" json:"totalSum"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
