package handlers

import (
	"net/http"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/services"
)

func SeedHandler(w http.ResponseWriter , r *http.Request){
 if r.Method != http.MethodPost {
   http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
   return
 }

 //seedData
 err := services.SeedCollections()
 if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
 }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Seeded successfully"))
}  
