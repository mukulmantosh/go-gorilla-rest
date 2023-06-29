package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("User is deleted successfully")
}
