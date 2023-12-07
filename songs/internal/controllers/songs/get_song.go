package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
	"tchipify/services/songs"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

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
