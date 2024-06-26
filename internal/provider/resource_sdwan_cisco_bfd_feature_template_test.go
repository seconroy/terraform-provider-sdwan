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
func TestAccSdwanCiscoBFDFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "poll_interval", "800000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "default_dscp", "48"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.color", "private5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.multiplier", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.pmtu_discovery", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.dscp", "46"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoBFDFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoBFDFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoBFDFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_bfd_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoBFDFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_bfd_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	multiplier = 3` + "\n"
	config += `	poll_interval = 800000` + "\n"
	config += `	default_dscp = 48` + "\n"
	config += `	colors = [{` + "\n"
	config += `	  color = "private5"` + "\n"
	config += `	  hello_interval = 1000` + "\n"
	config += `	  multiplier = 7` + "\n"
	config += `	  pmtu_discovery = true` + "\n"
	config += `	  dscp = 46` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
