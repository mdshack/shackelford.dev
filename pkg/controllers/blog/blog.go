package blog

import (
	"context"
	"net/http"

	"github.com/mdshack/shackelford.dev/internal/container"
	"github.com/mdshack/shackelford.dev/pkg/templates/pages"
)

type BlogController struct {
	container *container.Container
}

func New(c *container.Container) BlogController {
	return BlogController{
		container: c,
	}
}

func (b *BlogController) Index(w http.ResponseWriter, r *http.Request) {
	// posts := b.container.DB.Posts.All()

	pages.Blog().Render(context.Background(), w)
}

func (b *BlogController) Show(w http.ResponseWriter, r *http.Request) {
	// slug := mux.Vars(r)["slug"]
	// post := b.container.DB.Posts.WhereSlug(slug)

	// pages.BlogPost(post).Render(context.Background(), w)
}
