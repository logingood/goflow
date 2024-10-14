package clickhouse_transport_cgnat

/* cgnat clickhouse transport */

import (
	"context"
	"fmt"
	"net"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	flowprotob "github.com/cloudflare/goflow/v3/pb"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
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
	logger    *zap.Logger
}

func New(conn driver.Conn, queueSize int, dbName, tableName string) *ClickhouseClient {
	logger, err := zap.NewProduction() // NewProduction outputs JSON by default
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // Flushes buffer, if any

	return &ClickhouseClient{
		conn:      conn,
		logger:    logger,
		queue:     make(chan *flowprotob.FlowMessage, queueSize),
		dbName:    dbName,
		tableName: tableName,
	}
}

func (c *ClickhouseClient) Publish(flows []*flowprotob.FlowMessage) {
	for idx, flow := range flows {

		c.logger.Info("publishing flows",
			zap.Any("timestampRcvd", flow.TimeReceived),
			zap.Any("SamplerAddress", net.IP(flow.SamplerAddress).String()),
			zap.Any("SrcAddr", net.IP(flow.SrcAddr).String()),
			zap.Any("DstAddr", net.IP(flow.DstAddr).String()),
			zap.Any("postNATSourceIPv4Address", net.IP(flow.PostNATSourceIPv4Address).String()),
			zap.Any("postNATDestinationIPv4Address", net.IP(flow.PostNATDestinationIPv4Address).String()),
			zap.Any("SrcPort", flow.SrcPort),
			zap.Any("DstPort", flow.DstPort),
			zap.Any("postNAPTSourceTransportPort", flow.PostNAPTSourceTransportPort),
			zap.Any("postNAPTDestinationTransportPort", flow.PostNAPTDestinationTransportPort),
			zap.Any("SrcAddr", net.IP(flow.SrcAddr).String()),
			zap.Any("idex", idx),
		)

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
		log.Debug(flow)

		// CGNAT records do not have destiantion address or nextHop. IPv6 type conversion for clickhouse
		// doesn't check the length of the []byte slice and tries to go for at least 4 bytes.
		srcAddr := flow.GetSrcAddr()
		if len(srcAddr) == 0 {
			srcAddr = []byte{0, 0, 0, 0}
		}
		dstAddr := flow.GetDstAddr()
		if len(dstAddr) == 0 {
			dstAddr = []byte{0, 0, 0, 0}
		}
		nextHop := flow.GetNextHop()
		if len(nextHop) == 0 {
			nextHop = []byte{0, 0, 0, 0}
		}

		router := flow.GetSamplerAddress()
		if len(router) == 0 {
			router = append(router, 0)
		}

		postNatSrc := flow.GetPostNATSourceIPv4Address()
		if len(postNatSrc) == 0 {
			postNatSrc = []byte{0, 0, 0, 0}
		}

		postNatDst := flow.GetPostNATDestinationIPv4Address()
		if len(postNatDst) == 0 {
			postNatDst = []byte{0, 0, 0, 0}
		}

		natEvent := string(flow.GetNatEvent())
		natRealm := string(flow.GetNatOriginatingAddressRealm())

		batch.Append(
			flow.GetTimeReceived(),
			flow.GetSequenceNum(),
			flow.GetFlowDirection(),
			router,
			srcAddr,
			dstAddr,
			flow.GetProto(),
			flow.GetSrcPort(),
			flow.GetDstPort(),
			flow.GetTCPFlags(),
			flow.GetIcmpType(),
			flow.GetIcmpCode(),
			postNatSrc,
			postNatDst,
			flow.GetPostNAPTSourceTransportPort(),
			flow.GetPostNAPTDestinationTransportPort(),
			natEvent,
			natRealm,
		)
	}
	if err := batch.Send(); err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("sent successfully %d", len(flows)))
	return nil
}

// To keep the schema universal between ipv6 and ipv4, we suggest to use Ipv6 for everything.
// you can extract ipv4 address using `replaceOne(IPv6NumToString(ip), '::ffff:', ”)`
// a standard function might appear some time in future.
// https://github.com/ClickHouse/ClickHouse/issues/20469

func (c *ClickhouseClient) InitDb(ctx context.Context) error {
	stm := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.%s (
		time UInt64,
		SequenceNum UInt32,
		FlowDirection UInt32,
		Router IPv6,
		SrcAddr IPv6,
		DstAddr IPv6,
		Proto UInt32,
		SrcPort UInt32,
		DstPort UInt32,
		TCPFlags UInt32,
		IcmpType UInt32,
		IcmpCode UInt32,
		postNATSourceIPv4Address IPv6,
		postNATDestinationIPv4Address IPv6,
		postNAPTSourceTransportPort UInt32,
		postNAPTDestinationTransportPort UInt32,
		natEvent VARCHAR(255),
		natOriginatingAddressRealm VARCHAR(255)
	)
	ENGINE = MergeTree
	ORDER BY tuple()`,
		c.dbName, c.tableName)
	return c.conn.Exec(ctx, stm)
}
