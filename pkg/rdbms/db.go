package rdbms

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args []any) error
	QueryContext(ctx context.Context, query string, in []any, out []any) error
	QueryRowContext(ctx context.Context, query string, in []any, out []any) error
	Close() error
}

type db struct {
	db *sql.DB
}

var (
	ErrorNotFound       = "error no entry found with the given arguments"
	ErrDuplicate        = "error operation would result in a duplicate entry"
	ErrPrepareStatement = "error while trying to prepare statement"
)

func (db *db) ExecContext(ctx context.Context, query string, args []any) error {
	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("error closing statement: %v", err)
		}
	}(stmt)

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("%s: %w", ErrDuplicate, err)
		}

		return fmt.Errorf("error when trying to execute statement: %w", err)
	}

	if rows, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("error getting number of affected rows: %w", err)
	} else if rows == 0 {
		return fmt.Errorf("%s: %w", ErrorNotFound, err)
	}

	return nil
}

func (db *db) QueryContext(ctx context.Context, query string, in []any, out []any) error {
	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("error closing statement: %v", err)
		}
	}(stmt)

	rows, err := stmt.QueryContext(ctx, in...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", ErrorNotFound, err)
		}

		return fmt.Errorf("error while trying to execute statement: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("error closing rows: %v", err)
		}
	}(rows)

	var index = 0
	for ; rows.Next(); index++ {
		if err := rows.Scan(out[index]); err != nil {
			return fmt.Errorf("error while trying to scan row: %w", err)
		}
		index++
	}
	out = out[:index]

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error while trying to iterate over rows: %w", err)
	}

	return nil
}

func (db *db) QueryRowContext(ctx context.Context, query string, in []any, out []any) error {
	stmt, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("error closing statement: %v", err)
		}
	}(stmt)

	if err := stmt.QueryRowContext(ctx, in...).Scan(out...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", ErrorNotFound, err)
		}

		return fmt.Errorf("error while trying to execute statement: %w", err)
	}

	return nil
}

func (db *db) Close() error {
	return db.db.Close()
}
