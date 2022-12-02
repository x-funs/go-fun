package fun

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	timeStamp := Timestamp()
	timeStampMill := Timestamp(true)

	fmt.Println(timeStamp)
	fmt.Println(timeStampMill)

	assert.NotEmpty(t, timeStamp)
	assert.NotEmpty(t, timeStampMill)
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

func TestStrTotimeEn(t *testing.T) {
	dateStrs := []string{
		"30 August 2022",
		"02 Sep 2022",
		"02 Sep 2022 11:40",
		"02 Sep 2022 11:40:53am",
		"02 Sep 2022 11:40:53am",
		"02 Sep 2022 11:40 aM",
		"Sep 02 2022 11:40:00 aM",
		"Sep 02 2022 11:40 aM",
		"Sep 02 2022 11:40",
		"20220601",
	}
	for _, dateStr := range dateStrs {
		t.Log(Date(StrToTime(dateStr)))
	}
}

func TestStrToTime(t *testing.T) {

	date := "2015-04-06 16:03:03"
	timeStamp := StrToTime(date)

	// 极限值
	assert.Equal(t, "2023-01-31 23:59:59", Date(StrToTime("+1 month", StrToTime("2022-12-31 23:59:59"))))
	assert.Equal(t, "2023-01-01 23:59:59", Date(StrToTime("+1 day", StrToTime("2022-12-31 23:59:59"))))
	assert.Equal(t, "2023-01-01 00:59:59", Date(StrToTime("+1 hour", StrToTime("2022-12-31 23:59:59"))))
	assert.Equal(t, "2023-01-01 00:00:59", Date(StrToTime("+1 minute", StrToTime("2022-12-31 23:59:59"))))
	assert.Equal(t, "2023-01-01 00:00:00", Date(StrToTime("+1 second", StrToTime("2022-12-31 23:59:59"))))

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
