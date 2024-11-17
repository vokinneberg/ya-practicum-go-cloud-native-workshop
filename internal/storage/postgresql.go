package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const (
	createShortURLQuery = `INSERT INTO urls (original_url, short_url) VALUES ($1, $2) ON CONFLICT (original_url) DO NOTHING`
	getOriginalURLQuery = `SELECT original_url FROM urls WHERE short_url = $1`
)

var (
	ErrNotFound  = errors.New("short URL not found")
	ErrDuplicate = errors.New("short URL already exists")
)

type PostgresStorage struct {
	db                 *sql.DB
	createShortURLStmt *sql.Stmt
	getOriginalURLStmt *sql.Stmt
}

func NewPostgresStorage(db *sql.DB) (*PostgresStorage, error) {
	createShortURLStmt, err := db.Prepare(createShortURLQuery)
	if err != nil {
		return nil, err
	}

	getOriginalURLStmt, err := db.Prepare(getOriginalURLQuery)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db:                 db,
		createShortURLStmt: createShortURLStmt,
		getOriginalURLStmt: getOriginalURLStmt,
	}, nil
}

func (ps *PostgresStorage) CreateShortURL(ctx context.Context, originalURL string, shortURL string) error {
	_, err := ps.createShortURLStmt.ExecContext(ctx, originalURL, shortURL)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (ps *PostgresStorage) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	var originalURL string
	if err := ps.getOriginalURLStmt.QueryRowContext(ctx, shortURL).Scan(&originalURL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", err
	}
	return originalURL, nil
}
