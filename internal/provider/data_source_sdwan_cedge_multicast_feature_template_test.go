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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanCEdgeMulticastFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCEdgeMulticastFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cedge_multicast_feature_template.test", "spt_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_multicast_feature_template.test", "local_replicator", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_multicast_feature_template.test", "threshold", "200"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCEdgeMulticastFeatureTemplateConfig = `

resource "sdwan_cedge_multicast_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  spt_only = true
  local_replicator = true
  threshold = 200
}

data "sdwan_cedge_multicast_feature_template" "test" {
  id = sdwan_cedge_multicast_feature_template.test.id
}
`