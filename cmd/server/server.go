package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/mdshack/shackelford.dev/internal/container"
	"github.com/mdshack/shackelford.dev/pkg/controllers/auth"
	"github.com/mdshack/shackelford.dev/pkg/controllers/blog"
	"github.com/mdshack/shackelford.dev/pkg/controllers/guestbook"
	"github.com/mdshack/shackelford.dev/pkg/controllers/home"
)

func New(assets *embed.FS) Server {
	c := container.New(assets)

	return Server{
		container: c,
		home:      home.New(c),
		blog:      blog.New(c),
		guestbook: guestbook.New(c),
		auth:      auth.New(c),
	}
}

type Server struct {
	container *container.Container

	// Controllers
	home      home.HomeController
	blog      blog.BlogController
	guestbook guestbook.GuestbookController
	auth      auth.AuthController
}

func (s *Server) Serve() error {
	s.registerRoutes()
	s.serveStaticFiles()

	addr := "0.0.0.0:8000"

	server := &http.Server{
		Handler:      s.container.Router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening at %s\n", addr)

	return server.ListenAndServe()
}

func (s *Server) registerRoutes() *Server {
	s.container.Router.HandleFunc("/", s.home.Index).Methods("GET")

	s.container.Router.HandleFunc("/blog", s.blog.Index).Methods("GET")
	// s.container.Router.HandleFunc("/blog/{slug}", s.blog.Show).Methods("GET")

	s.container.Router.HandleFunc("/guestbook", s.guestbook.Index).Methods("GET")

	// s.container.Router.HandleFunc("/auth/login", s.auth.Show).Methods("GET")
	// s.container.Router.HandleFunc("/auth/redirect", s.auth.Redirect).Methods("GET")
	// s.container.Router.HandleFunc("/auth/callback", s.auth.Callback).Methods("GET")

	return s
}

func (s *Server) serveStaticFiles() *Server {
	staticDirectory, _ := fs.Sub(*s.container.Assets, "assets/dist")
	s.container.Router.PathPrefix("/").Handler(http.FileServer(http.FS(staticDirectory)))

	return s
}
