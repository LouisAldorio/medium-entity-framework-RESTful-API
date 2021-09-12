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

func UserGetAllController(w http.ResponseWriter, r *http.Request) {

	users, err := service.NewUserOps(r.Context()).UsersGetAll()
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, users)
}

func UserGetByIDController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	user, err := service.NewUserOps(r.Context()).UserGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, user)
}

func UserCreateController(w http.ResponseWriter, r *http.Request) {

	var newUser ent.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	user, err := service.NewUserOps(r.Context()).UserCreate(newUser)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, user)
}

func UserUpdateController(w http.ResponseWriter, r *http.Request) {

	var newUserData ent.User
	err := json.NewDecoder(r.Body).Decode(&newUserData)
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
	newUserData.ID = id

	updatedUser, err := service.NewUserOps(r.Context()).UserUpdate(newUserData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedUser)
}

func UserDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewUserOps(r.Context()).UserDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}