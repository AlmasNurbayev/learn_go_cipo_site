package main

import (
	"cipo_cite_server/internal/server"
)

var Version = "v0.1.0"

func main() {
	s := server.New(Version)
	s.Init()
	s.Run()
}
