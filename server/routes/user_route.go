package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task4/model"
	"task4/services"
)

// UserRouterInterface defines the interface for user-related HTTP routes.
type UserRouterInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

// UserRouter is the implementation of the UserRouterInterface.
type UserRouter struct {
	userService services.UserServiceInterface
}

// GetUsers handles the HTTP GET request to retrieve a list of users.
func (u *UserRouter) GetUsers(w http.ResponseWriter, r *http.Request) {
	err, users := u.userService.GetUsers()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

// AddUser handles the HTTP POST request to add a new user.
func (u *UserRouter) AddUser(w http.ResponseWriter, r *http.Request) {
	var user *model.UserModel

	// Assuming the request's content type is JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := u.userService.AddUser(user)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// UpdateUser handles the HTTP PUT request to update an existing user.
func (u *UserRouter) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.UserModel

	// Assuming the request's content type is JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := u.userService.UpdateUser(user)
	if err != nil {
		if err.Error() == fmt.Errorf("user not found for ID: %s", user.ID.Hex()).Error() {
			sendErr(w, http.StatusNotFound, err.Error())
			return
		}

		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// DeleteUser handles the HTTP DELETE request to delete a user.
func (u *UserRouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user *model.UserModel

	// Assuming the request's content type is JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := u.userService.DeleteUser(user)
	if err != nil {
		if err.Error() == fmt.Errorf("user not found for ID: %s", user.ID.Hex()).Error() {
			sendErr(w, http.StatusNotFound, err.Error())
			return
		}

		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// sendErr sends an error response to the client with the specified status code and message.
func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// NewUserRouter creates a new UserRouter instance with the specified UserServiceInterface.
func NewUserRouter(userService services.UserServiceInterface) UserRouterInterface {
	return &UserRouter{
		userService: userService,
	}
}
