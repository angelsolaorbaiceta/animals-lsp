package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func loadTestData(t *testing.T, filename string) (string, int) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("test_data", filename))
	if err != nil {
		t.Fatalf("failed to read test data %s: %v", filename, err)
	}

	return string(data), len(data)
}

func makeContentLengthHeader(length int) string {
	return fmt.Sprintf("Content-Length: %d\r\n", length)
}

func makeReqReaderFromString(content string) io.Reader {
	header := makeContentLengthHeader(len(content))
	return bufio.NewReader(strings.NewReader(header + "\r\n" + content))
}

func makeReqReaderFromFile(t *testing.T, filename string) io.Reader {
	var (
		content, length = loadTestData(t, filename)
		header          = makeContentLengthHeader(length)
	)

	return bufio.NewReader(strings.NewReader(header + "\r\n" + content))
}
