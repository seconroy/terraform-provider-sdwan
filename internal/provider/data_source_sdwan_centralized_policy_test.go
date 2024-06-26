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
func TestAccDataSourceSdwanCentralizedPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_centralized_policy.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_centralized_policy.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_centralized_policy.test", "definitions.0.type", "data"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_centralized_policy.test", "definitions.0.entries.0.direction", "service"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCentralizedPolicyPrerequisitesConfig + testAccDataSourceSdwanCentralizedPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanCentralizedPolicyPrerequisitesConfig = `
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

resource "sdwan_traffic_data_policy_definition" "data1" {
  name           = "TF_TEST"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id      = 1
      name    = "Seq1"
      type    = "applicationFirewall"
      ip_type = "ipv4"
      action_entries = [
        {
          type = "log"
          log  = true
        }
      ]
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCentralizedPolicyConfig() string {
	config := ""
	config += `resource "sdwan_centralized_policy" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	definitions = [{` + "\n"
	config += `	  id = sdwan_traffic_data_policy_definition.data1.id` + "\n"
	config += `	  type = "data"` + "\n"
	config += `	  entries = [{` + "\n"
	config += `		site_list_ids = [sdwan_site_list_policy_object.sites1.id]` + "\n"
	config += `		vpn_list_ids = [sdwan_vpn_list_policy_object.vpns1.id]` + "\n"
	config += `		direction = "service"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_centralized_policy" "test" {
			id = sdwan_centralized_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
