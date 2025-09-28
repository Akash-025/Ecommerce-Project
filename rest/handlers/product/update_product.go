package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/repo"
	"practice/utils"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImgaeUrl    string `json:"image_url"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	idx := r.PathValue("id")
	pid, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println("Atoi error", err)
		return
	}

	var updateProduct ReqUpdateProduct
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		fmt.Println("Decoding error", err)
		return
	}
	
	_, err = h.productRepo.Update(repo.Product{
		ID: pid,
		Title: updateProduct.Title,
		Description: updateProduct.Description,
		Price: updateProduct.Price,
		ImgaeUrl: updateProduct.ImgaeUrl,
	})
	if err != nil {
		http.Error(w, "plz give me valid json", http.StatusBadRequest)
		return
	}
	utils.SendData(w, "Product updated successfully", 201)
}
