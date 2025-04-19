package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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


func FindProductById(id primitive.ObjectID) (*models.Product, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() 

	filter := bson.M{"_id": id}

	var product models.Product

	err := config.DbInstance.ProductsCollection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil 
}


func FindProductBySlug(slug string) (*models.Product, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() 

	filter := bson.M{"slug": slug}

	var product models.Product

	err := config.DbInstance.ProductsCollection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil 
}
