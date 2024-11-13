package handlers

import (
	"encoding/json"
	"net/http"
	"tasklist_REST_API/internal/model/db"
	"time"
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
