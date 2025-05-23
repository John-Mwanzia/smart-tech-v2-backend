package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
  ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  Password string `json:"password"`
  IsAdmin bool `json:"isAdmin"`
}
type ReturnSafeUser struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Email   string             `json:"email"`
	IsAdmin bool               `json:"isAdmin"`
}

