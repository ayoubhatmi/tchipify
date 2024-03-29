package main

import (
	"net/http"
	"tchipify/internal/controllers/songs"
	"tchipify/internal/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs)
		r.Post("/", songs.AddSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(songs.Ctx)

			// Use *uuid.UUID as the type for the id parameter
			r.Get("/", songs.GetSong)
			r.Put("/", songs.UpdateSong)
			r.Delete("/", songs.DeleteSong)
		})
	})
	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{`
	CREATE TABLE IF NOT EXISTS songs (
		id TEXT PRIMARY KEY,
		artist VARCHAR(255) NOT NULL,
		file_name VARCHAR(255) NOT NULL,
		published_date DATETIME NOT NULL,
		title VARCHAR(255) NOT NULL
	);
`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}


