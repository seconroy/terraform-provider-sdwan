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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanTransportRoutingOSPFProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "reference_bandwidth", "101"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "rfc_1583_compatible", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "default_information_originate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "default_information_originate_always", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "default_information_originate_metric", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "default_information_originate_metric_type", "type1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "distance_external", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "distance_inter_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "distance_intra_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "spf_calculation_delay", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "spf_initial_hold_time", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "spf_maximum_hold_time", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "redistributes.0.protocol", "static"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "redistributes.0.nat_dia", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "router_lsas.0.type", "on-startup"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "router_lsas.0.time", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.area_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.area_type", "stub"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.no_summary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.name", "GigabitEthernet2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.dead_interval", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.lsa_retransmit_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.designated_router_priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.network_type", "broadcast"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.passive_interface", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.authentication_type", "message-digest"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.interfaces.0.message_digest_key_id", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.ranges.0.ip_address", "10.1.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.ranges.0.subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.ranges.0.cost", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_routing_ospf_feature.test", "areas.0.ranges.0.no_advertise", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportRoutingOSPFPrerequisitesProfileParcelConfig + testAccSdwanTransportRoutingOSPFProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportRoutingOSPFPrerequisitesProfileParcelConfig + testAccSdwanTransportRoutingOSPFProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportRoutingOSPFPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportRoutingOSPFProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_routing_ospf_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportRoutingOSPFProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_routing_ospf_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	reference_bandwidth = 101` + "\n"
	config += `	rfc_1583_compatible = true` + "\n"
	config += `	default_information_originate = false` + "\n"
	config += `	default_information_originate_always = false` + "\n"
	config += `	default_information_originate_metric = 1` + "\n"
	config += `	default_information_originate_metric_type = "type1"` + "\n"
	config += `	distance_external = 110` + "\n"
	config += `	distance_inter_area = 110` + "\n"
	config += `	distance_intra_area = 110` + "\n"
	config += `	spf_calculation_delay = 200` + "\n"
	config += `	spf_initial_hold_time = 1000` + "\n"
	config += `	spf_maximum_hold_time = 10000` + "\n"
	config += `	redistributes = [{` + "\n"
	config += `	  protocol = "static"` + "\n"
	config += `	  nat_dia = true` + "\n"
	config += `	}]` + "\n"
	config += `	router_lsas = [{` + "\n"
	config += `	  type = "on-startup"` + "\n"
	config += `	  time = 5` + "\n"
	config += `	}]` + "\n"
	config += `	areas = [{` + "\n"
	config += `	  area_number = 1` + "\n"
	config += `	  area_type = "stub"` + "\n"
	config += `	  no_summary = false` + "\n"
	config += `	  interfaces = [{` + "\n"
	config += `		name = "GigabitEthernet2"` + "\n"
	config += `		hello_interval = 10` + "\n"
	config += `		dead_interval = 40` + "\n"
	config += `		lsa_retransmit_interval = 5` + "\n"
	config += `		cost = 10` + "\n"
	config += `		designated_router_priority = 1` + "\n"
	config += `		network_type = "broadcast"` + "\n"
	config += `		passive_interface = false` + "\n"
	config += `		authentication_type = "message-digest"` + "\n"
	config += `		message_digest_key_id = 7` + "\n"
	config += `		message_digest_key = "sdjfhsghbjdjr"` + "\n"
	config += `	}]` + "\n"
	config += `	  ranges = [{` + "\n"
	config += `		ip_address = "10.1.1.0"` + "\n"
	config += `		subnet_mask = "255.255.255.0"` + "\n"
	config += `		cost = 1` + "\n"
	config += `		no_advertise = false` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
