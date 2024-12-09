package models

import "time"

type Post struct {
	Id               int       `json:"id"`
	Title            string    `json:"title"`
	IdAuthor         int       `json:"id_author"`
	Description      string    `json:"description"`
	Likes            int       `json:"likes"`
	DatePublication  time.Time `json:"date_publication"`
	DateLastModified time.Time `json:"date_last_modified"`
}
