package strtotime

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	reSpace    = "[ ]+"
	reSpaceOpt = "[ ]*"
	reMeridian = "(am|pm)"
	reHour24   = "(2[0-4]|[01]?[0-9])"
	reHour24lz = "([01][0-9]|2[0-4])"
	reHour12   = "(0?[1-9]|1[0-2])"
	reMinute   = "([0-5]?[0-9])"
	reMinutelz = "([0-5][0-9])"
	reSecond   = "(60|[0-5]?[0-9])"
	reSecondlz = "(60|[0-5][0-9])"
	reFrac     = "(?:\\.([0-9]+))"

	reDayfull = "sunday|monday|tuesday|wednesday|thursday|friday|saturday"
	reDayabbr = "sun|mon|tue|wed|thu|fri|sat"
	reDaytext = reDayfull + "|" + reDayabbr + "|weekdays?"

	reReltextnumber = "first|second|third|fourth|fifth|sixth|seventh|eighth?|ninth|tenth|eleventh|twelfth"
	reReltexttext   = "next|last|previous|this"
	reReltextunit   = "(?:second|sec|minute|min|hour|day|fortnight|forthnight|month|year)s?|weeks|" + reDaytext
	reRelmvttext    = "(back|front)"

	reYear          = "([0-9]{1,4})"
	reYear2         = "([0-9]{2})"
	reYear4         = "([0-9]{4})"
	reYear4withSign = "([+-]?[0-9]{4})"
	reMonth         = "(1[0-2]|0?[0-9])"
	reMonthlz       = "(0[0-9]|1[0-2])"
	reDay           = "(?:(3[01]|[0-2]?[0-9])(?:st|nd|rd|th)?)"
	reDaylz         = "(0[0-9]|[1-2][0-9]|3[01])"

	reMonthFull  = "january|february|march|april|may|june|july|august|september|october|november|december"
	reMonthAbbr  = "jan|feb|mar|apr|may|jun|jul|aug|sept?|oct|nov|dec"
	reMonthRoman = "i[vx]|vi{0,3}|xi{0,2}|i{1,3}"
	reMonthText  = "(" + reMonthFull + "|" + reMonthAbbr + "|" + reMonthRoman + ")"

	reTzCorrection = "((?:GMT)?([+-])" + reHour24 + ":?" + reMinute + "?)"
	reDayOfYear    = "(00[1-9]|0[1-9][0-9]|[12][0-9][0-9]|3[0-5][0-9]|36[0-6])"
	reWeekOfYear   = "(0[1-9]|[1-4][0-9]|5[0-3])"
)

type format struct {
	regex    string
	name     string
	callback func(r *result, inputs ...string) error
}

func pointer(x int) *int {
	return &x
}

