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
type VPNInterfaceDSLPPPoE struct {
	Id                                                 types.String                                        `tfsdk:"id"`
	Version                                            types.Int64                                         `tfsdk:"version"`
	TemplateType                                       types.String                                        `tfsdk:"template_type"`
	Name                                               types.String                                        `tfsdk:"name"`
	Description                                        types.String                                        `tfsdk:"description"`
	DeviceTypes                                        types.Set                                           `tfsdk:"device_types"`
	EthernetInterfaceName                              types.String                                        `tfsdk:"ethernet_interface_name"`
	EthernetInterfaceNameVariable                      types.String                                        `tfsdk:"ethernet_interface_name_variable"`
	Shutdown                                           types.Bool                                          `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                        `tfsdk:"shutdown_variable"`
	EthernetDescription                                types.String                                        `tfsdk:"ethernet_description"`
	EthernetDescriptionVariable                        types.String                                        `tfsdk:"ethernet_description_variable"`
	VdslConfiguration                                  []VPNInterfaceDSLPPPoEVdslConfiguration             `tfsdk:"vdsl_configuration"`
	EthernetVlanId                                     types.Int64                                         `tfsdk:"ethernet_vlan_id"`
	EthernetVlanIdVariable                             types.String                                        `tfsdk:"ethernet_vlan_id_variable"`
	EthernetDialerPoolNumber                           types.Int64                                         `tfsdk:"ethernet_dialer_pool_number"`
	EthernetDialerPoolNumberVariable                   types.String                                        `tfsdk:"ethernet_dialer_pool_number_variable"`
	EthernetPppMaximumPayload                          types.Int64                                         `tfsdk:"ethernet_ppp_maximum_payload"`
	EthernetPppMaximumPayloadVariable                  types.String                                        `tfsdk:"ethernet_ppp_maximum_payload_variable"`
	EthernetDialerAddressNegotiated                    types.Bool                                          `tfsdk:"ethernet_dialer_address_negotiated"`
	EthernetUnnumberedLoopbackInterface                types.String                                        `tfsdk:"ethernet_unnumbered_loopback_interface"`
	EthernetUnnumberedLoopbackInterfaceVariable        types.String                                        `tfsdk:"ethernet_unnumbered_loopback_interface_variable"`
	PppAuthenticationProtocol                          types.String                                        `tfsdk:"ppp_authentication_protocol"`
	PppAuthenticationProtocolVariable                  types.String                                        `tfsdk:"ppp_authentication_protocol_variable"`
	PppAuthenticationProtocolPap                       types.Bool                                          `tfsdk:"ppp_authentication_protocol_pap"`
	ChapHostname                                       types.String                                        `tfsdk:"chap_hostname"`
	ChapHostnameVariable                               types.String                                        `tfsdk:"chap_hostname_variable"`
	ChapPppAuthPassword                                types.String                                        `tfsdk:"chap_ppp_auth_password"`
	ChapPppAuthPasswordVariable                        types.String                                        `tfsdk:"chap_ppp_auth_password_variable"`
	PapUsername                                        types.String                                        `tfsdk:"pap_username"`
	PapUsernameVariable                                types.String                                        `tfsdk:"pap_username_variable"`
	PapPassword                                        types.Bool                                          `tfsdk:"pap_password"`
	PapPppAuthPassword                                 types.String                                        `tfsdk:"pap_ppp_auth_password"`
	PapPppAuthPasswordVariable                         types.String                                        `tfsdk:"pap_ppp_auth_password_variable"`
	PppAuthenticationType                              types.String                                        `tfsdk:"ppp_authentication_type"`
	EnableCoreRegion                                   types.Bool                                          `tfsdk:"enable_core_region"`
	EnableCoreRegionVariable                           types.String                                        `tfsdk:"enable_core_region_variable"`
	CoreRegion                                         types.String                                        `tfsdk:"core_region"`
	CoreRegionVariable                                 types.String                                        `tfsdk:"core_region_variable"`
	SecondaryRegion                                    types.String                                        `tfsdk:"secondary_region"`
	SecondaryRegionVariable                            types.String                                        `tfsdk:"secondary_region_variable"`
	TunnelInterfaceEncapsulations                      []VPNInterfaceDSLPPPoETunnelInterfaceEncapsulations `tfsdk:"tunnel_interface_encapsulations"`
	TunnelInterfaceGroups                              types.Set                                           `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String                                        `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceBorder                              types.Bool                                          `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String                                        `tfsdk:"tunnel_interface_border_variable"`
	PerTunnelQos                                       types.Bool                                          `tfsdk:"per_tunnel_qos"`
	PerTunnelQosVariable                               types.String                                        `tfsdk:"per_tunnel_qos_variable"`
	PerTunnelQosAggregator                             types.Bool                                          `tfsdk:"per_tunnel_qos_aggregator"`
	PerTunnelQosAggregatorVariable                     types.String                                        `tfsdk:"per_tunnel_qos_aggregator_variable"`
	TunnelQosMode                                      types.String                                        `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String                                        `tfsdk:"tunnel_qos_mode_variable"`
	TunnelInterfaceColor                               types.String                                        `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String                                        `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool                                          `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String                                        `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool                                          `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String                                        `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64                                         `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String                                        `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool                                          `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String                                        `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfaceNetworkBroadcast                    types.Bool                                          `tfsdk:"tunnel_interface_network_broadcast"`
	TunnelInterfaceNetworkBroadcastVariable            types.String                                        `tfsdk:"tunnel_interface_network_broadcast_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64                                         `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String                                        `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceControlConnections                  types.Bool                                          `tfsdk:"tunnel_interface_control_connections"`
	TunnelInterfaceControlConnectionsVariable          types.String                                        `tfsdk:"tunnel_interface_control_connections_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool                                          `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String                                        `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.Set                                           `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String                                        `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64                                         `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String                                        `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfacePortHop                             types.Bool                                          `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String                                        `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceColorRestrict                       types.Bool                                          `tfsdk:"tunnel_interface_color_restrict"`
	TunnelInterfaceCarrier                             types.String                                        `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String                                        `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64                                         `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String                                        `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	TunnelInterfaceHelloInterval                       types.Int64                                         `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String                                        `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64                                         `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String                                        `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String                                        `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String                                        `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceAllowAll                            types.Bool                                          `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String                                        `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool                                          `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String                                        `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool                                          `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String                                        `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowDns                            types.Bool                                          `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String                                        `tfsdk:"tunnel_interface_allow_dns_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool                                          `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String                                        `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowSsh                            types.Bool                                          `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String                                        `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowNtp                            types.Bool                                          `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String                                        `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool                                          `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String                                        `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowOspf                           types.Bool                                          `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String                                        `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowStun                           types.Bool                                          `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String                                        `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool                                          `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String                                        `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceAllowHttps                          types.Bool                                          `tfsdk:"tunnel_interface_allow_https"`
	TunnelInterfaceAllowHttpsVariable                  types.String                                        `tfsdk:"tunnel_interface_allow_https_variable"`
	Nat                                                types.Bool                                          `tfsdk:"nat"`
	NatRefreshMode                                     types.String                                        `tfsdk:"nat_refresh_mode"`
	NatRefreshModeVariable                             types.String                                        `tfsdk:"nat_refresh_mode_variable"`
	NatUdpTimeout                                      types.Int64                                         `tfsdk:"nat_udp_timeout"`
	NatUdpTimeoutVariable                              types.String                                        `tfsdk:"nat_udp_timeout_variable"`
	NatTcpTimeout                                      types.Int64                                         `tfsdk:"nat_tcp_timeout"`
	NatTcpTimeoutVariable                              types.String                                        `tfsdk:"nat_tcp_timeout_variable"`
	NatBlockIcmpError                                  types.Bool                                          `tfsdk:"nat_block_icmp_error"`
	NatBlockIcmpErrorVariable                          types.String                                        `tfsdk:"nat_block_icmp_error_variable"`
	NatResponseToPing                                  types.Bool                                          `tfsdk:"nat_response_to_ping"`
	NatResponseToPingVariable                          types.String                                        `tfsdk:"nat_response_to_ping_variable"`
	NatPortForwards                                    []VPNInterfaceDSLPPPoENatPortForwards               `tfsdk:"nat_port_forwards"`
	QosAdaptivePeriod                                  types.Int64                                         `tfsdk:"qos_adaptive_period"`
	QosAdaptivePeriodVariable                          types.String                                        `tfsdk:"qos_adaptive_period_variable"`
	QosAdaptiveBandwidthDownstream                     types.Int64                                         `tfsdk:"qos_adaptive_bandwidth_downstream"`
	QosAdaptiveBandwidthDownstreamVariable             types.String                                        `tfsdk:"qos_adaptive_bandwidth_downstream_variable"`
	QosAdaptiveMinDownstream                           types.Int64                                         `tfsdk:"qos_adaptive_min_downstream"`
	QosAdaptiveMinDownstreamVariable                   types.String                                        `tfsdk:"qos_adaptive_min_downstream_variable"`
	QosAdaptiveMaxDownstream                           types.Int64                                         `tfsdk:"qos_adaptive_max_downstream"`
	QosAdaptiveMaxDownstreamVariable                   types.String                                        `tfsdk:"qos_adaptive_max_downstream_variable"`
	QosAdaptiveBandwidthUpstream                       types.Int64                                         `tfsdk:"qos_adaptive_bandwidth_upstream"`
	QosAdaptiveBandwidthUpstreamVariable               types.String                                        `tfsdk:"qos_adaptive_bandwidth_upstream_variable"`
	QosAdaptiveMinUpstream                             types.Int64                                         `tfsdk:"qos_adaptive_min_upstream"`
	QosAdaptiveMinUpstreamVariable                     types.String                                        `tfsdk:"qos_adaptive_min_upstream_variable"`
	QosAdaptiveMaxUpstream                             types.Int64                                         `tfsdk:"qos_adaptive_max_upstream"`
	QosAdaptiveMaxUpstreamVariable                     types.String                                        `tfsdk:"qos_adaptive_max_upstream_variable"`
	ShapingRate                                        types.Int64                                         `tfsdk:"shaping_rate"`
	ShapingRateVariable                                types.String                                        `tfsdk:"shaping_rate_variable"`
	QosMap                                             types.String                                        `tfsdk:"qos_map"`
	QosMapVariable                                     types.String                                        `tfsdk:"qos_map_variable"`
	VpnQosMap                                          types.String                                        `tfsdk:"vpn_qos_map"`
	VpnQosMapVariable                                  types.String                                        `tfsdk:"vpn_qos_map_variable"`
	BandwidthUpstream                                  types.Int64                                         `tfsdk:"bandwidth_upstream"`
	BandwidthUpstreamVariable                          types.String                                        `tfsdk:"bandwidth_upstream_variable"`
	BandwidthDownstream                                types.Int64                                         `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String                                        `tfsdk:"bandwidth_downstream_variable"`
	WriteRule                                          types.String                                        `tfsdk:"write_rule"`
	WriteRuleVariable                                  types.String                                        `tfsdk:"write_rule_variable"`
	AccessLists                                        []VPNInterfaceDSLPPPoEAccessLists                   `tfsdk:"access_lists"`
	Policers                                           []VPNInterfaceDSLPPPoEPolicers                      `tfsdk:"policers"`
	IpMtu                                              types.Int64                                         `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String                                        `tfsdk:"ip_mtu_variable"`
	TcpMss                                             types.Int64                                         `tfsdk:"tcp_mss"`
	TcpMssVariable                                     types.String                                        `tfsdk:"tcp_mss_variable"`
	TlocExtension                                      types.String                                        `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String                                        `tfsdk:"tloc_extension_variable"`
	Tracker                                            types.Set                                           `tfsdk:"tracker"`
	TrackerVariable                                    types.String                                        `tfsdk:"tracker_variable"`
	IpDirectedBroadcast                                types.Bool                                          `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable                        types.String                                        `tfsdk:"ip_directed_broadcast_variable"`
}

type VPNInterfaceDSLPPPoEVdslConfiguration struct {
	Optional                       types.Bool   `tfsdk:"optional"`
	ControllerVdslSlot             types.String `tfsdk:"controller_vdsl_slot"`
	ControllerVdslSlotVariable     types.String `tfsdk:"controller_vdsl_slot_variable"`
	Sra                            types.Bool   `tfsdk:"sra"`
	ModeAdsl1                      types.Bool   `tfsdk:"mode_adsl1"`
	ModeAdsl2                      types.Bool   `tfsdk:"mode_adsl2"`
	ModeAdsl2plus                  types.Bool   `tfsdk:"mode_adsl2plus"`
	ModeVdsl2                      types.Bool   `tfsdk:"mode_vdsl2"`
	ModeAnsi                       types.Bool   `tfsdk:"mode_ansi"`
	VdslModemConfiguration         types.String `tfsdk:"vdsl_modem_configuration"`
	VdslModemConfigurationVariable types.String `tfsdk:"vdsl_modem_configuration_variable"`
}

type VPNInterfaceDSLPPPoETunnelInterfaceEncapsulations struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Encapsulation      types.String `tfsdk:"encapsulation"`
	Preference         types.Int64  `tfsdk:"preference"`
	PreferenceVariable types.String `tfsdk:"preference_variable"`
	Weight             types.Int64  `tfsdk:"weight"`
	WeightVariable     types.String `tfsdk:"weight_variable"`
}

type VPNInterfaceDSLPPPoENatPortForwards struct {
	Optional                 types.Bool   `tfsdk:"optional"`
	PortStartRange           types.Int64  `tfsdk:"port_start_range"`
	PortEndRange             types.Int64  `tfsdk:"port_end_range"`
	Protocol                 types.String `tfsdk:"protocol"`
	PrivateVpn               types.Int64  `tfsdk:"private_vpn"`
	PrivateVpnVariable       types.String `tfsdk:"private_vpn_variable"`
	PrivateIpAddress         types.String `tfsdk:"private_ip_address"`
	PrivateIpAddressVariable types.String `tfsdk:"private_ip_address_variable"`
}

type VPNInterfaceDSLPPPoEAccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type VPNInterfaceDSLPPPoEPolicers struct {
	Optional    types.Bool   `tfsdk:"optional"`
	Direction   types.String `tfsdk:"direction"`
	PolicerName types.String `tfsdk:"policer_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data VPNInterfaceDSLPPPoE) getModel() string {
	return "vpn-interface-pppoe"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data VPNInterfaceDSLPPPoE) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "vpn-interface-pppoe")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.EthernetInterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.EthernetInterfaceNameVariable.ValueString())
	} else if data.EthernetInterfaceName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.EthernetInterfaceName.ValueString())
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

	if !data.EthernetDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.EthernetDescriptionVariable.ValueString())
	} else if data.EthernetDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.EthernetDescription.ValueString())
	}
	if len(data.VdslConfiguration) > 0 {
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"controller.vdsl."+"vipValue", []interface{}{})
	}
	for _, item := range data.VdslConfiguration {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.ControllerVdslSlotVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.ControllerVdslSlotVariable.ValueString())
		} else if item.ControllerVdslSlot.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.ControllerVdslSlot.ValueString())
		}
		itemAttributes = append(itemAttributes, "sra")
		if !item.Sra.IsNull() {
			if item.Sra.ValueBool() {
				if !gjson.Get(itemBody, "sra").Exists() {
					itemBody, _ = sjson.Set(itemBody, "sra", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "sra."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "sra."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "operating")
		if !item.ModeAdsl1.IsNull() {
			if item.ModeAdsl1.ValueBool() {
				if !gjson.Get(itemBody, "operating.mode.auto.adsl1").Exists() {
					itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl1", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl1."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl1."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "operating")
		if !item.ModeAdsl2.IsNull() {
			if item.ModeAdsl2.ValueBool() {
				if !gjson.Get(itemBody, "operating.mode.auto.adsl2").Exists() {
					itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "operating")
		if !item.ModeAdsl2plus.IsNull() {
			if item.ModeAdsl2plus.ValueBool() {
				if !gjson.Get(itemBody, "operating.mode.auto.adsl2plus").Exists() {
					itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2plus", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2plus."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.adsl2plus."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "operating")
		if !item.ModeVdsl2.IsNull() {
			if item.ModeVdsl2.ValueBool() {
				if !gjson.Get(itemBody, "operating.mode.auto.vdsl2").Exists() {
					itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.vdsl2", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.vdsl2."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.vdsl2."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "operating")
		if !item.ModeAnsi.IsNull() {
			if item.ModeAnsi.ValueBool() {
				if !gjson.Get(itemBody, "operating.mode.auto.ansi").Exists() {
					itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.ansi", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.ansi."+"vipObjectType", "node-only")
				itemBody, _ = sjson.Set(itemBody, "operating.mode.auto.ansi."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "modem")

		if !item.VdslModemConfigurationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipVariableName", item.VdslModemConfigurationVariable.ValueString())
		} else if item.VdslModemConfiguration.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "modem."+"vipValue", item.VdslModemConfiguration.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"controller.vdsl."+"vipValue.-1", itemBody)
	}

	if !data.EthernetVlanIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipVariableName", data.EthernetVlanIdVariable.ValueString())
	} else if data.EthernetVlanId.IsNull() {
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bridge-dot1q.encap."+"vipValue", data.EthernetVlanId.ValueInt64())
	}

	if !data.EthernetDialerPoolNumberVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipVariableName", data.EthernetDialerPoolNumberVariable.ValueString())
	} else if data.EthernetDialerPoolNumber.IsNull() {
		if !gjson.Get(body, path+"pppoe-client").Exists() {
			body, _ = sjson.Set(body, path+"pppoe-client", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pppoe-client.dial-pool-number."+"vipValue", data.EthernetDialerPoolNumber.ValueInt64())
	}

	if !data.EthernetPppMaximumPayloadVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipVariableName", data.EthernetPppMaximumPayloadVariable.ValueString())
	} else if data.EthernetPppMaximumPayload.IsNull() {
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pppoe-client.ppp-max-payload."+"vipValue", data.EthernetPppMaximumPayload.ValueInt64())
	}
	if !data.EthernetDialerAddressNegotiated.IsNull() {
		if data.EthernetDialerAddressNegotiated.ValueBool() {
			if !gjson.Get(body, path+"ip.address-negotiated").Exists() {
				body, _ = sjson.Set(body, path+"ip.address-negotiated", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ip.address-negotiated."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ip.address-negotiated."+"vipType", "ignore")
		}
	}

	if !data.EthernetUnnumberedLoopbackInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipVariableName", data.EthernetUnnumberedLoopbackInterfaceVariable.ValueString())
	} else if data.EthernetUnnumberedLoopbackInterface.IsNull() {
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.unnumbered."+"vipValue", data.EthernetUnnumberedLoopbackInterface.ValueString())
	}

	if !data.PppAuthenticationProtocolVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ppp.authentication.method."+"vipVariableName", data.PppAuthenticationProtocolVariable.ValueString())
	} else if data.PppAuthenticationProtocol.IsNull() {
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
	if !data.Nat.IsNull() {
		if data.Nat.ValueBool() {
			if !gjson.Get(body, path+"nat").Exists() {
				body, _ = sjson.Set(body, path+"nat", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"nat."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"nat."+"vipType", "ignore")
		}
	}

	if !data.NatRefreshModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipVariableName", data.NatRefreshModeVariable.ValueString())
	} else if data.NatRefreshMode.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.refresh."+"vipValue", data.NatRefreshMode.ValueString())
	}

	if !data.NatUdpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipVariableName", data.NatUdpTimeoutVariable.ValueString())
	} else if data.NatUdpTimeout.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipValue", data.NatUdpTimeout.ValueInt64())
	}

	if !data.NatTcpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipVariableName", data.NatTcpTimeoutVariable.ValueString())
	} else if data.NatTcpTimeout.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipValue", data.NatTcpTimeout.ValueInt64())
	}

	if !data.NatBlockIcmpErrorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipVariableName", data.NatBlockIcmpErrorVariable.ValueString())
	} else if data.NatBlockIcmpError.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.block-icmp-error."+"vipValue", strconv.FormatBool(data.NatBlockIcmpError.ValueBool()))
	}

	if !data.NatResponseToPingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipVariableName", data.NatResponseToPingVariable.ValueString())
	} else if data.NatResponseToPing.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.respond-to-ping."+"vipValue", strconv.FormatBool(data.NatResponseToPing.ValueBool()))
	}
	if len(data.NatPortForwards) > 0 {
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipPrimaryKey", []string{"port-start", "port-end", "proto"})
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.NatPortForwards {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "port-start")
		if item.PortStartRange.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "port-start."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "port-start."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "port-start."+"vipValue", item.PortStartRange.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "port-end")
		if item.PortEndRange.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "port-end."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "port-end."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "port-end."+"vipValue", item.PortEndRange.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "proto")
		if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "private-vpn")

		if !item.PrivateVpnVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipVariableName", item.PrivateVpnVariable.ValueString())
		} else if item.PrivateVpn.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "private-vpn."+"vipValue", item.PrivateVpn.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "private-ip-address")

		if !item.PrivateIpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipVariableName", item.PrivateIpAddressVariable.ValueString())
		} else if item.PrivateIpAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "private-ip-address."+"vipValue", item.PrivateIpAddress.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.port-forward."+"vipValue.-1", itemBody)
	}

	if !data.QosAdaptivePeriodVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipVariableName", data.QosAdaptivePeriodVariable.ValueString())
	} else if data.QosAdaptivePeriod.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipValue", data.QosAdaptivePeriod.ValueInt64())
	}

	if !data.QosAdaptiveBandwidthDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipVariableName", data.QosAdaptiveBandwidthDownstreamVariable.ValueString())
	} else if data.QosAdaptiveBandwidthDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipValue", data.QosAdaptiveBandwidthDownstream.ValueInt64())
	}

	if !data.QosAdaptiveMinDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipVariableName", data.QosAdaptiveMinDownstreamVariable.ValueString())
	} else if data.QosAdaptiveMinDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipValue", data.QosAdaptiveMinDownstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipVariableName", data.QosAdaptiveMaxDownstreamVariable.ValueString())
	} else if data.QosAdaptiveMaxDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipValue", data.QosAdaptiveMaxDownstream.ValueInt64())
	}

	if !data.QosAdaptiveBandwidthUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipVariableName", data.QosAdaptiveBandwidthUpstreamVariable.ValueString())
	} else if data.QosAdaptiveBandwidthUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipValue", data.QosAdaptiveBandwidthUpstream.ValueInt64())
	}

	if !data.QosAdaptiveMinUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipVariableName", data.QosAdaptiveMinUpstreamVariable.ValueString())
	} else if data.QosAdaptiveMinUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipValue", data.QosAdaptiveMinUpstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipVariableName", data.QosAdaptiveMaxUpstreamVariable.ValueString())
	} else if data.QosAdaptiveMaxUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipValue", data.QosAdaptiveMaxUpstream.ValueInt64())
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

	if !data.VpnQosMapVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipVariableName", data.VpnQosMapVariable.ValueString())
	} else if data.VpnQosMap.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipValue", data.VpnQosMap.ValueString())
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
	if len(data.Policers) > 0 {
		body, _ = sjson.Set(body, path+"policer."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"policer."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"policer."+"vipPrimaryKey", []string{"policer-name", "direction"})
		body, _ = sjson.Set(body, path+"policer."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"policer."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"policer."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"policer."+"vipPrimaryKey", []string{"policer-name", "direction"})
		body, _ = sjson.Set(body, path+"policer."+"vipValue", []interface{}{})
	}
	for _, item := range data.Policers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "policer-name")
		if item.PolicerName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipValue", item.PolicerName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"policer."+"vipValue.-1", itemBody)
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

	if !data.TrackerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tracker."+"vipVariableName", data.TrackerVariable.ValueString())
	} else if data.Tracker.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "constant")
		var values []string
		data.Tracker.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", values)
	}

	if !data.IpDirectedBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipVariableName", data.IpDirectedBroadcastVariable.ValueString())
	} else if data.IpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipValue", strconv.FormatBool(data.IpDirectedBroadcast.ValueBool()))
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *VPNInterfaceDSLPPPoE) fromBody(ctx context.Context, res gjson.Result) {
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
			data.EthernetInterfaceName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.EthernetInterfaceNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetInterfaceName = types.StringNull()
			data.EthernetInterfaceNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.EthernetInterfaceName = types.StringValue(v.String())
			data.EthernetInterfaceNameVariable = types.StringNull()
		}
	} else {
		data.EthernetInterfaceName = types.StringNull()
		data.EthernetInterfaceNameVariable = types.StringNull()
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
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.EthernetDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetDescription = types.StringNull()
			data.EthernetDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.EthernetDescription = types.StringValue(v.String())
			data.EthernetDescriptionVariable = types.StringNull()
		}
	} else {
		data.EthernetDescription = types.StringNull()
		data.EthernetDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "controller.vdsl.vipValue"); len(value.Array()) > 0 {
		data.VdslConfiguration = make([]VPNInterfaceDSLPPPoEVdslConfiguration, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceDSLPPPoEVdslConfiguration{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ControllerVdslSlot = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.ControllerVdslSlotVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ControllerVdslSlot = types.StringNull()
					item.ControllerVdslSlotVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.ControllerVdslSlot = types.StringValue(cv.String())
					item.ControllerVdslSlotVariable = types.StringNull()
				}
			} else {
				item.ControllerVdslSlot = types.StringNull()
				item.ControllerVdslSlotVariable = types.StringNull()
			}
			if cValue := v.Get("sra.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Sra = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Sra = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("sra.vipValue")
					item.Sra = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("sra"); cValue.Exists() {
				item.Sra = types.BoolValue(true)

			} else {
				item.Sra = types.BoolNull()

			}
			if cValue := v.Get("operating.mode.auto.adsl1.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeAdsl1 = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ModeAdsl1 = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("operating.mode.auto.adsl1.vipValue")
					item.ModeAdsl1 = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("operating.mode.auto.adsl1"); cValue.Exists() {
				item.ModeAdsl1 = types.BoolValue(true)

			} else {
				item.ModeAdsl1 = types.BoolNull()

			}
			if cValue := v.Get("operating.mode.auto.adsl2.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeAdsl2 = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ModeAdsl2 = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("operating.mode.auto.adsl2.vipValue")
					item.ModeAdsl2 = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("operating.mode.auto.adsl2"); cValue.Exists() {
				item.ModeAdsl2 = types.BoolValue(true)

			} else {
				item.ModeAdsl2 = types.BoolNull()

			}
			if cValue := v.Get("operating.mode.auto.adsl2plus.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeAdsl2plus = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ModeAdsl2plus = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("operating.mode.auto.adsl2plus.vipValue")
					item.ModeAdsl2plus = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("operating.mode.auto.adsl2plus"); cValue.Exists() {
				item.ModeAdsl2plus = types.BoolValue(true)

			} else {
				item.ModeAdsl2plus = types.BoolNull()

			}
			if cValue := v.Get("operating.mode.auto.vdsl2.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeVdsl2 = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ModeVdsl2 = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("operating.mode.auto.vdsl2.vipValue")
					item.ModeVdsl2 = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("operating.mode.auto.vdsl2"); cValue.Exists() {
				item.ModeVdsl2 = types.BoolValue(true)

			} else {
				item.ModeVdsl2 = types.BoolNull()

			}
			if cValue := v.Get("operating.mode.auto.ansi.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeAnsi = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ModeAnsi = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("operating.mode.auto.ansi.vipValue")
					item.ModeAnsi = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("operating.mode.auto.ansi"); cValue.Exists() {
				item.ModeAnsi = types.BoolValue(true)

			} else {
				item.ModeAnsi = types.BoolNull()

			}
			if cValue := v.Get("modem.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VdslModemConfiguration = types.StringNull()

					cv := v.Get("modem.vipVariableName")
					item.VdslModemConfigurationVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VdslModemConfiguration = types.StringNull()
					item.VdslModemConfigurationVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("modem.vipValue")
					item.VdslModemConfiguration = types.StringValue(cv.String())
					item.VdslModemConfigurationVariable = types.StringNull()
				}
			} else {
				item.VdslModemConfiguration = types.StringNull()
				item.VdslModemConfigurationVariable = types.StringNull()
			}
			data.VdslConfiguration = append(data.VdslConfiguration, item)
			return true
		})
	} else {
		if len(data.VdslConfiguration) > 0 {
			data.VdslConfiguration = []VPNInterfaceDSLPPPoEVdslConfiguration{}
		}
	}
	if value := res.Get(path + "bridge-dot1q.encap.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetVlanId = types.Int64Null()

			v := res.Get(path + "bridge-dot1q.encap.vipVariableName")
			data.EthernetVlanIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetVlanId = types.Int64Null()
			data.EthernetVlanIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bridge-dot1q.encap.vipValue")
			data.EthernetVlanId = types.Int64Value(v.Int())
			data.EthernetVlanIdVariable = types.StringNull()
		}
	} else {
		data.EthernetVlanId = types.Int64Null()
		data.EthernetVlanIdVariable = types.StringNull()
	}
	if value := res.Get(path + "pppoe-client.dial-pool-number.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetDialerPoolNumber = types.Int64Null()

			v := res.Get(path + "pppoe-client.dial-pool-number.vipVariableName")
			data.EthernetDialerPoolNumberVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetDialerPoolNumber = types.Int64Null()
			data.EthernetDialerPoolNumberVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pppoe-client.dial-pool-number.vipValue")
			data.EthernetDialerPoolNumber = types.Int64Value(v.Int())
			data.EthernetDialerPoolNumberVariable = types.StringNull()
		}
	} else {
		data.EthernetDialerPoolNumber = types.Int64Null()
		data.EthernetDialerPoolNumberVariable = types.StringNull()
	}
	if value := res.Get(path + "pppoe-client.ppp-max-payload.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetPppMaximumPayload = types.Int64Null()

			v := res.Get(path + "pppoe-client.ppp-max-payload.vipVariableName")
			data.EthernetPppMaximumPayloadVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetPppMaximumPayload = types.Int64Null()
			data.EthernetPppMaximumPayloadVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pppoe-client.ppp-max-payload.vipValue")
			data.EthernetPppMaximumPayload = types.Int64Value(v.Int())
			data.EthernetPppMaximumPayloadVariable = types.StringNull()
		}
	} else {
		data.EthernetPppMaximumPayload = types.Int64Null()
		data.EthernetPppMaximumPayloadVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.address-negotiated.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetDialerAddressNegotiated = types.BoolNull()

		} else if value.String() == "ignore" {
			data.EthernetDialerAddressNegotiated = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address-negotiated.vipValue")
			data.EthernetDialerAddressNegotiated = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ip.address-negotiated"); value.Exists() {
		data.EthernetDialerAddressNegotiated = types.BoolValue(true)

	} else {
		data.EthernetDialerAddressNegotiated = types.BoolNull()

	}
	if value := res.Get(path + "ip.unnumbered.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EthernetUnnumberedLoopbackInterface = types.StringNull()

			v := res.Get(path + "ip.unnumbered.vipVariableName")
			data.EthernetUnnumberedLoopbackInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EthernetUnnumberedLoopbackInterface = types.StringNull()
			data.EthernetUnnumberedLoopbackInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.unnumbered.vipValue")
			data.EthernetUnnumberedLoopbackInterface = types.StringValue(v.String())
			data.EthernetUnnumberedLoopbackInterfaceVariable = types.StringNull()
		}
	} else {
		data.EthernetUnnumberedLoopbackInterface = types.StringNull()
		data.EthernetUnnumberedLoopbackInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "ppp.authentication.method.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PppAuthenticationProtocol = types.StringNull()

			v := res.Get(path + "ppp.authentication.method.vipVariableName")
			data.PppAuthenticationProtocolVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PppAuthenticationProtocol = types.StringNull()
			data.PppAuthenticationProtocolVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ppp.authentication.method.vipValue")
			data.PppAuthenticationProtocol = types.StringValue(v.String())
			data.PppAuthenticationProtocolVariable = types.StringNull()
		}
	} else {
		data.PppAuthenticationProtocol = types.StringNull()
		data.PppAuthenticationProtocolVariable = types.StringNull()
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
		data.TunnelInterfaceEncapsulations = make([]VPNInterfaceDSLPPPoETunnelInterfaceEncapsulations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceDSLPPPoETunnelInterfaceEncapsulations{}
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
			data.TunnelInterfaceEncapsulations = []VPNInterfaceDSLPPPoETunnelInterfaceEncapsulations{}
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
	if value := res.Get(path + "nat.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Nat = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "nat.vipValue")
			data.Nat = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "nat"); value.Exists() {
		data.Nat = types.BoolValue(true)

	} else {
		data.Nat = types.BoolNull()

	}
	if value := res.Get(path + "nat.refresh.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatRefreshMode = types.StringNull()

			v := res.Get(path + "nat.refresh.vipVariableName")
			data.NatRefreshModeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatRefreshMode = types.StringNull()
			data.NatRefreshModeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.refresh.vipValue")
			data.NatRefreshMode = types.StringValue(v.String())
			data.NatRefreshModeVariable = types.StringNull()
		}
	} else {
		data.NatRefreshMode = types.StringNull()
		data.NatRefreshModeVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.udp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatUdpTimeout = types.Int64Null()

			v := res.Get(path + "nat.udp-timeout.vipVariableName")
			data.NatUdpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatUdpTimeout = types.Int64Null()
			data.NatUdpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.udp-timeout.vipValue")
			data.NatUdpTimeout = types.Int64Value(v.Int())
			data.NatUdpTimeoutVariable = types.StringNull()
		}
	} else {
		data.NatUdpTimeout = types.Int64Null()
		data.NatUdpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.tcp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatTcpTimeout = types.Int64Null()

			v := res.Get(path + "nat.tcp-timeout.vipVariableName")
			data.NatTcpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatTcpTimeout = types.Int64Null()
			data.NatTcpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.tcp-timeout.vipValue")
			data.NatTcpTimeout = types.Int64Value(v.Int())
			data.NatTcpTimeoutVariable = types.StringNull()
		}
	} else {
		data.NatTcpTimeout = types.Int64Null()
		data.NatTcpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.block-icmp-error.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatBlockIcmpError = types.BoolNull()

			v := res.Get(path + "nat.block-icmp-error.vipVariableName")
			data.NatBlockIcmpErrorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatBlockIcmpError = types.BoolNull()
			data.NatBlockIcmpErrorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.block-icmp-error.vipValue")
			data.NatBlockIcmpError = types.BoolValue(v.Bool())
			data.NatBlockIcmpErrorVariable = types.StringNull()
		}
	} else {
		data.NatBlockIcmpError = types.BoolNull()
		data.NatBlockIcmpErrorVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.respond-to-ping.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatResponseToPing = types.BoolNull()

			v := res.Get(path + "nat.respond-to-ping.vipVariableName")
			data.NatResponseToPingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatResponseToPing = types.BoolNull()
			data.NatResponseToPingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.respond-to-ping.vipValue")
			data.NatResponseToPing = types.BoolValue(v.Bool())
			data.NatResponseToPingVariable = types.StringNull()
		}
	} else {
		data.NatResponseToPing = types.BoolNull()
		data.NatResponseToPingVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.port-forward.vipValue"); len(value.Array()) > 0 {
		data.NatPortForwards = make([]VPNInterfaceDSLPPPoENatPortForwards, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceDSLPPPoENatPortForwards{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("port-start.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PortStartRange = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.PortStartRange = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("port-start.vipValue")
					item.PortStartRange = types.Int64Value(cv.Int())

				}
			} else {
				item.PortStartRange = types.Int64Null()

			}
			if cValue := v.Get("port-end.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PortEndRange = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.PortEndRange = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("port-end.vipValue")
					item.PortEndRange = types.Int64Value(cv.Int())

				}
			} else {
				item.PortEndRange = types.Int64Null()

			}
			if cValue := v.Get("proto.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("proto.vipValue")
					item.Protocol = types.StringValue(cv.String())

				}
			} else {
				item.Protocol = types.StringNull()

			}
			if cValue := v.Get("private-vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivateVpn = types.Int64Null()

					cv := v.Get("private-vpn.vipVariableName")
					item.PrivateVpnVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrivateVpn = types.Int64Null()
					item.PrivateVpnVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("private-vpn.vipValue")
					item.PrivateVpn = types.Int64Value(cv.Int())
					item.PrivateVpnVariable = types.StringNull()
				}
			} else {
				item.PrivateVpn = types.Int64Null()
				item.PrivateVpnVariable = types.StringNull()
			}
			if cValue := v.Get("private-ip-address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivateIpAddress = types.StringNull()

					cv := v.Get("private-ip-address.vipVariableName")
					item.PrivateIpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrivateIpAddress = types.StringNull()
					item.PrivateIpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("private-ip-address.vipValue")
					item.PrivateIpAddress = types.StringValue(cv.String())
					item.PrivateIpAddressVariable = types.StringNull()
				}
			} else {
				item.PrivateIpAddress = types.StringNull()
				item.PrivateIpAddressVariable = types.StringNull()
			}
			data.NatPortForwards = append(data.NatPortForwards, item)
			return true
		})
	} else {
		if len(data.NatPortForwards) > 0 {
			data.NatPortForwards = []VPNInterfaceDSLPPPoENatPortForwards{}
		}
	}
	if value := res.Get(path + "qos-adaptive.period.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptivePeriod = types.Int64Null()

			v := res.Get(path + "qos-adaptive.period.vipVariableName")
			data.QosAdaptivePeriodVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptivePeriod = types.Int64Null()
			data.QosAdaptivePeriodVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.period.vipValue")
			data.QosAdaptivePeriod = types.Int64Value(v.Int())
			data.QosAdaptivePeriodVariable = types.StringNull()
		}
	} else {
		data.QosAdaptivePeriod = types.Int64Null()
		data.QosAdaptivePeriodVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveBandwidthDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipVariableName")
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveBandwidthDownstream = types.Int64Null()
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipValue")
			data.QosAdaptiveBandwidthDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveBandwidthDownstream = types.Int64Null()
		data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.range.dmin.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMinDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.range.dmin.vipVariableName")
			data.QosAdaptiveMinDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMinDownstream = types.Int64Null()
			data.QosAdaptiveMinDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.range.dmin.vipValue")
			data.QosAdaptiveMinDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveMinDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMinDownstream = types.Int64Null()
		data.QosAdaptiveMinDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.range.dmax.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMaxDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.range.dmax.vipVariableName")
			data.QosAdaptiveMaxDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMaxDownstream = types.Int64Null()
			data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.range.dmax.vipValue")
			data.QosAdaptiveMaxDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMaxDownstream = types.Int64Null()
		data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveBandwidthUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipVariableName")
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveBandwidthUpstream = types.Int64Null()
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipValue")
			data.QosAdaptiveBandwidthUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveBandwidthUpstream = types.Int64Null()
		data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.range.umin.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMinUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.range.umin.vipVariableName")
			data.QosAdaptiveMinUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMinUpstream = types.Int64Null()
			data.QosAdaptiveMinUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.range.umin.vipValue")
			data.QosAdaptiveMinUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveMinUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMinUpstream = types.Int64Null()
		data.QosAdaptiveMinUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.range.umax.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMaxUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.range.umax.vipVariableName")
			data.QosAdaptiveMaxUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMaxUpstream = types.Int64Null()
			data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.range.umax.vipValue")
			data.QosAdaptiveMaxUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMaxUpstream = types.Int64Null()
		data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
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
			data.VpnQosMap = types.StringNull()

			v := res.Get(path + "qos-map-vpn.vipVariableName")
			data.VpnQosMapVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.VpnQosMap = types.StringNull()
			data.VpnQosMapVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-map-vpn.vipValue")
			data.VpnQosMap = types.StringValue(v.String())
			data.VpnQosMapVariable = types.StringNull()
		}
	} else {
		data.VpnQosMap = types.StringNull()
		data.VpnQosMapVariable = types.StringNull()
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
		data.AccessLists = make([]VPNInterfaceDSLPPPoEAccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceDSLPPPoEAccessLists{}
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
			data.AccessLists = []VPNInterfaceDSLPPPoEAccessLists{}
		}
	}
	if value := res.Get(path + "policer.vipValue"); len(value.Array()) > 0 {
		data.Policers = make([]VPNInterfaceDSLPPPoEPolicers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceDSLPPPoEPolicers{}
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
			if cValue := v.Get("policer-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PolicerName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PolicerName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("policer-name.vipValue")
					item.PolicerName = types.StringValue(cv.String())

				}
			} else {
				item.PolicerName = types.StringNull()

			}
			data.Policers = append(data.Policers, item)
			return true
		})
	} else {
		if len(data.Policers) > 0 {
			data.Policers = []VPNInterfaceDSLPPPoEPolicers{}
		}
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
	if value := res.Get(path + "tracker.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.Tracker = types.SetNull(types.StringType)

			v := res.Get(path + "tracker.vipVariableName")
			data.TrackerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Tracker = types.SetNull(types.StringType)
			data.TrackerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tracker.vipValue")
			data.Tracker = helpers.GetStringSet(v.Array())
			data.TrackerVariable = types.StringNull()
		}
	} else {
		data.Tracker = types.SetNull(types.StringType)
		data.TrackerVariable = types.StringNull()
	}
	if value := res.Get(path + "ip-directed-broadcast.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpDirectedBroadcast = types.BoolNull()

			v := res.Get(path + "ip-directed-broadcast.vipVariableName")
			data.IpDirectedBroadcastVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpDirectedBroadcast = types.BoolNull()
			data.IpDirectedBroadcastVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip-directed-broadcast.vipValue")
			data.IpDirectedBroadcast = types.BoolValue(v.Bool())
			data.IpDirectedBroadcastVariable = types.StringNull()
		}
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
		data.IpDirectedBroadcastVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *VPNInterfaceDSLPPPoE) hasChanges(ctx context.Context, state *VPNInterfaceDSLPPPoE) bool {
	hasChanges := false
	if !data.EthernetInterfaceName.Equal(state.EthernetInterfaceName) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.EthernetDescription.Equal(state.EthernetDescription) {
		hasChanges = true
	}
	if len(data.VdslConfiguration) != len(state.VdslConfiguration) {
		hasChanges = true
	} else {
		for i := range data.VdslConfiguration {
			if !data.VdslConfiguration[i].ControllerVdslSlot.Equal(state.VdslConfiguration[i].ControllerVdslSlot) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].Sra.Equal(state.VdslConfiguration[i].Sra) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].ModeAdsl1.Equal(state.VdslConfiguration[i].ModeAdsl1) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].ModeAdsl2.Equal(state.VdslConfiguration[i].ModeAdsl2) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].ModeAdsl2plus.Equal(state.VdslConfiguration[i].ModeAdsl2plus) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].ModeVdsl2.Equal(state.VdslConfiguration[i].ModeVdsl2) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].ModeAnsi.Equal(state.VdslConfiguration[i].ModeAnsi) {
				hasChanges = true
			}
			if !data.VdslConfiguration[i].VdslModemConfiguration.Equal(state.VdslConfiguration[i].VdslModemConfiguration) {
				hasChanges = true
			}
		}
	}
	if !data.EthernetVlanId.Equal(state.EthernetVlanId) {
		hasChanges = true
	}
	if !data.EthernetDialerPoolNumber.Equal(state.EthernetDialerPoolNumber) {
		hasChanges = true
	}
	if !data.EthernetPppMaximumPayload.Equal(state.EthernetPppMaximumPayload) {
		hasChanges = true
	}
	if !data.EthernetDialerAddressNegotiated.Equal(state.EthernetDialerAddressNegotiated) {
		hasChanges = true
	}
	if !data.EthernetUnnumberedLoopbackInterface.Equal(state.EthernetUnnumberedLoopbackInterface) {
		hasChanges = true
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
	if !data.Nat.Equal(state.Nat) {
		hasChanges = true
	}
	if !data.NatRefreshMode.Equal(state.NatRefreshMode) {
		hasChanges = true
	}
	if !data.NatUdpTimeout.Equal(state.NatUdpTimeout) {
		hasChanges = true
	}
	if !data.NatTcpTimeout.Equal(state.NatTcpTimeout) {
		hasChanges = true
	}
	if !data.NatBlockIcmpError.Equal(state.NatBlockIcmpError) {
		hasChanges = true
	}
	if !data.NatResponseToPing.Equal(state.NatResponseToPing) {
		hasChanges = true
	}
	if len(data.NatPortForwards) != len(state.NatPortForwards) {
		hasChanges = true
	} else {
		for i := range data.NatPortForwards {
			if !data.NatPortForwards[i].PortStartRange.Equal(state.NatPortForwards[i].PortStartRange) {
				hasChanges = true
			}
			if !data.NatPortForwards[i].PortEndRange.Equal(state.NatPortForwards[i].PortEndRange) {
				hasChanges = true
			}
			if !data.NatPortForwards[i].Protocol.Equal(state.NatPortForwards[i].Protocol) {
				hasChanges = true
			}
			if !data.NatPortForwards[i].PrivateVpn.Equal(state.NatPortForwards[i].PrivateVpn) {
				hasChanges = true
			}
			if !data.NatPortForwards[i].PrivateIpAddress.Equal(state.NatPortForwards[i].PrivateIpAddress) {
				hasChanges = true
			}
		}
	}
	if !data.QosAdaptivePeriod.Equal(state.QosAdaptivePeriod) {
		hasChanges = true
	}
	if !data.QosAdaptiveBandwidthDownstream.Equal(state.QosAdaptiveBandwidthDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMinDownstream.Equal(state.QosAdaptiveMinDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMaxDownstream.Equal(state.QosAdaptiveMaxDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveBandwidthUpstream.Equal(state.QosAdaptiveBandwidthUpstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMinUpstream.Equal(state.QosAdaptiveMinUpstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMaxUpstream.Equal(state.QosAdaptiveMaxUpstream) {
		hasChanges = true
	}
	if !data.ShapingRate.Equal(state.ShapingRate) {
		hasChanges = true
	}
	if !data.QosMap.Equal(state.QosMap) {
		hasChanges = true
	}
	if !data.VpnQosMap.Equal(state.VpnQosMap) {
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
	if len(data.Policers) != len(state.Policers) {
		hasChanges = true
	} else {
		for i := range data.Policers {
			if !data.Policers[i].Direction.Equal(state.Policers[i].Direction) {
				hasChanges = true
			}
			if !data.Policers[i].PolicerName.Equal(state.Policers[i].PolicerName) {
				hasChanges = true
			}
		}
	}
	if !data.IpMtu.Equal(state.IpMtu) {
		hasChanges = true
	}
	if !data.TcpMss.Equal(state.TcpMss) {
		hasChanges = true
	}
	if !data.TlocExtension.Equal(state.TlocExtension) {
		hasChanges = true
	}
	if !data.Tracker.Equal(state.Tracker) {
		hasChanges = true
	}
	if !data.IpDirectedBroadcast.Equal(state.IpDirectedBroadcast) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
