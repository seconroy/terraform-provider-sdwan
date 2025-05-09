//go:build ignore
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
	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &{{camelCase .Name}}ProfileParcelResource{}
var _ resource.ResourceWithImportState = &{{camelCase .Name}}ProfileParcelResource{}

func New{{camelCase .Name}}ProfileParcelResource() resource.Resource {
	return &{{camelCase .Name}}ProfileParcelResource{}
}

type {{camelCase .Name}}ProfileParcelResource struct {
	client *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *{{camelCase .Name}}ProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{getProfileParcelName .}}"
}

func (r *{{camelCase .Name}}ProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResDescription}}").AddMinimumVersionDescription("{{.MinimumVersion}}").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the {{camelCase .ParcelType}}",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the {{camelCase .ParcelType}}",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the {{camelCase .ParcelType}}",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the {{camelCase .ParcelType}}",
				Optional:            true,
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if isStringInt64 .}}String{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
					{{- if and (len .EnumValues) (not .IgnoreEnum) -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- end}}
				{{- if and .Mandatory (not .Optional)}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- end}}
				{{- if and (len .EnumValues) (not .IgnoreEnum)}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
				Validators: []validator.String{
					{{- if and (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- else if and (ne .StringMinLength 0) (eq .StringMaxLength 0)}}
					stringvalidator.LengthAtLeast({{.StringMinLength}}),
					{{- else if and (eq .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthAtMost({{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if and (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					{{- if and (ne .MinInt 0) (ne .MaxInt 0)}}
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
					{{- else if and (ne .MinInt 0) (eq .MaxInt 0)}}
					int64validator.AtLeast({{.MinInt}}),
					{{- else if and (eq .MinInt 0) (ne .MaxInt 0)}}
					int64validator.AtMost({{.MaxInt}}),
					{{- end}}
				},
				{{- else if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
					{{- else if and (ne .MinFloat 0.0) (eq .MaxFloat 0.0)}}
					float64validator.AtLeast({{.MinFloat}}),
					{{- else if and (eq .MinFloat 0.0) (ne .MaxFloat 0.0)}}
					float64validator.AtMost({{.MaxFloat}}),
					{{- end}}
				},
				{{- end}}
				{{- if isNestedListSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if isStringInt64 .}}String{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
								{{- if and (len .EnumValues) (not .IgnoreEnum) -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- end}}
							Optional:            true,
							{{- if and (len .EnumValues) (not .IgnoreEnum)}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
							Validators: []validator.String{
								{{- if and (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
								{{- else if and (ne .StringMinLength 0) (eq .StringMaxLength 0)}}
								stringvalidator.LengthAtLeast({{.StringMinLength}}),
								{{- else if and (eq .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthAtMost({{.StringMaxLength}}),
								{{- end}}
								{{- range .StringPatterns}}
								stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
								{{- end}}
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								{{- if and (ne .MinInt 0) (ne .MaxInt 0)}}
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
								{{- else if and (ne .MinInt 0) (eq .MaxInt 0)}}
								int64validator.AtLeast({{.MinInt}}),
								{{- else if and (eq .MinInt 0) (ne .MaxInt 0)}}
								int64validator.AtMost({{.MaxInt}}),
								{{- end}}
							},
							{{- else if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
							Validators: []validator.Float64{
								{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
								float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
								{{- else if and (ne .MinFloat 0.0) (eq .MaxFloat 0.0)}}
								float64validator.AtLeast({{.MinFloat}}),
								{{- else if and (eq .MinFloat 0.0) (ne .MaxFloat 0.0)}}
								float64validator.AtMost({{.MaxFloat}}),
								{{- end}}
							},
							{{- end}}
							{{- if isNestedListSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if isStringInt64 .}}String{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
											{{- if and (len .EnumValues) (not .IgnoreEnum) -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- end}}
										Optional:            true,
										{{- if and (len .EnumValues) (not .IgnoreEnum)}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
										Validators: []validator.String{
											{{- if and (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
											{{- else if and (ne .StringMinLength 0) (eq .StringMaxLength 0)}}
											stringvalidator.LengthAtLeast({{.StringMinLength}}),
											{{- else if and (eq .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthAtMost({{.StringMaxLength}}),
											{{- end}}
											{{- range .StringPatterns}}
											stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
											{{- end}}
										},
										{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
										Validators: []validator.Int64{
											{{- if and (ne .MinInt 0) (ne .MaxInt 0)}}
											int64validator.Between({{.MinInt}}, {{.MaxInt}}),
											{{- else if and (ne .MinInt 0) (eq .MaxInt 0)}}
											int64validator.AtLeast({{.MinInt}}),
											{{- else if and (eq .MinInt 0) (ne .MaxInt 0)}}
											int64validator.AtMost({{.MaxInt}}),
											{{- end}}
										},
										{{- else if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
										Validators: []validator.Float64{
											{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
											float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
											{{- else if and (ne .MinFloat 0.0) (eq .MaxFloat 0.0)}}
											float64validator.AtLeast({{.MinFloat}}),
											{{- else if and (eq .MinFloat 0.0) (ne .MaxFloat 0.0)}}
											float64validator.AtMost({{.MaxFloat}}),
											{{- end}}
										},
										{{- end}}
										{{- if isNestedListSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isList .}}List{{else if isSet .}}Set{{else if isStringInt64 .}}String{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
														{{- if and (len .EnumValues) (not .IgnoreEnum) -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- end}}
													Optional:            true,
													{{- if and (len .EnumValues) (not .IgnoreEnum)}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
													Validators: []validator.String{
														{{- if and (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
														{{- else if and (ne .StringMinLength 0) (eq .StringMaxLength 0)}}
														stringvalidator.LengthAtLeast({{.StringMinLength}}),
														{{- else if and (eq .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthAtMost({{.StringMaxLength}}),
														{{- end}}
														{{- range .StringPatterns}}
														stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
														{{- end}}
													},
													{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
													Validators: []validator.Int64{
														{{- if and (ne .MinInt 0) (ne .MaxInt 0)}}
														int64validator.Between({{.MinInt}}, {{.MaxInt}}),
														{{- else if and (ne .MinInt 0) (eq .MaxInt 0)}}
														int64validator.AtLeast({{.MinInt}}),
														{{- else if and (eq .MinInt 0) (ne .MaxInt 0)}}
														int64validator.AtMost({{.MaxInt}}),
														{{- end}}
													},
													{{- else if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
													Validators: []validator.Float64{
														{{- if and (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
														float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
														{{- else if and (ne .MinFloat 0.0) (eq .MaxFloat 0.0)}}
														float64validator.AtLeast({{.MinFloat}}),
														{{- else if and (eq .MinFloat 0.0) (ne .MaxFloat 0.0)}}
														float64validator.AtMost({{.MaxFloat}}),
														{{- end}}
													},
													{{- end}}
												},
												{{- if .Variable}}
												"{{.TfName}}_variable": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Variable name{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}").String,
													Optional:            true,
												},
												{{- end}}
												{{- end}}
												{{- end}}
											},
										},
										{{- end}}
									},
									{{- if .Variable}}
									"{{.TfName}}_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}").String,
										Optional:            true,
									},
									{{- end}}
									{{- end}}
									{{- end}}
								},
							},
							{{- end}}
						},
						{{- if .Variable}}
						"{{.TfName}}_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}").String,
							Optional:            true,
						},
						{{- end}}
						{{- end}}
						{{- end}}
					},
				},
				{{- end}}
			},
			{{- if .Variable}}
			"{{.TfName}}_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}").String,
				Optional:            true,
			},
			{{- end}}
			{{- end}}
			{{- end}}
		},
	}
}

func (r *{{camelCase .Name}}ProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}
// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *{{camelCase .Name}}ProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{camelCase .Name}}

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

	plan.Id = types.StringValue(res.Get({{if .IdAttribute}}"{{.IdAttribute}}"{{else}}"parcelId"{{end}}).String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *{{camelCase .Name}}ProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	{{- if not .FullUpdate}}
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	{{- else}}
	state.fromBody(ctx, res)
	{{- end}}
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *{{camelCase .Name}}ProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state {{camelCase .Name}}

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
	res, err := r.client.Put(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64()+1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *{{camelCase .Name}}ProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}
// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *{{camelCase .Name}}ProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{- if hasReference .Attributes}}
	count := {{ countReferences .Attributes}}
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "{{getProfileParcelName .}}_id"{{range .Attributes}}{{if .Reference}} + ",{{.TfName}}"{{end}}{{end}}
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	{{- $count := 0}}
	{{- range .Attributes}}{{- if .Reference}}{{$count = add $count 1}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{.TfName}}"), parts[{{$count}}])...)
	{{- end}}{{- end}}
	{{- else}}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	{{- end}}
	
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end import
