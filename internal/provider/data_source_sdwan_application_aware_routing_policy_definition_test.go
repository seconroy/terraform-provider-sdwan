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
func TestAccDataSourceSdwanApplicationAwareRoutingPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.name", "Region1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.match_entries.0.type", "appList"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.action_entries.0.type", "backupSlaPreferredColor"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.action_entries.0.backup_sla_preferred_color", "bronze"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionPrerequisitesConfig + testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_application_list_policy_object" "test" {
  name = "TF_TEST"
  entries = [
    {
      application = "netflix"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_application_aware_routing_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  name = "Region1"` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		type = "appList"` + "\n"
	config += `		application_list_id = sdwan_application_list_policy_object.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `		type = "backupSlaPreferredColor"` + "\n"
	config += `		backup_sla_preferred_color = "bronze"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_application_aware_routing_policy_definition" "test" {
			id = sdwan_application_aware_routing_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
