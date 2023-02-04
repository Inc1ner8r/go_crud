package models

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Publication string `json:"publication"`
}
