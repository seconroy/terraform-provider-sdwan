// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type VPNInterfaceMultilink struct {
	Id                                                 types.String                                         `tfsdk:"id"`
	Version                                            types.Int64                                          `tfsdk:"version"`
	TemplateType                                       types.String                                         `tfsdk:"template_type"`
	Name                                               types.String                                         `tfsdk:"name"`
	Description                                        types.String                                         `tfsdk:"description"`
	DeviceTypes                                        types.Set                                            `tfsdk:"device_types"`
	InterfaceName                                      types.String                                         `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String                                         `tfsdk:"interface_name_variable"`
	MultilinkGroupNumber                               types.Int64                                          `tfsdk:"multilink_group_number"`
	MultilinkGroupNumberVariable                       types.String                                         `tfsdk:"multilink_group_number_variable"`
	InterfaceDescription                               types.String                                         `tfsdk:"interface_description"`
	InterfaceDescriptionVariable                       types.String                                         `tfsdk:"interface_description_variable"`
	Ipv4Address                                        types.String                                         `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                                types.String                                         `tfsdk:"ipv4_address_variable"`
	Ipv6Address                                        types.String                                         `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                                types.String                                         `tfsdk:"ipv6_address_variable"`
	Ipv6AccessLists                                    []VPNInterfaceMultilinkIpv6AccessLists               `tfsdk:"ipv6_access_lists"`
	PppAuthenticationProtocol                          types.String                                         `tfsdk:"ppp_authentication_protocol"`
	PppAuthenticationProtocolPap                       types.Bool                                           `tfsdk:"ppp_authentication_protocol_pap"`
	ChapHostname                                       types.String                                         `tfsdk:"chap_hostname"`
	ChapHostnameVariable                               types.String                                         `tfsdk:"chap_hostname_variable"`
	ChapPppAuthPassword                                types.String                                         `tfsdk:"chap_ppp_auth_password"`
	ChapPppAuthPasswordVariable                        types.String                                         `tfsdk:"chap_ppp_auth_password_variable"`
	PapUsername                                        types.String                                         `tfsdk:"pap_username"`
	PapUsernameVariable                                types.String                                         `tfsdk:"pap_username_variable"`
	PapPassword                                        types.Bool                                           `tfsdk:"pap_password"`
	PapPppAuthPassword                                 types.String                                         `tfsdk:"pap_ppp_auth_password"`
	PapPppAuthPasswordVariable                         types.String                                         `tfsdk:"pap_ppp_auth_password_variable"`
	PppAuthenticationType                              types.String                                         `tfsdk:"ppp_authentication_type"`
	EnableCoreRegion                                   types.Bool                                           `tfsdk:"enable_core_region"`
	EnableCoreRegionVariable                           types.String                                         `tfsdk:"enable_core_region_variable"`
	CoreRegion                                         types.String                                         `tfsdk:"core_region"`
	CoreRegionVariable                                 types.String                                         `tfsdk:"core_region_variable"`
	SecondaryRegion                                    types.String                                         `tfsdk:"secondary_region"`
	SecondaryRegionVariable                            types.String                                         `tfsdk:"secondary_region_variable"`
	TunnelInterfaceEncapsulations                      []VPNInterfaceMultilinkTunnelInterfaceEncapsulations `tfsdk:"tunnel_interface_encapsulations"`
	TunnelInterfaceGroups                              types.Set                                            `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String                                         `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceBorder                              types.Bool                                           `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String                                         `tfsdk:"tunnel_interface_border_variable"`
	PerTunnelQos                                       types.Bool                                           `tfsdk:"per_tunnel_qos"`
	PerTunnelQosVariable                               types.String                                         `tfsdk:"per_tunnel_qos_variable"`
	PerTunnelQosAggregator                             types.Bool                                           `tfsdk:"per_tunnel_qos_aggregator"`
	PerTunnelQosAggregatorVariable                     types.String                                         `tfsdk:"per_tunnel_qos_aggregator_variable"`
	TunnelQosMode                                      types.String                                         `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String                                         `tfsdk:"tunnel_qos_mode_variable"`
	TunnelInterfaceColor                               types.String                                         `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String                                         `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool                                           `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String                                         `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool                                           `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String                                         `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64                                          `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String                                         `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool                                           `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String                                         `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfaceNetworkBroadcast                    types.Bool                                           `tfsdk:"tunnel_interface_network_broadcast"`
	TunnelInterfaceNetworkBroadcastVariable            types.String                                         `tfsdk:"tunnel_interface_network_broadcast_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64                                          `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String                                         `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceControlConnections                  types.Bool                                           `tfsdk:"tunnel_interface_control_connections"`
	TunnelInterfaceControlConnectionsVariable          types.String                                         `tfsdk:"tunnel_interface_control_connections_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool                                           `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String                                         `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.Set                                            `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String                                         `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64                                          `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String                                         `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfacePortHop                             types.Bool                                           `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String                                         `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceColorRestrict                       types.Bool                                           `tfsdk:"tunnel_interface_color_restrict"`
	TunnelInterfaceCarrier                             types.String                                         `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String                                         `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64                                          `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String                                         `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	TunnelInterfaceHelloInterval                       types.Int64                                          `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String                                         `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64                                          `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String                                         `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String                                         `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String                                         `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceAllowAll                            types.Bool                                           `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String                                         `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool                                           `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String                                         `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool                                           `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String                                         `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowDns                            types.Bool                                           `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String                                         `tfsdk:"tunnel_interface_allow_dns_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool                                           `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String                                         `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowSsh                            types.Bool                                           `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String                                         `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowNtp                            types.Bool                                           `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String                                         `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool                                           `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String                                         `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowOspf                           types.Bool                                           `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String                                         `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowStun                           types.Bool                                           `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String                                         `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool                                           `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String                                         `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceAllowHttps                          types.Bool                                           `tfsdk:"tunnel_interface_allow_https"`
	TunnelInterfaceAllowHttpsVariable                  types.String                                         `tfsdk:"tunnel_interface_allow_https_variable"`
	DisableFragmentation                               types.Bool                                           `tfsdk:"disable_fragmentation"`
	FragmentMaxDelay                                   types.Int64                                          `tfsdk:"fragment_max_delay"`
	FragmentMaxDelayVariable                           types.String                                         `tfsdk:"fragment_max_delay_variable"`
	InterleavingFragment                               types.Bool                                           `tfsdk:"interleaving_fragment"`
	ClearDontFragmentBit                               types.Bool                                           `tfsdk:"clear_dont_fragment_bit"`
	ClearDontFragmentBitVariable                       types.String                                         `tfsdk:"clear_dont_fragment_bit_variable"`
	PmtuDiscovery                                      types.Bool                                           `tfsdk:"pmtu_discovery"`
	PmtuDiscoveryVariable                              types.String                                         `tfsdk:"pmtu_discovery_variable"`
	IpMtu                                              types.Int64                                          `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String                                         `tfsdk:"ip_mtu_variable"`
	StaticIngressQos                                   types.Int64                                          `tfsdk:"static_ingress_qos"`
	StaticIngressQosVariable                           types.String                                         `tfsdk:"static_ingress_qos_variable"`
	TcpMss                                             types.Int64                                          `tfsdk:"tcp_mss"`
	TcpMssVariable                                     types.String                                         `tfsdk:"tcp_mss_variable"`
	TlocExtension                                      types.String                                         `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String                                         `tfsdk:"tloc_extension_variable"`
	Shutdown                                           types.Bool                                           `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                         `tfsdk:"shutdown_variable"`
	Autonegotiate                                      types.Bool                                           `tfsdk:"autonegotiate"`
	AutonegotiateVariable                              types.String                                         `tfsdk:"autonegotiate_variable"`
	ShapingRate                                        types.Int64                                          `tfsdk:"shaping_rate"`
	ShapingRateVariable                                types.String                                         `tfsdk:"shaping_rate_variable"`
	QosMap                                             types.String                                         `tfsdk:"qos_map"`
	QosMapVariable                                     types.String                                         `tfsdk:"qos_map_variable"`
	QosMapVpn                                          types.String                                         `tfsdk:"qos_map_vpn"`
	QosMapVpnVariable                                  types.String                                         `tfsdk:"qos_map_vpn_variable"`
	BandwidthUpstream                                  types.Int64                                          `tfsdk:"bandwidth_upstream"`
	BandwidthUpstreamVariable                          types.String                                         `tfsdk:"bandwidth_upstream_variable"`
	BandwidthDownstream                                types.Int64                                          `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String                                         `tfsdk:"bandwidth_downstream_variable"`
	WriteRule                                          types.String                                         `tfsdk:"write_rule"`
	WriteRuleVariable                                  types.String                                         `tfsdk:"write_rule_variable"`
	AccessLists                                        []VPNInterfaceMultilinkAccessLists                   `tfsdk:"access_lists"`
	MultilinkInterfaces                                []VPNInterfaceMultilinkMultilinkInterfaces           `tfsdk:"multilink_interfaces"`
	NimInterfaceList                                   []VPNInterfaceMultilinkNimInterfaceList              `tfsdk:"nim_interface_list"`
}

type VPNInterfaceMultilinkIpv6AccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type VPNInterfaceMultilinkTunnelInterfaceEncapsulations struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Encapsulation      types.String `tfsdk:"encapsulation"`
	Preference         types.Int64  `tfsdk:"preference"`
	PreferenceVariable types.String `tfsdk:"preference_variable"`
	Weight             types.Int64  `tfsdk:"weight"`
	WeightVariable     types.String `tfsdk:"weight_variable"`
}

type VPNInterfaceMultilinkAccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type VPNInterfaceMultilinkMultilinkInterfaces struct {
	Optional            types.Bool                                                 `tfsdk:"optional"`
	InterfaceType       types.String                                               `tfsdk:"interface_type"`
	Slot                types.String                                               `tfsdk:"slot"`
	Framing             types.String                                               `tfsdk:"framing"`
	FramingVariable     types.String                                               `tfsdk:"framing_variable"`
	LineMode            types.String                                               `tfsdk:"line_mode"`
	LineModeVariable    types.String                                               `tfsdk:"line_mode_variable"`
	Internal            types.Bool                                                 `tfsdk:"internal"`
	Description         types.String                                               `tfsdk:"description"`
	DescriptionVariable types.String                                               `tfsdk:"description_variable"`
	Linecode            types.String                                               `tfsdk:"linecode"`
	LinecodeVariable    types.String                                               `tfsdk:"linecode_variable"`
	SetLengthForLong    types.String                                               `tfsdk:"set_length_for_long"`
	SetLengthForShort   types.String                                               `tfsdk:"set_length_for_short"`
	ChannelGroupList    []VPNInterfaceMultilinkMultilinkInterfacesChannelGroupList `tfsdk:"channel_group_list"`
}

type VPNInterfaceMultilinkNimInterfaceList struct {
	Optional                       types.Bool   `tfsdk:"optional"`
	NimSerialInterfaceType         types.String `tfsdk:"nim_serial_interface_type"`
	NimSerialInterfaceTypeVariable types.String `tfsdk:"nim_serial_interface_type_variable"`
	InterfaceName                  types.String `tfsdk:"interface_name"`
	InterfaceNameVariable          types.String `tfsdk:"interface_name_variable"`
	InterfaceDescription           types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable   types.String `tfsdk:"interface_description_variable"`
	Bandwidth                      types.Int64  `tfsdk:"bandwidth"`
	BandwidthVariable              types.String `tfsdk:"bandwidth_variable"`
	ClockRate                      types.Int64  `tfsdk:"clock_rate"`
	ClockRateVariable              types.String `tfsdk:"clock_rate_variable"`
	EncapsulationSerial            types.String `tfsdk:"encapsulation_serial"`
	EncapsulationSerialVariable    types.String `tfsdk:"encapsulation_serial_variable"`
}

type VPNInterfaceMultilinkMultilinkInterfacesChannelGroupList struct {
	Optional             types.Bool   `tfsdk:"optional"`
	ChannelGroup         types.Int64  `tfsdk:"channel_group"`
	ChannelGroupVariable types.String `tfsdk:"channel_group_variable"`
	TimeSlot             types.Set    `tfsdk:"time_slot"`
	TimeSlotVariable     types.String `tfsdk:"time_slot_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data VPNInterfaceMultilink) getModel() string {
	return "vpn-cedge-interface-multilink-controller"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data VPNInterfaceMultilink) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "vpn-cedge-interface-multilink-controller")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.InterfaceNameVariable.ValueString())
	} else if data.InterfaceName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.InterfaceName.ValueString())
	}

	if !data.MultilinkGroupNumberVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipVariableName", data.MultilinkGroupNumberVariable.ValueString())
	} else if data.MultilinkGroupNumber.IsNull() {
		if !gjson.Get(body, path+"ppp.multilink").Exists() {
			body, _ = sjson.Set(body, path+"ppp.multilink", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.multilink.group."+"vipValue", data.MultilinkGroupNumber.ValueInt64())
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.InterfaceDescription.ValueString())
	}

	if !data.Ipv4AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.address."+"vipVariableName", data.Ipv4AddressVariable.ValueString())
	} else if data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.address."+"vipValue", data.Ipv4Address.ValueString())
	}

	if !data.Ipv6AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipVariableName", data.Ipv6AddressVariable.ValueString())
	} else if data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipValue", data.Ipv6Address.ValueString())
	}
	if len(data.Ipv6AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.access-list."+"vipValue.-1", itemBody)
	}
	if data.PppAuthenticationProtocol.IsNull() {
		if !gjson.Get(body, path+"ppp.authentication").Exists() {
			body, _ = sjson.Set(body, path+"ppp.authentication", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipValue", data.PppAuthenticationProtocol.ValueString())
	}
	if !data.PppAuthenticationProtocolPap.IsNull() {
		if data.PppAuthenticationProtocolPap.ValueBool() {
			if !gjson.Get(body, path+"ppp.authentication.pap").Exists() {
				body, _ = sjson.Set(body, path+"ppp.authentication.pap", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ppp.authentication.pap."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ppp.authentication.pap."+"vipType", "ignore")
		}
	}

	if !data.ChapHostnameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipVariableName", data.ChapHostnameVariable.ValueString())
	} else if data.ChapHostname.IsNull() {
		if !gjson.Get(body, path+"ppp.chap").Exists() {
			body, _ = sjson.Set(body, path+"ppp.chap", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.chap.hostname."+"vipValue", data.ChapHostname.ValueString())
	}

	if !data.ChapPppAuthPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipVariableName", data.ChapPppAuthPasswordVariable.ValueString())
	} else if data.ChapPppAuthPassword.IsNull() {
		if !gjson.Get(body, path+"ppp.chap.password").Exists() {
			body, _ = sjson.Set(body, path+"ppp.chap.password", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.chap.password.ppp-auth-password."+"vipValue", data.ChapPppAuthPassword.ValueString())
	}

	if !data.PapUsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipVariableName", data.PapUsernameVariable.ValueString())
	} else if data.PapUsername.IsNull() {
		if !gjson.Get(body, path+"ppp.pap.sent-username.username").Exists() {
			body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.username-string."+"vipValue", data.PapUsername.ValueString())
	}
	if !data.PapPassword.IsNull() {
		if data.PapPassword.ValueBool() {
			if !gjson.Get(body, path+"ppp.pap.sent-username.username.password").Exists() {
				body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.password", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.password."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.password."+"vipType", "ignore")
		}
	}

	if !data.PapPppAuthPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipVariableName", data.PapPppAuthPasswordVariable.ValueString())
	} else if data.PapPppAuthPassword.IsNull() {
		if !gjson.Get(body, path+"ppp.pap.sent-username.username").Exists() {
			body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.pap.sent-username.username.ppp-auth-password."+"vipValue", data.PapPppAuthPassword.ValueString())
	}
	if data.PppAuthenticationType.IsNull() {
		if !gjson.Get(body, path+"ppp.authentication").Exists() {
			body, _ = sjson.Set(body, path+"ppp.authentication", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"ppp.authentication.callin."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ppp.authentication.callin."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.authentication.callin."+"vipValue", data.PppAuthenticationType.ValueString())
	}

	if !data.EnableCoreRegionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipVariableName", data.EnableCoreRegionVariable.ValueString())
	} else if data.EnableCoreRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipValue", strconv.FormatBool(data.EnableCoreRegion.ValueBool()))
	}

	if !data.CoreRegionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipVariableName", data.CoreRegionVariable.ValueString())
	} else if data.CoreRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipValue", data.CoreRegion.ValueString())
	}

	if !data.SecondaryRegionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipVariableName", data.SecondaryRegionVariable.ValueString())
	} else if data.SecondaryRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipValue", data.SecondaryRegion.ValueString())
	}
	if len(data.TunnelInterfaceEncapsulations) > 0 {
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipPrimaryKey", []string{"encap"})
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.TunnelInterfaceEncapsulations {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "encap")
		if item.Encapsulation.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipValue", item.Encapsulation.ValueString())
		}
		itemAttributes = append(itemAttributes, "preference")

		if !item.PreferenceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipVariableName", item.PreferenceVariable.ValueString())
		} else if item.Preference.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipValue", item.Preference.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "weight")

		if !item.WeightVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipVariableName", item.WeightVariable.ValueString())
		} else if item.Weight.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipValue", item.Weight.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tunnel-interface.encapsulation."+"vipValue.-1", itemBody)
	}

	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipVariableName", data.TunnelInterfaceGroupsVariable.ValueString())
	} else if data.TunnelInterfaceGroups.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipType", "constant")
		var values []int64
		data.TunnelInterfaceGroups.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipValue", values)
	}

	if !data.TunnelInterfaceBorderVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipVariableName", data.TunnelInterfaceBorderVariable.ValueString())
	} else if data.TunnelInterfaceBorder.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipValue", strconv.FormatBool(data.TunnelInterfaceBorder.ValueBool()))
	}

	if !data.PerTunnelQosVariable.IsNull() {
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipVariableName", data.PerTunnelQosVariable.ValueString())
	} else if data.PerTunnelQos.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"per-tunnel-qos."+"vipValue", strconv.FormatBool(data.PerTunnelQos.ValueBool()))
	}

	if !data.PerTunnelQosAggregatorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipVariableName", data.PerTunnelQosAggregatorVariable.ValueString())
	} else if data.PerTunnelQosAggregator.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"per-tunnel-qos-aggregator."+"vipValue", strconv.FormatBool(data.PerTunnelQosAggregator.ValueBool()))
	}

	if !data.TunnelQosModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipVariableName", data.TunnelQosModeVariable.ValueString())
	} else if data.TunnelQosMode.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipValue", data.TunnelQosMode.ValueString())
	}

	if !data.TunnelInterfaceColorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipVariableName", data.TunnelInterfaceColorVariable.ValueString())
	} else if data.TunnelInterfaceColor.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipValue", data.TunnelInterfaceColor.ValueString())
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipVariableName", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipValue", strconv.FormatBool(data.TunnelInterfaceLastResortCircuit.ValueBool()))
	}

	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipVariableName", data.TunnelInterfaceLowBandwidthLinkVariable.ValueString())
	} else if data.TunnelInterfaceLowBandwidthLink.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipValue", strconv.FormatBool(data.TunnelInterfaceLowBandwidthLink.ValueBool()))
	}

	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipVariableName", data.TunnelInterfaceTunnelTcpMssVariable.ValueString())
	} else if data.TunnelInterfaceTunnelTcpMss.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipValue", data.TunnelInterfaceTunnelTcpMss.ValueInt64())
	}

	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipVariableName", data.TunnelInterfaceClearDontFragmentVariable.ValueString())
	} else if data.TunnelInterfaceClearDontFragment.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipValue", strconv.FormatBool(data.TunnelInterfaceClearDontFragment.ValueBool()))
	}

	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipVariableName", data.TunnelInterfaceNetworkBroadcastVariable.ValueString())
	} else if data.TunnelInterfaceNetworkBroadcast.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipValue", strconv.FormatBool(data.TunnelInterfaceNetworkBroadcast.ValueBool()))
	}

	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipVariableName", data.TunnelInterfaceMaxControlConnectionsVariable.ValueString())
	} else if data.TunnelInterfaceMaxControlConnections.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipValue", data.TunnelInterfaceMaxControlConnections.ValueInt64())
	}

	if !data.TunnelInterfaceControlConnectionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipVariableName", data.TunnelInterfaceControlConnectionsVariable.ValueString())
	} else if data.TunnelInterfaceControlConnections.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipValue", strconv.FormatBool(data.TunnelInterfaceControlConnections.ValueBool()))
	}

	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipVariableName", data.TunnelInterfaceVbondAsStunServerVariable.ValueString())
	} else if data.TunnelInterfaceVbondAsStunServer.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipValue", strconv.FormatBool(data.TunnelInterfaceVbondAsStunServer.ValueBool()))
	}

	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipVariableName", data.TunnelInterfaceExcludeControllerGroupListVariable.ValueString())
	} else if data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipType", "constant")
		var values []int64
		data.TunnelInterfaceExcludeControllerGroupList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipValue", values)
	}

	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipVariableName", data.TunnelInterfaceVmanageConnectionPreferenceVariable.ValueString())
	} else if data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipValue", data.TunnelInterfaceVmanageConnectionPreference.ValueInt64())
	}

	if !data.TunnelInterfacePortHopVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipVariableName", data.TunnelInterfacePortHopVariable.ValueString())
	} else if data.TunnelInterfacePortHop.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipValue", strconv.FormatBool(data.TunnelInterfacePortHop.ValueBool()))
	}
	if !data.TunnelInterfaceColorRestrict.IsNull() {
		if data.TunnelInterfaceColorRestrict.ValueBool() {
			if !gjson.Get(body, path+"tunnel-interface.color.restrict").Exists() {
				body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipType", "ignore")
		}
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipVariableName", data.TunnelInterfaceCarrierVariable.ValueString())
	} else if data.TunnelInterfaceCarrier.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipValue", data.TunnelInterfaceCarrier.ValueString())
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipVariableName", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipValue", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipVariableName", data.TunnelInterfaceHelloIntervalVariable.ValueString())
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipValue", data.TunnelInterfaceHelloInterval.ValueInt64())
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipVariableName", data.TunnelInterfaceHelloToleranceVariable.ValueString())
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipValue", data.TunnelInterfaceHelloTolerance.ValueInt64())
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipVariableName", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipValue", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
	}

	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipVariableName", data.TunnelInterfaceAllowAllVariable.ValueString())
	} else if data.TunnelInterfaceAllowAll.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowAll.ValueBool()))
	}

	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipVariableName", data.TunnelInterfaceAllowBgpVariable.ValueString())
	} else if data.TunnelInterfaceAllowBgp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowBgp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipVariableName", data.TunnelInterfaceAllowDhcpVariable.ValueString())
	} else if data.TunnelInterfaceAllowDhcp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowDhcp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowDnsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipVariableName", data.TunnelInterfaceAllowDnsVariable.ValueString())
	} else if data.TunnelInterfaceAllowDns.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowDns.ValueBool()))
	}

	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipVariableName", data.TunnelInterfaceAllowIcmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowIcmp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowIcmp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipVariableName", data.TunnelInterfaceAllowSshVariable.ValueString())
	} else if data.TunnelInterfaceAllowSsh.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowSsh.ValueBool()))
	}

	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipVariableName", data.TunnelInterfaceAllowNtpVariable.ValueString())
	} else if data.TunnelInterfaceAllowNtp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowNtp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipVariableName", data.TunnelInterfaceAllowNetconfVariable.ValueString())
	} else if data.TunnelInterfaceAllowNetconf.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowNetconf.ValueBool()))
	}

	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipVariableName", data.TunnelInterfaceAllowOspfVariable.ValueString())
	} else if data.TunnelInterfaceAllowOspf.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowOspf.ValueBool()))
	}

	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipVariableName", data.TunnelInterfaceAllowStunVariable.ValueString())
	} else if data.TunnelInterfaceAllowStun.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowStun.ValueBool()))
	}

	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipVariableName", data.TunnelInterfaceAllowSnmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowSnmp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowSnmp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipVariableName", data.TunnelInterfaceAllowHttpsVariable.ValueString())
	} else if data.TunnelInterfaceAllowHttps.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowHttps.ValueBool()))
	}
	if !data.DisableFragmentation.IsNull() {
		if data.DisableFragmentation.ValueBool() {
			if !gjson.Get(body, path+"ppp.multilink.fragment.disable").Exists() {
				body, _ = sjson.Set(body, path+"ppp.multilink.fragment.disable", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ppp.multilink.fragment.disable."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ppp.multilink.fragment.disable."+"vipType", "ignore")
		}
	}

	if !data.FragmentMaxDelayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipVariableName", data.FragmentMaxDelayVariable.ValueString())
	} else if data.FragmentMaxDelay.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ppp.multilink.fragment.delay.delay-value."+"vipValue", data.FragmentMaxDelay.ValueInt64())
	}
	if !data.InterleavingFragment.IsNull() {
		if data.InterleavingFragment.ValueBool() {
			if !gjson.Get(body, path+"ppp.multilink.interleave").Exists() {
				body, _ = sjson.Set(body, path+"ppp.multilink.interleave", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ppp.multilink.interleave."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ppp.multilink.interleave."+"vipType", "ignore")
		}
	}

	if !data.ClearDontFragmentBitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipVariableName", data.ClearDontFragmentBitVariable.ValueString())
	} else if data.ClearDontFragmentBit.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipValue", strconv.FormatBool(data.ClearDontFragmentBit.ValueBool()))
	}

	if !data.PmtuDiscoveryVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pmtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pmtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pmtu."+"vipVariableName", data.PmtuDiscoveryVariable.ValueString())
	} else if data.PmtuDiscovery.IsNull() {
		body, _ = sjson.Set(body, path+"pmtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pmtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pmtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pmtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pmtu."+"vipValue", strconv.FormatBool(data.PmtuDiscovery.ValueBool()))
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mtu."+"vipVariableName", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mtu."+"vipValue", data.IpMtu.ValueInt64())
	}

	if !data.StaticIngressQosVariable.IsNull() {
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipVariableName", data.StaticIngressQosVariable.ValueString())
	} else if data.StaticIngressQos.IsNull() {
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"static-ingress-qos."+"vipValue", data.StaticIngressQos.ValueInt64())
	}

	if !data.TcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipVariableName", data.TcpMssVariable.ValueString())
	} else if data.TcpMss.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipValue", data.TcpMss.ValueInt64())
	}

	if !data.TlocExtensionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipVariableName", data.TlocExtensionVariable.ValueString())
	} else if data.TlocExtension.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipValue", data.TlocExtension.ValueString())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.AutonegotiateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipVariableName", data.AutonegotiateVariable.ValueString())
	} else if data.Autonegotiate.IsNull() {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipValue", strconv.FormatBool(data.Autonegotiate.ValueBool()))
	}

	if !data.ShapingRateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipVariableName", data.ShapingRateVariable.ValueString())
	} else if data.ShapingRate.IsNull() {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipValue", data.ShapingRate.ValueInt64())
	}

	if !data.QosMapVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-map."+"vipVariableName", data.QosMapVariable.ValueString())
	} else if data.QosMap.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-map."+"vipValue", data.QosMap.ValueString())
	}

	if !data.QosMapVpnVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipVariableName", data.QosMapVpnVariable.ValueString())
	} else if data.QosMapVpn.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipValue", data.QosMapVpn.ValueString())
	}

	if !data.BandwidthUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipVariableName", data.BandwidthUpstreamVariable.ValueString())
	} else if data.BandwidthUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipValue", data.BandwidthUpstream.ValueInt64())
	}

	if !data.BandwidthDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipVariableName", data.BandwidthDownstreamVariable.ValueString())
	} else if data.BandwidthDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipValue", data.BandwidthDownstream.ValueInt64())
	}

	if !data.WriteRuleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipVariableName", data.WriteRuleVariable.ValueString())
	} else if data.WriteRule.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipValue", data.WriteRule.ValueString())
	}
	if len(data.AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"access-list."+"vipValue.-1", itemBody)
	}
	if len(data.MultilinkInterfaces) > 0 {
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipPrimaryKey", []string{"name", "number"})
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipPrimaryKey", []string{"name", "number"})
		body, _ = sjson.Set(body, path+"controller.controller-tx-ex-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.MultilinkInterfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.InterfaceType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.InterfaceType.ValueString())
		}
		itemAttributes = append(itemAttributes, "number")
		if item.Slot.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipValue", item.Slot.ValueString())
		}
		itemAttributes = append(itemAttributes, "framing")

		if !item.FramingVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipVariableName", item.FramingVariable.ValueString())
		} else if item.Framing.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "framing."+"vipValue", item.Framing.ValueString())
		}
		itemAttributes = append(itemAttributes, "clock")

		if !item.LineModeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipVariableName", item.LineModeVariable.ValueString())
		} else if item.LineMode.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "clock.source.line.line-mode."+"vipValue", item.LineMode.ValueString())
		}
		itemAttributes = append(itemAttributes, "clock")
		if !item.Internal.IsNull() {
			if item.Internal.ValueBool() {
				if !gjson.Get(itemBody, "clock.source.internal").Exists() {
					itemBody, _ = sjson.Set(itemBody, "clock.source.internal", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "clock.source.internal."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "clock.source.internal."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "description")

		if !item.DescriptionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipVariableName", item.DescriptionVariable.ValueString())
		} else if item.Description.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipValue", item.Description.ValueString())
		}
		itemAttributes = append(itemAttributes, "linecode")

		if !item.LinecodeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipVariableName", item.LinecodeVariable.ValueString())
		} else if item.Linecode.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "linecode."+"vipValue", item.Linecode.ValueString())
		}
		itemAttributes = append(itemAttributes, "cablelength")
		if item.SetLengthForLong.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "cablelength.long."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "cablelength.long."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "cablelength.long."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "cablelength.long."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "cablelength.long."+"vipValue", item.SetLengthForLong.ValueString())
		}
		itemAttributes = append(itemAttributes, "cablelength")
		if item.SetLengthForShort.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "cablelength.short."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "cablelength.short."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "cablelength.short."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "cablelength.short."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "cablelength.short."+"vipValue", item.SetLengthForShort.ValueString())
		}
		itemAttributes = append(itemAttributes, "channel-group")
		if len(item.ChannelGroupList) > 0 {
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipPrimaryKey", []string{"number"})
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipPrimaryKey", []string{"number"})
			itemBody, _ = sjson.Set(itemBody, "channel-group."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.ChannelGroupList {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "number")

			if !childItem.ChannelGroupVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipVariableName", childItem.ChannelGroupVariable.ValueString())
			} else if childItem.ChannelGroup.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipValue", childItem.ChannelGroup.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "timeslots")

			if !childItem.TimeSlotVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipObjectType", "list")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipVariableName", childItem.TimeSlotVariable.ValueString())
			} else if childItem.TimeSlot.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipObjectType", "list")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipObjectType", "list")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipType", "constant")
				var values []string
				childItem.TimeSlot.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "timeslots."+"vipValue", values)
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "channel-group."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"controller.controller-tx-ex-list."+"vipValue.-1", itemBody)
	}
	if len(data.NimInterfaceList) > 0 {
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipPrimaryKey", []string{"nim", "if-name"})
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipPrimaryKey", []string{"nim", "if-name"})
		body, _ = sjson.Set(body, path+"controller.nim-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.NimInterfaceList {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "nim")

		if !item.NimSerialInterfaceTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipVariableName", item.NimSerialInterfaceTypeVariable.ValueString())
		} else if item.NimSerialInterfaceType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nim."+"vipValue", item.NimSerialInterfaceType.ValueString())
		}
		itemAttributes = append(itemAttributes, "if-name")

		if !item.InterfaceNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.InterfaceNameVariable.ValueString())
		} else if item.InterfaceName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.InterfaceName.ValueString())
		}
		itemAttributes = append(itemAttributes, "description")

		if !item.InterfaceDescriptionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipVariableName", item.InterfaceDescriptionVariable.ValueString())
		} else if item.InterfaceDescription.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipValue", item.InterfaceDescription.ValueString())
		}
		itemAttributes = append(itemAttributes, "bandwidth")

		if !item.BandwidthVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipVariableName", item.BandwidthVariable.ValueString())
		} else if item.Bandwidth.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "bandwidth."+"vipValue", item.Bandwidth.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "DCE-mode-config")

		if !item.ClockRateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipVariableName", item.ClockRateVariable.ValueString())
		} else if item.ClockRate.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "DCE-mode-config.clock-rate."+"vipValue", item.ClockRate.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "encapsulation-serial")

		if !item.EncapsulationSerialVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipVariableName", item.EncapsulationSerialVariable.ValueString())
		} else if item.EncapsulationSerial.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "encapsulation-serial."+"vipValue", item.EncapsulationSerial.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"controller.nim-list."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *VPNInterfaceMultilink) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.InterfaceNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceName = types.StringNull()
			data.InterfaceNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.InterfaceName = types.StringValue(v.String())
			data.InterfaceNameVariable = types.StringNull()
		}
	} else {
		data.InterfaceName = types.StringNull()
		data.InterfaceNameVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.multilink.group.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MultilinkGroupNumber = types.Int64Null()

			v := res.Get(path + "ppp.multilink.group.vipVariableName")
			data.MultilinkGroupNumberVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MultilinkGroupNumber = types.Int64Null()
			data.MultilinkGroupNumberVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.multilink.group.vipValue")
			data.MultilinkGroupNumber = types.Int64Value(v.Int())
			data.MultilinkGroupNumberVariable = types.StringNull()
		}
	} else {
		data.MultilinkGroupNumber = types.Int64Null()
		data.MultilinkGroupNumberVariable = types.StringNull()
	}
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.InterfaceDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceDescription = types.StringNull()
			data.InterfaceDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.InterfaceDescription = types.StringValue(v.String())
			data.InterfaceDescriptionVariable = types.StringNull()
		}
	} else {
		data.InterfaceDescription = types.StringNull()
		data.InterfaceDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4Address = types.StringNull()

			v := res.Get(path + "ip.address.vipVariableName")
			data.Ipv4AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4Address = types.StringNull()
			data.Ipv4AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address.vipValue")
			data.Ipv4Address = types.StringValue(v.String())
			data.Ipv4AddressVariable = types.StringNull()
		}
	} else {
		data.Ipv4Address = types.StringNull()
		data.Ipv4AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Address = types.StringNull()

			v := res.Get(path + "ipv6.address.vipVariableName")
			data.Ipv6AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Address = types.StringNull()
			data.Ipv6AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.address.vipValue")
			data.Ipv6Address = types.StringValue(v.String())
			data.Ipv6AddressVariable = types.StringNull()
		}
	} else {
		data.Ipv6Address = types.StringNull()
		data.Ipv6AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.access-list.vipValue"); len(value.Array()) > 0 {
		data.Ipv6AccessLists = make([]VPNInterfaceMultilinkIpv6AccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceMultilinkIpv6AccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.Ipv6AccessLists = append(data.Ipv6AccessLists, item)
			return true
		})
	} else {
		if len(data.Ipv6AccessLists) > 0 {
			data.Ipv6AccessLists = []VPNInterfaceMultilinkIpv6AccessLists{}
		}
	}
	if value := res.Get(path + "ppp.authentication.method.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PppAuthenticationProtocol = types.StringNull()

		} else if value.String() == "ignore" {
			data.PppAuthenticationProtocol = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.authentication.method.vipValue")
			data.PppAuthenticationProtocol = types.StringValue(v.String())

		}
	} else {
		data.PppAuthenticationProtocol = types.StringNull()

	}
	if value := res.Get(path + "ppp.authentication.pap.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PppAuthenticationProtocolPap = types.BoolNull()

		} else if value.String() == "ignore" {
			data.PppAuthenticationProtocolPap = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.authentication.pap.vipValue")
			data.PppAuthenticationProtocolPap = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ppp.authentication.pap"); value.Exists() {
		data.PppAuthenticationProtocolPap = types.BoolValue(true)

	} else {
		data.PppAuthenticationProtocolPap = types.BoolNull()

	}
	if value := res.Get(path + "ppp.chap.hostname.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ChapHostname = types.StringNull()

			v := res.Get(path + "ppp.chap.hostname.vipVariableName")
			data.ChapHostnameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ChapHostname = types.StringNull()
			data.ChapHostnameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.chap.hostname.vipValue")
			data.ChapHostname = types.StringValue(v.String())
			data.ChapHostnameVariable = types.StringNull()
		}
	} else {
		data.ChapHostname = types.StringNull()
		data.ChapHostnameVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.chap.password.ppp-auth-password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ChapPppAuthPassword = types.StringNull()

			v := res.Get(path + "ppp.chap.password.ppp-auth-password.vipVariableName")
			data.ChapPppAuthPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ChapPppAuthPassword = types.StringNull()
			data.ChapPppAuthPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.chap.password.ppp-auth-password.vipValue")
			data.ChapPppAuthPassword = types.StringValue(v.String())
			data.ChapPppAuthPasswordVariable = types.StringNull()
		}
	} else {
		data.ChapPppAuthPassword = types.StringNull()
		data.ChapPppAuthPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.pap.sent-username.username.username-string.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PapUsername = types.StringNull()

			v := res.Get(path + "ppp.pap.sent-username.username.username-string.vipVariableName")
			data.PapUsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PapUsername = types.StringNull()
			data.PapUsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.pap.sent-username.username.username-string.vipValue")
			data.PapUsername = types.StringValue(v.String())
			data.PapUsernameVariable = types.StringNull()
		}
	} else {
		data.PapUsername = types.StringNull()
		data.PapUsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.pap.sent-username.username.password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PapPassword = types.BoolNull()

		} else if value.String() == "ignore" {
			data.PapPassword = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.pap.sent-username.username.password.vipValue")
			data.PapPassword = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ppp.pap.sent-username.username.password"); value.Exists() {
		data.PapPassword = types.BoolValue(true)

	} else {
		data.PapPassword = types.BoolNull()

	}
	if value := res.Get(path + "ppp.pap.sent-username.username.ppp-auth-password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PapPppAuthPassword = types.StringNull()

			v := res.Get(path + "ppp.pap.sent-username.username.ppp-auth-password.vipVariableName")
			data.PapPppAuthPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PapPppAuthPassword = types.StringNull()
			data.PapPppAuthPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.pap.sent-username.username.ppp-auth-password.vipValue")
			data.PapPppAuthPassword = types.StringValue(v.String())
			data.PapPppAuthPasswordVariable = types.StringNull()
		}
	} else {
		data.PapPppAuthPassword = types.StringNull()
		data.PapPppAuthPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.authentication.callin.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PppAuthenticationType = types.StringNull()

		} else if value.String() == "ignore" {
			data.PppAuthenticationType = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.authentication.callin.vipValue")
			data.PppAuthenticationType = types.StringValue(v.String())

		}
	} else {
		data.PppAuthenticationType = types.StringNull()

	}
	if value := res.Get(path + "tunnel-interface.enable-core-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableCoreRegion = types.BoolNull()

			v := res.Get(path + "tunnel-interface.enable-core-region.vipVariableName")
			data.EnableCoreRegionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EnableCoreRegion = types.BoolNull()
			data.EnableCoreRegionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.enable-core-region.vipValue")
			data.EnableCoreRegion = types.BoolValue(v.Bool())
			data.EnableCoreRegionVariable = types.StringNull()
		}
	} else {
		data.EnableCoreRegion = types.BoolNull()
		data.EnableCoreRegionVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.core-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CoreRegion = types.StringNull()

			v := res.Get(path + "tunnel-interface.core-region.vipVariableName")
			data.CoreRegionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CoreRegion = types.StringNull()
			data.CoreRegionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.core-region.vipValue")
			data.CoreRegion = types.StringValue(v.String())
			data.CoreRegionVariable = types.StringNull()
		}
	} else {
		data.CoreRegion = types.StringNull()
		data.CoreRegionVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.secondary-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SecondaryRegion = types.StringNull()

			v := res.Get(path + "tunnel-interface.secondary-region.vipVariableName")
			data.SecondaryRegionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SecondaryRegion = types.StringNull()
			data.SecondaryRegionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.secondary-region.vipValue")
			data.SecondaryRegion = types.StringValue(v.String())
			data.SecondaryRegionVariable = types.StringNull()
		}
	} else {
		data.SecondaryRegion = types.StringNull()
		data.SecondaryRegionVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.encapsulation.vipValue"); len(value.Array()) > 0 {
		data.TunnelInterfaceEncapsulations = make([]VPNInterfaceMultilinkTunnelInterfaceEncapsulations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceMultilinkTunnelInterfaceEncapsulations{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("encap.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Encapsulation = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Encapsulation = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("encap.vipValue")
					item.Encapsulation = types.StringValue(cv.String())

				}
			} else {
				item.Encapsulation = types.StringNull()

			}
			if cValue := v.Get("preference.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Preference = types.Int64Null()

					cv := v.Get("preference.vipVariableName")
					item.PreferenceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Preference = types.Int64Null()
					item.PreferenceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("preference.vipValue")
					item.Preference = types.Int64Value(cv.Int())
					item.PreferenceVariable = types.StringNull()
				}
			} else {
				item.Preference = types.Int64Null()
				item.PreferenceVariable = types.StringNull()
			}
			if cValue := v.Get("weight.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Weight = types.Int64Null()

					cv := v.Get("weight.vipVariableName")
					item.WeightVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Weight = types.Int64Null()
					item.WeightVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("weight.vipValue")
					item.Weight = types.Int64Value(cv.Int())
					item.WeightVariable = types.StringNull()
				}
			} else {
				item.Weight = types.Int64Null()
				item.WeightVariable = types.StringNull()
			}
			data.TunnelInterfaceEncapsulations = append(data.TunnelInterfaceEncapsulations, item)
			return true
		})
	} else {
		if len(data.TunnelInterfaceEncapsulations) > 0 {
			data.TunnelInterfaceEncapsulations = []VPNInterfaceMultilinkTunnelInterfaceEncapsulations{}
		}
	}
	if value := res.Get(path + "tunnel-interface.group.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)

			v := res.Get(path + "tunnel-interface.group.vipVariableName")
			data.TunnelInterfaceGroupsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)
			data.TunnelInterfaceGroupsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.group.vipValue")
			data.TunnelInterfaceGroups = helpers.GetInt64Set(v.Array())
			data.TunnelInterfaceGroupsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)
		data.TunnelInterfaceGroupsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.border.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceBorder = types.BoolNull()

			v := res.Get(path + "tunnel-interface.border.vipVariableName")
			data.TunnelInterfaceBorderVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceBorder = types.BoolNull()
			data.TunnelInterfaceBorderVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.border.vipValue")
			data.TunnelInterfaceBorder = types.BoolValue(v.Bool())
			data.TunnelInterfaceBorderVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceBorder = types.BoolNull()
		data.TunnelInterfaceBorderVariable = types.StringNull()
	}
	if value := res.Get(path + "per-tunnel-qos.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PerTunnelQos = types.BoolNull()

			v := res.Get(path + "per-tunnel-qos.vipVariableName")
			data.PerTunnelQosVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PerTunnelQos = types.BoolNull()
			data.PerTunnelQosVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "per-tunnel-qos.vipValue")
			data.PerTunnelQos = types.BoolValue(v.Bool())
			data.PerTunnelQosVariable = types.StringNull()
		}
	} else {
		data.PerTunnelQos = types.BoolNull()
		data.PerTunnelQosVariable = types.StringNull()
	}
	if value := res.Get(path + "per-tunnel-qos-aggregator.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PerTunnelQosAggregator = types.BoolNull()

			v := res.Get(path + "per-tunnel-qos-aggregator.vipVariableName")
			data.PerTunnelQosAggregatorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PerTunnelQosAggregator = types.BoolNull()
			data.PerTunnelQosAggregatorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "per-tunnel-qos-aggregator.vipValue")
			data.PerTunnelQosAggregator = types.BoolValue(v.Bool())
			data.PerTunnelQosAggregatorVariable = types.StringNull()
		}
	} else {
		data.PerTunnelQosAggregator = types.BoolNull()
		data.PerTunnelQosAggregatorVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelQosMode = types.StringNull()

			v := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipVariableName")
			data.TunnelQosModeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelQosMode = types.StringNull()
			data.TunnelQosModeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipValue")
			data.TunnelQosMode = types.StringValue(v.String())
			data.TunnelQosModeVariable = types.StringNull()
		}
	} else {
		data.TunnelQosMode = types.StringNull()
		data.TunnelQosModeVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.color.value.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceColor = types.StringNull()

			v := res.Get(path + "tunnel-interface.color.value.vipVariableName")
			data.TunnelInterfaceColorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceColor = types.StringNull()
			data.TunnelInterfaceColorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.color.value.vipValue")
			data.TunnelInterfaceColor = types.StringValue(v.String())
			data.TunnelInterfaceColorVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceColor = types.StringNull()
		data.TunnelInterfaceColorVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.last-resort-circuit.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceLastResortCircuit = types.BoolNull()

			v := res.Get(path + "tunnel-interface.last-resort-circuit.vipVariableName")
			data.TunnelInterfaceLastResortCircuitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceLastResortCircuit = types.BoolNull()
			data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.last-resort-circuit.vipValue")
			data.TunnelInterfaceLastResortCircuit = types.BoolValue(v.Bool())
			data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceLastResortCircuit = types.BoolNull()
		data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.low-bandwidth-link.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolNull()

			v := res.Get(path + "tunnel-interface.low-bandwidth-link.vipVariableName")
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.low-bandwidth-link.vipValue")
			data.TunnelInterfaceLowBandwidthLink = types.BoolValue(v.Bool())
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
		data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Null()

			v := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipVariableName")
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipValue")
			data.TunnelInterfaceTunnelTcpMss = types.Int64Value(v.Int())
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
		data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.clear-dont-fragment.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceClearDontFragment = types.BoolNull()

			v := res.Get(path + "tunnel-interface.clear-dont-fragment.vipVariableName")
			data.TunnelInterfaceClearDontFragmentVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceClearDontFragment = types.BoolNull()
			data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.clear-dont-fragment.vipValue")
			data.TunnelInterfaceClearDontFragment = types.BoolValue(v.Bool())
			data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceClearDontFragment = types.BoolNull()
		data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.network-broadcast.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolNull()

			v := res.Get(path + "tunnel-interface.network-broadcast.vipVariableName")
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.network-broadcast.vipValue")
			data.TunnelInterfaceNetworkBroadcast = types.BoolValue(v.Bool())
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
		data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.max-control-connections.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Null()

			v := res.Get(path + "tunnel-interface.max-control-connections.vipVariableName")
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Null()
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.max-control-connections.vipValue")
			data.TunnelInterfaceMaxControlConnections = types.Int64Value(v.Int())
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceMaxControlConnections = types.Int64Null()
		data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.control-connections.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceControlConnections = types.BoolNull()

			v := res.Get(path + "tunnel-interface.control-connections.vipVariableName")
			data.TunnelInterfaceControlConnectionsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceControlConnections = types.BoolNull()
			data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.control-connections.vipValue")
			data.TunnelInterfaceControlConnections = types.BoolValue(v.Bool())
			data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceControlConnections = types.BoolNull()
		data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolNull()

			v := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipVariableName")
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipValue")
			data.TunnelInterfaceVbondAsStunServer = types.BoolValue(v.Bool())
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
		data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)

			v := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipVariableName")
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipValue")
			data.TunnelInterfaceExcludeControllerGroupList = helpers.GetInt64Set(v.Array())
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
		data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()

			v := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipVariableName")
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipValue")
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Value(v.Int())
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
		data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.port-hop.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfacePortHop = types.BoolNull()

			v := res.Get(path + "tunnel-interface.port-hop.vipVariableName")
			data.TunnelInterfacePortHopVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfacePortHop = types.BoolNull()
			data.TunnelInterfacePortHopVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.port-hop.vipValue")
			data.TunnelInterfacePortHop = types.BoolValue(v.Bool())
			data.TunnelInterfacePortHopVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfacePortHop = types.BoolNull()
		data.TunnelInterfacePortHopVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.color.restrict.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceColorRestrict = types.BoolNull()

		} else if value.String() == "ignore" {
			data.TunnelInterfaceColorRestrict = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.color.restrict.vipValue")
			data.TunnelInterfaceColorRestrict = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "tunnel-interface.color.restrict"); value.Exists() {
		data.TunnelInterfaceColorRestrict = types.BoolValue(true)

	} else {
		data.TunnelInterfaceColorRestrict = types.BoolNull()

	}
	if value := res.Get(path + "tunnel-interface.carrier.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceCarrier = types.StringNull()

			v := res.Get(path + "tunnel-interface.carrier.vipVariableName")
			data.TunnelInterfaceCarrierVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceCarrier = types.StringNull()
			data.TunnelInterfaceCarrierVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.carrier.vipValue")
			data.TunnelInterfaceCarrier = types.StringValue(v.String())
			data.TunnelInterfaceCarrierVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceCarrier = types.StringNull()
		data.TunnelInterfaceCarrierVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.nat-refresh-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Null()

			v := res.Get(path + "tunnel-interface.nat-refresh-interval.vipVariableName")
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.nat-refresh-interval.vipValue")
			data.TunnelInterfaceNatRefreshInterval = types.Int64Value(v.Int())
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
		data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.hello-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceHelloInterval = types.Int64Null()

			v := res.Get(path + "tunnel-interface.hello-interval.vipVariableName")
			data.TunnelInterfaceHelloIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceHelloInterval = types.Int64Null()
			data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.hello-interval.vipValue")
			data.TunnelInterfaceHelloInterval = types.Int64Value(v.Int())
			data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceHelloInterval = types.Int64Null()
		data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.hello-tolerance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceHelloTolerance = types.Int64Null()

			v := res.Get(path + "tunnel-interface.hello-tolerance.vipVariableName")
			data.TunnelInterfaceHelloToleranceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceHelloTolerance = types.Int64Null()
			data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.hello-tolerance.vipValue")
			data.TunnelInterfaceHelloTolerance = types.Int64Value(v.Int())
			data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceHelloTolerance = types.Int64Null()
		data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.bind.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()

			v := res.Get(path + "tunnel-interface.bind.vipVariableName")
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.bind.vipValue")
			data.TunnelInterfaceBindLoopbackTunnel = types.StringValue(v.String())
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
		data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.all.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowAll = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.all.vipVariableName")
			data.TunnelInterfaceAllowAllVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowAll = types.BoolNull()
			data.TunnelInterfaceAllowAllVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.all.vipValue")
			data.TunnelInterfaceAllowAll = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowAllVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowAll = types.BoolNull()
		data.TunnelInterfaceAllowAllVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.bgp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowBgp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.bgp.vipVariableName")
			data.TunnelInterfaceAllowBgpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowBgp = types.BoolNull()
			data.TunnelInterfaceAllowBgpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.bgp.vipValue")
			data.TunnelInterfaceAllowBgp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowBgpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowBgp = types.BoolNull()
		data.TunnelInterfaceAllowBgpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.dhcp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowDhcp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.dhcp.vipVariableName")
			data.TunnelInterfaceAllowDhcpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowDhcp = types.BoolNull()
			data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.dhcp.vipValue")
			data.TunnelInterfaceAllowDhcp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowDhcp = types.BoolNull()
		data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.dns.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowDns = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.dns.vipVariableName")
			data.TunnelInterfaceAllowDnsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowDns = types.BoolNull()
			data.TunnelInterfaceAllowDnsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.dns.vipValue")
			data.TunnelInterfaceAllowDns = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowDnsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowDns = types.BoolNull()
		data.TunnelInterfaceAllowDnsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.icmp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowIcmp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.icmp.vipVariableName")
			data.TunnelInterfaceAllowIcmpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowIcmp = types.BoolNull()
			data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.icmp.vipValue")
			data.TunnelInterfaceAllowIcmp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowIcmp = types.BoolNull()
		data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.sshd.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowSsh = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.sshd.vipVariableName")
			data.TunnelInterfaceAllowSshVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowSsh = types.BoolNull()
			data.TunnelInterfaceAllowSshVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.sshd.vipValue")
			data.TunnelInterfaceAllowSsh = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowSshVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowSsh = types.BoolNull()
		data.TunnelInterfaceAllowSshVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.ntp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowNtp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.ntp.vipVariableName")
			data.TunnelInterfaceAllowNtpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowNtp = types.BoolNull()
			data.TunnelInterfaceAllowNtpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.ntp.vipValue")
			data.TunnelInterfaceAllowNtp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowNtpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowNtp = types.BoolNull()
		data.TunnelInterfaceAllowNtpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.netconf.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowNetconf = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.netconf.vipVariableName")
			data.TunnelInterfaceAllowNetconfVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowNetconf = types.BoolNull()
			data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.netconf.vipValue")
			data.TunnelInterfaceAllowNetconf = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowNetconf = types.BoolNull()
		data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.ospf.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowOspf = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.ospf.vipVariableName")
			data.TunnelInterfaceAllowOspfVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowOspf = types.BoolNull()
			data.TunnelInterfaceAllowOspfVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.ospf.vipValue")
			data.TunnelInterfaceAllowOspf = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowOspfVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowOspf = types.BoolNull()
		data.TunnelInterfaceAllowOspfVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.stun.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowStun = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.stun.vipVariableName")
			data.TunnelInterfaceAllowStunVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowStun = types.BoolNull()
			data.TunnelInterfaceAllowStunVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.stun.vipValue")
			data.TunnelInterfaceAllowStun = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowStunVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowStun = types.BoolNull()
		data.TunnelInterfaceAllowStunVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.snmp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowSnmp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.snmp.vipVariableName")
			data.TunnelInterfaceAllowSnmpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowSnmp = types.BoolNull()
			data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.snmp.vipValue")
			data.TunnelInterfaceAllowSnmp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowSnmp = types.BoolNull()
		data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.https.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowHttps = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.https.vipVariableName")
			data.TunnelInterfaceAllowHttpsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowHttps = types.BoolNull()
			data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.https.vipValue")
			data.TunnelInterfaceAllowHttps = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowHttps = types.BoolNull()
		data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.multilink.fragment.disable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DisableFragmentation = types.BoolNull()

		} else if value.String() == "ignore" {
			data.DisableFragmentation = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.multilink.fragment.disable.vipValue")
			data.DisableFragmentation = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ppp.multilink.fragment.disable"); value.Exists() {
		data.DisableFragmentation = types.BoolValue(true)

	} else {
		data.DisableFragmentation = types.BoolNull()

	}
	if value := res.Get(path + "ppp.multilink.fragment.delay.delay-value.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.FragmentMaxDelay = types.Int64Null()

			v := res.Get(path + "ppp.multilink.fragment.delay.delay-value.vipVariableName")
			data.FragmentMaxDelayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.FragmentMaxDelay = types.Int64Null()
			data.FragmentMaxDelayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.multilink.fragment.delay.delay-value.vipValue")
			data.FragmentMaxDelay = types.Int64Value(v.Int())
			data.FragmentMaxDelayVariable = types.StringNull()
		}
	} else {
		data.FragmentMaxDelay = types.Int64Null()
		data.FragmentMaxDelayVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.multilink.interleave.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterleavingFragment = types.BoolNull()

		} else if value.String() == "ignore" {
			data.InterleavingFragment = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.multilink.interleave.vipValue")
			data.InterleavingFragment = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ppp.multilink.interleave"); value.Exists() {
		data.InterleavingFragment = types.BoolValue(true)

	} else {
		data.InterleavingFragment = types.BoolNull()

	}
	if value := res.Get(path + "clear-dont-fragment.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ClearDontFragmentBit = types.BoolNull()

			v := res.Get(path + "clear-dont-fragment.vipVariableName")
			data.ClearDontFragmentBitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ClearDontFragmentBit = types.BoolNull()
			data.ClearDontFragmentBitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "clear-dont-fragment.vipValue")
			data.ClearDontFragmentBit = types.BoolValue(v.Bool())
			data.ClearDontFragmentBitVariable = types.StringNull()
		}
	} else {
		data.ClearDontFragmentBit = types.BoolNull()
		data.ClearDontFragmentBitVariable = types.StringNull()
	}
	if value := res.Get(path + "pmtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PmtuDiscovery = types.BoolNull()

			v := res.Get(path + "pmtu.vipVariableName")
			data.PmtuDiscoveryVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PmtuDiscovery = types.BoolNull()
			data.PmtuDiscoveryVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pmtu.vipValue")
			data.PmtuDiscovery = types.BoolValue(v.Bool())
			data.PmtuDiscoveryVariable = types.StringNull()
		}
	} else {
		data.PmtuDiscovery = types.BoolNull()
		data.PmtuDiscoveryVariable = types.StringNull()
	}
	if value := res.Get(path + "mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpMtu = types.Int64Null()

			v := res.Get(path + "mtu.vipVariableName")
			data.IpMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpMtu = types.Int64Null()
			data.IpMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mtu.vipValue")
			data.IpMtu = types.Int64Value(v.Int())
			data.IpMtuVariable = types.StringNull()
		}
	} else {
		data.IpMtu = types.Int64Null()
		data.IpMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "static-ingress-qos.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.StaticIngressQos = types.Int64Null()

			v := res.Get(path + "static-ingress-qos.vipVariableName")
			data.StaticIngressQosVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.StaticIngressQos = types.Int64Null()
			data.StaticIngressQosVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "static-ingress-qos.vipValue")
			data.StaticIngressQos = types.Int64Value(v.Int())
			data.StaticIngressQosVariable = types.StringNull()
		}
	} else {
		data.StaticIngressQos = types.Int64Null()
		data.StaticIngressQosVariable = types.StringNull()
	}
	if value := res.Get(path + "tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpMss = types.Int64Null()

			v := res.Get(path + "tcp-mss-adjust.vipVariableName")
			data.TcpMssVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpMss = types.Int64Null()
			data.TcpMssVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tcp-mss-adjust.vipValue")
			data.TcpMss = types.Int64Value(v.Int())
			data.TcpMssVariable = types.StringNull()
		}
	} else {
		data.TcpMss = types.Int64Null()
		data.TcpMssVariable = types.StringNull()
	}
	if value := res.Get(path + "tloc-extension.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TlocExtension = types.StringNull()

			v := res.Get(path + "tloc-extension.vipVariableName")
			data.TlocExtensionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TlocExtension = types.StringNull()
			data.TlocExtensionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tloc-extension.vipValue")
			data.TlocExtension = types.StringValue(v.String())
			data.TlocExtensionVariable = types.StringNull()
		}
	} else {
		data.TlocExtension = types.StringNull()
		data.TlocExtensionVariable = types.StringNull()
	}
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "autonegotiate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Autonegotiate = types.BoolNull()

			v := res.Get(path + "autonegotiate.vipVariableName")
			data.AutonegotiateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Autonegotiate = types.BoolNull()
			data.AutonegotiateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "autonegotiate.vipValue")
			data.Autonegotiate = types.BoolValue(v.Bool())
			data.AutonegotiateVariable = types.StringNull()
		}
	} else {
		data.Autonegotiate = types.BoolNull()
		data.AutonegotiateVariable = types.StringNull()
	}
	if value := res.Get(path + "shaping-rate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ShapingRate = types.Int64Null()

			v := res.Get(path + "shaping-rate.vipVariableName")
			data.ShapingRateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ShapingRate = types.Int64Null()
			data.ShapingRateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shaping-rate.vipValue")
			data.ShapingRate = types.Int64Value(v.Int())
			data.ShapingRateVariable = types.StringNull()
		}
	} else {
		data.ShapingRate = types.Int64Null()
		data.ShapingRateVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-map.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosMap = types.StringNull()

			v := res.Get(path + "qos-map.vipVariableName")
			data.QosMapVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosMap = types.StringNull()
			data.QosMapVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-map.vipValue")
			data.QosMap = types.StringValue(v.String())
			data.QosMapVariable = types.StringNull()
		}
	} else {
		data.QosMap = types.StringNull()
		data.QosMapVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-map-vpn.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosMapVpn = types.StringNull()

			v := res.Get(path + "qos-map-vpn.vipVariableName")
			data.QosMapVpnVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosMapVpn = types.StringNull()
			data.QosMapVpnVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-map-vpn.vipValue")
			data.QosMapVpn = types.StringValue(v.String())
			data.QosMapVpnVariable = types.StringNull()
		}
	} else {
		data.QosMapVpn = types.StringNull()
		data.QosMapVpnVariable = types.StringNull()
	}
	if value := res.Get(path + "bandwidth-upstream.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BandwidthUpstream = types.Int64Null()

			v := res.Get(path + "bandwidth-upstream.vipVariableName")
			data.BandwidthUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BandwidthUpstream = types.Int64Null()
			data.BandwidthUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bandwidth-upstream.vipValue")
			data.BandwidthUpstream = types.Int64Value(v.Int())
			data.BandwidthUpstreamVariable = types.StringNull()
		}
	} else {
		data.BandwidthUpstream = types.Int64Null()
		data.BandwidthUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "bandwidth-downstream.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BandwidthDownstream = types.Int64Null()

			v := res.Get(path + "bandwidth-downstream.vipVariableName")
			data.BandwidthDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BandwidthDownstream = types.Int64Null()
			data.BandwidthDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bandwidth-downstream.vipValue")
			data.BandwidthDownstream = types.Int64Value(v.Int())
			data.BandwidthDownstreamVariable = types.StringNull()
		}
	} else {
		data.BandwidthDownstream = types.Int64Null()
		data.BandwidthDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "rewrite-rule.rule-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.WriteRule = types.StringNull()

			v := res.Get(path + "rewrite-rule.rule-name.vipVariableName")
			data.WriteRuleVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.WriteRule = types.StringNull()
			data.WriteRuleVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "rewrite-rule.rule-name.vipValue")
			data.WriteRule = types.StringValue(v.String())
			data.WriteRuleVariable = types.StringNull()
		}
	} else {
		data.WriteRule = types.StringNull()
		data.WriteRuleVariable = types.StringNull()
	}
	if value := res.Get(path + "access-list.vipValue"); len(value.Array()) > 0 {
		data.AccessLists = make([]VPNInterfaceMultilinkAccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceMultilinkAccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.AccessLists = append(data.AccessLists, item)
			return true
		})
	} else {
		if len(data.AccessLists) > 0 {
			data.AccessLists = []VPNInterfaceMultilinkAccessLists{}
		}
	}
	if value := res.Get(path + "controller.controller-tx-ex-list.vipValue"); len(value.Array()) > 0 {
		data.MultilinkInterfaces = make([]VPNInterfaceMultilinkMultilinkInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceMultilinkMultilinkInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.InterfaceType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.InterfaceType = types.StringValue(cv.String())

				}
			} else {
				item.InterfaceType = types.StringNull()

			}
			if cValue := v.Get("number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Slot = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Slot = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("number.vipValue")
					item.Slot = types.StringValue(cv.String())

				}
			} else {
				item.Slot = types.StringNull()

			}
			if cValue := v.Get("framing.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Framing = types.StringNull()

					cv := v.Get("framing.vipVariableName")
					item.FramingVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Framing = types.StringNull()
					item.FramingVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("framing.vipValue")
					item.Framing = types.StringValue(cv.String())
					item.FramingVariable = types.StringNull()
				}
			} else {
				item.Framing = types.StringNull()
				item.FramingVariable = types.StringNull()
			}
			if cValue := v.Get("clock.source.line.line-mode.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LineMode = types.StringNull()

					cv := v.Get("clock.source.line.line-mode.vipVariableName")
					item.LineModeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.LineMode = types.StringNull()
					item.LineModeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("clock.source.line.line-mode.vipValue")
					item.LineMode = types.StringValue(cv.String())
					item.LineModeVariable = types.StringNull()
				}
			} else {
				item.LineMode = types.StringNull()
				item.LineModeVariable = types.StringNull()
			}
			if cValue := v.Get("clock.source.internal.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Internal = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Internal = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("clock.source.internal.vipValue")
					item.Internal = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("clock.source.internal"); cValue.Exists() {
				item.Internal = types.BoolValue(true)

			} else {
				item.Internal = types.BoolNull()

			}
			if cValue := v.Get("description.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Description = types.StringNull()

					cv := v.Get("description.vipVariableName")
					item.DescriptionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Description = types.StringNull()
					item.DescriptionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("description.vipValue")
					item.Description = types.StringValue(cv.String())
					item.DescriptionVariable = types.StringNull()
				}
			} else {
				item.Description = types.StringNull()
				item.DescriptionVariable = types.StringNull()
			}
			if cValue := v.Get("linecode.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Linecode = types.StringNull()

					cv := v.Get("linecode.vipVariableName")
					item.LinecodeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Linecode = types.StringNull()
					item.LinecodeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("linecode.vipValue")
					item.Linecode = types.StringValue(cv.String())
					item.LinecodeVariable = types.StringNull()
				}
			} else {
				item.Linecode = types.StringNull()
				item.LinecodeVariable = types.StringNull()
			}
			if cValue := v.Get("cablelength.long.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SetLengthForLong = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SetLengthForLong = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("cablelength.long.vipValue")
					item.SetLengthForLong = types.StringValue(cv.String())

				}
			} else {
				item.SetLengthForLong = types.StringNull()

			}
			if cValue := v.Get("cablelength.short.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SetLengthForShort = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SetLengthForShort = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("cablelength.short.vipValue")
					item.SetLengthForShort = types.StringValue(cv.String())

				}
			} else {
				item.SetLengthForShort = types.StringNull()

			}
			if cValue := v.Get("channel-group.vipValue"); len(cValue.Array()) > 0 {
				item.ChannelGroupList = make([]VPNInterfaceMultilinkMultilinkInterfacesChannelGroupList, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := VPNInterfaceMultilinkMultilinkInterfacesChannelGroupList{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("number.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.ChannelGroup = types.Int64Null()

							ccv := cv.Get("number.vipVariableName")
							cItem.ChannelGroupVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.ChannelGroup = types.Int64Null()
							cItem.ChannelGroupVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("number.vipValue")
							cItem.ChannelGroup = types.Int64Value(ccv.Int())
							cItem.ChannelGroupVariable = types.StringNull()
						}
					} else {
						cItem.ChannelGroup = types.Int64Null()
						cItem.ChannelGroupVariable = types.StringNull()
					}
					if ccValue := cv.Get("timeslots.vipType"); len(ccValue.Array()) > 0 {
						if ccValue.String() == "variableName" {
							cItem.TimeSlot = types.SetNull(types.StringType)

							ccv := cv.Get("timeslots.vipVariableName")
							cItem.TimeSlotVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.TimeSlot = types.SetNull(types.StringType)
							cItem.TimeSlotVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("timeslots.vipValue")
							cItem.TimeSlot = helpers.GetStringSet(ccv.Array())
							cItem.TimeSlotVariable = types.StringNull()
						}
					} else {
						cItem.TimeSlot = types.SetNull(types.StringType)
						cItem.TimeSlotVariable = types.StringNull()
					}
					item.ChannelGroupList = append(item.ChannelGroupList, cItem)
					return true
				})
			} else {
				if len(item.ChannelGroupList) > 0 {
					item.ChannelGroupList = []VPNInterfaceMultilinkMultilinkInterfacesChannelGroupList{}
				}
			}
			data.MultilinkInterfaces = append(data.MultilinkInterfaces, item)
			return true
		})
	} else {
		if len(data.MultilinkInterfaces) > 0 {
			data.MultilinkInterfaces = []VPNInterfaceMultilinkMultilinkInterfaces{}
		}
	}
	if value := res.Get(path + "controller.nim-list.vipValue"); len(value.Array()) > 0 {
		data.NimInterfaceList = make([]VPNInterfaceMultilinkNimInterfaceList, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceMultilinkNimInterfaceList{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("nim.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NimSerialInterfaceType = types.StringNull()

					cv := v.Get("nim.vipVariableName")
					item.NimSerialInterfaceTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NimSerialInterfaceType = types.StringNull()
					item.NimSerialInterfaceTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nim.vipValue")
					item.NimSerialInterfaceType = types.StringValue(cv.String())
					item.NimSerialInterfaceTypeVariable = types.StringNull()
				}
			} else {
				item.NimSerialInterfaceType = types.StringNull()
				item.NimSerialInterfaceTypeVariable = types.StringNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceName = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.InterfaceNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceName = types.StringNull()
					item.InterfaceNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.InterfaceName = types.StringValue(cv.String())
					item.InterfaceNameVariable = types.StringNull()
				}
			} else {
				item.InterfaceName = types.StringNull()
				item.InterfaceNameVariable = types.StringNull()
			}
			if cValue := v.Get("description.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceDescription = types.StringNull()

					cv := v.Get("description.vipVariableName")
					item.InterfaceDescriptionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceDescription = types.StringNull()
					item.InterfaceDescriptionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("description.vipValue")
					item.InterfaceDescription = types.StringValue(cv.String())
					item.InterfaceDescriptionVariable = types.StringNull()
				}
			} else {
				item.InterfaceDescription = types.StringNull()
				item.InterfaceDescriptionVariable = types.StringNull()
			}
			if cValue := v.Get("bandwidth.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Bandwidth = types.Int64Null()

					cv := v.Get("bandwidth.vipVariableName")
					item.BandwidthVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Bandwidth = types.Int64Null()
					item.BandwidthVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("bandwidth.vipValue")
					item.Bandwidth = types.Int64Value(cv.Int())
					item.BandwidthVariable = types.StringNull()
				}
			} else {
				item.Bandwidth = types.Int64Null()
				item.BandwidthVariable = types.StringNull()
			}
			if cValue := v.Get("DCE-mode-config.clock-rate.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ClockRate = types.Int64Null()

					cv := v.Get("DCE-mode-config.clock-rate.vipVariableName")
					item.ClockRateVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ClockRate = types.Int64Null()
					item.ClockRateVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("DCE-mode-config.clock-rate.vipValue")
					item.ClockRate = types.Int64Value(cv.Int())
					item.ClockRateVariable = types.StringNull()
				}
			} else {
				item.ClockRate = types.Int64Null()
				item.ClockRateVariable = types.StringNull()
			}
			if cValue := v.Get("encapsulation-serial.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EncapsulationSerial = types.StringNull()

					cv := v.Get("encapsulation-serial.vipVariableName")
					item.EncapsulationSerialVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EncapsulationSerial = types.StringNull()
					item.EncapsulationSerialVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("encapsulation-serial.vipValue")
					item.EncapsulationSerial = types.StringValue(cv.String())
					item.EncapsulationSerialVariable = types.StringNull()
				}
			} else {
				item.EncapsulationSerial = types.StringNull()
				item.EncapsulationSerialVariable = types.StringNull()
			}
			data.NimInterfaceList = append(data.NimInterfaceList, item)
			return true
		})
	} else {
		if len(data.NimInterfaceList) > 0 {
			data.NimInterfaceList = []VPNInterfaceMultilinkNimInterfaceList{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *VPNInterfaceMultilink) hasChanges(ctx context.Context, state *VPNInterfaceMultilink) bool {
	hasChanges := false
	if !data.InterfaceName.Equal(state.InterfaceName) {
		hasChanges = true
	}
	if !data.MultilinkGroupNumber.Equal(state.MultilinkGroupNumber) {
		hasChanges = true
	}
	if !data.InterfaceDescription.Equal(state.InterfaceDescription) {
		hasChanges = true
	}
	if !data.Ipv4Address.Equal(state.Ipv4Address) {
		hasChanges = true
	}
	if !data.Ipv6Address.Equal(state.Ipv6Address) {
		hasChanges = true
	}
	if len(data.Ipv6AccessLists) != len(state.Ipv6AccessLists) {
		hasChanges = true
	} else {
		for i := range data.Ipv6AccessLists {
			if !data.Ipv6AccessLists[i].Direction.Equal(state.Ipv6AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.Ipv6AccessLists[i].AclName.Equal(state.Ipv6AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if !data.PppAuthenticationProtocol.Equal(state.PppAuthenticationProtocol) {
		hasChanges = true
	}
	if !data.PppAuthenticationProtocolPap.Equal(state.PppAuthenticationProtocolPap) {
		hasChanges = true
	}
	if !data.ChapHostname.Equal(state.ChapHostname) {
		hasChanges = true
	}
	if !data.ChapPppAuthPassword.Equal(state.ChapPppAuthPassword) {
		hasChanges = true
	}
	if !data.PapUsername.Equal(state.PapUsername) {
		hasChanges = true
	}
	if !data.PapPassword.Equal(state.PapPassword) {
		hasChanges = true
	}
	if !data.PapPppAuthPassword.Equal(state.PapPppAuthPassword) {
		hasChanges = true
	}
	if !data.PppAuthenticationType.Equal(state.PppAuthenticationType) {
		hasChanges = true
	}
	if !data.EnableCoreRegion.Equal(state.EnableCoreRegion) {
		hasChanges = true
	}
	if !data.CoreRegion.Equal(state.CoreRegion) {
		hasChanges = true
	}
	if !data.SecondaryRegion.Equal(state.SecondaryRegion) {
		hasChanges = true
	}
	if len(data.TunnelInterfaceEncapsulations) != len(state.TunnelInterfaceEncapsulations) {
		hasChanges = true
	} else {
		for i := range data.TunnelInterfaceEncapsulations {
			if !data.TunnelInterfaceEncapsulations[i].Encapsulation.Equal(state.TunnelInterfaceEncapsulations[i].Encapsulation) {
				hasChanges = true
			}
			if !data.TunnelInterfaceEncapsulations[i].Preference.Equal(state.TunnelInterfaceEncapsulations[i].Preference) {
				hasChanges = true
			}
			if !data.TunnelInterfaceEncapsulations[i].Weight.Equal(state.TunnelInterfaceEncapsulations[i].Weight) {
				hasChanges = true
			}
		}
	}
	if !data.TunnelInterfaceGroups.Equal(state.TunnelInterfaceGroups) {
		hasChanges = true
	}
	if !data.TunnelInterfaceBorder.Equal(state.TunnelInterfaceBorder) {
		hasChanges = true
	}
	if !data.PerTunnelQos.Equal(state.PerTunnelQos) {
		hasChanges = true
	}
	if !data.PerTunnelQosAggregator.Equal(state.PerTunnelQosAggregator) {
		hasChanges = true
	}
	if !data.TunnelQosMode.Equal(state.TunnelQosMode) {
		hasChanges = true
	}
	if !data.TunnelInterfaceColor.Equal(state.TunnelInterfaceColor) {
		hasChanges = true
	}
	if !data.TunnelInterfaceLastResortCircuit.Equal(state.TunnelInterfaceLastResortCircuit) {
		hasChanges = true
	}
	if !data.TunnelInterfaceLowBandwidthLink.Equal(state.TunnelInterfaceLowBandwidthLink) {
		hasChanges = true
	}
	if !data.TunnelInterfaceTunnelTcpMss.Equal(state.TunnelInterfaceTunnelTcpMss) {
		hasChanges = true
	}
	if !data.TunnelInterfaceClearDontFragment.Equal(state.TunnelInterfaceClearDontFragment) {
		hasChanges = true
	}
	if !data.TunnelInterfaceNetworkBroadcast.Equal(state.TunnelInterfaceNetworkBroadcast) {
		hasChanges = true
	}
	if !data.TunnelInterfaceMaxControlConnections.Equal(state.TunnelInterfaceMaxControlConnections) {
		hasChanges = true
	}
	if !data.TunnelInterfaceControlConnections.Equal(state.TunnelInterfaceControlConnections) {
		hasChanges = true
	}
	if !data.TunnelInterfaceVbondAsStunServer.Equal(state.TunnelInterfaceVbondAsStunServer) {
		hasChanges = true
	}
	if !data.TunnelInterfaceExcludeControllerGroupList.Equal(state.TunnelInterfaceExcludeControllerGroupList) {
		hasChanges = true
	}
	if !data.TunnelInterfaceVmanageConnectionPreference.Equal(state.TunnelInterfaceVmanageConnectionPreference) {
		hasChanges = true
	}
	if !data.TunnelInterfacePortHop.Equal(state.TunnelInterfacePortHop) {
		hasChanges = true
	}
	if !data.TunnelInterfaceColorRestrict.Equal(state.TunnelInterfaceColorRestrict) {
		hasChanges = true
	}
	if !data.TunnelInterfaceCarrier.Equal(state.TunnelInterfaceCarrier) {
		hasChanges = true
	}
	if !data.TunnelInterfaceNatRefreshInterval.Equal(state.TunnelInterfaceNatRefreshInterval) {
		hasChanges = true
	}
	if !data.TunnelInterfaceHelloInterval.Equal(state.TunnelInterfaceHelloInterval) {
		hasChanges = true
	}
	if !data.TunnelInterfaceHelloTolerance.Equal(state.TunnelInterfaceHelloTolerance) {
		hasChanges = true
	}
	if !data.TunnelInterfaceBindLoopbackTunnel.Equal(state.TunnelInterfaceBindLoopbackTunnel) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowAll.Equal(state.TunnelInterfaceAllowAll) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowBgp.Equal(state.TunnelInterfaceAllowBgp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowDhcp.Equal(state.TunnelInterfaceAllowDhcp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowDns.Equal(state.TunnelInterfaceAllowDns) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowIcmp.Equal(state.TunnelInterfaceAllowIcmp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowSsh.Equal(state.TunnelInterfaceAllowSsh) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowNtp.Equal(state.TunnelInterfaceAllowNtp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowNetconf.Equal(state.TunnelInterfaceAllowNetconf) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowOspf.Equal(state.TunnelInterfaceAllowOspf) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowStun.Equal(state.TunnelInterfaceAllowStun) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowSnmp.Equal(state.TunnelInterfaceAllowSnmp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowHttps.Equal(state.TunnelInterfaceAllowHttps) {
		hasChanges = true
	}
	if !data.DisableFragmentation.Equal(state.DisableFragmentation) {
		hasChanges = true
	}
	if !data.FragmentMaxDelay.Equal(state.FragmentMaxDelay) {
		hasChanges = true
	}
	if !data.InterleavingFragment.Equal(state.InterleavingFragment) {
		hasChanges = true
	}
	if !data.ClearDontFragmentBit.Equal(state.ClearDontFragmentBit) {
		hasChanges = true
	}
	if !data.PmtuDiscovery.Equal(state.PmtuDiscovery) {
		hasChanges = true
	}
	if !data.IpMtu.Equal(state.IpMtu) {
		hasChanges = true
	}
	if !data.StaticIngressQos.Equal(state.StaticIngressQos) {
		hasChanges = true
	}
	if !data.TcpMss.Equal(state.TcpMss) {
		hasChanges = true
	}
	if !data.TlocExtension.Equal(state.TlocExtension) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.Autonegotiate.Equal(state.Autonegotiate) {
		hasChanges = true
	}
	if !data.ShapingRate.Equal(state.ShapingRate) {
		hasChanges = true
	}
	if !data.QosMap.Equal(state.QosMap) {
		hasChanges = true
	}
	if !data.QosMapVpn.Equal(state.QosMapVpn) {
		hasChanges = true
	}
	if !data.BandwidthUpstream.Equal(state.BandwidthUpstream) {
		hasChanges = true
	}
	if !data.BandwidthDownstream.Equal(state.BandwidthDownstream) {
		hasChanges = true
	}
	if !data.WriteRule.Equal(state.WriteRule) {
		hasChanges = true
	}
	if len(data.AccessLists) != len(state.AccessLists) {
		hasChanges = true
	} else {
		for i := range data.AccessLists {
			if !data.AccessLists[i].Direction.Equal(state.AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.AccessLists[i].AclName.Equal(state.AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if len(data.MultilinkInterfaces) != len(state.MultilinkInterfaces) {
		hasChanges = true
	} else {
		for i := range data.MultilinkInterfaces {
			if !data.MultilinkInterfaces[i].InterfaceType.Equal(state.MultilinkInterfaces[i].InterfaceType) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].Slot.Equal(state.MultilinkInterfaces[i].Slot) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].Framing.Equal(state.MultilinkInterfaces[i].Framing) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].LineMode.Equal(state.MultilinkInterfaces[i].LineMode) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].Internal.Equal(state.MultilinkInterfaces[i].Internal) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].Description.Equal(state.MultilinkInterfaces[i].Description) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].Linecode.Equal(state.MultilinkInterfaces[i].Linecode) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].SetLengthForLong.Equal(state.MultilinkInterfaces[i].SetLengthForLong) {
				hasChanges = true
			}
			if !data.MultilinkInterfaces[i].SetLengthForShort.Equal(state.MultilinkInterfaces[i].SetLengthForShort) {
				hasChanges = true
			}
			if len(data.MultilinkInterfaces[i].ChannelGroupList) != len(state.MultilinkInterfaces[i].ChannelGroupList) {
				hasChanges = true
			} else {
				for ii := range data.MultilinkInterfaces[i].ChannelGroupList {
					if !data.MultilinkInterfaces[i].ChannelGroupList[ii].ChannelGroup.Equal(state.MultilinkInterfaces[i].ChannelGroupList[ii].ChannelGroup) {
						hasChanges = true
					}
					if !data.MultilinkInterfaces[i].ChannelGroupList[ii].TimeSlot.Equal(state.MultilinkInterfaces[i].ChannelGroupList[ii].TimeSlot) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.NimInterfaceList) != len(state.NimInterfaceList) {
		hasChanges = true
	} else {
		for i := range data.NimInterfaceList {
			if !data.NimInterfaceList[i].NimSerialInterfaceType.Equal(state.NimInterfaceList[i].NimSerialInterfaceType) {
				hasChanges = true
			}
			if !data.NimInterfaceList[i].InterfaceName.Equal(state.NimInterfaceList[i].InterfaceName) {
				hasChanges = true
			}
			if !data.NimInterfaceList[i].InterfaceDescription.Equal(state.NimInterfaceList[i].InterfaceDescription) {
				hasChanges = true
			}
			if !data.NimInterfaceList[i].Bandwidth.Equal(state.NimInterfaceList[i].Bandwidth) {
				hasChanges = true
			}
			if !data.NimInterfaceList[i].ClockRate.Equal(state.NimInterfaceList[i].ClockRate) {
				hasChanges = true
			}
			if !data.NimInterfaceList[i].EncapsulationSerial.Equal(state.NimInterfaceList[i].EncapsulationSerial) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
