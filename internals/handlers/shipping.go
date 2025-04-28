package handlers

import (
	"fmt"
	"net/http"

	middleware "github.com/John-Mwanzia/smart-tech-v2-backend/internals/middlewares"
)

func ShippingHandler(w http.ResponseWriter, r *http.Request)  {
	 userId := r.Context().Value(middleware.UserIDKey()).(string)

	 fmt.Println("User id :", userId)
	
}
