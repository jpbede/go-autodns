package autodns_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/autodns"
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
	c, err := autodns.NewWithOptions(autodns.WithAPIEndpoint(autodns.APIEndpointDemo), autodns.WithCredentials("abc", "abc-123", 1))
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
