package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
	"tchipify/services/songs"

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
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(int)

	collection, err := songs.GetSongById(collectionId)
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
