package handlers

import (
	"encoding/json"
	"net/http"
	"practice/database"
	"practice/utils"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var reqUser struct{
		email string
		pass string
	}
	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}
	
	usr := database.Find(reqUser.email, reqUser.pass)
	if usr == nil{
		http.Error(w, "Invalid credential the login user", 400)
		return
	}

	utils.SendData(w, usr, 201)
}
