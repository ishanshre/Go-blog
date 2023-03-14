package db

import (
	"fmt"
	"time"

	"github.com/ishanshre/Go-blog/api/v1/models"
)

func (s *PostgresStore) UserSignUp(user *models.RegsiterUser) error {
	// register new user
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			username,
			email,
			password,
			created_at,
			updated_at,
			last_login
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`
	rows, err := s.db.Exec(
		query,
		user.FirstName,
		user.LastName,
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
	// return id and password
	query := `
		SELECT id, password FROM users
		WHERE username = $1
	`
	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanUserPass(rows)
	}
	return nil, fmt.Errorf("username: %v not found", username)
}

func (s *PostgresStore) UpdateLastLogin(id int) error {
	// update login date
	query := `
		UPDATE users
		SET last_login = $2
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, id, time.Now())
	return err
}

func (s *PostgresStore) UserInfoById(id int) (*models.User, error) {
	// return user info by id
	query := `
		SELECT * FROM users
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanUser(rows)
	}
	return nil, fmt.Errorf("account with id %v not found", id)
}

func (s *PostgresStore) UsersAll() ([]*models.User, error) {
	// returns all user registered
	query := `
		SELECT * FROM users
	`
	users := []*models.User{}
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user, err := ScanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

func (s *PostgresStore) UserDelete(id int) error {
	// delete user account
	query := `
		DELETE FROM users
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
		return fmt.Errorf("error in deleting user")
	}
	return nil
}

func (s *PostgresStore) UserGetUsername(id int) (*models.GetUsername, error) {
	query := `
		SELECT username FROM users
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanUsername(rows)
	}
	return nil, fmt.Errorf("error in getting username")
}
