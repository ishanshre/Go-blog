package models

import "time"

type Comment struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Post_id    int       `json:"post_id"`
	User_id    int       `json:"user_id"`
}

type CommentReq struct {
	Content string `json:"content"`
	Rating  int    `json:"rating"`
}

type NewComment struct {
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Post_id    int       `json:"post_id"`
	User_id    int       `json:"user_id"`
}

type CommentUpdate struct {
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
	Updated_at time.Time `json:"updated_at"`
}

type CommentOwner struct {
	User_id int `json:"user_id"`
}

func NewCommentUpdate(content string, rating int) *CommentUpdate {
	return &CommentUpdate{
		Content:    content,
		Rating:     rating,
		Updated_at: time.Now(),
	}
}

func NewCommentCreate(content string, rating, post_id, user_id int) *NewComment {
	return &NewComment{
		Content:    content,
		Rating:     rating,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Post_id:    post_id,
		User_id:    user_id,
	}
}
