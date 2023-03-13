package cloudwatchLogCollector

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tp := &TimeParser{
		location: "Asia/Seoul",
	}
	result := tp.stringToTimestamp("hello")
	fmt.Println(result)
}
