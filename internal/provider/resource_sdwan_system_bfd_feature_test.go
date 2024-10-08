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
func TestAccSdwanSystemBFDProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "poll_interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "default_dscp", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "colors.0.color", "3g"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "colors.0.hello_interval", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "colors.0.multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "colors.0.pmtu_discovery", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_bfd_feature.test", "colors.0.dscp", "16"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemBFDPrerequisitesProfileParcelConfig + testAccSdwanSystemBFDProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemBFDPrerequisitesProfileParcelConfig + testAccSdwanSystemBFDProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemBFDPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemBFDProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_bfd_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemBFDProfileParcelConfig_all() string {
	config := `resource "sdwan_system_bfd_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	multiplier = 3` + "\n"
	config += `	poll_interval = 100` + "\n"
	config += `	default_dscp = 8` + "\n"
	config += `	colors = [{` + "\n"
	config += `	  color = "3g"` + "\n"
	config += `	  hello_interval = 200` + "\n"
	config += `	  multiplier = 3` + "\n"
	config += `	  pmtu_discovery = true` + "\n"
	config += `	  dscp = 16` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
