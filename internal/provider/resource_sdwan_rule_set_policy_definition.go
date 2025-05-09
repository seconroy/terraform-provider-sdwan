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
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &RuleSetPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &RuleSetPolicyDefinitionResource{}

func NewRuleSetPolicyDefinitionResource() resource.Resource {
	return &RuleSetPolicyDefinitionResource{}
}

type RuleSetPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *RuleSetPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_set_policy_definition"
}

func (r *RuleSetPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Rule Set Policy Definition .").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition").String,
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of rules").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the rule").String,
							Required:            true,
						},
						"order": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The order of the rule").String,
							Required:            true,
						},
						"source_object_group_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source object group ID").String,
							Optional:            true,
						},
						"source_object_group_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source object group version").String,
							Optional:            true,
						},
						"source_data_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source data IPv4 prefix list ID").String,
							Optional:            true,
						},
						"source_data_ipv4_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source data IPv4 prefix list version").String,
							Optional:            true,
						},
						"source_ipv4_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IPv4 prefix").String,
							Optional:            true,
						},
						"source_ipv4_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IPv4 prefix variable name").String,
							Optional:            true,
						},
						"source_data_fqdn_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source data FQDN prefix list ID").String,
							Optional:            true,
						},
						"source_data_fqdn_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source data FQDN prefix list version").String,
							Optional:            true,
						},
						"source_fqdn": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source fully qualified domain name").String,
							Optional:            true,
						},
						"source_port_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source port list ID").String,
							Optional:            true,
						},
						"source_port_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source port list version").String,
							Optional:            true,
						},
						"source_port": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source port or range of ports").String,
							Optional:            true,
						},
						"source_geo_location_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source geo location list ID").String,
							Optional:            true,
						},
						"source_geo_location_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source geo location list version").String,
							Optional:            true,
						},
						"source_geo_location": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source geo location").String,
							Optional:            true,
						},
						"destination_object_group_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination object group ID").String,
							Optional:            true,
						},
						"destination_object_group_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination object group version").String,
							Optional:            true,
						},
						"destination_data_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination data IPv4 prefix list ID").String,
							Optional:            true,
						},
						"destination_data_ipv4_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination data IPv4 prefix list version").String,
							Optional:            true,
						},
						"destination_ipv4_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination IPv4 prefix").String,
							Optional:            true,
						},
						"destination_ipv4_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination IPv4 prefix variable name").String,
							Optional:            true,
						},
						"destination_data_fqdn_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination data FQDN prefix list ID").String,
							Optional:            true,
						},
						"destination_data_fqdn_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination data FQDN prefix list version").String,
							Optional:            true,
						},
						"destination_fqdn": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination fully qualified domain name").String,
							Optional:            true,
						},
						"destination_port_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination port list ID").String,
							Optional:            true,
						},
						"destination_port_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination port list version").String,
							Optional:            true,
						},
						"destination_port": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination port or range of ports").String,
							Optional:            true,
						},
						"destination_geo_location_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination geo location list ID").String,
							Optional:            true,
						},
						"destination_geo_location_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination geo location list version").String,
							Optional:            true,
						},
						"destination_geo_location": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination geo location").String,
							Optional:            true,
						},
						"protocol_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol list ID").String,
							Optional:            true,
						},
						"protocol_list_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol list version").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol name").String,
							Optional:            true,
						},
						"protocol_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol number").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
					},
				},
			},
		},
	}
}

func (r *RuleSetPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *RuleSetPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RuleSetPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *RuleSetPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RuleSetPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.ValueString()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
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
func (r *RuleSetPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RuleSetPolicyDefinition

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

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else if strings.Contains(res.Get("error.message").String(), "Template locked in edit mode") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *RuleSetPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RuleSetPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *RuleSetPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
