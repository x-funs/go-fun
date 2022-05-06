package strtotime

import (
	"fmt"
	"math"
	"time"
)

//result holds all the integers tha make up the final Time object returned.
// we use pointers for some properties because we need to verify if they'be been
// initialized or not
type result struct {
	// date
	y *int
	m *int
	d *int
	// time
	h *int
	i *int
	s *int
	f *int

	// relative shifts
	ry int
	rm int
	rd int
	rh int
	ri int
	rs int
	rf int

	// weekday related shifts
	weekday         *int
	weekdayBehavior int

	// first or last day of month
	// 0 none, 1 first, -1 last
	firstOrLastDayOfMonth int

	// timezone correction in minutes
	z *int

	// counters
	dates int
	times int
	zones int
}

func (r *result) ymd(y, m, d int) error {
	if r.dates > 0 {
		return fmt.Errorf("strtotime: The string contains two conflicting date/months")
	}

	r.dates++
	r.y = pointer(y)
	r.m = pointer(m)
	r.d = pointer(d)
	return nil
}

func (r *result) time(h, i, s, f int) error {
	if r.times > 0 {
		return fmt.Errorf("strtotime: The string contains two conflicting hours")
	}

	r.times++
	r.h = &h
	r.i = &i
	r.s = &s
	r.f = &f

	return nil
}

func (r *result) resetTime() error {
	r.h = pointer(0)
	r.i = pointer(0)
	r.s = pointer(0)
	r.f = pointer(0)
	r.times = 0

	return nil
}

func (r *result) zone(minutes int) error {
	if r.zones > 0 {
		return fmt.Errorf("strtotime: The string contains two conflicting time zones")

	}
	r.zones++
	r.z = pointer(minutes)
	return nil
}

func (r *result) toDate(re int64) time.Time {

	relativeTo := time.Unix(re, 0).Local()

	if r.dates > 0 && r.times <= 0 {
		r.h = pointer(0)
		r.i = pointer(0)
		r.s = pointer(0)
		r.f = pointer(0)
	}

	// fill holes
	if r.y == nil {
		y := relativeTo.Year()
		r.y = &y
	}

	if r.m == nil {
		m := lookupMonth(relativeTo.Month().String())
		r.m = &m
	}

	if r.d == nil {
		d := relativeTo.Day()
		r.d = &d
	}

	if r.h == nil {
		h := relativeTo.Hour()
		r.h = &h
	}

	if r.i == nil {
		i := relativeTo.Minute()
		r.i = &i
	}

	if r.s == nil {
		s := relativeTo.Second()
		r.s = &s
	}

	if r.f == nil {
		f := relativeTo.Nanosecond() / 1000000
		r.f = &f
	}

	// adjust special early
	switch r.firstOrLastDayOfMonth {
	case 1:
		*r.d = 1
		break
	case -1:
		*r.d = 0
		break
	}

	if r.weekday != nil {

		var dow = lookupWeekday(relativeTo.Weekday().String(), 1)

		if r.weekdayBehavior == 2 {
			// To make "r week" work, where the current day of week is a "sunday"
			if dow == 0 && *r.weekday != 0 {
				*r.weekday = -6
			}

			// To make "sunday r week" work, where the current day of week is not a "sunday"
			if *r.weekday == 0 && dow != 0 {
				*r.weekday = 7
			}

			*r.d -= dow
			*r.d += *r.weekday
		} else {
			var diff = *r.weekday - dow

			//TODO: Fix this madness
			if (r.rd < 0 && diff < 0) || (r.rd >= 0 && diff <= -r.weekdayBehavior) {
				diff += 7
			}

			if *r.weekday >= 0 {
				*r.d += diff
			} else {
				//TODO: Fix this madness
				*r.d -= int(7 - (math.Abs(float64(*r.weekday)) - float64(dow)))
			}

			r.weekday = nil
		}
	}

	// adjust relative
	*r.y += r.ry
	*r.m += r.rm
	*r.d += r.rd

	*r.h += r.rh
	*r.i += r.ri
	*r.s += r.rs
	*r.f += r.rf

	r.ry = 0
	r.rm = 0
	r.rd = 0
	r.rh = 0
	r.ri = 0
	r.rs = 0
	r.rf = 0

	// note: this is done twice in PHP
	// early when processing special relatives
	// and late
	// todo: check if the logic can be reduced
	// to just one time action
	switch r.firstOrLastDayOfMonth {
	case 1:
		*r.d = 1
		break
	case -1:
		m := lookupNumberToMonth(*r.m)
		firstOfMonth := time.Date(*r.y, m, 1, 0, 0, 0, 0, time.Local)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		_, _, *r.d = lastOfMonth.Date()
		break
	}

	// TODO: process and adjust timezone
	if r.z != nil {
		*r.i += *r.z
	}

	return time.Date(*r.y, lookupNumberToMonth(*r.m), *r.d, *r.h, *r.i, *r.s, *r.f, time.Local)
}
