package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
	"tchipify/services/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func AddSong(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body to get the new song data
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the song ID
	uuid, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("error generating UUID: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the generated UUID as the ID for the new song
	newSong.ID = uuid.String()

	// Call the service layer to create the new song
	createdSong, err := songs.CreateSong(newSong)
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())
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

	// Return the created song in the response
	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(createdSong)
	_, _ = w.Write(response)
}