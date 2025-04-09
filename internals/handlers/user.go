package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
)

func SignIn(w http.ResponseWriter, r *http.Request){
  var user models.User

  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
    http.Error(w, "Invalid Json Format", http.StatusBadRequest)
    return 
  }

  //check if the user exists 
}

  
