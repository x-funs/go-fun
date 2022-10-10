package datastruct

import (
	"fmt"
	"testing"
)

func TestTire_FindAll(t *testing.T) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box")

	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。to box, into box"
	all := tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true})
	fmt.Printf("%+v\n", all)
}

func BenchmarkTire_FindAll(b *testing.B) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box")
	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。to box, into box"

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true})
	}
	b.StopTimer()
}
