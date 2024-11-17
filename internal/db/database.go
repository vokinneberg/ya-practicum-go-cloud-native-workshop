package db

import (
	"context"
	"database/sql"
)

type DBFunc func(tx *sql.Tx) error

type Database struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		DB: db,
	}
}

// WithTx возвращает функцию, которая выполняет переданную функцию в транзакции
// Делает Rollback , если произошла ошибка
func (d *Database) WithTx(ctx context.Context, fn DBFunc) error {
	tx, err := d.beginTx(ctx)
	if err != nil {
		return err
	}
	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr // rollback Error
		}
		return err // function error
	}
	return tx.Commit()
}

// BeginTx создает транзакцию
func (d *Database) beginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
