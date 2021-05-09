package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Avatar       string `json:"avatar"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=restapitest port=5432"
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection success")
	}

	db.Debug().AutoMigrate(&User{})

	handleRequest()
}

func handleRequest() {
	log.Println("Start the development server at http://127.0.0.1:4000")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/users", createUser).Methods("POST")
	myRouter.HandleFunc("/api/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(payloads, &user)

	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := Result{Code: 200, Data: user, Message: "User created successfully"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	db.Find(&users)
	res := Result{Code: 200, Data: users, Message: "Users loaded successfully."}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
