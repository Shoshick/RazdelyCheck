package repo_impl

import (
	"context"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type itemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) repo.ItemRepo {
	return &itemRepo{db: db}
}

func (r *itemRepo) Create(i *dto.Item) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO Item (id, check_id, position, name, price, quantity) VALUES ($1,$2,$3,$4,$5,$6)`,
		i.ID, i.CheckID, i.Position, i.Name, i.Price, i.Quantity,
	)
	return err
}

func (r *itemRepo) GetByID(id uuid.UUID) (*dto.Item, error) {
	i := &dto.Item{}
	err := r.db.GetContext(
		context.Background(),
		i,
		`SELECT id, check_id, position, name, price, quantity FROM Item WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (r *itemRepo) ListByCheckID(checkID uuid.UUID) ([]*dto.Item, error) {
	var items []*dto.Item
	err := r.db.SelectContext(
		context.Background(),
		&items,
		`SELECT id, check_id, position, name, price, quantity FROM Item WHERE check_id=$1`,
		checkID,
	)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepo) Update(i *dto.Item) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`UPDATE Item SET check_id=$1, position=$2, name=$3, price=$4, quantity=$5 WHERE id=$6`,
		i.CheckID, i.Position, i.Name, i.Price, i.Quantity, i.ID,
	)
	return err
}

func (r *itemRepo) Delete(id uuid.UUID) error {
	_, err := r.db.ExecContext(
		context.Background(),
		`DELETE FROM Item WHERE id=$1`,
		id,
	)
	return err
}
