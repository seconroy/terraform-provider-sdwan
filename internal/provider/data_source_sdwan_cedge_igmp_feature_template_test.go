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
func TestAccDataSourceSdwanCEdgeIGMPFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_igmp_feature_template.test", "interfaces.0.name", "Ethernet0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_igmp_feature_template.test", "interfaces.0.join_groups.0.group_address", "235.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_igmp_feature_template.test", "interfaces.0.join_groups.0.source", "1.2.3.4"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCEdgeIGMPFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCEdgeIGMPFeatureTemplateConfig() string {
	config := `resource "sdwan_cedge_igmp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  name = "Ethernet0"` + "\n"
	config += `	  join_groups = [{` + "\n"
	config += `		group_address = "235.1.1.1"` + "\n"
	config += `		source = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cedge_igmp_feature_template" "test" {
			id = sdwan_cedge_igmp_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
