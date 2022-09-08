package fun

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	fmt.Println(MemoryBytes())
	fmt.Println(Memory(SizeKB))
	fmt.Println(Memory(SizeMB))
}

func TestTime(t *testing.T) {
	timeStamp := Timestamp()
	timeStampMill := Timestamp(true)

	fmt.Println(timeStamp)
	fmt.Println(timeStampMill)

	assert.NotEmpty(t, timeStamp)
	assert.NotEmpty(t, timeStampMill)
}

func TestMd5(t *testing.T) {
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", Md5(""))
	assert.Equal(t, "df10ef8509dc176d733d59549e7dbfaf", Md5("123456abc"))
	assert.Equal(t, "21232f297a57a5a743894a0e4a801fc3", Md5("admin"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", Md5("123456"))
	assert.Equal(t, "49ba59abbe56e057", Md5Bit16("123456"))
	assert.Equal(t, "a32b4da32d9a67a5", Md5Bit16("df"))
}

func TestSha(t *testing.T) {
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", Sha1(""))
	assert.Equal(t, "a172ffc990129fe6f68b50f6037c54a1894ee3fd", Sha1("123456abc"))
	assert.Equal(t, "931145d4ddd1811be545e4ac88a81f1fdbfaf0779c437efba16b884595274d11", Sha256("123456abc"))
	assert.Equal(t, "2545507ada3a26999b11ec0324ae76e0cdef629c4a28b24be3965d24e1c406180a8ef79626c77fb3f556bfd59ab54920", Sha384("123456abc"))
	assert.Equal(t, "8756869d440a13e93979197b5d7839c823de87c2b115bce0dee62030af3b5b63114a517f1ab02509bfefa88527369ae0ad7946990f27dcb37711a7d6fb9bc893", Sha512("123456abc"))
}

func TestDate(t *testing.T) {

	fmt.Println(Date(Timestamp()))
	fmt.Println(Date(DatetimeMilliPattern, Timestamp(true)))

	timeStamp := 1650732457
	now := time.Now()

	assert.NotEmpty(t, Date())
	assert.NotEmpty(t, Date(""))
	assert.NotEmpty(t, Date(timeStamp))
	assert.NotEmpty(t, Date(DatetimePattern, timeStamp))

	// 无效的参数, 返回空
	assert.Empty(t, Date(123, ""))
	assert.Empty(t, Date("", ""))
	assert.Empty(t, Date("", "", ""))

	assert.Equal(t, now.Format(DatetimePattern), Date())
	assert.Equal(t, now.Format(DatetimePattern), Date(DatetimePattern))
	assert.Equal(t, "2022-04-24 00:47:37", Date(timeStamp))
	assert.Equal(t, "2022-04-24 00:47:37", Date(DatetimePattern, timeStamp))
}

func TestPad(t *testing.T) {
	assert.Equal(t, "12345", PadLeft("12345", "", 10))
	assert.Equal(t, "12345", PadLeft("12345", " ", 3))

	assert.Equal(t, "0000012345", PadLeft("12345", "0", 10))
	assert.Equal(t, "0101012345", PadLeft("12345", "01", 10))
	assert.Equal(t, "1234500000", PadRight("12345", "0", 10))
	assert.Equal(t, "0012345000", PadBoth("12345", "0", 10))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, 0, ToInt(""))
	assert.Equal(t, 0, ToInt(" "))
	assert.Equal(t, 0, ToInt(" 123 "))
	assert.Equal(t, 123, ToInt("123"))
	assert.Equal(t, 123, ToInt("0123"))
	assert.Equal(t, 0, ToInt("1.1"))
	assert.Equal(t, -123, ToInt("-123"))
}

func TestToInt64(t *testing.T) {
	assert.Equal(t, int64(0), ToInt64(""))
	assert.Equal(t, int64(0), ToInt64(" "))
	assert.Equal(t, int64(0), ToInt64(" 123 "))
	assert.Equal(t, int64(123), ToInt64("123"))
	assert.Equal(t, int64(123), ToInt64("0123"))
	assert.Equal(t, int64(0), ToInt64("1.1"))
	assert.Equal(t, int64(0), ToLong("1.1"))
	assert.Equal(t, int64(-123), ToLong("-123"))
}

func TestToUnit(t *testing.T) {
	assert.Equal(t, uint(0), ToUint(""))
	assert.Equal(t, uint(123), ToUint("0123"))
	assert.Equal(t, uint8(0), ToUint8("-1"))
}

func TestToUtf8AndCharset(t *testing.T) {
	s := []byte{0xd7, 0xd4}
	t.Log(string(s))

	result, err := ToUtf8(s, "gbk")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(result))
	}

	ss := "Hello，世界"
	gbk, _ := Utf8To([]byte(ss), "gbk")
	t.Log(string(gbk))

	utf8, _ := ToUtf8(gbk, "gbk")
	t.Log(string(utf8))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, false, ToBool(""))
	assert.Equal(t, true, ToBool("true"))
	assert.Equal(t, false, ToBool("false"))
	assert.Equal(t, false, ToBool(" "))
	assert.Equal(t, false, ToBool("a"))
}

