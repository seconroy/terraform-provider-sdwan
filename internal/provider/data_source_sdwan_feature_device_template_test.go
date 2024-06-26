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
func TestAccDataSourceSdwanFeatureDeviceTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_feature_device_template.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_feature_device_template.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_feature_device_template.test", "device_type", "vedge-ISR-4331"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_feature_device_template.test", "general_templates.0.type", "cisco_system"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanFeatureDeviceTemplatePrerequisitesConfig + testAccDataSourceSdwanFeatureDeviceTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanFeatureDeviceTemplatePrerequisitesConfig = `
resource "sdwan_cisco_system_feature_template" "system" {
  name = "TF_SYSTEM_1"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  hostname = "Router1"
  system_ip = "5.5.5.5"
  site_id = 1
  console_baud_rate = "115200"
  multi_tenant = true
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanFeatureDeviceTemplateConfig() string {
	config := ""
	config += `resource "sdwan_feature_device_template" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	device_type = "vedge-ISR-4331"` + "\n"
	config += `	general_templates = [{` + "\n"
	config += `	  id = sdwan_cisco_system_feature_template.system.id` + "\n"
	config += `	  type = "cisco_system"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_feature_device_template" "test" {
			id = sdwan_feature_device_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
