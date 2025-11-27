package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type CheckRepo interface {
	Create(c *dto.Check) error
	GetByID(id uuid.UUID) (*dto.Check, error)
	ListByUserID(userID uuid.UUID) ([]*dto.Check, error)
	Delete(id uuid.UUID) error
}
