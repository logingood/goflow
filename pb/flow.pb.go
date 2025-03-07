// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/flow.proto

package goflow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FlowMessage_FlowType int32

const (
	FlowMessage_FLOWUNKNOWN FlowMessage_FlowType = 0
	FlowMessage_SFLOW_5     FlowMessage_FlowType = 1
	FlowMessage_NETFLOW_V5  FlowMessage_FlowType = 2
	FlowMessage_NETFLOW_V9  FlowMessage_FlowType = 3
	FlowMessage_IPFIX       FlowMessage_FlowType = 4
)

// Enum value maps for FlowMessage_FlowType.
var (
	FlowMessage_FlowType_name = map[int32]string{
		0: "FLOWUNKNOWN",
		1: "SFLOW_5",
		2: "NETFLOW_V5",
		3: "NETFLOW_V9",
		4: "IPFIX",
	}
	FlowMessage_FlowType_value = map[string]int32{
		"FLOWUNKNOWN": 0,
		"SFLOW_5":     1,
		"NETFLOW_V5":  2,
		"NETFLOW_V9":  3,
		"IPFIX":       4,
	}
)

func (x FlowMessage_FlowType) Enum() *FlowMessage_FlowType {
	p := new(FlowMessage_FlowType)
	*p = x
	return p
}

func (x FlowMessage_FlowType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlowMessage_FlowType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_flow_proto_enumTypes[0].Descriptor()
}

func (FlowMessage_FlowType) Type() protoreflect.EnumType {
	return &file_pb_flow_proto_enumTypes[0]
}

