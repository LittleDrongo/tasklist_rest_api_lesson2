package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tasklist_REST_API/internal/model/db"
	"time"

	"github.com/gorilla/mux"
)

// POST   /tasks/              :  создаёт задачу и возвращает её ID
func PostTasks(writer http.ResponseWriter, request *http.Request) {

	initHeaders(writer)

	var task db.Task

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response := ResponseModel[int]{
			IsSuccess: false,
			Message:   "provide json file is invalid, look at sample request model.",
			Data:      nil,
			SampleRequestModel: &db.Task{
				Text: "sample task text",
				Tags: []string{"work", "bussines"},
				Due:  time.Now().Add(time.Hour * 48),
			},
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	id, err := db.InsertTask(task)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		response := ResponseModel[int]{
			IsSuccess: false,
			IsError:   true,
			Message:   err.Error(),
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	type taskId struct {
		TaskId int `json:"task_id"`
	}

	writer.WriteHeader(http.StatusOK) // 200 error
	response := ResponseModel[taskId]{
		IsSuccess: true,
		Message:   "Task added successfully",
		Data:      &taskId{TaskId: id},
	}
	json.NewEncoder(writer).Encode(response)
}

// GET    /tasks/<taskid>      :  возвращает одну задачу по её ID
func GetTaskById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	idStr := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response := ResponseModel[any]{
			IsSuccess: false,
			IsError:   true,
			Message:   fmt.Sprintf("symbols expected [0-1], actual [%v], error occurs while parsing id field %v", idStr, err),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	task, err, rowExists := db.FindTaskById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := ResponseModel[any]{
			IsSuccess: false,
			IsError:   true,
			Message:   err.Error(),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	if !rowExists {
		writer.WriteHeader(http.StatusNotFound)
		response := ResponseModel[any]{
			IsSuccess: false,
			Message:   "Task not found.",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := ResponseModel[db.Task]{
		IsSuccess: true,
		Message:   "Task found successfully.",
		Data:      &task,
	}
	json.NewEncoder(writer).Encode(response)
}

// GET    /tasks/              :  возвращает все задачи
func GetAllTasks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	tasks, err := db.GetAllTasks()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := ResponseModel[any]{
			IsSuccess: false,
			IsError:   true,
			Message:   fmt.Sprintf("Fail getting tasks from database: %v", err.Error()),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := ResponseModel[[]db.Task]{
		IsSuccess: true,
		Message:   "Tasks loading successfully.",
		Data:      &tasks,
	}
	json.NewEncoder(writer).Encode(response)
}

// GET    /tags/<tagname>      :  возвращает список задач с заданным тегом
func GetTasksByTags(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	tagsStr := mux.Vars(request)["tags"]
	tags := strings.Split(tagsStr, ",")

	tasks, err := db.FindTasksByTags(tags...)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := ResponseModel[any]{
			IsSuccess: false,
			IsError:   true,
			Message:   fmt.Sprintf("Fail getting tasks from database: %v", err.Error()),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	if len(tasks) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		response := ResponseModel[any]{ //? Нужна ли такая проверка или пустой возврат списка тоже 200-й код?
			IsSuccess: true,
			Message:   fmt.Sprintf("Tasks with tags [%v] not found in database", tagsStr),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := ResponseModel[[]db.Task]{
		IsSuccess: true,
		Message:   "Tasks loading successfully.",
		Data:      &tasks,
	}
	json.NewEncoder(writer).Encode(response)
}
