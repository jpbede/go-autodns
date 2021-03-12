package autodns_test

import (
	"github.com/jpbede/go-autodns"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Stat(t *testing.T) {
	c, err := autodns.New("abc", "abc-123", 1)
	assert.NoError(t, err)

	statAPI := c.Domain()
	if statAPI == nil {
		t.Error("Failed to get 'Domain' endpoint")
	}
}
