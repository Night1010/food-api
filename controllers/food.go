package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"project/models"
	"project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

type FoodInput struct {
	Name      string `json:"name" validate:"required"`
	Available bool   `json:"available" `
}

func GetAllFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var foods []models.Food
	models.DB.Find(&foods)

	json.NewEncoder(w).Encode(foods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var food models.Food

	if err := models.DB.Where("id = ?", id).First(&food).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "food not found")
		return
	}

	json.NewEncoder(w).Encode(food)
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
	var input FoodInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		log.Fatal(err)
		return
	}

	food := &models.Food{
		Name:      input.Name,
		Available: input.Available,
	}

	models.DB.Create(food)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(food)

}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var food models.Food

	if err := models.DB.Where("id = ?", id).First(&food).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "food not found")
		return
	}

	var input FoodInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	food.Name = input.Name
	food.Available = input.Available

	models.DB.Save(&food)

	json.NewEncoder(w).Encode(food)
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var food models.Food

	if err := models.DB.Where("id = ?", id).First(&food).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Food not found")
		return
	}

	models.DB.Delete(&food)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(food)
}
