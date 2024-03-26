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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Mode struct {
	Id           types.String `tfsdk:"id"`
	Version      types.Int64  `tfsdk:"version"`
	TemplateType types.String `tfsdk:"template_type"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	DeviceTypes  types.Set    `tfsdk:"device_types"`
	Bay1         types.String `tfsdk:"bay_1"`
	Bay1Variable types.String `tfsdk:"bay_1_variable"`
	Bay2         types.String `tfsdk:"bay_2"`
	Bay2Variable types.String `tfsdk:"bay_2_variable"`
}

func (data Mode) getModel() string {
	return "mode"
}

func (data Mode) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "mode")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.Bay1Variable.IsNull() {
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipVariableName", data.Bay1Variable.ValueString())
	} else if data.Bay1.IsNull() {
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"subslot.bay1.mode."+"vipValue", data.Bay1.ValueString())
	}

	if !data.Bay2Variable.IsNull() {
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipVariableName", data.Bay2Variable.ValueString())
	} else if data.Bay2.IsNull() {
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"subslot.bay2.mode."+"vipValue", data.Bay2.ValueString())
	}
	return body
}

func (data *Mode) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "subslot.bay1.mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Bay1 = types.StringNull()

			v := res.Get(path + "subslot.bay1.mode.vipVariableName")
			data.Bay1Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Bay1 = types.StringNull()
			data.Bay1Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "subslot.bay1.mode.vipValue")
			data.Bay1 = types.StringValue(v.String())
			data.Bay1Variable = types.StringNull()
		}
	} else {
		data.Bay1 = types.StringNull()
		data.Bay1Variable = types.StringNull()
	}
	if value := res.Get(path + "subslot.bay2.mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Bay2 = types.StringNull()

			v := res.Get(path + "subslot.bay2.mode.vipVariableName")
			data.Bay2Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Bay2 = types.StringNull()
			data.Bay2Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "subslot.bay2.mode.vipValue")
			data.Bay2 = types.StringValue(v.String())
			data.Bay2Variable = types.StringNull()
		}
	} else {
		data.Bay2 = types.StringNull()
		data.Bay2Variable = types.StringNull()
	}
}

func (data *Mode) hasChanges(ctx context.Context, state *Mode) bool {
	hasChanges := false
	if !data.Bay1.Equal(state.Bay1) {
		hasChanges = true
	}
	if !data.Bay2.Equal(state.Bay2) {
		hasChanges = true
	}
	return hasChanges
}
