package models

import (
	"fmt"
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

type PostPic struct {
	Pic string `json:"pic"`
}

type PostOwner struct {
	User_id int `json:"user_id"`
}

type PostUpdate struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Pic        string    `json:"pic"`
	Updated_at time.Time `json:"updated_at"`
}

func NewPostUpdate(id int, title, slug, content, pic string) *PostUpdate {
	return &PostUpdate{
		Id:         id,
		Title:      title,
		Slug:       slug,
		Content:    content,
		Pic:        pic,
		Updated_at: time.Now(),
	}
}

func NewPostCreate(title, content, pic string, user_id int) *NewPost {
	text := fmt.Sprintf("%s %s", title, time.Now().Format(time.DateTime))
	slugCreate := slug.Make(text)
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
