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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CellularController struct {
	Id                          types.String                     `tfsdk:"id"`
	Version                     types.Int64                      `tfsdk:"version"`
	TemplateType                types.String                     `tfsdk:"template_type"`
	Name                        types.String                     `tfsdk:"name"`
	Description                 types.String                     `tfsdk:"description"`
	DeviceTypes                 types.Set                        `tfsdk:"device_types"`
	CellularInterfaceId         types.String                     `tfsdk:"cellular_interface_id"`
	CellularInterfaceIdVariable types.String                     `tfsdk:"cellular_interface_id_variable"`
	DataProfiles                []CellularControllerDataProfiles `tfsdk:"data_profiles"`
	PrimarySimSlot              types.Int64                      `tfsdk:"primary_sim_slot"`
	PrimarySimSlotVariable      types.String                     `tfsdk:"primary_sim_slot_variable"`
	SimFailoverRetries          types.Int64                      `tfsdk:"sim_failover_retries"`
	SimFailoverRetriesVariable  types.String                     `tfsdk:"sim_failover_retries_variable"`
	SimFailoverTimeout          types.Int64                      `tfsdk:"sim_failover_timeout"`
	SimFailoverTimeoutVariable  types.String                     `tfsdk:"sim_failover_timeout_variable"`
	FirmwareAutoSim             types.Bool                       `tfsdk:"firmware_auto_sim"`
	FirmwareAutoSimVariable     types.String                     `tfsdk:"firmware_auto_sim_variable"`
}

