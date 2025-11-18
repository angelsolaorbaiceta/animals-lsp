package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

type Response struct {
	JSONRPC string         `json:"jsonrpc"` // Always "2.0"
	ID      *int           `json:"id"`
	Result  any            `json:"result,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}

func makeResultResponse(ID *int, result any) *Response {
	return &Response{
		JSONRPC: jsonRPCVersion,
		ID:      ID,
		Result:  result,
	}
}

type ResponseError struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data,omitempty"`
}

func writeResponse(writer io.Writer, res *Response) error {
	content, err := json.Marshal(res)
	if err != nil {
		return err
	}

	header := fmt.Appendf(nil, "Content-Length: %d%s", len(content), doubleSeparator)
	message := make([]byte, 0, len(header)+len(content))
	message = append(message, header...)
	message = append(message, content...)

	_, err = writer.Write(message)
	if err != nil {
		return err
	}

	return nil
}
