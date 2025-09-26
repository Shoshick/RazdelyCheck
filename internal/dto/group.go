package dto

import "github.com/google/uuid"

type Group struct {
	ID uuid.UUID `db:"id" json:"id"`
}

type GroupToCheck struct {
	ID      uuid.UUID `db:"id" json:"id"`
	GroupID uuid.UUID `db:"group_id" json:"groupId"`
	CheckID uuid.UUID `db:"check_id" json:"checkId"`
}
