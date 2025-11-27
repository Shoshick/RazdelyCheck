package dto

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID `db:"id"`
	CheckID  uuid.UUID `db:"check_id"`
	Position int       `db:"position"`
	Name     string    `db:"name"`
	Price    int64     `db:"price"`
	Quantity float64   `db:"quantity"`
}

type ItemToUser struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ItemID   uuid.UUID `db:"item_id" json:"itemId"`
	UserID   uuid.UUID `db:"user_id" json:"userId"`
	Quantity float64   `db:"quantity" json:"quantity"`
}
