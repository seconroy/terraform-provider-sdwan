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

func TestAccSdwanCiscoSIGCredentialsFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_organization", "org1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_base_uri", "abc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_username", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_cloud_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_username", "partner1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_api_key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_secret", "secret123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_organization_id", "org1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_sig_credentials_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_sig_credentials_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	zscaler_organization = "org1"` + "\n"
	config += `	zscaler_partner_base_uri = "abc"` + "\n"
	config += `	zscaler_username = "user1"` + "\n"
	config += `	zscaler_password = "password123"` + "\n"
	config += `	zscaler_cloud_name = 1` + "\n"
	config += `	zscaler_partner_username = "partner1"` + "\n"
	config += `	zscaler_partner_password = "password123"` + "\n"
	config += `	zscaler_partner_api_key = "key123"` + "\n"
	config += `	umbrella_api_key = "key123"` + "\n"
	config += `	umbrella_api_secret = "secret123"` + "\n"
	config += `	umbrella_organization_id = "org1"` + "\n"
	config += `}` + "\n"
	return config
}
