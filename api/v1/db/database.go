package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/Go-blog/api/v1/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	// User Interface
	UserSignUp(*models.RegsiterUser) error
	UserLogin(string) (*models.UserPass, error)
	UpdateLastLogin(int) error
	UserInfoById(int) (*models.User, error)
	UsersAll() ([]*models.User, error)
	UserDelete(int) error

	// Tag Interface
	TagCreate(*models.Tag) error
	TagAll(int, int) ([]*models.Tag, error)
	TagDelete(int) error
	TagUpdate(int, *models.CreateTagRequest) error
	TagByID(int) (*models.Tag, error)

	// Post Interface
	PostCreate(*models.NewPost) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	if err := godotenv.Load("./.env"); err != nil {
		return nil, fmt.Errorf("error in loading environment files: %s", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error in connetiong to database: %s", err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}
