package dto

import "github.com/google/uuid"

type Item struct {
	ID         uuid.UUID `db:"id"`
	CheckID    uuid.UUID `db:"check_id"`
	Position   int       `db:"position"`
	Name       string    `db:"name"`
	Price      int64     `db:"price"`
	Quantity   float64   `db:"quantity"`
	IsExcluded bool      `db:"is_excluded"`
}

type ItemToCheckResult struct {
	ID            uuid.UUID `db:"id" json:"id"`
	ItemID        uuid.UUID `db:"item_id" json:"itemId"`
	CheckResultID uuid.UUID `db:"check_result_id" json:"checkResultId"`
	Quantity      float64   `db:"quantity" json:"quantity"`
}
