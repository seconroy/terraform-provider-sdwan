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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DNSSecurityPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &DNSSecurityPolicyDefinitionDataSource{}
)

func NewDNSSecurityPolicyDefinitionDataSource() datasource.DataSource {
	return &DNSSecurityPolicyDefinitionDataSource{}
}

type DNSSecurityPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *DNSSecurityPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_security_policy_definition"
}

func (d *DNSSecurityPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the DNS Security Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition.",
				Computed:            true,
			},
			"domain_list_id": schema.StringAttribute{
				MarkdownDescription: "Local domain bypass list ID",
				Computed:            true,
			},
			"domain_list_version": schema.Int64Attribute{
				MarkdownDescription: "Local domain bypass list version",
				Computed:            true,
			},
			"local_domain_bypass_enabled": schema.BoolAttribute{
				MarkdownDescription: "Should the local domain bypass list be enabled",
				Computed:            true,
			},
			"match_all_vpn": schema.BoolAttribute{
				MarkdownDescription: "Should use match all VPN",
				Computed:            true,
			},
			"target_vpns": schema.ListNestedAttribute{
				MarkdownDescription: "Only relevant when `match_all_vpn` is `false`",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_ids": schema.SetAttribute{
							MarkdownDescription: "VPN ID's separated by Comma",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"umbrella_dns_default": schema.BoolAttribute{
							MarkdownDescription: "Should use umbrella as DNS Server",
							Computed:            true,
						},
						"custom_dns_server_ip": schema.StringAttribute{
							MarkdownDescription: "Only relevant when `umbrella_dns_default` is `false`",
							Computed:            true,
						},
						"local_domain_bypass_enabled": schema.BoolAttribute{
							MarkdownDescription: "Should the local domain bypass list be enabled",
							Computed:            true,
						},
					},
				},
			},
			"dnscrypt": schema.BoolAttribute{
				MarkdownDescription: "Should DNSCrypt be enabled",
				Computed:            true,
			},
			"umbrella_dns_default": schema.BoolAttribute{
				MarkdownDescription: "Should use umbrella as DNS Server",
				Computed:            true,
			},
			"custom_dns_server_ip": schema.StringAttribute{
				MarkdownDescription: "Only relevant when `umbrella_dns_default` is `false`",
				Computed:            true,
			},
			"cisco_sig_credentials_feature_template_id": schema.StringAttribute{
				MarkdownDescription: "Credentials feature template ID",
				Computed:            true,
			},
			"cisco_sig_credentials_feature_template_version": schema.Int64Attribute{
				MarkdownDescription: "Credentials feature template version",
				Computed:            true,
			},
		},
	}
}

func (d *DNSSecurityPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *DNSSecurityPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DNSSecurityPolicyDefinition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
