package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olamilekan000/sosivio-app/auth-service/interfaces"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"
)

type Auth struct {
	Service interfaces.AuthService
}

func (s *Auth) Login(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var request repository.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		fmt.Println(err)
	}

	res, sErr := s.Service.Login(request)

	if sErr != nil {
		w.WriteHeader(sErr.Code)
		json.NewEncoder(w).Encode(sErr.AsMessage())
		return
	}

	w.Header().Add("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func (s *Auth) VerifyAuthToken(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	reqToken := r.Header.Get("Authorization")

	res, err := s.Service.VerifyAuthToken(reqToken)

	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.AsMessage())
		return
	}

	w.Header().Add("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Opps, this route does not exist", "statusCode": http.StatusNotFound})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
