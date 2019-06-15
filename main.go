package main

import (
	"classtime/app"
	"classtime/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// users
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", controllers.GetUser).Methods("GET")

	// disciplines
	router.HandleFunc("/discipline/create", controllers.CreateDiscipline).Methods("POST")
	router.HandleFunc("/discipline/{id:[0-9]+}", controllers.UpdateDiscipline).Methods("PUT")
	router.HandleFunc("/disciplines", controllers.GetDisciplines).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})

	fmt.Println(port)

	//Launch the app, visit localhost:8000/api
	err := http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
	if err != nil {
		fmt.Print(err)
	}
}
