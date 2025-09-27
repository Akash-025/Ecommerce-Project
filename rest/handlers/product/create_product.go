package product

import (
	"encoding/json"
	"net/http"
	"practice/database"
	"practice/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}
	
	createdProduct := database.Stor(newProduct)

	utils.SendData(w, createdProduct, 201)
}
