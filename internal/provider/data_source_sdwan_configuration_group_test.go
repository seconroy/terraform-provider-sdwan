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
func TestAccDataSourceSdwanConfigurationGroup(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group.test", "name", "CG_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group.test", "description", "My config group 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group.test", "solution", "sdwan"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanConfigurationGroupPrerequisitesConfig + testAccDataSourceSdwanConfigurationGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanConfigurationGroupPrerequisitesConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanConfigurationGroupConfig() string {
	config := ""
	config += `resource "sdwan_configuration_group" "test" {` + "\n"
	config += `	name = "CG_1"` + "\n"
	config += `	description = "My config group 1"` + "\n"
	config += `	solution = "sdwan"` + "\n"
	config += `	feature_profiles = [{` + "\n"
	config += `	  id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_configuration_group" "test" {
			id = sdwan_configuration_group.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
