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

// DeleteSong
// @Tags         songs
// @Summary      Delete a single song by ID.
// @Description  Delete a single song by its ID from the database.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Success      204 "No Content"
// @Failure      400 "Invalid ID format"
// @Failure      404 "Song not found"
// @Failure      500 "Internal Server Error"
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Extract song ID from URL parameter
	songIDStr := chi.URLParam(r, "id")

	// Parse the ID as UUID
	songID, err := uuid.FromString(songIDStr)
	if err != nil {
		logrus.Errorf("parsing error: %s", err.Error())
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

	w.WriteHeader(http.StatusNoContent)
}
