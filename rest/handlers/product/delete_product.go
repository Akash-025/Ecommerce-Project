package product

import (
	"fmt"
	"net/http"
	"practice/database"
	"practice/utils"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	idx := r.PathValue("id")
	pid, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println("Atoi error", err)
		return
	}

	database.Delete(pid)
	utils.SendData(w, "Product deleted successfully", 201)
}
