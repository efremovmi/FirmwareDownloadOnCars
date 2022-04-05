package repository

import "github.com/jmoiron/sqlx"

type Repo struct {
	connection *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *Repo {
	return &Repo{connection: conn}
}
