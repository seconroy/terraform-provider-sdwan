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
func TestAccDataSourceSdwanTLSSSLProfilePolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "mode", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "decrypt_threshold", "high-risk"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "reputation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_tls_ssl_profile_policy_definition.test", "fail_decrypt", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTLSSSLProfilePolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTLSSSLProfilePolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_tls_ssl_profile_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	mode = "security"` + "\n"
	config += `	decrypt_categories = ["alcohol-and-tobacco"]` + "\n"
	config += `	never_decrypt_categories = ["auctions"]` + "\n"
	config += `	skip_decrypt_categories = ["cdns"]` + "\n"
	config += `	decrypt_threshold = "high-risk"` + "\n"
	config += `	reputation = false` + "\n"
	config += `	fail_decrypt = true` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_tls_ssl_profile_policy_definition" "test" {
			id = sdwan_tls_ssl_profile_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
