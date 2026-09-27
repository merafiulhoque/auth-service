package database

import (
	"auth-service/internal/shared/domainerrors"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, domainerrors.ErrDbConnection
	}
	return db, nil
}
