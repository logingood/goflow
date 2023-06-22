package kinesis_transport

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	flowmessage "github.com/cloudflare/goflow/v3/pb"
	log "github.com/sirupsen/logrus"
)

type KinesisClient struct {
	kinesisAPI API
	streamName string
}

func NewKinesisClient(kAPI API, streamName string) *KinesisClient {
	return &KinesisClient{
		kinesisAPI: kAPI,
		streamName: streamName,
	}
}

func (k *KinesisClient) Publish(messages []*flowmessage.FlowMessage) {

	records := []types.PutRecordsRequestEntry{}
	for _, msg := range messages {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Error(err)
			continue
		}

		key, _ := os.Hostname()
		if key == "" {
			key = fmt.Sprintf("%d", time.Now().Unix())
		}
		records = append(records, types.PutRecordsRequestEntry{
			Data:         data,
			PartitionKey: &key,
		})

	}

	for i := 0; i < min(500, len(records)); i += 499 {
		// TODO handle errors correctly
		if _, err := k.kinesisAPI.PutRecords(context.TODO(), &kinesis.PutRecordsInput{
			Records:    records,
			StreamName: aws.String(k.streamName),
		}); err != nil {
			log.Error(err)
		}
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
