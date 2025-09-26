package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type UserRepo interface {
	Create(u *dto.User) error
	GetByID(id uuid.UUID) (*dto.User, error)
	List() ([]*dto.User, error)
	Update(u *dto.User) error
	Delete(id uuid.UUID) error
}
