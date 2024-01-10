package songs

import (
	"database/sql"
	"net/http"
	"tchipify/internal/models"
	repository "tchipify/internal/repositories/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllSongs()
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

func GetSongById(id uuid.UUID) (*models.Song, error) {
	collection, err := repository.GetSongById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "Song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving song: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return collection, err
}

func CreateSong(newSong models.Song) (*models.Song, error) {

	var err error

	// calling repository to insert the new song
	createdSong, err := repository.CreateSong(newSong)

	// managing errors
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())

		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdSong, nil
}


func UpdateSongById(updatedSong models.Song) error {
	err := repository.UpdateSongById(updatedSong)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CustomError{
				Message: "Song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating song in repository: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func DeleteSongById(id uuid.UUID) error {
	// Call the repository to delete the song
	err := repository.DeleteSongById(id)
	if err != nil {
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
