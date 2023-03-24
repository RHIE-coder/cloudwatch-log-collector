package cloudwatchLogCollector

import (
	"log"
	"time"
)

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)

const (
	Second = iota
	Millisecond
	Microsecond
	Nanosecond
)

type TimeParser struct {
	Location string
	Format   string
	Unit     int8
}

func (tp TimeParser) StringToTimestamp(timeString string) int64 {

	timestamp, err := time.Parse(tp.Format, timeString)
	if err != nil {
		log.Fatalf("time string should match a [%s] format", tp.Format)
	}

	if tp.Unit == Second {
		return timestamp.Unix()
	}

	if tp.Unit == Millisecond {
		return timestamp.UnixMilli()
	}

	if tp.Unit == Microsecond {
		return timestamp.UnixMicro()
	}

	if tp.Unit == Nanosecond {
		return timestamp.UnixNano()
	}

	return -1
}

func (tp TimeParser) TimestampToString(timestamp int64) string {
	if tp.Format == "" {
		log.Fatal("format should not be empty")
	}

	zone, err := time.LoadLocation(tp.Location)

	if err != nil {
		log.Fatalf("timezone error is occured: set value [%s]", tp.Location)
	}

	var timeInfo time.Time

	if tp.Unit == Second {
		timeInfo = time.Unix(timestamp, 0)
	}

	if tp.Unit == Millisecond {
		timeInfo = time.UnixMilli(timestamp)
	}

	if tp.Unit == Microsecond {
		timeInfo = time.UnixMicro(timestamp)
	}

	if tp.Unit == Nanosecond {
		timeInfo = time.Unix(0, timestamp)
	}

	return timeInfo.In(zone).Format(tp.Format)
}
