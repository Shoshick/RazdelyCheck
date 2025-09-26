package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type UserRepo interface {
	Create(u *dto.User) error
	GetByID(id uuid.UUID) (*dto.User, error)
	ListByOwner(ownerID uuid.UUID) ([]*dto.User, error)
	ExistsByEmail(email string) (bool, error)

	UpdateName(id uuid.UUID, name string) error
	UpdateEmail(id uuid.UUID, email string) error
	UpdateOwner(id uuid.UUID, ownerID uuid.UUID) error

	Delete(id uuid.UUID) error
}
