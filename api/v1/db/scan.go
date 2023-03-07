package db

import (
	"database/sql"
	"fmt"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func ScanUserPass(rows *sql.Rows) (*models.UserPass, error) {
	user := new(models.UserPass)
	err := rows.Scan(
		&user.ID,
		&user.Password,
	)
	return user, err
}

func ScanUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)
	return user, err
}

func ScanTags(rows *sql.Rows) (*models.Tag, error) {
	tag := new(models.Tag)
	err := rows.Scan(
		&tag.ID,
		&tag.Name,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	return tag, err
}

func ScanPosts(rows *sql.Rows, domain string) (*models.Post, error) {
	post := new(models.Post)
	err := rows.Scan(
		&post.Id,
		&post.Title,
		&post.Slug,
		&post.Pic,
		&post.Content,
		&post.Created_at,
		&post.Updated_at,
		&post.User_id,
	)
	pic := fmt.Sprintf("%s/media/image/%s", domain, post.Pic)
	post.Pic = pic
	return post, err
}

func ScanPostPic(rows *sql.Rows) (*models.PostPic, error) {
	post := new(models.PostPic)
	err := rows.Scan(&post.Pic)
	return post, err
}

func ScanPostOwner(rows *sql.Rows) (*models.PostOwner, error) {
	Owner := new(models.PostOwner)
	err := rows.Scan(&Owner.User_id)
	return Owner, err
}

func ScanCommentOwner(rows *sql.Rows) (*models.CommentOwner, error) {
	owner := new(models.CommentOwner)
	err := rows.Scan(&owner.User_id)
	return owner, err
}

func ScanComment(rows *sql.Rows) (*models.Comment, error) {
	comment := new(models.Comment)
	err := rows.Scan(
		&comment.Id,
		&comment.Content,
		&comment.Rating,
		&comment.Created_at,
		&comment.Updated_at,
		&comment.Post_id,
		&comment.User_id,
	)
	return comment, err
}
