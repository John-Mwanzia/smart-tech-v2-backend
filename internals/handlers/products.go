package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request){
	path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	switch r.Method {
	case http.MethodGet:
		switch len(path){
		case 1:
			switch path[0]{
			case "":
				getProducts(w, r)
			case "search":
				searchProducts(w,r)
			case "categories":
				getCategories(w, r)
			default:
				http.Error(w, "Path not found", http.StatusNotFound)
			}
		case 2:
			switch path[0]{
			case "id":
				getProductById(w,r,path[1])
			case "slug":
				getProductBySlug(w, r, path[1])
			default:
				http.Error(w, "path not found", http.StatusNotFound)
			}

		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}




func getProducts(w http.ResponseWriter, _*http.Request){

	products, err := repository.FetchProducts() 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)

	if err != nil {
		http.Error(w,"Failed to Decode Response", http.StatusInternalServerError)
		return
	}


}



func searchProducts(w http.ResponseWriter, r *http.Request){
}


func getCategories(w http.ResponseWriter, _*http.Request){
	categories, err := repository.FindCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Failed to decode categories", http.StatusInternalServerError)
		return 
	}

}



func getProductBySlug(w http.ResponseWriter, _*http.Request, slug string){

	product, err := repository.FindProductBySlug(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)

	if err != nil {
		http.Error(w,"Failed to Decode Response", http.StatusInternalServerError)
		return
	}

}



func getProductById(w http.ResponseWriter, _*http.Request, id string){

	objId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
		http.Error(w, "failed to convert hex string id to object id", http.StatusInternalServerError)
		return
	}

	product, err := repository.FindProductById(objId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	err =	json.NewEncoder(w).Encode(product)

	if err != nil {
		http.Error(w,"Failed to Decode Response", http.StatusInternalServerError)
		return
	}

}






