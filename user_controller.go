package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var user User

func createUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(payloads, &user)

	db.Create(&user)

	writeResponse(&user, w)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	db.Find(&users)

	writeResponse(&users, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	db.First(&user, userId)

	writeResponse(&user, w)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	userUpdates := user
	json.Unmarshal(payloads, &userUpdates)

	db.First(&user, userId)
	db.Model(&user).Update(userUpdates)

	writeResponse(&user, w)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	db.First(&user, userId)
	db.Delete(&user)

	writeResponse(&user, w)
}
