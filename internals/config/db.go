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
  userCollection  *mongo.Collection
  productsCollection *mongo.Collection
  featuredProductsCollection *mongo.Collection
  ordersCollection *mongo.Collection
  shippingCollection *mongo.Collection

}


var DbInstance *DB

func InitDB(){
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  clientOptions := options.Client().ApplyURI(AppEnv.mongodb_url)

  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal("Error Trying to connect to db", err)
  }

  db := client.Database("smartTechDB")

  DbInstance = &DB{
    Client: client,
    userCollection : db.Collection("users"),
    productsCollection : db.Collection("products"),
    featuredProductsCollection : db.Collection("featuredProducts"),
    ordersCollection : db.Collection("orders"),
    shippingCollection : db.Collection("shipping"),

  }

  fmt.Println("Connected to DB successfully")

}
