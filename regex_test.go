package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatches(t *testing.T) {
	assert.Equal(t, true, Matches("abc@abc.com", RegexEmail))
}
