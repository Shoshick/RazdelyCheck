package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type CheckSourceRepo interface {
	Create(cs *dto.CheckSource) error
	GetByCheckID(checkID uuid.UUID) (*dto.CheckSource, error)
}
