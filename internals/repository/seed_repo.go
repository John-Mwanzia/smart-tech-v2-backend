package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"go.mongodb.org/mongo-driver/bson"
)

func ClearCollections() error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  _, err := config.DbInstance.ProductsCollection.DeleteMany(ctx, bson.M{}) 
  if err != nil {
    return fmt.Errorf("failed to delete products : %w", err)
  }

  _, err = config.DbInstance.FeaturedProductsCollection.DeleteMany(context.Background(), bson.M{})
  if err != nil{
    return fmt.Errorf("failed to delete featured products products %w ", err)
  }

  _, err = config.DbInstance.UserCollection.DeleteMany(context.Background(), bson.M{})
  if err != nil{
    return  fmt.Errorf("failed to delete users %w", err)
  }
  return nil 

}  


func SeedInitialdata(users []interface{}, products []interface{}, fps []interface{}) error {
  _, err := config.DbInstance.UserCollection.InsertMany(context.Background(), users)

  if err != nil {
    return err
  }

  _, err = config.DbInstance.ProductsCollection.InsertMany(context.Background(), products)

  if err != nil {
    return err
  }

  _, err = config.DbInstance.FeaturedProductsCollection.InsertMany(context.Background(), fps)

  if err != nil {
    return err
  }


  return nil
}