func TestBase64(t *testing.T) {
	assert.Equal(t, "", Base64Encode(""))
	assert.Equal(t, "MTIzNDU2YWJj", Base64Encode("123456abc"))
	assert.Equal(t, "aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M/aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1", Base64Encode("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu"))
	assert.Equal(t, "aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M_aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1", Base64UrlEncode("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu"))

	assert.Equal(t, "123456abc", Base64Decode("MTIzNDU2YWJj"))
	assert.Equal(t, "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu", Base64Decode("aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M/aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1"))
	assert.Equal(t, "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu", Base64UrlDecode("aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M_aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1"))
}

func TestBlank(t *testing.T) {
	assert.True(t, Blank(""))
	assert.True(t, Blank("  "))
	assert.True(t, Blank("	"))
	assert.True(t, BlankAny("a", "b", "		", "123"))
	assert.True(t, BlankAll("", "  ", "		"))
	assert.False(t, BlankAll("", "  ", "		", "123"))
}

func TestStrTotimeParse(t *testing.T) {
	dateStrs := []string{
		// "30 August 2022",
		// "02 Sep 2022",
		// "02 Sep 2022 11:40",
		// "02 Sep 2022 11:40:53am",
		// "02 Sep 2022 11:40:53am",
		"02 Sep 2022 11:40 aM",
		// "02 Sep 2022 11:40:53am",
		"Sep 02 2022 11:40:00 aM",
		"Sep 02 2022 11:40 aM",
		"Sep 02 2022 11:40",
	}
	for _, dateStr := range dateStrs {
		t.Log(Date(StrToTime(dateStr)))
	}
}

