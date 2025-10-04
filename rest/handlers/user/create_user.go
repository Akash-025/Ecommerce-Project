package user

import (
	"encoding/json"
	"net/http"
	"practice/domain"
	"practice/utils"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser ReqCreateUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}

	createdUser, err := h.svc.Create(domain.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}

	utils.SendData(w, createdUser, 201)
}
