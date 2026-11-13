package product

import (
	"command/database"
	"command/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var Odata []models.ProductDetails
	err := json.NewDecoder(r.Body).Decode(&Odata)
	if err != nil {
		http.Error(w, "error in request decoding", 500)
		return
	}
	err = database.Db.Create(&Odata).Error
	if err != nil {
		fmt.Fprint(w, err.Error())
		http.Error(w, "record not created in database", 500)
		return
	}
	fmt.Fprint(w, "Product created successfully...!")
}

func GetProductDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProductDetails called")
	var products []models.ProductDetails

	result := database.Db.Find(&products)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var Odata models.ProductDetails
	err := database.Db.Model(&Odata).Where("id = ?", id).UpdateColumn("name", "Horlicks").Error
	if err != nil {
		http.Error(w, "product name not updated in database..!", 500)
		return
	}
	fmt.Fprint(w, Odata)
}
