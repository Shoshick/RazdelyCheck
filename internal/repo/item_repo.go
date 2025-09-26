package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type ItemRepo interface {
	Create(i *dto.Item) error
	GetByID(id uuid.UUID) (*dto.Item, error)
	ListByCheckID(checkID uuid.UUID) ([]*dto.Item, error)
	Update(i *dto.Item) error
	Delete(id uuid.UUID) error
}
