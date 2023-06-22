package kinesis_transport

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

type API interface {
	PutRecords(ctx context.Context, params *kinesis.PutRecordsInput, optFns ...func(*kinesis.Options)) (*kinesis.PutRecordsOutput, error)
}
