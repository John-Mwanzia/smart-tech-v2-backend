package repository

import (
  "context"
  "fmt"
  "time"

  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
  "go.mongodb.org/mongo-driver/bson"
)

func FetchProducts() (*[]models.Product, error) {
  var products []models.Product
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  cursor, err := config.DbInstance.ProductsCollection.Find(ctx, bson.M{})

  if err != nil {
    return nil, fmt.Errorf("failed to fetch products") 
  }

  for cursor.Next(ctx){
    var product models.Product

    if err := cursor.Decode(&product); err != nil {
      return nil, fmt.Errorf("failed to decode products")
    }

    products = append(products, product)
  }


  return &products,nil 



}




