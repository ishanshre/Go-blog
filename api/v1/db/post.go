package db

import (
	"fmt"
	"os"

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

func (s *PostgresStore) PostGetAll(limit, offset int, domain string) ([]*models.Post, error) {
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
		post, err := ScanPosts(rows, domain)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostgresStore) PostGetById(id int, domain string) (*models.Post, error) {
	query := `
		SELECT * FROM posts
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanPosts(rows, domain)
	}
	return nil, fmt.Errorf("post with id %v does not exists", id)
}

func (s *PostgresStore) PostDelete(id int) (*models.PostPic, error) {
	query1 := `
	SELECT pic FROM posts
	WHERE id = $1
	`
	query2 := `
		DELETE FROM posts
		WHERE id = $1
	`
	rows, err := s.db.Query(query1, id)
	if err != nil {
		return nil, err
	}
	var postPic *models.PostPic
	for rows.Next() {
		postPic, err = ScanPostPic(rows)
		if err != nil {
			return nil, err
		}
	}
	s.db.Exec("COMMIT")
	row2, err := s.db.Exec(query2, id)
	if err != nil {
		return nil, err
	}
	rows_affected, err := row2.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows_affected == 0 {
		return nil, fmt.Errorf("post with id %v does not exists", id)
	}
	return postPic, nil
}

func (s *PostgresStore) PostUpdate(id int, post *models.PostUpdate) error {
	picQuery := `
		SELECT pic FROM posts
		WHERE id = $1
	`
	picRows, err := s.db.Query(picQuery, id)
	if err != nil {
		return err
	}
	pic := new(models.PostPic)
	for picRows.Next() {
		picRows.Scan(&pic.Pic)
	}
	query := `
		UPDATE posts
		SET title = $2, slug = $3, content = $4,  pic = $5, updated_at = $6
		WHERE id= $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(
		query,
		id,
		post.Title,
		post.Slug,
		post.Content,
		post.Pic,
		post.Updated_at,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("post not updated")
	}
	if err := os.Remove(fmt.Sprintf("./media/uploads/posts/%s", pic.Pic)); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) PostGetOwner(id int) (*models.PostOwner, error) {
	query := `
		SELECT user_id FROM posts
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanPostOwner(rows)
	}
	return nil, fmt.Errorf("post does not exists")
}

func (s *PostgresStore) PostExist(post_id int) error {
	post, err := s.db.Exec(`SELECT id FROM posts WHERE id = $1`, post_id)
	if err != nil {
		return err
	}
	rows_affected, err := post.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("post with id %v does not exists", post_id)
	}
	return nil
}
