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

func TestAccSdwanCiscoVPNInterfaceFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoVPNInterfaceFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "interface_name", "ge0/0"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_type", "interface"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_start", "10.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_end", "10.1.1.255"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_prefix_length", "24"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "per_tunnel_qos", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "per_tunnel_qos_aggregator", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_qos_mode", "spoke"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_bandwidth", "50"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_downstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_downstream", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_downstream", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_upstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_upstream", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_upstream", "100000"),
				),
			},
			{
				Config: testAccSdwanCiscoVPNInterfaceFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "interface_name", "ge0/0"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "interface_description", "My Interface Description"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "poe", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "address", "1.1.1.1/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_secondary_addresses.0.address", "2.2.2.2/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "dhcp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "dhcp_distance", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_address", "2001:1::1/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "dhcpv6", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_secondary_addresses.0.address", "2.2.2.2/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_access_lists.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_access_lists.0.acl_name", "ACL1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_dhcp_helpers.0.address", "2001:7::7/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_dhcp_helpers.0.vpn_id", "5"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "auto_bandwidth_detect", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "iperf_server", "8.8.8.8"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_type", "interface"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "udp_timeout", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tcp_timeout", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_start", "10.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_end", "10.1.1.255"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_overload", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_inside_source_loopback_interface", "lo1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_prefix_length", "24"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_nat", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat64_interface", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "nat66_interface", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.source_prefix", "2001:7::/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.translated_source_prefix", "2001:8::/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.source_vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.source_ip", "10.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.translate_ip", "100.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.static_nat_direction", "inside"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.source_vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_ip", "10.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.translate_ip", "100.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.static_nat_direction", "inside"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_port", "8000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.translate_port", "9000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "enable_core_region", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "core_region", "core"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "secondary_region", "off"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.preference", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.weight", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_border", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "per_tunnel_qos", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "per_tunnel_qos_aggregator", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_qos_mode", "spoke"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_bandwidth", "50"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_color", "gold"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_max_control_connections", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_control_connections", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_vbond_as_stun_server", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_vmanage_connection_preference", "5"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_port_hop", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_color_restrict", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_gre_tunnel_destination_ip", "5.5.5.5"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_carrier", "carrier1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_nat_refresh_interval", "5"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_hello_interval", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_hello_tolerance", "12"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_bind_loopback_tunnel", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_last_resort_circuit", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_low_bandwidth_link", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_tunnel_tcp_mss", "1460"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_clear_dont_fragment", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_propagate_sgt", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_network_broadcast", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_all", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_bgp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_dhcp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_dns", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_icmp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ssh", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_netconf", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ntp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ospf", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_stun", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_snmp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_https", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "media_type", "auto-select"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "interface_mtu", "9216"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ip_mtu", "1500"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tcp_mss_adjust", "1460"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "tloc_extension", "123"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "load_interval", "30"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "gre_tunnel_source_ip", "3.3.3.3"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "gre_tunnel_xconnect", "a123"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "mac_address", "00-B0-D0-63-C2-26"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "speed", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "duplex", "full"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "arp_timeout", "1200"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "autonegotiate", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ip_directed_broadcast", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "icmp_redirect_disable", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_period", "15"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_downstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_downstream", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_downstream", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_upstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_upstream", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_upstream", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "shaping_rate", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_map", "QOSMAP1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "qos_map_vpn", "QOSMAP2"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "bandwidth_upstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "bandwidth_downstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "block_non_source_ip", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "rewrite_rule_name", "RULE1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "access_lists.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "access_lists.0.acl_name", "ACL1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_arps.0.ip_address", "8.8.8.8"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_arps.0.mac", "00-B0-D0-63-C2-26"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.group_id", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.priority", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.timer", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.track_omp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.track_prefix_list", "PL1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.ip_address", "2.2.2.2"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.ipv4_secondary_addresses.0.ip_address", "2.2.2.3"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tloc_preference_change", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tloc_preference_change_value", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.tracker_id", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.track_action", "decrement"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.decrement_value", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.group_id", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.priority", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.timer", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.track_omp", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.track_prefix_list", "PL1"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.ipv6_link_local", "fe80::260:8ff:fe52:f9d8"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.prefix", "2001:9::/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "propagate_sgt", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_sgt", "1003"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "static_sgt_trusted", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "enable_sgt", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "sgt_enforcement", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_feature_template.test", "sgt_enforcement_sgt", "1004"),
				),
			},
		},
	})
}

func testAccSdwanCiscoVPNInterfaceFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_vpn_interface_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		interface_name = "ge0/0"
		nat_type = "interface"
		nat_pool_range_start = "10.1.1.1"
		nat_pool_range_end = "10.1.1.255"
		nat_pool_prefix_length = 24
		per_tunnel_qos = false
		per_tunnel_qos_aggregator = false
		tunnel_qos_mode = "spoke"
		tunnel_bandwidth = 50
		qos_adaptive_bandwidth_downstream = 10000
		qos_adaptive_min_downstream = 100
		qos_adaptive_max_downstream = 100000
		qos_adaptive_bandwidth_upstream = 10000
		qos_adaptive_min_upstream = 100
		qos_adaptive_max_upstream = 100000
	}
	`
}

func testAccSdwanCiscoVPNInterfaceFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_vpn_interface_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		interface_name = "ge0/0"
		interface_description = "My Interface Description"
		poe = false
		address = "1.1.1.1/24"
		ipv4_secondary_addresses = [{
			address = "2.2.2.2/24"
		}]
		dhcp = false
		dhcp_distance = 10
		ipv6_address = "2001:1::1/48"
		dhcpv6 = false
		ipv6_secondary_addresses = [{
			address = "2.2.2.2/24"
		}]
		ipv6_access_lists = [{
			direction = "in"
			acl_name = "ACL1"
		}]
		ipv4_dhcp_helper = ["6.6.6.6"]
		ipv6_dhcp_helpers = [{
			address = "2001:7::7/48"
			vpn_id = 5
		}]
		tracker = ["tracker1"]
		auto_bandwidth_detect = false
		iperf_server = "8.8.8.8"
		nat = true
		nat_type = "interface"
		udp_timeout = 1
		tcp_timeout = 60
		nat_pool_range_start = "10.1.1.1"
		nat_pool_range_end = "10.1.1.255"
		nat_overload = false
		nat_inside_source_loopback_interface = "lo1"
		nat_pool_prefix_length = 24
		ipv6_nat = false
		nat64_interface = false
		nat66_interface = false
		static_nat66_entries = [{
			source_prefix = "2001:7::/48"
			translated_source_prefix = "2001:8::/48"
			source_vpn_id = 1
		}]
		static_nat_entries = [{
			source_ip = "10.1.1.1"
			translate_ip = "100.1.1.1"
			static_nat_direction = "inside"
			source_vpn_id = 1
		}]
		static_port_forward_entries = [{
			source_ip = "10.1.1.1"
			translate_ip = "100.1.1.1"
			static_nat_direction = "inside"
			source_port = 8000
			translate_port = 9000
			protocol = "tcp"
			source_vpn_id = 1
		}]
		enable_core_region = false
		core_region = "core"
		secondary_region = "off"
		tunnel_interface_encapsulations = [{
			encapsulation = "gre"
			preference = 10
			weight = 100
		}]
		tunnel_interface_border = false
		per_tunnel_qos = false
		per_tunnel_qos_aggregator = false
		tunnel_qos_mode = "spoke"
		tunnel_bandwidth = 50
		tunnel_interface_groups = [5]
		tunnel_interface_color = "gold"
		tunnel_interface_max_control_connections = 10
		tunnel_interface_control_connections = false
		tunnel_interface_vbond_as_stun_server = false
		tunnel_interface_exclude_controller_group_list = [10]
		tunnel_interface_vmanage_connection_preference = 5
		tunnel_interface_port_hop = false
		tunnel_interface_color_restrict = false
		tunnel_interface_gre_tunnel_destination_ip = "5.5.5.5"
		tunnel_interface_carrier = "carrier1"
		tunnel_interface_nat_refresh_interval = 5
		tunnel_interface_hello_interval = 1000
		tunnel_interface_hello_tolerance = 12
		tunnel_interface_bind_loopback_tunnel = "1"
		tunnel_interface_last_resort_circuit = false
		tunnel_interface_low_bandwidth_link = false
		tunnel_interface_tunnel_tcp_mss = 1460
		tunnel_interface_clear_dont_fragment = false
		tunnel_interface_propagate_sgt = false
		tunnel_interface_network_broadcast = false
		tunnel_interface_allow_all = false
		tunnel_interface_allow_bgp = false
		tunnel_interface_allow_dhcp = false
		tunnel_interface_allow_dns = false
		tunnel_interface_allow_icmp = false
		tunnel_interface_allow_ssh = false
		tunnel_interface_allow_netconf = false
		tunnel_interface_allow_ntp = false
		tunnel_interface_allow_ospf = false
		tunnel_interface_allow_stun = false
		tunnel_interface_allow_snmp = false
		tunnel_interface_allow_https = false
		media_type = "auto-select"
		interface_mtu = 9216
		ip_mtu = 1500
		tcp_mss_adjust = 1460
		tloc_extension = "123"
		load_interval = 30
		gre_tunnel_source_ip = "3.3.3.3"
		gre_tunnel_xconnect = "a123"
		mac_address = "00-B0-D0-63-C2-26"
		speed = "1000"
		duplex = "full"
		shutdown = false
		arp_timeout = 1200
		autonegotiate = true
		ip_directed_broadcast = false
		icmp_redirect_disable = false
		qos_adaptive = false
		qos_adaptive_period = 15
		qos_adaptive_bandwidth_downstream = 10000
		qos_adaptive_min_downstream = 100
		qos_adaptive_max_downstream = 100000
		qos_adaptive_bandwidth_upstream = 10000
		qos_adaptive_min_upstream = 100
		qos_adaptive_max_upstream = 100000
		shaping_rate = 1000
		qos_map = "QOSMAP1"
		qos_map_vpn = "QOSMAP2"
		bandwidth_upstream = 10000
		bandwidth_downstream = 10000
		block_non_source_ip = false
		rewrite_rule_name = "RULE1"
		access_lists = [{
			direction = "in"
			acl_name = "ACL1"
		}]
		static_arps = [{
			ip_address = "8.8.8.8"
			mac = "00-B0-D0-63-C2-26"
		}]
		ipv4_vrrps = [{
			group_id = 100
			priority = 100
			timer = 100
			track_omp = false
			track_prefix_list = "PL1"
			ip_address = "2.2.2.2"
			ipv4_secondary_addresses = [{
				ip_address = "2.2.2.3"
			}]
			tloc_preference_change = false
			tloc_preference_change_value = 10
			tracking_objects = [{
				tracker_id = 10
				track_action = "decrement"
				decrement_value = 100
			}]
		}]
		ipv6_vrrps = [{
			group_id = 100
			priority = 100
			timer = 100
			track_omp = false
			track_prefix_list = "PL1"
			ipv6_addresses = [{
				ipv6_link_local = "fe80::260:8ff:fe52:f9d8"
				prefix = "2001:9::/48"
			}]
		}]
		propagate_sgt = false
		static_sgt = 1003
		static_sgt_trusted = false
		enable_sgt = true
		sgt_enforcement = true
		sgt_enforcement_sgt = 1004
	}
	`
}
