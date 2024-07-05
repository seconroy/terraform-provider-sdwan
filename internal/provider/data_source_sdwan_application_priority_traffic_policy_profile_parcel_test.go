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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "default_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "target_direction", "all"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.name", "RULE_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.dscp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.packet_length", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.tcp", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.traffic_to", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.counter", "COUNTER_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.log", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.dscp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.local_tloc_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.local_tloc_list_encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.tloc_encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_type", "FW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_encapsulation", "ipsec"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.next_hop", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.redirect_dns_field", "redirectDns"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.redirect_dns_value", "umbrella"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_pool", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_vpn", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_fallback", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.fallback_to_routing", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig + testAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig = `
resource "sdwan_application_priority_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_cisco_vpn_feature_template" "test" {
  name                    = "TF_TEST"
  description             = "Terraform test"
  device_types            = ["vedge-C8000V"]
  vpn_id                  = 1
  vpn_name                = "VPN1"
  tenant_vpn_id           = 1
  organization_name       = "org1"
  omp_admin_distance_ipv4 = 10
  omp_admin_distance_ipv6 = 10
  enhance_ecmp_keying     = true
  dns_ipv4_servers = [
    {
      address = "9.9.9.9"
      role    = "primary"
    }
  ]
  dns_ipv6_servers = [
    {
      address = "2001::9"
      role    = "primary"
    }
  ]
  dns_hosts = [
    {
      hostname = "abc1"
      ip       = ["7.7.7.7"]
    }
  ]
  services = [
    {
      service_types = "FW"
      address       = ["8.8.8.8"]
      interface     = "e1"
      track_enable  = true
    }
  ]
  ipv4_static_service_routes = [
    {
      prefix  = "2.2.2.0/24"
      vpn_id  = 2
      service = "sig"
    }
  ]
  ipv4_static_routes = [
    {
      prefix   = "3.3.3.0/24"
      null0    = false
      distance = 10
      vpn_id   = 5
      dhcp     = false
      next_hops = [
        {
          address  = "11.1.1.1"
          distance = 20
        }
      ]
      track_next_hops = [
        {
          address  = "12.1.1.1"
          distance = 20
          tracker  = "tracker1"
        }
      ]
    }
  ]
  ipv6_static_routes = [
    {
      prefix = "2001::/48"
      null0  = false
      vpn_id = 5
      nat    = "NAT64"
      next_hops = [
        {
          address  = "2001::11"
          distance = 20
        }
      ]
    }
  ]
  ipv4_static_gre_routes = [
    {
      prefix    = "3.3.3.0/24"
      vpn_id    = 2
      interface = ["e1"]
    }
  ]
  ipv4_static_ipsec_routes = [
    {
      prefix    = "4.4.4.0/24"
      vpn_id    = 2
      interface = ["e1"]
    }
  ]
  omp_advertise_ipv4_routes = [
    {
      protocol          = "bgp"
      route_policy      = "rp1"
      protocol_sub_type = ["external"]
      prefixes = [
        {
          prefix_entry   = "1.1.1.0/24"
          aggregate_only = true
        }
      ]
    }
  ]
  omp_advertise_ipv6_routes = [
    {
      protocol          = "bgp"
      route_policy      = "rp1"
      protocol_sub_type = ["external"]
      prefixes = [
        {
          prefix_entry   = "2001:2::/48"
          aggregate_only = true
        }
      ]
    }
  ]
  nat64_pools = [
    {
      name                      = "POOL1"
      start_address             = "100.1.1.1"
      end_address               = "100.1.2.255"
      overload                  = true
      leak_from_global          = true
      leak_from_global_protocol = "rip"
      leak_to_global            = true
    }
  ]
  nat_pools = [
    {
      name          = 1
      prefix_length = 24
      range_start   = "101.1.1.1"
      range_end     = "101.1.2.255"
      overload      = true
      direction     = "inside"
      tracker_id    = 10
    }
  ]
  static_nat_rules = [
    {
      pool_name            = 1
      source_ip            = "10.1.1.1"
      translate_ip         = "105.1.1.1"
      static_nat_direction = "inside"
      tracker_id           = 10
    }
  ]
  static_nat_subnet_rules = [
    {
      source_ip_subnet     = "10.2.1.0"
      translate_ip_subnet  = "105.2.1.0"
      prefix_length        = 24
      static_nat_direction = "inside"
      tracker_id           = 10
    }
  ]
  port_forward_rules = [
    {
      pool_name      = 1
      source_port    = 5000
      translate_port = 6000
      source_ip      = "10.3.1.1"
      translate_ip   = "120.3.1.1"
      protocol       = "tcp"
    }
  ]
  route_global_imports = [
    {
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
  route_vpn_imports = [
    {
      source_vpn_id     = 5
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
  route_global_exports = [
    {
      protocol          = "ospf"
      protocol_sub_type = ["external"]
      route_policy      = "policy1"
      redistributes = [
        {
          protocol     = "bgp"
          route_policy = "policy1"
        }
      ]
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcelConfig() string {
	config := `resource "sdwan_application_priority_traffic_policy_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_application_priority_feature_profile.test.id` + "\n"
	config += `	default_action = "accept"` + "\n"
	config += `	vpn = sdwan_cisco_vpn_feature_template.test.name` + "\n"
	config += `	target_direction = "all"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  sequence_id = 1` + "\n"
	config += `	  name = "RULE_1"` + "\n"
	config += `	  protocol = "ipv4"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  matches = [{` + "\n"
	config += `		dscp = 1` + "\n"
	config += `		packet_length = "123"` + "\n"
	config += `		protocol = ["2"]` + "\n"
	config += `		tcp = "gre"` + "\n"
	config += `		traffic_to = "core"` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		counter = "COUNTER_1"` + "\n"
	config += `		log = false` + "\n"
	config += `      sla_class = [{` + "\n"
	config += `			preferred_color = ["default"]` + "\n"
	config += `		}]` + "\n"
	config += `      sets = [{` + "\n"
	config += `			dscp = 1` + "\n"
	config += `			local_tloc_list_color = ["default"]` + "\n"
	config += `			local_tloc_restrict = "false"` + "\n"
	config += `			local_tloc_list_encapsulation = "gre"` + "\n"
	config += `			tloc_ip = "1.2.3.4"` + "\n"
	config += `			tloc_color = ["default"]` + "\n"
	config += `			tloc_encapsulation = "gre"` + "\n"
	config += `			service_type = "FW"` + "\n"
	config += `			service_color = ["default"]` + "\n"
	config += `			service_encapsulation = "ipsec"` + "\n"
	config += `			service_tloc_ip = "1.2.3.4"` + "\n"
	config += `			next_hop = "1.2.3.4"` + "\n"
	config += `		}]` + "\n"
	config += `		redirect_dns_field = "redirectDns"` + "\n"
	config += `		redirect_dns_value = "umbrella"` + "\n"
	config += `		nat_pool = 2` + "\n"
	config += `		nat_vpn = false` + "\n"
	config += `		nat_fallback = false` + "\n"
	config += `		fallback_to_routing = true` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_application_priority_traffic_policy_profile_parcel" "test" {
			id = sdwan_application_priority_traffic_policy_profile_parcel.test.id
			feature_profile_id = sdwan_application_priority_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
