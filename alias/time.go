package alias

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/x-funs/go-fun"
)

type Date struct {
	time.Time
}

type DateTime struct {
	time.Time
}

type DateTimeLayout struct {
	Time   time.Time
	Layout string
}

func parseTime(data []byte) time.Time {
	dataStr := strings.Trim(string(data), `"`)
	timestamp := fun.StrToTime(dataStr)
	if timestamp == 0 {
		return time.Time{}
	} else {
		return time.Unix(timestamp, 0)
	}
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Format(fun.DatePattern))
	return []byte(str), nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	d.Time = parseTime(data)
	return nil
}

func (d *Date) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

func (d Date) String() string {
	return d.Format(fun.DatePattern)
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Format(fun.DatetimePattern))
	return []byte(str), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	d.Time = parseTime(data)
	return nil
}

func (d *DateTime) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d DateTime) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

func (d DateTime) String() string {
	return d.Format(fun.DatetimePattern)
}

func (d DateTimeLayout) MarshalJSON() ([]byte, error) {
	if fun.Blank(d.Layout) {
		d.Layout = time.RFC3339
	}
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Time.Format(d.Layout))
	return []byte(str), nil
}

func (d *DateTimeLayout) UnmarshalJSON(data []byte) (err error) {
	d.Time = parseTime(data)
	return nil
}

func (d *DateTimeLayout) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d DateTimeLayout) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

func (d DateTimeLayout) String() string {
	return d.Time.Format(d.Layout)
}
