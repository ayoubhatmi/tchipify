package users

import (
	"database/sql"
	"net/http"
	"tchipify/users/internal/models"
	repository "tchipify/users/internal/repositories/users"

	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func CreateUser(newUser models.User) (*models.User, error) {

	var err error

	createdUser, err := repository.CreateUser(newUser)

	// managing errors
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())



		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdUser, nil
}


func DeleteUserById(id int) error {
	err := repository.DeleteUserById(id)
	if err != nil {
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func GetUserById(id int) (*models.User, error) {
	collection, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}

func UpdateUserById(updatedUser models.User) error {
	err := repository.UpdateUserById(updatedUser)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user in repository: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
