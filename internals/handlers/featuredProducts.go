package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FeaturedProducts(w http.ResponseWriter, r *http.Request){
	pathStr := strings.Split(strings.Trim(r.URL.Path, "/"),"/")

	switch r.Method{
	case http.MethodGet :
		switch len(pathStr){
		case 1:
			switch pathStr[0]{
			case "":
				getFeaturedProducts(w, r)
			case "search":
				searchFeaturedProducts(w, r)
			case "categories":
				getFeaturedCategories(w, r)

			default:
				http.Error(w, "Path Not Found ", http.StatusNotFound)

			}
		case 2:
			switch pathStr[0]{
			case "id":
				getFeaturedProductById(w,r, pathStr[1])
			case "slug":
				getFeaturedProductBySlug(w, r ,pathStr[1])
				default :
				http.Error(w, "Path Not Found ", http.StatusNotFound)
			}

		}
		default :
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	}
}



func getFeaturedProducts(w http.ResponseWriter,  _*http.Request){

	products, err := repository.FindFeaturedProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode the products", http.StatusInternalServerError)
		return 
	}

}

func searchFeaturedProducts(w http.ResponseWriter, r *http.Request){
}

func getFeaturedCategories(w http.ResponseWriter, r *http.Request){

}

func getFeaturedProductById(w http.ResponseWriter, _*http.Request, id string){
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Not a valid object id ", http.StatusBadRequest)
		return
	}
	product, err := repository.FindFeaturedProductById(objId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	if 	err = json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode product", http.StatusInternalServerError)
		return
	}

}

func getFeaturedProductBySlug(w http.ResponseWriter, _*http.Request, slug string ){
	product, err := repository.FindFeaturedProductBySlug(slug)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	 if err = json.NewEncoder(w).Encode(product); err != nil {
     http.Error(w, "Failed to encode product ", http.StatusInternalServerError)
	 }

}




















