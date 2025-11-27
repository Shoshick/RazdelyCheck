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

type CheckResponse struct {
	Code int `json:"code"`
	Data struct {
		JSON struct {
			User     string `json:"user"`
			DateTime string `json:"dateTime"`
			TotalSum int    `json:"totalSum"`
			Items    []struct {
				Name     string  `json:"name"`
				Price    int     `json:"price"`
				Quantity float64 `json:"quantity"`
				Sum      int     `json:"sum"`
			} `json:"items"`
		} `json:"json"`
	} `json:"data"`
}
