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

type DateTimeFormat struct {
	Time   time.Time
	Format string
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Format(fun.DatePattern))
	return []byte(str), nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	dataStr := strings.Trim(string(data), `"`)
	timestamp := fun.StrToTime(dataStr)
	if timestamp == 0 {
		d.Time = time.Time{}
	} else {
		d.Time = time.Unix(timestamp, 0)
	}
	return nil
}

func (d *Date) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Format(fun.DatetimePattern))
	return []byte(str), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	dataStr := strings.Trim(string(data), `"`)
	timestamp := fun.StrToTime(dataStr)
	if timestamp == 0 {
		d.Time = time.Time{}
	} else {
		d.Time = time.Unix(timestamp, 0)
	}
	return nil
}

func (d *DateTime) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d DateTime) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d DateTimeFormat) MarshalJSON() ([]byte, error) {
	if fun.Blank(d.Format) {
		d.Format = time.RFC3339
	}
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	str := fmt.Sprintf(`"%s"`, d.Time.Format(d.Format))
	return []byte(str), nil
}

func (d *DateTimeFormat) UnmarshalJSON(data []byte) (err error) {
	dataStr := strings.Trim(string(data), `"`)
	timestamp := fun.StrToTime(dataStr)
	if timestamp == 0 {
		d.Time = time.Time{}
	} else {
		d.Time = time.Unix(timestamp, 0)
	}
	return nil
}

func (d *DateTimeFormat) Scan(src any) error {
	var err error
	switch x := src.(type) {
	case time.Time:
		d.Time = x
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", src, src)
	}
	return err
}

func (d DateTimeFormat) Value() (driver.Value, error) {
	return d.Time, nil
}
