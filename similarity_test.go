package fun

import (
	"testing"
)

func TestLcs(t *testing.T) {
	t.Log(Lcs("ABC", "B"))
	t.Log(Lcs("ABC", "AC"))
	t.Log(Lcs("ABC", "CA"))
	t.Log(Lcs("ABCDEF", "CD"))
	t.Log(Lcs("ABCDEFGHIG", "CEGML"))
	t.Log(Lcs("ABCDEFGHIG", "GIH"))
	t.Log(Lcs("我是中国人我说中国话", "中国人米说"))
}
