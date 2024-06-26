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
func TestAccDataSourceSdwanCiscoDHCPServerFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "address_pool", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "lease_time", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "domain_name", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "default_gateway", "10.1.1.254"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.mac_address", "11:11:11:11:11:11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.ip_address", "10.1.1.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.hostname", "HOST1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "options.0.option_code", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_dhcp_server_feature_template.test", "options.0.ascii", "abc"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoDHCPServerFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoDHCPServerFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_dhcp_server_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	address_pool = "10.1.1.0/24"` + "\n"
	config += `	exclude_addresses = ["10.1.1.1-10.1.1.5", "10.1.1.254"]` + "\n"
	config += `	lease_time = 600` + "\n"
	config += `	interface_mtu = 1500` + "\n"
	config += `	domain_name = "cisco.com"` + "\n"
	config += `	default_gateway = "10.1.1.254"` + "\n"
	config += `	dns_servers = ["1.2.3.4"]` + "\n"
	config += `	tftp_servers = ["1.2.3.4"]` + "\n"
	config += `	static_leases = [{` + "\n"
	config += `	  mac_address = "11:11:11:11:11:11"` + "\n"
	config += `	  ip_address = "10.1.1.10"` + "\n"
	config += `	  hostname = "HOST1"` + "\n"
	config += `	}]` + "\n"
	config += `	options = [{` + "\n"
	config += `	  option_code = 10` + "\n"
	config += `	  ascii = "abc"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_dhcp_server_feature_template" "test" {
			id = sdwan_cisco_dhcp_server_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
