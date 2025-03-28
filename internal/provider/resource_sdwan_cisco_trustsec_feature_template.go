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
var _ resource.Resource = &CiscoTrustSecFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoTrustSecFeatureTemplateResource{}

func NewCiscoTrustSecFeatureTemplateResource() resource.Resource {
	return &CiscoTrustSecFeatureTemplateResource{}
}

type CiscoTrustSecFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoTrustSecFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_trustsec_feature_template"
}

func (r *CiscoTrustSecFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco TrustSec feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"device_sgt": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Local device security group <2..65519>").AddIntegerRangeDescription(2, 65519).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 65519),
				},
			},
			"device_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"credentials_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the TrustSec Network Access Device ID, should be same as mentioned in the Identity Services Engine (upto 32 char)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"credentials_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"credentials_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the password for the device").String,
				Optional:            true,
			},
			"credentials_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_enforcement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Role-based Access Control enforcement").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_enforcement_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_sxp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CTS SXP support").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"sxp_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SXP Source IP").String,
				Optional:            true,
			},
			"sxp_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_default_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP default password").String,
				Optional:            true,
			},
			"sxp_default_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_key_chain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP key-chain").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"sxp_key_chain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_log_binding_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables logging for IP-to-SGT binding changes").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"sxp_log_binding_changes_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_reconciliation_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the SXP reconciliation period in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 64000),
				},
			},
			"sxp_reconciliation_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_retry_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Retry period for SXP connection in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 64000),
				},
			},
			"sxp_retry_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"speaker_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Speaker hold-time in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"speaker_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"minimum_listener_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Minimum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("90").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"minimum_listener_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"maximum_listener_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Maximum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("180").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"maximum_listener_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_node_id_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define SXP Node ID type <IP, 8 char string or interface name>").AddStringEnumDescription("ip", "interface-name", "8-char-hex-string").AddDefaultValueDescription("ip").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ip", "interface-name", "8-char-hex-string"),
				},
			},
			"sxp_node_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Node ID <IP, 8 char string or interface name>").String,
				Optional:            true,
			},
			"sxp_node_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_connections": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Connections").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"peer_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Peer IP address (IPv4)").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Source IP address (IPv4)").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"preshared_key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Preshared Key type").AddStringEnumDescription("default", "key-chain", "none").AddDefaultValueDescription("none").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("default", "key-chain", "none"),
							},
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Mode of connection").AddStringEnumDescription("local", "peer").AddDefaultValueDescription("local").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("local", "peer"),
							},
						},
						"mode_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Role of a device <speaker/listener/both>").AddStringEnumDescription("listener", "speaker", "both").AddDefaultValueDescription("speaker").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("listener", "speaker", "both"),
							},
						},
						"minimum_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Minimum hold time <0..65535>").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"minimum_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"maximum_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Maximum hold time <0..65535>").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"maximum_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection VPN (VRF) ID").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65527),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
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
		},
	}
}

func (r *CiscoTrustSecFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoTrustSecFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoTrustSec

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
func (r *CiscoTrustSecFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoTrustSec

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
func (r *CiscoTrustSecFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoTrustSec

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
func (r *CiscoTrustSecFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoTrustSec

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
func (r *CiscoTrustSecFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
