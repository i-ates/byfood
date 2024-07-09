package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"task4/model"
)

// DB defines the interface for interacting with the database.
type DB interface {
	GetUsers() ([]*model.UserModel, error)
	AddUser(user *model.UserModel) error
	UpdateUser(user *model.UserModel) error
	DeleteUser(userID primitive.ObjectID) error
}

// MongoDB is the implementation of the DB interface for MongoDB.
type MongoDB struct {
	collection *mongo.Collection
}

// NewMongo creates a new MongoDB instance and returns it as a DB interface.
func NewMongo(client *mongo.Client) DB {
	tech := client.Database("tech").Collection("tech")
	return MongoDB{collection: tech}
}

// GetUsers retrieves a list of all users from the database.
func (m MongoDB) GetUsers() ([]*model.UserModel, error) {
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching users:", err.Error())
		return nil, err
	}
	var users []*model.UserModel
	err = res.All(context.TODO(), &users)
	if err != nil {
		log.Println("Error while decoding users:", err.Error())
		return nil, err
	}
	return users, nil
}

// AddUser inserts a new user into the database.
func (m MongoDB) AddUser(user *model.UserModel) error {
	_, err := m.collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error while inserting user:", err.Error())
		return err
	}
	return nil
}

// UpdateUser updates an existing user in the database.
func (m MongoDB) UpdateUser(user *model.UserModel) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	res, err := m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error while updating user:", err.Error())
		return err
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("user not found for ID: %s", user.ID.Hex())
	}
	return nil
}

// DeleteUser deletes a user from the database based on their ObjectID.
func (m MongoDB) DeleteUser(userID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	res, err := m.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error while deleting user:", err.Error())
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("user not found for ID: %s", userID.Hex())
	}
	return nil
}
