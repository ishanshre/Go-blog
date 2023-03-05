package db

import (
	"database/sql"

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
