package home

import (
	"context"
	"net/http"

	"github.com/mdshack/shackelford.dev/internal/container"
	"github.com/mdshack/shackelford.dev/pkg/templates/pages"
)

type HomeController struct {
	container *container.Container
}

func New(c *container.Container) HomeController {
	return HomeController{
		container: c,
	}
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(context.Background(), w)
}
