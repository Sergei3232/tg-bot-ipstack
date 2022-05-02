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
	s.router.HandleFunc("/getListUser", s.getListUser())
	s.router.HandleFunc("/getUserById/:Id", s.getUserById())
}

func (s *APIServer) getListUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listUser, err := s.DB.GetListUsers()
		if err != nil {
			s.logger.Error("APIServer.getListUser ", err.Error())
		}
		jsonText, _ := json.MarshalIndent(listUser, "", "  ")

		io.WriteString(w, string(jsonText))
	}
}

func (s *APIServer) getUserById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "getUserById")
	}
}
