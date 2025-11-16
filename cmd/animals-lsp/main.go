package main

import (
	"log"
	"os"
)

func main() {
	// Redirect log to stderr so it doesn't interfere with LSP IO
	log.SetOutput(os.Stderr)
	log.SetPrefix("[animals-lsp]")
	log.Println("Animals LSP server started.")

	// reader := bufio.NewReader(os.Stdin)
}
