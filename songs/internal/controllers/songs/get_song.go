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

// GetSongByID
// @Tags         songs
// @Summary      Get a single song by ID.
// @Description  Get a single song by its ID from the database.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Produce      json
// @Success      200 {object} models.Song
// @Failure      400 "Invalid ID format"
// @Failure      404 "Song not found"
// @Failure      500 "Internal Server Error"
// @Router       /songs/{id} [get]
func GetSong(w http.ResponseWriter, r *http.Request) {
	// Extract song ID from URL parameter
	songIDStr := chi.URLParam(r, "id")

	// Parse the song ID as a UUID
	songID, err := uuid.FromString(songIDStr)
	if err != nil {
		logrus.Errorf("error parsing song ID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service layer to get the song by ID
	collection, err := songs.GetSongById(songID)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
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
