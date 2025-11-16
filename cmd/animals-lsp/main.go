package main

import (
	"log"
	"os"

	"github.com/angelsolaorbaiceta/animals-lsp/internal"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("[animals-lsp]")
	log.Println("Animals LSP server starting...")

	internal.MakeStdIOServer().Start()
}
