package product

import (
	"net/http"
	"practice/database"
	"practice/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), 200)

}
