package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
			Success: false,
			Message: "provide json file is invalid, look at sample request model.",
			Data:    nil,
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
			Success: false,
			Error:   true,
			Message: err.Error(),
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	type taskId struct {
		TaskId int `json:"task_id"`
	}

	writer.WriteHeader(http.StatusOK) // 200 error
	response := ResponseModel[taskId]{
		Success: true,
		Message: "Task added successfully",
		Data:    &taskId{TaskId: id},
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
			Success: false,
			Error:   true,
			Message: fmt.Sprintf("symbols expected [0-1], actual [%v], error occurs while parsing id field %v", idStr, err),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	task, err, rowExists := db.FindTaskById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := ResponseModel[any]{
			Success: false,
			Error:   true,
			Message: err.Error(),
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	if !rowExists {
		writer.WriteHeader(http.StatusNotFound)
		response := ResponseModel[any]{
			Success: false,
			Message: "Task not found.",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.WriteHeader(http.StatusNotFound)
	response := ResponseModel[db.Task]{
		Success: true,
		Message: "Task found successfully.",
		Data:    &task,
	}
	json.NewEncoder(writer).Encode(response)

}
