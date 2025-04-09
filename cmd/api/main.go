package main

import (
	"fmt"
	"net/http"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/handlers"
)


func main(){
  config.LoadEnv()
  config.InitDB() 

  
  mux := http.NewServeMux()
  
  mux.Handle("/api/seed", http.StripPrefix("/api/seed", handlers.SeedHandler))

  // http.StripPrefix("/api/seed", seedHandler);
  // http.StripPrefix("/api/users", userHandler)
  // http.StripPrefix("/api/products", productsHandler);
  // http.StripPrefix("/api/featuredProducts", featuredProductsHandler);
  // http.StripPrefix("/api/lipaNaMpesa", lipaNaMpesaHandler);
  // http.StripPrefix("/api/shipping", protect, shippingHandler);
  // http.StripPrefix("/api/checkout", stripeHandler);
  // http.StripPrefix("/api/order", protect, orderHandler);

  // fmt.Println("Server started on port:",config.AppEnv.port)
  // http.ListenAndServe(port, mux)
}
