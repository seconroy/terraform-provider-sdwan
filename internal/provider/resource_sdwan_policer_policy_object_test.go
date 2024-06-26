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
func TestAccSdwanPolicerPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policer_policy_object.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policer_policy_object.test", "burst", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policer_policy_object.test", "exceed_action", "remark"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policer_policy_object.test", "rate", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanPolicerPolicyObjectConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanPolicerPolicyObjectConfig_all() string {
	config := `resource "sdwan_policer_policy_object" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	burst = 100000` + "\n"
	config += `	exceed_action = "remark"` + "\n"
	config += `	rate = 100` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
