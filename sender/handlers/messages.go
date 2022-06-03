package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/olamilekan000/sosivio-app/auth-service/interfaces"
)

type Sender struct {
	Service interfaces.Messenger
}

func (s *Sender) SendMessage(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	reqToken := r.Header.Get("Authorization")
	res, sErr := s.Service.SendMessage(reqToken)

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

func SayHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Opps, this route does not exist", "statusCode": http.StatusNotFound})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
