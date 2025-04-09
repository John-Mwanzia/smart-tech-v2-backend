package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Brand        string `json:"brand"`
	CompName     string `json:"compName"`
	Category     string `json:"category"`
	Slug         string `json:"slug"`
	ImageUrl     string `json:"imageUrl"`
	Price        int    `json:"price"`
	CountInStock int    `json:"countInStock"`
	Specs        string `json:"specs"`
}

type FeaturedProduct struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GadgetName   string `json:"Gadget_Name"`
	Category     string `json:"category"`
	Slug         string `json:"slug"`
	ImageUrl     string `json:"Img_Url"`
	Price        int    `json:"price"`
	CountInStock int    `json:"countInStock"`
	Specs        string `json:"specs"`
}
