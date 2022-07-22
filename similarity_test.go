package fun

import (
	"testing"
)

func TestLcs(t *testing.T) {
	t.Log(LongestCommonSubString("ABC", "B"))
	t.Log(LongestCommonSubString("ABC", "AC"))
	t.Log(LongestCommonSubString("ABC", "CA"))
	t.Log(LongestCommonSubString("ABCDEF", "CD"))
	t.Log(LongestCommonSubString("ABCDEFGHIG", "CEGML"))
	t.Log(LongestCommonSubString("ABCDEFGHIG", "GIH"))
	t.Log(LongestCommonSubString("我是中国人我说中国话", "中国人米说"))
}
