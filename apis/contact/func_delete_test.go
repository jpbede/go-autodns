package contact_test

import (
	"context"
	"github.com/jpbede/go-autodns/apis/contact"
	"github.com/jpbede/go-autodns/internal/transport"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Delete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/contact/31364492", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)

		rw.Write([]byte("{\"stid\":\"20210312-app1-demo-55588\",\"status\":{\"code\":\"S0303\",\"text\":\"Domain-Kontakt wurde erfolgreich geloescht.\",\"type\":\"SUCCESS\"},\"object\":{\"type\":\"Contact\",\"value\":\"31364492\"}}"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	resp, err := cl.Delete(31364492, context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "31364492", resp.Object.Value)
}

func TestClient_Delete_InvalidJson(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("no json"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	_, err := cl.Delete(31364492, context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}
