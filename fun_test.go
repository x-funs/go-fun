package fun

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ExampleMd5() {
	fmt.Println(Md5(""))
	fmt.Println(Md5("123456"))
	// Output:
	// d41d8cd98f00b204e9800998ecf8427e
	// e10adc3949ba59abbe56e057f20f883e
}

func TestMemory(t *testing.T) {
	fmt.Println(MemoryBytes())
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

	// 无效的参数，返回空
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
}

func TestToInt64(t *testing.T) {
	assert.Equal(t, int64(0), ToInt64(""))
	assert.Equal(t, int64(0), ToInt64(" "))
	assert.Equal(t, int64(0), ToInt64(" 123 "))
	assert.Equal(t, int64(123), ToInt64("123"))
	assert.Equal(t, int64(123), ToInt64("0123"))
	assert.Equal(t, int64(0), ToInt64("1.1"))
	assert.Equal(t, int64(0), ToLong("1.1"))
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
	assert.Equal(t, StrToTime(date), StrToTime("2015年04月06日 16时03分03秒"))
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
	assert.Equal(t, "_abc", UnderToCamel("_abc"))
	assert.Equal(t, "AAbc", UnderToCamel("a_abc"))
	assert.Equal(t, "AAbc", UnderToCamel("a__abc"))
	assert.Equal(t, "TestAbc", UnderToCamel("test_abc_"))
	assert.Equal(t, "TestAbc", UnderToCamel("test_abc"))
	assert.Equal(t, "TestAbc", UnderToCamel("Test_Abc"))
	assert.Equal(t, "TestAbcDe", UnderToCamel("test_aBC_DE"))

	assert.Equal(t, "test_abc_de", CamelToUnder("TestAbcDe"))
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
}

func TestInSlice(t *testing.T) {
	assert.Equal(t, true, InSlice("a", []string{"a", "b", "c"}))
	assert.Equal(t, true, InSlice(7, []int{3, 5, 7}))
	assert.Equal(t, true, InSlice(0.2, []float64{0.1, 0.2, 0.3}))
	assert.Equal(t, false, InSlice(0.4, []float64{0.1, 0.2, 0.3}))
	assert.Equal(t, false, InSlice(2, []int{3, 5, 7}))
	assert.Equal(t, false, InSlice(2, []int{}))
	assert.Equal(t, false, InSlice(0.4, []float64{0.1, 0.2, 0.3}))
}

func TestUniqueSlice(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, UniqueSlice([]string{"a", "b", "c", "a", "b", "c"}))
	assert.Equal(t, []string{""}, UniqueSlice([]string{"", "", ""}))
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
