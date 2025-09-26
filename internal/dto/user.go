package dto

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID  `db:"id" json:"id"`
	Email   *string    `db:"email" json:"email,omitempty"`
	Name    *string    `db:"name" json:"name,omitempty"`
	OwnerID *uuid.UUID `db:"owner" json:"owner,omitempty"` // self-reference
}

type UserToGroup struct {
	ID      uuid.UUID `db:"id" json:"id"`
	UserID  uuid.UUID `db:"user_id" json:"userId"`
	GroupID uuid.UUID `db:"group_id" json:"groupId"`
}
