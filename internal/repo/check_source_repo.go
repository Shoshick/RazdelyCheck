package repo

import (
	"RazdelyCheck/internal/dto"
	"database/sql"
	"github.com/google/uuid"
)

type CheckSourceRepo interface {
	Create(cs *dto.CheckSource) error
	GetByCheckID(checkID uuid.UUID) (*dto.CheckSource, error)
	CreateTx(tx *sql.Tx, cs *dto.CheckSource) error
	CreateItemTx(tx *sql.Tx, item *dto.Item) error
}
