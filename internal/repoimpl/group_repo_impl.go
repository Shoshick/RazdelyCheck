package repo_impl

import (
	"RazdelyCheck/internal/util"
	"context"
	"database/sql"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GroupRepo struct {
	db              *sqlx.DB
	CheckResultRepo repo.CheckResultRepo
}

func NewGroupRepo(db *sqlx.DB, crRepo repo.CheckResultRepo) *GroupRepo {
	return &GroupRepo{
		db:              db,
		CheckResultRepo: crRepo,
	}
}

func (r *GroupRepo) Create(g *dto.Group) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO "Group" (id) VALUES ($1)`,
		g.ID,
	)
	return err
}

func (r *GroupRepo) GetByID(id uuid.UUID) (*dto.Group, error) {
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

func (r *GroupRepo) List() ([]*dto.Group, error) {
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

func (r *GroupRepo) ListByUser(userID uuid.UUID) ([]*dto.Group, error) {
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

func (r *GroupRepo) ListByGroup(groupID uuid.UUID) ([]*dto.User, error) {
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

func (r *GroupRepo) Update(g *dto.Group) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE "Group" SET id=$1 WHERE id=$2`,
		g.ID, g.ID,
	)
	return err
}

func (r *GroupRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM "Group" WHERE id=$1`,
		id,
	)
	return err
}

func (r *GroupRepo) ExistsUserInGroup(userID, groupID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Get(
		&exists,
		`SELECT EXISTS(SELECT 1 FROM user_to_group WHERE user_id=$1 AND group_id=$2)`,
		userID, groupID,
	)
	return exists, err
}

func (r *GroupRepo) AddUserToGroup(userID, groupID uuid.UUID) error {
	return util.WithTransaction(r.db.DB, func(tx *sql.Tx) error {

		_, err := tx.Exec(`
			INSERT INTO user_to_group (id, group_id, user_id)
			VALUES ($1, $2, $3)
		`, uuid.New(), groupID, userID)
		if err != nil {
			return err
		}

		var checkID uuid.UUID
		err = tx.QueryRow(`
			SELECT id
			FROM "Check"
			WHERE group_id = $1
		`, groupID).Scan(&checkID)
		if err != nil {
			return err
		}

		cr := &dto.CheckResult{
			ID:       uuid.New(),
			CheckID:  checkID,
			UserID:   userID,
			TotalDue: 0,
		}
		err = r.CheckResultRepo.CreateCheckResultTx(tx, cr)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *GroupRepo) RemoveUserFromGroup(userID, groupID uuid.UUID, checkResultRepo repo.CheckResultRepo) error {
	return util.WithTransaction(r.db.DB, func(tx *sql.Tx) error {

		var checkID uuid.UUID
		err := tx.QueryRow(`
			SELECT id
			FROM "Check"
			WHERE group_id = $1
		`, groupID).Scan(&checkID)
		if err != nil {
			return err
		}

		results, err := checkResultRepo.GetCheckResultsByCheckID(checkID)
		if err != nil {
			return err
		}

		for _, cr := range results {
			if cr.UserID == userID {

				if err := checkResultRepo.DeleteCheckResult(cr.ID); err != nil {
					return err
				}
			}
		}

		_, err = tx.Exec(`
			DELETE FROM user_to_group
			WHERE user_id=$1 AND group_id=$2
		`, userID, groupID)
		return err
	})
}