func (x FlowMessage_FlowType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlowMessage_FlowType.Descriptor instead.
func (FlowMessage_FlowType) EnumDescriptor() ([]byte, []int) {
	return file_pb_flow_proto_rawDescGZIP(), []int{0, 0}
}

type FlowMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          FlowMessage_FlowType `protobuf:"varint,1,opt,name=Type,proto3,enum=flowprotob.FlowMessage_FlowType" json:"Type,omitempty"`
	TimeReceived  uint64               `protobuf:"varint,2,opt,name=TimeReceived,proto3" json:"TimeReceived,omitempty"`
	SequenceNum   uint32               `protobuf:"varint,4,opt,name=SequenceNum,proto3" json:"SequenceNum,omitempty"`
	SamplingRate  uint64               `protobuf:"varint,3,opt,name=SamplingRate,proto3" json:"SamplingRate,omitempty"`
	FlowDirection uint32               `protobuf:"varint,42,opt,name=FlowDirection,proto3" json:"FlowDirection,omitempty"`
	// Sampler information
	SamplerAddress []byte `protobuf:"bytes,11,opt,name=SamplerAddress,proto3" json:"SamplerAddress,omitempty"`
	// Found inside packet
	TimeFlowStart uint64 `protobuf:"varint,38,opt,name=TimeFlowStart,proto3" json:"TimeFlowStart,omitempty"`
	TimeFlowEnd   uint64 `protobuf:"varint,5,opt,name=TimeFlowEnd,proto3" json:"TimeFlowEnd,omitempty"`
	// Size of the sampled packet
	Bytes   uint64 `protobuf:"varint,9,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	Packets uint64 `protobuf:"varint,10,opt,name=Packets,proto3" json:"Packets,omitempty"`
	// Source/destination addresses
	SrcAddr []byte `protobuf:"bytes,6,opt,name=SrcAddr,proto3" json:"SrcAddr,omitempty"`
	DstAddr []byte `protobuf:"bytes,7,opt,name=DstAddr,proto3" json:"DstAddr,omitempty"`
	// Layer 3 protocol (IPv4/IPv6/ARP/MPLS...)
	Etype uint32 `protobuf:"varint,30,opt,name=Etype,proto3" json:"Etype,omitempty"`
	// Layer 4 protocol
	Proto uint32 `protobuf:"varint,20,opt,name=Proto,proto3" json:"Proto,omitempty"`
	// Ports for UDP and TCP
	SrcPort uint32 `protobuf:"varint,21,opt,name=SrcPort,proto3" json:"SrcPort,omitempty"`
	DstPort uint32 `protobuf:"varint,22,opt,name=DstPort,proto3" json:"DstPort,omitempty"`
	// Interfaces
	InIf  uint32 `protobuf:"varint,18,opt,name=InIf,proto3" json:"InIf,omitempty"`
	OutIf uint32 `protobuf:"varint,19,opt,name=OutIf,proto3" json:"OutIf,omitempty"`
	// Ethernet information
	SrcMac uint64 `protobuf:"varint,27,opt,name=SrcMac,proto3" json:"SrcMac,omitempty"`
	DstMac uint64 `protobuf:"varint,28,opt,name=DstMac,proto3" json:"DstMac,omitempty"`
	// Vlan
	SrcVlan uint32 `protobuf:"varint,33,opt,name=SrcVlan,proto3" json:"SrcVlan,omitempty"`
	DstVlan uint32 `protobuf:"varint,34,opt,name=DstVlan,proto3" json:"DstVlan,omitempty"`
	// 802.1q VLAN in sampled packet
	VlanId uint32 `protobuf:"varint,29,opt,name=VlanId,proto3" json:"VlanId,omitempty"`
	// VRF
	IngressVrfID uint32 `protobuf:"varint,39,opt,name=IngressVrfID,proto3" json:"IngressVrfID,omitempty"`
	EgressVrfID  uint32 `protobuf:"varint,40,opt,name=EgressVrfID,proto3" json:"EgressVrfID,omitempty"`
	// IP and TCP special flags
	IPTos            uint32 `protobuf:"varint,23,opt,name=IPTos,proto3" json:"IPTos,omitempty"`
	ForwardingStatus uint32 `protobuf:"varint,24,opt,name=ForwardingStatus,proto3" json:"ForwardingStatus,omitempty"`
	IPTTL            uint32 `protobuf:"varint,25,opt,name=IPTTL,proto3" json:"IPTTL,omitempty"`
	TCPFlags         uint32 `protobuf:"varint,26,opt,name=TCPFlags,proto3" json:"TCPFlags,omitempty"`
	IcmpType         uint32 `protobuf:"varint,31,opt,name=IcmpType,proto3" json:"IcmpType,omitempty"`
	IcmpCode         uint32 `protobuf:"varint,32,opt,name=IcmpCode,proto3" json:"IcmpCode,omitempty"`
	IPv6FlowLabel    uint32 `protobuf:"varint,37,opt,name=IPv6FlowLabel,proto3" json:"IPv6FlowLabel,omitempty"`
	// Fragments (IPv4/IPv6)
	FragmentId      uint32 `protobuf:"varint,35,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
	FragmentOffset  uint32 `protobuf:"varint,36,opt,name=FragmentOffset,proto3" json:"FragmentOffset,omitempty"`
	BiFlowDirection uint32 `protobuf:"varint,41,opt,name=BiFlowDirection,proto3" json:"BiFlowDirection,omitempty"`
	// Autonomous system information
	SrcAS     uint32 `protobuf:"varint,14,opt,name=SrcAS,proto3" json:"SrcAS,omitempty"`
	DstAS     uint32 `protobuf:"varint,15,opt,name=DstAS,proto3" json:"DstAS,omitempty"`
	NextHop   []byte `protobuf:"bytes,12,opt,name=NextHop,proto3" json:"NextHop,omitempty"`
	NextHopAS uint32 `protobuf:"varint,13,opt,name=NextHopAS,proto3" json:"NextHopAS,omitempty"`
	// Prefix size
	SrcNet uint32 `protobuf:"varint,16,opt,name=SrcNet,proto3" json:"SrcNet,omitempty"`
	DstNet uint32 `protobuf:"varint,17,opt,name=DstNet,proto3" json:"DstNet,omitempty"`
	// IP encapsulation information
	HasEncap            bool   `protobuf:"varint,43,opt,name=HasEncap,proto3" json:"HasEncap,omitempty"`
	SrcAddrEncap        []byte `protobuf:"bytes,44,opt,name=SrcAddrEncap,proto3" json:"SrcAddrEncap,omitempty"`
	DstAddrEncap        []byte `protobuf:"bytes,45,opt,name=DstAddrEncap,proto3" json:"DstAddrEncap,omitempty"`
	ProtoEncap          uint32 `protobuf:"varint,46,opt,name=ProtoEncap,proto3" json:"ProtoEncap,omitempty"`
	EtypeEncap          uint32 `protobuf:"varint,47,opt,name=EtypeEncap,proto3" json:"EtypeEncap,omitempty"`
	IPTosEncap          uint32 `protobuf:"varint,48,opt,name=IPTosEncap,proto3" json:"IPTosEncap,omitempty"`
	IPTTLEncap          uint32 `protobuf:"varint,49,opt,name=IPTTLEncap,proto3" json:"IPTTLEncap,omitempty"`
	IPv6FlowLabelEncap  uint32 `protobuf:"varint,50,opt,name=IPv6FlowLabelEncap,proto3" json:"IPv6FlowLabelEncap,omitempty"`
	FragmentIdEncap     uint32 `protobuf:"varint,51,opt,name=FragmentIdEncap,proto3" json:"FragmentIdEncap,omitempty"`
	FragmentOffsetEncap uint32 `protobuf:"varint,52,opt,name=FragmentOffsetEncap,proto3" json:"FragmentOffsetEncap,omitempty"`
	// MPLS information
	HasMPLS       bool   `protobuf:"varint,53,opt,name=HasMPLS,proto3" json:"HasMPLS,omitempty"`
	MPLSCount     uint32 `protobuf:"varint,54,opt,name=MPLSCount,proto3" json:"MPLSCount,omitempty"`
	MPLS1TTL      uint32 `protobuf:"varint,55,opt,name=MPLS1TTL,proto3" json:"MPLS1TTL,omitempty"`           // First TTL
	MPLS1Label    uint32 `protobuf:"varint,56,opt,name=MPLS1Label,proto3" json:"MPLS1Label,omitempty"`       // First Label
	MPLS2TTL      uint32 `protobuf:"varint,57,opt,name=MPLS2TTL,proto3" json:"MPLS2TTL,omitempty"`           // Second TTL
	MPLS2Label    uint32 `protobuf:"varint,58,opt,name=MPLS2Label,proto3" json:"MPLS2Label,omitempty"`       // Second Label
	MPLS3TTL      uint32 `protobuf:"varint,59,opt,name=MPLS3TTL,proto3" json:"MPLS3TTL,omitempty"`           // Third TTL
	MPLS3Label    uint32 `protobuf:"varint,60,opt,name=MPLS3Label,proto3" json:"MPLS3Label,omitempty"`       // Third Label
	MPLSLastTTL   uint32 `protobuf:"varint,61,opt,name=MPLSLastTTL,proto3" json:"MPLSLastTTL,omitempty"`     // Last TTL
	MPLSLastLabel uint32 `protobuf:"varint,62,opt,name=MPLSLastLabel,proto3" json:"MPLSLastLabel,omitempty"` // Last Label
	// PPP information
	HasPPP            bool   `protobuf:"varint,63,opt,name=HasPPP,proto3" json:"HasPPP,omitempty"`
	PPPAddressControl uint32 `protobuf:"varint,64,opt,name=PPPAddressControl,proto3" json:"PPPAddressControl,omitempty"`
	// cgnat
	PostNATSourceIPv4Address         []byte `protobuf:"bytes,225,opt,name=postNATSourceIPv4Address,proto3" json:"postNATSourceIPv4Address,omitempty"`
	PostNATDestinationIPv4Address    []byte `protobuf:"bytes,226,opt,name=postNATDestinationIPv4Address,proto3" json:"postNATDestinationIPv4Address,omitempty"`
	PostNAPTSourceTransportPort      uint32 `protobuf:"varint,227,opt,name=postNAPTSourceTransportPort,proto3" json:"postNAPTSourceTransportPort,omitempty"`
	PostNAPTDestinationTransportPort uint32 `protobuf:"varint,228,opt,name=postNAPTDestinationTransportPort,proto3" json:"postNAPTDestinationTransportPort,omitempty"`
	NatOriginatingAddressRealm       []byte `protobuf:"bytes,229,opt,name=natOriginatingAddressRealm,proto3" json:"natOriginatingAddressRealm,omitempty"`
	NatEvent                         []byte `protobuf:"bytes,230,opt,name=natEvent,proto3" json:"natEvent,omitempty"`
}

