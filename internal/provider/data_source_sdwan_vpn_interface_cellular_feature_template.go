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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VPNInterfaceCellularFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNInterfaceCellularFeatureTemplateDataSource{}
)

func NewVPNInterfaceCellularFeatureTemplateDataSource() datasource.DataSource {
	return &VPNInterfaceCellularFeatureTemplateDataSource{}
}

type VPNInterfaceCellularFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *VPNInterfaceCellularFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_interface_cellular_feature_template"
}

func (d *VPNInterfaceCellularFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VPN Interface Cellular feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"cellular_interface_name": schema.StringAttribute{
				MarkdownDescription: "Cellular Interface Name <0-1>",
				Computed:            true,
			},
			"cellular_interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "Interface description",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply ACL",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: "List of DHCP server addresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tracker": schema.SetAttribute{
				MarkdownDescription: "Enable tracker for this interface",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tracker_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat": schema.BoolAttribute{
				MarkdownDescription: "Network Address Translation on this interface",
				Computed:            true,
			},
			"nat_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"refresh_mode": schema.StringAttribute{
				MarkdownDescription: "Set NAT refresh mode",
				Computed:            true,
			},
			"refresh_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT UDP session timeout, in minutes",
				Computed:            true,
			},
			"udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT TCP session timeout, in minutes",
				Computed:            true,
			},
			"tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"block_icmp_error": schema.BoolAttribute{
				MarkdownDescription: "Block inbound ICMP error messages",
				Computed:            true,
			},
			"block_icmp_error_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"respond_to_ping": schema.BoolAttribute{
				MarkdownDescription: "Respond to ping requests to NAT interface ip address from the public side",
				Computed:            true,
			},
			"respond_to_ping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"port_forward": schema.ListNestedAttribute{
				MarkdownDescription: "Set port-forwarding rules for NAT on this interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_start_range": schema.Int64Attribute{
							MarkdownDescription: "Starting port of port range",
							Computed:            true,
						},
						"port_end_range": schema.Int64Attribute{
							MarkdownDescription: "Ending port of port range",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Layer 4 protocol to apply port forwarding to",
							Computed:            true,
						},
						"private_vpn": schema.Int64Attribute{
							MarkdownDescription: "VPN in which private IP Address resides",
							Computed:            true,
						},
						"private_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"private_ip_address": schema.StringAttribute{
							MarkdownDescription: "Private IP Address to translate to",
							Computed:            true,
						},
						"private_ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"enable_core_region": schema.BoolAttribute{
				MarkdownDescription: "Enable core region",
				Computed:            true,
			},
			"enable_core_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"core_region": schema.StringAttribute{
				MarkdownDescription: "Enable core region",
				Computed:            true,
			},
			"core_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_region": schema.StringAttribute{
				MarkdownDescription: "Enable secondary region",
				Computed:            true,
			},
			"secondary_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"encapsulation": schema.ListNestedAttribute{
				MarkdownDescription: "Encapsulation for TLOC",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encapsulation_type": schema.StringAttribute{
							MarkdownDescription: "Encapsulation",
							Computed:            true,
						},
						"preference": schema.Int64Attribute{
							MarkdownDescription: "Set preference for TLOC",
							Computed:            true,
						},
						"preference_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"weight": schema.Int64Attribute{
							MarkdownDescription: "Set weight for TLOC",
							Computed:            true,
						},
						"weight_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"groups": schema.SetAttribute{
				MarkdownDescription: "List of groups",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"border": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as border TLOC",
				Computed:            true,
			},
			"border_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"per_tunnel_qos": schema.BoolAttribute{
				MarkdownDescription: "Per-tunnel Qos",
				Computed:            true,
			},
			"per_tunnel_qos_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"per_tunnel_qos_aggregator": schema.BoolAttribute{
				MarkdownDescription: "Per-tunnel QoS Aggregator",
				Computed:            true,
			},
			"per_tunnel_qos_aggregator_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"color": schema.StringAttribute{
				MarkdownDescription: "Set color for TLOC",
				Computed:            true,
			},
			"color_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as last resort",
				Computed:            true,
			},
			"last_resort_circuit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"low_bandwidth_link": schema.BoolAttribute{
				MarkdownDescription: "Set the interface as a low-bandwidth circuit",
				Computed:            true,
			},
			"low_bandwidth_link_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "Tunnel TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tunnel_tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: "Clear don't fragment bit",
				Computed:            true,
			},
			"enable_clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"network_broadcast_1": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts)",
				Computed:            true,
			},
			"network_broadcast_1_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"max_control_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of control connections for this TLOC",
				Computed:            true,
			},
			"max_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"control_connections": schema.BoolAttribute{
				MarkdownDescription: "Allow Control Connection",
				Computed:            true,
			},
			"control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: "Put this wan interface in STUN mode only",
				Computed:            true,
			},
			"vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: "Exclude the following controller groups defined in this list",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"exclude_controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vmanage_connection_preference": schema.Int64Attribute{
				MarkdownDescription: "Set interface preference for control connection to vManage <0..8>",
				Computed:            true,
			},
			"vmanage_connection_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"port_hop": schema.BoolAttribute{
				MarkdownDescription: "Disallow port hopping on the tunnel interface",
				Computed:            true,
			},
			"port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"restrict": schema.BoolAttribute{
				MarkdownDescription: "Restrict this TLOC behavior",
				Computed:            true,
			},
			"restrict_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"carrier": schema.StringAttribute{
				MarkdownDescription: "Set carrier for TLOC",
				Computed:            true,
			},
			"carrier_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_refresh_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of nat refresh packets <1...60> seconds",
				Computed:            true,
			},
			"nat_refresh_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of control hello packets <100..600000> milli seconds",
				Computed:            true,
			},
			"hello_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hello_tolerance": schema.Int64Attribute{
				MarkdownDescription: "Set tolerance of control hello packets <12..6000> seconds",
				Computed:            true,
			},
			"hello_tolerance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bind_loopback_tunnel": schema.StringAttribute{
				MarkdownDescription: "Bind loopback tunnel interface to a physical interface",
				Computed:            true,
			},
			"bind_loopback_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"all": schema.BoolAttribute{
				MarkdownDescription: "Allow all traffic. Overrides all other allow-service options if allow-service all is set",
				Computed:            true,
			},
			"all_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"network_broadcast_2": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts)",
				Computed:            true,
			},
			"network_broadcast_2_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bgp": schema.BoolAttribute{
				MarkdownDescription: "Allow/deny BGP",
				Computed:            true,
			},
			"bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dhcp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DHCP",
				Computed:            true,
			},
			"dhcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dns": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DNS",
				Computed:            true,
			},
			"dns_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"icmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny ICMP",
				Computed:            true,
			},
			"icmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ssh": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SSH",
				Computed:            true,
			},
			"ssh_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ntp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NTP",
				Computed:            true,
			},
			"ntp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"netconf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NETCONF",
				Computed:            true,
			},
			"netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ospf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny OSPF",
				Computed:            true,
			},
			"ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"stun": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny STUN",
				Computed:            true,
			},
			"stun_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"snmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SNMP",
				Computed:            true,
			},
			"snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"https": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny Https",
				Computed:            true,
			},
			"https_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"clear_dont_fragment_bit": schema.BoolAttribute{
				MarkdownDescription: "Clear don't fragment bit",
				Computed:            true,
			},
			"clear_dont_fragment_bit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"pmtu_discovery": schema.BoolAttribute{
				MarkdownDescription: "Path MTU Discovery",
				Computed:            true,
			},
			"pmtu_discovery_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU <68...1500>, in bytes",
				Computed:            true,
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"static_ingress_qos": schema.Int64Attribute{
				MarkdownDescription: "Static ingress QoS for the port",
				Computed:            true,
			},
			"static_ingress_qos_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tloc_extension": schema.StringAttribute{
				MarkdownDescription: "Extends a local TLOC to a remote node only for vpn 0",
				Computed:            true,
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: "IP Directed-Broadcast",
				Computed:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"administrative_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"administrative_shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"link_autonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Link autonegotiation",
				Computed:            true,
			},
			"link_autonegotiate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"adaptive_qos": schema.BoolAttribute{
				MarkdownDescription: "Adaptive QoS",
				Computed:            true,
			},
			"adapt_period": schema.Int64Attribute{
				MarkdownDescription: "Periodic timer for adaptive QoS in minutes",
				Computed:            true,
			},
			"adapt_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"adaptive_qos_downstream_default": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default downstream bandwidth",
				Computed:            true,
			},
			"adaptive_qos_downstream_default_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"downstream_min_bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: "Downstream min bandwidth limit",
				Computed:            true,
			},
			"downstream_min_bandwidth_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"downstream_max_bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: "Downstream max bandwidth limit",
				Computed:            true,
			},
			"downstream_max_bandwidth_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"adaptive_qos_upstream_default": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default upstream bandwidth",
				Computed:            true,
			},
			"adaptive_qos_upstream_default_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"upstream_min_bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: "Upstream min bandwidth limit",
				Computed:            true,
			},
			"upstream_min_bandwidth_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"upstream_max_bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: "Upstream max bandwidth limit",
				Computed:            true,
			},
			"upstream_max_bandwidth_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shaping_rate": schema.Int64Attribute{
				MarkdownDescription: "1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps",
				Computed:            true,
			},
			"shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_map": schema.StringAttribute{
				MarkdownDescription: "Name of QoS map",
				Computed:            true,
			},
			"qos_map_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vpn_qos_map": schema.StringAttribute{
				MarkdownDescription: "Name of VPN QoS map",
				Computed:            true,
			},
			"vpn_qos_map_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bandwidth_upstream": schema.Int64Attribute{
				MarkdownDescription: "Interface upstream bandwidth capacity, in kbps",
				Computed:            true,
			},
			"bandwidth_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bandwidth_downstream": schema.Int64Attribute{
				MarkdownDescription: "Interface downstream bandwidth capacity, in kbps",
				Computed:            true,
			},
			"bandwidth_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"write_rule": schema.StringAttribute{
				MarkdownDescription: "Name of rewrite rule",
				Computed:            true,
			},
			"write_rule_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"access_list": schema.ListNestedAttribute{
				MarkdownDescription: "Apply ACL",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"policer": schema.ListNestedAttribute{
				MarkdownDescription: "Enable policer",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"policer_name": schema.StringAttribute{
							MarkdownDescription: "Name of policer",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"static_arp_entries": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static ARP entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "MAC address",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VPNInterfaceCellularFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *VPNInterfaceCellularFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *VPNInterfaceCellularFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNInterfaceCellular

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/template/feature")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing templates, got error: %s", err))
			return
		}
		if value := res.Get("data"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("templateName").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found feature template with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find feature template with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get("/template/feature/object/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