func TestStrToTime(t *testing.T) {

	date := "2015-04-06 16:03:03"
	timeStamp := StrToTime(date)

	assert.Equal(t, Timestamp(), StrToTime())
	assert.Equal(t, int64(0), StrToTime(""))
	assert.Equal(t, Timestamp(), StrToTime("now"))
	assert.Equal(t, date, Date(StrToTime(date)))

	assert.Equal(t, "2015-04-06 00:00:00", Date(StrToTime("2015-04-06")))

	assert.Equal(t, "2015-04-07 16:03:03", Date(StrToTime("tomorrow", timeStamp)))
	assert.Equal(t, "2015-04-06 17:03:03", Date(StrToTime("+1 hour", timeStamp)))
	assert.Equal(t, "2015-05-06 16:03:03", Date(StrToTime("+30 days", timeStamp)))
	assert.Equal(t, "2015-06-06 16:03:03", Date(StrToTime("+2 months", timeStamp)))
	assert.Equal(t, "2015-03-30 16:03:03", Date(StrToTime("-1 week", timeStamp)))
	assert.Equal(t, "2013-04-06 16:03:03", Date(StrToTime("-2 year", timeStamp)))
	assert.Equal(t, "2013-04-06 16:03:03", Date(StrToTime("-2 years", timeStamp)))

	assert.Equal(t, "2020-02-29 16:03:03", Date(StrToTime("-1 day", StrToTime("2020-03-01 16:03:03"))))
	assert.Equal(t, "2021-03-01 16:03:03", Date(StrToTime("+1 year", StrToTime("2020-02-29 16:03:03"))))

	assert.Equal(t, StrToTime(date), StrToTime("2015-04-06 16:03:03"))
	assert.Equal(t, StrToTime(date), StrToTime("2015年04月06日 16:03:03"))
	assert.Equal(t, StrToTime(date), StrToTime("2015年04月06 16:03:03"))
	assert.Equal(t, StrToTime(date), StrToTime("2015年04月06日16时03分03秒"))
	assert.Equal(t, StrToTime(date), StrToTime("2015年04月06 16时03分03"))
	assert.Equal(t, StrToTime(date), StrToTime("2015/04/06 16:03:03"))
	assert.Equal(t, StrToTime(date), StrToTime("2015-4-6 16:3:3"))
	assert.Equal(t, StrToTime(date), StrToTime("2015-4-06 16:03:3"))
	assert.Equal(t, StrToTime(date), StrToTime("2015-4-06 16:03:3"))

	dates := []string{
		"2022-01-24T14:19:02+08:00", // "yyyy-MM-dd'T'HH:mm:ssX"
		"2022-01-24T14:19:02+02:00", // "yyyy-MM-dd'T'HH:mm:ssX"
		"2022-01-24T14:19:02-0500",  // "yyyy-MM-dd'T'HH:mm:ssX"
		"2022-01-24T14:19:00Z",      // "yyyy-MM-dd'T'HH:mm:ssX"
		"2022-01-24T14:19:01z",      // "yyyy-MM-dd'T'HH:mm:ss.SSS'z'"
		"2022-01-24T14:19:01",       // "yyyy-MM-dd'T'HH:mm:ss"

		"2022-01-24T14:19:01-0700",     // "2006-01-02T15:04:05-0700"
		"2022-01-24T14:19:01-070000",   // "2006-01-02T15:04:05-070000"
		"2022-01-24T14:19:01-07",       // "2006-01-02T15:04:05-07"
		"2022-01-24T14:19:01+07:00",    // "2006-01-02T15:04:05-07:00"
		"2022-01-24T14:19:01+07:00:00", // "2006-01-02T15:04:05-07:00:00"
		"2022/01/24T14:19:01-0700",     // "2006/01/02T15:04:05-0700"
		"2022/01/24T14:19:01+070000",   // "2006/01/02T15:04:05-070000"
		"2022/01/24T14:19:01-07",       // "2006/01/02T15:04:05-07"
		"2022/01/24T14:19:01-07:00",    // "2006/01/02T15:04:05-07:00"
		"2022/01/24T14:19:01-07:00:00", // "2006/01/02T15:04:05-07:00:00"
		"2022-01-24T14:19:01Z",         // "2006-01-02T15:04:05Z"
		"2022-01-24T14:19:01z",         // "2006-01-02T15:04:05z"
		"2022/01/24T14:19:01Z",         // "2006/01/02T15:04:05Z"
		"2022/01/24T14:19:01z",         // "2006/01/02T15:04:05z"

	}

	for _, d := range dates {
		timeStamp := StrToTime(d)
		result := Date(timeStamp)
		fmt.Printf("%s -> %s\n", d, result)
	}
}

func TestIP2Long(t *testing.T) {
	assert.Equal(t, uint32(2130706433), Ip2Long("127.0.0.1"))
	assert.Equal(t, uint32(3221234342), Ip2Long("192.0.34.166"))
	assert.Equal(t, uint32(659439616), Ip2Long("39.78.64.0"))
	assert.Equal(t, uint32(659439617), Ip2Long("39.78.64.1"))
	assert.Equal(t, uint32(659439870), Ip2Long("39.78.64.254"))
	assert.Equal(t, uint32(659439871), Ip2Long("39.78.64.255"))

	assert.Equal(t, "39.78.64.255", Long2Ip(659439871))
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, true, Empty(nil))
	assert.Equal(t, true, Empty(0))
	assert.Equal(t, true, Empty(""))
	assert.Equal(t, true, Empty(false))

	assert.Equal(t, false, Empty(" "))
	assert.Equal(t, false, Empty(1))
	assert.Equal(t, false, Empty(true))

	var a1 [0]int
	a2 := [2]int{3, 5}
	assert.Equal(t, true, Empty(a1))
	assert.Equal(t, false, Empty(a2))

	var s1 []string
	s2 := []int{2, 3, 5}
	assert.Equal(t, true, Empty(s1))
	assert.Equal(t, false, Empty(s2))

	m1 := map[string]int{}
	m2 := map[string]string{"k1": "v1", "k2": "v2"}
	assert.Equal(t, true, Empty(m1))
	assert.Equal(t, false, Empty(m2))

	assert.Equal(t, true, EmptyAll("", nil))
	assert.Equal(t, true, EmptyAny("a", "", "b"))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "1", ToString(1))
	assert.Equal(t, "0.123", ToString(0.123))
	assert.Equal(t, "<nil>", ToString(nil))
	assert.Equal(t, "[1 2 3]", ToString([]int{1, 2, 3}))
}

