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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanCiscoOSPFFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "auto_cost_reference_bandwidth", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "compatible_rfc1583", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "default_information_originate_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "default_information_originate_metric", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "default_information_originate_metric_type", "type1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "distance_external", "111"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "distance_inter_area", "111"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "distance_intra_area", "112"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "timers_spf_delay", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "timers_spf_initial_hold", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "timers_spf_max_hold", "20000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "redistribute.0.protocol", "static"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "redistribute.0.route_policy", "RP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "redistribute.0.nat_dia", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "max_metric_router_lsa.0.ad_type", "on-startup"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "max_metric_router_lsa.0.time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "route_policies.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "route_policies.0.policy_name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.area_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.stub", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.stub_no_summary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.nssa", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.nssa_no_summary", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.name", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.hello_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.dead_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.retransmit_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.cost", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.network", "point-to-point"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.passive_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.authentication_type", "message-digest"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.authentication_message_digest_key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.interfaces.0.authentication_message_digest_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.ranges.0.address", "1.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.ranges.0.cost", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ospf_feature_template.test", "areas.0.ranges.0.no_advertise", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoOSPFFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoOSPFFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoOSPFFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_ospf_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoOSPFFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_ospf_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	auto_cost_reference_bandwidth = 100000` + "\n"
	config += `	compatible_rfc1583 = true` + "\n"
	config += `	default_information_originate = true` + "\n"
	config += `	default_information_originate_always = true` + "\n"
	config += `	default_information_originate_metric = 100` + "\n"
	config += `	default_information_originate_metric_type = "type1"` + "\n"
	config += `	distance_external = 111` + "\n"
	config += `	distance_inter_area = 111` + "\n"
	config += `	distance_intra_area = 112` + "\n"
	config += `	timers_spf_delay = 300` + "\n"
	config += `	timers_spf_initial_hold = 2000` + "\n"
	config += `	timers_spf_max_hold = 20000` + "\n"
	config += `	redistribute = [{` + "\n"
	config += `	  protocol = "static"` + "\n"
	config += `	  route_policy = "RP1"` + "\n"
	config += `	  nat_dia = true` + "\n"
	config += `	}]` + "\n"
	config += `	max_metric_router_lsa = [{` + "\n"
	config += `	  ad_type = "on-startup"` + "\n"
	config += `	  time = 100` + "\n"
	config += `	}]` + "\n"
	config += `	route_policies = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  policy_name = "POLICY1"` + "\n"
	config += `	}]` + "\n"
	config += `	areas = [{` + "\n"
	config += `	  area_number = 1` + "\n"
	config += `	  stub = false` + "\n"
	config += `	  stub_no_summary = false` + "\n"
	config += `	  nssa = false` + "\n"
	config += `	  nssa_no_summary = true` + "\n"
	config += `	  interfaces = [{` + "\n"
	config += `		name = "e1"` + "\n"
	config += `		hello_interval = 20` + "\n"
	config += `		dead_interval = 60` + "\n"
	config += `		retransmit_interval = 10` + "\n"
	config += `		cost = 100` + "\n"
	config += `		priority = 10` + "\n"
	config += `		network = "point-to-point"` + "\n"
	config += `		passive_interface = true` + "\n"
	config += `		authentication_type = "message-digest"` + "\n"
	config += `		authentication_message_digest_key_id = 1` + "\n"
	config += `		authentication_message_digest_key = "cisco123"` + "\n"
	config += `	}]` + "\n"
	config += `	  ranges = [{` + "\n"
	config += `		address = "1.1.1.0/24"` + "\n"
	config += `		cost = 100` + "\n"
	config += `		no_advertise = true` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
