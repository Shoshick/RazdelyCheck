package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CheckSourceRepo interface {
	Create(cs *dto.CheckSource) error
	GetByCheckID(checkID uuid.UUID) (*dto.CheckSource, error)
	CreateTx(tx *sqlx.Tx, cs *dto.CheckSource) error
	CreateItemTx(tx *sqlx.Tx, item *dto.Item) error
}
