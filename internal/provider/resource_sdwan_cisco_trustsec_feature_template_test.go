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
func TestAccSdwanCiscoTrustSecFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "device_sgt", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "credentials_id", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "credentials_password", "MyPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "enable_enforcement", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "enable_sxp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_default_password", "MyPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_key_chain", "keychain1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_log_binding_changes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_reconciliation_period", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_retry_period", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "speaker_hold_time", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "minimum_listener_hold_time", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "maximum_listener_hold_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id_type", "interface-name"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id", "VirtualPortGroup"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.peer_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.source_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.preshared_key", "default"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.mode", "local"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.mode_type", "listener"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.minimum_hold_time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.maximum_hold_time", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connections.0.vpn_id", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoTrustSecFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoTrustSecFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoTrustSecFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_trustsec_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoTrustSecFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_trustsec_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	device_sgt = 100` + "\n"
	config += `	credentials_id = "user1"` + "\n"
	config += `	credentials_password = "MyPassword"` + "\n"
	config += `	enable_enforcement = true` + "\n"
	config += `	enable_sxp = true` + "\n"
	config += `	sxp_source_ip = "1.2.3.4"` + "\n"
	config += `	sxp_default_password = "MyPassword"` + "\n"
	config += `	sxp_key_chain = "keychain1"` + "\n"
	config += `	sxp_log_binding_changes = false` + "\n"
	config += `	sxp_reconciliation_period = 120` + "\n"
	config += `	sxp_retry_period = 120` + "\n"
	config += `	speaker_hold_time = 120` + "\n"
	config += `	minimum_listener_hold_time = 90` + "\n"
	config += `	maximum_listener_hold_time = 180` + "\n"
	config += `	sxp_node_id_type = "interface-name"` + "\n"
	config += `	sxp_node_id = "VirtualPortGroup"` + "\n"
	config += `	sxp_connections = [{` + "\n"
	config += `	  peer_ip = "1.2.3.4"` + "\n"
	config += `	  source_ip = "2.3.4.5"` + "\n"
	config += `	  preshared_key = "default"` + "\n"
	config += `	  mode = "local"` + "\n"
	config += `	  mode_type = "listener"` + "\n"
	config += `	  minimum_hold_time = 100` + "\n"
	config += `	  maximum_hold_time = 200` + "\n"
	config += `	  vpn_id = 0` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
