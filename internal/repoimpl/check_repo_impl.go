package repo_impl

import (
	"context"
	"database/sql"
	"fmt"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type checkRepo struct {
	db *sqlx.DB
}

func NewCheckRepo(db *sqlx.DB) repo.CheckRepo {
	return &checkRepo{db: db}
}

func (r *checkRepo) Create(c *dto.Check) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO "Check" (id, user_id, total_sum, group_id, created_at) VALUES ($1,$2,$3,$4,$5)`,
		c.ID, c.UserID, c.TotalSum, c.GroupID, c.CreatedAt,
	)
	return err
}

func (r *checkRepo) GetByID(id uuid.UUID) (*dto.Check, error) {
	c := &dto.Check{}
	err := r.db.GetContext(
		context.Background(),
		c,
		`SELECT id, user_id, group_id, total_sum, created_at FROM "Check" WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *checkRepo) GetCheckByIDTx(tx *sql.Tx, id uuid.UUID) (*dto.Check, error) {
	var ch dto.Check
	err := tx.QueryRow(`
		SELECT id, user_id, total_sum
		FROM public."Check"
		WHERE id=$1
	`, id).Scan(&ch.ID, &ch.UserID, &ch.TotalSum)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("check not found")
		}
		return nil, err
	}
	return &ch, nil
}

func (r *checkRepo) ListByUserID(userID uuid.UUID) ([]*dto.Check, error) {
	var checks []*dto.Check
	err := r.db.SelectContext(
		context.Background(),
		&checks,
		`SELECT id, user_id, group_id, total_sum, created_at FROM "Check" WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return checks, nil
}

func (r *checkRepo) UpdateTotalSum(id uuid.UUID) error {
	var total int64
	err := r.db.Get(&total, `
		SELECT COALESCE(SUM(price * quantity),0)
		FROM Item
		WHERE check_id = $1 AND is_excluded = false
	`, id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE "Check"
		SET total_sum = $2
		WHERE id = $1
	`, id, total)
	return err
}

func (r *checkRepo) UpdateTotalSumTx(tx *sql.Tx, checkID uuid.UUID) error {
	var total int64
	err := tx.QueryRowContext(
		context.Background(),
		`SELECT COALESCE(SUM(price * quantity),0)
		 FROM Item
		 WHERE check_id = $1 AND is_excluded = false`,
		checkID,
	).Scan(&total)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		context.Background(),
		`UPDATE "Check"
		 SET total_sum = $2
		 WHERE id = $1`,
		checkID, total,
	)
	return err
}

func (r *checkRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "Check" WHERE id=$1`,
		id,
	)
	return err
}

func (r *checkRepo) GetCheckByGroupID(groupID uuid.UUID) (uuid.UUID, error) {
	var checkID uuid.UUID
	err := r.db.Get(&checkID, `
		SELECT id
		FROM "Check"
		WHERE group_id = $1
	`, groupID)
	if err != nil {
		return uuid.Nil, err
	}
	return checkID, nil
}
