package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMin(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, int64(1), MinInt64(1, 2))
	assert.Equal(t, int64(2), MaxInt64(1, 2))

	assert.Equal(t, 1, MinN(1, 2, 3, 4, 5, 6, 7, 9))
	assert.Equal(t, 9, MaxN(1, 2, 3, 4, 5, 6, 7, 9))
	assert.Equal(t, int64(1), MinN[int64](1, 2, 3, 4, 5, 6, 7, 9))
	assert.Equal(t, int64(9), MaxN[int64](1, 2, 3, 4, 5, 6, 7, 9))
	assert.Equal(t, 1.36, MinN(1.36, 2.69, 5.68, 8.62, 22.96))
	assert.Equal(t, 22.96, MaxN(1.36, 2.69, 5.68, 8.62, 22.96))
}
