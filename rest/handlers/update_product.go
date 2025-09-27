package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/database"
	"practice/utils"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	idx := r.PathValue("id")
	pid, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println("Atoi error", err)
		return
	}

	var updateProduct database.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		fmt.Println("Decoding error", err)
		return 
	}
	updateProduct.ID = pid
	database.Update(updateProduct)
	utils.SendData(w, "Product updated successfully", 201)
}
