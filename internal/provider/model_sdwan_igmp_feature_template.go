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

type Igmp struct {
	Id               types.String    `tfsdk:"id"`
	Version          types.Int64     `tfsdk:"version"`
	TemplateType     types.String    `tfsdk:"template_type"`
	Name             types.String    `tfsdk:"name"`
	Description      types.String    `tfsdk:"description"`
	DeviceTypes      types.Set       `tfsdk:"device_types"`
	Shutdown         types.Bool      `tfsdk:"shutdown"`
	ShutdownVariable types.String    `tfsdk:"shutdown_variable"`
	Interface        []IgmpInterface `tfsdk:"interface"`
}

type IgmpInterface struct {
	Optional              types.Bool                 `tfsdk:"optional"`
	InterfaceName         types.String               `tfsdk:"interface_name"`
	InterfaceNameVariable types.String               `tfsdk:"interface_name_variable"`
	StaticJoins           []IgmpInterfaceStaticJoins `tfsdk:"static_joins"`
}

type IgmpInterfaceStaticJoins struct {
	Optional             types.Bool   `tfsdk:"optional"`
	GroupAddress         types.String `tfsdk:"group_address"`
	GroupAddressVariable types.String `tfsdk:"group_address_variable"`
}

func (data Igmp) getModel() string {
	return "igmp"
}

func (data Igmp) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "igmp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"igmp.shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}
	if len(data.Interface) > 0 {
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"igmp.interface."+"vipValue", []interface{}{})
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
		itemAttributes = append(itemAttributes, "join-group")
		if len(item.StaticJoins) > 0 {
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipPrimaryKey", []string{"group-address"})
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipPrimaryKey", []string{"group-address"})
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.StaticJoins {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "group-address")

			if !childItem.GroupAddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipVariableName", childItem.GroupAddressVariable.ValueString())
			} else if childItem.GroupAddress.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "group-address."+"vipValue", childItem.GroupAddress.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "join-group."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"igmp.interface."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *Igmp) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "igmp.shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "igmp.shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "igmp.shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "igmp.interface.vipValue"); len(value.Array()) > 0 {
		data.Interface = make([]IgmpInterface, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IgmpInterface{}
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
			if cValue := v.Get("join-group.vipValue"); len(cValue.Array()) > 0 {
				item.StaticJoins = make([]IgmpInterfaceStaticJoins, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := IgmpInterfaceStaticJoins{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("group-address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.GroupAddress = types.StringNull()

							ccv := cv.Get("group-address.vipVariableName")
							cItem.GroupAddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.GroupAddress = types.StringNull()
							cItem.GroupAddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("group-address.vipValue")
							cItem.GroupAddress = types.StringValue(ccv.String())
							cItem.GroupAddressVariable = types.StringNull()
						}
					} else {
						cItem.GroupAddress = types.StringNull()
						cItem.GroupAddressVariable = types.StringNull()
					}
					item.StaticJoins = append(item.StaticJoins, cItem)
					return true
				})
			}
			data.Interface = append(data.Interface, item)
			return true
		})
	}
}

func (data *Igmp) hasChanges(ctx context.Context, state *Igmp) bool {
	hasChanges := false
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if len(data.Interface) != len(state.Interface) {
		hasChanges = true
	} else {
		for i := range data.Interface {
			if !data.Interface[i].InterfaceName.Equal(state.Interface[i].InterfaceName) {
				hasChanges = true
			}
			if len(data.Interface[i].StaticJoins) != len(state.Interface[i].StaticJoins) {
				hasChanges = true
			} else {
				for ii := range data.Interface[i].StaticJoins {
					if !data.Interface[i].StaticJoins[ii].GroupAddress.Equal(state.Interface[i].StaticJoins[ii].GroupAddress) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}
