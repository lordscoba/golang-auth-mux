package routes

import (
	"github.com/gorilla/mux"
	"github.com/lordscoba/golang-auth-mux/controllers"

)

var AuthRoutes = func(router *mux.Router) {

	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
}

var AdminRoutes = func(router *mux.Router) {

	router.HandleFunc("/admin/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/admin/update/{Id}", controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/admin", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/admin/show/{Id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/admin/delete/{Id}", controllers.DeleteUser).Methods("DELETE")
}
