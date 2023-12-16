package users

import (
	"encoding/json"
	"net/http"
	"tchipify/users/internal/models"
	"tchipify/users/services/users"

	"github.com/sirupsen/logrus"
)

// GetUsers
// @Tags         users
// @Summary      Get all users.
// @Description  Get all users from the database.
// @Produce      json
// @Success      200 {array} models.User
// @Failure      500 "Internal Server Error"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service
	users, err := users.GetAllUsers()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
	return
}