func TestSplitTrim(t *testing.T) {
	assert.Equal(t, []string{}, SplitTrim("abc", ""))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("a b c", " "))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("a,,b,c", ","))
	assert.Equal(t, []string{"a", "b", "c"}, SplitTrim("  a ,  , b ,		c", ","))

	assert.Equal(t, []int{2, 3, 5}, SplitTrimToInts("2,,3,5", ","))

	assert.Equal(t, []string{}, SplitTrim("abc", "/"))
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, false, IsNumber(""))
	assert.Equal(t, true, IsNumber("1"))
	assert.Equal(t, true, IsNumber("012345"))
}

func TestIsLetter(t *testing.T) {
	assert.Equal(t, false, IsLetter(""))
	assert.Equal(t, false, IsLetter("   "))
	assert.Equal(t, true, IsLetter("a"))
	assert.Equal(t, true, IsLetter("abc"))
	assert.Equal(t, false, IsLetter("abc123"))
	assert.Equal(t, false, IsLetter("123"))
}

func TestIsEmail(t *testing.T) {
	assert.Equal(t, false, IsEmail(""))
	assert.Equal(t, true, IsEmail("aaa@aa.aa"))
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

func TestToJson(t *testing.T) {
	assert.Equal(t, `null`, ToJson(nil))
	assert.Equal(t, `""`, ToJson(""))
	assert.Equal(t, `"abc"`, ToJson("abc"))
	assert.Equal(t, `123`, ToJson(123))
	assert.Equal(t, `["a","1","b","2"]`, ToJson([]string{"a", "1", "b", "2"}))
	assert.Equal(t, `{"a":"1","b":"2"}`, ToJson(map[string]string{"a": "1", "b": "2"}))
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

func TestMatches(t *testing.T) {
	assert.Equal(t, true, Matches("abc@abc.com", RegexEmail))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, "ac", Remove("abc", "b"))
	assert.Equal(t, "a", RemoveAny("abc", "b", "c"))
	assert.Equal(t, "abcdefg", RemovePrefix("abcdefg", ""))
	assert.Equal(t, "abcdefg", RemovePrefix("abcdefg", "b"))
	assert.Equal(t, "cdefg", RemovePrefix("abcdefg", "ab"))
	assert.Equal(t, "abcd", RemoveSuffix("abcdefg", "efg"))
}

func TestRandom(t *testing.T) {
	t.Log(Random())
	t.Log(RandomInt(1, 3))
	t.Log(RandomInt64(10, 20))
	t.Log(RandomNumber(10))
	t.Log(RandomLetter(10))
	t.Log(RandomString(10))
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

func TestSliceContains(t *testing.T) {
	assert.Equal(t, true, SliceContains([]string{"a", "b", "c"}, "a"))
	assert.Equal(t, true, SliceContains([]int{3, 5, 7}, 7))
	assert.Equal(t, false, SliceContains([]int{3, 5, 7}, 2))
	assert.Equal(t, false, SliceContains([]int{}, 2))
}

func TestMapKeysValues(t *testing.T) {
	assert.Equal(t, 3, len(MapKeys(map[string]string{"a": "1", "b": "2", "c": "3"})))
	assert.Equal(t, []string{"1"}, MapValues(map[string]string{"a": "1"}))
}

func TestMapMerge(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	m2 := map[string]int{
		"b": 4,
		"c": 3,
	}
	assert.Equal(t, map[string]int{"a": 1, "b": 4, "c": 3, "d": 4}, MapMerge(m1, m2))
}

func TestSliceUnique(t *testing.T) {
	assert.Equal(t, 3, len(SliceUnique([]string{"a", "b", "c", "a", "b", "c"})))
	assert.Equal(t, []string{""}, SliceUnique([]string{"", "", ""}))
}

func TestSliceSplit(t *testing.T) {
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 1))
	assert.Equal(t, [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 2))
	assert.Equal(t, [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 3))
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 9))
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g"}}, SliceSplit([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
}

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, []int{}, SliceRemove([]int{}, 0))
	assert.Equal(t, []int{2, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 0))
	assert.Equal(t, []int{2, 0, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 1))
	assert.Equal(t, []string{"b", "c", "d"}, SliceRemove([]string{"a", "b", "c", "d"}, "a"))
	assert.Equal(t, []string{"a", "d"}, SliceRemove([]string{"a", "", "", "d"}, ""))
}

