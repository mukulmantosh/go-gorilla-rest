package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	var count int64

	DB.First(&user, params["id"]).Count(&count)
	if count <= 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Not Found!"})
	} else {
		json.NewEncoder(w).Encode(user)
	}

}
