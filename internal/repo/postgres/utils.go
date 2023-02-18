package postgres

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/maskot/internal/repo"
)

func wrapErrNoRows(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return repo.ErrNotFound
	}
	return err
}

func isUniqueViolation(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

func wrapUniqueViolation(err error) error {
	if isUniqueViolation(err) {
		return repo.ErrAlreadyExists
	}
	return err
}
