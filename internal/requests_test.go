package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadRequestWithoutParams(t *testing.T) {
	reader := makeReqReaderFromString(
		"{ \"jsonrpc\": \"2.0\", \"id\": 34, \"method\": \"textDocument/completion\" }",
	)
	got, err := readRequest(reader)
	assert.Nil(t, err)

	want := &Request{
		JSONRPC: "2.0",
		ID:      34,
		Method:  "textDocument/completion",
	}

	assert.Equal(t, want, got)
}

func TestReadRequestWithParams(t *testing.T) {
	reader := makeReqReaderFromString(
		"{ \"jsonrpc\": \"2.0\", \"id\": 34, \"method\": \"textDocument/completion\", \"params\": { \"foo\": 3 } }",
	)
	got, err := readRequest(reader)
	assert.Nil(t, err)

	want := &Request{
		JSONRPC: "2.0",
		ID:      34,
		Method:  "textDocument/completion",
		Params:  []byte("{ \"foo\": 3 }"),
	}

	assert.Equal(t, want, got)
}
