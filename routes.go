package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	log.Println("Start the development server at http://127.0.0.1:4000")
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/api", homePage)
	myRouter.HandleFunc("/api/users", createUser).Methods("POST")
	myRouter.HandleFunc("/api/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}
