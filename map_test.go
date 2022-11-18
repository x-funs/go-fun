package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeysValues(t *testing.T) {
	assert.Equal(t, 3, len(MapKeys(map[string]string{"a": "1", "b": "2", "c": "3"})))
	assert.Equal(t, []string{"1"}, MapValues(map[string]string{"a": "1"}))
}

func TestMapMerge(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	m2 := map[string]int{
		"b": 4,
		"c": 3,
	}
	m3 := map[string]int{
		"d": 6,
	}
	assert.Equal(t, map[string]int{"a": 1, "b": 4, "c": 3, "d": 6}, MapMerge(m1, m2, m3))
}
