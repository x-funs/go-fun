package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomUserAgent(t *testing.T) {
	assert.NotEmpty(t, UserAgentRandom())
	assert.NotEmpty(t, UserAgentRandom())
}

func TestRandomMobileUserAgent(t *testing.T) {
	assert.NotEmpty(t, UserAgentRandomMobile())
	assert.NotEmpty(t, UserAgentRandomMobile())

}
