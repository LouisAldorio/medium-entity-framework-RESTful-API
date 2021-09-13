package router

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerUserRouter(r *mux.Router) {

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", controller.UserGetAllController).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", controller.UserGetByIDController).Methods(http.MethodGet)
	userRouter.HandleFunc("/", controller.UserCreateController).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", controller.UserUpdateController).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", controller.UserDeleteController).Methods(http.MethodDelete)
}
