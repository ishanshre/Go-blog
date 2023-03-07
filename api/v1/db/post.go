package db

import (
	"fmt"
	"os"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) PostCreate(post *models.NewPost) error {
	// create new post by authenticated user
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
	// get all post in pages according to request
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
	// returns specific post using id
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
	// delete post
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
	// update post
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
	// return owner id of post
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
	// checks and returns nill if post exists
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

func (s *PostgresStore) PostTagAdd(post_id, tag_id int) error {
	// add tag to the post
	if err := s.PostExist(post_id); err != nil {
		return err
	}
	if err := s.TagExist(tag_id); err != nil {
		return err
	}
	query := `
		INSERT INTO tag_post (post_id, tag_id)
		VALUES ($1, $2);
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, post_id, tag_id)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("tag was not added to the post")
	}
	return nil
}

func (s *PostgresStore) PostTagDelete(post_id, tag_id int) error {
	// remove tag from the post
	if err := s.PostExist(post_id); err != nil {
		return err
	}
	if err := s.TagExist(tag_id); err != nil {
		return err
	}
	query := `
		DELETE FROM tag_post
		WHERE post_id = $1 AND tag_id = $2
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, post_id, tag_id)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("tag not found to delete")
	}
	return nil
}

func (s *PostgresStore) PostTagsAll(post_id, limit, offset int) ([]*models.TagPost, error) {
	// returns all tags add to the specific post
	query := `
		SELECT * FROM tag_post
		WHERE post_id = $1
		LIMIT $2 OFFSET $3
	`
	rows, err := s.db.Query(query, post_id, limit, offset)
	if err != nil {
		return nil, err
	}
	tags := []*models.TagPost{}
	for rows.Next() {
		tag, err := ScanTagPost(rows)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
