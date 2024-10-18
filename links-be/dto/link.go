package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Link string             `json:"link" bson:"link"`
}
