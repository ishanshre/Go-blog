package db

import (
	"fmt"
	"time"

	"github.com/ishanshre/Go-blog/api/v1/middlewares"
	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) UserSignUp(user *models.RegsiterUser) error {
	query := `
		INSERT INTO users (
			username,
			email,
			password,
			created_at,
			updated_at,
			last_login
		) VALUES (
			$1, $2, $3, $4, $5, $6 
		)
	`
	rows, err := s.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.LastLogin,
	)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error in creating new user")
	}
	return nil
}

func (s *PostgresStore) UserLogin(username string) (*models.UserPass, error) {
	query := `
		SELECT id, password FROM users
		WHERE username = $1
	`
	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return middlewares.ScanUserPass(rows)
	}
	return nil, fmt.Errorf("username: %v not found", username)
}

func (s *PostgresStore) UpdateLastLogin(id int) error {
	query := `
		UPDATE users
		SET last_login = $2
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, id, time.Now())
	return err
}
