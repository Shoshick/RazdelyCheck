package repo

import (
	"RazdelyCheck/internal/dto"
	"database/sql"
	"github.com/google/uuid"
)

type CheckRepo interface {
	Create(c *dto.Check) error
	GetByID(id uuid.UUID) (*dto.Check, error)
	GetCheckByGroupID(id uuid.UUID) (uuid.UUID, error)

	UpdateTotalSum(id uuid.UUID) error
	ListByUserID(id uuid.UUID) ([]*dto.Check, error)

	UpdateTotalSumTx(tx *sql.Tx, checkID uuid.UUID) error

	Delete(id uuid.UUID) error
}
