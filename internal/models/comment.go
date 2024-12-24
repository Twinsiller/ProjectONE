package models

import "time"

type Comment struct {
	Id               int       `json:"id"`
	IdAuthor         int       `json:"id_author"`
	IdPost           int       `json:"id_post"`
	Text             string    `json:"text_comment"`
	DatePublication  time.Time `json:"date_publication"`
	DateLastModified time.Time `json:"date_last_modified"`
}
