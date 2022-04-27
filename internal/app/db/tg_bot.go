package db

func (r *repository) HasAdministratorRools(userTgId int) (bool, error) {
	return false, nil
}

func (r *repository) AddNewUserBot(id int, nameUser string) error {
	return nil
}

func (r *repository) GetUsersTelegram() ([]int, error) {
	return []int{}, nil
}

func (r *repository) DeleteAdmin(id int) error {
	return nil
}

func (r *repository) AddAdmin(id int) error {
	return nil
}
