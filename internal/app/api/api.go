package api

import (
	"database/sql"
	"net/http"
	"tasklist_REST_API/internal/model/db"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api struct {
	config *config
	logger *logrus.Logger
	router *mux.Router
	db     *sql.DB
}

func New(config *config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *Api) Start() error {

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
