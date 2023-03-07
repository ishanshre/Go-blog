package models

import "time"

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type TagReq struct {
	ID int `json:"id"`
}

type TagPost struct {
	Post_id int `json:"post_id"`
	Tag_id  int `json:"tag_id"`
}

func CreateNewTag(name string) *Tag {
	return &Tag{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
