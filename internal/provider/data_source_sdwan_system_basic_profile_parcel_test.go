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
func TestAccDataSourceSdwanSystemBasicProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "timezone", "UTC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "config_description", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "location", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_longitude", "-77"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_latitude", "38"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_geo_fencing_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_geo_fencing_range", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_sms_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "gps_sms_mobile_numbers.0.number", "+11111233"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "overlay_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "port_offset", "19"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "port_hopping", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "control_session_pps", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "track_transport", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "track_interface_tag", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "console_baud_rate", "9600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "max_omp_sessions", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "multi_tenant", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "track_default_gateway", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "admin_tech_on_failure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "idle_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "on_demand_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "on_demand_idle_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "transport_gateway", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "enhanced_app_aware_routing", "aggressive"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "affinity_group_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "affinity_preference_auto", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "affinity_per_vrfs.0.affinity_group_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_basic_profile_parcel.test", "affinity_per_vrfs.0.vrf_range", "123-456"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemBasicPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemBasicProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemBasicPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemBasicProfileParcelConfig() string {
	config := `resource "sdwan_system_basic_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	timezone = "UTC"` + "\n"
	config += `	config_description = "example"` + "\n"
	config += `	location = "example"` + "\n"
	config += `	gps_longitude = -77` + "\n"
	config += `	gps_latitude = 38` + "\n"
	config += `	gps_geo_fencing_enable = true` + "\n"
	config += `	gps_geo_fencing_range = 100` + "\n"
	config += `	gps_sms_enable = true` + "\n"
	config += `	gps_sms_mobile_numbers = [{` + "\n"
	config += `	  number = "+11111233"` + "\n"
	config += `	}]` + "\n"
	config += `	device_groups = ["example"]` + "\n"
	config += `	controller_groups = [1]` + "\n"
	config += `	overlay_id = 1` + "\n"
	config += `	port_offset = 19` + "\n"
	config += `	port_hopping = true` + "\n"
	config += `	control_session_pps = 300` + "\n"
	config += `	track_transport = true` + "\n"
	config += `	track_interface_tag = 2` + "\n"
	config += `	console_baud_rate = "9600"` + "\n"
	config += `	max_omp_sessions = 24` + "\n"
	config += `	multi_tenant = false` + "\n"
	config += `	track_default_gateway = true` + "\n"
	config += `	admin_tech_on_failure = true` + "\n"
	config += `	idle_timeout = 10` + "\n"
	config += `	on_demand_enable = true` + "\n"
	config += `	on_demand_idle_timeout = 10` + "\n"
	config += `	transport_gateway = false` + "\n"
	config += `	enhanced_app_aware_routing = "aggressive"` + "\n"
	config += `	site_types = ["type-1"]` + "\n"
	config += `	affinity_group_number = 1` + "\n"
	config += `	affinity_group_preferences = [1]` + "\n"
	config += `	affinity_preference_auto = false` + "\n"
	config += `	affinity_per_vrfs = [{` + "\n"
	config += `	  affinity_group_number = 1` + "\n"
	config += `	  vrf_range = "123-456"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_basic_profile_parcel" "test" {
			id = sdwan_system_basic_profile_parcel.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