func (x *FlowMessage) Reset() {
	*x = FlowMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowMessage) ProtoMessage() {}

func (x *FlowMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowMessage.ProtoReflect.Descriptor instead.
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return file_pb_flow_proto_rawDescGZIP(), []int{0}
}

func (x *FlowMessage) GetType() FlowMessage_FlowType {
	if x != nil {
		return x.Type
	}
	return FlowMessage_FLOWUNKNOWN
}

func (x *FlowMessage) GetTimeReceived() uint64 {
	if x != nil {
		return x.TimeReceived
	}
	return 0
}

func (x *FlowMessage) GetSequenceNum() uint32 {
	if x != nil {
		return x.SequenceNum
	}
	return 0
}

func (x *FlowMessage) GetSamplingRate() uint64 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *FlowMessage) GetFlowDirection() uint32 {
	if x != nil {
		return x.FlowDirection
	}
	return 0
}

func (x *FlowMessage) GetSamplerAddress() []byte {
	if x != nil {
		return x.SamplerAddress
	}
	return nil
}

func (x *FlowMessage) GetTimeFlowStart() uint64 {
	if x != nil {
		return x.TimeFlowStart
	}
	return 0
}

func (x *FlowMessage) GetTimeFlowEnd() uint64 {
	if x != nil {
		return x.TimeFlowEnd
	}
	return 0
}

