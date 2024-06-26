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
func TestAccSdwanCiscoThousandEyesFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.instance_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.application_type", "te"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_account_group_token", "1234567"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_agent_ip", "1.1.1.2/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_default_gateway", "1.1.1.255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_name_server", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_hostname", "agent1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_web_proxy_type", "static"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_proxy_host", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_proxy_port", "80"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoThousandEyesFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoThousandEyesFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoThousandEyesFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_thousandeyes_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoThousandEyesFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_thousandeyes_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	virtual_applications = [{` + "\n"
	config += `	  instance_id = "1"` + "\n"
	config += `	  application_type = "te"` + "\n"
	config += `	  te_account_group_token = "1234567"` + "\n"
	config += `	  te_vpn = 1` + "\n"
	config += `	  te_agent_ip = "1.1.1.2/24"` + "\n"
	config += `	  te_default_gateway = "1.1.1.255"` + "\n"
	config += `	  te_name_server = "10.2.2.2"` + "\n"
	config += `	  te_hostname = "agent1"` + "\n"
	config += `	  te_web_proxy_type = "static"` + "\n"
	config += `	  te_proxy_host = "3.3.3.3"` + "\n"
	config += `	  te_proxy_port = 80` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
