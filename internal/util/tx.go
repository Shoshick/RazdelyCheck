package util

import (
	"github.com/jmoiron/sqlx"
)

// Функция для выполнения транзакции с sqlx
func WithTransaction(db *sqlx.DB, fn func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
