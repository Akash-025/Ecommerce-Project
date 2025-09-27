package handlers

import (
	"net/http"
	"practice/database"
	"practice/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), 200)

}
