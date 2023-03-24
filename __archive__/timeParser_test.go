package cloudwatchLogCollector_test

import (
	clc "cloudwatch-log-collector"
	"testing"
)

func TestStringToTimestamp(t *testing.T) {
	tp := &clc.TimeParser{
		Location: "Asia/Seoul",
		Format:   clc.DateOnly,
		Unit:     clc.Second, // Default
	}

	var result int64

	result = tp.StringToTimestamp("2023-03-13")

	if result == -1 {
		t.Error("maybe unit doesn't match")
	}

	if result != 1678665600 {
		t.Error("cannot convert time string to second-timestamp")
	}

	tp.Unit = clc.Millisecond
	result = tp.StringToTimestamp("2023-03-13")

	if result != 1678665600000 {
		t.Error("cannot convert time string to milli-timestamp")
	}

	tp.Unit = clc.Microsecond
	result = tp.StringToTimestamp("2023-03-13")

	if result != 1678665600000000 {
		t.Error("cannot convert time string to micro-timestamp")
	}

	tp.Unit = clc.Nanosecond
	result = tp.StringToTimestamp("2023-03-13")

	if result != 1678665600000000000 {
		t.Error("cannot convert time string to nano-timestamp")
	}

}

func TestTimestampToString(t *testing.T) {

	tp := &clc.TimeParser{
		Location: "Asia/Seoul",
		Format:   clc.DateTime,
		Unit:     clc.Millisecond,
	}

	result := tp.TimestampToString(1678755751151)

	if result != "2023-03-14 10:02:31" {
		t.Error("the time string doesn't match")
	}

}

func TestConvertWorking(t *testing.T) {

	tp := &clc.TimeParser{
		Location: "Asia/Seoul",
		Format:   clc.RFC3339Nano,
		Unit:     clc.Nanosecond,
	}

	var originTimestamp int64 = 1678759158377246848 //nano

	timestampString := tp.TimestampToString(originTimestamp)

	if timestampString != "2023-03-14T10:59:18.377246848+09:00" {
		t.Error("fail to convert timestamp to string")
	}

	var outputTimestamp int64 = tp.StringToTimestamp(timestampString)

	if originTimestamp != outputTimestamp {
		t.Error("original timestamp and converted timestamp should be same")
	}

}
