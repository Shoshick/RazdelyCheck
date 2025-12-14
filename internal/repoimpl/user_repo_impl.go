package repo_impl

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"context"

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

func (r *userRepo) ListByOwner(id uuid.UUID) ([]*dto.User, error) {
	var users []*dto.User
	err := r.db.SelectContext(
		context.Background(),
		&users,
		`SELECT id, email, name, owner FROM "User" WHERE owner = $1`, id,
	)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, `SELECT EXISTS(SELECT 1 FROM "User" WHERE email = $1)`, email)
	return exists, err
}

func (r *userRepo) UpdateName(id uuid.UUID, name string) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "User" SET name = $1 WHERE id = $2`,
		name, id,
	)
	return err
}

// Обновить email пользователя
func (r *userRepo) UpdateEmail(id uuid.UUID, email string) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "User" SET email = $1 WHERE id = $2`,
		email, id,
	)
	return err
}

// Обновить owner пользователя
func (r *userRepo) UpdateOwner(id uuid.UUID, ownerID uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "User" SET owner = $1 WHERE id = $2`,
		ownerID, id,
	)
	return err
}

func (r *userRepo) GetOwnerForTempUser(tempUserID uuid.UUID) (uuid.UUID, error) {
	var ownerID uuid.UUID
	err := r.db.Get(&ownerID, `
		SELECT c.user_id
		FROM "Check" c
		JOIN "Group" g ON c.group_id = g.id
		JOIN "user_to_group" gu ON gu.group_id = g.id
		WHERE gu.user_id = $1
		LIMIT 1
	`, tempUserID)
	return ownerID, err
}

func (r *userRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "User" WHERE id=$1`,
		id,
	)
	return err
}

func (r *userRepo) DeleteByOwner(ownerID uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "User" WHERE owner = $1`,
		ownerID,
	)
	return err
}
