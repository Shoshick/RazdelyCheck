package repo_impl

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"RazdelyCheck/internal/util"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type checkResultRepoImpl struct {
	db *sqlx.DB
}

func NewCheckResultRepo(db *sqlx.DB) repo.CheckResultRepo {
	return &checkResultRepoImpl{db: db}
}

func (r *checkResultRepoImpl) CreateCheckResultTx(tx *sql.Tx, cr *dto.CheckResult) error {
	_, err := tx.Exec(`
		INSERT INTO public.check_result (id, check_id, user_id, total_due)
		VALUES ($1, $2, $3, $4)
	`, cr.ID, cr.CheckID, cr.UserID, cr.TotalDue)
	return err
}

func (r *checkResultRepoImpl) DeleteCheckResult(id uuid.UUID) error {
	_, err := r.db.Exec(`DELETE FROM public.check_result WHERE id=$1`, id)
	return err
}

func (r *checkResultRepoImpl) GetCheckResultsByCheckID(checkID uuid.UUID) ([]dto.CheckResult, error) {
	var results []dto.CheckResult
	err := r.db.Select(&results, `
		SELECT id, check_id, user_id, total_due
		FROM public.check_result
		WHERE check_id=$1
	`, checkID)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *checkResultRepoImpl) AddItemToCheckResult(item *dto.ItemToCheckResult) error {
	return util.WithTransaction(r.db.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			INSERT INTO public.item_to_check_result (id, item_id, check_result_id, quantity)
			VALUES ($1, $2, $3, $4)
		`, item.ID, item.ItemID, item.CheckResultID, item.Quantity)
		if err != nil {
			return err
		}

		var total float64
		err = tx.QueryRow(`
			SELECT COALESCE(SUM(i.price * itcr.quantity),0)
			FROM public.item_to_check_result itcr
			JOIN public.item i ON i.id = itcr.item_id
			WHERE itcr.check_result_id=$1
		`, item.CheckResultID).Scan(&total)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE public.check_result SET total_due=$1 WHERE id=$2`, total, item.CheckResultID)
		return err
	})
}

func (r *checkResultRepoImpl) UpdateItemQuantityInCheckResult(itemID, checkResultID uuid.UUID, quantity float64) error {
	return util.WithTransaction(r.db.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			UPDATE public.item_to_check_result
			SET quantity=$1
			WHERE item_id=$2 AND check_result_id=$3
		`, quantity, itemID, checkResultID)
		if err != nil {
			return err
		}

		var total float64
		err = tx.QueryRow(`
			SELECT COALESCE(SUM(i.price * itcr.quantity),0)
			FROM public.item_to_check_result itcr
			JOIN public.item i ON i.id = itcr.item_id
			WHERE itcr.check_result_id=$1
		`, checkResultID).Scan(&total)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE public.check_result SET total_due=$1 WHERE id=$2`, total, checkResultID)
		return err
	})
}

func (r *checkResultRepoImpl) RemoveItemFromCheckResult(itemID, checkResultID uuid.UUID) error {
	return util.WithTransaction(r.db.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			DELETE FROM public.item_to_check_result
			WHERE item_id=$1 AND check_result_id=$2
		`, itemID, checkResultID)
		if err != nil {
			return err
		}

		var total float64
		err = tx.QueryRow(`
			SELECT COALESCE(SUM(i.price * itcr.quantity),0)
			FROM public.item_to_check_result itcr
			JOIN public.item i ON i.id = itcr.item_id
			WHERE itcr.check_result_id=$1
		`, checkResultID).Scan(&total)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`UPDATE public.check_result SET total_due=$1 WHERE id=$2`, total, checkResultID)
		return err
	})
}

func (r *checkResultRepoImpl) GetItemsByCheckResultID(checkResultID uuid.UUID) ([]dto.ItemToCheckResult, error) {
	rows, err := r.db.Query(`
		SELECT id, item_id, check_result_id, quantity
		FROM public.item_to_check_result
		WHERE check_result_id=$1
	`, checkResultID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var items []dto.ItemToCheckResult
	for rows.Next() {
		var it dto.ItemToCheckResult
		if err := rows.Scan(&it.ID, &it.ItemID, &it.CheckResultID, &it.Quantity); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

func (r *checkResultRepoImpl) UpdateCheckResultTotal(crID uuid.UUID) error {
	_, err := r.db.Exec(`
		UPDATE public.check_result
		SET total_due = (
			SELECT COALESCE(SUM(i.price * itcr.quantity),0)
			FROM public.item_to_check_result itcr
			JOIN public.item i ON i.id = itcr.item_id
			WHERE itcr.check_result_id=$1
		)
		WHERE id=$1
	`, crID)
	return err
}
