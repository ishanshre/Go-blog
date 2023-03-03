package db

import (
	"fmt"

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
