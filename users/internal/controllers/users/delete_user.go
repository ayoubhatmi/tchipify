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

// DeleteUser
// @Tags         users
// @Summary      Delete a single user by ID.
// @Description  Delete a single user by its ID from the database.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Success      204 "No Content"
// @Failure      400 "Invalid ID format"
// @Failure      404 "User not found"
// @Failure      500 "Internal Server Error"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL parameter
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logrus.Errorf("error parsing user ID: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call the service layer to delete the user
	err = users.DeleteUserById(userID)
	if err != nil {
		logrus.Errorf("error deleting user: %s", err.Error())
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
}
