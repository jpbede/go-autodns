package contact_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.bnck.me/autodns/apis/contact"
	"go.bnck.me/autodns/internal/transport"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_All(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/contact/_search", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)

		rw.Write([]byte("{\"stid\":\"20210312-app1-demo-55340\",\"status\":{\"code\":\"S0304\",\"text\":\"Die Daten des Domain-Kontaktes wurden erfolgreich ermittelt.\",\"type\":\"SUCCESS\"},\"object\":{\"type\":\"Contact\",\"summary\":1},\"data\":[{\"created\":\"2021-03-12T15:06:54.000+0100\",\"updated\":\"2021-03-12T22:07:49.000+0100\",\"id\":31364475,\"owner\":{\"context\":1,\"user\":\"2021_03_11_jpbe_la\"},\"alias\":\"Max Mustermann\",\"type\":\"PERSON\",\"organization\":\"\",\"title\":\"\",\"city\":\"Musterhausen\",\"country\":\"DE\",\"state\":\"DE\",\"fname\":\"Jan-Philipp\",\"lname\":\"Benecke\",\"address\":[\"Musterstraße 1\"],\"pcode\":\"12345\"}]}"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	resp, err := cl.All(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp, 1)
}

func TestClient_All_InvalidJson(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("no json"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	_, err := cl.All(context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}
