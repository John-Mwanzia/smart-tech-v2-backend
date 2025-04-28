package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func FindFeaturedProducts() ([]models.FeaturedProduct, error ) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var featuredProducts []models.FeaturedProduct

	cursor,err := config.DbInstance.FeaturedProductsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil , fmt.Errorf("failed to fectch featured products")
	}

	for	cursor.Next(ctx){
		var product models.FeaturedProduct
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("failed to decode featured product")
		}

		featuredProducts = append(featuredProducts, product)
	}

	return featuredProducts , nil

}

func FindFeaturedProductById(id primitive.ObjectID)(*models.FeaturedProduct, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.FeaturedProduct

	filter := bson.M{"_id": id}

  err := config.DbInstance.FeaturedProductsCollection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch product")
	}

	return &product, nil 
}




func FindFeaturedProductBySlug(slug string) (*models.FeaturedProduct, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
  var product models.FeaturedProduct
	
	filter := bson.M{"slug": slug}

	 err := config.DbInstance.FeaturedProductsCollection.FindOne(ctx, filter).Decode(&product)

	 if err != nil {
		 if err == mongo.ErrNoDocuments {
			 return nil, fmt.Errorf("no product found")
		 }
		 return nil, fmt.Errorf("Error trying to Fetch product")
	 }

	 return &product, nil 
 }



 func FindFeaturedProductsCategories()(any, error){
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := config.DbInstance.FeaturedProductsCollection.Distinct(ctx, "category", bson.M{})

  if err != nil {
		return nil, fmt.Errorf("failed to fetch categories")
	}

	return result, nil 

}





















