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
func TestAccSdwanSecurityAppHostingFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.instance_id", "2e89c1fe-440a-43f5-9f3a-54a9836fdbb5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.application_type", "utd"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.database_url", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.resource_profile", "low"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.service_gateway_ip", "1.2.3.4/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.service_ip", "1.2.3.5/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.data_gateway_ip", "192.0.2.1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_security_app_hosting_feature_template.test", "virtual_applications.0.data_service_ip", "192.0.2.2/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSecurityAppHostingFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanSecurityAppHostingFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSecurityAppHostingFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_security_app_hosting_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSecurityAppHostingFeatureTemplateConfig_all() string {
	config := `resource "sdwan_security_app_hosting_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	virtual_applications = [{` + "\n"
	config += `	  instance_id = "2e89c1fe-440a-43f5-9f3a-54a9836fdbb5"` + "\n"
	config += `	  application_type = "utd"` + "\n"
	config += `	  nat = true` + "\n"
	config += `	  database_url = false` + "\n"
	config += `	  resource_profile = "low"` + "\n"
	config += `	  service_gateway_ip = "1.2.3.4/24"` + "\n"
	config += `	  service_ip = "1.2.3.5/24"` + "\n"
	config += `	  data_gateway_ip = "192.0.2.1/24"` + "\n"
	config += `	  data_service_ip = "192.0.2.2/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
