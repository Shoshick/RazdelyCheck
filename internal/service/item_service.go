package service

import (
	"database/sql"

	"RazdelyCheck/internal/repo"
	"github.com/google/uuid"
)

type ItemService struct {
	itemRepo  repo.ItemRepo
	checkRepo repo.CheckRepo
}

// конструктор
func NewItemService(itemRepo repo.ItemRepo, checkRepo repo.CheckRepo) *ItemService {
	return &ItemService{
		itemRepo:  itemRepo,
		checkRepo: checkRepo,
	}
}

func (s *ItemService) ExcludeItem(tx *sql.Tx, itemID, checkID uuid.UUID) error {
	if err := s.itemRepo.ExcludeItemTx(tx, itemID); err != nil {
		return err
	}
	return s.checkRepo.UpdateTotalSumTx(tx, checkID)
}

func (s *ItemService) IncludeItem(tx *sql.Tx, itemID, checkID uuid.UUID) error {
	if err := s.itemRepo.IncludeItemTx(tx, itemID); err != nil {
		return err
	}
	return s.checkRepo.UpdateTotalSumTx(tx, checkID)
}
