package main

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/joho/godotenv"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)


var userCollection *mongo.Collection

func initDB(){

  MONGO_URI := os.Getenv("MONGODB_URL")

  ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  clientOptions := options.Client().ApplyURI(MONGO_URI)
  client,err := mongo.Connect(ctx,clientOptions)

  if err != nil {
    log.Fatal("Failed to Connect to DB")
  }

  fmt.Println("Connected to db successfully")

  db := client.Database("smartTechDB")

  userCollection = db.Collection("users")

}


func main(){
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Failed to Load env file")
  }

  port := os.Getenv("PORT")
  mux := http.NewServeMux()

  initDB()

  fmt.Println("Server started on port:",port)
  http.ListenAndServe(port, mux)
}
