package fun

import (
	"sync"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Log(Random())
	t.Log(RandomInt(1, 3))
	t.Log(RandomInt(2, 4))
	t.Log(RandomInt64(10, 20))
	t.Log(RandomNumber(10))
	t.Log(RandomLetter(10))
	t.Log(RandomString(10))
}

func TestRandomConcurrent(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = Random()
			_ = RandomInt(1, 100)
			_ = RandomInt64(1, 100)
			_ = RandomString(16)
		}()
	}

	wg.Wait()
}
