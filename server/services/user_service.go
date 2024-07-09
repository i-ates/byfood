package services

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task4/db"
	"task4/model"
)

// UserServiceInterface defines the interface for user-related services.
type UserServiceInterface interface {
	GetUsers() (error, []*model.UserModel)
	AddUser(user *model.UserModel) error
	UpdateUser(user *model.UserModel) error
	DeleteUser(user *model.UserModel) error
}

// UserService is the implementation of the UserServiceInterface.
type UserService struct {
	db db.DB
}

// AddUser adds a new user to the database.
func (u UserService) AddUser(user *model.UserModel) error {
	err := u.db.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers retrieves a list of all users from the database.
func (u UserService) GetUsers() (error, []*model.UserModel) {
	users, err := u.db.GetUsers()
	if err != nil {
		return err, nil
	}

	return nil, users
}

// UpdateUser updates an existing user in the database.
func (u UserService) UpdateUser(userModel *model.UserModel) error {
	err := u.db.UpdateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database.
func (u UserService) DeleteUser(user *model.UserModel) error {
	if user.ID == (primitive.ObjectID{}) {
		return fmt.Errorf("userID can not be nil")
	}
	err := u.db.DeleteUser(user.ID)
	if err != nil {
		return err
	}

	return nil
}

// NewUserService creates a new UserService instance with the specified DB implementation.
func NewUserService(db db.DB) UserServiceInterface {
	return &UserService{
		db: db,
	}
}
