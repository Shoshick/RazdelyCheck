package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CheckResultService struct {
	repo repo.CheckResultRepo
}

func NewCheckResultService(repo repo.CheckResultRepo) *CheckResultService {
	return &CheckResultService{repo: repo}
}

// Создает миничек для владельца из неразобранных товаров
func (s *CheckResultService) CreateOwnerMiniCheck(tx *sqlx.Tx, checkID, ownerID uuid.UUID) (*dto.CheckResult, error) {
	results, err := s.repo.GetCheckResultsByCheckID(checkID)
	if err != nil {
		return nil, err
	}

	for _, cr := range results {
		if cr.UserID == ownerID {
			return &cr, nil
		}
	}

	newCR := &dto.CheckResult{
		ID:       uuid.New(),
		CheckID:  checkID,
		UserID:   ownerID,
		TotalDue: 0,
	}

	if err := s.repo.CreateCheckResultTx(tx, newCR); err != nil {
		return nil, err
	}

	usedQty, err := s.repo.GetUsedQuantitiesByCheckIDTx(tx, checkID)
	if err != nil {
		return nil, err
	}

	allItems, err := s.repo.GetItemsByCheckResultID(checkID)
	if err != nil {
		return nil, err
	}

	for _, item := range allItems {
		remaining := item.Quantity - usedQty[item.ItemID]
		if remaining > 0 {
			if err := s.repo.AddItemToCheckResultTx(tx, item.ItemID, newCR.ID, remaining); err != nil {
				return nil, err
			}
		}
	}

	return newCR, nil
}

func (s *CheckResultService) AddItemToCheckResult(tx *sqlx.Tx, checkResultID, itemID uuid.UUID, qty float64) error {
	if err := s.repo.AddItemToCheckResultTx(tx, itemID, checkResultID, qty); err != nil {
		return err
	}
	_, err := s.RecalculateMiniCheckTotal(tx, checkResultID)
	return err
}

func (s *CheckResultService) RemoveItemFromCheckResult(tx *sqlx.Tx, checkResultID, itemID uuid.UUID) error {
	if err := s.repo.RemoveItemFromCheckResult(itemID, checkResultID); err != nil {
		return err
	}
	_, err := s.RecalculateMiniCheckTotal(tx, checkResultID)
	return err
}

func (s *CheckResultService) UpdateItemQuantityInCheckResult(tx *sqlx.Tx, checkResultID, itemID uuid.UUID, qty float64) error {
	if err := s.repo.UpdateItemQuantityInCheckResult(itemID, checkResultID, qty); err != nil {
		return err
	}
	_, err := s.RecalculateMiniCheckTotal(tx, checkResultID)
	return err
}

func (s *CheckResultService) GetCheckItems(checkResultID uuid.UUID) (*dto.CheckResultWithItems, error) {
	crItems, err := s.repo.GetItemsByCheckResultID(checkResultID)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.GetTotalSumByCheckResultID(checkResultID)
	if err != nil {
		return nil, err
	}

	return &dto.CheckResultWithItems{
		TotalSum: total,
		Items:    crItems,
	}, nil
}

func (s *CheckResultService) GetAllMiniChecks(checkID uuid.UUID) ([]dto.CheckResult, error) {
	return s.repo.GetCheckResultsByCheckID(checkID)
}

// Пересчет суммы миничека
func (s *CheckResultService) RecalculateMiniCheckTotal(tx *sqlx.Tx, checkResultID uuid.UUID) (int64, error) {
	total, err := s.repo.GetTotalSumByCheckResultID(checkResultID)
	if err != nil {
		return 0, err
	}

	err = s.repo.UpdateTotalDueTx(tx, checkResultID, float64(total)/100)
	if err != nil {
		return 0, err
	}

	return total, nil
}
