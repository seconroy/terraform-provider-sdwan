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
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CEdgeAAAFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CEdgeAAAFeatureTemplateResource{}

func NewCEdgeAAAFeatureTemplateResource() resource.Resource {
	return &CEdgeAAAFeatureTemplateResource{}
}

type CEdgeAAAFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CEdgeAAAFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cedge_aaa_feature_template"
}

func (r *CEdgeAAAFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a cEdge AAA feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"dot1x_authentication": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication configurations parameters").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"dot1x_authentication_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"dot1x_accounting": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accounting configurations parameters").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"dot1x_accounting_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"server_groups_priority_order": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ServerGroups priority order").AddDefaultValueDescription("local").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 220),
				},
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Create local login account").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the username").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the user password").String,
							Optional:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the user scrypt password/hash").String,
							Optional:            true,
						},
						"privilege_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Privilege Level for this user").AddStringEnumDescription("1", "15").AddDefaultValueDescription("15").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "15"),
							},
						},
						"privilege_level_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ssh_pubkeys": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of RSA public-keys per user").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key_string": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the RSA key string").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 1024),
										},
									},
									"key_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Only RSA is supported").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 32),
										},
									},
									"key_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"radius_server_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the Radius serverGroup").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Radius server Group Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which Radius server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65530),
							},
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach Radius server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the Radius server").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set IP address of Radius server").String,
										Optional:            true,
									},
									"authentication_port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Authentication port to use to connect to Radius server").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("1812").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65534),
										},
									},
									"authentication_port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"accounting_port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Accounting port to use to connect to Radius server").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("1813").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65534),
										},
									},
									"accounting_port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure how long to wait for replies from the Radius server").AddIntegerRangeDescription(1, 1000).AddDefaultValueDescription("5").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 1000),
										},
									},
									"timeout_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"retransmit": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure how many times to contact this Radius server").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("3").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 100),
										},
									},
									"retransmit_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the Radius server shared key").String,
										Optional:            true,
									},
									"secret_key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the Radius server shared type 7 encrypted key").String,
										Optional:            true,
									},
									"secret_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"encryption_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of encyption. To be used for type 6").AddStringEnumDescription("6", "7").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("6", "7"),
										},
									},
									"key_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("key type").AddStringEnumDescription("key", "pac").AddDefaultValueDescription("key").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("key", "pac"),
										},
									},
									"key_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"radius_clients": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a RADIUS client").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Client IP").String,
							Optional:            true,
						},
						"client_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_configurations": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN configuration").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"vpn_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VPN ID").String,
										Optional:            true,
									},
									"vpn_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"server_key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Specify a RADIUS client server-key").String,
										Optional:            true,
									},
									"server_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"radius_dynamic_author_server_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a radius dynamic author server-key").String,
				Optional:            true,
			},
			"radius_dynamic_author_server_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"radius_dynamic_author_domain_stripping": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain Stripping").AddStringEnumDescription("yes", "no", "right-to-left").AddDefaultValueDescription("no").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("yes", "no", "right-to-left"),
				},
			},
			"radius_dynamic_author_domain_stripping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"radius_dynamic_author_authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Type").AddStringEnumDescription("any", "all", "session-key").AddDefaultValueDescription("any").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("any", "all", "session-key"),
				},
			},
			"radius_dynamic_author_authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"radius_dynamic_author_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify Radius Dynamic Author Port").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("1700").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"radius_dynamic_author_port_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"radius_trustsec_cts_authorization_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CTS Authorization List").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"radius_trustsec_cts_authorization_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"radius_trustsec_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RADIUS trustsec group").String,
				Optional:            true,
			},
			"tacacs_server_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the TACACS serverGroup").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set TACACS server Group Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which TACACS server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65530),
							},
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach TACACS server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the TACACS server").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set IP address of TACACS server").String,
										Optional:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("TACACS Port").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("49").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure how long to wait for replies from the TACACS server").AddIntegerRangeDescription(1, 1000).AddDefaultValueDescription("5").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 1000),
										},
									},
									"timeout_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the TACACS server shared key").String,
										Optional:            true,
									},
									"secret_key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the TACACS server shared type 7 encrypted key").String,
										Optional:            true,
									},
									"secret_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"encryption_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of encyption. To be used for type 6").AddStringEnumDescription("6", "7").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("6", "7"),
										},
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"accounting_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the accounting rules").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Accounting Rule ID").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"method": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Accounting Method").AddStringEnumDescription("commands", "exec", "network", "system").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("commands", "exec", "network", "system"),
							},
						},
						"privilege_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Privilege level when method is commands").AddStringEnumDescription("1", "15").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "15"),
							},
						},
						"start_stop": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Record start and stop without waiting").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"start_stop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"groups": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma separated list of groups").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"authorization_console": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For enabling console authorization").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"authorization_console_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"authorization_config_commands": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For configuration mode commands.").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"authorization_config_commands_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"authorization_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the Authorization Rules").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Authorization Rule ID").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"method": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Method").AddStringEnumDescription("commands").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("commands"),
							},
						},
						"privilege_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Privilege level when method is commands").AddStringEnumDescription("1", "15").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "15"),
							},
						},
						"groups": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma separated list of groups").String,
							Optional:            true,
						},
						"authenticated": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Succeed if user has authenticated").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *CEdgeAAAFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CEdgeAAAFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CEdgeAAA

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CEdgeAAAFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CEdgeAAA

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CEdgeAAAFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CEdgeAAA

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CEdgeAAAFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CEdgeAAA

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CEdgeAAAFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
