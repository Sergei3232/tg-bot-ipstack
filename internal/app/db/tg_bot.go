package db

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const adminRools string = "admin"

// HasAdministratorRols checking whether the administrator role is available to the user
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
		Where(sq.Eq{"rol_id": idRoolAdmin}).
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
func (r *repository) GetUsersTelegram(idAdmin int) ([]UserDb, error) {
	ok, err := r.HasAdministratorRols(idAdmin)
	if err != nil {
		return []UserDb{}, err
	}
	if !ok {
		return []UserDb{}, errors.New("Команда недоступна!")
	}

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

// getIdRool get the role id
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

// GetUserTelegram get all telegram users
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
		var idHistory, telegramId int
		var name string
		errScan := rows.Scan(&idHistory, &name, &telegramId)

		if errScan != nil {
			return nil, errScan
		}

		return &UserDb{idHistory, name, telegramId}, nil
	}
	return nil, errors.New("user does not exist")
}

// DeleteAdmin remove the administrator role from the user
func (r *repository) DeleteAdmin(idUser, idAdmin int) error {
	ok, err := r.HasAdministratorRols(idAdmin)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Команда недоступна!")
	}

	userDb, errU := r.GetUserTelegram(idUser)
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
		Where(sq.Eq{"rol_id": idRoolAdmin}).
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

// AddAdmin add an administrator role to a user
func (r *repository) AddAdmin(idUser, idAdmin int) error {
	ok, err := r.HasAdministratorRols(idAdmin)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Команда недоступна!")
	}

	userDb, errU := r.GetUserTelegram(idUser)
	if errU != nil {
		return errU
	}

	idRoolAdmin, errAd := r.getIdRol(adminRools)
	if errAd != nil {
		return errAd
	}

	ok, errRol := r.recordRolExists(userDb.Id, idRoolAdmin)
	if errRol != nil {
		return errRol
	}

	if !ok {
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
	}

	return nil
}

// recordRolExists сhecking the existence of a user role
func (r *repository) recordRolExists(idUser, idRol int) (bool, error) {
	queryUserExists, args, err := r.qb.
		Select("id").
		From("user_rols").
		Where(sq.Eq{"user_id": idUser}).
		Where(sq.Eq{"rol_id": idRol}).
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

//AddUserHistoryQuery add Adding a user request history
func (r *repository) AddUserHistoryQuery(
	idUser int,
	ip, queryResult string,
	timeQuery time.Time) error {

	queryAddUserHistory, args, err := r.qb.Insert("user_request_history").
		Columns("userid, ip, query_result, time_query").
		Values(idUser, ip, queryResult, timeQuery).
		ToSql()

	if err != nil {
		return err
	}

	rows, errDB := r.db.Query(queryAddUserHistory, args...)
	defer rows.Close()

	if errDB != nil {
		return errDB
	}

	return nil
}

//GetHistoryUserQuery Getting a list of user requests
func (r *repository) GetHistoryUserQuery(idUser int) (string, error) {
	var listQueryUser string

	userDb, errU := r.GetUserTelegram(idUser)
	if errU != nil {
		return "", errU
	}

	queryGetUserTelegram, args, err := r.qb.
		Select("ip, query_result").
		From("user_request_history").
		Where(sq.Eq{"userid": userDb.Id}).
		GroupBy("ip", "query_result").
		ToSql()

	if err != nil {
		return "", err
	}

	rows, errDB := r.db.Query(queryGetUserTelegram, args...)
	defer rows.Close()
	if errDB != nil {
		return "", errDB
	}

	for rows.Next() {
		var ip, query string
		errScan := rows.Scan(&ip, &query)

		if errScan != nil {
			return "", errScan
		}
		listQueryUser += ip + "\n" + query + "\n" + "------------------------" + "\n"
	}

	return listQueryUser, nil
}

//GetHistoryUserQueryAdmin Get a list of the user's ip search
func (r *repository) GetHistoryUserQueryAdmin(idUser, adminId int) (string, error) {
	var listQueryUser string

	userDb, errU := r.GetUserTelegram(idUser)
	if errU != nil {
		return "", errU
	}

	ok, err := r.HasAdministratorRols(adminId)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("Команда недоступна!")
	}

	queryGetUserTelegram, args, err := r.qb.
		Select("ip").
		From("user_request_history").
		Where(sq.Eq{"userid": userDb.Id}).
		GroupBy("ip").
		ToSql()

	if err != nil {
		return "", err
	}

	rows, errDB := r.db.Query(queryGetUserTelegram, args...)
	defer rows.Close()
	if errDB != nil {
		return "", errDB
	}

	for rows.Next() {
		var ip string
		errScan := rows.Scan(&ip)

		if errScan != nil {
			return "", errScan
		}
		listQueryUser += ip + "\n" + "------------------------" + "\n"
	}

	return listQueryUser, nil
}

//GetListUsers Get a list of users to GET request
func (r *repository) GetListUsers() ([]UserDb, error) {

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

//GetUserRequestHistory Get user request history to the GET request
func (r *repository) GetUserRequestHistory(idUser int) (HistoryUser, error) {
	listQuery := HistoryUser{}

	userDb, errU := r.GetUserTelegram(idUser)
	if errU != nil {
		return HistoryUser{}, errU
	}
	listQuery.User = *userDb
	listQuery.ListQuery = make([]UserRequestHistory, 0)

	queryGetUserTelegram, args, err := r.qb.
		Select("id, ip, query_result, time_query").
		From("user_request_history").
		Where(sq.Eq{"userid": userDb.Id}).
		GroupBy("id", "ip", "query_result").
		ToSql()

	if err != nil {
		return HistoryUser{}, err
	}

	rows, errDB := r.db.Query(queryGetUserTelegram, args...)
	defer rows.Close()
	if errDB != nil {
		return HistoryUser{}, errDB
	}

	for rows.Next() {
		var id int
		var ip, query, timeQuery string
		errScan := rows.Scan(&id, &ip, &query, &timeQuery)

		if errScan != nil {
			return HistoryUser{}, errScan
		}
		listQuery.ListQuery = append(listQuery.ListQuery, UserRequestHistory{
			id,
			ip,
			query,
			timeQuery})
	}

	return listQuery, nil
}

func (r *repository) DeleteRecordUserHistory(id int) error {
	queryGetUserTelegram, args, err := r.qb.
		Delete("user_request_history").Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, errDB := r.db.Query(queryGetUserTelegram, args...)

	if errDB != nil {
		return errDB
	}

	return nil
}