func (x *FlowMessage) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *FlowMessage) GetPackets() uint64 {
	if x != nil {
		return x.Packets
	}
	return 0
}

func (x *FlowMessage) GetSrcAddr() []byte {
	if x != nil {
		return x.SrcAddr
	}
	return nil
}

func (x *FlowMessage) GetDstAddr() []byte {
	if x != nil {
		return x.DstAddr
	}
	return nil
}

func (x *FlowMessage) GetEtype() uint32 {
	if x != nil {
		return x.Etype
	}
	return 0
}

func (x *FlowMessage) GetProto() uint32 {
	if x != nil {
		return x.Proto
	}
	return 0
}

func (x *FlowMessage) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *FlowMessage) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *FlowMessage) GetInIf() uint32 {
	if x != nil {
		return x.InIf
	}
	return 0
}

func (x *FlowMessage) GetOutIf() uint32 {
	if x != nil {
		return x.OutIf
	}
	return 0
}

func (x *FlowMessage) GetSrcMac() uint64 {
	if x != nil {
		return x.SrcMac
	}
	return 0
}

func (x *FlowMessage) GetDstMac() uint64 {
	if x != nil {
		return x.DstMac
	}
	return 0
}

func (x *FlowMessage) GetSrcVlan() uint32 {
	if x != nil {
		return x.SrcVlan
	}
	return 0
}

func (x *FlowMessage) GetDstVlan() uint32 {
	if x != nil {
		return x.DstVlan
	}
	return 0
}

func (x *FlowMessage) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *FlowMessage) GetIngressVrfID() uint32 {
	if x != nil {
		return x.IngressVrfID
	}
	return 0
}

func (x *FlowMessage) GetEgressVrfID() uint32 {
	if x != nil {
		return x.EgressVrfID
	}
	return 0
}

func (x *FlowMessage) GetIPTos() uint32 {
	if x != nil {
		return x.IPTos
	}
	return 0
}

func (x *FlowMessage) GetForwardingStatus() uint32 {
	if x != nil {
		return x.ForwardingStatus
	}
	return 0
}

func (x *FlowMessage) GetIPTTL() uint32 {
	if x != nil {
		return x.IPTTL
	}
	return 0
}

func (x *FlowMessage) GetTCPFlags() uint32 {
	if x != nil {
		return x.TCPFlags
	}
	return 0
}

func (x *FlowMessage) GetIcmpType() uint32 {
	if x != nil {
		return x.IcmpType
	}
	return 0
}

