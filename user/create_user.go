package user

import (
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	r.Body.Close()
	var user User
	json.Unmarshal(reqBody, &user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