type CellularControllerDataProfiles struct {
	Optional              types.Bool   `tfsdk:"optional"`
	SlotNumber            types.Int64  `tfsdk:"slot_number"`
	SlotNumberVariable    types.String `tfsdk:"slot_number_variable"`
	DataProfile           types.Int64  `tfsdk:"data_profile"`
	DataProfileVariable   types.String `tfsdk:"data_profile_variable"`
	AttachProfile         types.Int64  `tfsdk:"attach_profile"`
	AttachProfileVariable types.String `tfsdk:"attach_profile_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CellularController) getModel() string {
	return "cellular-cedge-controller"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CellularController) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cellular-cedge-controller")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.CellularInterfaceIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"id."+"vipVariableName", data.CellularInterfaceIdVariable.ValueString())
	} else if data.CellularInterfaceId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"id."+"vipValue", data.CellularInterfaceId.ValueString())
	}
	if len(data.DataProfiles) > 0 {
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipPrimaryKey", []string{"slot"})
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipPrimaryKey", []string{"slot"})
		body, _ = sjson.Set(body, path+"lte.sim.data-profile-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.DataProfiles {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "slot")

		if !item.SlotNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipVariableName", item.SlotNumberVariable.ValueString())
		} else if item.SlotNumber.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "slot."+"vipValue", item.SlotNumber.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "data-profile")

		if !item.DataProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipVariableName", item.DataProfileVariable.ValueString())
		} else if item.DataProfile.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "data-profile."+"vipValue", item.DataProfile.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "attach-profile")

		if !item.AttachProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipVariableName", item.AttachProfileVariable.ValueString())
		} else if item.AttachProfile.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "attach-profile."+"vipValue", item.AttachProfile.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"lte.sim.data-profile-list."+"vipValue.-1", itemBody)
	}

	if !data.PrimarySimSlotVariable.IsNull() {
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipVariableName", data.PrimarySimSlotVariable.ValueString())
	} else if data.PrimarySimSlot.IsNull() {
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lte.sim.primary.slot."+"vipValue", data.PrimarySimSlot.ValueInt64())
	}

	if !data.SimFailoverRetriesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipVariableName", data.SimFailoverRetriesVariable.ValueString())
	} else if data.SimFailoverRetries.IsNull() {
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lte.sim.max-retry."+"vipValue", data.SimFailoverRetries.ValueInt64())
	}

	if !data.SimFailoverTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipVariableName", data.SimFailoverTimeoutVariable.ValueString())
	} else if data.SimFailoverTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lte.failovertimer."+"vipValue", data.SimFailoverTimeout.ValueInt64())
	}

	if !data.FirmwareAutoSimVariable.IsNull() {
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipVariableName", data.FirmwareAutoSimVariable.ValueString())
	} else if data.FirmwareAutoSim.IsNull() {
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lte.firmware.auto-sim."+"vipValue", strconv.FormatBool(data.FirmwareAutoSim.ValueBool()))
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CellularController) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CellularInterfaceId = types.StringNull()

			v := res.Get(path + "id.vipVariableName")
			data.CellularInterfaceIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CellularInterfaceId = types.StringNull()
			data.CellularInterfaceIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "id.vipValue")
			data.CellularInterfaceId = types.StringValue(v.String())
			data.CellularInterfaceIdVariable = types.StringNull()
		}
	} else {
		data.CellularInterfaceId = types.StringNull()
		data.CellularInterfaceIdVariable = types.StringNull()
	}
	if value := res.Get(path + "lte.sim.data-profile-list.vipValue"); len(value.Array()) > 0 {
		data.DataProfiles = make([]CellularControllerDataProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CellularControllerDataProfiles{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("slot.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SlotNumber = types.Int64Null()

					cv := v.Get("slot.vipVariableName")
					item.SlotNumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SlotNumber = types.Int64Null()
					item.SlotNumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("slot.vipValue")
					item.SlotNumber = types.Int64Value(cv.Int())
					item.SlotNumberVariable = types.StringNull()
				}
			} else {
				item.SlotNumber = types.Int64Null()
				item.SlotNumberVariable = types.StringNull()
			}
			if cValue := v.Get("data-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DataProfile = types.Int64Null()

					cv := v.Get("data-profile.vipVariableName")
					item.DataProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DataProfile = types.Int64Null()
					item.DataProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("data-profile.vipValue")
					item.DataProfile = types.Int64Value(cv.Int())
					item.DataProfileVariable = types.StringNull()
				}
			} else {
				item.DataProfile = types.Int64Null()
				item.DataProfileVariable = types.StringNull()
			}
			if cValue := v.Get("attach-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AttachProfile = types.Int64Null()

					cv := v.Get("attach-profile.vipVariableName")
					item.AttachProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AttachProfile = types.Int64Null()
					item.AttachProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("attach-profile.vipValue")
					item.AttachProfile = types.Int64Value(cv.Int())
					item.AttachProfileVariable = types.StringNull()
				}
			} else {
				item.AttachProfile = types.Int64Null()
				item.AttachProfileVariable = types.StringNull()
			}
			data.DataProfiles = append(data.DataProfiles, item)
			return true
		})
	} else {
		if len(data.DataProfiles) > 0 {
			data.DataProfiles = []CellularControllerDataProfiles{}
		}
	}
	if value := res.Get(path + "lte.sim.primary.slot.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PrimarySimSlot = types.Int64Null()

			v := res.Get(path + "lte.sim.primary.slot.vipVariableName")
			data.PrimarySimSlotVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PrimarySimSlot = types.Int64Null()
			data.PrimarySimSlotVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "lte.sim.primary.slot.vipValue")
			data.PrimarySimSlot = types.Int64Value(v.Int())
			data.PrimarySimSlotVariable = types.StringNull()
		}
	} else {
		data.PrimarySimSlot = types.Int64Null()
		data.PrimarySimSlotVariable = types.StringNull()
	}
	if value := res.Get(path + "lte.sim.max-retry.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SimFailoverRetries = types.Int64Null()

			v := res.Get(path + "lte.sim.max-retry.vipVariableName")
			data.SimFailoverRetriesVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SimFailoverRetries = types.Int64Null()
			data.SimFailoverRetriesVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "lte.sim.max-retry.vipValue")
			data.SimFailoverRetries = types.Int64Value(v.Int())
			data.SimFailoverRetriesVariable = types.StringNull()
		}
	} else {
		data.SimFailoverRetries = types.Int64Null()
		data.SimFailoverRetriesVariable = types.StringNull()
	}
	if value := res.Get(path + "lte.failovertimer.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SimFailoverTimeout = types.Int64Null()

			v := res.Get(path + "lte.failovertimer.vipVariableName")
			data.SimFailoverTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SimFailoverTimeout = types.Int64Null()
			data.SimFailoverTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "lte.failovertimer.vipValue")
			data.SimFailoverTimeout = types.Int64Value(v.Int())
			data.SimFailoverTimeoutVariable = types.StringNull()
		}
	} else {
		data.SimFailoverTimeout = types.Int64Null()
		data.SimFailoverTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "lte.firmware.auto-sim.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.FirmwareAutoSim = types.BoolNull()

			v := res.Get(path + "lte.firmware.auto-sim.vipVariableName")
			data.FirmwareAutoSimVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.FirmwareAutoSim = types.BoolNull()
			data.FirmwareAutoSimVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "lte.firmware.auto-sim.vipValue")
			data.FirmwareAutoSim = types.BoolValue(v.Bool())
			data.FirmwareAutoSimVariable = types.StringNull()
		}
	} else {
		data.FirmwareAutoSim = types.BoolNull()
		data.FirmwareAutoSimVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CellularController) hasChanges(ctx context.Context, state *CellularController) bool {
	hasChanges := false
	if !data.CellularInterfaceId.Equal(state.CellularInterfaceId) {
		hasChanges = true
	}
	if len(data.DataProfiles) != len(state.DataProfiles) {
		hasChanges = true
	} else {
		for i := range data.DataProfiles {
			if !data.DataProfiles[i].SlotNumber.Equal(state.DataProfiles[i].SlotNumber) {
				hasChanges = true
			}
			if !data.DataProfiles[i].DataProfile.Equal(state.DataProfiles[i].DataProfile) {
				hasChanges = true
			}
			if !data.DataProfiles[i].AttachProfile.Equal(state.DataProfiles[i].AttachProfile) {
				hasChanges = true
			}
		}
	}
	if !data.PrimarySimSlot.Equal(state.PrimarySimSlot) {
		hasChanges = true
	}
	if !data.SimFailoverRetries.Equal(state.SimFailoverRetries) {
		hasChanges = true
	}
	if !data.SimFailoverTimeout.Equal(state.SimFailoverTimeout) {
		hasChanges = true
	}
	if !data.FirmwareAutoSim.Equal(state.FirmwareAutoSim) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
