package db

type UserDb struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	TelegramId int    `json:"telegram_id"`
}
