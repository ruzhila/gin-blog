package themes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRender(t *testing.T) {
	r, err := NewRender()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}
