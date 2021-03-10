package autodns_test

import (
	"github.com/jpbede/go-autodns"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := autodns.New("abc", "abc-123", 1)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
	c, err := autodns.NewWithClient(&http.Client{}, "abc", "abc-123", 1)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
	c, err := autodns.NewWithOptions(autodns.WithAPIEndpoint(autodns.APIEndpointProduction), autodns.WithCredentials("abc", "abc-123", 1))
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
