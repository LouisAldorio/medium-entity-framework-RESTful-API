package router

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerTodoRouter(r *mux.Router) {

	todoRouter := r.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("/{id}", controller.TodoGetByIDController).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", controller.TodoCreateController).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id}", controller.TodoUpdateController).Methods(http.MethodPut)
	todoRouter.HandleFunc("/{id}", controller.TodoDeleteController).Methods(http.MethodDelete)
}
