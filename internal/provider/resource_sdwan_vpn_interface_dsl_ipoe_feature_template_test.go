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

func TestAccSdwanVPNInterfaceDSLIPoEFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_interface_name", "Example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "internal_controller_type", "ipoe"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dialer_pool_number", "255"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_authentication_protocol", "chap"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_hostname", "chap-example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_password", "myPassword"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_username", "pap-username"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_password", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_outbound_password", "myPassword"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos_aggregator", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_default", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_min", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_max", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_default", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_min", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_max", "10000000"),
				),
			},
			{
				Config: testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_interface_name", "Example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ipv4_address", "1.2.3.4/24"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "enable_dhcp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dhcp_distance", "1234"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "internal_controller_type", "ipoe"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "interface_description", "My Description"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.controller_vdsl_slot", "Example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.sra", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.mode_adsl1", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.mode_adsl2", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.mode_adsl2plus", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.mode_vdsl2", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.mode_ansi", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl.0.vdsl_modem_configuration", "100"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "encap", "4094"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dialer_pool_number", "255"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_maximum_payload", "1790"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dialer_address_negotiated", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "unnumbered_loopback_interface", "example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_authentication_protocol", "chap"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_authentication_protocol_pap", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_hostname", "chap-example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_password", "myPassword"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_username", "pap-username"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_password", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_outbound_password", "myPassword"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "encapsulation.0.encapsulation_type", "gre"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "encapsulation.0.preference", "4294967"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "encapsulation.0.weight", "250"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "border", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos_aggregator", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "color", "custom1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "last_resort_circuit", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "low_bandwidth_link", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_tcp_mss", "1460"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "clear_dont_fragment", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "network_broadcast_1", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "max_control_connections", "8"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "control_connections", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vbond_as_stun_server", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vmanage_connection_preference", "5"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_hop", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "restrict", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "carrier", "carrier1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_refresh_interval", "15"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "hello_interval", "1000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "hello_tolerance", "12"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bind_loopback_tunnel", "12"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "all", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "network_broadcast_2", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bgp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dhcp", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dns", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "icmp", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ssh", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "netconf", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ospf", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "stun", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "snmp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "https", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "refresh_mode", "outbound"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "udp_timeout", "1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tcp_timeout", "60"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "block_icmp_error", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "respond_to_ping", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_forward.0.port_start_range", "0"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_forward.0.port_end_range", "65530"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_forward.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_forward.0.private_vpn", "65530"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "port_forward.0.private_ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "adaptive_qos", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "adapt_period", "15"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_default", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_min", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_downstream_max", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_default", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_min", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate_upstream_max", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_map", "test"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vpn_qos_map", "test"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bandwidth_upstream", "214748300"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bandwidth_downstream", "214748300"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "write_rule", "test_rule"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "access_list.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "access_list.0.acl_name", "Egress ACL - IPv4"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "policer.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "policer.0.policer_name", "example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ip_mtu", "1500"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tcp_mss", "720"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tloc_extension", "tloc"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ip_directed_broadcast", "true"),
				),
			},
		},
	})
}

func testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_vpn_interface_dsl_ipoe_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		ethernet_interface_name = "Example"
		internal_controller_type = "ipoe"
		dialer_pool_number = 255
		ppp_authentication_protocol = "chap"
		chap_hostname = "chap-example"
		chap_password = "myPassword"
		pap_username = "pap-username"
		pap_password = true
		pap_outbound_password = "myPassword"
		per_tunnel_qos = true
		per_tunnel_qos_aggregator = false
		shaping_rate_downstream_default = 10000000
		shaping_rate_downstream_min = 10000000
		shaping_rate_downstream_max = 10000000
		shaping_rate_upstream_default = 10000000
		shaping_rate_upstream_min = 10000000
		shaping_rate_upstream_max = 10000000
	}
	`
}

func testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_vpn_interface_dsl_ipoe_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		ethernet_interface_name = "Example"
		ipv4_address = "1.2.3.4/24"
		enable_dhcp = false
		dhcp_distance = 1234
		dhcp_helper = ["3"]
		internal_controller_type = "ipoe"
		shutdown = true
		interface_description = "My Description"
		vdsl = [{
			controller_vdsl_slot = "Example"
			sra = true
			mode_adsl1 = false
			mode_adsl2 = false
			mode_adsl2plus = false
			mode_vdsl2 = false
			mode_ansi = false
			vdsl_modem_configuration = "100"
		}]
		encap = 4094
		dialer_pool_number = 255
		ppp_maximum_payload = 1790
		dialer_address_negotiated = false
		unnumbered_loopback_interface = "example"
		ppp_authentication_protocol = "chap"
		ppp_authentication_protocol_pap = false
		chap_hostname = "chap-example"
		chap_password = "myPassword"
		pap_username = "pap-username"
		pap_password = true
		pap_outbound_password = "myPassword"
		encapsulation = [{
			encapsulation_type = "gre"
			preference = 4294967
			weight = 250
		}]
		groups = [42949672]
		border = true
		per_tunnel_qos = true
		per_tunnel_qos_aggregator = false
		color = "custom1"
		last_resort_circuit = false
		low_bandwidth_link = false
		tunnel_tcp_mss = 1460
		clear_dont_fragment = false
		network_broadcast_1 = false
		max_control_connections = 8
		control_connections = true
		vbond_as_stun_server = false
		exclude_controller_group_list = [100]
		vmanage_connection_preference = 5
		port_hop = false
		restrict = false
		carrier = "carrier1"
		nat_refresh_interval = 15
		hello_interval = 1000
		hello_tolerance = 12
		bind_loopback_tunnel = "12"
		all = false
		network_broadcast_2 = false
		bgp = false
		dhcp = true
		dns = true
		icmp = true
		ssh = false
		netconf = false
		ospf = false
		stun = false
		snmp = false
		https = true
		nat = true
		refresh_mode = "outbound"
		udp_timeout = 1
		tcp_timeout = 60
		block_icmp_error = true
		respond_to_ping = false
		port_forward = [{
			port_start_range = 0
			port_end_range = 65530
			protocol = "tcp"
			private_vpn = 65530
			private_ip_address = "1.2.3.4"
		}]
		adaptive_qos = true
		adapt_period = 15
		shaping_rate_downstream_default = 10000000
		shaping_rate_downstream_min = 10000000
		shaping_rate_downstream_max = 10000000
		shaping_rate_upstream_default = 10000000
		shaping_rate_upstream_min = 10000000
		shaping_rate_upstream_max = 10000000
		shaping_rate = 10000000
		qos_map = "test"
		vpn_qos_map = "test"
		bandwidth_upstream = 214748300
		bandwidth_downstream = 214748300
		write_rule = "test_rule"
		access_list = [{
			direction = "in"
			acl_name = "Egress ACL - IPv4"
		}]
		policer = [{
			direction = "in"
			policer_name = "example"
		}]
		ip_mtu = 1500
		tcp_mss = 720
		tloc_extension = "tloc"
		tracker = ["tracker1"]
		ip_directed_broadcast = true
	}
	`
}
