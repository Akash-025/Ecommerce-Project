package handlers

import (
	"encoding/json"
	"net/http"
	"practice/database"
	"practice/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}
	
	createdUser := newUser.Store()

	utils.SendData(w, createdUser, 201)
}
