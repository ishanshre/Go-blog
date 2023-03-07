package models

import (
	"time"

	"github.com/gosimple/slug"
)

type Post struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Pic        string    `json:"pic"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id    int       `json:"user_id"`
}
type NewPost struct {
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Pic        string    `json:"pic"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id    int       `json:"user_id"`
}

type NewPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPostCreate(title, content, pic string, user_id int) *NewPost {
	slugCreate := slug.Make(title)
	return &NewPost{
		Title:      title,
		Slug:       slugCreate,
		Pic:        pic,
		Content:    content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		User_id:    user_id,
	}
}
