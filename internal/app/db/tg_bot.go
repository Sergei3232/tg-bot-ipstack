package db

import (
	sq "github.com/Masterminds/squirrel"
)

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
		if errDB != nil {
			return errDB
		}
		for rows.Next() {
			return nil
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

func (r *repository) DeleteAdmin(id int) error {
	return nil
}

func (r *repository) AddAdmin(id int) error {
	//r.qb.Insert("")
	return nil
}
