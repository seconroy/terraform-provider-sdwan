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
	_ datasource.DataSource              = &SystemOMPProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemOMPProfileParcelDataSource{}
)

func NewSystemOMPProfileParcelDataSource() datasource.DataSource {
	return &SystemOMPProfileParcelDataSource{}
}

type SystemOMPProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemOMPProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_omp_feature"
}

func (d *SystemOMPProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System OMP Feature.",

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
			"graceful_restart": schema.BoolAttribute{
				MarkdownDescription: "Graceful Restart for OMP",
				Computed:            true,
			},
			"graceful_restart_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"overlay_as": schema.Int64Attribute{
				MarkdownDescription: "Overlay AS Number",
				Computed:            true,
			},
			"overlay_as_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"paths_advertised_per_prefix": schema.Int64Attribute{
				MarkdownDescription: "Number of Paths Advertised per Prefix",
				Computed:            true,
			},
			"paths_advertised_per_prefix_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ecmp_limit": schema.Int64Attribute{
				MarkdownDescription: "Set maximum number of OMP paths to install in cEdge route table",
				Computed:            true,
			},
			"ecmp_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: "OMP Admin Distance IPv4",
				Computed:            true,
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: "OMP Admin Distance IPv6",
				Computed:            true,
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertisement_interval": schema.Int64Attribute{
				MarkdownDescription: "Advertisement Interval (seconds)",
				Computed:            true,
			},
			"advertisement_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"graceful_restart_timer": schema.Int64Attribute{
				MarkdownDescription: "Graceful Restart Timer (seconds)",
				Computed:            true,
			},
			"graceful_restart_timer_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"eor_timer": schema.Int64Attribute{
				MarkdownDescription: "EOR Timer",
				Computed:            true,
			},
			"eor_timer_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"holdtime": schema.Int64Attribute{
				MarkdownDescription: "Hold Time (seconds)",
				Computed:            true,
			},
			"holdtime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_bgp": schema.BoolAttribute{
				MarkdownDescription: "BGP",
				Computed:            true,
			},
			"advertise_ipv4_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_ospf": schema.BoolAttribute{
				MarkdownDescription: "OSPF",
				Computed:            true,
			},
			"advertise_ipv4_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_ospf_v3": schema.BoolAttribute{
				MarkdownDescription: "OSPFV3",
				Computed:            true,
			},
			"advertise_ipv4_ospf_v3_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_connected": schema.BoolAttribute{
				MarkdownDescription: "Connected",
				Computed:            true,
			},
			"advertise_ipv4_connected_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_static": schema.BoolAttribute{
				MarkdownDescription: "Static",
				Computed:            true,
			},
			"advertise_ipv4_static_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_eigrp": schema.BoolAttribute{
				MarkdownDescription: "EIGRP",
				Computed:            true,
			},
			"advertise_ipv4_eigrp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_lisp": schema.BoolAttribute{
				MarkdownDescription: "LISP",
				Computed:            true,
			},
			"advertise_ipv4_lisp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv4_isis": schema.BoolAttribute{
				MarkdownDescription: "ISIS",
				Computed:            true,
			},
			"advertise_ipv4_isis_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_bgp": schema.BoolAttribute{
				MarkdownDescription: "BGP",
				Computed:            true,
			},
			"advertise_ipv6_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_ospf": schema.BoolAttribute{
				MarkdownDescription: "OSPF",
				Computed:            true,
			},
			"advertise_ipv6_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_connected": schema.BoolAttribute{
				MarkdownDescription: "Connected",
				Computed:            true,
			},
			"advertise_ipv6_connected_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_static": schema.BoolAttribute{
				MarkdownDescription: "Static",
				Computed:            true,
			},
			"advertise_ipv6_static_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_eigrp": schema.BoolAttribute{
				MarkdownDescription: "EIGRP",
				Computed:            true,
			},
			"advertise_ipv6_eigrp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_lisp": schema.BoolAttribute{
				MarkdownDescription: "LISP",
				Computed:            true,
			},
			"advertise_ipv6_lisp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"advertise_ipv6_isis": schema.BoolAttribute{
				MarkdownDescription: "ISIS",
				Computed:            true,
			},
			"advertise_ipv6_isis_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ignore_region_path_length": schema.BoolAttribute{
				MarkdownDescription: "Treat hierarchical and direct (secondary region) paths equally",
				Computed:            true,
			},
			"ignore_region_path_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"transport_gateway": schema.StringAttribute{
				MarkdownDescription: "Transport Gateway Path Behavior",
				Computed:            true,
			},
			"transport_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"site_types": schema.SetAttribute{
				MarkdownDescription: "Site Types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"site_types_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *SystemOMPProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemOMPProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemOMP

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
