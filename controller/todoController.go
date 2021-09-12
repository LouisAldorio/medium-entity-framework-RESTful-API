package controller

import (
	"encoding/json"
	"myapp/ent"
	"myapp/utils"
	"net/http"
	"strconv"

	"myapp/service"

	"github.com/gorilla/mux"
)

func TodoGetByIDController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	todo, err := service.NewTodoOps(r.Context()).TodoGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, todo)
}

func TodoCreateController(w http.ResponseWriter, r *http.Request) {

	var newTodo ent.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdTodo, err := service.NewTodoOps(r.Context()).TodoCreate(newTodo)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, createdTodo)
}

func TodoUpdateController(w http.ResponseWriter, r *http.Request) {

	var newTodoData ent.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodoData)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	newTodoData.ID = id

	updatedTodo, err := service.NewTodoOps(r.Context()).TodoUpdate(newTodoData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedTodo)
}

func TodoDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewTodoOps(r.Context()).TodoDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
