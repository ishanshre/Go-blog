package db

import (
	"fmt"
	"time"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) TagCreate(tag *models.Tag) error {
	// create a new tag
	query := `
		INSERT INTO tags (name, created_at, updated_at)
		VALUES ($1, $2, $3)
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, tag.Name, tag.CreatedAt, tag.UpdatedAt)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error in creating new record")
	}
	return nil
}

func (s *PostgresStore) TagAll(limit, offset int) ([]*models.Tag, error) {
	// return all tag
	query := `
		SELECT * FROM tags
		LIMIT $1 OFFSET $2;
	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	tags := []*models.Tag{}
	defer rows.Close()
	for rows.Next() {
		tag, err := ScanTags(rows)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *PostgresStore) TagDelete(id int) error {
	// delete tag using id
	query := `
		DELETE FROM tags
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error in deleting tags or does not exist")
	}
	return nil
}

func (s *PostgresStore) TagUpdate(id int, tag *models.CreateTagRequest) error {
	// update tag using id
	query := `
		UPDATE tags
		SET name = $2, updated_at = $3
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, id, tag.Name, time.Now())
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("tag update failed or tag does not exists")
	}
	return nil
}

func (s *PostgresStore) TagByID(id int) (*models.Tag, error) {
	// returns tag using id
	query := `
		SELECT * FROM tags
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return ScanTags(rows)
	}
	return nil, fmt.Errorf("error in fetch tag or tag does not exists")
}

func (s *PostgresStore) TagExist(tag_id int) error {
	// checks if tag exists and if exists returns nil
	post, err := s.db.Exec(`SELECT id FROM tags WHERE id = $1`, tag_id)
	if err != nil {
		return err
	}
	rows_affected, err := post.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("post with id %v does not exists", tag_id)
	}
	return nil
}
