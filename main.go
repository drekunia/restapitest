package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB
var err error

type User struct {
	ID           uint       `json:"id" sql:"primary key;not null; unique"`
	Username     string     `json:"username" sql:"not null; unique"`
	Email        string     `json:"email" sql:"not null; unique"`
	PasswordHash string     `json:"password_hash" sql:"not null"`
	FirstName    string     `json:"first_name" sql:"not null"`
	LastName     string     `json:"last_name" sql:"not null"`
	City         string     `json:"city"`
	Country      string     `json:"country"`
	Avatar       string     `json:"avatar"`
	Bio          string     `json:"bio"`
	CreatedAt    time.Time  `json:"created_at" sql:"not null"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func init() {
	dbHost := os.Getenv("APP_DB_HOST")
	dbName := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("APP_DB_PORT")
	dbUsername := os.Getenv("APP_DB_USERNAME")
	dbPassword := os.Getenv("APP_DB_PASSWORD")

	dbUri := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbPort, dbUsername, dbPassword)

	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&User{})
}

func main() {
	handleRequest()
}

func handleRequest() {
	log.Println("Start the development server at http://127.0.0.1:4000")
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/users", createUser).Methods("POST")
	myRouter.HandleFunc("/api/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/users/deleted", getUsersUnscoped).Methods("GET")
	myRouter.HandleFunc("/api/users/deleted/{id}", getUserUnscoped).Methods("GET")
	myRouter.HandleFunc("/api/users/deleted/{id}", deleteUserUnscoped).Methods("DELETE")
	myRouter.HandleFunc("/api/users/deleted/{id}", restoreUserUnscoped).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(payloads, &user)

	db.Create(&user)

	res := Result{Code: 200, Data: &user, Message: "User created successfully"}
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
	res := Result{Code: 200, Data: &users, Message: "Users loaded successfully"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User

	db.First(&user, userId)
	res := Result{Code: 200, Data: &user, Message: "User loaded successfully"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var userUpdates User
	json.Unmarshal(payloads, &userUpdates)

	var user User

	db.First(&user, userId)
	db.Model(&user).Update(userUpdates)

	res := Result{Code: 200, Data: &user, Message: "User updated successfully"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User

	db.First(&user, userId)
	db.Delete(&user)

	res := Result{Code: 200, Data: &user, Message: "User deleted successfully"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getUsersUnscoped(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	db.Unscoped().Find(&users)
	res := Result{Code: 200, Data: &users, Message: "Deleted users loaded successfully"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getUserUnscoped(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User

	db.Unscoped().First(&user, userId)
	res := Result{Code: 200, Data: &user, Message: "Deleted user loaded successfully"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteUserUnscoped(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User

	db.First(&user, userId)
	db.Unscoped().Delete(&user)

	res := Result{Code: 200, Data: &user, Message: "User permanently deleted"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func restoreUserUnscoped(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User

	db.Unscoped().First(&user, userId)
	user.DeletedAt = nil
	db.Save(&user)

	res := Result{Code: 200, Data: &user, Message: "Deleted user restored successfully"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
