package domain_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.bnck.me/autodns/apis/contact"
	"go.bnck.me/autodns/apis/domain"
	"go.bnck.me/autodns/internal/transport"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Create(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/domain")
		assert.Equal(t, req.Method, http.MethodPost)

		rw.Write([]byte{0x7b, 0x22, 0x73, 0x74, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x31, 0x30, 0x33, 0x31,
			0x32, 0x2d, 0x61, 0x70, 0x70, 0x31, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x34, 0x30, 0x31, 0x38, 0x38,
			0x22, 0x2c, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x63, 0x6f, 0x64, 0x65,
			0x22, 0x3a, 0x22, 0x4e, 0x30, 0x31, 0x30, 0x31, 0x22, 0x2c, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a,
			0x22, 0x44, 0x69, 0x65, 0x20, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2d, 0x52, 0x65, 0x67, 0x69, 0x73,
			0x74, 0x72, 0x69, 0x65, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x77, 0x75, 0x72, 0x64, 0x65, 0x20, 0x65, 0x72,
			0x66, 0x6f, 0x6c, 0x67, 0x72, 0x65, 0x69, 0x63, 0x68, 0x20, 0x67, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
			0x65, 0x74, 0x2e, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x4e, 0x4f, 0x54, 0x49,
			0x46, 0x59, 0x22, 0x7d, 0x2c, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74,
			0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x2c, 0x22, 0x76, 0x61,
			0x6c, 0x75, 0x65, 0x22, 0x3a, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x32, 0x32, 0x32,
			0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x5b,
			0x7b, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x31, 0x2d,
			0x30, 0x33, 0x2d, 0x31, 0x32, 0x54, 0x31, 0x36, 0x3a, 0x30, 0x36, 0x3a, 0x32, 0x34, 0x2e, 0x30, 0x37,
			0x33, 0x2b, 0x30, 0x31, 0x30, 0x30, 0x22, 0x2c, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
			0x3a, 0x22, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x30, 0x33, 0x2d, 0x31, 0x32, 0x54, 0x31, 0x36, 0x3a, 0x30,
			0x36, 0x3a, 0x32, 0x34, 0x2e, 0x30, 0x30, 0x30, 0x2b, 0x30, 0x31, 0x30, 0x30, 0x22, 0x2c, 0x22, 0x6f,
			0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a,
			0x31, 0x2c, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x30, 0x33,
			0x5f, 0x31, 0x31, 0x5f, 0x6a, 0x70, 0x62, 0x65, 0x5f, 0x6c, 0x61, 0x22, 0x7d, 0x2c, 0x22, 0x75, 0x70,
			0x64, 0x61, 0x74, 0x65, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22,
			0x3a, 0x31, 0x2c, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x30,
			0x33, 0x5f, 0x31, 0x31, 0x5f, 0x6a, 0x70, 0x62, 0x65, 0x5f, 0x6c, 0x61, 0x22, 0x7d, 0x2c, 0x22, 0x73,
			0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x22, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x22, 0x2c,
			0x22, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x31,
			0x2d, 0x30, 0x33, 0x2d, 0x31, 0x32, 0x54, 0x31, 0x36, 0x3a, 0x30, 0x36, 0x3a, 0x32, 0x33, 0x2e, 0x30,
			0x30, 0x30, 0x2b, 0x30, 0x31, 0x30, 0x30, 0x22, 0x2c, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x39, 0x33, 0x34,
			0x33, 0x33, 0x38, 0x34, 0x7d, 0x5d, 0x7d})
	}))

	transportC := transport.New(srv.URL)
	transportC.HTTPClient = srv.Client()
	transportC.Credentials = &transport.APICredentials{Username: "abc", Password: "abc123", Context: 1}
	cl := domain.New(transportC)

	ct := contact.Contact{
		ID: 31364475,
	}
	d := domain.Domain{
		Name:   "example222222.com",
		OwnerC: &ct,
		AdminC: &ct,
		TechC:  &ct,
		ZoneC:  &ct,
		Nameserver: []*domain.Nameserver{
			{
				Name: "a.example.com",
			},
			{
				Name: "b.example.com",
			},
		},
	}

	resp, err := cl.Create(d, context.Background())

	assert.NoError(t, err)
	assert.Equal(t, d.Name, resp.Object.Value)
}
