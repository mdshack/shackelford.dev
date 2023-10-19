package container

import (
	"embed"

	"github.com/gorilla/mux"
	"github.com/mdshack/shackelford.dev/internal/config"
	"github.com/mdshack/shackelford.dev/pkg/database"
)

type Container struct {
	Router *mux.Router
	DB     *database.Database
	Config *config.Config
	Assets *embed.FS
}

func New(assets *embed.FS) *Container {
	return &Container{
		Router: mux.NewRouter(),
		// DB:     database.New(),
		Config: config.New(),
		Assets: assets,
	}
}
