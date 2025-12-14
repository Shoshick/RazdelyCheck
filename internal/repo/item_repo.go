package repo

import (
	"RazdelyCheck/internal/dto"
	"database/sql"
	"github.com/google/uuid"
)

type ItemRepo interface {
	Create(i *dto.Item) error
	GetByID(id uuid.UUID) (*dto.Item, error)
	ListByCheckID(checkID uuid.UUID) ([]*dto.Item, error)
	Update(i *dto.Item) error

	ExcludeItemTx(tx *sql.Tx, id uuid.UUID) error
	IncludeItemTx(tx *sql.Tx, id uuid.UUID) error
	GetItemsByCheckIDTx(tx *sql.Tx, checkID uuid.UUID) ([]*dto.Item, error)

	Delete(id uuid.UUID) error
}
