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
	UserGetUsername(id int) (*models.GetUsername, error)

	// Tag Interface
	TagCreate(*models.Tag) error
	TagAll(int, int) ([]*models.Tag, error)
	TagDelete(int) error
	TagUpdate(int, *models.CreateTagRequest) error
	TagByID(int) (*models.Tag, error)
	TagExist(int) error

	// Post Interface
	PostCreate(*models.NewPost) error
	PostGetAll(int, int, string) ([]*models.Post, error)
	PostGetById(int, string) (*models.Post, error)
	PostDelete(int) (*models.PostPic, error)
	PostUpdate(int, *models.PostUpdate) error
	PostGetOwner(int) (*models.PostOwner, error)
	PostExist(int) error
	PostTagAdd(int, int) error
	PostTagDelete(int, int) error
	PostTagsAll(post_id, limit, offset int) ([]*models.TagPost, error)

	// Comment Interface
	CommentCreate(*models.NewComment) error
	CommentAllByPost(int) ([]*models.Comment, error)
	CommentUpdate(int, int, *models.CommentUpdate) error
	CommentDelete(int, int) error
	CommentOwner(int) (*models.CommentOwner, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// create, connect, ping database and returns database
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
