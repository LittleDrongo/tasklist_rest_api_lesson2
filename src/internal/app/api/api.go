package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type api struct {
	config *config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *config) *api {
	return &api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *api) Start() error {

	err := a.configureLoggerField()
	if err != nil {
		return err
	}

	a.logger.Info("starting api server at port:", (a.config.Host + a.config.Port))

	a.configureRouterField()

	return http.ListenAndServe(a.config.Port, a.router)
}
