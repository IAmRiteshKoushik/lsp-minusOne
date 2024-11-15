package rpc

import (
	"encoding/json"
	"fmt"
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
