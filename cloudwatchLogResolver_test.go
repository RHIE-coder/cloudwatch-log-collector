package cloudwatchLogCollector_test

import (
	clc "cloudwatch-log-collector"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	ACCESS_KEY := os.Getenv("AWS_ACCESS_KEY")
	SECRET_KEY := os.Getenv("AWS_SECRET_KEY")
	REGION_NAME := os.Getenv("REGION_NAME")

	clc.NewClient3(ACCESS_KEY, SECRET_KEY, REGION_NAME)

}
