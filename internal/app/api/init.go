package api

import (
	"net/http"
	"tasklist_REST_API/internal/handlers"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const startMessage = `use /info`

func (a *api) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *api) configureRouterField() {

	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(startMessage))
	})

	a.router.HandleFunc("/info", handlers.GetInfo).Methods("GET")
}
