package product

import (
	"encoding/json"
	"net/http"
	"practice/repo"
	"practice/utils"
)

type ReqCreateProduct struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	ImageUrl    string `json:"image_url" db:"image_url"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title: newProduct.Title,
		Description: newProduct.Description,
		Price: newProduct.Price,
		ImageUrl: newProduct.ImageUrl,
	})
	
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, createdProduct, 201)
}
