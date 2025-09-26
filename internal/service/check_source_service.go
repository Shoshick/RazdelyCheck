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
	repo repo.CheckSourceRepo
	db   *sql.DB
}

func NewCheckSourceService(r repo.CheckSourceRepo, db *sql.DB) *CheckSourceService {
	return &CheckSourceService{
		repo: r,
		db:   db,
	}
}

// ProcessQR сохраняет QR-код и товары в рамках одной транзакции
func (s *CheckSourceService) ProcessQR(input dto.QRScanInput, checkID uuid.UUID, jsonData []byte) error {
	if input.QRData == "" {
		return errors.New("QR is empty")
	}

	// Разбор JSON в Items
	items, err := ParseCheckJSON(jsonData, checkID)
	if err != nil {
		return err
	}

	checkSource := dto.CheckSource{
		CheckID: checkID,
		QR:      input.QRData,
	}

	// Атомарное сохранение через транзакцию
	return util.WithTransaction(s.db, func(tx *sql.Tx) error {
		if err := s.repo.CreateTx(tx, &checkSource); err != nil {
			return err
		}

		for _, item := range items {
			if err := s.repo.CreateItemTx(tx, &item); err != nil {
				return err
			}
		}

		return nil
	})
}

// ParseCheckJSON разбирает JSON и возвращает список товаров с привязкой к checkID
func ParseCheckJSON(jsonData []byte, checkID uuid.UUID) ([]dto.Item, error) {
	var resp dto.CheckResponse
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 1 {
		return nil, fmt.Errorf("check not ready, code: %d", resp.Code)
	}

	items := make([]dto.Item, len(resp.Data.JSON.Items))
	for i, it := range resp.Data.JSON.Items {
		items[i] = dto.Item{
			ID:       uuid.New(),
			CheckID:  checkID,
			Position: i + 1,
			Name:     it.Name,
			Price:    float64(it.Price) / 100,
			Quantity: it.Quantity,
		}
	}

	return items, nil
}
