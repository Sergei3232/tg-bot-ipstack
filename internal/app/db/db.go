package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"time"
)

type Repository interface {
	HasAdministratorRols(id int) (bool, error)
	AddNewUserBot(id int, nameUser string) error
	GetUsersTelegram(idAdmin int) ([]UserDb, error)
	DeleteAdmin(idUser, idAdmin int) error
	AddAdmin(idUser, idAdmin int) error
	GetUserTelegram(id int) (*UserDb, error)
	getIdRol(name string) (int, error)
	recordRolExists(idUser, idRol int) (bool, error)
	AddUserHistoryQuery(idUser int, ip, queryResult string, timeQuery time.Time) error
	GetHistoryUserQuery(idUser int) ([]string, error)
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
