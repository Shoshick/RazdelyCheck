package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"RazdelyCheck/internal/util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type CheckSourceService struct {
	repo         repo.CheckSourceRepo
	db           *sql.DB
	checkService *CheckService
}

func NewCheckSourceService(r repo.CheckSourceRepo, db *sql.DB, cs *CheckService) *CheckSourceService {
	return &CheckSourceService{
		repo:         r,
		db:           db,
		checkService: cs,
	}
}

func (s *CheckSourceService) ProcessQR(userID uuid.UUID, input dto.QRScanInput, jsonData []byte) (*dto.Check, error) {
	if input.QRData == "" {
		return nil, errors.New("QR is empty")
	}

	var check *dto.Check
	var items []dto.Item
	var totalSum int64
	var err error

	err = util.WithTransaction(s.db, func(tx *sql.Tx) error {

		items, totalSum, err = ParseCheckJSON(jsonData, uuid.Nil)
		if err != nil {
			return err
		}

		check, err = s.checkService.Create(userID, nil, totalSum)
		if err != nil {
			return err
		}

		checkSource := dto.CheckSource{
			CheckID: check.ID,
			QR:      input.QRData,
		}
		if err := s.repo.CreateTx(tx, &checkSource); err != nil {
			return err
		}

		for i := range items {
			items[i].CheckID = check.ID
			if err := s.repo.CreateItemTx(tx, &items[i]); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return check, nil
}

func ParseCheckJSON(jsonData []byte, checkID uuid.UUID) ([]dto.Item, int64, error) {
	var raw dto.CheckResponse
	if err := json.Unmarshal(jsonData, &raw); err != nil {
		return nil, 0, err
	}

	if raw.Code != 1 {
		return nil, 0, fmt.Errorf("API returned code %d", raw.Code)
	}

	items := make([]dto.Item, len(raw.Data.JSON.Items))
	for i, r := range raw.Data.JSON.Items {
		items[i] = dto.Item{
			ID:       uuid.New(),
			CheckID:  checkID,
			Position: i + 1,
			Name:     r.Name,
			Price:    int64(r.Price),
			Quantity: r.Quantity,
		}
	}

	totalSum := int64(raw.Data.JSON.TotalSum)

	return items, totalSum, nil
}
