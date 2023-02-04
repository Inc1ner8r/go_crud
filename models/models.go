package models

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Publication string `json:"publication"`
}
