package main

import (
	"embed"
	"log"

	"github.com/mdshack/shackelford.dev/cmd/server"
)

var (
	//go:embed assets/dist
	assets embed.FS
)

func main() {
	s := server.New(&assets)

	log.Fatal(s.Serve())
}
