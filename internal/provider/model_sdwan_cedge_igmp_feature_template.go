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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CEdgeIGMP struct {
	Id           types.String          `tfsdk:"id"`
	Version      types.Int64           `tfsdk:"version"`
	TemplateType types.String          `tfsdk:"template_type"`
	Name         types.String          `tfsdk:"name"`
	Description  types.String          `tfsdk:"description"`
	DeviceTypes  types.Set             `tfsdk:"device_types"`
	Interfaces   []CEdgeIGMPInterfaces `tfsdk:"interfaces"`
}

type CEdgeIGMPInterfaces struct {
	Optional     types.Bool                      `tfsdk:"optional"`
	Name         types.String                    `tfsdk:"name"`
	NameVariable types.String                    `tfsdk:"name_variable"`
	JoinGroups   []CEdgeIGMPInterfacesJoinGroups `tfsdk:"join_groups"`
}

type CEdgeIGMPInterfacesJoinGroups struct {
	Optional             types.Bool   `tfsdk:"optional"`
	GroupAddress         types.String `tfsdk:"group_address"`
	GroupAddressVariable types.String `tfsdk:"group_address_variable"`
	Source               types.String `tfsdk:"source"`
	SourceVariable       types.String `tfsdk:"source_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CEdgeIGMP) getModel() string {
	return "cedge_igmp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CEdgeIGMP) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cedge_igmp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if len(data.Interfaces) > 0 {
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
	for _, item := range data.Interfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "join-group")
		if len(item.JoinGroups) > 0 {
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipPrimaryKey", []string{"group-address", "source"})
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipPrimaryKey", []string{"group-address", "source"})
			itemBody, _ = sjson.Set(itemBody, "join-group."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.JoinGroups {
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
			itemChildAttributes = append(itemChildAttributes, "source")

			if !childItem.SourceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipVariableName", childItem.SourceVariable.ValueString())
			} else if childItem.Source.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "source."+"vipValue", childItem.Source.ValueString())
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

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CEdgeIGMP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "igmp.interface.vipValue"); len(value.Array()) > 0 {
		data.Interfaces = make([]CEdgeIGMPInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeIGMPInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("join-group.vipValue"); len(cValue.Array()) > 0 {
				item.JoinGroups = make([]CEdgeIGMPInterfacesJoinGroups, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CEdgeIGMPInterfacesJoinGroups{}
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
					if ccValue := cv.Get("source.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Source = types.StringNull()

							ccv := cv.Get("source.vipVariableName")
							cItem.SourceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Source = types.StringNull()
							cItem.SourceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("source.vipValue")
							cItem.Source = types.StringValue(ccv.String())
							cItem.SourceVariable = types.StringNull()
						}
					} else {
						cItem.Source = types.StringNull()
						cItem.SourceVariable = types.StringNull()
					}
					item.JoinGroups = append(item.JoinGroups, cItem)
					return true
				})
			} else {
				if len(item.JoinGroups) > 0 {
					item.JoinGroups = []CEdgeIGMPInterfacesJoinGroups{}
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		if len(data.Interfaces) > 0 {
			data.Interfaces = []CEdgeIGMPInterfaces{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CEdgeIGMP) hasChanges(ctx context.Context, state *CEdgeIGMP) bool {
	hasChanges := false
	if len(data.Interfaces) != len(state.Interfaces) {
		hasChanges = true
	} else {
		for i := range data.Interfaces {
			if !data.Interfaces[i].Name.Equal(state.Interfaces[i].Name) {
				hasChanges = true
			}
			if len(data.Interfaces[i].JoinGroups) != len(state.Interfaces[i].JoinGroups) {
				hasChanges = true
			} else {
				for ii := range data.Interfaces[i].JoinGroups {
					if !data.Interfaces[i].JoinGroups[ii].GroupAddress.Equal(state.Interfaces[i].JoinGroups[ii].GroupAddress) {
						hasChanges = true
					}
					if !data.Interfaces[i].JoinGroups[ii].Source.Equal(state.Interfaces[i].JoinGroups[ii].Source) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
