package product

import (
	"net/http"
	"practice/utils"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		return
	}

	product, err := h.productRepo.Get(pid)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if product == nil {
		utils.SendData(w, "Product Not Found", 404)
		return
	}
	utils.SendData(w, product, 200)

}
