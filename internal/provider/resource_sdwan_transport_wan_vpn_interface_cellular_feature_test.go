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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanTransportWANVPNInterfaceCellularProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "interface_description", "WAN"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "service_provider", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "bandwidth_upstream", "21474836"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "bandwidth_downstream", "21474836"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_qos_mode", "hub"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_bandwidth_percent", "82"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_bind_loopback_tunnel", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_carrier", "default"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_color", "default"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_color_restrict", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_groups", "42949672"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_border", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_max_control_connections", "62"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_nat_refresh_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_vmanage_connection_preference", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_port_hop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_ntp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_dns", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_icmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_https", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_allow_bfd", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_encapsulations.0.preference", "4294967"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tunnel_interface_encapsulations.0.weight", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "nat_ipv4", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "nat_udp_timeout", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "nat_tcp_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "qos_adaptive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "qos_shaping_rate", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "arps.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "arps.0.mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tcp_mss", "505"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tloc_extension", "tloc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "tracker", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_cellular_feature.test", "ip_directed_broadcast", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportWANVPNInterfaceCellularPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceCellularProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportWANVPNInterfaceCellularPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceCellularProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportWANVPNInterfaceCellularPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name                       = "TF_TEST_WAN"
  description                = "Terraform test"
  feature_profile_id         = sdwan_transport_feature_profile.test.id
  vpn                        = 0
  enhance_ecmp_keying        = true
  primary_dns_address_ipv4   = "1.2.3.4"
  secondary_dns_address_ipv4 = "2.3.4.5"
  primary_dns_address_ipv6   = "2001:0:0:1::0"
  secondary_dns_address_ipv6 = "2001:0:0:2::0"
  new_host_mappings = [
    {
      host_name            = "example"
      list_of_ip_addresses = ["1.2.3.4"]
    }
  ]
  ipv4_static_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      gateway         = "nextHop"
      next_hops = [
        {
          address                 = "1.2.3.4"
          administrative_distance = 1
        }
      ]
    }
  ]
  ipv6_static_routes = [
    {
      prefix  = "2002::/16"
      gateway = "nextHop"
      next_hops = [
        {
          address                 = "2001:0:0:1::0"
          administrative_distance = 1
        }
      ]
    }
  ]
  services = [
    {
      service_type = "TE"
    }
  ]
  nat_64_v4_pools = [
    {
      nat64_v4_pool_name        = "example"
      nat64_v4_pool_range_start = "203.0.113.50"
      nat64_v4_pool_range_end   = "203.0.113.100"
      nat64_v4_pool_overload    = false
    }
  ]
}

resource "sdwan_transport_ipv4_acl_feature" "test" {
  name               = "TF_TEST_ACL_IPV4"
  description        = "Terraform Test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  default_action     = "drop"
  sequences = [
    {
      sequence_id   = 1
      sequence_name = "AccessControlList1"
      match_entries = [
        {
          dscps         = [16]
          packet_length = 1500
          protocols     = [1]
          source_ports = [
            {
              port = 8000
            }
          ]
          tcp_state = "syn"
        }
      ]
      actions = [
        {
          accept_set_dscp     = 60
          accept_counter_name = "COUNTER_1"
          accept_log          = false
          accept_set_next_hop = "1.2.3.4"
        }
      ]
    }
  ]
}

