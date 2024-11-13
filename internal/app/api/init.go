package api

import (
	"net/http"
	"tasklist_REST_API/internal/handlers"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const startMessage = `use /info`

func (a *Api) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *Api) configureRouterField() {

	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(startMessage))
	})

	a.router.HandleFunc("/info", handlers.GetInfo).Methods("GET")
	a.router.HandleFunc("/tasks", handlers.PostTasks).Methods("POST")
	a.router.HandleFunc("/tasks"+"/{id}", handlers.GetTaskById).Methods("GET")
	a.router.HandleFunc("/tasks", handlers.GetAllTasks).Methods("GET")
	a.router.HandleFunc("/tags"+"/{tags}", handlers.GetTasksByTags).Methods("GET")
}
