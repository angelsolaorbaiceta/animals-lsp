package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type RequestMethod string

const (
	ReqMethodInit             RequestMethod = "initialize"
	ReqMethodDocCompletion    RequestMethod = "textDocument/completion"
	ReqMethodDocHover         RequestMethod = "textDocument/hover"
	ReqMethodDocSignatureHelp RequestMethod = "textDocument/signatureHelp"
	ReqMethodDocDefinition    RequestMethod = "textDocument/definition"
)

var contentLenRe = regexp.MustCompile(`(?i)^Content-Length:\s*(\d+)\s*$`)

type Request struct {
	JSONRPC string          `json:"jsonrpc"` // Always "2.0"
	ID      int             `json:"id"`
	Method  RequestMethod   `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// readRequest reads a request from the reader. A request has two parts:
//  1. Headers: "<Name>: <Value>\r\n"
//  2. Content: A JSON-RPC payload
//
// NOTE: Header names are case-insensitive as per RFC 7230, section 3.2.
// https://datatracker.ietf.org/doc/html/rfc7230#section-3.2
//
// > Each header field consists of a case-insensitive field name followed
// > by a colon (":"), optional leading whitespace, the field value, and
// > optional trailing whitespace.
//
// The header and content are separated by "\r\n".
//
// According to the protocol, there are currently two supported headers:
//   - `Content-Length: <int>` -- The lenght of the content part in bytes.
//   - `Content-Type`: <string> -- Defaults to `application/vscode-jsonrpc; charset=utf-8`.
func readRequest(reader io.Reader) (*Request, error) {
	var (
		bufReader  = bufio.NewReader(reader)
		contentLen = 0
		request    Request
	)

	// 1. Read headers until the empty line (but we only recognize the content-length)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil, err
			}

			return nil, fmt.Errorf("failed to read header: %w", err)
		}

		// Trim whitespace (handles both \r\n and \n line endings)
		line = strings.TrimSpace(line)

		// A blank line signals the end of the headers section
		if line == "" {
			break
		}

		// Try to parse the content length header
		matches := contentLenRe.FindStringSubmatch(line)
		if len(matches) == 2 {
			length, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse the content length value: %w", err)
			}

			contentLen = length
		}
	}

	// 2. Make sure we got a content length header
	if contentLen == 0 {
		return nil, fmt.Errorf("missing Content-Length header")
	}

	// 3. Read contentLen bytes from the JSON payload
	content := make([]byte, contentLen)
	if _, err := io.ReadFull(bufReader, content); err != nil {
		return nil, fmt.Errorf("failed to read content: %s", err)
	}

	if err := json.Unmarshal(content, &request); err != nil {
		return nil, fmt.Errorf("failed to parse JSON request: %w", err)
	}

	return &request, nil
}
