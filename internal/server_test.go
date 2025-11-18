package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
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
		assert.Fail(t, "Couldn't unmarshal response content")
	}
	assert.Equal(t, jsonRPCVersion, gotResponse.JSONRPC)
	assert.Equal(t, 1123, *gotResponse.ID)
	assert.Nil(t, gotResponse.Error)

	resultJSON, err := json.Marshal(gotResponse.Result)
	assert.Nil(t, err)
	if err := json.Unmarshal(resultJSON, &gotInitResult); err != nil {
		assert.Fail(t, "Couldn't unmarshal result from response")
	}

	assert.Equal(t, *makeInitializeResult(), gotInitResult)
}
