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
func TestAccSdwanSystemSecurityProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "rekey", "86400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "anti_replay_window", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "extended_anti_replay_window", "256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "ipsec_pairwise_keying", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keychains.0.key_chain_name", "aaa"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keychains.0.key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.name", "aaa"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.send_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.receiver_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.include_tcp_options", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.accept_ao_mismatch", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.crypto_algorithm", "aes-128-cmac"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.send_life_time_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.send_life_time_start_epoch", "1659284400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.send_life_time_infinite", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.accept_life_time_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.accept_life_time_start_epoch", "1659284400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_security_feature.test", "keys.0.accept_life_time_infinite", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemSecurityPrerequisitesProfileParcelConfig + testAccSdwanSystemSecurityProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemSecurityPrerequisitesProfileParcelConfig + testAccSdwanSystemSecurityProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemSecurityPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemSecurityProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_security_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemSecurityProfileParcelConfig_all() string {
	config := `resource "sdwan_system_security_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	rekey = 86400` + "\n"
	config += `	anti_replay_window = "512"` + "\n"
	config += `	extended_anti_replay_window = 256` + "\n"
	config += `	ipsec_pairwise_keying = false` + "\n"
	config += `	integrity_type = ["esp"]` + "\n"
	config += `	keychains = [{` + "\n"
	config += `	  key_chain_name = "aaa"` + "\n"
	config += `	  key_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	keys = [{` + "\n"
	config += `	  id = 0` + "\n"
	config += `	  name = "aaa"` + "\n"
	config += `	  send_id = 1` + "\n"
	config += `	  receiver_id = 2` + "\n"
	config += `	  include_tcp_options = false` + "\n"
	config += `	  accept_ao_mismatch = false` + "\n"
	config += `	  crypto_algorithm = "aes-128-cmac"` + "\n"
	config += `	  key_string = "abcabc"` + "\n"
	config += `	  send_life_time_local = true` + "\n"
	config += `	  send_life_time_start_epoch = 1659284400` + "\n"
	config += `	  send_life_time_infinite = true` + "\n"
	config += `	  accept_life_time_local = true` + "\n"
	config += `	  accept_life_time_start_epoch = 1659284400` + "\n"
	config += `	  accept_life_time_infinite = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
