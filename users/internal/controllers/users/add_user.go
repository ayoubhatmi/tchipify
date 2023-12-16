package users

import (
	"encoding/json"
	"net/http"
	"tchipify/users/internal/models"
	"tchipify/users/services/users"

	"github.com/sirupsen/logrus"
)

// AddUser
// @Summary Add a new user
// @Description Adds a new user to the system.
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object to be added"
// @Success 201 {object} models.User "User added successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func AddUser(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body to get the new user data
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service layer to create the new user
	createdUser, err := users.CreateUser(newUser)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// Return a custom error response
			w.WriteHeader(customError.Code)
			response, _ := json.Marshal(customError)
			_, _ = w.Write(response)
		} else {
			// Return a generic error response for internal server errors
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Return the created user in the response
	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(createdUser)
	_, _ = w.Write(response)
}
