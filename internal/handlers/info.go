package handlers

import (
	"log"
	"net/http"
)

func GetInfo(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	message := `Привет!
Данный сервер поддерживает следующие методы REST API:
POST	/tasks/              :  создаёт задачу и возвращает её ID
GET		/tasks/<taskid>      :  возвращает одну задачу по её ID
GET		/tasks/              :  возвращает все задачи
DELETE	/tasks/<taskid>      :  удаляет задачу по ID
DELETE	/tasks/              :  удаляет все задачи
GET		/tags/<tagname>      :  возвращает список задач с заданным тегом
GET		/due/<yy>/<mm>/<dd>  :  возвращает список задач, запланированных на указанную дату
	`
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(message))
	log.Println("started GetInfo method \n",
		"user:", request.Header, "\n",
		request.RemoteAddr)
}
