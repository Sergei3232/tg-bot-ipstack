package db

type UserDb struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	TelegramId int    `json:"telegram_id"`
}

type UserRequestHistory struct {
	Id          int    `json:"id"`
	Ip          string `json:"ip"`
	QueryResult string `json:"query_result"`
	TimeQuery   string `json:"time_query"`
}

type UserRequests []UserRequestHistory

type HistoryUser struct {
	UserDb       `json:"user"`
	UserRequests `json:"query"`
}
