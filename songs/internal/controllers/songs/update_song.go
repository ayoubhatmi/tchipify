package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
	"tchipify/services/songs"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateSong
// @Tags         songs
// @Summary      Update a single song by ID.
// @Description  Update a single song by its ID in the database.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Param        title         formData  string  true  "New title"
// @Param        artist        formData  string  true  "New artist"
// @Produce      json
// @Success      200 {object} models.Song
// @Failure      400 "Invalid ID format"
// @Failure      404 "Song not found"
// @Failure      500 "Internal Server Error"
// @Router       /songs/{id} [put]
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

	// Parse song ID from string to UUID
	songID, err := uuid.FromString(songIDStr)
	if err != nil {
		logrus.Errorf("error parsing song ID: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Set the ID for the updated song
	updatedSong.ID = songID.String()

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
