package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Book   string             `json:"title" bson:"title,omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
}
