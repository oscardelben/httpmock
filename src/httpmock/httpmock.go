// Package httpmock provides http mocking functionalities to Go.
package httpmock

import (
	"bytes"
	"io"
	"net/http"
)

// Mock Transfer objet used by the mock http client
type MockTransport struct {
	Responder func(*http.Request) (*http.Response, error)
}

type Body struct {
	io.Reader
}

func (Body) Close() error { return nil }

// Returns a buffered string of type io.Reader
// This is a convenience method that can be used in tests
//    http.Response{Body: httpmock.NewBody("hello")
func NewBody(body string) Body {
	return Body{bytes.NewBufferString(body)}
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Responder != nil {
		response, err := t.Responder(req)

		if response == nil {
			panic("URL not registered: " + req.URL.String())
		}

		return response, err
	}

	panic("[httpmock] No responder registered")
}

// TODO: this is not threadsafe, execute in a mutex
func Activate(f func(), responder func(*http.Request) (*http.Response,
	error)) {
	originalClient := http.DefaultClient

	mockClient := &http.Client{Transport: &MockTransport{Responder: responder}}

	http.DefaultClient = mockClient

	f()

	http.DefaultClient = originalClient
}
