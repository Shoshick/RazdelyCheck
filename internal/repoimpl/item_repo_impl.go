package repo_impl

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"context"

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
		`INSERT INTO Item (id, check_id, position, name, price, quantity, is_excluded)
 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		i.ID, i.CheckID, i.Position, i.Name, i.Price, i.Quantity, i.IsExcluded,
	)
	return err
}

func (r *itemRepo) GetByID(id uuid.UUID) (*dto.Item, error) {
	i := &dto.Item{}
	err := r.db.GetContext(
		context.Background(),
		i,
		`SELECT id, check_id, position, name, price, quantity, is_excluded
 FROM Item WHERE id=$1`,
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
		`SELECT id, check_id, position, name, price, quantity, is_excluded
 FROM Item WHERE check_id=$1`,
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
		`UPDATE Item
 SET check_id=$1, position=$2, name=$3, price=$4, quantity=$5, is_excluded=$6
 WHERE id=$7`,
		i.CheckID, i.Position, i.Name, i.Price, i.Quantity, i.ID, i.IsExcluded,
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

func (r *itemRepo) ExcludeItemTx(tx *sqlx.Tx, id uuid.UUID) error {
	_, err := tx.ExecContext(
		context.Background(),
		`UPDATE Item SET is_excluded = true WHERE id=$1`,
		id,
	)
	return err
}

func (r *itemRepo) IncludeItemTx(tx *sqlx.Tx, id uuid.UUID) error {
	_, err := tx.ExecContext(
		context.Background(),
		`UPDATE Item SET is_excluded = false WHERE id=$1`,
		id,
	)
	return err
}

func (r *itemRepo) GetItemsByCheckIDTx(tx *sqlx.Tx, checkID uuid.UUID) ([]*dto.Item, error) {
	var items []*dto.Item
	err := tx.Select(&items, `
		SELECT id, check_id, position, name, price, quantity
		FROM public.item
		WHERE check_id=$1
	`, checkID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
