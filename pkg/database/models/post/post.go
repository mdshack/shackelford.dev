package post

import "github.com/jmoiron/sqlx"

type Post struct {
	db *sqlx.DB
}

type Model struct {
	ID        string `db:"id"`
	Title     string `db:"title"`
	Slug      string `db:"slug"`
	Contents  string `db:"contents"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func New(db *sqlx.DB) *Post {
	return &Post{
		db: db,
	}
}

func (p *Post) All() []Model {
	var posts []Model

	p.db.Select(&posts, "SELECT * FROM posts ORDER BY created_at DESC")

	return posts
}

func (p *Post) WhereSlug(slug string) Model {
	var post Model

	p.db.Get(&post, "SELECT * FROM posts WHERE slug=$1", slug)

	return post
}
