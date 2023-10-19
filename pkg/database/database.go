package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mdshack/shackelford.dev/pkg/database/models/post"
)

type Database struct {
	db *sqlx.DB

	Posts *post.Post
}

func New() *Database {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return &Database{
		db: db,

		Posts: post.New(db),
	}
}
