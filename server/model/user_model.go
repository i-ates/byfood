package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Mail string             `bson:"mail" json:"mail"`
	Age  int                `bson:"age" json:"age"`
}
