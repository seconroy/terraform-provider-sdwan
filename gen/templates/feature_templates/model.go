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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
	Version types.Int64 `tfsdk:"version"`
	TemplateType types.String `tfsdk:"template_type"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	DeviceTypes types.Set `tfsdk:"device_types"`
{{- range .Attributes}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{toGoName .TfName}} struct {
	Optional types.Bool `tfsdk:"optional"`
{{- range .Attributes}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
	Optional types.Bool `tfsdk:"optional"`
{{- range .Attributes}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
	Optional types.Bool `tfsdk:"optional"`
{{- range .Attributes}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}
// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data {{camelCase .Name}}) getModel() string {
	return "{{.Model}}"
}
// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "{{.Model}}")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})
	{{- if .FeatureTemplateIsGlobal}}
	body, _ = sjson.Set(body, "isGlobal", true)
	{{ end}}

	path := "templateDefinition."
	{{- range .Attributes}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
		{{- if .RequiresConstAndVar }}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", data.{{toGoName .TfName}}.Value{{.Type}}())
		{{end}}
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		{{- else if and (.DataPath) (not .ExcludeIgnore)}}
		if !gjson.Get(body, path+"{{path .DataPath}}").Exists() {
			body, _ = sjson.Set(body, path+"{{path .DataPath}}", map[string]interface{}{})
		}
		{{- end}}
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", data.{{toGoName .TfName}}.Value{{.Type}}())
		{{- if and (.RequiresConstAndVar) (.Variable) }}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.Value{{.Type}}())
		{{end}}
	}
	{{- else if and (eq .Type "Bool") (not .NodeOnlyContainer)}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		{{- else if and (.DataPath) (not .ExcludeIgnore)}}
		if !gjson.Get(body, path+"{{path .DataPath}}").Exists() {
			body, _ = sjson.Set(body, path+"{{path .DataPath}}", map[string]interface{}{})
		}
		{{- end}}
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(data.{{toGoName .TfName}}.ValueBool()))
	}
	{{- else if and (eq .Type "Bool") .NodeOnlyContainer}}
	if !data.{{toGoName .TfName}}.IsNull() {
		if data.{{toGoName .TfName}}.ValueBool() {
			if !gjson.Get(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Exists() {
				body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		}
	}
	{{- else if isListSet .}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		{{- else if and (.DataPath) (not .ExcludeIgnore)}}
		if !gjson.Get(body, path+"{{path .DataPath}}").Exists() {
			body, _ = sjson.Set(body, path+"{{path .DataPath}}", map[string]interface{}{})
		}
		{{- end}}
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
	}
	{{- else if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
	} else {
		{{- if not .ExcludeIgnore}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
		{{- end}}
	}
	for _, item := range data.{{toGoName .TfName}} {
		itemBody := ""
		itemAttributes := make([]string, 0)
		{{- range .Attributes}}
		itemAttributes = append(itemAttributes, "{{getParentModelName .}}")
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
		{{if .Variable}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
			{{- if .RequiresConstAndVar }}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", item.{{toGoName .TfName}}.Value{{.Type}}())
			{{end}}
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			{{- else if and (.DataPath) (not .ExcludeIgnore)}}
			if !gjson.Get(itemBody, "{{path .DataPath}}").Exists() {
				itemBody, _ = sjson.Set(itemBody, "{{path .DataPath}}", map[string]interface{}{})
			}
			{{- end}}
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", item.{{toGoName .TfName}}.Value{{.Type}}())
			{{- if and (.RequiresConstAndVar) (.Variable) }}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.Value{{.Type}}())
			{{end}}
		}
		{{- else if and (eq .Type "Bool") (not .NodeOnlyContainer)}}
		{{if .Variable}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			{{- else if and (.DataPath) (not .ExcludeIgnore)}}
			if !gjson.Get(itemBody, "{{path .DataPath}}").Exists() {
				itemBody, _ = sjson.Set(itemBody, "{{path .DataPath}}", map[string]interface{}{})
			}
			{{- end}}
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(item.{{toGoName .TfName}}.ValueBool()))
		}
		{{- else if and (eq .Type "Bool") .NodeOnlyContainer}}
		if !item.{{toGoName .TfName}}.IsNull() {
			if item.{{toGoName .TfName}}.ValueBool() {
				if !gjson.Get(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Exists() {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			}
		}
		{{- else if isListSet .}}
		{{if .Variable}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			{{- else if and (.DataPath) (not .ExcludeIgnore)}}
			if !gjson.Get(temBody, "{{path .DataPath}}").Exists() {
				itemBody, _ = sjson.Set(itemBody, "{{path .DataPath}}", map[string]interface{}{})
			}
			{{- end}}
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
			item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
		}
		{{- else if isNestedListSet .}}
		if len(item.{{toGoName .TfName}}) > 0 {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
		} else {
			{{- if not .ExcludeIgnore}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
			{{- end}}
		}
		for _, childItem := range item.{{toGoName .TfName}} {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			{{- range .Attributes}}
			itemChildAttributes = append(itemChildAttributes, "{{getParentModelName .}}")
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			{{if .Variable}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
				{{- if .RequiresConstAndVar }}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", childItem.{{toGoName .TfName}}.Value{{.Type}}())
				{{end}}
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
				{{- else if and (.DataPath) (not .ExcludeIgnore)}}
				if !gjson.Get(itemChildBody, "{{path .DataPath}}").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{path .DataPath}}", map[string]interface{}{})
				}
				{{- end}}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", childItem.{{toGoName .TfName}}.Value{{.Type}}())
				{{- if and (.RequiresConstAndVar) (.Variable) }}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.Value{{.Type}}())
				{{end}}
			}
			{{- else if and (eq .Type "Bool") (not .NodeOnlyContainer)}}
			{{if .Variable}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
				{{- else if and (.DataPath) (not .ExcludeIgnore)}}
				if !gjson.Get(itemChildBody, "{{path .DataPath}}").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{path .DataPath}}", map[string]interface{}{})
				}
				{{- end}}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(childItem.{{toGoName .TfName}}.ValueBool()))
			}
			{{- else if and (eq .Type "Bool") .NodeOnlyContainer}}
			if !childItem.{{toGoName .TfName}}.IsNull() {
				if childItem.{{toGoName .TfName}}.ValueBool() {
					if !gjson.Get(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Exists() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", map[string]interface{}{})
					}
				} else {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
				}
			}
			{{- else if isListSet .}}
			{{if .Variable}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
				{{- else if and (.DataPath) (not .ExcludeIgnore)}}
				if !gjson.Get(itemChildBody, "{{path .DataPath}}").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{path .DataPath}}", map[string]interface{}{})
				}
				{{- end}}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
			}
			{{- else if isNestedListSet .}}
			if len(childItem.{{toGoName .TfName}}) > 0 {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
			} else {
				{{- if not .ExcludeIgnore}}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
				{{- end}}
			}
			for _, childChildItem := range childItem.{{toGoName .TfName}} {
				itemChildChildBody := ""
				itemChildChildAttributes := make([]string, 0)
				{{- range .Attributes}}
				itemChildChildAttributes = append(itemChildChildAttributes, "{{getParentModelName .}}")
				{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
				{{if .Variable}}
				if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childChildItem.{{toGoName .TfName}}Variable.ValueString())
					{{- if .RequiresConstAndVar }}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
					{{end}}
				} else
				{{- end}} if childChildItem.{{toGoName .TfName}}.IsNull() {
					{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
					{{- else if and (.DataPath) (not .ExcludeIgnore)}}
					if !gjson.Get(itemChildChildBody, "{{path .DataPath}}").Exists() {
						itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{path .DataPath}}", map[string]interface{}{})
					}
					{{- end}}
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
					{{- if and (.RequiresConstAndVar) (.Variable) }}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childChildItem.{{toGoName .TfName}}Variable.Value{{.Type}}())
					{{end}}
				}
				{{- else if and (eq .Type "Bool") (not .NodeOnlyContainer)}}
				{{if .Variable}}
				if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childChildItem.{{toGoName .TfName}}Variable.ValueString())
				} else
				{{- end}} if childChildItem.{{toGoName .TfName}}.IsNull() {
					{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
					{{- else if and (.DataPath) (not .ExcludeIgnore)}}
					if !gjson.Get(itemChildChildBody, "{{path .DataPath}}").Exists() {
						itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{path .DataPath}}", map[string]interface{}{})
					}
					{{- end}}
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(childChildItem.{{toGoName .TfName}}.ValueBool()))
				}
				{{- else if and (eq .Type "Bool") .NodeOnlyContainer}}
				if !childChildItem.{{toGoName .TfName}}.IsNull() {
					if childChildItem.{{toGoName .TfName}}.ValueBool() {
						if !gjson.Get(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Exists() {
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", map[string]interface{}{})
						}
					} else {
						itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
						itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
					}
				}
				{{- else if isListSet .}}
				{{if .Variable}}
				if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childChildItem.{{toGoName .TfName}}Variable.ValueString())
				} else
				{{- end}} if childChildItem.{{toGoName .TfName}}.IsNull() {
					{{- if and (not .Mandatory) (not .ExcludeIgnore)}}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
					{{- else if and (.DataPath) (not .ExcludeIgnore)}}
					if !gjson.Get(itemChildChildBody, "{{path .DataPath}}").Exists() {
						itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{path .DataPath}}", map[string]interface{}{})
					}
					{{- end}}
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
					var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
					childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
				}
				{{- end}}
				{{- end}}
				if !childChildItem.Optional.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "vipOptional", childChildItem.Optional.ValueBool())
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "priority-order", itemChildChildAttributes)
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue.-1", itemChildChildBody)
			}
			{{- end}}
			{{- end}}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue.-1", itemChildBody)
		}
		{{- end}}
		{{- end}}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue.-1", itemBody)
	}
	{{- end}}
	{{- end}}
	return body
}
// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.StringNull()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.StringNull()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.StringValue(v.String())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.Int64Null()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.Int64Null()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.Int64Value(v.Int())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.Float64Null()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.Float64Null()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.Float64Value(v.Float())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.BoolNull()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			{{- if .NodeOnlyContainer}}
			data.{{toGoName .TfName}} = types.BoolValue(false)
			{{- else}}
			data.{{toGoName .TfName}} = types.BoolNull()
			{{- end}}
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.BoolValue(v.Bool())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	{{- if .NodeOnlyContainer}}
	} else if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.BoolValue(true)
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if isListSet .}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(v.Array())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if isNestedListSet .}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue"); len(value.Array()) > 0 {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.StringNull()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.StringNull()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.StringValue(cv.String())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.Int64Null()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.Int64Null()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.Int64Value(cv.Int())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.Float64Null()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.Float64Null()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.Float64Value(cv.Float())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.BoolNull()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					{{- if .NodeOnlyContainer}}
					item.{{toGoName .TfName}} = types.BoolValue(false)
					{{- else}}
					item.{{toGoName .TfName}} = types.BoolNull()
					{{- end}}
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.BoolValue(cv.Bool())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			{{- if .NodeOnlyContainer}}
			} else if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.BoolValue(true)
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			{{- end}}
			} else {
				item.{{toGoName .TfName}} = types.BoolNull()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if isListSet .}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cv.Array())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if isNestedListSet .}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue"); len(cValue.Array()) > 0 {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					{{- range .Attributes}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.StringNull()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.StringNull()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.StringValue(ccv.String())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.Int64Null()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.Int64Null()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.Int64Value(ccv.Int())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.Float64Null()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.Float64Null()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.Float64Value(ccv.Float())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.BoolNull()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							{{- if .NodeOnlyContainer}}
							cItem.{{toGoName .TfName}} = types.BoolValue(false)
							{{- else}}
							cItem.{{toGoName .TfName}} = types.BoolNull()
							{{- end}}
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.BoolValue(ccv.Bool())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					{{- if .NodeOnlyContainer}}
					} else if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.BoolValue(true)
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					{{- end}}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolNull()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if isListSet .}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(ccValue.Array()) > 0 {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(ccv.Array())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if isNestedListSet .}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue"); len(ccValue.Array()) > 0 {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							if cccValue := ccv.Get("vipOptional"); cccValue.Exists() {
								ccItem.Optional = types.BoolValue(cccValue.Bool())
							} else {
								ccItem.Optional = types.BoolNull()
							}
							{{- range .Attributes}}
							{{- if eq .Type "String"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.{{toGoName .TfName}} = types.StringNull()
									{{if .Variable}}
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(cccv.String())
									{{end}}
								} else if cccValue.String() == "ignore" {
									ccItem.{{toGoName .TfName}} = types.StringNull()
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
									ccItem.{{toGoName .TfName}} = types.StringValue(cccv.String())
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								}
							} else {
								ccItem.{{toGoName .TfName}} = types.StringNull()
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							}
							{{- else if eq .Type "Int64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.{{toGoName .TfName}} = types.Int64Null()
									{{if .Variable}}
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(cccv.String())
									{{end}}
								} else if cccValue.String() == "ignore" {
									ccItem.{{toGoName .TfName}} = types.Int64Null()
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
									ccItem.{{toGoName .TfName}} = types.Int64Value(cccv.Int())
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								}
							} else {
								ccItem.{{toGoName .TfName}} = types.Int64Null()
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							}
							{{- else if eq .Type "Float64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.{{toGoName .TfName}} = types.Float64Null()
									{{if .Variable}}
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(cccv.String())
									{{end}}
								} else if cccValue.String() == "ignore" {
									ccItem.{{toGoName .TfName}} = types.Float64Null()
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								} else if cccValue.String() == "constant" {
									cccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
									ccItem.{{toGoName .TfName}} = types.Float64Value(cccv.Float())
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								}
							} else {
								ccItem.{{toGoName .TfName}} = types.Float64Null()
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							}
							{{- else if eq .Type "Bool"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.{{toGoName .TfName}} = types.BoolNull()
									{{if .Variable}}
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(cccv.String())
									{{end}}
								} else if cccValue.String() == "ignore" {
									{{- if .NodeOnlyContainer}}
									ccItem.{{toGoName .TfName}} = types.BoolValue(false)
									{{- else}}
									ccItem.{{toGoName .TfName}} = types.BoolNull()
									{{- end}}
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
									ccItem.{{toGoName .TfName}} = types.BoolValue(cccv.Bool())
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								}
							{{- if .NodeOnlyContainer}}
							} else if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() {
								ccItem.{{toGoName .TfName}} = types.BoolValue(true)
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							{{- end}}
							} else {
								ccItem.{{toGoName .TfName}} = types.BoolNull()
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							}
							{{- else if isListSet .}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(cccValue.Array()) > 0 {
								if cccValue.String() == "variableName" {
									ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
									{{if .Variable}}
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(cccv.String())
									{{end}}
								} else if cccValue.String() == "ignore" {
									ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
									ccItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cccv.Array())
									{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
								}
							} else {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
								{{if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					} else {
						if len(cItem.{{toGoName .TfName}}) > 0 {
							cItem.{{toGoName .TfName}} = []{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
						}
					}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			} else {
				if len(item.{{toGoName .TfName}}) > 0 {
					item.{{toGoName .TfName}} = []{{$name}}{{$cname}}{{toGoName .TfName}}{}
				}
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	} else {
		if len(data.{{toGoName .TfName}}) > 0 {
			data.{{toGoName .TfName}} = []{{$name}}{{toGoName .TfName}}{}
		}
	}
	{{- end}}
	{{- end}}
}
// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *{{camelCase .Name}}) hasChanges(ctx context.Context, state *{{camelCase .Name}}) bool {
	hasChanges := false
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if and (not .Value) (not .TfOnly)}}
	{{- if not (isNestedListSet .)}}
	if !data.{{toGoName .TfName}}.Equal(state.{{toGoName .TfName}}) {
		hasChanges = true
	}
	{{- else}}
	if len(data.{{toGoName .TfName}}) != len(state.{{toGoName .TfName}}) {
		hasChanges = true
	} else {
		for i := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			{{- $cname := toGoName .TfName}}
			{{- if and (not .Value) (not .TfOnly)}}
			{{- if not (isNestedListSet .)}}
			if !data.{{$name}}[i].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			}
			{{- else}}
			if len(data.{{$name}}[i].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			} else {
				for ii := range data.{{$name}}[i].{{toGoName .TfName}} {
					{{- range .Attributes}}
					{{- $ccname := toGoName .TfName}}
					{{- if and (not .Value) (not .TfOnly)}}
					{{- if not (isNestedListSet .)}}
					if !data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					}
					{{- else}}
					if len(data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					} else {
						for iii := range data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
							{{- range .Attributes}}
							{{- if and (not .Value) (not .TfOnly)}}
							if !data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}) {
								hasChanges = true
							}
							{{- end}}
							{{- end}}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return hasChanges
}
// End of section. //template:end hasChanges
