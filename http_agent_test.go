package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomUserAgent(t *testing.T) {
	assert.NotEmpty(t, UserAgentRandom())
	assert.NotEmpty(t, UserAgentRandom())
}

func TestRandomMobileUserAgent(t *testing.T) {
	assert.NotEmpty(t, UserAgentRandomMobile())
	assert.NotEmpty(t, UserAgentRandomMobile())

}
