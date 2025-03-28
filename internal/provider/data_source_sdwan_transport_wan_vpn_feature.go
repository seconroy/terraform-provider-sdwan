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
	_ datasource.DataSource              = &TransportWANVPNProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &TransportWANVPNProfileParcelDataSource{}
)

func NewTransportWANVPNProfileParcelDataSource() datasource.DataSource {
	return &TransportWANVPNProfileParcelDataSource{}
}

type TransportWANVPNProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportWANVPNProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_wan_vpn_feature"
}

func (d *TransportWANVPNProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport WAN VPN Feature.",

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
			"vpn": schema.Int64Attribute{
				MarkdownDescription: "VPN",
				Computed:            true,
			},
			"enhance_ecmp_keying": schema.BoolAttribute{
				MarkdownDescription: "Enhance ECMP Keying",
				Computed:            true,
			},
			"enhance_ecmp_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"primary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: "Primary DNS Address (IPv4)",
				Computed:            true,
			},
			"primary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS Address (IPv4)",
				Computed:            true,
			},
			"secondary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"primary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: "Primary DNS Address (IPv6)",
				Computed:            true,
			},
			"primary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS Address (IPv6)",
				Computed:            true,
			},
			"secondary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"new_host_mappings": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Hostname",
							Computed:            true,
						},
						"host_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"list_of_ip_addresses": schema.SetAttribute{
							MarkdownDescription: "List of IP",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"list_of_ip_addresses_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Static Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: "Gateway",
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 Route Gateway Next Hop",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"administrative_distance": schema.Int64Attribute{
							MarkdownDescription: "Administrative distance",
							Computed:            true,
						},
						"administrative_distance_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Static Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: "Prefix",
							Computed:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: "Gateway",
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Route Gateway Next Hop",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: "IPv6 Route Gateway Next Hop",
							Computed:            true,
						},
						"nat": schema.StringAttribute{
							MarkdownDescription: "IPv6 Nat",
							Computed:            true,
						},
						"nat_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"services": schema.ListNestedAttribute{
				MarkdownDescription: "Service",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_type": schema.StringAttribute{
							MarkdownDescription: "Service Type",
							Computed:            true,
						},
					},
				},
			},
			"nat_64_v4_pools": schema.ListNestedAttribute{
				MarkdownDescription: "NAT64 V4 Pool",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat64_v4_pool_name": schema.StringAttribute{
							MarkdownDescription: "NAT64 v4 Pool Name",
							Computed:            true,
						},
						"nat64_v4_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat64_v4_pool_range_start": schema.StringAttribute{
							MarkdownDescription: "NAT64 Pool Range Start",
							Computed:            true,
						},
						"nat64_v4_pool_range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat64_v4_pool_range_end": schema.StringAttribute{
							MarkdownDescription: "NAT64 Pool Range End",
							Computed:            true,
						},
						"nat64_v4_pool_range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat64_v4_pool_overload": schema.BoolAttribute{
							MarkdownDescription: "NAT64 Overload",
							Computed:            true,
						},
						"nat64_v4_pool_overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *TransportWANVPNProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportWANVPNProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportWANVPN

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
