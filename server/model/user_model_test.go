package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUserModel(t *testing.T) {
	// Create a test UserModel instance
	user := UserModel{
		Name: "John Doe",
		Mail: "johndoe@example.com",
		Age:  30,
	}

	// Test that the ID field is initialized as empty
	if user.ID != primitive.NilObjectID {
		t.Errorf("Expected ID to be empty, but got %v", user.ID)
	}

	// Test the Name field
	expectedName := "John Doe"
	if user.Name != expectedName {
		t.Errorf("Expected Name to be %s, but got %s", expectedName, user.Name)
	}

	// Test the Mail field
	expectedMail := "johndoe@example.com"
	if user.Mail != expectedMail {
		t.Errorf("Expected Mail to be %s, but got %s", expectedMail, user.Mail)
	}

	// Test the Age field
	expectedAge := 30
	if user.Age != expectedAge {
		t.Errorf("Expected Age to be %d, but got %d", expectedAge, user.Age)
	}
}
