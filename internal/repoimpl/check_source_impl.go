package repo_impl

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type checkSourceRepo struct {
	db *sqlx.DB
}

func NewCheckSourceRepo(db *sqlx.DB) repo.CheckSourceRepo {
	return &checkSourceRepo{db: db}
}

func (r *checkSourceRepo) Create(cs *dto.CheckSource) error {
	query := `
        INSERT INTO check_source (id, check_id, qr, data)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, cs.ID, cs.CheckID, cs.QR, cs.Data)
	return err
}

func (r *checkSourceRepo) GetByCheckID(checkID uuid.UUID) (*dto.CheckSource, error) {
	var cs dto.CheckSource
	err := r.db.Get(&cs, `SELECT id, check_id, data FROM check_sources WHERE check_id = $1`, checkID)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

func (r *checkSourceRepo) CreateTx(tx *sql.Tx, cs *dto.CheckSource) error {
	_, err := tx.ExecContext(
		context.Background(),
		`INSERT INTO check_sources (check_id, qr) VALUES ($1, $2)`,
		cs.CheckID,
		cs.QR,
	)
	return err
}

func (r *checkSourceRepo) CreateItemTx(tx *sql.Tx, item *dto.Item) error {
	_, err := tx.ExecContext(
		context.Background(),
		`INSERT INTO items (id, check_id, position, name, price, quantity)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		item.ID,
		item.CheckID,
		item.Position,
		item.Name,
		item.Price,
		item.Quantity,
	)
	return err
}
