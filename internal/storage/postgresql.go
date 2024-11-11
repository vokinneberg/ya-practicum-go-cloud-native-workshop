package storage

import (
	"context"
	"database/sql"
)

const (
	createShortURLQuery = `INSERT INTO urls (original_url, short_url) VALUES ($1, $2) RETURNING short_url`
	getOriginalURLQuery = `SELECT original_url FROM urls WHERE short_url = $1`
)

type PostgresStorage struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

func NewPostgresStorage(db *sql.DB) (*PostgresStorage, error) {
	statements := make(map[string]*sql.Stmt)
	createShortURLStmt, err := db.Prepare(createShortURLQuery)
	if err != nil {
		return nil, err
	}
	statements["createShortURL"] = createShortURLStmt

	getOriginalURLStmt, err := db.Prepare(getOriginalURLQuery)
	if err != nil {
		return nil, err
	}
	statements["getOriginalURL"] = getOriginalURLStmt

	return &PostgresStorage{db: db, stmts: statements}, nil
}

func (ps *PostgresStorage) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	var shortURL string
	if err := ps.stmts["createShortURL"].QueryRowContext(ctx, originalURL, shortURL).Scan(&shortURL); err != nil {
		return "", err
	}
	return shortURL, nil
}

func (ps *PostgresStorage) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	var originalURL string
	if err := ps.stmts["getOriginalURL"].QueryRowContext(ctx, shortURL).Scan(&originalURL); err != nil {
		return "", err
	}
	return originalURL, nil
}
