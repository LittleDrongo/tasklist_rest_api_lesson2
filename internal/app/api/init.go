package api

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const startMessage = `Привет!
Данный сервер поддерживает следующие методы REST API:
POST	/tasks/              :  создаёт задачу и возвращает её ID
GET		/tasks/<taskid>      :  возвращает одну задачу по её ID
GET		/tasks/              :  возвращает все задачи
DELETE	/tasks/<taskid>      :  удаляет задачу по ID
DELETE	/tasks/              :  удаляет все задачи
GET		/tags/<tagname>      :  возвращает список задач с заданным тегом
GET		/due/<yy>/<mm>/<dd>  :  возвращает список задач, запланированных на указанную дату

Запущен: %s
`

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

	msg := fmt.Sprintf(startMessage, time.Now().Format("02-01-2006 15:04"))

	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	})
}
