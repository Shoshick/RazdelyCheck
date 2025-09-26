package repo_impl

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type checkResultRepo struct {
	db *sqlx.DB
}

func NewCheckResultRepo(db *sqlx.DB) repo.CheckResultRepo {
	return &checkResultRepo{db: db}
}

func (r *checkResultRepo) ListByCheckID(checkID uuid.UUID) ([]*dto.CheckResult, error) {
	var results []*dto.CheckResult
	err := r.db.Select(&results, "SELECT id, check_id, user_id, amount FROM check_results WHERE check_id = $1", checkID)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *checkResultRepo) ListByUserID(userID uuid.UUID) ([]*dto.CheckResult, error) {
	var results []*dto.CheckResult
	err := r.db.Select(&results, "SELECT id, check_id, user_id, amount FROM check_results WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return results, nil
}
