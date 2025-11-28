package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type CheckRepo interface {
	Create(c *dto.Check) error
	GetByID(id uuid.UUID) (*dto.Check, error)
	GetItemsFromCheck(userID, checkID uuid.UUID) ([]*dto.Item, error)
	UpdateTotalSum(id uuid.UUID, totalSum int64) error
	ListByUserID(id uuid.UUID) ([]*dto.Check, error)
	Delete(id uuid.UUID) error
}
