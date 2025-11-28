package repo

import (
	"RazdelyCheck/internal/dto"
	"database/sql"
	"github.com/google/uuid"
)

type CheckResultRepo interface {
	CreateCheckResultTx(tx *sql.Tx, cr *dto.CheckResult) error
	DeleteCheckResult(id uuid.UUID) error
	GetCheckResultsByCheckID(checkID uuid.UUID) ([]dto.CheckResult, error)

	AddItemToCheckResult(item *dto.ItemToCheckResult) error
	UpdateItemQuantityInCheckResult(itemID, checkResultID uuid.UUID, quantity float64) error
	RemoveItemFromCheckResult(itemID, checkResultID uuid.UUID) error
	GetItemsByCheckResultID(checkResultID uuid.UUID) ([]dto.ItemToCheckResult, error)

	UpdateCheckResultTotal(crID uuid.UUID) error
}