func formats() []format {

	yesterday := format{
		regex: `(yesterday)`,
		name:  "yesterday",
		callback: func(r *result, inputs ...string) error {
			r.rd--
			// HACK: Original code had call to r.resetTime()
			// Might have to do with timezone adjustment
			return nil
		},
	}

	now := format{
		regex: `(now)`,
		name:  "now",
		callback: func(r *result, inputs ...string) error {
			return nil
		},
	}

	noon := format{
		regex: `(noon)`,
		name:  "noon",
		callback: func(r *result, inputs ...string) error {
			err := r.resetTime()
			if err != nil {
				return err
			}
			return r.time(12, 0, 0, 0)
		},
	}

	midnightOrToday := format{
		regex: `(midnight|today)`,
		name:  "midnight | today",
		callback: func(r *result, inputs ...string) error {
			return r.resetTime()
		},
	}

	tomorrow := format{
		regex: "(tomorrow)",
		name:  "tomorrow",
		callback: func(r *result, inputs ...string) error {
			r.rd++
			// Original code calls r.resetTime() here.
			return nil
		},
	}

	timestamp := format{
		regex: `^@(-?\d+)`,
		name:  "timestamp",
		callback: func(r *result, inputs ...string) error {
			s, err := strconv.Atoi(inputs[0])

			if err != nil {
				return err
			}

			r.rs += s
			r.y = pointer(1970)
			r.m = pointer(0)
			r.d = pointer(1)
			r.dates = 0

			return r.resetTime()
			// original code called r.zone(0)
		},
	}

	firstOrLastDay := format{
		regex: `^(first|last) day of`,
		name:  "firstdayof | lastdayof",
		callback: func(r *result, inputs ...string) error {
			if strings.ToLower(inputs[0]) == "first" {
				r.firstOrLastDayOfMonth = 1
				return nil
			}
			r.firstOrLastDayOfMonth = -1
			return nil
		},
	}

	// weekdayOf := format{
	// 	regex: "(?i)^(" + reReltextnumber + "|" + reReltexttext + ") (" + reDayfull + "|" + reDayabbr + ")" + " of",
	// 	name:  "weekdayof",
	// 	callback: func(r *result, inputs ...string) error {
	// 		relValue := inputs[0]
	// 		relUnit := inputs[1]
	// 		//TODO: implement handling of 'this time-unit'
	// 		amount, _ := lookupRelative(relValue)
	// 		return nil
	// 	},
	// 	//TODO:Implement
	// }

	backOrFrontOf := format{
		regex: "(?i)^(" + reMonthFull + ") " + reDaylz + " " + reYear + " " + reRelmvttext + " of " + reHour24 + reMeridian,
		name:  "backof | frontof",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[2])
			if err != nil {
				return nil
			}
			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return nil
			}
			hour, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}
			r.m = pointer(lookupMonth(inputs[0]))
			r.y = pointer(year)
			r.d = pointer(day)

			minute, diffhour := lookupRelative(inputs[3])

			return r.time(processMeridian(hour+diffhour, inputs[5]), minute, 0, 0)
		},
	}

	mssqltime := format{
		regex: "^" + reHour24 + ":" + reMinutelz + ":" + reSecondlz + "[:.]([0-9]+)" + reMeridian + "?",
		name:  "mssqltime",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			frac, err := strconv.Atoi(inputs[3][0:3])
			if err != nil {
				return err
			}

			if len(inputs) == 5 {
				meridian := inputs[4]
				hour = processMeridian(hour, meridian)
			}

			return r.time(hour, minute, second, frac)
		},
	}

	timeLong12 := format{
		regex: "^" + reHour12 + "[:.]" + reMinute + "[:.]" + reSecondlz + reSpaceOpt + reMeridian,
		name:  "timeLong12",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			meridian := inputs[3]

			return r.time(processMeridian(hour, meridian), minute, second, 0)
		},
	}

	timeShort12 := format{
		regex: "^" + reHour12 + "[:.]" + reMinutelz + reSpaceOpt + reMeridian,
		name:  "timeShort12",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			meridian := inputs[2]

			return r.time(processMeridian(hour, meridian), minute, 0, 0)
		},
	}

	timeTiny12 := format{
		regex: "^" + reHour12 + reSpaceOpt + reMeridian,
		name:  "timeTiny12",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			meridian := inputs[1]

			return r.time(processMeridian(hour, meridian), 0, 0, 0)
		},
	}

	soap := format{
		regex: "^" + reYear4 + "-" + reMonthlz + "-" + reDaylz + "T" + reHour24lz + ":" + reMinutelz + ":" + reSecondlz + reFrac + "(z|Z)?" + reTzCorrection + "?",
		name:  "soap",
		callback: func(r *result, inputs ...string) error {

			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}
			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			mili := inputs[6]

			if len(mili) > 3 {
				mili = mili[0:3]
			}

			frac, err := strconv.Atoi(mili)
			if err != nil {
				return err
			}

			tzCorrection := inputs[8]

			err = r.ymd(year, month-1, day)
			if err != nil {
				return err
			}
			err = r.time(hour, minute, second, frac)
			if err != nil {
				return err
			}
			if len(tzCorrection) > 0 {
				err := r.zone(processTzCorrection(tzCorrection, 0))
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	wddx := format{
		regex: "^" + reYear4 + "-" + reMonth + "-" + reDay + "T" + reHour24 + ":" + reMinute + ":" + reSecond + ".*",
		name:  "wddx",
		callback: func(r *result, inputs ...string) error {

			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}
			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			err = r.ymd(year, month-1, day)
			if err != nil {
				return err
			}

			err = r.time(hour, minute, second, 0)
			return err
		},
	}

	exif := format{
		regex: "(?i)" + "^" + reYear4 + ":" + reMonthlz + ":" + reDaylz + " " + reHour24lz + ":" + reMinutelz + ":" + reSecondlz,
		name:  "exif",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}
			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			err = r.ymd(year, month-1, day)
			if err != nil {
				return err
			}

			err = r.time(hour, minute, second, 0)
			return err
		},
	}

	xmlRpc := format{
		regex: "^" + reYear4 + reMonthlz + reDaylz + "T" + reHour24 + ":" + reMinutelz + ":" + reSecondlz,
		name:  "xmlrpc",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}
			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			err = r.ymd(year, month-1, day)
			if err != nil {
				return err
			}

			err = r.time(+hour, +minute, +second, 0)
			return err
		},
	}

	xmlRpcNoColon := format{
		regex: "^" + reYear4 + reMonthlz + reDaylz + "[Tt]" + reHour24 + reMinutelz + reSecondlz,
		name:  "xmlrpcnocolon",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}
			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			err = r.ymd(year, month-1, day)
			if err != nil {
				return err
			}

			err = r.time(hour, minute, second, 0)
			return err
		},
	}

	clf := format{
		regex: "(?i)^" + reDay + "/(" + reMonthAbbr + ")/" + reYear4 + ":" + reHour24lz + ":" + reMinutelz + ":" + reSecondlz + reSpace + reTzCorrection,
		name:  "clf",
		callback: func(r *result, inputs ...string) error {

			day, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month := inputs[1]

			year, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			hour, err := strconv.Atoi(inputs[3])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[4])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[5])
			if err != nil {
				return err
			}

			tzCorrection := inputs[6]

			err = r.ymd(year, lookupMonth(month), day)
			if err != nil {
				return err
			}

			err = r.time(hour, minute, second, 0)
			if err != nil {
				return err
			}

			err = r.zone(processTzCorrection(tzCorrection, 0))
			return err
		},
	}

	iso8601long := format{
		regex: "^[Tt]?" + reHour24 + "[:.]" + reMinute + "[:.]" + reSecond + reFrac,
		name:  "iso8601long",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			mili := inputs[3]

			if len(mili) > 3 {
				mili = mili[0:3]
			}

			frac, err := strconv.Atoi(mili)
			if err != nil {
				return err
			}
			return r.time(hour, minute, second, frac)
		},
	}

	dateTextual := format{
		regex: "(?i)^" + reMonthText + "[ .\\t-]*" + reDay + "[,.stndrh\\t ]+" + reYear,
		name:  "datetextual",
		callback: func(r *result, inputs ...string) error {

			month := inputs[0]

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			year := inputs[2]
			y, err := processYear(year)
			if err != nil {
				return err
			}

			err = r.ymd(y, lookupMonth(month), day)
			return err
		},
	}

	pointedDate4 := format{
		regex: "^" + reDay + "[.\\t-]" + reMonth + "[.-]" + reYear4,
		name:  "pointeddate4",
		callback: func(r *result, inputs ...string) error {
			day, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			year, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			return r.ymd(year, month-1, day)
		},
	}

	pointedDate2 := format{
		regex: "^" + reDay + "[.\\t]" + reMonth + "\\." + reYear2,
		name:  "pointeddate2",
		callback: func(r *result, inputs ...string) error {
			day, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			y, err := processYear(inputs[2])
			if err != nil {
				return err
			}

			return r.ymd(y, month-1, day)
		},
	}

	datePointed := format{
		regex: "^" + reYear4 + "\\." + reMonth + "\\." + reDay + "\\.?",
		name:  "datePointed",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			return r.ymd(year, month-1, day)
		},
	}

	timeLong24 := format{
		regex: "^t?" + reHour24 + "[:.]" + reMinute + "[:.]" + reSecond,
		name:  "timelong24",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			err = r.time(hour, minute, second, 0)
			return err
		},
	}

	dateNoColon := format{
		regex: "^" + reYear4 + reMonthlz + reDaylz,
		name:  "datenocolon",
		callback: func(r *result, inputs ...string) error {

			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			err = r.ymd(year, month-1, day)
			return err
		},
	}

	pgydotd := format{
		// also known as julian date format
		regex: "^" + reYear4 + `\.?` + reDayOfYear,
		name:  "pgydotd",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			return r.ymd(year, 0, day)
		},
	}

	timeShort24 := format{
		regex: "^t?" + reHour24 + "[:.]" + reMinute,
		name:  "timeshort24",
		callback: func(r *result, inputs ...string) error {
			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			return r.time(hour, minute, 0, 0)
		},
	}

	iso8601noColon := format{
		regex: "^t?" + reHour24lz + reMinutelz + reSecondlz,
		name:  "iso8601nocolon",
		callback: func(r *result, inputs ...string) error {
			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			return r.time(hour, minute, second, 0)
		},
	}

	dateSlash := format{
		regex: "^" + reYear4 + "/" + reMonth + "/" + reDay + "/?",
		name:  "dateslash",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			return r.ymd(year, month-1, day)
		},
	}

	american := format{
		regex: "^" + reMonth + "/" + reDay + "/" + reYear,
		name:  "american",
		callback: func(r *result, inputs ...string) error {
			month, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			year, err := processYear(inputs[2])
			if err != nil {
				return err
			}
			return r.ymd(year, month-1, +day)
		},
	}

	americanShort := format{
		regex: "^" + reMonth + "/" + reDay,
		name:  "americanshort",
		callback: func(r *result, inputs ...string) error {
			month, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			if r.dates > 0 {
				return fmt.Errorf("strtotime: The string contains two conflicting date/months")
			}

			r.dates++
			r.m = pointer(month - 1)
			r.d = pointer(day)

			return nil
		},
	}

	gnuDateShortOrIso8601date2 := format{
		// iso8601date2 is complete subset of gnudateshort
		regex: "^" + reYear + "-" + reMonth + "-" + reDay,
		name:  "gnudateshort | iso8601date2",
		callback: func(r *result, inputs ...string) error {
			year, err := processYear(inputs[0])

			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			return r.ymd(year, month-1, day)
		},
	}

	iso8601date4 := format{
		regex: "^" + reYear4withSign + "-" + reMonthlz + "-" + reDaylz,
		name:  "iso8601date4",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])

			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}
			return r.ymd(year, month-1, day)
		},
	}

	gnuNoColon := format{
		regex: "^t" + reHour24lz + reMinutelz,
		name:  "gnunocolon",
		callback: func(r *result, inputs ...string) error {
			hour, err := strconv.Atoi(inputs[0])

			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			if r.times > 0 {
				return fmt.Errorf("strtotime: The string contains two conflicting hours")
			}

			r.times++
			r.i = pointer(minute)
			r.h = pointer(hour)
			r.s = pointer(0)
			return nil
		},
	}

	gnuDateShorter := format{
		regex: "^" + reYear4 + "-" + reMonth,
		name:  "gnudateshorter",
		callback: func(r *result, inputs ...string) error {
			year, err := strconv.Atoi(inputs[0])

			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			return r.ymd(year, month-1, 1)
		},
	}

	pgTextReverse := format{
		// note: allowed years are from 32-9999
		// years below 32 should be treated as days in datefull
		regex: "(?i)^" + `(\d{3,4}|[4-9]\d|3[2-9])-(` + reMonthAbbr + ")-" + reDaylz,
		name:  "pgtextreverse",
		callback: func(r *result, inputs ...string) error {
			year, err := processYear(inputs[0])

			if err != nil {
				return err
			}

			month := lookupMonth(inputs[1])

			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			return r.ymd(year, month, day)
		},
	}

	dateFull := format{
		regex: "(?i)^" + reDay + `[ \t.-]*` + reMonthText + `[ \t.-]*` + reYear,
		name:  "datefull",
		callback: func(r *result, inputs ...string) error {

			day, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month := lookupMonth(inputs[1])

			year, err := processYear(inputs[2])

			if err != nil {
				return err
			}

			return r.ymd(year, month, day)
		},
	}

	dateNoDay := format{
		regex: "(?i)^" + reMonthText + `[ .\t-]*` + reYear4,
		name:  "datenoday",
		callback: func(r *result, inputs ...string) error {
			month := lookupMonth(inputs[0])

			year, err := processYear(inputs[1])

			if err != nil {
				return err
			}

			return r.ymd(year, month, 1)
		},
	}

	dateNoDayRev := format{
		regex: "(?i)^" + reYear4 + `[ .\t-]*` + reMonthText,
		name:  "datenodayrev",
		callback: func(r *result, inputs ...string) error {
			year, err := processYear(inputs[0])

			if err != nil {
				return err
			}

			month := lookupMonth(inputs[1])

			return r.ymd(year, month, 1)
		},
	}

	pgTextShort := format{
		regex: "(?i)^(" + reMonthAbbr + ")-" + reDaylz + "-" + reYear,
		name:  "pgtextshort",
		callback: func(r *result, inputs ...string) error {

			month := lookupMonth(inputs[0])

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			year, err := processYear(inputs[2])

			if err != nil {
				return err
			}

			return r.ymd(year, month, day)
		},
	}

	dateNoYear := format{
		regex: "(?i)^" + reMonthText + `[ .\t-]*` + reDay + `[,.stndrh\t ]*`,
		name:  "datenoyear",
		callback: func(r *result, inputs ...string) error {

			if r.dates > 0 {
				return fmt.Errorf("strtotime: The string contains two conflicting date/months")
			}

			r.dates++

			month := lookupMonth(inputs[0])

			day, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			r.m = pointer(month)
			r.d = pointer(day)
			return nil
		},
	}

	dateNoYearRev := format{
		regex: "(?i)^" + reDay + `[ .\t-]*` + reMonthText,
		name:  "datenoyearrev",
		callback: func(r *result, inputs ...string) error {

			if r.dates > 0 {
				return fmt.Errorf("strtotime: The string contains two conflicting date/months")
			}
			r.dates++

			day, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			month := lookupMonth(inputs[1])

			r.m = pointer(month)
			r.d = pointer(day)
			return nil
		},
	}

	isoWeekDay := format{
		regex: "^" + reYear4 + "-?W" + reWeekOfYear + `(?:-?([0-7]))?`,
		name:  "isoweekday",
		callback: func(r *result, inputs ...string) error {

			day := 1

			if len(inputs) > 2 {
				d, err := strconv.Atoi(inputs[2])
				if err != nil {
					return err
				}
				day = d
			}

			year, err := strconv.Atoi(inputs[0])

			if err != nil {
				return err
			}

			week, err := strconv.Atoi(inputs[1])

			if err != nil {
				return err
			}

			// reset date to January 1st of given year
			err = r.ymd(year, 0, 1)

			if err != nil {
				return err
			}

			// get weekday for Jan 1st
			weekday := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local).Weekday()

			// and use the day to figure out the offset for day 1 of week 1
			diff := int(weekday)
			if diff > 4 {
				diff = -(diff - 7)
			} else {
				diff = -diff
			}
			// rd is number of days after Jan 1st
			r.rd += diff + ((week - 1) * 7) + day
			return nil
		},
	}

	relativeText := format{
		regex: "(?i)^(" + reReltextnumber + "|" + reReltexttext + ")" + reSpace + "(" + reReltextunit + ")",
		name:  "relativetext",
		callback: func(r *result, inputs ...string) error {
			relValue := inputs[0]
			relUnit := inputs[1]
			// TODO: implement handling of 'this time-unit'
			amount, _ := lookupRelative(relValue)

			switch strings.ToLower(relUnit) {
			case "sec", "secs", "second", "seconds":
				r.rs += amount
				break
			case "min", "mins", "minute", "minutes":
				r.ri += amount
				break
			case "hour", "hours":
				r.rh += amount
				break
			case "day", "days":
				r.rd += amount
				break
			case "fortnight", "fortnights", "forthnight", "forthnights":
				r.rd += amount * 14
				break
			case "week", "weeks":
				r.rd += amount * 7
				break
			case "month", "months":
				r.rm += amount
				break
			case "year", "years":
				r.ry += amount
				break
			case "mon", "monday", "tue", "tuesday", "wed", "wednesday", "thu", "thursday", "fri", "friday", "sat", "saturday", "sun", "sunday":
				err := r.resetTime()
				if err != nil {
					return err
				}
				r.weekday = pointer(lookupWeekday(relUnit, 7))
				r.weekdayBehavior = 1
				if amount > 0 {
					r.rd += (amount - 1) * 7
				}
				if amount <= 0 {
					r.rd += amount * 7
				}
				break
			case "weekday", "weekdays":
				// TODO: Implement
				break
			}
			return nil
		},
	}

	relative := format{
		regex: "(?i)^([+-]*)[ \\t]*(\\d+)" + reSpaceOpt + "(" + reReltextunit + "|week)",
		name:  "relative",
		callback: func(r *result, inputs ...string) error {
			signs := inputs[0]

			relValue, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}
			relUnit := inputs[2]
			minuses := float64(strings.Count(signs, "-"))
			amount := relValue * int(math.Pow(float64(-1), minuses))

			switch strings.ToLower(relUnit) {
			case "sec", "secs", "second", "seconds":
				r.rs += amount
				break
			case "min", "mins", "minute", "minutes":
				r.ri += amount
				break
			case "hour", "hours":
				r.rh += amount
				break
			case "day", "days":
				r.rd += amount
				break
			case "fortnight", "fortnights", "forthnight", "forthnights":
				r.rd += amount * 14
				break
			case "week", "weeks":
				r.rd += amount * 7
				break
			case "month", "months":
				r.rm += amount
				break
			case "year", "years":
				r.ry += amount
				break
			case "mon", "monday", "tue", "tuesday", "wed", "wednesday", "thu", "thursday", "fri", "friday", "sat", "saturday", "sun", "sunday":
				err := r.resetTime()
				if err != nil {
					return err
				}
				r.weekday = pointer(lookupWeekday(relUnit, 7))
				r.weekdayBehavior = 1
				rd := amount * 7
				if amount > 0 {
					rd = (amount - 1) * 7
				}
				r.rd += rd
				break
			case "weekday", "weekdays":
				// todo
				break
			}
			return nil
		},
	}

	dayText := format{
		regex: "(?i)^(" + reDaytext + ")",
		name:  "daytext",
		callback: func(r *result, inputs ...string) error {
			err := r.resetTime()
			if err != nil {
				return err
			}
			r.weekday = pointer(lookupWeekday(inputs[0], 0))

			if r.weekdayBehavior != 2 {
				r.weekdayBehavior = 1
			}
			return nil
		},
	}

	relativeTextWeek := format{
		regex: "(?i)^(" + reReltexttext + ")" + reSpace + "week",
		name:  "relativetextweek",
		callback: func(r *result, inputs ...string) error {
			r.weekdayBehavior = 2

			switch strings.ToLower(inputs[0]) {
			case "this":
				r.rd += 0
				break
			case "next":
				r.rd += 7
				break
			case "last", "previous":
				r.rd -= 7
				break
			}

			if r.weekday == nil {
				r.weekday = pointer(1)
			}
			return nil
		},
	}

	monthFullOrMonthAbbr := format{
		regex: "(?i)^(" + reMonthFull + "|" + reMonthAbbr + ")",
		name:  "monthfull | monthabbr",
		callback: func(r *result, inputs ...string) error {
			month := inputs[0]
			if r.dates > 0 {
				return fmt.Errorf("strtotime: The string contains two conflicting date/months")
			}
			r.dates++
			r.m = pointer(lookupMonth(month))
			return nil
		},
	}

	tzCorrection := format{
		regex: "(?i)^" + reTzCorrection,
		name:  "tzcorrection",
		callback: func(r *result, inputs ...string) error {
			return r.zone(processTzCorrection(inputs[0], 0))
		},
	}

	ago := format{
		regex: "(?i)^ago",
		name:  "ago",
		callback: func(r *result, inputs ...string) error {
			r.ry = -r.ry
			r.rm = -r.rm
			r.rd = -r.rd
			r.rh = -r.rh
			r.ri = -r.ri
			r.rs = -r.rs
			r.rf = -r.rf
			return nil
		},
	}

	gnuNoColon2 := format{
		// second instance of gnunocolon, without leading 't'
		// it's down here, because it is very generic (4 digits in a row)
		// thus conflicts with many rules above
		// only year4 should come afterwards
		regex: "(?i)^" + reHour24lz + reMinutelz,
		name:  "gnunocolon",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			f := r.f

			if f == nil {
				f = pointer(0)
			}

			return r.time(hour, minute, 0, *f)
		},
	}

	year4 := format{
		regex: "^" + reYear4,
		name:  "year4",
		callback: func(r *result, inputs ...string) error {

			year, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			r.y = pointer(year)

			return nil
		},
	}

	whitespace := format{
		regex: "^[ .,\t]+",
		name:  "whitespace",
		callback: func(r *result, inputs ...string) error {
			return nil
		},
	}

	zhYMDformat := format{
		// 匹配中文年月日
		regex: "^" + reYear + "年" + reMonth + "月" + reDay + "(日)?",
		name:  "zhformat",
		callback: func(r *result, inputs ...string) error {
			year, err := processYear(inputs[0])

			if err != nil {
				return err
			}

			month, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			day, err := strconv.Atoi(inputs[2])
			if err != nil {
				return err
			}

			return r.ymd(year, month-1, day)
		},
	}

	zhHMSformat := format{
		// 匹配中文时分秒
		regex: "^" + reHour24 + "[时|点]" + reMinute + "分" + reSecond + "?(秒)?",
		name:  "zhHMS",
		callback: func(r *result, inputs ...string) error {

			hour, err := strconv.Atoi(inputs[0])
			if err != nil {
				return err
			}

			minute, err := strconv.Atoi(inputs[1])
			if err != nil {
				return err
			}

			second := 0
			if inputs[2] != "" {
				second, err = strconv.Atoi(inputs[2])
				if err != nil {
					return err
				}
			}

			err = r.time(hour, minute, second, 0)
			return err
		},
	}

	formats := []format{
		zhYMDformat,
		zhHMSformat,
		yesterday,
		now,
		noon,
		midnightOrToday,
		tomorrow,
		timestamp,
		firstOrLastDay,
		backOrFrontOf,
		// weekdayOf,
		mssqltime,
		timeLong12,
		timeShort12,
		timeTiny12,
		soap,
		wddx,
		exif,
		xmlRpc,
		xmlRpcNoColon,
		clf,
		iso8601long,
		dateTextual,
		pointedDate4,
		pointedDate2,
		datePointed,
		timeLong24,
		dateNoColon,
		pgydotd,
		timeShort24,
		iso8601noColon,
		dateSlash,
		american,
		americanShort,
		gnuDateShortOrIso8601date2,
		iso8601date4,
		gnuNoColon,
		gnuDateShorter,
		pgTextReverse,
		dateFull,
		dateNoDay,
		dateNoDayRev,
		pgTextShort,
		dateNoYear,
		dateNoYearRev,
		isoWeekDay,
		relativeText,
		relative,
		dayText,
		relativeTextWeek,
		monthFullOrMonthAbbr,
		tzCorrection,
		ago,
		gnuNoColon2,
		year4,
		whitespace,
	}

	return formats
}
