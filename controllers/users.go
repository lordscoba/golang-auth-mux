package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	// "io/ioutil"
	"github.com/lordscoba/golang-auth-mux/config"
	// "github.com/gorilla/mux"
	"github.com/lordscoba/golang-auth-mux/models"
	// "github.com/jinzhu/gorm"
)


// var db *gorm.DB


func init(){
	// config.Connect(){}
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.User{})
}

func SignUp(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result := db.Create(&user)
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)

	// // another way
	// CreateUser := &models.User{}
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil{
	// 	fmt.Println("could not get your data")
	// }
	// err = json.Unmarshal([]byte(body), CreateUser)
	// if err != nil{
	// 	fmt.Println("could not unmarshal your data")
	// }
	// result := db.Create(&CreateUser)

	// res, _ := json.Marshal(result)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)

}

func Login(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("HEllo welcome")
	w.Header().Set("Content-Type", "application/json")
	var users []models.User

	result := db.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
