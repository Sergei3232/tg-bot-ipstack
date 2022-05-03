package apiserver

import (
	"encoding/json"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//APIServer ...
type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
	DB     db.Repository
}

//NewAPIServer ...
func NewAPIServer(config *config.Config) *APIServer {

	dbClient, errDb := db.NewDbConnectClient(config.DnsDB)
	if errDb != nil {
		log.Panic(errDb)
	}
	return &APIServer{
		config,
		logrus.New(),
		mux.NewRouter(),
		dbClient,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting API server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/get_users", s.getListUser())
	s.router.HandleFunc("/get_user", s.getUserById())
	s.router.HandleFunc("/get_history_by_tg", s.getHistoryByTg())
	s.router.HandleFunc("/history_by_id/{id}", s.deleteUserHistoryRecord()).Methods("DELETE")
}

func (s *APIServer) getListUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listUser, err := s.DB.GetListUsers()
		if err != nil {
			s.logger.Error("APIServer.getListUser ", err.Error())
		}
		jsonText, _ := json.MarshalIndent(listUser, "", "  ")

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(jsonText))
	}
}

func (s *APIServer) getUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			s.logger.Error(err.Error())
		}

		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			s.logger.Error(err.Error())
		}

		val, ok := m["id"]
		id, err := strconv.Atoi(val[0])

		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := s.DB.GetUserTelegram(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		jsonText, _ := json.Marshal(user)
		io.WriteString(w, string(jsonText))
	}
}

func (s *APIServer) getHistoryByTg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			s.logger.Error(err.Error())
		}

		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			s.logger.Error(err.Error())
		}

		val, ok := m["id"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, errConv := strconv.Atoi(val[0])
		if errConv != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := s.DB.GetUserRequestHistory(id)
		jsonText, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(jsonText))
	}
}

func (s *APIServer) deleteUserHistoryRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		idStr := param["id"]

		id, errConv := strconv.Atoi(idStr)
		if errConv != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := s.DB.DeleteRecordUserHistory(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
