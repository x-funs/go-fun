package tree

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/x-funs/go-fun"
)

func TestSeparator(t *testing.T) {
	runes := []rune(" []{}()【】（）,.《》<>，。、：:；;'\"‘’“”|=-/*！@#%…&*!~`·")
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	for _, r := range runes {
		fmt.Printf("%d(%s), ", r, string(r))
	}
	fmt.Println()
}

func TestTireFindAll(t *testing.T) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box").Add("watch")

	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。"
	fmt.Printf("%+v\n", tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true}))

	text = "to box, into box, sport watch, watching tv"
	fmt.Printf("%+v\n", tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true, WordGroup: true}))
}

func BenchmarkTireFindAll(b *testing.B) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box").Add("watch")

	wordPath := "./word.txt"
	if fun.IsExist(wordPath) {
		buf, _ := ioutil.ReadFile(wordPath)
		for _, word := range strings.Split(fun.String(buf), fun.LF) {
			if word != "" {
				tire.Add(word)
			}
		}
	}

	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。to box, into box, sport watch, watching tv"

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true})
	}
	b.StopTimer()
	b.ReportAllocs()
}
