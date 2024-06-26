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
func TestAccDataSourceSdwanPreferredColorGroupPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "primary_color_preference", "blue bronze"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "primary_path_preference", "direct-path"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "secondary_color_preference", "3g"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "secondary_path_preference", "multi-hop-path"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "tertiary_color_preference", "custom1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_preferred_color_group_policy_object.test", "tertiary_path_preference", "all-paths"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPreferredColorGroupPolicyObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPreferredColorGroupPolicyObjectConfig() string {
	config := ""
	config += `resource "sdwan_preferred_color_group_policy_object" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	primary_color_preference = "blue bronze"` + "\n"
	config += `	primary_path_preference = "direct-path"` + "\n"
	config += `	secondary_color_preference = "3g"` + "\n"
	config += `	secondary_path_preference = "multi-hop-path"` + "\n"
	config += `	tertiary_color_preference = "custom1"` + "\n"
	config += `	tertiary_path_preference = "all-paths"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_preferred_color_group_policy_object" "test" {
			id = sdwan_preferred_color_group_policy_object.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
