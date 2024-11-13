package api

import (
	"net/http"
	"tasklist_REST_API/internal/model/db"

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

	err = db.DeployDB(a.config.DataBasePath)
	if err != nil {
		a.logger.Fatal("fatal deploy database \n", err)
	}

	a.logger.Info("starting api server at port:", (a.config.Host + a.config.Port))
	a.router.StrictSlash(true)
	a.configureRouterField()

	return http.ListenAndServe(a.config.Port, a.router)
}
