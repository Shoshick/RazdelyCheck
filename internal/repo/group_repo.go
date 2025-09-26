package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type GroupRepo interface {
	Create(g *dto.Group) error
	GetByID(id uuid.UUID) (*dto.Group, error)
	List() ([]*dto.Group, error)
	Update(g *dto.Group) error
	Delete(id uuid.UUID) error
}
