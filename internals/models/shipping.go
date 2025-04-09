package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shipping struct {
  ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  FullName string `json:"fullname"`
  PhoneNumber int `json:"phoneNumber"`
  Address string `json:"address"`
  City string `json:"city"`
  PostalCode string `json:"postalCode"`
  PaymentMethod string `json:"paymentMethod"`
 }

