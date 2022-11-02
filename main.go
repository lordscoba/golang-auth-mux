package main

import (
	"os"
	"fmt"
	"log"
	"errors"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lordscoba/golang-auth-mux/routes"
)


func main() {

	// loading the .env file
	err := godotenv.Load(".env")

		if err != nil {
			log.Println("Error loading .env file")
		}
		port := os.Getenv("PORT")
		if port == "" {
			port = "9000"
		}



	// loading the mux routers
	r := mux.NewRouter()
	routes.AuthRoutes(r)
	routes.AdminRoutes(r)
	http.Handle("/", r)


	// log.Fatal(http.ListenAndServe("localhost:9081", r))
	err1 := http.ListenAndServe(":"+port, nil)
  if errors.Is(err1, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err1 != nil {
		fmt.Printf("error starting server: %s\n", err1)
		os.Exit(1)
	}
}
