### HttpMock

HttpMock is a mocking library for http requests. It works by overwriting the
http default client so that you can provide your own responder.

Example:

```go
var resp *http.Response
var err error

f := func() { resp, err = http.Get("http://example.com") }

responder := func(req *http.Request) (*http.Response, error) {
   if req.URL.Host == "example.com" {
     return &http.Response{Status: "200", Body:
     httpmock.NewBody("Hello, World!")}, nil
   }
   return &http.Response{}, nil
}

httpmock.Activate(f, responder)
```

### Status

Work in progress. Not all the api has been defined yet but it's a good starting
point if you need to test http interactions.
