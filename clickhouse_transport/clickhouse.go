package clickhouse_transport

import (
	"context"
	"fmt"
	"math/big"

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

		batch.Append(
			flow.GetTimeReceived(),
			flow.GetSequenceNum(),
			flow.GetSamplingRate(),
			flow.GetFlowDirection(),
			ip2Int(flow.GetSamplerAddress()),
			flow.GetTimeFlowStart(),
			flow.GetTimeFlowEnd(),
			flow.GetBytes(),
			flow.GetPackets(),
			ip2Int(flow.GetSrcAddr()),
			ip2Int(flow.GetDstAddr()),
			flow.GetEtype(),
			flow.GetProto(),
			flow.GetSrcPort(),
			flow.GetDstPort(),
			flow.GetInIf(),
			flow.GetOutIf(),
			flow.GetSrcMac(),
			flow.GetDstMac(),
			flow.GetSrcVlan(),
			flow.GetDstVlan(),
			flow.GetVlanId(),
			flow.GetIngressVrfID(),
			flow.GetEgressVrfID(),
			flow.GetIPTos(),
			flow.GetForwardingStatus(),
			flow.GetIPTTL(),
			flow.GetTCPFlags(),
			flow.GetIcmpType(),
			flow.GetIcmpCode(),
			flow.GetIPv6FlowLabel(),
			flow.GetFragmentId(),
			flow.GetFragmentOffset(),
			flow.GetBiFlowDirection(),
			flow.GetSrcAS(),
			flow.GetDstAS(),
			ip2Int(flow.GetNextHop()),
			flow.GetNextHopAS(),
			flow.GetSrcNet(),
			flow.GetDstNet(),
			flow.GetHasEncap(),
			ip2Int(flow.GetSrcAddrEncap()),
			ip2Int(flow.GetDstAddrEncap()),
			flow.GetProtoEncap(),
			flow.GetEtypeEncap(),
			flow.GetIPTosEncap(),
			flow.GetIPTTLEncap(),
			flow.GetIPv6FlowLabelEncap(),
			flow.GetFragmentIdEncap(),
			flow.GetFragmentOffsetEncap(),
			flow.GetHasMPLS(),
			flow.GetMPLSCount(),
			flow.GetMPLS1TTL(),
			flow.GetMPLS1Label(),
			flow.GetMPLS2TTL(),
			flow.GetMPLS2Label(),
			flow.GetMPLS3TTL(),
			flow.GetMPLS3Label(),
			flow.GetMPLSLastTTL(),
			flow.GetMPLSLastLabel(),
			flow.GetHasPPP(),
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
			SequenceNum UInt32,
			SamplingRate UInt64,
			FlowDirection UInt32,
			Router UInt128,
			TimeFlowStart UInt64,
			TimeFlowEnd UInt64,
			Bytes UInt64,
			Packets UInt64,
			SrcAddr UInt128,
			DstAddr UInt128,
			Etype UInt32,
			Proto UInt32,
			SrcPort UInt32,
			DstPort UInt32,
			InIf UInt32,
			OutIf UInt32,
			SrcMac UInt64,
			DstMac UInt64,
			SrcVlan UInt32,
			DstVlan UInt32,
			VlanId UInt32,
			IngressVrfID UInt32,
			EgressVrfID UInt32,
			IPTos UInt32,
			ForwardingStatus UInt32,
			IPTTL UInt32,
			TCPFlags UInt32,
			IcmpType UInt32,
			IcmpCode UInt32,
			IPv6FlowLabel UInt32,
			FragmentId UInt32,
			FragmentOffset UInt32,
			BiFlowDirection UInt32,
			SrcAS UInt32,
			DstAS UInt32,
			NextHop UInt128,
			NextHopAS UInt32,
			SrcNet UInt32,
			DstNet UInt32,
			HasEncap Bool,
			SrcAddrEncap UInt128,
			DstAddrEncap UInt128,
			ProtoEncap UInt32,
			EtypeEncap UInt32,
			IPTosEncap UInt32,
			IPTTLEncap UInt32,
			IPv6FlowLabelEncap UInt32,
			FragmentIdEncap UInt32,
			FragmentOffsetEncap UInt32,
			HasMPLS Bool,
			MPLSCount UInt32,
			MPLS1TTL UInt32,
			MPLS1Label UInt32,
			MPLS2TTL UInt32,
			MPLS2Label UInt32,
			MPLS3TTL UInt32,
			MPLS3Label UInt32,
			MPLSLastTTL UInt32,
			MPLSLastLabel UInt32,
			HasPPP Bool
	)
	ENGINE = MergeTree
	ORDER BY tuple()`,
		c.dbName, c.tableName)
	return c.conn.Exec(ctx, stm)
}

func ip2Int(b []byte) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(b)
	return i
}
