package repo_impl

import (
	"context"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(u *dto.User) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO "User" (id, email, name, owner) VALUES ($1, $2, $3, $4)`,
		u.ID, u.Email, u.Name, u.OwnerID,
	)
	return err
}

func (r *userRepo) GetByID(id uuid.UUID) (*dto.User, error) {
	u := &dto.User{}
	err := r.db.GetContext(
		context.Background(),
		u,
		`SELECT id, email, name, owner FROM "User" WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepo) List() ([]*dto.User, error) {
	var users []*dto.User
	err := r.db.SelectContext(
		context.Background(),
		&users,
		`SELECT id, email, name, owner FROM "User"`,
	)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Update(u *dto.User) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "User" SET email=$1, name=$2, owner=$3 WHERE id=$4`,
		u.Email, u.Name, u.OwnerID, u.ID,
	)
	return err
}

func (r *userRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "User" WHERE id=$1`,
		id,
	)
	return err
}
