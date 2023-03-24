package cloudwatchLogCollector

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type CloudWatchLogOutputFormat struct {
	EventId       string
	IngestionTime int64
	LogStreamName string
	Message       string
	Timestamp     int64
}

type CloudWatch struct {
	client *cloudwatchlogs.Client
}

func NewClient2(accessKey string, secretKey string, regionName string) *cloudwatchlogs.Client {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(regionName),
	)

	if err != nil {
		log.Fatal("Error get client credential")
	}

	return cloudwatchlogs.NewFromConfig(cfg)
}

// func NewClient3(accessKey string, secretKey string, regionName string) aws.Config {
// 	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

// 	cfg, _ := config.LoadDefaultConfig(context.TODO(),
// 		config.WithCredentialsProvider(creds),
// 		config.WithRegion(regionName),
// 	)

// 	fmt.Println(*cfg)
// 	fmt.Println(reflect.TypeOf(cfg))

// 	return cfg
// }

func (cw *CloudWatch) NewClient(accessKey string, secretKey string, regionName string) {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(regionName),
	)

	if err != nil {
		log.Fatal("Error get client credential")
	}

	cw.client = cloudwatchlogs.NewFromConfig(cfg)
	fmt.Println(cw.client)
	fmt.Println(reflect.TypeOf(*cw.client))
}

// func (cw *CloudWatch) GetAllLogs(logGroupName string, filterPattern string, action func(cwof CloudWatchLogOutputFormat) string) []string {

// 	// var start int64
// 	// var end int64

// 	// start = 1678320000000
// 	// end = 1678406399999

// 	var nextToken *string
// 	messages := []string{}
// 	for {
// 		output, err := cw.client.FilterLogEvents(context.TODO(), &cloudwatchlogs.FilterLogEventsInput{
// 			LogGroupName:  &logGroupName,
// 			FilterPattern: &filterPattern,
// 			NextToken:     nextToken,
// 			// StartTime:     &start,
// 			// EndTime:       &end,
// 		})

// 		if err != nil {
// 			log.Fatal("failed to filter log events, ", err)
// 		}

// 		for _, event := range output.Events {
// 			resultMsg := action(CloudWatchLogOutputFormat{
// 				EventId:       *event.EventId,
// 				IngestionTime: *event.IngestionTime,
// 				LogStreamName: *event.LogStreamName,
// 				Message:       *event.Message,
// 				Timestamp:     *event.Timestamp,
// 			})
// 			messages = append(messages, resultMsg)
// 		}

// 		if output.NextToken == nil {
// 			break
// 		}

// 		nextToken = output.NextToken
// 		log.Println(nextToken) //test
// 	}

// 	return messages
// }
