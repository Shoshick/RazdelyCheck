package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"RazdelyCheck/internal/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CheckSourceService struct {
	repo         repo.CheckSourceRepo
	db           *sqlx.DB
	checkService *CheckService
	httpClient   *http.Client // <-- пригодится
	token        string       // <-- токен вынесем в конфиг
}

func NewCheckSourceService(r repo.CheckSourceRepo, db *sqlx.DB, cs *CheckService, token string) *CheckSourceService {
	return &CheckSourceService{
		repo:         r,
		db:           db,
		checkService: cs,
		httpClient:   &http.Client{},
		token:        token,
	}
}

func (s *CheckSourceService) ProcessQR(userID uuid.UUID, input dto.QRScanInput, jsonData []byte) (*dto.Check, error) {
	if input.QRData == "" {
		return nil, errors.New("QR is empty")
	}

	jsonData, err := s.FetchCheckJSON(input.QRData)
	if err != nil {
		return nil, fmt.Errorf("fetch JSON failed: %w", err)
	}

	var check *dto.Check
	var items []dto.Item
	var totalSum int64

	err = util.WithTransaction(s.db, func(tx *sqlx.Tx) error {

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

func (s *CheckSourceService) FetchCheckJSON(qrRaw string) ([]byte, error) {
	apiURL := "https://proverkacheka.com/api/v1/check/get"

	data := url.Values{}
	data.Set("token", s.token)
	data.Set("qrraw", qrRaw)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
