package main

import (
	"context"
	"fmt"
	"github.com/blr-coder/m_tasker/database"
	"github.com/blr-coder/m_tasker/handlers"
	"github.com/blr-coder/m_tasker/repositories"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("m_tasker start!")

	ctx := context.TODO()

	db := database.ConnectDB(ctx)
	collection := db.Collection("tasks")

	client := &repositories.TaskClient{
		Ctx:        ctx,
		Collection: collection,
	}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", handlers.CreateTask(client)).Methods(http.MethodPost)
	router.HandleFunc("/all_tasks", handlers.AllTasks(client)).Methods(http.MethodGet)
	//router.HandleFunc("/tasks", handlers.AllTasksWithFilter(client)).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", handlers.GetTask(client)).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}/done", handlers.DoneTask(client)).Methods(http.MethodPatch)
	//router.HandleFunc("/tasks/{id}", handlers.UpdateTask(client)).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask(client)).Methods(http.MethodDelete)
	_ = http.ListenAndServe(":8090", router)
}
