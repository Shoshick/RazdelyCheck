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

func (r *groupRepo) ListByUser(userID uuid.UUID) ([]*dto.Group, error) {
	var groups []*dto.Group
	err := r.db.Select(
		&groups,
		`SELECT g.id
		 FROM "Group" g
		 JOIN user_to_group ug ON g.id = ug.group_id
		 WHERE ug.user_id = $1`,
		userID,
	)
	return groups, err
}

func (r *groupRepo) ListByGroup(groupID uuid.UUID) ([]*dto.User, error) {
	var users []*dto.User
	err := r.db.Select(
		&users,
		`SELECT u.id, u.email, u.name, u.owner
		 FROM "User" u
		 JOIN user_to_group ug ON u.id = ug.user_id
		 WHERE ug.group_id = $1`,
		groupID,
	)
	return users, err
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

func (r *groupRepo) ExistsUserInGroup(userID, groupID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Get(
		&exists,
		`SELECT EXISTS(SELECT 1 FROM user_to_group WHERE user_id=$1 AND group_id=$2)`,
		userID, groupID,
	)
	return exists, err
}

func (r *groupRepo) AddUserToGroup(userID, groupID uuid.UUID) error {
	_, err := r.db.Exec(
		`INSERT INTO user_to_group (id, group_id, user_id) VALUES ($1, $2, $3)`,
		uuid.New(), groupID, userID,
	)
	return err
}

func (r *groupRepo) RemoveUserFromGroup(userID, groupID uuid.UUID) error {
	_, err := r.db.Exec(
		`DELETE FROM user_to_group WHERE user_id=$1 AND group_id=$2`,
		userID, groupID,
	)
	return err
}