func (x *FlowMessage) GetIcmpCode() uint32 {
	if x != nil {
		return x.IcmpCode
	}
	return 0
}

func (x *FlowMessage) GetIPv6FlowLabel() uint32 {
	if x != nil {
		return x.IPv6FlowLabel
	}
	return 0
}

func (x *FlowMessage) GetFragmentId() uint32 {
	if x != nil {
		return x.FragmentId
	}
	return 0
}

func (x *FlowMessage) GetFragmentOffset() uint32 {
	if x != nil {
		return x.FragmentOffset
	}
	return 0
}

func (x *FlowMessage) GetBiFlowDirection() uint32 {
	if x != nil {
		return x.BiFlowDirection
	}
	return 0
}

func (x *FlowMessage) GetSrcAS() uint32 {
	if x != nil {
		return x.SrcAS
	}
	return 0
}

func (x *FlowMessage) GetDstAS() uint32 {
	if x != nil {
		return x.DstAS
	}
	return 0
}

func (x *FlowMessage) GetNextHop() []byte {
	if x != nil {
		return x.NextHop
	}
	return nil
}

func (x *FlowMessage) GetNextHopAS() uint32 {
	if x != nil {
		return x.NextHopAS
	}
	return 0
}

func (x *FlowMessage) GetSrcNet() uint32 {
	if x != nil {
		return x.SrcNet
	}
	return 0
}

func (x *FlowMessage) GetDstNet() uint32 {
	if x != nil {
		return x.DstNet
	}
	return 0
}

func (x *FlowMessage) GetHasEncap() bool {
	if x != nil {
		return x.HasEncap
	}
	return false
}

func (x *FlowMessage) GetSrcAddrEncap() []byte {
	if x != nil {
		return x.SrcAddrEncap
	}
	return nil
}

func (x *FlowMessage) GetDstAddrEncap() []byte {
	if x != nil {
		return x.DstAddrEncap
	}
	return nil
}

func (x *FlowMessage) GetProtoEncap() uint32 {
	if x != nil {
		return x.ProtoEncap
	}
	return 0
}

func (x *FlowMessage) GetEtypeEncap() uint32 {
	if x != nil {
		return x.EtypeEncap
	}
	return 0
}

func (x *FlowMessage) GetIPTosEncap() uint32 {
	if x != nil {
		return x.IPTosEncap
	}
	return 0
}

func (x *FlowMessage) GetIPTTLEncap() uint32 {
	if x != nil {
		return x.IPTTLEncap
	}
	return 0
}

func (x *FlowMessage) GetIPv6FlowLabelEncap() uint32 {
	if x != nil {
		return x.IPv6FlowLabelEncap
	}
	return 0
}

func (x *FlowMessage) GetFragmentIdEncap() uint32 {
	if x != nil {
		return x.FragmentIdEncap
	}
	return 0
}

func (x *FlowMessage) GetFragmentOffsetEncap() uint32 {
	if x != nil {
		return x.FragmentOffsetEncap
	}
	return 0
}

func (x *FlowMessage) GetHasMPLS() bool {
	if x != nil {
		return x.HasMPLS
	}
	return false
}

func (x *FlowMessage) GetMPLSCount() uint32 {
	if x != nil {
		return x.MPLSCount
	}
	return 0
}

func (x *FlowMessage) GetMPLS1TTL() uint32 {
	if x != nil {
		return x.MPLS1TTL
	}
	return 0
}

func (x *FlowMessage) GetMPLS1Label() uint32 {
	if x != nil {
		return x.MPLS1Label
	}
	return 0
}

func (x *FlowMessage) GetMPLS2TTL() uint32 {
	if x != nil {
		return x.MPLS2TTL
	}
	return 0
}

func (x *FlowMessage) GetMPLS2Label() uint32 {
	if x != nil {
		return x.MPLS2Label
	}
	return 0
}

func (x *FlowMessage) GetMPLS3TTL() uint32 {
	if x != nil {
		return x.MPLS3TTL
	}
	return 0
}

func (x *FlowMessage) GetMPLS3Label() uint32 {
	if x != nil {
		return x.MPLS3Label
	}
	return 0
}

