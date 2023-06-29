package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user"
)

func initRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	user.InitMigration()
	initRouter()
}
