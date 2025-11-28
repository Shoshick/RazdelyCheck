package service

import (
	"database/sql"

	"RazdelyCheck/internal/repo"
	"RazdelyCheck/internal/util"

	"github.com/google/uuid"
)

type ItemService struct {
	itemRepo  repo.ItemRepo
	checkRepo repo.CheckRepo
	db        *sql.DB
}

func NewItemService(itemRepo repo.ItemRepo, checkRepo repo.CheckRepo, db *sql.DB) *ItemService {
	return &ItemService{
		itemRepo:  itemRepo,
		checkRepo: checkRepo,
		db:        db,
	}
}

func (s *ItemService) ExcludeItem(itemID, checkID uuid.UUID) error {
	return util.WithTransaction(s.db, func(tx *sql.Tx) error {
		if err := s.itemRepo.ExcludeItemTx(tx, itemID); err != nil {
			return err
		}
		return s.checkRepo.UpdateTotalSumTx(tx, checkID)
	})
}

func (s *ItemService) IncludeItem(itemID, checkID uuid.UUID) error {
	return util.WithTransaction(s.db, func(tx *sql.Tx) error {
		if err := s.itemRepo.IncludeItemTx(tx, itemID); err != nil {
			return err
		}
		return s.checkRepo.UpdateTotalSumTx(tx, checkID)
	})
}
