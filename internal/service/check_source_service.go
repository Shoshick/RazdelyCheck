package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type CheckSourceService struct {
	repo repo.CheckSourceRepo
}

func NewCheckSourceService(r repo.CheckSourceRepo) *CheckSourceService {
	return &CheckSourceService{repo: r}
}

// Создание записи QR в БД
func (s *CheckSourceService) ProcessQR(input dto.QRScanInput, checkID uuid.UUID) error {
	if input.QRData == "" {
		return errors.New("QR код пустой")
	}

	checkSource := dto.CheckSource{
		CheckID: checkID,
		QR:      input.QRData,
	}

	return s.repo.Create(&checkSource)
}

// Парсинг JSON из API на товары
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
			Price:    float64(it.Price) / 100, // если в копейках
			Quantity: it.Quantity,
		}
	}

	return items, nil
}
