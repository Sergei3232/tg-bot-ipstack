package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Repository interface {
	HasAdministratorRools(id int) (bool, error)
	AddNewUserBot(id int, nameUser string) error
	GetUsersTelegram() ([]int, error)
	DeleteAdmin(id int) (err error)
	AddAdmin(id int) (err error)
}

type repository struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

func NewDbConnectClient(sqlConnect string) (Repository, error) {
	bd, err := sql.Open("postgres", sqlConnect) //postgres
	if err != nil {
		return nil, err
	}
	return &repository{bd, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
}
