package product

import (
	"net/http"
	"practice/database"
	"practice/utils"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)
	if err != nil {
		return
	}

	product := database.Get(pid)
	if product == nil {
		utils.SendData(w, "Product Not Found", 404)
		return
	}
	utils.SendData(w, product, 200)

}
