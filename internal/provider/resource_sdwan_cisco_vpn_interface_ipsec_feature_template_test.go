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
func TestAccSdwanCiscoVPNInterfaceIPSecFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "interface_name", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "interface_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ip_address", "1.1.1.1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "tunnel_source", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "tunnel_source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "tunnel_destination", "3.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "application", "sig"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "tcp_mss_adjust", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "clear_dont_fragment", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "dead_peer_detection_interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "dead_peer_detection_retries", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_mode", "main"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_rekey_interval", "20000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_ciphersuite", "aes256-cbc-sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_group", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_pre_shared_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_pre_shared_key_local_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ike_pre_shared_key_remote_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ipsec_rekey_interval", "7200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ipsec_replay_window", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ipsec_ciphersuite", "aes256-cbc-sha256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "ipsec_perfect_forward_secrecy", "group-20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_vpn_interface_ipsec_feature_template.test", "tunnel_route_via", "g0/0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoVPNInterfaceIPSecFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoVPNInterfaceIPSecFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoVPNInterfaceIPSecFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_vpn_interface_ipsec_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoVPNInterfaceIPSecFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_vpn_interface_ipsec_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	interface_name = "ipsec1"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	interface_description = "My Description"` + "\n"
	config += `	ip_address = "1.1.1.1/24"` + "\n"
	config += `	tunnel_source = "1.2.3.4"` + "\n"
	config += `	tunnel_source_interface = "e1"` + "\n"
	config += `	tunnel_destination = "3.4.5.6"` + "\n"
	config += `	application = "sig"` + "\n"
	config += `	tcp_mss_adjust = 1400` + "\n"
	config += `	clear_dont_fragment = true` + "\n"
	config += `	mtu = 1500` + "\n"
	config += `	dead_peer_detection_interval = 100` + "\n"
	config += `	dead_peer_detection_retries = 4` + "\n"
	config += `	ike_version = 2` + "\n"
	config += `	ike_mode = "main"` + "\n"
	config += `	ike_rekey_interval = 20000` + "\n"
	config += `	ike_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	ike_group = "20"` + "\n"
	config += `	ike_pre_shared_key = "cisco123"` + "\n"
	config += `	ike_pre_shared_key_local_id = "1"` + "\n"
	config += `	ike_pre_shared_key_remote_id = "2"` + "\n"
	config += `	ipsec_rekey_interval = 7200` + "\n"
	config += `	ipsec_replay_window = 128` + "\n"
	config += `	ipsec_ciphersuite = "aes256-cbc-sha256"` + "\n"
	config += `	ipsec_perfect_forward_secrecy = "group-20"` + "\n"
	config += `	tracker = ["TRACKER1"]` + "\n"
	config += `	tunnel_route_via = "g0/0"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
