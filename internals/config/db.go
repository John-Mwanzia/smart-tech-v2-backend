package config

import (
  "context"
  "fmt"
  "log"
  "time"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
  Client *mongo.Client
  UserCollection  *mongo.Collection
  ProductsCollection *mongo.Collection
  FeaturedProductsCollection *mongo.Collection
  OrdersCollection *mongo.Collection
  ShippingCollection *mongo.Collection

}


var DbInstance *DB

func InitDB(){
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  clientOptions := options.Client().ApplyURI(AppEnv.Mongodb_url)

  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal("Error Trying to connect to db", err)
  }

  db := client.Database("smartTechDB")

  DbInstance = &DB{
    Client: client,
    UserCollection : db.Collection("users"),
    ProductsCollection : db.Collection("products"),
    FeaturedProductsCollection : db.Collection("featuredProducts"),
    OrdersCollection : db.Collection("orders"),
    ShippingCollection : db.Collection("shipping"),

  }

  fmt.Println("Connected to DB successfully")

}
