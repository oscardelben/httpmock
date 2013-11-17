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
		if req.URL.Host == "example.com" {
			return &http.Response{Status: "200", Body: NewBody("Hello, World!")}, nil
		}
		return &http.Response{}, nil
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
