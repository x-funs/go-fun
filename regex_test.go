package fun

import "testing"

func TestRegexDateTime(t *testing.T) {
	datatimes := []string{
		"2001-9-8T1:2:5+12:00",
		"2022-12-10T14:33:05+01:00",
		"2001-09-10T14:22:05+12:00",
		"2018-07-01T16:41:27.123+07:00",
		"2010-02-31T14:14:14.205+05:30",
		"2022-01-24T16:17:49-0500",
		"2022-01-24T16:17:49.123-0500",

		"2014-09-10T14:12:05Z",
		"2014-09-10T14:12:05z",
		"2019-01-01T12:34:56.789Z",
		"2019-01-01T12:34:56.789z",
		"2019-10-16T22:30",

		"2018-01-17 18:00:49.123",
		"2014-09-25 10:25:10",
		"2014-09-25 10:25",
		"2014-09-10",

		"2014/09/25 10:25:10",
		"2014/09/2510:25",

		"2015.03.10 10:00:11",
		"2015.03.10 23:00",

		"2014年09月25日 10时25分",
		"2014年09月25日  10时25分",
		"2014年09月25日 10时25分10秒",
		"2014年09月25日10时25分10秒",
		"2014年09月25日 9时5分5秒",
		"2014年09月25日 9时5分",

		"2014-6-8",
		"2014-6-8 5:6:2",
		"2014-6-8",
		"2014-6-8 5:6",
	}

	for _, datatime := range datatimes {
		if RegexDateTimePattern.MatchString(datatime) {
			t.Log(datatime)
		} else {
			t.Error(datatime)
		}
	}
}
