package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatches(t *testing.T) {
	assert.Equal(t, true, Matches("abc@abc.com", RegexEmail))
}
func TestMail(t *testing.T) {
	assert.Equal(t, true, Matches("dhjwauihdaiu@163.com", RegexEmail))
}