func TestSliceTrim(t *testing.T) {
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{"b", "c", "d"}))
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{" b ", "c", "    d "}))
	assert.Equal(t, []string{"a", "b"}, SliceTrim([]string{" a  ", "b", " ", "   	"}))
}

func TestSliceRemoveBlank(t *testing.T) {
	var s []string
	assert.Equal(t, []string{"a", "b"}, SliceRemoveBlank([]string{"a", "b"}))
	assert.Equal(t, []string{"b"}, SliceRemoveBlank([]string{"", "b"}))
	assert.Equal(t, []string{" a "}, SliceRemoveBlank([]string{" a ", " "}))
	assert.Equal(t, s, SliceRemoveBlank([]string{"  ", " "}))
}

func TestSliceIndex(t *testing.T) {
	assert.Equal(t, -1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "e"))
	assert.Equal(t, 1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "b"))
	assert.Equal(t, 2, SliceIndex([]string{"a", "b", "c", "d", "c"}, "c"))
	assert.Equal(t, 4, SliceLastIndex([]string{"a", "b", "c", "d", "c"}, "c"))
}

func TestIntsStrings(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, IntsToStrings([]int{1, 2, 3}))
	assert.Equal(t, []string{}, IntsToStrings([]int{}))
	assert.Equal(t, []int{}, StringsToInts([]string{}))
	assert.Equal(t, []int{23, 45}, StringsToInts([]string{"23", "45"}))
	assert.Equal(t, []int{12}, StringsToInts([]string{"a", "12"}))
}

func TestIsASCII(t *testing.T) {
	assert.Equal(t, true, IsASCII(""))
	assert.Equal(t, true, IsASCII("#"))
	assert.Equal(t, false, IsASCII("中文"))
}

func TestMaxMin(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, int64(1), MinInt64(1, 2))
	assert.Equal(t, int64(2), MaxInt64(1, 2))
}

func TestIsDir(t *testing.T) {
	t.Log(IsDir("/tmp"))
	t.Log(IsDir("/tmp/test"))
}

func TestIsExist(t *testing.T) {
	t.Log(IsExist("/tmp"))
	t.Log(IsExist("/tmp/test"))
}

func TestSimilarity(t *testing.T) {
	t.Log(Similarity("", ""))
	t.Log(Similarity("ABC", "B"))
	t.Log(Similarity("ABC", "AC"))
	t.Log(Similarity("ABC", "CA"))
	t.Log(Similarity("ABCDEF", "CD"))
	t.Log(Similarity("ABCDEFGHIG", "CEGML"))
	t.Log(Similarity("ABCDEFGHIG", "GIH"))
	t.Log(Similarity("我是中国人我说中国话", "我米说"))

	t.Log(Similarity("国家统计局上半年GDP同比增长5", "上半年GDP同比增长"))
}

func TestSimilarityText(t *testing.T) {
	t.Log(SimilarityText("国家统计局：上半年GDP同比增长5%", "上半年GDP同比增长"))
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

func TestIsUtf8(t *testing.T) {
	assert.Equal(t, true, IsUtf8(Bytes("中文")))
}

func TestRemoveLines(t *testing.T) {
	assert.Equal(t, "acb", RemoveLines("a\n\nc\nb"))
}

func BenchmarkRemoveLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveLines("https://www.163.com/news/article/HE8G1HQ8000189FH.html")
	}
}

func TestUrlParse(t *testing.T) {
	urlStrs := []string{
		"baidu.com",
		"www.baidu.com",
		"http://www.baidu.com",
		"abc",
		"javascript:;",
		"http://xf.house.163.com/{cityCode}/calculator/4.html#163-3-FDJS",
	}

	for _, urlStr := range urlStrs {
		t.Log(UrlParse(urlStr))
	}
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

func TestNormaliseSpace(t *testing.T) {
	t.Log(NormaliseSpace("中   国\n世   	界\n\n\n\n\n, hello      world     "))
	t.Log(NormaliseSpace("\n  \n\n\n\n\n    "))
}

func TestNormaliseLine(t *testing.T) {
	t.Log(NormaliseLine("中   国\n世   	界\n\n\n\n\n, hello      world     "))
}

func BenchmarkUrlParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = url.Parse("http://www.baidu.com")
	}
}

func BenchmarkString(b *testing.B) {
	resp, _ := HttpGetResp("https://www.qq.com", nil, 10000)

	body := resp.Body

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = String(body)
	}
}
