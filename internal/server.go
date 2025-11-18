package internal

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type LangServer struct {
	reader io.Reader
	writer io.Writer
	// The initialization parameters received from the client.
	// This is nil before the initialization messages have been exchanged.
	initParams *InitializeParams
}

func MakeServer(reader io.Reader, writer io.Writer) *LangServer {
	return &LangServer{
		reader: reader,
		writer: writer,
	}
}

// MakeStdIOServer creates a new server that uses Stdin and Stdout for exchanging
// messages.
func MakeStdIOServer() *LangServer {
	return MakeServer(os.Stdin, os.Stdout)
}

// Start reads from the reader indifinitely until the EOF, and writes responses
// to the writer. This method blocks.
func (s *LangServer) Start() {
	for {
		request, err := readRequest(s.reader)
		if err != nil {
			if err == io.EOF {
				log.Println("EOF reached. Client disconnected.")
				return
			}

			log.Printf("Error reading request: %v", err)
			continue
		}

		s.handleRequest(request)
	}
}

func (s *LangServer) handleRequest(req *Request) {
	switch req.Method {
	case ReqMethodInit:
		var params InitializeParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			log.Printf("Can't parse initialization params from request with ID: %v", req.ID)
		}
		s.handleInit(params, req.ID)
	default:
		log.Printf("Received unhandled method: %s with ID: %v", req.Method, req.ID)
	}
}

func (s *LangServer) handleInit(params InitializeParams, requestID int) {
	response := makeResultResponse(&requestID, makeInitializeResult())
	if err := writeResponse(s.writer, response); err != nil {
		log.Printf("Error marshaling initialize response: %v", err)
		return
	}

	s.initParams = &params
}
