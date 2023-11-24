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

func TestAccDataSourceSdwanServiceFeatureProfile(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceFeatureProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_service_feature_profile.test", "name", "SERVICE_FP_1"),
					resource.TestCheckResourceAttr("data.sdwan_service_feature_profile.test", "description", "My service feature profile 1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanServiceFeatureProfileConfig = `

resource "sdwan_service_feature_profile" "test" {
  name = "SERVICE_FP_1"
  description = "My service feature profile 1"
}

data "sdwan_service_feature_profile" "test" {
  id = sdwan_service_feature_profile.test.id
}
`
