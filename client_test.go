package autodns_test

import (
	"github.com/jpbede/go-autodns"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Domain(t *testing.T) {
	c, err := autodns.New("abc", "abc-123", 1)
	assert.NoError(t, err)

	domainAPI := c.Domain()
	assert.NotNil(t, domainAPI)
}

func TestClient_Contact(t *testing.T) {
	c, err := autodns.New("abc", "abc-123", 1)
	assert.NoError(t, err)

	contactAPI := c.Contact()
	assert.NotNil(t, contactAPI)
}

func TestClient_Job(t *testing.T) {
	c, err := autodns.New("abc", "abc-123", 1)
	assert.NoError(t, err)

	jobAPI := c.Job()
	assert.NotNil(t, jobAPI)
}
