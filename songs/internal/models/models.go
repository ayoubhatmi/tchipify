package models

  
type Song struct {
	ID            int    `json:"id"`
	Artist        string    `json:"artist"`
	FileName      string    `json:"file_name"`
	PublishedDate string `json:"published_date"`
	Title         string    `json:"title"`
}

 