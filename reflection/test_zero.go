package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsZero(t *testing.T) {
	assert.Equal(t, IsZero(0), true)
	assert.Equal(t, IsZero(1), false)
	assert.Equal(t, IsZero(-1), false)

	assert.Equal(t, IsZero(""), true)
	assert.Equal(t, IsZero("0"), false)
	assert.Equal(t, IsZero(" "), false)

	assert.Equal(t, IsZero([]string{}), true)
	assert.Equal(t, IsZero([]string{""}), false)
}
