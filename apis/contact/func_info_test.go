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

func TestClient_Info(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/contact/31364475", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		rw.Write([]byte("{\"stid\":\"20210312-app1-demo-55474\",\"status\":{\"code\":\"S0304\",\"text\":\"Die Daten des Domain-Kontaktes wurden erfolgreich ermittelt.\",\"type\":\"SUCCESS\"},\"object\":{\"type\":\"Contact\",\"value\":\"31364475\"},\"data\":[{\"created\":\"2021-03-12T15:06:54.000+0100\",\"updated\":\"2021-03-12T22:07:49.000+0100\",\"id\":31364475,\"owner\":{\"context\":1,\"user\":\"2021_03_11_jpbe_la\"},\"updater\":{\"context\":1,\"user\":\"2021_03_11_jpbe_la\"},\"alias\":\"Max Mustermann\",\"type\":\"PERSON\",\"organization\":\"\",\"title\":\"\",\"city\":\"Musterhausen\",\"country\":\"DE\",\"state\":\"DE\",\"email\":\"hello@example.com\",\"protection\":\"SHOW_NONE\",\"domainsafe\":false,\"linked\":true,\"fname\":\"Jan-Philipp\",\"lname\":\"Benecke\",\"address\":[\"Musterstraße 1\"],\"pcode\":\"12345\",\"nicRef\":[{\"nic\":{\"label\":\"DE-TEST\",\"name\":\"de\"},\"status\":\"SUCCESS\",\"type\":\"ALL\",\"name\":\"DEMO-31364475\"}]}]}"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	resp, err := cl.Info(31364475, context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "hello@example.com", resp.EMail)
}