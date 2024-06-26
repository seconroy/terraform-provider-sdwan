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
func TestAccDataSourceSdwanCiscoSNMPFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "contact", "Max"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "location", "Building 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "views.0.name", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "views.0.object_identifiers.0.id", "1.2.3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "views.0.object_identifiers.0.exclude", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "communities.0.name", "community1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "communities.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "communities.0.authorization", "read-only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "groups.0.name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "groups.0.security_level", "auth-priv"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "groups.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.name", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.authentication_protocol", "sha"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.authentication_password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.privacy_protocol", "aes-cfb-128"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.privacy_password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "users.0.group", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.udp_port", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.community_name", "community1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.user", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_snmp_feature_template.test", "trap_targets.0.source_interface", "e1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoSNMPFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoSNMPFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_snmp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	contact = "Max"` + "\n"
	config += `	location = "Building 1"` + "\n"
	config += `	views = [{` + "\n"
	config += `	  name = "VIEW1"` + "\n"
	config += `	  object_identifiers = [{` + "\n"
	config += `		id = "1.2.3"` + "\n"
	config += `		exclude = true` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	communities = [{` + "\n"
	config += `	  name = "community1"` + "\n"
	config += `	  view = "VIEW1"` + "\n"
	config += `	  authorization = "read-only"` + "\n"
	config += `	}]` + "\n"
	config += `	groups = [{` + "\n"
	config += `	  name = "GROUP1"` + "\n"
	config += `	  security_level = "auth-priv"` + "\n"
	config += `	  view = "VIEW1"` + "\n"
	config += `	}]` + "\n"
	config += `	users = [{` + "\n"
	config += `	  name = "user1"` + "\n"
	config += `	  authentication_protocol = "sha"` + "\n"
	config += `	  authentication_password = "password123"` + "\n"
	config += `	  privacy_protocol = "aes-cfb-128"` + "\n"
	config += `	  privacy_password = "password123"` + "\n"
	config += `	  group = "GROUP1"` + "\n"
	config += `	}]` + "\n"
	config += `	trap_targets = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  ip = "1.1.1.1"` + "\n"
	config += `	  udp_port = 12345` + "\n"
	config += `	  community_name = "community1"` + "\n"
	config += `	  user = "user1"` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_snmp_feature_template" "test" {
			id = sdwan_cisco_snmp_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
