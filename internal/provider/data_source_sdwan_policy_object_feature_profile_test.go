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
func TestAccDataSourceSdwanPolicyObjectFeatureProfile(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_feature_profile.test", "name", "POLICY_OBJECT_FP_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_feature_profile.test", "description", "My policy object feature profile 1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectFeatureProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

func testAccDataSourceSdwanPolicyObjectFeatureProfileConfig() string {
	config := ""
	config += `resource "sdwan_policy_object_feature_profile" "test" {` + "\n"
	config += `	name = "POLICY_OBJECT_FP_1"` + "\n"
	config += `	description = "My policy object feature profile 1"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_feature_profile" "test" {
			depends_on = [sdwan_policy_object_feature_profile.test]
		}
	`
	return config
}