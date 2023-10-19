package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mdshack/shackelford.dev/internal/config"
	"github.com/mdshack/shackelford.dev/pkg/database/models/post"
)

type Database struct {
	db *sqlx.DB

	Posts *post.Post
}

func New(cfg *config.Config) *Database {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host,
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Database,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	return &Database{
		db: db,

		Posts: post.New(db),
	}
}
