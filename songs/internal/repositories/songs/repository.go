package songs

import (
	"tchipify/internal/helpers"
	"tchipify/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.ID, &data.Artist, &data.FileName, &data.PublishedDate, &data.Title)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}
func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())

	var song models.Song
	err = row.Scan(&song.ID, &song.Artist, &song.FileName, &song.PublishedDate, &song.Title)
	if err != nil {
		return nil, err
	}
	return &song, nil
}


func CreateSong(newSong models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db) // Ensure that the database connection is closed even if an error occurs.

	// Generate a new UUID for the song ID
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO songs (id, artist, file_name, published_date, title) VALUES (?, ?, ?, ?, ?)",
    uuid.String(), newSong.Artist, newSong.FileName, newSong.PublishedDate, newSong.Title)

	if err != nil {
		return nil, err
	}

	// Update the Song object with the generated ID
	newSong.ID = uuid.String()

	return &newSong, nil
}


func UpdateSongById(updatedSong models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Update the song in the database
	_, err = db.Exec("UPDATE songs SET artist=?, file_name=?, published_date=?, title=? WHERE id=?",
		updatedSong.Artist, updatedSong.FileName, updatedSong.PublishedDate, updatedSong.Title, updatedSong.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSongById(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id=?", id.String())
	if err != nil {
		return err
	}

	return nil
}