// Code generated by protoc-gen-go.
// source: ipv6_dhcpv6d_proxy_binding.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_binding_clients_client is a generated protocol buffer package.

It is generated from these files:
	ipv6_dhcpv6d_proxy_binding.proto

It has these top-level messages:
	Ipv6Dhcpv6DProxyBinding_KEYS
	Ipv6Dhcpv6DProxyBinding
	IPV6AddressType
	TimeBag
	BagDhcpv6DAddrAttrb
	BagDhcpv6DAddrAttrbItem
	BagDhcpv6DIaIdPdInfo
	BagDhcpv6DIaIdPdInfoItem
*/
package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_binding_clients_client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DHCPv6 proxy binding entry
type Ipv6Dhcpv6DProxyBinding_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
}

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) Reset()                    { *m = Ipv6Dhcpv6DProxyBinding_KEYS{} }
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyBinding_KEYS) ProtoMessage()               {}
func (*Ipv6Dhcpv6DProxyBinding_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type Ipv6Dhcpv6DProxyBinding struct {
	// Client DUID
	Duid string `protobuf:"bytes,50,opt,name=duid" json:"duid,omitempty"`
	// DHCPV6 client flag
	ClientFlag uint32 `protobuf:"varint,51,opt,name=client_flag,json=clientFlag" json:"client_flag,omitempty"`
	// DHCPV6 subscriber label
	SubscriberLabel uint32 `protobuf:"varint,52,opt,name=subscriber_label,json=subscriberLabel" json:"subscriber_label,omitempty"`
	// DHCPVV6 client/subscriber VRF name
	VrfName string `protobuf:"bytes,53,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Client MAC address
	MacAddress string `protobuf:"bytes,54,opt,name=mac_address,json=macAddress" json:"mac_address,omitempty"`
	// Number of ia_id/pd
	IaIdPDs uint32 `protobuf:"varint,55,opt,name=ia_id_p_ds,json=iaIdPDs" json:"ia_id_p_ds,omitempty"`
	// List of DHCPv6 IA_ID/PDs
	IaIdPd *BagDhcpv6DIaIdPdInfo `protobuf:"bytes,56,opt,name=ia_id_pd,json=iaIdPd" json:"ia_id_pd,omitempty"`
	// DHCPV6 access interface to client
	InterfaceName string `protobuf:"bytes,57,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// DHCPV6 access VRF name to client
	AccessVrfName string `protobuf:"bytes,58,opt,name=access_vrf_name,json=accessVrfName" json:"access_vrf_name,omitempty"`
	// DHCPV6 VLAN tag count
	ProxyBindingTags uint32 `protobuf:"varint,59,opt,name=proxy_binding_tags,json=proxyBindingTags" json:"proxy_binding_tags,omitempty"`
	// DHCPV6 VLAN Outer VLAN
	ProxyBindingOuterTag uint32 `protobuf:"varint,60,opt,name=proxy_binding_outer_tag,json=proxyBindingOuterTag" json:"proxy_binding_outer_tag,omitempty"`
	// DHCPV6 VLAN Inner VLAN
	ProxyBindingInnerTag uint32 `protobuf:"varint,61,opt,name=proxy_binding_inner_tag,json=proxyBindingInnerTag" json:"proxy_binding_inner_tag,omitempty"`
	// DHCPV6 class name
	ClassName string `protobuf:"bytes,62,opt,name=class_name,json=className" json:"class_name,omitempty"`
	// DHCPV6 pool name
	PoolName string `protobuf:"bytes,63,opt,name=pool_name,json=poolName" json:"pool_name,omitempty"`
	// DHCPV6 received Remote ID
	RxRemoteId string `protobuf:"bytes,64,opt,name=rx_remote_id,json=rxRemoteId" json:"rx_remote_id,omitempty"`
	// DHCPV6 transmitted Remote ID
	TxRemoteId string `protobuf:"bytes,65,opt,name=tx_remote_id,json=txRemoteId" json:"tx_remote_id,omitempty"`
	// DHCPV6 received Interface ID
	RxInterfaceId string `protobuf:"bytes,66,opt,name=rx_interface_id,json=rxInterfaceId" json:"rx_interface_id,omitempty"`
	// DHCPV6 transmitted Interface ID
	TxInterfaceId string `protobuf:"bytes,67,opt,name=tx_interface_id,json=txInterfaceId" json:"tx_interface_id,omitempty"`
	// DHCPV6 server IPv6 address
	ServerIpv6Address string `protobuf:"bytes,68,opt,name=server_ipv6_address,json=serverIpv6Address" json:"server_ipv6_address,omitempty"`
	// DHCPV6 profile name
	ProfileName string `protobuf:"bytes,69,opt,name=profile_name,json=profileName" json:"profile_name,omitempty"`
	// DHCPV6 framed ipv6 addess used by ND
	FramedIpv6Prefix string `protobuf:"bytes,70,opt,name=framed_ipv6_prefix,json=framedIpv6Prefix" json:"framed_ipv6_prefix,omitempty"`
	// DHCPV6 framed ipv6 prefix length used by ND
	FramedPrefixLength uint32 `protobuf:"varint,71,opt,name=framed_prefix_length,json=framedPrefixLength" json:"framed_prefix_length,omitempty"`
	// Is true if DHCP next renew from client will be NAK'd
	IsNakNextRenew bool `protobuf:"varint,72,opt,name=is_nak_next_renew,json=isNakNextRenew" json:"is_nak_next_renew,omitempty"`
	// DHCPV6 SRG state
	SrgState uint32 `protobuf:"varint,73,opt,name=srg_state,json=srgState" json:"srg_state,omitempty"`
	// DHCPV6 SRG Intf Role
	SrgIntfRole uint32 `protobuf:"varint,74,opt,name=srg_intf_role,json=srgIntfRole" json:"srg_intf_role,omitempty"`
	// SRG P2P Status
	Srgp2P bool `protobuf:"varint,75,opt,name=srgp2_p,json=srgp2P" json:"srgp2_p,omitempty"`
	// DHCPV6 SRG VRF NAME
	SrgVrfName string `protobuf:"bytes,76,opt,name=srg_vrf_name,json=srgVrfName" json:"srg_vrf_name,omitempty"`
}

func (m *Ipv6Dhcpv6DProxyBinding) Reset()                    { *m = Ipv6Dhcpv6DProxyBinding{} }
func (m *Ipv6Dhcpv6DProxyBinding) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyBinding) ProtoMessage()               {}
func (*Ipv6Dhcpv6DProxyBinding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6Dhcpv6DProxyBinding) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetClientFlag() uint32 {
	if m != nil {
		return m.ClientFlag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSubscriberLabel() uint32 {
	if m != nil {
		return m.SubscriberLabel
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIaIdPDs() uint32 {
	if m != nil {
		return m.IaIdPDs
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIaIdPd() *BagDhcpv6DIaIdPdInfo {
	if m != nil {
		return m.IaIdPd
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyBinding) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetAccessVrfName() string {
	if m != nil {
		return m.AccessVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingTags() uint32 {
	if m != nil {
		return m.ProxyBindingTags
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingOuterTag() uint32 {
	if m != nil {
		return m.ProxyBindingOuterTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingInnerTag() uint32 {
	if m != nil {
		return m.ProxyBindingInnerTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetRxRemoteId() string {
	if m != nil {
		return m.RxRemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetTxRemoteId() string {
	if m != nil {
		return m.TxRemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetRxInterfaceId() string {
	if m != nil {
		return m.RxInterfaceId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetTxInterfaceId() string {
	if m != nil {
		return m.TxInterfaceId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetServerIpv6Address() string {
	if m != nil {
		return m.ServerIpv6Address
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetFramedIpv6Prefix() string {
	if m != nil {
		return m.FramedIpv6Prefix
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetFramedPrefixLength() uint32 {
	if m != nil {
		return m.FramedPrefixLength
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIsNakNextRenew() bool {
	if m != nil {
		return m.IsNakNextRenew
	}
	return false
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgState() uint32 {
	if m != nil {
		return m.SrgState
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgIntfRole() uint32 {
	if m != nil {
		return m.SrgIntfRole
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgp2P() bool {
	if m != nil {
		return m.Srgp2P
	}
	return false
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgVrfName() string {
	if m != nil {
		return m.SrgVrfName
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6AddressType) Reset()                    { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()               {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Time in diffrent formats
type TimeBag struct {
	// DHCPV6 client lease in seconds
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Time in format HH:MM:SS
	Time string `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
}

func (m *TimeBag) Reset()                    { *m = TimeBag{} }
func (m *TimeBag) String() string            { return proto.CompactTextString(m) }
func (*TimeBag) ProtoMessage()               {}
func (*TimeBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TimeBag) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeBag) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type BagDhcpv6DAddrAttrb struct {
	// Next address attribute information
	BagDhcpv6DAddrAttrb []*BagDhcpv6DAddrAttrbItem `protobuf:"bytes,1,rep,name=bag_dhcpv6d_addr_attrb,json=bagDhcpv6dAddrAttrb" json:"bag_dhcpv6d_addr_attrb,omitempty"`
}

func (m *BagDhcpv6DAddrAttrb) Reset()                    { *m = BagDhcpv6DAddrAttrb{} }
func (m *BagDhcpv6DAddrAttrb) String() string            { return proto.CompactTextString(m) }
func (*BagDhcpv6DAddrAttrb) ProtoMessage()               {}
func (*BagDhcpv6DAddrAttrb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BagDhcpv6DAddrAttrb) GetBagDhcpv6DAddrAttrb() []*BagDhcpv6DAddrAttrbItem {
	if m != nil {
		return m.BagDhcpv6DAddrAttrb
	}
	return nil
}

type BagDhcpv6DAddrAttrbItem struct {
	// IPv6 prefix
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	// Prefix length
	PrefixLength uint32 `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Lease time in seconds
	LeaseTime uint32 `protobuf:"varint,3,opt,name=lease_time,json=leaseTime" json:"lease_time,omitempty"`
	// Remaining lease time in seconds
	RemainingLeaseTime uint32 `protobuf:"varint,4,opt,name=remaining_lease_time,json=remainingLeaseTime" json:"remaining_lease_time,omitempty"`
}

func (m *BagDhcpv6DAddrAttrbItem) Reset()                    { *m = BagDhcpv6DAddrAttrbItem{} }
func (m *BagDhcpv6DAddrAttrbItem) String() string            { return proto.CompactTextString(m) }
func (*BagDhcpv6DAddrAttrbItem) ProtoMessage()               {}
func (*BagDhcpv6DAddrAttrbItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BagDhcpv6DAddrAttrbItem) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *BagDhcpv6DAddrAttrbItem) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *BagDhcpv6DAddrAttrbItem) GetLeaseTime() uint32 {
	if m != nil {
		return m.LeaseTime
	}
	return 0
}

func (m *BagDhcpv6DAddrAttrbItem) GetRemainingLeaseTime() uint32 {
	if m != nil {
		return m.RemainingLeaseTime
	}
	return 0
}

type BagDhcpv6DIaIdPdInfo struct {
	// Next ia_id_pd information
	BagDhcpv6DIaIdPdInfo []*BagDhcpv6DIaIdPdInfoItem `protobuf:"bytes,1,rep,name=bag_dhcpv6d_ia_id_pd_info,json=bagDhcpv6dIaIdPdInfo" json:"bag_dhcpv6d_ia_id_pd_info,omitempty"`
}

func (m *BagDhcpv6DIaIdPdInfo) Reset()                    { *m = BagDhcpv6DIaIdPdInfo{} }
func (m *BagDhcpv6DIaIdPdInfo) String() string            { return proto.CompactTextString(m) }
func (*BagDhcpv6DIaIdPdInfo) ProtoMessage()               {}
func (*BagDhcpv6DIaIdPdInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BagDhcpv6DIaIdPdInfo) GetBagDhcpv6DIaIdPdInfo() []*BagDhcpv6DIaIdPdInfoItem {
	if m != nil {
		return m.BagDhcpv6DIaIdPdInfo
	}
	return nil
}

type BagDhcpv6DIaIdPdInfoItem struct {
	// IA type
	IaType string `protobuf:"bytes,1,opt,name=ia_type,json=iaType" json:"ia_type,omitempty"`
	// IA_ID of this IA
	IaId uint32 `protobuf:"varint,2,opt,name=ia_id,json=iaId" json:"ia_id,omitempty"`
	// FSM Flag for this IA
	Flags uint32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	// Total address in this IA
	TotalAddress uint32 `protobuf:"varint,4,opt,name=total_address,json=totalAddress" json:"total_address,omitempty"`
	// State
	State string `protobuf:"bytes,5,opt,name=state" json:"state,omitempty"`
	// List of addresses in this IA
	Addresses *BagDhcpv6DAddrAttrb `protobuf:"bytes,6,opt,name=addresses" json:"addresses,omitempty"`
}

func (m *BagDhcpv6DIaIdPdInfoItem) Reset()                    { *m = BagDhcpv6DIaIdPdInfoItem{} }
func (m *BagDhcpv6DIaIdPdInfoItem) String() string            { return proto.CompactTextString(m) }
func (*BagDhcpv6DIaIdPdInfoItem) ProtoMessage()               {}
func (*BagDhcpv6DIaIdPdInfoItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BagDhcpv6DIaIdPdInfoItem) GetIaType() string {
	if m != nil {
		return m.IaType
	}
	return ""
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetIaId() uint32 {
	if m != nil {
		return m.IaId
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetTotalAddress() uint32 {
	if m != nil {
		return m.TotalAddress
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetAddresses() *BagDhcpv6DAddrAttrb {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DProxyBinding_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.ipv6_dhcpv6d_proxy_binding_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DProxyBinding)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.ipv6_dhcpv6d_proxy_binding")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.IPV6AddressType")
	proto.RegisterType((*TimeBag)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.time_bag")
	proto.RegisterType((*BagDhcpv6DAddrAttrb)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_addr_attrb")
	proto.RegisterType((*BagDhcpv6DAddrAttrbItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_addr_attrb_item")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfo)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_ia_id_pd_info")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfoItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_ia_id_pd_info_item")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_proxy_binding.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 967 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xb6, 0x89, 0x63, 0x1f, 0xc7, 0x4d, 0x3a, 0xb1, 0xda, 0x29, 0x51, 0x89, 0x31, 0xa2,
	0xb8, 0x12, 0xb2, 0x50, 0x4a, 0x43, 0xf9, 0x27, 0x25, 0x2d, 0x2c, 0x8d, 0x42, 0xd8, 0x46, 0x95,
	0x10, 0x17, 0xa3, 0xf1, 0xee, 0xec, 0x76, 0xd4, 0xfd, 0x63, 0x66, 0xe2, 0x3a, 0x2f, 0x00, 0x0f,
	0xc2, 0x15, 0xf0, 0x0e, 0x5c, 0xf3, 0x3e, 0xbc, 0x00, 0x9a, 0x33, 0xb3, 0xb6, 0x43, 0x95, 0x5c,
	0x41, 0x6f, 0x92, 0x9d, 0xef, 0xfb, 0xce, 0x99, 0xd9, 0x73, 0xbe, 0x39, 0x6b, 0x18, 0xc8, 0x7a,
	0xba, 0xc7, 0x92, 0xe7, 0x71, 0x3d, 0xdd, 0x4b, 0x58, 0xad, 0xaa, 0xd9, 0x19, 0x9b, 0xc8, 0x32,
	0x91, 0x65, 0x36, 0xae, 0x55, 0x65, 0x2a, 0xf2, 0x7d, 0x2c, 0x75, 0x5c, 0x31, 0x59, 0x69, 0x36,
	0x53, 0x0c, 0xe5, 0xa5, 0x78, 0x39, 0x0f, 0xa9, 0x6a, 0xa1, 0xc6, 0x6e, 0x31, 0x2e, 0xab, 0x44,
	0x68, 0xfc, 0x3b, 0xc6, 0x4c, 0xe3, 0x26, 0x53, 0x9c, 0x4b, 0x51, 0x1a, 0xed, 0xff, 0x0f, 0x7f,
	0x84, 0x9d, 0x8b, 0xb7, 0x65, 0x4f, 0x1e, 0xfd, 0xf0, 0x94, 0x6c, 0x43, 0xc7, 0xe6, 0x61, 0x25,
	0x2f, 0x04, 0x0d, 0x06, 0xc1, 0xa8, 0x13, 0xb5, 0x2d, 0x70, 0xc4, 0x0b, 0x61, 0x49, 0x97, 0x89,
	0xc9, 0x84, 0x5e, 0x71, 0xa4, 0x03, 0xc2, 0x64, 0xf8, 0x77, 0x1b, 0xde, 0xb8, 0x38, 0x3b, 0x21,
	0xb0, 0x92, 0x9c, 0xca, 0x84, 0xee, 0x62, 0x18, 0x3e, 0x93, 0x1d, 0xe8, 0xfa, 0x7c, 0x69, 0xce,
	0x33, 0x7a, 0x6f, 0x10, 0x8c, 0x7a, 0x11, 0x38, 0xe8, 0x71, 0xce, 0x33, 0x72, 0x17, 0x36, 0xf5,
	0xe9, 0x44, 0xc7, 0x4a, 0x4e, 0x84, 0x62, 0x39, 0x9f, 0x88, 0x9c, 0x7e, 0x80, 0xaa, 0x8d, 0x05,
	0x7e, 0x68, 0x61, 0x72, 0x0b, 0xda, 0x53, 0x95, 0xba, 0x73, 0xdf, 0xc7, 0x3d, 0xd6, 0xa6, 0x2a,
	0xc5, 0x63, 0xef, 0x40, 0xb7, 0xe0, 0x31, 0xe3, 0x49, 0xa2, 0x84, 0xd6, 0x74, 0x0f, 0x59, 0x28,
	0x78, 0xbc, 0xef, 0x10, 0xb2, 0x0d, 0x20, 0x39, 0x93, 0x09, 0xab, 0x59, 0xa2, 0xe9, 0x87, 0xb8,
	0xc1, 0x9a, 0xe4, 0x61, 0x72, 0x7c, 0xa0, 0xc9, 0xcf, 0x01, 0xb4, 0x3d, 0x9b, 0xd0, 0x07, 0x83,
	0x60, 0xd4, 0xdd, 0xcd, 0xc7, 0xff, 0x79, 0x6f, 0xc6, 0x13, 0x9e, 0xcd, 0xe3, 0x9b, 0xed, 0x98,
	0x2c, 0xd3, 0x2a, 0x6a, 0xe1, 0x49, 0x12, 0xf2, 0x0e, 0x5c, 0x93, 0xa5, 0x11, 0x2a, 0xe5, 0xb1,
	0xef, 0xcf, 0x47, 0xf8, 0x26, 0xbd, 0x39, 0x8a, 0x6f, 0x7b, 0x07, 0x36, 0x78, 0x1c, 0x0b, 0xad,
	0xd9, 0xbc, 0x1e, 0x1f, 0x3b, 0x9d, 0x83, 0x9f, 0xf9, 0xaa, 0xbc, 0x07, 0xe4, 0x7c, 0xff, 0x0d,
	0xcf, 0x34, 0xfd, 0x04, 0x5f, 0x7e, 0x13, 0x99, 0x87, 0x8e, 0x38, 0xe1, 0x99, 0x26, 0xf7, 0xe1,
	0xe6, 0x79, 0x75, 0x75, 0x6a, 0x84, 0xb2, 0x31, 0xf4, 0x53, 0x0c, 0xe9, 0x2f, 0x87, 0x7c, 0x67,
	0xc9, 0x13, 0x9e, 0xbd, 0x1a, 0x26, 0xcb, 0xd2, 0x87, 0x7d, 0xf6, 0x6a, 0x58, 0x68, 0x49, 0x1b,
	0x76, 0x1b, 0x20, 0xce, 0xb9, 0xd6, 0xee, 0xf8, 0x9f, 0xe3, 0xf1, 0x3b, 0x88, 0x34, 0x3e, 0xac,
	0xab, 0x2a, 0x77, 0xec, 0x17, 0xce, 0x87, 0x16, 0x40, 0x72, 0x00, 0xeb, 0x6a, 0xc6, 0x94, 0x28,
	0x2a, 0x23, 0xac, 0x4f, 0xbf, 0x74, 0xed, 0x56, 0xb3, 0x08, 0xa1, 0x30, 0xb1, 0x0a, 0xb3, 0xac,
	0xd8, 0x77, 0x0a, 0xb3, 0x50, 0xdc, 0x81, 0x0d, 0x35, 0x63, 0x8b, 0x6a, 0xcb, 0x84, 0x3e, 0x74,
	0x35, 0x54, 0xb3, 0xb0, 0x41, 0x9d, 0xce, 0xfc, 0x4b, 0xf7, 0x95, 0xd3, 0x99, 0x73, 0xba, 0x31,
	0x6c, 0x69, 0xa1, 0xa6, 0xc2, 0x7b, 0xa5, 0x71, 0xe2, 0x01, 0x6a, 0xaf, 0x3b, 0x2a, 0xac, 0xa7,
	0x7b, 0x8d, 0x21, 0xdf, 0x82, 0xf5, 0x5a, 0x55, 0xa9, 0xcc, 0x7d, 0xa3, 0x1f, 0xa1, 0xb0, 0xeb,
	0xb1, 0xa6, 0x7d, 0xa9, 0xe2, 0x85, 0x48, 0x5c, 0xca, 0x5a, 0x89, 0x54, 0xce, 0xe8, 0x63, 0x14,
	0x6e, 0x3a, 0xc6, 0x66, 0x3c, 0x46, 0x9c, 0xbc, 0x0f, 0x7d, 0xaf, 0x76, 0x42, 0x96, 0x8b, 0x32,
	0x33, 0xcf, 0xe9, 0xd7, 0xd8, 0x04, 0x9f, 0xc9, 0x69, 0x0f, 0x91, 0x21, 0x77, 0xe1, 0xba, 0xb4,
	0xf5, 0x7f, 0xc1, 0x4a, 0x31, 0x33, 0x4c, 0x89, 0x52, 0xbc, 0xa4, 0xdf, 0x0c, 0x82, 0x51, 0x3b,
	0xba, 0x26, 0xf5, 0x11, 0x7f, 0x71, 0x24, 0x66, 0x26, 0xb2, 0xa8, 0x6d, 0x87, 0x56, 0x19, 0xd3,
	0x86, 0x1b, 0x41, 0x43, 0xcc, 0xd8, 0xd6, 0x2a, 0x7b, 0x6a, 0xd7, 0x64, 0x08, 0x3d, 0x4b, 0xca,
	0xd2, 0xa4, 0x4c, 0x55, 0xb9, 0xa0, 0xdf, 0xa2, 0xa0, 0xab, 0x55, 0x16, 0x96, 0x26, 0x8d, 0xaa,
	0x5c, 0x90, 0x9b, 0xb0, 0xa6, 0x55, 0x56, 0xef, 0xb2, 0x9a, 0x3e, 0xc1, 0x1d, 0x5a, 0xb8, 0x3c,
	0xb6, 0x9d, 0xb2, 0xc1, 0x73, 0x23, 0x1f, 0xba, 0x4e, 0x69, 0x95, 0x79, 0x17, 0x0f, 0xdf, 0x85,
	0x8d, 0xf0, 0xf8, 0x59, 0x53, 0xb8, 0x93, 0xb3, 0x5a, 0x90, 0x3e, 0xac, 0x4e, 0x79, 0x7e, 0xda,
	0x8c, 0x2f, 0xb7, 0x18, 0x3e, 0x80, 0xb6, 0x91, 0x85, 0x60, 0x13, 0x9e, 0x11, 0x0a, 0x6b, 0x5a,
	0xc4, 0x55, 0x99, 0x68, 0xd4, 0xf4, 0xa2, 0x66, 0x69, 0xa7, 0x94, 0x55, 0xf9, 0xe1, 0x86, 0xcf,
	0xc3, 0x3f, 0x03, 0xb8, 0xb1, 0x7c, 0x3b, 0x6d, 0xf7, 0x18, 0x37, 0x46, 0x4d, 0xc8, 0xaf, 0x17,
	0x52, 0x34, 0x18, 0x5c, 0x1d, 0x75, 0x77, 0xcb, 0xff, 0x79, 0x52, 0x2c, 0x36, 0x64, 0xd2, 0x88,
	0x22, 0xda, 0x9a, 0xf0, 0xec, 0xc0, 0x71, 0xb6, 0x24, 0xfb, 0x96, 0x19, 0xfe, 0x11, 0xc0, 0xf6,
	0x25, 0x41, 0xe4, 0x06, 0xb4, 0xbc, 0x7d, 0x5c, 0xc5, 0xfc, 0x8a, 0xbc, 0x0d, 0xbd, 0xf3, 0x6e,
	0xb9, 0x82, 0xc5, 0x5a, 0xaf, 0x97, 0x7d, 0x72, 0x1b, 0x20, 0x17, 0x5c, 0x0b, 0x86, 0x75, 0xbb,
	0x8a, 0x8a, 0x0e, 0x22, 0x27, 0xb2, 0x10, 0xd6, 0x78, 0x4a, 0x14, 0x5c, 0x96, 0xf6, 0xf2, 0x2f,
	0x09, 0x57, 0x9c, 0xf1, 0xe6, 0xdc, 0x61, 0x13, 0x31, 0xfc, 0x2b, 0x80, 0x5b, 0x17, 0x0e, 0x43,
	0xf2, 0xfb, 0x65, 0xac, 0x2f, 0xfa, 0x4f, 0xaf, 0x73, 0x3c, 0xbb, 0xba, 0xf7, 0x17, 0x75, 0x0f,
	0x71, 0x5a, 0x87, 0x65, 0x5a, 0x0d, 0x7f, 0xbb, 0x02, 0x6f, 0x5e, 0x1e, 0x68, 0xad, 0x2f, 0x39,
	0x33, 0x67, 0x75, 0x63, 0xd7, 0x96, 0xe4, 0xe8, 0xe2, 0x2d, 0x58, 0x45, 0xb9, 0x2f, 0xfa, 0x8a,
	0xfd, 0x08, 0x58, 0x6b, 0xdb, 0x2f, 0xa5, 0xf6, 0x75, 0x76, 0x0b, 0xdb, 0x27, 0x53, 0x19, 0x9e,
	0xcf, 0xe7, 0x8a, 0x2b, 0xee, 0x3a, 0x82, 0xcd, 0x48, 0xe9, 0xc3, 0xaa, 0xbb, 0xa0, 0xab, 0xee,
	0x56, 0xe0, 0x82, 0xfc, 0x12, 0x40, 0xc7, 0x47, 0x09, 0x4d, 0x5b, 0xf8, 0x75, 0x93, 0xaf, 0xcd,
	0xb3, 0xd1, 0x62, 0xef, 0x49, 0x0b, 0x7f, 0xf5, 0xdc, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0xfb, 0x23, 0x55, 0x19, 0x09, 0x00, 0x00,
}
