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

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Extract song ID from URL parameter
	songIDStr := chi.URLParam(r, "id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		logrus.Errorf("error parsing song ID: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call the service layer to delete the song
	err = songs.DeleteSongById(songID)
	if err != nil {
		logrus.Errorf("error deleting song: %s", err.Error())
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
