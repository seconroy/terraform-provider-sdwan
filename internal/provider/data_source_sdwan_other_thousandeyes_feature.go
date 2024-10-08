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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OtherThousandEyesProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &OtherThousandEyesProfileParcelDataSource{}
)

func NewOtherThousandEyesProfileParcelDataSource() datasource.DataSource {
	return &OtherThousandEyesProfileParcelDataSource{}
}

type OtherThousandEyesProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *OtherThousandEyesProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_other_thousandeyes_feature"
}

func (d *OtherThousandEyesProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Other ThousandEyes Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"virtual_application": schema.ListNestedAttribute{
				MarkdownDescription: "Virtual application Instance",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"account_group_token": schema.StringAttribute{
							MarkdownDescription: "Set the Account Group Token",
							Computed:            true,
						},
						"account_group_token_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "VPN number",
							Computed:            true,
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"management_ip": schema.StringAttribute{
							MarkdownDescription: "Set the Agent IP Address",
							Computed:            true,
						},
						"management_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"management_subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Set the Agent SubnetMask",
							Computed:            true,
						},
						"management_subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"agent_default_gateway": schema.StringAttribute{
							MarkdownDescription: "Set the Agent default gateway",
							Computed:            true,
						},
						"agent_default_gateway_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"name_server_ip": schema.StringAttribute{
							MarkdownDescription: "Set the name server",
							Computed:            true,
						},
						"name_server_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Set the host name",
							Computed:            true,
						},
						"hostname_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"proxy_type": schema.StringAttribute{
							MarkdownDescription: "Select Web Proxy Type",
							Computed:            true,
						},
						"proxy_host": schema.StringAttribute{
							MarkdownDescription: "Set the Proxy Host",
							Computed:            true,
						},
						"proxy_host_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"proxy_port": schema.Int64Attribute{
							MarkdownDescription: "Set the Proxy Port",
							Computed:            true,
						},
						"proxy_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"pac_url": schema.StringAttribute{
							MarkdownDescription: "Set the proxy PAC url",
							Computed:            true,
						},
						"pac_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OtherThousandEyesProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *OtherThousandEyesProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OtherThousandEyes

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
