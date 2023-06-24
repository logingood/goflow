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
	dbName    string
	tableName string
	conn      driver.Conn
	queue     chan *flowprotob.FlowMessage
}

func New(conn driver.Conn, queueSize int, dbName, tableName string) *ClickhouseClient {
	return &ClickhouseClient{
		conn:      conn,
		queue:     make(chan *flowprotob.FlowMessage, queueSize),
		dbName:    dbName,
		tableName: tableName,
	}
}

func (c *ClickhouseClient) Publish(flows []*flowprotob.FlowMessage) {
	for idx, flow := range flows {
		log.Debug("publish flow", idx, flow.SequenceNum)
		c.enqueue(flow)
	}
	return
}

func (c *ClickhouseClient) enqueue(flow *flowprotob.FlowMessage) {
	log.Debug("enqueue flow ", flow.SequenceNum)
	c.queue <- flow
}

func (c *ClickhouseClient) StartQueue(ctx context.Context, errGroup *errgroup.Group) {
	errGroup.Go(func() error {
		flows := []*flowprotob.FlowMessage{}
		for j := range c.queue {
			log.Debug("new job ", j.SequenceNum)
			flows = append(flows, j)
			if len(flows) == flushBatchSize {
				if err := c.insert(flows); err != nil {
					log.Error(err)
					//return err
				}
				log.Debug("flush flows ", len(flows))
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
	batch, err := c.conn.PrepareBatch(context.Background(), fmt.Sprintf("INSERT INTO %s.%s", c.dbName, c.tableName))
	if err != nil {
		return err
	}

	log.Debug(fmt.Sprintf("about to process  %d", len(flows)))
	for _, flow := range flows {
		srcAddr := flow.GetSrcAddr()
		dstAddr := flow.GetDstAddr()

		if len(srcAddr) != 4 || len(dstAddr) != 4 {
			// bad records, like CGNAT
			log.Debug("strange flow", flow)
			continue
		}

		batch.Append(
			flow.TimeReceived,
			flow.Bytes,
			flow.GetEtype(),
			flow.Packets,
			ipv4BytesToUint32(srcAddr[:4]),
			ipv4BytesToUint32(dstAddr[:4]),
			flow.GetSrcPort(),
			flow.GetDstPort(),
			flow.GetProto(),
			flow.GetSrcAS(),
			flow.GetDstAS(),
		)
	}
	if err := batch.Send(); err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("sent successfully %d", len(flows)))
	return nil
}

func (c *ClickhouseClient) InitDb(ctx context.Context) error {
	stm := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.%s (
	    time UInt64,
	    Bytes UInt64,
	    Etype UInt32,
	    Packets UInt64,
	    SrcAddr UInt32,
	    DstAddr UInt32,
	    SrcPort UInt32,
	    DstPort UInt32,
	    Proto UInt32,
	    SrcAs UInt32,
	    DstAs UInt32
	)
	ENGINE = MergeTree
	ORDER BY tuple()`,
		c.dbName, c.tableName)
	return c.conn.Exec(ctx, stm)
}

func ipv4BytesToUint32(b []byte) uint32 {
	// TODO deal with this later, also work out ipv6
	if len(b) != 4 {
		return 0
	}
	return uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
}
