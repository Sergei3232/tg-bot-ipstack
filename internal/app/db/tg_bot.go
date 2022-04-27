package db

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const adminRools string = "admin"

func (r *repository) HasAdministratorRools(userTgId int) (bool, error) {

	return false, nil
}

func (r *repository) AddNewUserBot(id int, nameUser string) error {
	ok, err := r.UserExists(id)
	if err != nil {
		return err
	}

	if !ok {
		queryAddNewUser, args, err := r.qb.Insert("users").
			Columns("    name,telegram_id").
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

func (r *repository) GetUsersTelegram() ([]int, error) {
	return []int{}, nil
}

func (r *repository) getIdRool(nameRool string) (int, error) {
	queryGetIdRool, args, err := r.qb.
		Select("id").
		From("rools").
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

	idRoolAdmin, errAd := r.getIdRool(adminRools)
	if errAd != nil {
		return errAd
	}

	queryDeleteAdmin, args, err := r.qb.
		Delete("user_rools").
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

	idRoolAdmin, errAd := r.getIdRool(adminRools)
	if errAd != nil {
		return errAd
	}

	queryAddAdmin, args, err := r.qb.Insert("user_rools").
		Columns("user_id, rool_id").
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
