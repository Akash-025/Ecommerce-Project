package product

import (
	"fmt"
	"net/http"
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

	err = h.svc.Delete(pid)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, "Product deleted successfully", 201)
}
