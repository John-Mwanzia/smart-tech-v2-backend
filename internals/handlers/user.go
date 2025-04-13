package handlers

import (
  "encoding/json"
  "net/http"
  "strings"

  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/services"
)

func UserHandler(w http.ResponseWriter, r *http.Request){
  path := strings.Trim(r.URL.Path, "/") 

  switch r.Method {
  case http.MethodPost:
    switch path {
    case "signup":
      registerHandler(w,r)
    case "signin":
      signInHandler(w,r)
    default:
      http.Error(w, "Path not found", http.StatusNotFound)
    }

  default:
    http.Error(w, "Invalid request Method", http.StatusMethodNotAllowed)
  }
}


func registerHandler(w http.ResponseWriter, r *http.Request){
  var user models.User 

  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    http.Error(w, "Invalid request Json", http.StatusBadRequest)
    return
  }

  //check for empty fields

  if user.Name == "" || user.Email == "" || user.Password == "" {
    http.Error(w, "All input fields are Required", http.StatusBadRequest)
    return 
  }

  createdUser, err := services.RegisterUser(user) 

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return 
  }

  newUser := models.ReturnSafeUser{
    ID: createdUser.ID,
    Name: createdUser.Name,
    Email: createdUser.Email,
    IsAdmin: createdUser.IsAdmin,
  }

  response := map[string]interface{}{
    "success": true,
    "message": "Registered user successfully",
    "user": newUser,
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(response)


}




func signInHandler(w http.ResponseWriter, r *http.Request) {


}


