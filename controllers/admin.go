package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	// "io/ioutil"
	"github.com/lordscoba/golang-auth-mux/config"
	"github.com/gorilla/mux"
	"github.com/lordscoba/golang-auth-mux/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


func init(){
	// config.Connect(){}
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.User{})
}


func checkIfUserExists(userId string) bool {
	var user models.User
	db.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}



func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("Admin Create User")

	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result := db.Create(&user)
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}


func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	fmt.Println("Hello update users")
	userId := mux.Vars(r)["Id"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	db.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)

}


func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	result := db.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}


func GetUserById(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
	userId := mux.Vars(r)["Id"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var user models.User
	db.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["Id"]
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var user models.User
	db.Delete(&user, userId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}