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
func TestAccDataSourceSdwanHubAndSpokeTopologyPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "topologies.0.name", "Topology1"))
	if os.Getenv("SDWAN_2012") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "topologies.0.all_hubs_are_equal", "false"))
	}
	if os.Getenv("SDWAN_2012") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "topologies.0.advertise_hub_tlocs", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_hub_and_spoke_topology_policy_definition.test", "topologies.0.spokes.0.hubs.0.preference", "30"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanHubAndSpokeTopologyPolicyDefinitionPrerequisitesConfig + testAccDataSourceSdwanHubAndSpokeTopologyPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanHubAndSpokeTopologyPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_site_list_policy_object" "sites1" {
  name = "TF_TEST"
  entries = [
    {
      site_id = "100-200"
    }
  ]
}

resource "sdwan_vpn_list_policy_object" "vpns1" {
  name = "TF_TEST"
  entries = [
    {
      vpn_id = "100-200"
    }
  ]
}

resource "sdwan_tloc_list_policy_object" "tloc1" {
  name = "Example"
  entries = [
    {
      tloc_ip       = "1.1.1.2"
      color         = "blue"
      encapsulation = "gre"
      preference    = 10
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanHubAndSpokeTopologyPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_hub_and_spoke_topology_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	vpn_list_id = sdwan_vpn_list_policy_object.vpns1.id` + "\n"
	config += `	topologies = [{` + "\n"
	config += `	  name = "Topology1"` + "\n"
	if os.Getenv("SDWAN_2012") != "" {
		config += `	  all_hubs_are_equal = false` + "\n"
	}
	if os.Getenv("SDWAN_2012") != "" {
		config += `	  advertise_hub_tlocs = true` + "\n"
	}
	if os.Getenv("SDWAN_2012") != "" {
		config += `	  tloc_list_id = sdwan_tloc_list_policy_object.tloc1.id` + "\n"
	}
	config += `	  spokes = [{` + "\n"
	config += `		site_list_id = sdwan_site_list_policy_object.sites1.id` + "\n"
	config += `      hubs = [{` + "\n"
	config += `			site_list_id = sdwan_site_list_policy_object.sites1.id` + "\n"
	config += `			preference = "30"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_hub_and_spoke_topology_policy_definition" "test" {
			id = sdwan_hub_and_spoke_topology_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