func (x *FlowMessage) GetMPLSLastTTL() uint32 {
	if x != nil {
		return x.MPLSLastTTL
	}
	return 0
}

func (x *FlowMessage) GetMPLSLastLabel() uint32 {
	if x != nil {
		return x.MPLSLastLabel
	}
	return 0
}

func (x *FlowMessage) GetHasPPP() bool {
	if x != nil {
		return x.HasPPP
	}
	return false
}

func (x *FlowMessage) GetPPPAddressControl() uint32 {
	if x != nil {
		return x.PPPAddressControl
	}
	return 0
}

func (x *FlowMessage) GetPostNATSourceIPv4Address() []byte {
	if x != nil {
		return x.PostNATSourceIPv4Address
	}
	return nil
}

func (x *FlowMessage) GetPostNATDestinationIPv4Address() []byte {
	if x != nil {
		return x.PostNATDestinationIPv4Address
	}
	return nil
}

func (x *FlowMessage) GetPostNAPTSourceTransportPort() uint32 {
	if x != nil {
		return x.PostNAPTSourceTransportPort
	}
	return 0
}

func (x *FlowMessage) GetPostNAPTDestinationTransportPort() uint32 {
	if x != nil {
		return x.PostNAPTDestinationTransportPort
	}
	return 0
}

func (x *FlowMessage) GetNatOriginatingAddressRealm() []byte {
	if x != nil {
		return x.NatOriginatingAddressRealm
	}
	return nil
}

func (x *FlowMessage) GetNatEvent() []byte {
	if x != nil {
		return x.NatEvent
	}
	return nil
}

var File_pb_flow_proto protoreflect.FileDescriptor

