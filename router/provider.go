package router

import "github.com/gorilla/mux"

func RegisterRouter(r *mux.Router) {
	registerTodoRouter(r)
	registerUserRouter(r)
}