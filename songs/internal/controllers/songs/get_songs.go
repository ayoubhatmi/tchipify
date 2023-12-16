package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
	"tchipify/services/songs"

	"github.com/sirupsen/logrus"
)

// GetSongs
// @Tags         songs
// @Summary      Get all songs.
// @Description  Get all songs from the database.
// @Produce      json
// @Success      200 {array} models.Song
// @Failure      500 "Internal Server Error"
// @Router       /songs [get]
func GetSongs(w http.ResponseWriter, _ *http.Request) {
	// calling service
	songs, err := songs.GetAllSongs()
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
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
	return
}