resource "sdwan_transport_ipv6_acl_feature" "test" {
  name               = "TF_TEST_ACL_IPV6"
  description        = "Terraform Test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  default_action     = "drop"
  sequences = [
    {
      sequence_id   = 1
      sequence_name = "AccessControlList1"
      match_entries = [
        {
          next_header   = 10
          packet_length = 1500
          source_ports = [
            {
              port = 8000
            }
          ]
          tcp_state     = "syn"
          traffic_class = [10]
        }
      ]
      actions = [
        {
          accept_counter_name  = "COUNTER_1"
          accept_log           = false
          accept_set_next_hop  = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
          accept_traffic_class = 10
        }
      ]
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportWANVPNInterfaceCellularProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_wan_vpn_interface_cellular_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	tunnel_interface = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportWANVPNInterfaceCellularProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_wan_vpn_interface_cellular_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	interface_description = "WAN"` + "\n"
	config += `	ipv4_dhcp_helper = ["1.2.3.4"]` + "\n"
	config += `	service_provider = "example"` + "\n"
	config += `	bandwidth_upstream = 21474836` + "\n"
	config += `	bandwidth_downstream = 21474836` + "\n"
	config += `	tunnel_interface = true` + "\n"
	config += `	per_tunnel_qos = true` + "\n"
	config += `	tunnel_qos_mode = "hub"` + "\n"
	config += `	tunnel_bandwidth_percent = 82` + "\n"
	config += `	tunnel_interface_bind_loopback_tunnel = "example"` + "\n"
	config += `	tunnel_interface_carrier = "default"` + "\n"
	config += `	tunnel_interface_color = "default"` + "\n"
	config += `	tunnel_interface_hello_interval = 1000` + "\n"
	config += `	tunnel_interface_hello_tolerance = 12` + "\n"
	config += `	tunnel_interface_last_resort_circuit = false` + "\n"
	config += `	tunnel_interface_color_restrict = true` + "\n"
	config += `	tunnel_interface_groups = 42949672` + "\n"
	config += `	tunnel_interface_border = false` + "\n"
	config += `	tunnel_interface_max_control_connections = 62` + "\n"
	config += `	tunnel_interface_nat_refresh_interval = 5` + "\n"
	config += `	tunnel_interface_vbond_as_stun_server = false` + "\n"
	config += `	tunnel_interface_exclude_controller_group_list = [2]` + "\n"
	config += `	tunnel_interface_vmanage_connection_preference = 8` + "\n"
	config += `	tunnel_interface_port_hop = true` + "\n"
	config += `	tunnel_interface_low_bandwidth_link = false` + "\n"
	config += `	tunnel_interface_tunnel_tcp_mss = 1460` + "\n"
	config += `	tunnel_interface_clear_dont_fragment = false` + "\n"
	config += `	tunnel_interface_network_broadcast = false` + "\n"
	config += `	tunnel_interface_allow_all = false` + "\n"
	config += `	tunnel_interface_allow_bgp = false` + "\n"
	config += `	tunnel_interface_allow_dhcp = true` + "\n"
	config += `	tunnel_interface_allow_ntp = false` + "\n"
	config += `	tunnel_interface_allow_ssh = false` + "\n"
	config += `	tunnel_interface_allow_dns = true` + "\n"
	config += `	tunnel_interface_allow_icmp = true` + "\n"
	config += `	tunnel_interface_allow_https = true` + "\n"
	config += `	tunnel_interface_allow_ospf = false` + "\n"
	config += `	tunnel_interface_allow_stun = false` + "\n"
	config += `	tunnel_interface_allow_snmp = false` + "\n"
	config += `	tunnel_interface_allow_netconf = false` + "\n"
	config += `	tunnel_interface_allow_bfd = false` + "\n"
	config += `	tunnel_interface_encapsulations = [{` + "\n"
	config += `	  encapsulation = "gre"` + "\n"
	config += `	  preference = 4294967` + "\n"
	config += `	  weight = 250` + "\n"
	config += `	}]` + "\n"
	config += `	nat_ipv4 = true` + "\n"
	config += `	nat_udp_timeout = 1` + "\n"
	config += `	nat_tcp_timeout = 60` + "\n"
	config += `	qos_adaptive = false` + "\n"
	config += `	qos_shaping_rate = 16` + "\n"
	config += `	acl_ipv4_egress_feature_id = sdwan_transport_ipv4_acl_feature.test.id` + "\n"
	config += `	acl_ipv6_ingress_feature_id = sdwan_transport_ipv6_acl_feature.test.id` + "\n"
	config += `	arps = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	interface_mtu = 1500` + "\n"
	config += `	tcp_mss = 505` + "\n"
	config += `	tloc_extension = "tloc"` + "\n"
	config += `	tracker = "example"` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
