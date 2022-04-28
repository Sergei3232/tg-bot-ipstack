package db

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const adminRools string = "admin"

// HasAdministratorRools checking whether the administrator role is available to the user
func (r *repository) HasAdministratorRols(userTgId int) (bool, error) {
	userDb, errT := r.GetUserTelegram(userTgId)
	if errT != nil {
		return false, errT
	}

	idRoolAdmin, errAd := r.getIdRol(adminRools)
	if errAd != nil {
		return false, errAd
	}

	queryHasAdministratorRools, args, errQ := r.qb.
		Select("user_id").
		From("user_rols").
		Where(sq.Eq{"user_id": userDb.Id}).
		Where(sq.Eq{"rool_id": idRoolAdmin}).
		ToSql()

	if errQ != nil {
		return false, errQ
	}

	rows, errDB := r.db.Query(queryHasAdministratorRools, args...)
	defer rows.Close()

	if errDB != nil {
		return false, errDB
	}

	for rows.Next() {
		return true, nil
	}
	return false, nil
}

// AddNewUserBot adding a new bot user to the database
func (r *repository) AddNewUserBot(id int, nameUser string) error {
	ok, err := r.UserExists(id)
	if err != nil {
		return err
	}

	if !ok {
		queryAddNewUser, args, err := r.qb.Insert("users").
			Columns("name,telegram_id").
			Values(nameUser, id).
			ToSql()

		if err != nil {
			return err
		}

		rows, errDB := r.db.Query(queryAddNewUser, args...)
		defer rows.Close()

		if errDB != nil {
			return errDB
		}
	}
	return nil
}

// UserExists checking the user's existence
func (r *repository) UserExists(userTgID int) (bool, error) {
	queryUserExists, args, err := r.qb.
		Select("id").
		From("users").
		Where(sq.Eq{"telegram_id": userTgID}).
		ToSql()

	if err != nil {
		return false, err
	}

	rows, errDB := r.db.Query(queryUserExists, args...)
	defer rows.Close()

	if errDB != nil {
		return false, errDB
	}

	for rows.Next() {
		return true, nil
	}
	return false, nil
}

// GetUsersTelegram get the telegram bot user from the database
func (r *repository) GetUsersTelegram() ([]UserDb, error) {
	listUsers := make([]UserDb, 0)
	queryGetUsersTelegram, _, err := r.qb.
		Select("id, name, telegram_id").
		From("users").
		ToSql()

	if err != nil {
		return []UserDb{}, err
	}

	rows, errDB := r.db.Query(queryGetUsersTelegram)
	defer rows.Close()

	if errDB != nil {
		return []UserDb{}, errDB
	}

	for rows.Next() {
		var id, telegramId int
		var name string
		errScan := rows.Scan(&id, &name, &telegramId)
		if errScan != nil {
			return nil, errScan
		}

		listUsers = append(listUsers, UserDb{id, name, telegramId})
	}

	return listUsers, nil
}

// getIdRool
func (r *repository) getIdRol(nameRool string) (int, error) {
	queryGetIdRool, args, err := r.qb.
		Select("id").
		From("rols").
		Where(sq.Eq{"name": nameRool}).
		ToSql()

	if err != nil {
		return 0, err
	}

	rows, errDB := r.db.Query(queryGetIdRool, args...)
	defer rows.Close()

	if errDB != nil {
		return 0, errDB
	}

	for rows.Next() {
		var id int
		errScan := rows.Scan(&id)
		if errScan != nil {
			return 0, errScan
		}
		return id, nil
	}
	return 0, errors.New("role with this id was not found")
}

func (r *repository) GetUserTelegram(id int) (*UserDb, error) {
	queryGetUserTelegram, args, err := r.qb.
		Select("id, name, telegram_id").
		From("users").
		Where(sq.Eq{"telegram_id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, errDB := r.db.Query(queryGetUserTelegram, args...)
	defer rows.Close()
	if errDB != nil {
		return nil, errDB
	}

	for rows.Next() {
		var id, telegramId int
		var name string
		errScan := rows.Scan(&id, &name, &telegramId)

		if errScan != nil {
			return nil, errScan
		}

		return &UserDb{id, name, telegramId}, nil
	}
	return nil, errors.New("user does not exist")
}

func (r *repository) DeleteAdmin(id int) error {
	userDb, errU := r.GetUserTelegram(id)
	if errU != nil {
		return errU
	}

	idRoolAdmin, errAd := r.getIdRol(adminRools)
	if errAd != nil {
		return errAd
	}

	queryDeleteAdmin, args, err := r.qb.
		Delete("user_rols").
		Where(sq.Eq{"user_id": userDb.Id}).
		Where(sq.Eq{"rool_id": idRoolAdmin}).
		ToSql()

	if err != nil {
		return err
	}

	rows, errDB := r.db.Query(queryDeleteAdmin, args...)
	defer rows.Close()

	if errDB != nil {
		return errDB
	}

	return nil
}

func (r *repository) AddAdmin(id int) error {
	userDb, errU := r.GetUserTelegram(id)
	if errU != nil {
		return errU
	}

	idRoolAdmin, errAd := r.getIdRol(adminRools)
	if errAd != nil {
		return errAd
	}

	queryAddAdmin, args, err := r.qb.Insert("user_rols").
		Columns("user_id, rol_id").
		Values(userDb.Id, idRoolAdmin).
		ToSql()

	if err != nil {
		return err
	}

	rows, errDB := r.db.Query(queryAddAdmin, args...)
	defer rows.Close()

	if errDB != nil {
		return errDB
	}

	return nil
}
