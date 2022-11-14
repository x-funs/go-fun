package fun

import "testing"

func TestRandom(t *testing.T) {
	t.Log(Random())
	t.Log(RandomInt(1, 3))
	t.Log(RandomInt64(10, 20))
	t.Log(RandomNumber(10))
	t.Log(RandomLetter(10))
	t.Log(RandomString(10))
}
