package dto

import (
	"github.com/google/uuid"
)

type CheckResult struct {
	ID       uuid.UUID `db:"id" json:"id"`
	CheckID  uuid.UUID `db:"check_id" json:"checkId"`
	UserID   uuid.UUID `db:"user_id" json:"userId"`
	TotalDue float64   `db:"total_due" json:"totalDue"`
}

type CheckResultWithItems struct {
	TotalSum int64               `json:"totalSum"`
	Items    []ItemToCheckResult `json:"items"`
}
