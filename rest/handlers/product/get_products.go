package product

import (
	"fmt"
	"net/http"
	"practice/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	productList, err := h.productRepo.List()
	fmt.Println("Get:", productList)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, productList, 200)

}
