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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanTLSSSLDecryptionPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "mode", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "default_action", "noIntent"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.base_action", "doNotDecrypt"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.rule_id", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.rule_name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.rule_type", "sslDecryption"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.source_and_destination_configuration.0.option", "destinationIp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "network_rules.0.source_and_destination_configuration.0.value", "10.0.0.0/12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "ssl_decryption_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "expired_certificate", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "untrusted_certificate", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "certificate_revocation_status", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "unknown_revocation_status", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "unsupported_protocol_versions", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "unsupported_cipher_suites", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "failure_mode", "close"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "rsa_key_pair_modulus", "2048"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "ec_key_type", "P384"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "certificate_lifetime_in_days", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "minimal_tls_version", "TLSv1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_decryption_policy_definition.test", "use_default_ca_cert_bundle", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTLSSSLDecryptionPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTLSSSLDecryptionPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_tls_ssl_decryption_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	mode = "security"` + "\n"
	config += `	default_action = "noIntent"` + "\n"
	config += `	network_rules = [{` + "\n"
	config += `	  base_action = "doNotDecrypt"` + "\n"
	config += `	  rule_id = 4` + "\n"
	config += `	  rule_name = "Example"` + "\n"
	config += `	  rule_type = "sslDecryption"` + "\n"
	config += `	  source_and_destination_configuration = [{` + "\n"
	config += `		option = "destinationIp"` + "\n"
	config += `		value = "10.0.0.0/12"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ssl_decryption_enabled = "true"` + "\n"
	config += `	expired_certificate = "drop"` + "\n"
	config += `	untrusted_certificate = "drop"` + "\n"
	config += `	certificate_revocation_status = "none"` + "\n"
	config += `	unknown_revocation_status = "drop"` + "\n"
	config += `	unsupported_protocol_versions = "drop"` + "\n"
	config += `	unsupported_cipher_suites = "drop"` + "\n"
	config += `	failure_mode = "close"` + "\n"
	config += `	rsa_key_pair_modulus = "2048"` + "\n"
	config += `	ec_key_type = "P384"` + "\n"
	config += `	certificate_lifetime_in_days = 1` + "\n"
	config += `	minimal_tls_version = "TLSv1.2"` + "\n"
	config += `	use_default_ca_cert_bundle = true` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_tls_ssl_decryption_policy_definition" "test" {
			id = sdwan_tls_ssl_decryption_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
