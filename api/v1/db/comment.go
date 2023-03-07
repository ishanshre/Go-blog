package db

import (
	"fmt"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) CommentCreate(comment *models.NewComment) error {
	query := `
		INSERT INTO comments (
			content,
			rating,
			created_at,
			updated_at,
			post_id,
			user_id
		) VALUES ($1, $2, $3, $4, $5, $6);
		
	`
	rows, err := s.db.Exec(
		query,
		comment.Content,
		comment.Rating,
		comment.Created_at,
		comment.Updated_at,
		comment.Post_id,
		comment.User_id,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error in creating comment")
	}
	return nil
}

func (s *PostgresStore) CommentAllByPost(post_id int) ([]*models.Comment, error) {
	if err := s.PostExist(post_id); err != nil {
		return nil, err
	}
	query := `
		SELECT * FROM comments
		WHERE post_id = $1
	`
	rows, err := s.db.Query(query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []*models.Comment{}
	for rows.Next() {
		comment, err := ScanComment(rows)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (s *PostgresStore) CommentOwner(comment_id int) (*models.CommentOwner, error) {
	query := `
		SELECT user_id FROM comments
		WHERE id = $1
	`
	rows, err := s.db.Query(query, comment_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanCommentOwner(rows)
	}
	return nil, fmt.Errorf("comment with id %v does not exists", comment_id)
}

func (s *PostgresStore) CommentUpdate(post_id, comment_id int, comment *models.CommentUpdate) error {
	if err := s.PostExist(post_id); err != nil {
		return err
	}
	query := `
		UPDATE comments 
		SET content = $2, rating = $3, updated_at = $4
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(
		query,
		comment_id,
		comment.Content,
		comment.Rating,
		comment.Updated_at,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("comment with id %v does not exists", comment_id)
	}
	return nil
}

func (s *PostgresStore) CommentDelete(post_id, comment_id int) error {
	if err := s.PostExist(post_id); err != nil {
		return err
	}
	query := `
		DELETE FROM comments
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, comment_id)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("comment with id %v does not exists", comment_id)
	}
	return nil
}
