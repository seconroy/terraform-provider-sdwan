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
func TestAccDataSourceSdwanVEdgeInventory(t *testing.T) {
	if os.Getenv("SDWAN_209") == "" {
		t.Skip("skipping test, set environment variable SDWAN_209")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vedge_inventory.test", "devices.0.chassis_number", "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vedge_inventory.test", "devices.0.serial_number", "52FD47D8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vedge_inventory.test", "devices.0.device_type", "vedge"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanVEdgeInventoryConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanVEdgeInventoryConfig() string {
	config := ""

	config += `
		data "sdwan_vedge_inventory" "test" {
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
