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
func TestAccDataSourceSdwanSLAClassPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "jitter", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "latency", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "loss", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_criteria", "jitter-loss-latency"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_jitter", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_latency", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_loss", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSLAClassPolicyObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSLAClassPolicyObjectConfig() string {
	config := ""
	config += `resource "sdwan_sla_class_policy_object" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	jitter = 100` + "\n"
	config += `	latency = 10` + "\n"
	config += `	loss = 1` + "\n"
	config += `	fallback_best_tunnel_criteria = "jitter-loss-latency"` + "\n"
	config += `	fallback_best_tunnel_jitter = 100` + "\n"
	config += `	fallback_best_tunnel_latency = 10` + "\n"
	config += `	fallback_best_tunnel_loss = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_sla_class_policy_object" "test" {
			id = sdwan_sla_class_policy_object.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
