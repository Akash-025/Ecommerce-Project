package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/config"
	"practice/database"
	"practice/utils"
)

type ReqUser struct {
	Email string `json:"email"`
	Password  string `json:"password"`
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var reqUser ReqUser
	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}

	fmt.Println("req:",reqUser)

	usr := database.Find(reqUser.Email, reqUser.Password)
	if usr == nil {
		http.Error(w, "Invalid credential the login user", 400)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreatJwt(cnf.JwtSecretKey, utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return

	}

	utils.SendData(w, accessToken, 201)
}
