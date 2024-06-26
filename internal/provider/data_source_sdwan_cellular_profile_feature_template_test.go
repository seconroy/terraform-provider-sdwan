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
func TestAccDataSourceSdwanCellularProfileFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "if_name", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "profile_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "access_point_name", "APN1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "authentication_type", "CHAP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "profile_name", "PROFILE1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "packet_data_network_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "profile_username", "MyUsername"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "profile_password", "MyPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "primary_dns_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cellular_profile_feature_template.test", "secondary_dns_address", "1.2.3.5"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCellularProfileFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCellularProfileFeatureTemplateConfig() string {
	config := `resource "sdwan_cellular_profile_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	if_name = "Ethernet1"` + "\n"
	config += `	profile_id = 1` + "\n"
	config += `	access_point_name = "APN1"` + "\n"
	config += `	authentication_type = "CHAP"` + "\n"
	config += `	ip_address = "1.2.3.4"` + "\n"
	config += `	profile_name = "PROFILE1"` + "\n"
	config += `	packet_data_network_type = "ipv4"` + "\n"
	config += `	profile_username = "MyUsername"` + "\n"
	config += `	profile_password = "MyPassword"` + "\n"
	config += `	primary_dns_address = "1.2.3.4"` + "\n"
	config += `	secondary_dns_address = "1.2.3.5"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cellular_profile_feature_template" "test" {
			id = sdwan_cellular_profile_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
