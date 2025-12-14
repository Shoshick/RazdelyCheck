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
	GetTotalSumByCheckResultID(checkResultID uuid.UUID) (int64, error)

	GetUsedQuantitiesByCheckIDTx(tx *sql.Tx, checkID uuid.UUID) (map[uuid.UUID]float64, error)
	AddItemToCheckResultTx(tx *sql.Tx, itemID, checkResultID uuid.UUID, qty float64) error
	UpdateTotalDueTx(tx *sql.Tx, checkResultID uuid.UUID, total float64) error

	UpdateCheckResultTotal(crID uuid.UUID) error
}
