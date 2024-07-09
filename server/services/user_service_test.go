package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task4/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) AddUser(user *model.UserModel) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) GetUsers() ([]*model.UserModel, error) {
	args := m.Called()
	return args.Get(0).([]*model.UserModel), args.Error(1)
}

func (m *MockDB) UpdateUser(user *model.UserModel) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) DeleteUser(userID primitive.ObjectID) error {
	args := m.Called(userID)
	return args.Error(0)
}

func TestAddUser(t *testing.T) {
	mockDB := new(MockDB)
	userService := NewUserService(mockDB)

	user := &model.UserModel{
		// Initialize user fields as needed.
	}

	// Mock the behavior of the DB.
	mockDB.On("AddUser", user).Return(nil)

	err := userService.AddUser(user)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestGetUsers(t *testing.T) {
	mockDB := new(MockDB)
	userService := NewUserService(mockDB)

	users := []*model.UserModel{
		// Initialize user models as needed.
	}

	// Mock the behavior of the DB.
	mockDB.On("GetUsers").Return(users, nil)

	err, returnedUsers := userService.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, returnedUsers)
	mockDB.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockDB := new(MockDB)
	userService := NewUserService(mockDB)

	user := &model.UserModel{
		// Initialize user fields as needed.
	}

	// Mock the behavior of the DB.
	mockDB.On("UpdateUser", user).Return(nil)

	err := userService.UpdateUser(user)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockDB := new(MockDB)
	userService := NewUserService(mockDB)

	userID := primitive.NewObjectID()
	user := &model.UserModel{
		ID: userID,
	}

	// Mock the behavior of the DB.
	mockDB.On("DeleteUser", userID).Return(nil)

	err := userService.DeleteUser(user)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteUserWithNilID(t *testing.T) {
	mockDB := new(MockDB)
	userService := NewUserService(mockDB)

	user := &model.UserModel{
		ID: primitive.ObjectID{},
	}

	err := userService.DeleteUser(user)

	assert.Error(t, err)
	assert.EqualError(t, err, "userID can not be nil")
	mockDB.AssertExpectations(t)
}
