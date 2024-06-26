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
func TestAccSdwanVPNInterfaceSVIFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "if_name", "Vlan100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "interface_description", "VPN Interface SVI"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_secondary_addresses.0.ipv4_address", "4.5.6.7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_address", "2001:db8:85a3::8a2e:370:7334"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_dhcp_client", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_dhcp_distance", "101"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_dhcp_rapid_commit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_secondary_addresses.0.ipv6_address", "2001:db8:85a3::8a2e:370:7334"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_dhcp_helpers.0.address", "2001:db8:85a3::8a2e:370:7334"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_dhcp_helpers.0.vpn_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ip_directed_broadcast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "tcp_mss_adjust", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "arp_timeout", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_access_lists.0.acl_name", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_access_lists.0.acl_name", "ACL2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "policers.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "policers.0.policer_name", "POLICER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "static_arp_entries.0.ipv4_address", "3.4.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "static_arp_entries.0.mac_address", "00:00:00:00:00:00"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.track_omp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.track_prefix_list", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.ipv4_address", "5.6.7.8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.ipv4_secondary_addresses.0.ipv4_address", "8.8.8.8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.tloc_preference_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.tloc_preference_change_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.name", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.track_action", "decrement"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.decrement_value", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.track_omp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.track_prefix_list", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.link_local_address", "FE80::1/64"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.prefix", "2001:db8:85a3::8a2e:370:7335"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_svi_feature_template.test", "ipv6_vrrps.0.ipv6_secondary_addresses.0.prefix", "2001:db8:85a3::8a2e:370:7336"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceSVIFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanVPNInterfaceSVIFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanVPNInterfaceSVIFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_vpn_interface_svi_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanVPNInterfaceSVIFeatureTemplateConfig_all() string {
	config := `resource "sdwan_vpn_interface_svi_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	if_name = "Vlan100"` + "\n"
	config += `	interface_description = "VPN Interface SVI"` + "\n"
	config += `	ipv4_address = "2.3.4.5"` + "\n"
	config += `	ipv4_secondary_addresses = [{` + "\n"
	config += `	  ipv4_address = "4.5.6.7"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_address = "2001:db8:85a3::8a2e:370:7334"` + "\n"
	config += `	ipv6_dhcp_client = false` + "\n"
	config += `	ipv6_dhcp_distance = 101` + "\n"
	config += `	ipv6_dhcp_rapid_commit = false` + "\n"
	config += `	ipv6_secondary_addresses = [{` + "\n"
	config += `	  ipv6_address = "2001:db8:85a3::8a2e:370:7334"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_dhcp_helper = ["7.7.7.7"]` + "\n"
	config += `	ipv6_dhcp_helpers = [{` + "\n"
	config += `	  address = "2001:db8:85a3::8a2e:370:7334"` + "\n"
	config += `	  vpn_id = 100` + "\n"
	config += `	}]` + "\n"
	config += `	ip_directed_broadcast = true` + "\n"
	config += `	mtu = 1500` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	tcp_mss_adjust = 1400` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	arp_timeout = 100` + "\n"
	config += `	ipv4_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL2"` + "\n"
	config += `	}]` + "\n"
	config += `	policers = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  policer_name = "POLICER1"` + "\n"
	config += `	}]` + "\n"
	config += `	static_arp_entries = [{` + "\n"
	config += `	  ipv4_address = "3.4.4.5"` + "\n"
	config += `	  mac_address = "00:00:00:00:00:00"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = true` + "\n"
	config += `	  track_prefix_list = "TRACK1"` + "\n"
	config += `	  ipv4_address = "5.6.7.8"` + "\n"
	config += `	  ipv4_secondary_addresses = [{` + "\n"
	config += `		ipv4_address = "8.8.8.8"` + "\n"
	config += `	}]` + "\n"
	config += `	  tloc_preference_change = true` + "\n"
	config += `	  tloc_preference_change_value = 100` + "\n"
	config += `	  tracking_objects = [{` + "\n"
	config += `		name = 100` + "\n"
	config += `		track_action = "decrement"` + "\n"
	config += `		decrement_value = 10` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = true` + "\n"
	config += `	  track_prefix_list = "TRACK1"` + "\n"
	config += `	  ipv6_addresses = [{` + "\n"
	config += `		link_local_address = "FE80::1/64"` + "\n"
	config += `		prefix = "2001:db8:85a3::8a2e:370:7335"` + "\n"
	config += `	}]` + "\n"
	config += `	  ipv6_secondary_addresses = [{` + "\n"
	config += `		prefix = "2001:db8:85a3::8a2e:370:7336"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
