package dto

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID `db:"id" json:"id"`
	CheckID  uuid.UUID `db:"check_id" json:"checkId"`
	Position int       `db:"position" json:"position"`
	Name     string    `db:"name" json:"name"`
	Price    float64   `db:"price" json:"price"`
	Quantity float64   `db:"quantity" json:"quantity"`
}

type ItemToUser struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ItemID   uuid.UUID `db:"item_id" json:"itemId"`
	UserID   uuid.UUID `db:"user_id" json:"userId"`
	Quantity float64   `db:"quantity" json:"quantity"`
}
