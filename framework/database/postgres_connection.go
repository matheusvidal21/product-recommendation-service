package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
)

func NewPostgresConnection(postgresUrl, port, user, password, dbName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresUrl, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	logger.Info(fmt.Sprintf("Connected to Postgres: %s", postgresUrl))
	return db, nil
}