var file_pb_flow_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x22, 0xfc, 0x12, 0x0a, 0x0b,
	0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x6f, 0x77,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46,
	0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x2a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x26, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x45, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x44,
	0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x44, 0x73,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x53, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44,
	0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x44, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x49, 0x66, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x49, 0x6e, 0x49, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x75, 0x74,
	0x49, 0x66, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4f, 0x75, 0x74, 0x49, 0x66, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x72, 0x63, 0x4d, 0x61, 0x63, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x53, 0x72, 0x63, 0x4d, 0x61, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x73, 0x74, 0x4d, 0x61,
	0x63, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x44, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x56, 0x6c, 0x61, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x53, 0x72, 0x63, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x73, 0x74,
	0x56, 0x6c, 0x61, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x44, 0x73, 0x74, 0x56,
	0x6c, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x1d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x56, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x72, 0x66, 0x49, 0x44, 0x18, 0x27, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x72, 0x66, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x72, 0x66, 0x49, 0x44, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x72, 0x66, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x50, 0x54, 0x6f, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x49, 0x50, 0x54, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x50, 0x54, 0x54, 0x4c, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x49, 0x50, 0x54, 0x54, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x43, 0x50,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x54, 0x43, 0x50,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x49, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x63, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x49, 0x63, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x49, 0x50, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49, 0x50, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x42,
	0x69, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x42, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x72, 0x63, 0x41, 0x53, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x72, 0x63, 0x41, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x44,
	0x73, 0x74, 0x41, 0x53, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x44, 0x73, 0x74, 0x41,
	0x53, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x53, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x72, 0x63,
	0x4e, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x72, 0x63, 0x4e, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x44, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x61, 0x73,
	0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x48, 0x61, 0x73,
	0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x53, 0x72, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x44, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x2e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x45, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x2f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x45, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x49, 0x50, 0x54, 0x6f, 0x73, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x30, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x49, 0x50, 0x54, 0x6f, 0x73, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x49, 0x50, 0x54, 0x54, 0x4c, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x31, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x49, 0x50, 0x54, 0x54, 0x4c, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x2e, 0x0a,
	0x12, 0x49, 0x50, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e,
	0x63, 0x61, 0x70, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x49, 0x50, 0x76, 0x36, 0x46,
	0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x28, 0x0a,
	0x0f, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x45, 0x6e, 0x63, 0x61, 0x70,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x30, 0x0a, 0x13, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x34,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x73,
	0x4d, 0x50, 0x4c, 0x53, 0x18, 0x35, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x73, 0x4d,
	0x50, 0x4c, 0x53, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x50, 0x4c, 0x53, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x4d, 0x50, 0x4c, 0x53, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x50, 0x4c, 0x53, 0x31, 0x54, 0x54, 0x4c, 0x18, 0x37, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x4d, 0x50, 0x4c, 0x53, 0x31, 0x54, 0x54, 0x4c, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x50, 0x4c, 0x53, 0x31, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x38, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x4d, 0x50, 0x4c, 0x53, 0x31, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x4d, 0x50, 0x4c, 0x53, 0x32, 0x54, 0x54, 0x4c, 0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x4d, 0x50, 0x4c, 0x53, 0x32, 0x54, 0x54, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x50, 0x4c,
	0x53, 0x32, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4d,
	0x50, 0x4c, 0x53, 0x32, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x50, 0x4c,
	0x53, 0x33, 0x54, 0x54, 0x4c, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4d, 0x50, 0x4c,
	0x53, 0x33, 0x54, 0x54, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x50, 0x4c, 0x53, 0x33, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4d, 0x50, 0x4c, 0x53, 0x33,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x61, 0x73,
	0x74, 0x54, 0x54, 0x4c, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x50, 0x4c, 0x53,
	0x4c, 0x61, 0x73, 0x74, 0x54, 0x54, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x50, 0x4c, 0x53, 0x4c,
	0x61, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x61, 0x73, 0x50, 0x50, 0x50, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x48,
	0x61, 0x73, 0x50, 0x50, 0x50, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x50, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x40, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x50, 0x50, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x3b, 0x0a, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x41, 0x54, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x50, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0xe1, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x41, 0x54, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x50, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x45, 0x0a, 0x1d, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x41, 0x54, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x50, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1d, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x41,
	0x54, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x50, 0x76, 0x34,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x1b, 0x70, 0x6f, 0x73, 0x74, 0x4e,
	0x41, 0x50, 0x54, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0xe3, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x70,
	0x6f, 0x73, 0x74, 0x4e, 0x41, 0x50, 0x54, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x20, 0x70, 0x6f,
	0x73, 0x74, 0x4e, 0x41, 0x50, 0x54, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0xe4,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x41, 0x50, 0x54, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x1a, 0x6e, 0x61, 0x74, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0xe5, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1a, 0x6e, 0x61,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x1b, 0x0a, 0x08, 0x6e, 0x61, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0xe6, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x61, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4c, 0x4f, 0x57, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x35, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x35, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x39, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x50, 0x46, 0x49, 0x58, 0x10, 0x04, 0x42, 0x48, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x65, 0x74,
	0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x61, 0x67, 0x67, 0x42, 0x0d, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x62, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x6f,
	0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_flow_proto_rawDescOnce sync.Once
	file_pb_flow_proto_rawDescData = file_pb_flow_proto_rawDesc
)

func file_pb_flow_proto_rawDescGZIP() []byte {
	file_pb_flow_proto_rawDescOnce.Do(func() {
		file_pb_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_flow_proto_rawDescData)
	})
	return file_pb_flow_proto_rawDescData
}

var file_pb_flow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_flow_proto_goTypes = []interface{}{
	(FlowMessage_FlowType)(0), // 0: flowprotob.FlowMessage.FlowType
	(*FlowMessage)(nil),       // 1: flowprotob.FlowMessage
}
var file_pb_flow_proto_depIdxs = []int32{
	0, // 0: flowprotob.FlowMessage.Type:type_name -> flowprotob.FlowMessage.FlowType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_flow_proto_init() }
func file_pb_flow_proto_init() {
	if File_pb_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_flow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_flow_proto_goTypes,
		DependencyIndexes: file_pb_flow_proto_depIdxs,
		EnumInfos:         file_pb_flow_proto_enumTypes,
		MessageInfos:      file_pb_flow_proto_msgTypes,
	}.Build()
	File_pb_flow_proto = out.File
	file_pb_flow_proto_rawDesc = nil
	file_pb_flow_proto_goTypes = nil
	file_pb_flow_proto_depIdxs = nil
}
