package contact_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.bnck.me/autodns/apis/contact"
	"go.bnck.me/autodns/internal/transport"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Create(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/contact", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)

		bodyBytes, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		var cActual contact.Contact
		err = json.Unmarshal(bodyBytes, &cActual)
		assert.NoError(t, err)

		cExpected := contact.Contact{
			Type:      contact.ContactTypePerson,
			Firstname: "Max",
			Lastname:  "Mustermann",
			Address: []string{
				"Musterstraße 4",
			},
			PostCode: "12345",
			City:     "Musterhausen",
			EMail:    "max@example.org",
		}
		assert.Equal(t, cExpected, cActual)

		rw.Write([]byte("{\"stid\":\"20210312-app1-demo-55742\",\"status\":{\"code\":\"S0301\",\"text\":\"Domain-Kontakt wurde erfolgreich angelegt.\",\"type\":\"SUCCESS\"},\"object\":{\"type\":\"Contact\",\"value\":\"31364493\"},\"data\":[{\"created\":\"2021-03-12T22:23:21.000+0100\",\"updated\":\"2021-03-12T22:23:21.000+0100\",\"id\":31364493,\"owner\":{\"context\":1,\"user\":\"2021_03_11_jpbe_la\"},\"updater\":{\"context\":1,\"user\":\"2021_03_11_jpbe_la\"},\"alias\":\"Max-Mustermann\",\"type\":\"PERSON\",\"city\":\"Musterhausen\",\"country\":\"DE\",\"email\":\"max@example.org\",\"protection\":\"SHOW_NONE\",\"domainsafe\":false,\"fname\":\"Max\",\"lname\":\"Mustermann\",\"address\":[\"Musterstraße 4\"],\"pcode\":\"12345\"}]}"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	c := contact.Contact{
		Type:      contact.ContactTypePerson,
		Firstname: "Max",
		Lastname:  "Mustermann",
		Address: []string{
			"Musterstraße 4",
		},
		PostCode: "12345",
		City:     "Musterhausen",
		EMail:    "max@example.org",
	}

	resp, err := cl.Create(c, context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "31364493", resp.Object.Value)
}

func TestClient_Create_InvalidJson(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("no json"))
	}))

	tc := transport.New(srv.URL)
	tc.HTTPClient = srv.Client()
	tc.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := contact.New(tc)

	c := contact.Contact{
		Type:      contact.ContactTypePerson,
		Firstname: "Max",
		Lastname:  "Mustermann",
		Address: []string{
			"Musterstraße 4",
		},
		PostCode: "12345",
		City:     "Musterhausen",
		EMail:    "max@example.org",
	}

	_, err := cl.Create(c, context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}
