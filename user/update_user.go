package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("IN UPDATE")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
