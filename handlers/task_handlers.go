package handlers

import (
	"encoding/json"
	"github.com/blr-coder/m_tasker/interfaces"
	"github.com/blr-coder/m_tasker/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateTask(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		task := models.Task{}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		}

		err = json.Unmarshal(body, &task)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		}

		res, err := db.Create(task)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func GetTask(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		res, err := db.Get(id)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func AllTasks(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		res, err := db.List()
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func AllTasksWithFilter(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var decoder = schema.NewDecoder()

		var taskFilter models.TaskFilter
		err := decoder.Decode(&taskFilter, request.URL.Query())
		if err != nil {
			log.Println("Error in GET parameters : ", err)
		}

		res, err := db.Search(&taskFilter)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func DoneTask(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		res, err := db.Done(id)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func UpdateTask(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		task := models.Task{}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		}

		err = json.Unmarshal(body, &task)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		}

		res, err := db.Update(id, &task)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)

	}
}

func DeleteTask(db interfaces.TaskInterface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]

		res, err := db.Delete(id)
		if err != nil {
			WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(writer, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(res)
}
