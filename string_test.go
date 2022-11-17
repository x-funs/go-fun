package fun

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPad(t *testing.T) {
	assert.Equal(t, "12345", PadLeft("12345", "", 10))
	assert.Equal(t, "12345", PadLeft("12345", " ", 3))

	assert.Equal(t, "0000012345", PadLeft("12345", "0", 10))
	assert.Equal(t, "0101012345", PadLeft("12345", "01", 10))
	assert.Equal(t, "1234500000", PadRight("12345", "0", 10))
	assert.Equal(t, "0012345000", PadBoth("12345", "0", 10))
	assert.Equal(t, "12345678901", PadBoth("12345678901", "0", 10))
	assert.Equal(t, "0001230000", PadBoth("123", "0", 10))
}

func TestBlank(t *testing.T) {
	assert.True(t, Blank(""))
	assert.True(t, Blank("  "))
	assert.True(t, Blank("	"))
	assert.True(t, Blank("	       "))
	assert.True(t, BlankAny("a", "b", "		", "123"))
	assert.True(t, BlankAll("", "  ", "		"))
	assert.False(t, BlankAll("", "  ", "		", "123"))
}

func TestSplitTrim(t *testing.T) {
	assert.Equal(t, []string{}, SplitTrim("abc", ""))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("a b c", " "))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("a,,b,c", ","))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("  a ,  , b ,		c", ","))

	assert.Equal(t, []int{2, 3, 5}, SplitTrimToInts("2,,3,5", ","))

	assert.Equal(t, []string{}, SplitTrim("", "/"))
	assert.Equal(t, []string{}, SplitTrim("/", "/"))
	assert.Equal(t, []string{"index"}, SplitTrim("/index", "/"))
	assert.Equal(t, []string{"abc"}, SplitTrim("abc", "/"))
}

func TestContains(t *testing.T) {
	assert.Equal(t, true, Contains("", ""))
	assert.Equal(t, true, Contains("hello", "el"))
	assert.Equal(t, true, Contains("hello", ""))
	assert.Equal(t, false, Contains("hello", " "))
	assert.Equal(t, false, Contains("hello", "a"))

	assert.Equal(t, true, ContainsCase("HELLO", "el"))
	assert.Equal(t, true, ContainsAny("hello", "aa", "el"))
	assert.Equal(t, false, ContainsAny("hello", "aa", "eo"))

}

func TestReverse(t *testing.T) {
	assert.Equal(t, "", Reverse(""))
	assert.Equal(t, "olleh", Reverse("hello"))
}

func TestCamel(t *testing.T) {
	assert.Equal(t, "_abc", SnakeToCamel("_abc", true))
	assert.Equal(t, "AAbc", SnakeToCamel("a_abc", true))
	assert.Equal(t, "AAbc", SnakeToCamel("a__abc", true))
	assert.Equal(t, "TestAbc", SnakeToCamel("test_abc_", true))
	assert.Equal(t, "TestAbc", SnakeToCamel("test_abc", true))
	assert.Equal(t, "TestAbc", SnakeToCamel("Test_Abc", true))
	assert.Equal(t, "TestAbcDe", SnakeToCamel("test_aBC_DE", true))

	assert.Equal(t, "testAbc", SnakeToCamel("test_abc", false))
	assert.Equal(t, "aAbc", SnakeToCamel("a__abc", false))

	assert.Equal(t, "test_abc_de", CamelToSnake("TestAbcDe"))
	assert.Equal(t, "test_abc_de", CamelToSnake("testAbcDe"))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, "ac", Remove("abc", "b"))
	assert.Equal(t, "a", RemoveAny("abc", "b", "c"))
	assert.Equal(t, "abcdefg", RemovePrefix("abcdefg", ""))
	assert.Equal(t, "abcdefg", RemovePrefix("abcdefg", "b"))
	assert.Equal(t, "cdefg", RemovePrefix("abcdefg", "ab"))
	assert.Equal(t, "abcd", RemoveSuffix("abcdefg", "efg"))
}

func TestSubString(t *testing.T) {
	assert.Equal(t, "abcdefg", SubString("abcdefg", 0, 0))
	assert.Equal(t, "bcde", SubString("abcdefg", 1, 4))
	assert.Equal(t, "abcdefg", SubString("abcdefg", 0, 100))
	assert.Equal(t, "abcdefg", SubString("abcdefg", 0, -1))
	assert.Equal(t, "试he", SubString("测试hello中文", 1, 3))
}

func TestWrap(t *testing.T) {
	assert.Equal(t, "`abcdefg`", Wrap("abcdefg", "`"))
	assert.Equal(t, "abcdefg", Unwrap("`abcdefg`", "`"))
	assert.Equal(t, "`abcdefg`", Unwrap("``abcdefg``", "`"))
}

func TestRemoveLines(t *testing.T) {
	assert.Equal(t, "acb", RemoveLines("a\n\nc\nb"))
}

func BenchmarkRemoveLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveLines("https://www.163.com/news/article/HE8G1HQ8000189FH.html")
	}
}

func TestNormaliseSpace(t *testing.T) {
	t.Log(NormaliseSpace("中   国\n世   	界\n\n\n\n\n, hello      world     "))
	t.Log(NormaliseSpace("\n  \n\n\n\n\n    "))
}

func TestNormaliseLine(t *testing.T) {
	t.Log(NormaliseLine("中   国\n世   	界\n\n\n\n\n, hello      world     "))
}

func TestRemoveSign(t *testing.T) {
	text := ",.!，，D_NAME。！；‘’”“\n《》\r\n**dfs#%^&()-+		我1431221     中国 123漢字\n\n\nかどうかのjavaを<決定>$¥"
	fmt.Println(RemoveSign(text))
}

func TestHasPrefixSuffix(t *testing.T) {
	// assert.Equal(t, true, HasPrefixCase("Abc", "ab"))
	assert.Equal(t, true, HasPrefixCase("http://d.house.163.com/{cityCode}/", "http"))
	assert.Equal(t, true, HasPrefixCase("Abc", "ab"))
	assert.Equal(t, false, HasPrefixCase("Abc", "bc"))
	assert.Equal(t, true, HasSuffixCase("Abc", "BC"))
}

func TestTemplate(t *testing.T) {
	tpl := `{
  "name": {{ .name }},
  "age": {{ .age }}
}`

	data := map[string]string{
		"name": `"张三"`,
		"age":  "18",
	}

	t.Log(Template(tpl, data))
}
