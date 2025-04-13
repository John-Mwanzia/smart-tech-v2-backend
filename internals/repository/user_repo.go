package repository

import (
  "context"
  "time"

  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
  "github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

//find user by email
func FindUserByEmail(email string) (*models.User, error){
  ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  var user models.User 
  filter := bson.M{"email": email}
  err := config.DbInstance.UserCollection.FindOne(ctx, filter).Decode(&user)

  if err != nil {
    return nil, err
  }

  return &user, nil
}


func CreateUser(user models.User) (*models.User, error){
  ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  result,err := config.DbInstance.UserCollection.InsertOne(ctx, user)

  if err != nil {
    return nil, err
  }

  user.ID = result.InsertedID.(primitive.ObjectID)

  return &user, nil 
  }
