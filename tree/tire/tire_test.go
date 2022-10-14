package tire

import (
	"fmt"
	"sort"
	"testing"
)

var (
	textZh   = "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。"
	textEn   = "The English Garden is an e-learning platform for learning the English language. to box, into box, sport watching tv, It is a fully interactive."
	textZhEn = "去上海帮忙挖土豆，土豆地瓜哪里挖，The English Garden is an e-learning platform for learning the English language. to box, into box, sport watching tv, It is a fully interactive. 一挖一麻袋。"
	textNone = "The English Garden is an, e-learning platform for learning the English language"

	keywords = []string{
		"挖土豆",
		"上海",
		"上海帮忙",
		"土豆",
		"to",
		"to box",
		"watch",
	}
)

func TestTireContains(t *testing.T) {
	tireTree := NewTire()
	tireTree.AddAll(keywords)

	contains := tireTree.Contains(textZh, false)
	t.Log(contains)

	containsWord := tireTree.Contains(textZhEn, true)
	t.Log(containsWord)

	containsNone := tireTree.Contains(textNone, false)
	t.Log(containsNone)
}

func TestTireFind(t *testing.T) {
	tireTree := NewTire()
	tireTree.AddAll(keywords)

	containsZh := tireTree.Find(textZh, false)
	t.Log(containsZh)

	containsEn := tireTree.Find(textEn, true)
	t.Log(containsEn)

	containsZhEn := tireTree.Find(textZhEn, false)
	t.Log(containsZhEn)

	containsWord := tireTree.Find(textZhEn, true)
	t.Log(containsWord)
}

func TestTireFindAll(t *testing.T) {
	tireTree := NewTire()
	tireTree.AddAll(keywords)

	foundZh := tireTree.FindWithOptions(textZh, Option{Limit: -1, Greed: true, Density: true})
	t.Log(foundZh)

	foundEn := tireTree.FindWithOptions(textEn, Option{Limit: -1, Greed: true, Density: true, WordMode: true})
	t.Log(foundEn)
}

func TestSeparator(t *testing.T) {
	runes := []rune(" []{}()【】（）,.《》<>，。、：:；;'\"‘’“”|=-/*！@#%…&*!~`·")
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	for _, r := range runes {
		fmt.Printf("%d(%s), ", r, string(r))
	}
	fmt.Println()
}
