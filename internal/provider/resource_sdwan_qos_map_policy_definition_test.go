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
func TestAccSdwanQoSMapPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.queue", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.bandwidth_percent", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.buffer_percent", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.burst", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.drop_type", "red-drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_qos_map_policy_definition.test", "qos_schedulers.0.scheduling_type", "wrr"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanQoSMapPolicyDefinitionPrerequisitesConfig + testAccSdwanQoSMapPolicyDefinitionConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanQoSMapPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_class_map_policy_object" "test" {
  name = "TF_TEST"
  queue = 6
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanQoSMapPolicyDefinitionConfig_all() string {
	config := `resource "sdwan_qos_map_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	qos_schedulers = [{` + "\n"
	config += `	  queue = 6` + "\n"
	config += `	  class_map_id = sdwan_class_map_policy_object.test.id` + "\n"
	config += `	  bandwidth_percent = 10` + "\n"
	config += `	  buffer_percent = 10` + "\n"
	config += `	  burst = 100000` + "\n"
	config += `	  drop_type = "red-drop"` + "\n"
	config += `	  scheduling_type = "wrr"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
