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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanVPNInterfaceEthernetPPPoEFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ethernet_interface_name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "interface_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "encap", "4094"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "dialer_pool_number", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ppp_maximum_payload", "1790"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "dialer_address_negotiated", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ip_directed_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "unnumbered_loopback_interface", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ppp_authentication_protocol", "chap"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ppp_authentication_protocol_pap", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "authentication_type", "callin"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "chap_hostname", "chap-example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "chap_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "pap_username", "pap-username"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "pap_password", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "pap_outbound_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "enable_core_region", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "core_region", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "secondary_region", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "encapsulation.0.encapsulation_type", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "encapsulation.0.preference", "4294967"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "encapsulation.0.weight", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "border", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "per_tunnel_qos_aggregator", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "color", "custom1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "network_broadcast_1", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "max_control_connections", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "control_connections", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "vmanage_connection_preference", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_hop", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "carrier", "carrier1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "nat_refresh_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "bind_loopback_tunnel", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "network_broadcast_2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "dns", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "icmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "https", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "refresh_mode", "outbound"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "udp_timeout", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "tcp_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "block_icmp_error", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "respond_to_ping", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_forward.0.port_start_range", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_forward.0.port_end_range", "65530"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_forward.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_forward.0.private_vpn", "65530"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "port_forward.0.private_ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "adaptive_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "adapt_period", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_downstream_default", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_downstream_min", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_downstream_max", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_upstream_default", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_upstream_min", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate_upstream_max", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "shaping_rate", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "vpn_qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "bandwidth_upstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "bandwidth_downstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "write_rule", "test_rule"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "access_list.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "access_list.0.acl_name", "Egress ACL - IPv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "policer.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "policer.0.policer_name", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "ip_mtu_for_dialer_interface", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "dialer_tcp_mss", "720"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_ethernet_pppoe_feature_template.test", "tloc_extension", "tloc"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceEthernetPPPoEFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanVPNInterfaceEthernetPPPoEFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccSdwanVPNInterfaceEthernetPPPoEFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_vpn_interface_ethernet_pppoe_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	ethernet_interface_name = "Example"` + "\n"
	config += `	dialer_pool_number = 255` + "\n"
	config += `	ppp_authentication_protocol = "chap"` + "\n"
	config += `	authentication_type = "callin"` + "\n"
	config += `	chap_hostname = "chap-example"` + "\n"
	config += `	chap_password = "myPassword"` + "\n"
	config += `	pap_username = "pap-username"` + "\n"
	config += `	pap_password = true` + "\n"
	config += `	pap_outbound_password = "myPassword"` + "\n"
	config += `	per_tunnel_qos = true` + "\n"
	config += `	per_tunnel_qos_aggregator = false` + "\n"
	config += `	shaping_rate_downstream_default = 10000000` + "\n"
	config += `	shaping_rate_downstream_min = 10000000` + "\n"
	config += `	shaping_rate_downstream_max = 10000000` + "\n"
	config += `	shaping_rate_upstream_default = 10000000` + "\n"
	config += `	shaping_rate_upstream_min = 10000000` + "\n"
	config += `	shaping_rate_upstream_max = 10000000` + "\n"
	config += `}` + "\n"
	return config
}

func testAccSdwanVPNInterfaceEthernetPPPoEFeatureTemplateConfig_all() string {
	config := `resource "sdwan_vpn_interface_ethernet_pppoe_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	ethernet_interface_name = "Example"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	interface_description = "My Description"` + "\n"
	config += `	encap = 4094` + "\n"
	config += `	dialer_pool_number = 255` + "\n"
	config += `	ppp_maximum_payload = 1790` + "\n"
	config += `	dialer_address_negotiated = false` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `	unnumbered_loopback_interface = "example"` + "\n"
	config += `	ppp_authentication_protocol = "chap"` + "\n"
	config += `	ppp_authentication_protocol_pap = false` + "\n"
	config += `	authentication_type = "callin"` + "\n"
	config += `	chap_hostname = "chap-example"` + "\n"
	config += `	chap_password = "myPassword"` + "\n"
	config += `	pap_username = "pap-username"` + "\n"
	config += `	pap_password = true` + "\n"
	config += `	pap_outbound_password = "myPassword"` + "\n"
	config += `	enable_core_region = true` + "\n"
	config += `	core_region = "core"` + "\n"
	config += `	secondary_region = "off"` + "\n"
	config += `	encapsulation = [{` + "\n"
	config += `	  encapsulation_type = "gre"` + "\n"
	config += `	  preference = 4294967` + "\n"
	config += `	  weight = 250` + "\n"
	config += `	}]` + "\n"
	config += `	groups = [42949672]` + "\n"
	config += `	border = true` + "\n"
	config += `	per_tunnel_qos = true` + "\n"
	config += `	per_tunnel_qos_aggregator = false` + "\n"
	config += `	color = "custom1"` + "\n"
	config += `	last_resort_circuit = false` + "\n"
	config += `	low_bandwidth_link = false` + "\n"
	config += `	tunnel_tcp_mss = 1460` + "\n"
	config += `	clear_dont_fragment = false` + "\n"
	config += `	network_broadcast_1 = false` + "\n"
	config += `	max_control_connections = 8` + "\n"
	config += `	control_connections = true` + "\n"
	config += `	vbond_as_stun_server = false` + "\n"
	config += `	exclude_controller_group_list = [100]` + "\n"
	config += `	vmanage_connection_preference = 5` + "\n"
	config += `	port_hop = false` + "\n"
	config += `	restrict = false` + "\n"
	config += `	carrier = "carrier1"` + "\n"
	config += `	nat_refresh_interval = 15` + "\n"
	config += `	hello_interval = 1000` + "\n"
	config += `	hello_tolerance = 12` + "\n"
	config += `	bind_loopback_tunnel = "12"` + "\n"
	config += `	all = false` + "\n"
	config += `	network_broadcast_2 = false` + "\n"
	config += `	bgp = false` + "\n"
	config += `	dhcp = true` + "\n"
	config += `	dns = true` + "\n"
	config += `	icmp = true` + "\n"
	config += `	ssh = false` + "\n"
	config += `	netconf = false` + "\n"
	config += `	ospf = false` + "\n"
	config += `	stun = false` + "\n"
	config += `	snmp = false` + "\n"
	config += `	https = true` + "\n"
	config += `	nat = true` + "\n"
	config += `	refresh_mode = "outbound"` + "\n"
	config += `	udp_timeout = 1` + "\n"
	config += `	tcp_timeout = 60` + "\n"
	config += `	block_icmp_error = true` + "\n"
	config += `	respond_to_ping = false` + "\n"
	config += `	port_forward = [{` + "\n"
	config += `	  port_start_range = 0` + "\n"
	config += `	  port_end_range = 65530` + "\n"
	config += `	  protocol = "tcp"` + "\n"
	config += `	  private_vpn = 65530` + "\n"
	config += `	  private_ip_address = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	adaptive_qos = true` + "\n"
	config += `	adapt_period = 15` + "\n"
	config += `	shaping_rate_downstream_default = 10000000` + "\n"
	config += `	shaping_rate_downstream_min = 10000000` + "\n"
	config += `	shaping_rate_downstream_max = 10000000` + "\n"
	config += `	shaping_rate_upstream_default = 10000000` + "\n"
	config += `	shaping_rate_upstream_min = 10000000` + "\n"
	config += `	shaping_rate_upstream_max = 10000000` + "\n"
	config += `	shaping_rate = 10000000` + "\n"
	config += `	qos_map = "test"` + "\n"
	config += `	vpn_qos_map = "test"` + "\n"
	config += `	bandwidth_upstream = 214748300` + "\n"
	config += `	bandwidth_downstream = 214748300` + "\n"
	config += `	write_rule = "test_rule"` + "\n"
	config += `	access_list = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "Egress ACL - IPv4"` + "\n"
	config += `	}]` + "\n"
	config += `	policer = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  policer_name = "example"` + "\n"
	config += `	}]` + "\n"
	config += `	ip_mtu_for_dialer_interface = 1500` + "\n"
	config += `	dialer_tcp_mss = 720` + "\n"
	config += `	tloc_extension = "tloc"` + "\n"
	config += `	tracker = ["tracker1"]` + "\n"
	config += `}` + "\n"
	return config
}
