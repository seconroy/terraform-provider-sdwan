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
	_ datasource.DataSource              = &CflowdPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &CflowdPolicyDefinitionDataSource{}
)

func NewCflowdPolicyDefinitionDataSource() datasource.DataSource {
	return &CflowdPolicyDefinitionDataSource{}
}

type CflowdPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *CflowdPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cflowd_policy_definition"
}

func (d *CflowdPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cflowd Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"active_flow_timeout": schema.Int64Attribute{
				MarkdownDescription: "Active flow timeout in seconds",
				Computed:            true,
			},
			"inactive_flow_timeout": schema.Int64Attribute{
				MarkdownDescription: "Inactive flow timeout in seconds",
				Computed:            true,
			},
			"sampling_interval": schema.Int64Attribute{
				MarkdownDescription: "Flow sampling interval",
				Computed:            true,
			},
			"flow_refresh": schema.Int64Attribute{
				MarkdownDescription: "Flow refresh in seconds",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Protocol, either `ipv4`, `ipv6` or `all`",
				Computed:            true,
			},
			"tos": schema.BoolAttribute{
				MarkdownDescription: "Collect TOS record field",
				Computed:            true,
			},
			"remarked_dscp": schema.BoolAttribute{
				MarkdownDescription: "Collect remarked DSCP",
				Computed:            true,
			},
			"collectors": schema.ListNestedAttribute{
				MarkdownDescription: "List of collectors",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN ID",
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"transport": schema.StringAttribute{
							MarkdownDescription: "Transport protocol",
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Source interface",
							Computed:            true,
						},
						"export_spreading": schema.StringAttribute{
							MarkdownDescription: "Export spreading",
							Computed:            true,
						},
						"bfd_metrics_exporting": schema.BoolAttribute{
							MarkdownDescription: "BFD metrics exporting",
							Computed:            true,
						},
						"exporting_interval": schema.Int64Attribute{
							MarkdownDescription: "Exporting interval",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CflowdPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CflowdPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CflowdPolicyDefinition

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
	config.Type = types.StringValue("cflowd")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
