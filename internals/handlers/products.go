package handlers

import (
  "net/http"
  "strings"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request){
  path := strings.Trim(r.URL.Path, "/")

  switch r.Method {
  case http.MethodGet:
    switch path{
    case "" :
      getProducts(w,r)
    case "/search":
      searchProducts(w,r)
    case "/categories":
      getCategories(w, r)
    case "/slug/:slug":
      getProductBySlug(w,r)
    case "id" :
      getProductById(w,r)

    } 
  }

}


func getProducts(w http.ResponseWriter, r *http.Request){

}



func searchProducts(w http.ResponseWriter, r *http.Request){
}


func getCategories(w http.ResponseWriter, r *http.Request){
}



func getProductBySlug(w http.ResponseWriter, r *http.Request){
}



func getProductById(w http.ResponseWriter, r *http.Request){
}






