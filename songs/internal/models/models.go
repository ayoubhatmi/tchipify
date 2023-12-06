package models

  
type Song struct {
	//ID     *uuid.UUID    `json:"id"`
	ID     int   `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Songs is a slice to store sample songs.
var Songs = []Song{
	{ID: 1, Title: "Song 1", Artist: "Artist 1"},
	{ID: 2, Title: "Song 2", Artist: "Artist 2"},
	{ID: 3, Title: "Song 3", Artist: "Artist 3"},
}