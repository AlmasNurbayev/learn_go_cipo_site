package main

import (
	"cipo_cite_server/internal/parser"
)

var Version = "v0.1.0"

func main() {

	p := parser.New(Version)
	p.Init()
	p.Run()
}
