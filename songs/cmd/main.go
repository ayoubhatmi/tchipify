package main

import (
	"net/http"
	"tchipify/internal/controllers/songs"

	"github.com/go-chi/chi/v5"
)




func main() {
	r := chi.NewRouter()

	// Routes
	r.Get("/songs", songs.GetAllSongs)
	r.Get("/songs/{id}", songs.GetSongByID)

	// Start the server
	http.ListenAndServe(":8080", r)
}

