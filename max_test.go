package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxMin(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, int64(1), MinInt64(1, 2))
	assert.Equal(t, int64(2), MaxInt64(1, 2))
}
