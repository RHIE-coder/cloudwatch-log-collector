package main

import (
	clc "cloudwatch-log-collector"
	"fmt"
)

func main() {
	tp := &clc.TimeParser{
		Location: "Asia/Seoul",
		Format:   clc.DateTime,
		Unit:     clc.Millisecond,
	}
	result1 := tp.StringToTimestamp("2023-03-13")
	fmt.Println(result1)
}
