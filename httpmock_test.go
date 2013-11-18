package httpmock

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	var resp *http.Response
	var err error

	f := func() { resp, err = http.Get("http://example.com") }

	responder := func(req *http.Request) (*http.Response, error) {
		if req.URL.String() == "http://example.com" {
			return &http.Response{Status: "200", Body: NewBody("Hello, World!")}, nil
		}
		return nil, nil
	}

	Activate(f, responder)

	if err != nil {
		t.Error("Expected no error")
	}

	if resp.Status != "200" {
		t.Error("expected 200, got ", resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if !bytes.Equal(body, []byte("Hello, World!")) {
		t.Error("Expected 'Hello, World!', got ", string(body))
	}
}

func TestNoResponderFound(t *testing.T) {
	var resp *http.Response
	var err error

	f := func() { resp, err = http.Get("http://example.com") }

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected to panic")
		}
	}()

	Activate(f, nil)
}

func TestResponderReturnsNil(t *testing.T) {
	var resp *http.Response
	var err error

	f := func() { resp, err = http.Get("http://example.com?foo=bar") }

	responder := func(req *http.Request) (*http.Response, error) {
		return nil, nil
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected to panic")
		}
		if r != "URL not registered: http://example.com?foo=bar" {
			t.Error("Got", r)
		}
	}()

	Activate(f, responder)
}
