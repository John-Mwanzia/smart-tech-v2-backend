package main

import (
	"fmt"
	"net/http"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/handlers"
	middleware "github.com/John-Mwanzia/smart-tech-v2-backend/internals/middlewares"
)


func main(){
  config.LoadEnv()
  config.InitDB() 

  mux := http.NewServeMux()
  
  mux.HandleFunc("/api/seed", handlers.SeedHandler)
  mux.Handle("/api/users/", http.StripPrefix("/api/users",http.HandlerFunc(handlers.UserHandler)))
  mux.Handle("/api/products/", http.StripPrefix("/api/products", http.HandlerFunc(handlers.ProductsHandler)))
  mux.Handle("/api/featuredProducts/", http.StripPrefix("/api/featuredProducts", http.HandlerFunc(handlers.FeaturedProducts)));
	mux.Handle("/api/shipping/", middleware.JwtAuthMiddleware(http.HandlerFunc(handlers.ShippingHandler)))
  // http.StripPrefix("/api/lipaNaMpesa", lipaNaMpesaHandler);
  // http.StripPrefix("/api/shipping", protect, shippingHandler);
  // http.StripPrefix("/api/checkout", stripeHandler);
  // http.StripPrefix("/api/order", protect, orderHandler);

  fmt.Println("Server started on port:", config.AppEnv.Port)
  if err := http.ListenAndServe(":"+config.AppEnv.Port, mux); err != nil {
    fmt.Println("Server error:", err)
}

}

