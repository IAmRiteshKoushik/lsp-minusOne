package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// We need to encode and send the message to our LSP so that it can decode,
// perform some actions and return back some results. Our serialization standard
// is going to be JSON.
func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

// The bytes.Cut() function accepts two byte-slices and works as a partition
// funcion by identifying the first occurance of the subsequence of bytes in
// the msg. and partitioning around it. This is also the reason why we are using
// []byte as type of msg instead of doing strings. Also, []byte is faster to
// work with as compared to strings or chars
func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Did not find a separator")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, err
	}

	// TODO: Use content somewhere at a later point
	_ = content

	return contentLength, nil
}
