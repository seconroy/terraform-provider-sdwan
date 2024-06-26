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
func TestAccSdwanRoutePolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "default_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.name", "Sequence 10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.match_entries.0.type", "metric"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.match_entries.0.metric", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.type", "aggregator"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator_ip_address", "10.1.2.3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanRoutePolicyDefinitionConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanRoutePolicyDefinitionConfig_all() string {
	config := `resource "sdwan_route_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	default_action = "reject"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 10` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  name = "Sequence 10"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		type = "metric"` + "\n"
	config += `		metric = 100` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `		type = "aggregator"` + "\n"
	config += `		aggregator = 10` + "\n"
	config += `		aggregator_ip_address = "10.1.2.3"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
