package transport

import (
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

// ClientLogging is http.RoundTripper to log client requests and responses
type ClientLogging struct {
	Logger *zerolog.Logger
}

// RoundTrip executes a single HTTP transaction, returning a Response for the provided Request.
func (cl *ClientLogging) RoundTrip(r *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := http.DefaultTransport.RoundTrip(r)
	elapsed := time.Now().Sub(start)

	evt := cl.Logger.WithLevel(zerolog.DebugLevel).
		Str("method", r.Method).
		Str("path", r.URL.String()).
		Dur("elapsed", elapsed)

	if res != nil {
		evt.Int("status", res.StatusCode).
			Int64("size", res.ContentLength)
	} else {
		evt.Int("status", -1).
			Int64("size", -1)
	}

	evt.Msg("autodns_request")
	return res, err
}
