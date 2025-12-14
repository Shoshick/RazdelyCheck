package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	Create(i *dto.Item) error
	GetByID(id uuid.UUID) (*dto.Item, error)
	ListByCheckID(checkID uuid.UUID) ([]*dto.Item, error)
	Update(i *dto.Item) error

	ExcludeItemTx(tx *sqlx.Tx, id uuid.UUID) error
	IncludeItemTx(tx *sqlx.Tx, id uuid.UUID) error
	GetItemsByCheckIDTx(tx *sqlx.Tx, checkID uuid.UUID) ([]*dto.Item, error)

	Delete(id uuid.UUID) error
}
