package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type GroupRepo interface {
	Create(g *dto.Group) error
	GetByID(id uuid.UUID) (*dto.Group, error)

	List() ([]*dto.Group, error)
	ListByUser(userID uuid.UUID) ([]*dto.Group, error)
	ListByGroup(groupID uuid.UUID) ([]*dto.User, error)

	Update(g *dto.Group) error
	Delete(id uuid.UUID) error

	ExistsUserInGroup(userID, groupID uuid.UUID) (bool, error)
	AddUserToGroup(userID, groupID uuid.UUID) error
	RemoveUserFromGroup(userID, groupID uuid.UUID) error
}
