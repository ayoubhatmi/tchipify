package songs

import (
	"tchipify/internal/helpers"
	"tchipify/internal/models"
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

func GetSongById(id int) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id)
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.ID, &song.Artist, &song.FileName, &song.PublishedDate, &song.Title )
	if err != nil {
		return nil, err
	}
	return &song, err
}


func CreateSong(newSong models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("INSERT INTO songs (artist, file_name, published_date, title) VALUES (?, ?, ?, ?)",
		newSong.Artist, newSong.FileName, newSong.PublishedDate, newSong.Title)

	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Update the Song object with the generated ID
	newSong.ID = int(id)
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

func DeleteSongById(id int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}
