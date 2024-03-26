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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Pim struct {
	Id                   types.String   `tfsdk:"id"`
	Version              types.Int64    `tfsdk:"version"`
	TemplateType         types.String   `tfsdk:"template_type"`
	Name                 types.String   `tfsdk:"name"`
	Description          types.String   `tfsdk:"description"`
	DeviceTypes          types.Set      `tfsdk:"device_types"`
	Shutdown             types.Bool     `tfsdk:"shutdown"`
	ShutdownVariable     types.String   `tfsdk:"shutdown_variable"`
	AutoRp               types.Bool     `tfsdk:"auto_rp"`
	AutoRpVariable       types.String   `tfsdk:"auto_rp_variable"`
	SptThreshold         types.Int64    `tfsdk:"spt_threshold"`
	SptThresholdVariable types.String   `tfsdk:"spt_threshold_variable"`
	Replicator           types.String   `tfsdk:"replicator"`
	ReplicatorVariable   types.String   `tfsdk:"replicator_variable"`
	Interface            []PimInterface `tfsdk:"interface"`
}

type PimInterface struct {
	Optional                  types.Bool   `tfsdk:"optional"`
	InterfaceName             types.String `tfsdk:"interface_name"`
	InterfaceNameVariable     types.String `tfsdk:"interface_name_variable"`
	HelloInterval             types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable     types.String `tfsdk:"hello_interval_variable"`
	JoinPruneInterval         types.Int64  `tfsdk:"join_prune_interval"`
	JoinPruneIntervalVariable types.String `tfsdk:"join_prune_interval_variable"`
}

func (data Pim) getModel() string {
	return "pim"
}

func (data Pim) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "pim")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.AutoRpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipVariableName", data.AutoRpVariable.ValueString())
	} else if data.AutoRp.IsNull() {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipValue", strconv.FormatBool(data.AutoRp.ValueBool()))
	}

	if !data.SptThresholdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipVariableName", data.SptThresholdVariable.ValueString())
	} else if data.SptThreshold.IsNull() {
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipValue", data.SptThreshold.ValueInt64())
	}

	if !data.ReplicatorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipVariableName", data.ReplicatorVariable.ValueString())
	} else if data.Replicator.IsNull() {
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.replicator-selection."+"vipValue", data.Replicator.ValueString())
	}
	if len(data.Interface) > 0 {
		body, _ = sjson.Set(body, path+"pim.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"pim.interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"pim.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"pim.interface."+"vipValue", []interface{}{})
	}
	for _, item := range data.Interface {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.InterfaceNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.InterfaceNameVariable.ValueString())
		} else if item.InterfaceName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.InterfaceName.ValueString())
		}
		itemAttributes = append(itemAttributes, "hello-interval")

		if !item.HelloIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipVariableName", item.HelloIntervalVariable.ValueString())
		} else if item.HelloInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipValue", item.HelloInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "join-prune-interval")

		if !item.JoinPruneIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipVariableName", item.JoinPruneIntervalVariable.ValueString())
		} else if item.JoinPruneInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipValue", item.JoinPruneInterval.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"pim.interface."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *Pim) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "pim.shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "pim.shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.auto-rp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AutoRp = types.BoolNull()

			v := res.Get(path + "pim.auto-rp.vipVariableName")
			data.AutoRpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AutoRp = types.BoolNull()
			data.AutoRpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.auto-rp.vipValue")
			data.AutoRp = types.BoolValue(v.Bool())
			data.AutoRpVariable = types.StringNull()
		}
	} else {
		data.AutoRp = types.BoolNull()
		data.AutoRpVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.spt-threshold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SptThreshold = types.Int64Null()

			v := res.Get(path + "pim.spt-threshold.vipVariableName")
			data.SptThresholdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SptThreshold = types.Int64Null()
			data.SptThresholdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.spt-threshold.vipValue")
			data.SptThreshold = types.Int64Value(v.Int())
			data.SptThresholdVariable = types.StringNull()
		}
	} else {
		data.SptThreshold = types.Int64Null()
		data.SptThresholdVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.replicator-selection.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Replicator = types.StringNull()

			v := res.Get(path + "pim.replicator-selection.vipVariableName")
			data.ReplicatorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Replicator = types.StringNull()
			data.ReplicatorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.replicator-selection.vipValue")
			data.Replicator = types.StringValue(v.String())
			data.ReplicatorVariable = types.StringNull()
		}
	} else {
		data.Replicator = types.StringNull()
		data.ReplicatorVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.interface.vipValue"); len(value.Array()) > 0 {
		data.Interface = make([]PimInterface, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PimInterface{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceName = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.InterfaceNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceName = types.StringNull()
					item.InterfaceNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.InterfaceName = types.StringValue(cv.String())
					item.InterfaceNameVariable = types.StringNull()
				}
			} else {
				item.InterfaceName = types.StringNull()
				item.InterfaceNameVariable = types.StringNull()
			}
			if cValue := v.Get("hello-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HelloInterval = types.Int64Null()

					cv := v.Get("hello-interval.vipVariableName")
					item.HelloIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HelloInterval = types.Int64Null()
					item.HelloIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("hello-interval.vipValue")
					item.HelloInterval = types.Int64Value(cv.Int())
					item.HelloIntervalVariable = types.StringNull()
				}
			} else {
				item.HelloInterval = types.Int64Null()
				item.HelloIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("join-prune-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.JoinPruneInterval = types.Int64Null()

					cv := v.Get("join-prune-interval.vipVariableName")
					item.JoinPruneIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.JoinPruneInterval = types.Int64Null()
					item.JoinPruneIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("join-prune-interval.vipValue")
					item.JoinPruneInterval = types.Int64Value(cv.Int())
					item.JoinPruneIntervalVariable = types.StringNull()
				}
			} else {
				item.JoinPruneInterval = types.Int64Null()
				item.JoinPruneIntervalVariable = types.StringNull()
			}
			data.Interface = append(data.Interface, item)
			return true
		})
	}
}

func (data *Pim) hasChanges(ctx context.Context, state *Pim) bool {
	hasChanges := false
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.AutoRp.Equal(state.AutoRp) {
		hasChanges = true
	}
	if !data.SptThreshold.Equal(state.SptThreshold) {
		hasChanges = true
	}
	if !data.Replicator.Equal(state.Replicator) {
		hasChanges = true
	}
	if len(data.Interface) != len(state.Interface) {
		hasChanges = true
	} else {
		for i := range data.Interface {
			if !data.Interface[i].InterfaceName.Equal(state.Interface[i].InterfaceName) {
				hasChanges = true
			}
			if !data.Interface[i].HelloInterval.Equal(state.Interface[i].HelloInterval) {
				hasChanges = true
			}
			if !data.Interface[i].JoinPruneInterval.Equal(state.Interface[i].JoinPruneInterval) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
