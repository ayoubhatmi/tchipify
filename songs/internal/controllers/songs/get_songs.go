package songs

import (
	"encoding/json"
	"net/http"
	"tchipify/internal/models"
)

// GetAllSongs handles the request to get all songs.
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	// Convert songs slice to JSON
	songsJSON, err := json.Marshal(models.Songs)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songsJSON)
}