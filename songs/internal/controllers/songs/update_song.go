package songs

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tchipify/internal/models"
	"tchipify/services/songs"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var updatedSong models.Song
	err := json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract song ID from URL parameter
	songIDStr := chi.URLParam(r, "id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		logrus.Errorf("error parsing song ID: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Set the ID for the updated song
	updatedSong.ID = songID

	// Call the service layer to update the song
	err = songs.UpdateSongById(updatedSong)
	if err != nil {
		logrus.Errorf("error updating song: %s", err.Error())
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
	response, _ := json.Marshal(updatedSong)
	_, _ = w.Write(response)
}
