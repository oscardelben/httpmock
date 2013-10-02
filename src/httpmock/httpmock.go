package httpmock

import (
  "io"
  "bytes"
  "net/http"
)

type MockTransport struct{
  Responder func(*http.Request) (*http.Response, error)
}

type Body struct {
  io.Reader
}

func (Body) Close() error { return nil }

// Used to create Body strings that can be passed to http.Response.
//    http.Response{Body: httpmock.NewBody("hello")
func NewBody(body string) Body {
  return Body{bytes.NewBufferString(body)}
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
  if t.Responder != nil {
    return t.Responder(req)
  }

  panic("no responder registered")
}

// TODO: this is not threadsafe, execute in a mutex
func Activate(f func(), responder func(*http.Request) (*http.Response,
error)) {
  originalClient := http.DefaultClient
  mockClient := &http.Client{Transport: &MockTransport{Responder:
  responder}}
  http.DefaultClient = mockClient

  f()

  http.DefaultClient = originalClient
}
