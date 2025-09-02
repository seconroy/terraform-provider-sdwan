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
func TestAccDataSourceSdwanTopologyHubAndSpokeProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_hub_and_spoke_feature.test", "spoke_groups.0.name", "sj_11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_hub_and_spoke_feature.test", "spoke_groups.0.hub_preferences.0.preference", "30"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTopologyHubAndSpokePrerequisitesProfileParcelConfig + testAccDataSourceSdwanTopologyHubAndSpokeProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTopologyHubAndSpokePrerequisitesProfileParcelConfig = `
resource "sdwan_topology_feature_profile" "test" {
  name        = "TF_TEST"
  description = "My topology feature profile 1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTopologyHubAndSpokeProfileParcelConfig() string {
	config := `resource "sdwan_topology_hub_and_spoke_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	vpn = ["edge_basic_vpn1"]` + "\n"
	config += `	hub_sites = ["site_100"]` + "\n"
	config += `	spoke_groups = [{` + "\n"
	config += `	  name = "sj_11"` + "\n"
	config += `	  sites = ["site_100"]` + "\n"
	config += `	  hub_preferences = [{` + "\n"
	config += `		sites = ["site_1000"]` + "\n"
	config += `		preference = 30` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_topology_hub_and_spoke_feature" "test" {
			id = sdwan_topology_hub_and_spoke_feature.test.id
			feature_profile_id = sdwan_topology_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
