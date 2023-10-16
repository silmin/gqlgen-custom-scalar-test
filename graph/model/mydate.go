package model

import (
	"fmt"
	"io"
	"time"
)

type MyDate struct {
	Year  int
	Month int
	Day   int
}

// MarshalGQL implements the graphql.Marshaler interface
func (d MyDate) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%d-%d-%d"`, d.Year, d.Month, d.Day))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *MyDate) UnmarshalGQL(v interface{}) error {
	layout := "2006-01-02"
	switch v := v.(type) {
	case string:
		if result, err := time.Parse(layout, v); err != nil {
			return err
		} else {
			u = &MyDate{
				Year:  result.Year(),
				Month: int(result.Month()),
				Day:   result.Day(),
			}
		}
		return nil
	case []byte:
		if result, err := time.Parse(layout, string(v)); err != nil {
			return err
		} else {
			u = &MyDate{
				Year:  result.Year(),
				Month: int(result.Month()),
				Day:   result.Day(),
			}
		}
		return nil
	default:
		return fmt.Errorf("%T is not a date format string (YYYY-MM-DD)", v)
	}
}
