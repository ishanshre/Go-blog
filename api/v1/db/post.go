package db

import (
	"fmt"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) PostCreate(post *models.NewPost) error {
	query := `
		INSERT INTO posts (
			title,
			slug,
			pic,
			content,
			created_at,
			updated_at,
			user_id
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	rows, err := s.db.Exec(
		query,
		post.Title,
		post.Slug,
		post.Pic,
		post.Content,
		post.Created_at,
		post.Updated_at,
		post.User_id,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error in creating post")
	}
	return nil
}

func (s *PostgresStore) PostGetAll(limit, offset int, url string) ([]*models.Post, error) {
	query := `
		SELECT * FROM posts
		LIMIT $1 OFFSET $2
	`
	posts := []*models.Post{}
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post, err := ScanPosts(rows, url)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
