package main

import (
	"net/http"
	"unicode"
	"unicode/utf8"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

// UpCodec creates a CodecRequest to process each request.
type UpCodec struct {
}

// NewUpCodec returns a new UpCodec.
func NewUpCodec() *UpCodec {
	return &UpCodec{}
}

// NewRequest returns a new CodecRequest of type UpCodecRequest.
func (c *UpCodec) NewRequest(r *http.Request) rpc.CodecRequest {
	outerCR := &UpCodecRequest{}   // Our custom CR
	jsonC := json2.NewCodec()      // json Codec to create json CR
	innerCR := jsonC.NewRequest(r) // create the json CR, sort of.

	// NOTE - innerCR is of the interface type rpc.CodecRequest.
	// Because innerCR is of the rpc.CR interface type, we need a
	// type assertion in order to assign it to our struct field's type.
	// We defined the source of the interface implementation here, so
	// we can be confident that innerCR will be of the correct underlying type
	outerCR.CodecRequest = innerCR.(*json2.CodecRequest)
	return outerCR
}

// UpCodecRequest decodes and encodes a single request. UpCodecRequest
// implements gorilla/rpc.CodecRequest interface primarily by embedding
// the CodecRequest from gorilla/rpc/json. By selectively adding
// CodecRequest methods to UpCodecRequest, we can modify that behaviour
// while maintaining all the other remaining CodecRequest methods from
// gorilla's rpc/json implementation
type UpCodecRequest struct {
	*json2.CodecRequest
}

// Method returns the decoded method as a string of the form "Service.Method"
// after checking for, and correcting a lowercase method name
// By being of lower depth in the struct , Method will replace the implementation
// of Method() on the embedded CodecRequest. Because the request data is part
// of the embedded json.CodecRequest, and unexported, we have to get the
// requested method name via the embedded CR's own method Method().
// Essentially, this just intercepts the return value from the embedded
// gorilla/rpc/json.CodecRequest.Method(), checks/modifies it, and passes it
// on to the calling rpc server.
func (c *UpCodecRequest) Method() (string, error) {
	method, err := c.CodecRequest.Method()
	if len(method) > 1 && err == nil {
		r, n := utf8.DecodeRuneInString(method) // get the first rune, and it's length
		if unicode.IsLower(r) {
			upMethod := "rpc" + "." + string(unicode.ToUpper(r)) + method[n:]
			return upMethod, err
		}
	}
	return method, err
}
