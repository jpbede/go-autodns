package transport

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestOption is a special type of function. It is used to modify a HTTP reques
type RequestOption func(*http.Request) error

// WithQueryValue adds a query parameter to a request's URL.
func WithQueryValue(key, value string) RequestOption {
	return func(req *http.Request) error {
		q := req.URL.Query()
		q.Set(key, value)

		req.URL.RawQuery = q.Encode()
		return nil
	}
}

// WithJSONRequestBody adds a JSON body to a request.
func WithJSONRequestBody(in interface{}) RequestOption {
	return func(req *http.Request) error {
		if in == nil {
			return nil
		}

		buf := bytes.Buffer{}
		enc := json.NewEncoder(&buf)
		err := enc.Encode(in)

		if err != nil {
			return err
		}

		rc := ioutil.NopCloser(&buf)

		copyBuf := buf.Bytes()

		req.Body = rc
		req.Header.Set("Content-Type", "application/json")
		req.ContentLength = int64(buf.Len())
		req.GetBody = func() (io.ReadCloser, error) {
			r := bytes.NewReader(copyBuf)
			return ioutil.NopCloser(r), nil
		}

		return nil
	}
}
