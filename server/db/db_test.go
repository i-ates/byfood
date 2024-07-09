package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task4/model"
)

// TestMongoDB is a test suite for the MongoDB struct and its methods.
// You need to run docker pull mongo && docker run -d --name mongodbtest -p 27017:27017 mongo if not tests will fail.
func TestMongoDB(t *testing.T) {
	// Establish a connection to a test MongoDB server.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to create MongoDB client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Create a MongoDB instance for testing.
	db := NewMongo(client)

	// Define a UserModel for testing.
	user := &model.UserModel{
		ID:   primitive.ObjectID{},
		Name: "testuser",
		Mail: "test@example.com",
	}

	t.Run("Test AddUser and GetUsers", func(t *testing.T) {
		// Add a user to the database.
		err := db.AddUser(user)
		assert.NoError(t, err)

		// Retrieve users from the database.
		users, err := db.GetUsers()
		user = users[0]
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, user.Name, users[0].Name)
	})

	t.Run("Test UpdateUser", func(t *testing.T) {
		// Update the user's email.
		user.Mail = "updated@example.com"
		err := db.UpdateUser(user)
		assert.NoError(t, err)

		// Retrieve the updated user from the database.
		updatedUser, err := db.GetUsers()
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, user.Mail, updatedUser[0].Mail)
	})

	t.Run("Test DeleteUser", func(t *testing.T) {
		// Delete the user from the database.
		err := db.DeleteUser(user.ID)
		assert.NoError(t, err)

		// Attempt to retrieve the deleted user (should return an error).
		deletedUser, err := db.GetUsers()
		assert.Error(t, err)
		assert.Nil(t, deletedUser)
	})
}
