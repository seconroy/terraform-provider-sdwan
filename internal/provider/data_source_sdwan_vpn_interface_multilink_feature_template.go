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

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VPNInterfaceMultilinkFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNInterfaceMultilinkFeatureTemplateDataSource{}
)

func NewVPNInterfaceMultilinkFeatureTemplateDataSource() datasource.DataSource {
	return &VPNInterfaceMultilinkFeatureTemplateDataSource{}
}

type VPNInterfaceMultilinkFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *VPNInterfaceMultilinkFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_interface_multilink_feature_template"
}

func (d *VPNInterfaceMultilinkFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VPN Interface Multilink feature template.",

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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface Name",
				Computed:            true,
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"multilink_group_number": schema.Int64Attribute{
				MarkdownDescription: "MultiLink Group Number",
				Computed:            true,
			},
			"multilink_group_number_variable": schema.StringAttribute{
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
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv4 address",
				Computed:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv6 address",
				Computed:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply IPv6 access list",
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
			"ppp_authentication_protocol": schema.StringAttribute{
				MarkdownDescription: "PPP Link Authentication Protocol",
				Computed:            true,
			},
			"ppp_authentication_protocol_pap": schema.BoolAttribute{
				MarkdownDescription: "PPP Authentication Protocol PAP",
				Computed:            true,
			},
			"chap_hostname": schema.StringAttribute{
				MarkdownDescription: "CHAP Hostname",
				Computed:            true,
			},
			"chap_hostname_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"chap_ppp_auth_password": schema.StringAttribute{
				MarkdownDescription: "Specify ppp authentication Password",
				Computed:            true,
			},
			"chap_ppp_auth_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"pap_username": schema.StringAttribute{
				MarkdownDescription: "PAP outbound Sent Username",
				Computed:            true,
			},
			"pap_username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"pap_password": schema.BoolAttribute{
				MarkdownDescription: "PAP outbound Password",
				Computed:            true,
			},
			"pap_ppp_auth_password": schema.StringAttribute{
				MarkdownDescription: "Specify ppp authentication Password",
				Computed:            true,
			},
			"pap_ppp_auth_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ppp_authentication_type": schema.StringAttribute{
				MarkdownDescription: "Authenticate remote on incoming call only",
				Computed:            true,
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
			"tunnel_interface_encapsulations": schema.ListNestedAttribute{
				MarkdownDescription: "Encapsulation for TLOC",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encapsulation": schema.StringAttribute{
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
			"tunnel_interface_groups": schema.SetAttribute{
				MarkdownDescription: "List of groups",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"tunnel_interface_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_border": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as border TLOC",
				Computed:            true,
			},
			"tunnel_interface_border_variable": schema.StringAttribute{
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
			"tunnel_qos_mode": schema.StringAttribute{
				MarkdownDescription: "Set tunnel QoS mode",
				Computed:            true,
			},
			"tunnel_qos_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_color": schema.StringAttribute{
				MarkdownDescription: "Set color for TLOC",
				Computed:            true,
			},
			"tunnel_interface_color_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as last resort",
				Computed:            true,
			},
			"tunnel_interface_last_resort_circuit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_low_bandwidth_link": schema.BoolAttribute{
				MarkdownDescription: "Set the interface as a low-bandwidth circuit",
				Computed:            true,
			},
			"tunnel_interface_low_bandwidth_link_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_tunnel_tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "Tunnel TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tunnel_interface_tunnel_tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: "Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)",
				Computed:            true,
			},
			"tunnel_interface_clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts)",
				Computed:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_max_control_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of control connections for this TLOC",
				Computed:            true,
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_control_connections": schema.BoolAttribute{
				MarkdownDescription: "Allow Control Connection",
				Computed:            true,
			},
			"tunnel_interface_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: "Put this wan interface in STUN mode only",
				Computed:            true,
			},
			"tunnel_interface_vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: "Exclude the following controller groups defined in this list",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"tunnel_interface_exclude_controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_vmanage_connection_preference": schema.Int64Attribute{
				MarkdownDescription: "Set interface preference for control connection to vManage <0..8>",
				Computed:            true,
			},
			"tunnel_interface_vmanage_connection_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_port_hop": schema.BoolAttribute{
				MarkdownDescription: "Disallow port hopping on the tunnel interface",
				Computed:            true,
			},
			"tunnel_interface_port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_color_restrict": schema.BoolAttribute{
				MarkdownDescription: "Restrict this TLOC behavior",
				Computed:            true,
			},
			"tunnel_interface_carrier": schema.StringAttribute{
				MarkdownDescription: "Set carrier for TLOC",
				Computed:            true,
			},
			"tunnel_interface_carrier_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_nat_refresh_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of nat refresh packets <1...60> seconds",
				Computed:            true,
			},
			"tunnel_interface_nat_refresh_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of control hello packets <100..600000> milli seconds",
				Computed:            true,
			},
			"tunnel_interface_hello_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_hello_tolerance": schema.Int64Attribute{
				MarkdownDescription: "Set tolerance of control hello packets <12..6000> seconds",
				Computed:            true,
			},
			"tunnel_interface_hello_tolerance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_bind_loopback_tunnel": schema.StringAttribute{
				MarkdownDescription: "Bind loopback tunnel interface to a physical interface",
				Computed:            true,
			},
			"tunnel_interface_bind_loopback_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_all": schema.BoolAttribute{
				MarkdownDescription: "Allow all traffic. Overrides all other allow-service options if allow-service all is set",
				Computed:            true,
			},
			"tunnel_interface_allow_all_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_bgp": schema.BoolAttribute{
				MarkdownDescription: "Allow/deny BGP",
				Computed:            true,
			},
			"tunnel_interface_allow_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DHCP",
				Computed:            true,
			},
			"tunnel_interface_allow_dhcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_dns": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DNS",
				Computed:            true,
			},
			"tunnel_interface_allow_dns_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_icmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny ICMP",
				Computed:            true,
			},
			"tunnel_interface_allow_icmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ssh": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SSH",
				Computed:            true,
			},
			"tunnel_interface_allow_ssh_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NTP",
				Computed:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NETCONF",
				Computed:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ospf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny OSPF",
				Computed:            true,
			},
			"tunnel_interface_allow_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_stun": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny STUN",
				Computed:            true,
			},
			"tunnel_interface_allow_stun_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_snmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SNMP",
				Computed:            true,
			},
			"tunnel_interface_allow_snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_https": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny Https",
				Computed:            true,
			},
			"tunnel_interface_allow_https_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"disable_fragmentation": schema.BoolAttribute{
				MarkdownDescription: "Suppresss multilink fragmentation",
				Computed:            true,
			},
			"fragment_max_delay": schema.Int64Attribute{
				MarkdownDescription: "Maximum delay for each fragment",
				Computed:            true,
			},
			"fragment_max_delay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interleaving_fragment": schema.BoolAttribute{
				MarkdownDescription: "Allow interleaving of packets with fragments",
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
				MarkdownDescription: "Interface MTU <68...2000>, in bytes",
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
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"autonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Link autonegotiation",
				Computed:            true,
			},
			"autonegotiate_variable": schema.StringAttribute{
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
			"qos_map_vpn": schema.StringAttribute{
				MarkdownDescription: "Name of VPN QoS map",
				Computed:            true,
			},
			"qos_map_vpn_variable": schema.StringAttribute{
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
			"access_lists": schema.ListNestedAttribute{
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
			"multilink_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Controller tx-ex List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Card Type",
							Computed:            true,
						},
						"slot": schema.StringAttribute{
							MarkdownDescription: "Slot number",
							Computed:            true,
						},
						"framing": schema.StringAttribute{
							MarkdownDescription: "Framing",
							Computed:            true,
						},
						"framing_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"line_mode": schema.StringAttribute{
							MarkdownDescription: "Line Mode",
							Computed:            true,
						},
						"line_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"internal": schema.BoolAttribute{
							MarkdownDescription: "Internal",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"linecode": schema.StringAttribute{
							MarkdownDescription: "LineCode",
							Computed:            true,
						},
						"linecode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"set_length_for_long": schema.StringAttribute{
							MarkdownDescription: "Set length for long",
							Computed:            true,
						},
						"set_length_for_short": schema.StringAttribute{
							MarkdownDescription: "Set Length for short",
							Computed:            true,
						},
						"channel_group_list": schema.ListNestedAttribute{
							MarkdownDescription: "Channel Group List",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"channel_group": schema.Int64Attribute{
										MarkdownDescription: "Number",
										Computed:            true,
									},
									"channel_group_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"time_slot": schema.SetAttribute{
										MarkdownDescription: "Time slots",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"time_slot_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"nim_interface_list": schema.ListNestedAttribute{
				MarkdownDescription: "Nim Interface List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nim_serial_interface_type": schema.StringAttribute{
							MarkdownDescription: "NIM Serial interface type",
							Computed:            true,
						},
						"nim_serial_interface_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Interface Name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
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
						"bandwidth": schema.Int64Attribute{
							MarkdownDescription: "Interface bandwidth capacity, in kbps",
							Computed:            true,
						},
						"bandwidth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"clock_rate": schema.Int64Attribute{
							MarkdownDescription: "Set preference for interface Clock speed",
							Computed:            true,
						},
						"clock_rate_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"encapsulation_serial": schema.StringAttribute{
							MarkdownDescription: "Configure Encapsulation for interface",
							Computed:            true,
						},
						"encapsulation_serial_variable": schema.StringAttribute{
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

func (d *VPNInterfaceMultilinkFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *VPNInterfaceMultilinkFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *VPNInterfaceMultilinkFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNInterfaceMultilink

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

// End of section. //template:end read
