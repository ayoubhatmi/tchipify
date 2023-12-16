package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tchipify/users/internal/models"
	"tchipify/users/services/users"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         users
// @Summary      Update a single user by ID.
// @Description  Update a single user by its ID in the database.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Param        title         formData  string  true  "New title"
// @Param        artist        formData  string  true  "New artist"
// @Produce      json
// @Success      200 {object} models.User
// @Failure      400 "Invalid ID format"
// @Failure      404 "User not found"
// @Failure      500 "Internal Server Error"
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract user ID from URL parameter
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logrus.Errorf("error parsing user ID: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Set the ID for the updated user
	updatedUser.ID = userID

	// Call the service layer to update the user
	err = users.UpdateUserById(updatedUser)
	if err != nil {
		logrus.Errorf("error updating user: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			response, _ := json.Marshal(customError)
			_, _ = w.Write(response)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedUser)
	_, _ = w.Write(response)
}
