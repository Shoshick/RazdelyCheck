package repo_impl

import (
	"context"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type groupRepo struct {
	db *sqlx.DB
}

func NewGroupRepo(db *sqlx.DB) repo.GroupRepo {
	return &groupRepo{db: db}
}

func (r *groupRepo) Create(g *dto.Group) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO "Group" (id) VALUES ($1)`,
		g.ID,
	)
	return err
}

func (r *groupRepo) GetByID(id uuid.UUID) (*dto.Group, error) {
	g := &dto.Group{}
	err := r.db.GetContext(
		context.Background(),
		g,
		`SELECT id FROM "Group" WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *groupRepo) List() ([]*dto.Group, error) {
	var groups []*dto.Group
	err := r.db.SelectContext(
		context.Background(),
		&groups,
		`SELECT id FROM "Group"`,
	)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *groupRepo) Update(g *dto.Group) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "Group" SET id=$1 WHERE id=$2`,
		g.ID, g.ID,
	)
	return err
}

func (r *groupRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "Group" WHERE id=$1`,
		id,
	)
	return err
}
