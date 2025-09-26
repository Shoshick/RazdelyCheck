package repo

import (
	"RazdelyCheck/internal/dto"
	"github.com/google/uuid"
)

type CheckResultRepo interface {
	ListByCheckID(checkID uuid.UUID) ([]*dto.CheckResult, error)
	ListByUserID(userID uuid.UUID) ([]*dto.CheckResult, error)
}
