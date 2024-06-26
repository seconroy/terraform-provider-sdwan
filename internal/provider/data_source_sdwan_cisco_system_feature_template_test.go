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
func TestAccDataSourceSdwanCiscoSystemFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "timezone", "UTC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "hostname", "Router1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "system_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "location", "Building 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "latitude", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "longitude", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_range", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_sms", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_sms_phone_numbers.0.number", "+1234567"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "system_ip", "5.5.5.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "overlay_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "site_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "port_offset", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "port_hopping", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "control_session_pps", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_transport", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_interface_tag", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "console_baud_rate", "115200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "max_omp_sessions", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "multi_tenant", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_default_gateway", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "admin_tech_on_failure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "idle_timeout", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_ip", "5.6.7.8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_dns_name", "abc.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_api_url", "https://1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.boolean", "or"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.threshold", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.type", "interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.object_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.sig", "sig1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.ip", "6.6.6.6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.group_tracks_ids.0.track_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.boolean", "and"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "on_demand_tunnel", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "on_demand_tunnel_idle_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "affinity_group_number", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "transport_gateway", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "enable_mrf_migration", "enabled"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "migration_bgp_community", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoSystemFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoSystemFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_system_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	timezone = "UTC"` + "\n"
	config += `	hostname = "Router1"` + "\n"
	config += `	system_description = "My Description"` + "\n"
	config += `	location = "Building 1"` + "\n"
	config += `	latitude = 40` + "\n"
	config += `	longitude = 50` + "\n"
	config += `	geo_fencing = true` + "\n"
	config += `	geo_fencing_range = 1000` + "\n"
	config += `	geo_fencing_sms = true` + "\n"
	config += `	geo_fencing_sms_phone_numbers = [{` + "\n"
	config += `	  number = "+1234567"` + "\n"
	config += `	}]` + "\n"
	config += `	device_groups = ["group1"]` + "\n"
	config += `	controller_group_list = [1]` + "\n"
	config += `	system_ip = "5.5.5.5"` + "\n"
	config += `	overlay_id = 1` + "\n"
	config += `	site_id = 1` + "\n"
	config += `	port_offset = 1` + "\n"
	config += `	port_hopping = true` + "\n"
	config += `	control_session_pps = 300` + "\n"
	config += `	track_transport = true` + "\n"
	config += `	track_interface_tag = 1` + "\n"
	config += `	console_baud_rate = "115200"` + "\n"
	config += `	max_omp_sessions = 5` + "\n"
	config += `	multi_tenant = true` + "\n"
	config += `	track_default_gateway = true` + "\n"
	config += `	admin_tech_on_failure = true` + "\n"
	config += `	idle_timeout = 100` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_ip = "5.6.7.8"` + "\n"
	config += `	  endpoint_dns_name = "abc.com"` + "\n"
	config += `	  endpoint_api_url = "https://1.1.1.1"` + "\n"
	config += `	  elements = ["abc", "def"]` + "\n"
	config += `	  boolean = "or"` + "\n"
	config += `	  threshold = 300` + "\n"
	config += `	  interval = 60` + "\n"
	config += `	  multiplier = 3` + "\n"
	config += `	  type = "interface"` + "\n"
	config += `	}]` + "\n"
	config += `	object_trackers = [{` + "\n"
	config += `	  object_number = 1` + "\n"
	config += `	  interface = "e1"` + "\n"
	config += `	  sig = "sig1"` + "\n"
	config += `	  ip = "6.6.6.6"` + "\n"
	config += `	  mask = "0.0.0.0"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  group_tracks_ids = [{` + "\n"
	config += `		track_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	  boolean = "and"` + "\n"
	config += `	}]` + "\n"
	config += `	on_demand_tunnel = true` + "\n"
	config += `	on_demand_tunnel_idle_timeout = 10` + "\n"
	config += `	affinity_group_number = 5` + "\n"
	config += `	affinity_group_preference = [1]` + "\n"
	config += `	transport_gateway = true` + "\n"
	config += `	enable_mrf_migration = "enabled"` + "\n"
	config += `	migration_bgp_community = 100` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_system_feature_template" "test" {
			id = sdwan_cisco_system_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
