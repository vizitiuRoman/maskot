package postgres

import (
	"database/sql"
	"errors"

	"github.com/maskot/internal/repo"
)

func wrapErrNoRows(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return repo.ErrNotFound
	}
	return err
}
