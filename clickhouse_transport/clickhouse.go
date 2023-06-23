package clickhouse_transport

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	flowprotob "github.com/cloudflare/goflow/v3/pb"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	flushBatchSize = 10000
)

type ClickhouseClient struct {
	conn  driver.Conn
	queue chan *flowprotob.FlowMessage
}

func New(conn driver.Conn, queueSize int) *ClickhouseClient {
	return &ClickhouseClient{
		conn:  conn,
		queue: make(chan *flowprotob.FlowMessage, queueSize),
	}
}

func (c *ClickhouseClient) Publish(flows []*flowprotob.FlowMessage) {
	for _, flow := range flows {
		c.enqueu(flow)
	}
	return
}

func (c *ClickhouseClient) enqueu(flow *flowprotob.FlowMessage) {
	c.queue <- flow
}

func (c *ClickhouseClient) StartQueue(ctx context.Context, errGroup *errgroup.Group) {
	errGroup.Go(func() error {
		flows := []*flowprotob.FlowMessage{}
		for j := range c.queue {
			flows = append(flows, j)
			if len(flows) == flushBatchSize {
				if err := c.insert(flows); err != nil {
					return err
				}
				flows = nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				// keep working
			}
		}
		return nil
	})
}

func (c *ClickhouseClient) insert(flows []*flowprotob.FlowMessage) error {
	batch, err := c.conn.PrepareBatch(context.Background(), "INSERT INTO deafult.nflow")
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("about to process  %d", len(flows)))
	for _, flow := range flows {
		batch.Append(
			flow.TimeReceived,
			flow.Bytes,
			flow.GetEtype(),
			flow.Packets,
			flow.GetSrcAddr(),
			flow.GetDstAddr(),
			flow.GetSrcPort(),
			flow.GetDstPort(),
			flow.GetProto(),
			flow.GetType(),
			flow.GetSrcAS(),
			flow.GetDstAS(),
		)
	}
	if err := batch.Send(); err != nil {
		return err
	}
	log.Info(fmt.Sprintf("sent successfully %d", len(flows)))
	return nil
}

func (c *ClickhouseClient) InitDb(ctx context.Context, dbName, tableName string) error {
	stm := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.%s (
	    time UInt32,
	    Bytes UInt16,
	    Etype UInt32,
	    Packets UInt64,
	    SrcAddr UInt32,
	    DstAddr UInt32,
	    SrcPort UInt32,
	    DstPort UInt32,
	    Proto UInt32,
	    SrcAs UInt32,
	    DstAs UInt32
	) ENGINE = MergeTree()
	ORDER BY (time, SrcAddr, SrcPort, DstAddr, DstPort)
	PARTITION BY DstAddr
	SAMPLE BY SrcAddr`, dbName, tableName)
	log.Info(stm)
	return c.conn.Exec(ctx, stm)
}
