package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CheckRepo interface {
	Create(c *dto.Check) error
	GetByID(id uuid.UUID) (*dto.Check, error)
	GetCheckByGroupID(id uuid.UUID) (uuid.UUID, error)

	UpdateTotalSum(id uuid.UUID) error
	ListByUserID(id uuid.UUID) ([]*dto.Check, error)

	UpdateTotalSumTx(tx *sqlx.Tx, checkID uuid.UUID) error
	GetCheckByIDTx(tx *sqlx.Tx, id uuid.UUID) (*dto.Check, error)

	Delete(id uuid.UUID) error
}
