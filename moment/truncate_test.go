package moment

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTruncateByDay(t *testing.T) {
	loc := time.Now().Location()

	baseTime := time.Date(2021, 3, 23, 16, 19, 11, 111, loc)

	truncatedTime := TruncateByDay(baseTime)

	assert.Equal(t, truncatedTime, time.Date(2021, 3, 23, 0, 0, 0, 0, loc))
}
