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

type CheckItem struct {
	Name     string  `json:"name"`
	Price    int64   `json:"price"`
	Quantity float64 `json:"quantity"`
	Sum      int64   `json:"sum"`
	Nds18    int64   `json:"nds18,omitempty"`
	Nds10    int64   `json:"nds,omitempty"`
	Nds0     int64   `json:"nds0,omitempty"`
	NdsNo    int64   `json:"ndsNo,omitempty"`
}

type CheckJSONData struct {
	User       string      `json:"user"`
	Items      []CheckItem `json:"items"`
	TicketDate string      `json:"ticketDate"`
	TotalSum   int64       `json:"totalSum"`
}

type CheckResponse struct {
	Code  int `json:"code"`
	First int `json:"first"`
	Data  struct {
		JSON CheckJSONData `json:"json"`
	} `json:"data"`
}
