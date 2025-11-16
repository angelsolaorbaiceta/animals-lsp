package internal

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadRequestWithoutParams(t *testing.T) {
	var (
		content = "{ \"jsonrpc\": \"2.0\", \"id\": 34, \"method\": \"textDocument/completion\" }"
		header  = fmt.Sprintf("Content-Length: %d\r\n", len(content))
		sc      = bufio.NewReader(strings.NewReader(header + "\r\n" + content))
	)

	got, err := readRequest(sc)
	assert.Nil(t, err)

	wantID := 34
	want := &Request{
		JSONRPC: "2.0",
		ID:      &wantID,
		Method:  "textDocument/completion",
	}

	assert.Equal(t, want, got)
}

func TestReadRequestWithParams(t *testing.T) {
	var (
		content = "{ \"jsonrpc\": \"2.0\", \"id\": 34, \"method\": \"textDocument/completion\", \"params\": { \"foo\": 3 } }"
		header  = fmt.Sprintf("Content-Length: %d\r\n", len(content))
		sc      = bufio.NewReader(strings.NewReader(header + "\r\n" + content))
	)

	got, err := readRequest(sc)
	assert.Nil(t, err)

	wantID := 34
	want := &Request{
		JSONRPC: "2.0",
		ID:      &wantID,
		Method:  "textDocument/completion",
		Params:  []byte("{ \"foo\": 3 }"),
	}

	assert.Equal(t, want, got)
}
