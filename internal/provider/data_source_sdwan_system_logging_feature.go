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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SystemLoggingProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemLoggingProfileParcelDataSource{}
)

func NewSystemLoggingProfileParcelDataSource() datasource.DataSource {
	return &SystemLoggingProfileParcelDataSource{}
}

type SystemLoggingProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemLoggingProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_logging_feature"
}

func (d *SystemLoggingProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System Logging Feature.",

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
			"disk_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable logging to local disk",
				Computed:            true,
			},
			"disk_enable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"disk_file_size": schema.Int64Attribute{
				MarkdownDescription: "Set maximum size of file before it is rotated",
				Computed:            true,
			},
			"disk_file_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"disk_file_rotate": schema.Int64Attribute{
				MarkdownDescription: "Set number of syslog files to create before discarding oldest files",
				Computed:            true,
			},
			"disk_file_rotate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tls_profiles": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a TLS profile",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"profile": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the TLS profile",
							Computed:            true,
						},
						"profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_version": schema.StringAttribute{
							MarkdownDescription: "TLS Version",
							Computed:            true,
						},
						"tls_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"cipher_suites": schema.SetAttribute{
							MarkdownDescription: "Syslog secure server ciphersuites",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"cipher_suites_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_servers": schema.ListNestedAttribute{
				MarkdownDescription: "Enable logging to remote server",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip": schema.StringAttribute{
							MarkdownDescription: "Set hostname or IPv4 address of server",
							Computed:            true,
						},
						"hostname_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Set hostname or IPv4 address of server",
							Computed:            true,
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set interface to use to reach syslog server",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Set logging level for messages logged to server",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_enable": schema.BoolAttribute{
							MarkdownDescription: "Enable TLS Profile",
							Computed:            true,
						},
						"tls_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_properties_custom_profile": schema.BoolAttribute{
							MarkdownDescription: "Define custom profile",
							Computed:            true,
						},
						"tls_properties_custom_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_properties_profile": schema.StringAttribute{
							MarkdownDescription: "Configure a TLS profile",
							Computed:            true,
						},
						"tls_properties_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_servers": schema.ListNestedAttribute{
				MarkdownDescription: "Enable logging to remote ipv6 server",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip": schema.StringAttribute{
							MarkdownDescription: "Set IPv6 hostname or IPv6 address of server",
							Computed:            true,
						},
						"hostname_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Set hostname or IPv4 address of server",
							Computed:            true,
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set interface to use to reach syslog server",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Set logging level for messages logged to server",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_enable": schema.BoolAttribute{
							MarkdownDescription: "Enable TLS Profile",
							Computed:            true,
						},
						"tls_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_properties_custom_profile": schema.BoolAttribute{
							MarkdownDescription: "Define custom profile",
							Computed:            true,
						},
						"tls_properties_custom_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tls_properties_profile": schema.StringAttribute{
							MarkdownDescription: "Configure a TLS profile",
							Computed:            true,
						},
						"tls_properties_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemLoggingProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemLoggingProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemLogging

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
