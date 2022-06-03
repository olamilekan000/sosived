package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/olamilekan000/sosivio-app/auth-service/handlers"
	"github.com/olamilekan000/sosivio-app/auth-service/helpers"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"
	"github.com/olamilekan000/sosivio-app/auth-service/services"
)

func envChecks() {
	envProps := []string{
		"JWT_SECRET",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatalf("Environment variable %s value is missing", k)
		}
	}
}

func Start() {
	fmt.Println("Starting Auth Service")

	envChecks()

	router := mux.NewRouter()

	repo := repository.NewUserRepository()
	ser := services.NewAuthService(repo, &helpers.JWTStruct{})

	h := handlers.Auth{Service: ser}
	notFoundHanler := http.HandlerFunc(handlers.NotFound)
	router.NotFoundHandler = notFoundHanler

	router.HandleFunc("/auth/login", h.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/verify", h.VerifyAuthToken).Methods("GET", "OPTIONS")

	fmt.Println("Starting applicaion at port: 8989")

	_ = http.ListenAndServe(":8989", router)

}
