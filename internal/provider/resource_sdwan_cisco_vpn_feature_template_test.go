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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanCiscoVPNFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "vpn_name", "VPN1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "tenant_vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "organization_name", "org1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_admin_distance_ipv4", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_admin_distance_ipv6", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "enhance_ecmp_keying", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "dns_ipv4_servers.0.address", "9.9.9.9"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "dns_ipv4_servers.0.role", "primary"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "dns_ipv6_servers.0.address", "2001::9"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "dns_ipv6_servers.0.role", "primary"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "dns_hosts.0.hostname", "abc1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "services.0.service_types", "FW"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "services.0.interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "services.0.track_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.prefix", "2.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.vpn_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.service", "sig"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.prefix", "3.3.3.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.null0", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.distance", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.vpn_id", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.dhcp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.next_hops.0.address", "11.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.next_hops.0.distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.address", "12.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.tracker", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.prefix", "2001::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.null0", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.vpn_id", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.nat", "NAT64"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.next_hops.0.address", "2001::11"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.next_hops.0.distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_gre_routes.0.prefix", "3.3.3.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_gre_routes.0.vpn_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_ipsec_routes.0.prefix", "4.4.4.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "ipv4_static_ipsec_routes.0.vpn_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.route_policy", "rp1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.prefixes.0.prefix_entry", "1.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.prefixes.0.aggregate_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.route_policy", "rp1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.prefixes.0.prefix_entry", "2001:2::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.prefixes.0.aggregate_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.name", "POOL1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.start_address", "100.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.end_address", "100.1.2.255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.overload", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_from_global", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_from_global_protocol", "rip"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_to_global", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.range_start", "101.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.range_end", "101.1.2.255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.overload", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "nat_pools.0.tracker_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.pool_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.source_ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.translate_ip", "105.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.static_nat_direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.tracker_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.source_ip_subnet", "10.2.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.translate_ip_subnet", "105.2.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.static_nat_direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.tracker_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.pool_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.source_port", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.translate_port", "6000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.source_ip", "10.3.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.translate_ip", "120.3.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.protocol", "ospf"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.route_policy", "policy1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.redistributes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.redistributes.0.route_policy", "policy1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.source_vpn_id", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.protocol", "ospf"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.route_policy", "policy1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.redistributes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.redistributes.0.route_policy", "policy1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.protocol", "ospf"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.route_policy", "policy1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.redistributes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.redistributes.0.route_policy", "policy1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoVPNFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoVPNFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoVPNFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_vpn_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoVPNFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_vpn_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	vpn_id = 1` + "\n"
	config += `	vpn_name = "VPN1"` + "\n"
	config += `	tenant_vpn_id = 1` + "\n"
	config += `	organization_name = "org1"` + "\n"
	config += `	omp_admin_distance_ipv4 = 10` + "\n"
	config += `	omp_admin_distance_ipv6 = 10` + "\n"
	config += `	enhance_ecmp_keying = true` + "\n"
	config += `	dns_ipv4_servers = [{` + "\n"
	config += `	  address = "9.9.9.9"` + "\n"
	config += `	  role = "primary"` + "\n"
	config += `	}]` + "\n"
	config += `	dns_ipv6_servers = [{` + "\n"
	config += `	  address = "2001::9"` + "\n"
	config += `	  role = "primary"` + "\n"
	config += `	}]` + "\n"
	config += `	dns_hosts = [{` + "\n"
	config += `	  hostname = "abc1"` + "\n"
	config += `	  ip = ["7.7.7.7"]` + "\n"
	config += `	}]` + "\n"
	config += `	services = [{` + "\n"
	config += `	  service_types = "FW"` + "\n"
	config += `	  address = ["8.8.8.8"]` + "\n"
	config += `	  interface = "e1"` + "\n"
	config += `	  track_enable = true` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_service_routes = [{` + "\n"
	config += `	  prefix = "2.2.2.0/24"` + "\n"
	config += `	  vpn_id = 2` + "\n"
	config += `	  service = "sig"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_routes = [{` + "\n"
	config += `	  prefix = "3.3.3.0/24"` + "\n"
	config += `	  null0 = false` + "\n"
	config += `	  distance = 10` + "\n"
	config += `	  vpn_id = 5` + "\n"
	config += `	  dhcp = false` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "11.1.1.1"` + "\n"
	config += `		distance = 20` + "\n"
	config += `	}]` + "\n"
	config += `	  track_next_hops = [{` + "\n"
	config += `		address = "12.1.1.1"` + "\n"
	config += `		distance = 20` + "\n"
	config += `		tracker = "tracker1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_static_routes = [{` + "\n"
	config += `	  prefix = "2001::/48"` + "\n"
	config += `	  null0 = false` + "\n"
	config += `	  vpn_id = 5` + "\n"
	config += `	  nat = "NAT64"` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "2001::11"` + "\n"
	config += `		distance = 20` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_gre_routes = [{` + "\n"
	config += `	  prefix = "3.3.3.0/24"` + "\n"
	config += `	  vpn_id = 2` + "\n"
	config += `	  interface = ["e1"]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_ipsec_routes = [{` + "\n"
	config += `	  prefix = "4.4.4.0/24"` + "\n"
	config += `	  vpn_id = 2` + "\n"
	config += `	  interface = ["e1"]` + "\n"
	config += `	}]` + "\n"
	config += `	omp_advertise_ipv4_routes = [{` + "\n"
	config += `	  protocol = "bgp"` + "\n"
	config += `	  route_policy = "rp1"` + "\n"
	config += `	  protocol_sub_type = ["external"]` + "\n"
	config += `	  prefixes = [{` + "\n"
	config += `		prefix_entry = "1.1.1.0/24"` + "\n"
	config += `		aggregate_only = true` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	omp_advertise_ipv6_routes = [{` + "\n"
	config += `	  protocol = "bgp"` + "\n"
	config += `	  route_policy = "rp1"` + "\n"
	config += `	  protocol_sub_type = ["external"]` + "\n"
	config += `	  prefixes = [{` + "\n"
	config += `		prefix_entry = "2001:2::/48"` + "\n"
	config += `		aggregate_only = true` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	nat64_pools = [{` + "\n"
	config += `	  name = "POOL1"` + "\n"
	config += `	  start_address = "100.1.1.1"` + "\n"
	config += `	  end_address = "100.1.2.255"` + "\n"
	config += `	  overload = true` + "\n"
	config += `	  leak_from_global = true` + "\n"
	config += `	  leak_from_global_protocol = "rip"` + "\n"
	config += `	  leak_to_global = true` + "\n"
	config += `	}]` + "\n"
	config += `	nat_pools = [{` + "\n"
	config += `	  name = 1` + "\n"
	config += `	  prefix_length = 24` + "\n"
	config += `	  range_start = "101.1.1.1"` + "\n"
	config += `	  range_end = "101.1.2.255"` + "\n"
	config += `	  overload = true` + "\n"
	config += `	  direction = "inside"` + "\n"
	config += `	  tracker_id = 10` + "\n"
	config += `	}]` + "\n"
	config += `	static_nat_rules = [{` + "\n"
	config += `	  pool_name = 1` + "\n"
	config += `	  source_ip = "10.1.1.1"` + "\n"
	config += `	  translate_ip = "105.1.1.1"` + "\n"
	config += `	  static_nat_direction = "inside"` + "\n"
	config += `	  tracker_id = 10` + "\n"
	config += `	}]` + "\n"
	config += `	static_nat_subnet_rules = [{` + "\n"
	config += `	  source_ip_subnet = "10.2.1.0"` + "\n"
	config += `	  translate_ip_subnet = "105.2.1.0"` + "\n"
	config += `	  prefix_length = 24` + "\n"
	config += `	  static_nat_direction = "inside"` + "\n"
	config += `	  tracker_id = 10` + "\n"
	config += `	}]` + "\n"
	config += `	port_forward_rules = [{` + "\n"
	config += `	  pool_name = 1` + "\n"
	config += `	  source_port = 5000` + "\n"
	config += `	  translate_port = 6000` + "\n"
	config += `	  source_ip = "10.3.1.1"` + "\n"
	config += `	  translate_ip = "120.3.1.1"` + "\n"
	config += `	  protocol = "tcp"` + "\n"
	config += `	}]` + "\n"
	config += `	route_global_imports = [{` + "\n"
	config += `	  protocol = "ospf"` + "\n"
	config += `	  protocol_sub_type = ["external"]` + "\n"
	config += `	  route_policy = "policy1"` + "\n"
	config += `	  redistributes = [{` + "\n"
	config += `		protocol = "bgp"` + "\n"
	config += `		route_policy = "policy1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	route_vpn_imports = [{` + "\n"
	config += `	  source_vpn_id = 5` + "\n"
	config += `	  protocol = "ospf"` + "\n"
	config += `	  protocol_sub_type = ["external"]` + "\n"
	config += `	  route_policy = "policy1"` + "\n"
	config += `	  redistributes = [{` + "\n"
	config += `		protocol = "bgp"` + "\n"
	config += `		route_policy = "policy1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	route_global_exports = [{` + "\n"
	config += `	  protocol = "ospf"` + "\n"
	config += `	  protocol_sub_type = ["external"]` + "\n"
	config += `	  route_policy = "policy1"` + "\n"
	config += `	  redistributes = [{` + "\n"
	config += `		protocol = "bgp"` + "\n"
	config += `		route_policy = "policy1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
