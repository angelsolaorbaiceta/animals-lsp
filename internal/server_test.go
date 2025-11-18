package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitalize(t *testing.T) {
	var (
		reader = makeReqReaderFromFile(t, "initialize_request_00.json")
		writer = new(bytes.Buffer)
		server = MakeServer(reader, writer)
	)

	server.Start()

	var (
		gotParts      = strings.Split(writer.String(), doubleSeparator)
		gotHeader     = gotParts[0]
		gotContent    = gotParts[1]
		gotResponse   Response
		gotInitResult InitializeResult
	)

	assert.Equal(t, fmt.Sprintf("Content-Length: %d", len(gotContent)), gotHeader)

	if err := json.Unmarshal([]byte(gotContent), &gotResponse); err != nil {
		assert.Fail(t, "Couldn't unmarshall response content")
	}
	assert.Equal(t, jsonRpcVersion, gotResponse.JSONRPC)
	assert.Equal(t, 1123, *gotResponse.ID)
	assert.Nil(t, gotResponse.Error)

	if err := json.Unmarshal(gotResponse.Result.([]byte), &gotInitResult); err != nil {
		assert.Fail(t, "Couldn't unmarshall response content")
	}
}
