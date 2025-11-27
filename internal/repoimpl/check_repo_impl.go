package repo_impl

import (
	"context"

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

func (r *checkRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "Check" WHERE id=$1`,
		id,
	)
	return err
}
