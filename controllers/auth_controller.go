package controllers

import (
	"net/http"

	"encoding/json"
	"github.com/christiangda/gogo-rest-api/models"
	"github.com/christiangda/gogo-rest-api/services"
)

func Login(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	whoIsComing := new(models.Credential)
	err := decoder.Decode(&whoIsComing)
	if err != nil {
		w.Write([]byte("Error decoding request, isent a valif JSON or "))
	}

	auth_service, err := services.Authenticate(whoIsComing)
	if err != nil {
		w.Write([]byte("Error validating User"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(auth_service.Token)
}
