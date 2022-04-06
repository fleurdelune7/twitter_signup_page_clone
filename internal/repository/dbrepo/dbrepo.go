package dbrepo

import (
	"database/sql"

	"bitbucket.org/janpavtel/site/internal/repository"
)


func NewPostgresRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}