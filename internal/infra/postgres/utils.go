package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxFunc func(*sql.Tx) error

func WithTx(db *sqlx.DB, fn TxFunc) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("cannot begin tx: %w", err)
	}

	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)

		case err != nil:
			rollbackErr := tx.Rollback()
			if rollbackErr != nil { // make the original err accessible by errors.Is
				err = fmt.Errorf("cannot rollback transaction: %s. original error: %w", rollbackErr.Error(), err)
			}

		default:
			err = tx.Commit()
			if err != nil {
				err = fmt.Errorf("cannot commit tx: %w", err)
			}
		}
	}()

	err = fn(tx)
	return
}
