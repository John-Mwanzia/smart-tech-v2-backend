package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	Name  string `bson:"name" json:"name"`
	Qty   int    `bson:"qty" json:"qty"`
	Image string `bson:"image" json:"image"`
	Price int    `bson:"price" json:"price"`
}

type ShippingAddress struct {
	FullName   string `bson:"fullName" json:"fullName"`
	Address    string `bson:"address" json:"address"`
	City       string `bson:"city" json:"city"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
	Country    string `bson:"country" json:"country"`
}

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderItems      []OrderItem        `bson:"orderItems" json:"orderItems"`
	ShippingAddress ShippingAddress    `bson:"shippingAddress" json:"shippingAddress"`
	PaymentMethod   string             `bson:"paymentMethod" json:"paymentMethod"`
	ItemsPrice      int                `bson:"itemsPrice" json:"itemsPrice"`
	ShippingPrice   int                `bson:"shippingPrice" json:"shippingPrice"`
	TaxPrice        int                `bson:"taxPrice" json:"taxPrice"`
	TotalPrice      int                `bson:"totalPrice" json:"totalPrice"`
	UserID          primitive.ObjectID `bson:"user" json:"user"` 
	IsPaid          bool               `bson:"isPaid" json:"isPaid"`
	PaidAt          *time.Time         `bson:"paidAt,omitempty" json:"paidAt,omitempty"`
	IsDelivered     bool               `bson:"isDelivered" json:"isDelivered"`
	DeliveredAt     *time.Time         `bson:"deliveredAt,omitempty" json:"deliveredAt,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
}

