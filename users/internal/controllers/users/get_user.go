package users

import (
	"encoding/json"
	"net/http"
	"tchipify/users/internal/models"
	"tchipify/users/services/users"

	"github.com/sirupsen/logrus"
)

// GetUserByID
// @Tags         users
// @Summary      Get a single user by ID.
// @Description  Get a single user by its ID from the database.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Produce      json
// @Success      200 {object} models.User
// @Failure      400 "Invalid ID format"
// @Failure      404 "User not found"
// @Failure      500 "Internal Server Error"
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(int)

	collection, err := users.GetUserById(collectionId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(collection)
	_, _ = w.Write(body)
	return
}
