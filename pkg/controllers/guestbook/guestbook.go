package guestbook

import (
	"context"
	"net/http"

	"github.com/mdshack/shackelford.dev/internal/container"
	"github.com/mdshack/shackelford.dev/pkg/templates/pages"
)

type GuestbookController struct {
	container *container.Container
}

func New(c *container.Container) GuestbookController {
	return GuestbookController{
		container: c,
	}
}

func (g *GuestbookController) Index(w http.ResponseWriter, r *http.Request) {
	pages.Guestbook().Render(context.Background(), w)
}
