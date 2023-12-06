package songs

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tchipify/internal/models"

	"github.com/go-chi/chi/v5"
)

// GetSongByID handles the request to get a single song by its ID.
func GetSongByID(w http.ResponseWriter, r *http.Request) {
	// Extract the song ID from the URL parameters
	idStr := chi.URLParam(r, "id")

	// Parse the ID from the string
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Find the song in the slice based on the ID
	var foundSong *models.Song
	for _, song := range models.Songs {
		if song.ID == id {
			foundSong = &song
			break
		}
	}

	// If the song is not found, return a 404 Not Found response
	if foundSong == nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	// Convert the found song to JSON
	songJSON, err := json.Marshal(foundSong)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}