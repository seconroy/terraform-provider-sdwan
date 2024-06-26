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
type Switchport struct {
	Id                 types.String                   `tfsdk:"id"`
	Version            types.Int64                    `tfsdk:"version"`
	TemplateType       types.String                   `tfsdk:"template_type"`
	Name               types.String                   `tfsdk:"name"`
	Description        types.String                   `tfsdk:"description"`
	DeviceTypes        types.Set                      `tfsdk:"device_types"`
	Slot               types.Int64                    `tfsdk:"slot"`
	SubSlot            types.Int64                    `tfsdk:"sub_slot"`
	ModuleType         types.String                   `tfsdk:"module_type"`
	Interfaces         []SwitchportInterfaces         `tfsdk:"interfaces"`
	AgeOutTime         types.Int64                    `tfsdk:"age_out_time"`
	AgeOutTimeVariable types.String                   `tfsdk:"age_out_time_variable"`
	StaticMacAddresses []SwitchportStaticMacAddresses `tfsdk:"static_mac_addresses"`
}

type SwitchportInterfaces struct {
	Optional                                     types.Bool   `tfsdk:"optional"`
	Name                                         types.String `tfsdk:"name"`
	NameVariable                                 types.String `tfsdk:"name_variable"`
	SwitchportMode                               types.String `tfsdk:"switchport_mode"`
	Shutdown                                     types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                             types.String `tfsdk:"shutdown_variable"`
	Speed                                        types.String `tfsdk:"speed"`
	SpeedVariable                                types.String `tfsdk:"speed_variable"`
	Duplex                                       types.String `tfsdk:"duplex"`
	DuplexVariable                               types.String `tfsdk:"duplex_variable"`
	SwitchportAccessVlan                         types.Int64  `tfsdk:"switchport_access_vlan"`
	SwitchportAccessVlanVariable                 types.String `tfsdk:"switchport_access_vlan_variable"`
	SwitchportTrunkAllowedVlans                  types.String `tfsdk:"switchport_trunk_allowed_vlans"`
	SwitchportTrunkAllowedVlansVariable          types.String `tfsdk:"switchport_trunk_allowed_vlans_variable"`
	SwitchportTrunkNativeVlan                    types.Int64  `tfsdk:"switchport_trunk_native_vlan"`
	SwitchportTrunkNativeVlanVariable            types.String `tfsdk:"switchport_trunk_native_vlan_variable"`
	Dot1xEnable                                  types.Bool   `tfsdk:"dot1x_enable"`
	Dot1xEnableVariable                          types.String `tfsdk:"dot1x_enable_variable"`
	Dot1xPortControl                             types.String `tfsdk:"dot1x_port_control"`
	Dot1xPortControlVariable                     types.String `tfsdk:"dot1x_port_control_variable"`
	Dot1xAuthenticationOrder                     types.Set    `tfsdk:"dot1x_authentication_order"`
	Dot1xAuthenticationOrderVariable             types.String `tfsdk:"dot1x_authentication_order_variable"`
	VoiceVlan                                    types.Int64  `tfsdk:"voice_vlan"`
	VoiceVlanVariable                            types.String `tfsdk:"voice_vlan_variable"`
	Dot1xPaeEnable                               types.Bool   `tfsdk:"dot1x_pae_enable"`
	Dot1xPaeEnableVariable                       types.String `tfsdk:"dot1x_pae_enable_variable"`
	Dot1xMacAuthenticationBypass                 types.Bool   `tfsdk:"dot1x_mac_authentication_bypass"`
	Dot1xMacAuthenticationBypassVariable         types.String `tfsdk:"dot1x_mac_authentication_bypass_variable"`
	Dot1xHostMode                                types.String `tfsdk:"dot1x_host_mode"`
	Dot1xHostModeVariable                        types.String `tfsdk:"dot1x_host_mode_variable"`
	Dot1xEnablePeriodicReauth                    types.Bool   `tfsdk:"dot1x_enable_periodic_reauth"`
	Dot1xEnablePeriodicReauthVariable            types.String `tfsdk:"dot1x_enable_periodic_reauth_variable"`
	Dot1xPeriodicReauthInactivityTimeout         types.Int64  `tfsdk:"dot1x_periodic_reauth_inactivity_timeout"`
	Dot1xPeriodicReauthInactivityTimeoutVariable types.String `tfsdk:"dot1x_periodic_reauth_inactivity_timeout_variable"`
	Dot1xPeriodicReauthInterval                  types.Int64  `tfsdk:"dot1x_periodic_reauth_interval"`
	Dot1xPeriodicReauthIntervalVariable          types.String `tfsdk:"dot1x_periodic_reauth_interval_variable"`
	Dot1xControlDirection                        types.String `tfsdk:"dot1x_control_direction"`
	Dot1xControlDirectionVariable                types.String `tfsdk:"dot1x_control_direction_variable"`
	Dot1xRestrictedVlan                          types.Int64  `tfsdk:"dot1x_restricted_vlan"`
	Dot1xRestrictedVlanVariable                  types.String `tfsdk:"dot1x_restricted_vlan_variable"`
	Dot1xGuestVlan                               types.Int64  `tfsdk:"dot1x_guest_vlan"`
	Dot1xGuestVlanVariable                       types.String `tfsdk:"dot1x_guest_vlan_variable"`
	Dot1xCriticalVlan                            types.Int64  `tfsdk:"dot1x_critical_vlan"`
	Dot1xCriticalVlanVariable                    types.String `tfsdk:"dot1x_critical_vlan_variable"`
	Dot1xEnableCriticialVoiceVlan                types.Bool   `tfsdk:"dot1x_enable_criticial_voice_vlan"`
	Dot1xEnableCriticialVoiceVlanVariable        types.String `tfsdk:"dot1x_enable_criticial_voice_vlan_variable"`
}

type SwitchportStaticMacAddresses struct {
	Optional           types.Bool   `tfsdk:"optional"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
	IfName             types.String `tfsdk:"if_name"`
	IfNameVariable     types.String `tfsdk:"if_name_variable"`
	Vlan               types.Int64  `tfsdk:"vlan"`
	VlanVariable       types.String `tfsdk:"vlan_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data Switchport) getModel() string {
	return "switchport"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Switchport) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "switchport")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if data.Slot.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"slot."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"slot."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"slot."+"vipValue", data.Slot.ValueInt64())
	}
	if data.SubSlot.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"subslot."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"subslot."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"subslot."+"vipValue", data.SubSlot.ValueInt64())
	}
	if data.ModuleType.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"module."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"module."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"module."+"vipValue", data.ModuleType.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, path+"interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"interface."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"interface."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.Interfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "if-name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "switchport")
		if item.SwitchportMode.IsNull() {
			if !gjson.Get(itemBody, "switchport").Exists() {
				itemBody, _ = sjson.Set(itemBody, "switchport", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "switchport.mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.mode."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "switchport.mode."+"vipValue", item.SwitchportMode.ValueString())
		}
		itemAttributes = append(itemAttributes, "shutdown")

		if !item.ShutdownVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipVariableName", item.ShutdownVariable.ValueString())
		} else if item.Shutdown.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipValue", strconv.FormatBool(item.Shutdown.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "speed")

		if !item.SpeedVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipVariableName", item.SpeedVariable.ValueString())
		} else if item.Speed.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "speed."+"vipValue", item.Speed.ValueString())
		}
		itemAttributes = append(itemAttributes, "duplex")

		if !item.DuplexVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipVariableName", item.DuplexVariable.ValueString())
		} else if item.Duplex.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "duplex."+"vipValue", item.Duplex.ValueString())
		}
		itemAttributes = append(itemAttributes, "switchport")

		if !item.SwitchportAccessVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipVariableName", item.SwitchportAccessVlanVariable.ValueString())
		} else if item.SwitchportAccessVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "switchport.access.vlan.vlan."+"vipValue", item.SwitchportAccessVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "switchport")

		if !item.SwitchportTrunkAllowedVlansVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipVariableName", item.SwitchportTrunkAllowedVlansVariable.ValueString())
		} else if item.SwitchportTrunkAllowedVlans.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.allowed.vlan.vlans."+"vipValue", item.SwitchportTrunkAllowedVlans.ValueString())
		}
		itemAttributes = append(itemAttributes, "switchport")

		if !item.SwitchportTrunkNativeVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipVariableName", item.SwitchportTrunkNativeVlanVariable.ValueString())
		} else if item.SwitchportTrunkNativeVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "switchport.trunk.native.vlan."+"vipValue", item.SwitchportTrunkNativeVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xEnableVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipVariableName", item.Dot1xEnableVariable.ValueString())
		} else if item.Dot1xEnable.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.dot1x-enable."+"vipValue", strconv.FormatBool(item.Dot1xEnable.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xPortControlVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipVariableName", item.Dot1xPortControlVariable.ValueString())
		} else if item.Dot1xPortControl.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.port-control."+"vipValue", item.Dot1xPortControl.ValueString())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xAuthenticationOrderVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipVariableName", item.Dot1xAuthenticationOrderVariable.ValueString())
		} else if item.Dot1xAuthenticationOrder.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipType", "constant")
			var values []string
			item.Dot1xAuthenticationOrder.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "dot1x.auth-order."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.VoiceVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipVariableName", item.VoiceVlanVariable.ValueString())
		} else if item.VoiceVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.voice-vlan."+"vipValue", item.VoiceVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xPaeEnableVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipVariableName", item.Dot1xPaeEnableVariable.ValueString())
		} else if item.Dot1xPaeEnable.IsNull() {
			if !gjson.Get(itemBody, "dot1x").Exists() {
				itemBody, _ = sjson.Set(itemBody, "dot1x", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.pae-enable."+"vipValue", strconv.FormatBool(item.Dot1xPaeEnable.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xMacAuthenticationBypassVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipVariableName", item.Dot1xMacAuthenticationBypassVariable.ValueString())
		} else if item.Dot1xMacAuthenticationBypass.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.mac-authentication-bypass."+"vipValue", strconv.FormatBool(item.Dot1xMacAuthenticationBypass.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xHostModeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipVariableName", item.Dot1xHostModeVariable.ValueString())
		} else if item.Dot1xHostMode.IsNull() {
			if !gjson.Get(itemBody, "dot1x.authentication").Exists() {
				itemBody, _ = sjson.Set(itemBody, "dot1x.authentication", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.host-mode."+"vipValue", item.Dot1xHostMode.ValueString())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xEnablePeriodicReauthVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipVariableName", item.Dot1xEnablePeriodicReauthVariable.ValueString())
		} else if item.Dot1xEnablePeriodicReauth.IsNull() {
			if !gjson.Get(itemBody, "dot1x.authentication").Exists() {
				itemBody, _ = sjson.Set(itemBody, "dot1x.authentication", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.enable-periodic-reauth."+"vipValue", strconv.FormatBool(item.Dot1xEnablePeriodicReauth.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xPeriodicReauthInactivityTimeoutVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipVariableName", item.Dot1xPeriodicReauthInactivityTimeoutVariable.ValueString())
		} else if item.Dot1xPeriodicReauthInactivityTimeout.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.inactivity."+"vipValue", item.Dot1xPeriodicReauthInactivityTimeout.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xPeriodicReauthIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipVariableName", item.Dot1xPeriodicReauthIntervalVariable.ValueString())
		} else if item.Dot1xPeriodicReauthInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.periodic-reauthentication.reauthentication."+"vipValue", item.Dot1xPeriodicReauthInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xControlDirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipVariableName", item.Dot1xControlDirectionVariable.ValueString())
		} else if item.Dot1xControlDirection.IsNull() {
			if !gjson.Get(itemBody, "dot1x.authentication").Exists() {
				itemBody, _ = sjson.Set(itemBody, "dot1x.authentication", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.control-direction."+"vipValue", item.Dot1xControlDirection.ValueString())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xRestrictedVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipVariableName", item.Dot1xRestrictedVlanVariable.ValueString())
		} else if item.Dot1xRestrictedVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.restricted-vlan."+"vipValue", item.Dot1xRestrictedVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xGuestVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipVariableName", item.Dot1xGuestVlanVariable.ValueString())
		} else if item.Dot1xGuestVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.guest-vlan."+"vipValue", item.Dot1xGuestVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xCriticalVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipVariableName", item.Dot1xCriticalVlanVariable.ValueString())
		} else if item.Dot1xCriticalVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.critical-vlan."+"vipValue", item.Dot1xCriticalVlan.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dot1x")

		if !item.Dot1xEnableCriticialVoiceVlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipVariableName", item.Dot1xEnableCriticialVoiceVlanVariable.ValueString())
		} else if item.Dot1xEnableCriticialVoiceVlan.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dot1x.authentication.event.enable-voice."+"vipValue", strconv.FormatBool(item.Dot1xEnableCriticialVoiceVlan.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"interface."+"vipValue.-1", itemBody)
	}

	if !data.AgeOutTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"age-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"age-time."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"age-time."+"vipVariableName", data.AgeOutTimeVariable.ValueString())
	} else if data.AgeOutTime.IsNull() {
		body, _ = sjson.Set(body, path+"age-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"age-time."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"age-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"age-time."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"age-time."+"vipValue", data.AgeOutTime.ValueInt64())
	}
	if len(data.StaticMacAddresses) > 0 {
		body, _ = sjson.Set(body, path+"static-mac-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"static-mac-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"static-mac-address."+"vipPrimaryKey", []string{"macaddr", "vlan"})
		body, _ = sjson.Set(body, path+"static-mac-address."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.StaticMacAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "macaddr")

		if !item.MacAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipVariableName", item.MacAddressVariable.ValueString())
		} else if item.MacAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "macaddr."+"vipValue", item.MacAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "if-name")

		if !item.IfNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.IfNameVariable.ValueString())
		} else if item.IfName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.IfName.ValueString())
		}
		itemAttributes = append(itemAttributes, "vlan")

		if !item.VlanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipVariableName", item.VlanVariable.ValueString())
		} else if item.Vlan.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vlan."+"vipValue", item.Vlan.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"static-mac-address."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Switchport) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "slot.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Slot = types.Int64Null()

		} else if value.String() == "ignore" {
			data.Slot = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "slot.vipValue")
			data.Slot = types.Int64Value(v.Int())

		}
	} else {
		data.Slot = types.Int64Null()

	}
	if value := res.Get(path + "subslot.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SubSlot = types.Int64Null()

		} else if value.String() == "ignore" {
			data.SubSlot = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "subslot.vipValue")
			data.SubSlot = types.Int64Value(v.Int())

		}
	} else {
		data.SubSlot = types.Int64Null()

	}
	if value := res.Get(path + "module.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ModuleType = types.StringNull()

		} else if value.String() == "ignore" {
			data.ModuleType = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "module.vipValue")
			data.ModuleType = types.StringValue(v.String())

		}
	} else {
		data.ModuleType = types.StringNull()

	}
	if value := res.Get(path + "interface.vipValue"); len(value.Array()) > 0 {
		data.Interfaces = make([]SwitchportInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SwitchportInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("switchport.mode.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SwitchportMode = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SwitchportMode = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("switchport.mode.vipValue")
					item.SwitchportMode = types.StringValue(cv.String())

				}
			} else {
				item.SwitchportMode = types.StringNull()

			}
			if cValue := v.Get("shutdown.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Shutdown = types.BoolNull()

					cv := v.Get("shutdown.vipVariableName")
					item.ShutdownVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Shutdown = types.BoolNull()
					item.ShutdownVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("shutdown.vipValue")
					item.Shutdown = types.BoolValue(cv.Bool())
					item.ShutdownVariable = types.StringNull()
				}
			} else {
				item.Shutdown = types.BoolNull()
				item.ShutdownVariable = types.StringNull()
			}
			if cValue := v.Get("speed.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Speed = types.StringNull()

					cv := v.Get("speed.vipVariableName")
					item.SpeedVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Speed = types.StringNull()
					item.SpeedVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("speed.vipValue")
					item.Speed = types.StringValue(cv.String())
					item.SpeedVariable = types.StringNull()
				}
			} else {
				item.Speed = types.StringNull()
				item.SpeedVariable = types.StringNull()
			}
			if cValue := v.Get("duplex.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Duplex = types.StringNull()

					cv := v.Get("duplex.vipVariableName")
					item.DuplexVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Duplex = types.StringNull()
					item.DuplexVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("duplex.vipValue")
					item.Duplex = types.StringValue(cv.String())
					item.DuplexVariable = types.StringNull()
				}
			} else {
				item.Duplex = types.StringNull()
				item.DuplexVariable = types.StringNull()
			}
			if cValue := v.Get("switchport.access.vlan.vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SwitchportAccessVlan = types.Int64Null()

					cv := v.Get("switchport.access.vlan.vlan.vipVariableName")
					item.SwitchportAccessVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SwitchportAccessVlan = types.Int64Null()
					item.SwitchportAccessVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("switchport.access.vlan.vlan.vipValue")
					item.SwitchportAccessVlan = types.Int64Value(cv.Int())
					item.SwitchportAccessVlanVariable = types.StringNull()
				}
			} else {
				item.SwitchportAccessVlan = types.Int64Null()
				item.SwitchportAccessVlanVariable = types.StringNull()
			}
			if cValue := v.Get("switchport.trunk.allowed.vlan.vlans.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SwitchportTrunkAllowedVlans = types.StringNull()

					cv := v.Get("switchport.trunk.allowed.vlan.vlans.vipVariableName")
					item.SwitchportTrunkAllowedVlansVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SwitchportTrunkAllowedVlans = types.StringNull()
					item.SwitchportTrunkAllowedVlansVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("switchport.trunk.allowed.vlan.vlans.vipValue")
					item.SwitchportTrunkAllowedVlans = types.StringValue(cv.String())
					item.SwitchportTrunkAllowedVlansVariable = types.StringNull()
				}
			} else {
				item.SwitchportTrunkAllowedVlans = types.StringNull()
				item.SwitchportTrunkAllowedVlansVariable = types.StringNull()
			}
			if cValue := v.Get("switchport.trunk.native.vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SwitchportTrunkNativeVlan = types.Int64Null()

					cv := v.Get("switchport.trunk.native.vlan.vipVariableName")
					item.SwitchportTrunkNativeVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SwitchportTrunkNativeVlan = types.Int64Null()
					item.SwitchportTrunkNativeVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("switchport.trunk.native.vlan.vipValue")
					item.SwitchportTrunkNativeVlan = types.Int64Value(cv.Int())
					item.SwitchportTrunkNativeVlanVariable = types.StringNull()
				}
			} else {
				item.SwitchportTrunkNativeVlan = types.Int64Null()
				item.SwitchportTrunkNativeVlanVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.dot1x-enable.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xEnable = types.BoolNull()

					cv := v.Get("dot1x.dot1x-enable.vipVariableName")
					item.Dot1xEnableVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xEnable = types.BoolNull()
					item.Dot1xEnableVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.dot1x-enable.vipValue")
					item.Dot1xEnable = types.BoolValue(cv.Bool())
					item.Dot1xEnableVariable = types.StringNull()
				}
			} else {
				item.Dot1xEnable = types.BoolNull()
				item.Dot1xEnableVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.port-control.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xPortControl = types.StringNull()

					cv := v.Get("dot1x.port-control.vipVariableName")
					item.Dot1xPortControlVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xPortControl = types.StringNull()
					item.Dot1xPortControlVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.port-control.vipValue")
					item.Dot1xPortControl = types.StringValue(cv.String())
					item.Dot1xPortControlVariable = types.StringNull()
				}
			} else {
				item.Dot1xPortControl = types.StringNull()
				item.Dot1xPortControlVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.auth-order.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Dot1xAuthenticationOrder = types.SetNull(types.StringType)

					cv := v.Get("dot1x.auth-order.vipVariableName")
					item.Dot1xAuthenticationOrderVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xAuthenticationOrder = types.SetNull(types.StringType)
					item.Dot1xAuthenticationOrderVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.auth-order.vipValue")
					item.Dot1xAuthenticationOrder = helpers.GetStringSet(cv.Array())
					item.Dot1xAuthenticationOrderVariable = types.StringNull()
				}
			} else {
				item.Dot1xAuthenticationOrder = types.SetNull(types.StringType)
				item.Dot1xAuthenticationOrderVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.voice-vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VoiceVlan = types.Int64Null()

					cv := v.Get("dot1x.voice-vlan.vipVariableName")
					item.VoiceVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VoiceVlan = types.Int64Null()
					item.VoiceVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.voice-vlan.vipValue")
					item.VoiceVlan = types.Int64Value(cv.Int())
					item.VoiceVlanVariable = types.StringNull()
				}
			} else {
				item.VoiceVlan = types.Int64Null()
				item.VoiceVlanVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.pae-enable.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xPaeEnable = types.BoolNull()

					cv := v.Get("dot1x.pae-enable.vipVariableName")
					item.Dot1xPaeEnableVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xPaeEnable = types.BoolNull()
					item.Dot1xPaeEnableVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.pae-enable.vipValue")
					item.Dot1xPaeEnable = types.BoolValue(cv.Bool())
					item.Dot1xPaeEnableVariable = types.StringNull()
				}
			} else {
				item.Dot1xPaeEnable = types.BoolNull()
				item.Dot1xPaeEnableVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.mac-authentication-bypass.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xMacAuthenticationBypass = types.BoolNull()

					cv := v.Get("dot1x.mac-authentication-bypass.vipVariableName")
					item.Dot1xMacAuthenticationBypassVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xMacAuthenticationBypass = types.BoolNull()
					item.Dot1xMacAuthenticationBypassVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.mac-authentication-bypass.vipValue")
					item.Dot1xMacAuthenticationBypass = types.BoolValue(cv.Bool())
					item.Dot1xMacAuthenticationBypassVariable = types.StringNull()
				}
			} else {
				item.Dot1xMacAuthenticationBypass = types.BoolNull()
				item.Dot1xMacAuthenticationBypassVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.host-mode.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xHostMode = types.StringNull()

					cv := v.Get("dot1x.authentication.host-mode.vipVariableName")
					item.Dot1xHostModeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xHostMode = types.StringNull()
					item.Dot1xHostModeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.host-mode.vipValue")
					item.Dot1xHostMode = types.StringValue(cv.String())
					item.Dot1xHostModeVariable = types.StringNull()
				}
			} else {
				item.Dot1xHostMode = types.StringNull()
				item.Dot1xHostModeVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.enable-periodic-reauth.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xEnablePeriodicReauth = types.BoolNull()

					cv := v.Get("dot1x.authentication.enable-periodic-reauth.vipVariableName")
					item.Dot1xEnablePeriodicReauthVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xEnablePeriodicReauth = types.BoolNull()
					item.Dot1xEnablePeriodicReauthVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.enable-periodic-reauth.vipValue")
					item.Dot1xEnablePeriodicReauth = types.BoolValue(cv.Bool())
					item.Dot1xEnablePeriodicReauthVariable = types.StringNull()
				}
			} else {
				item.Dot1xEnablePeriodicReauth = types.BoolNull()
				item.Dot1xEnablePeriodicReauthVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.periodic-reauthentication.inactivity.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xPeriodicReauthInactivityTimeout = types.Int64Null()

					cv := v.Get("dot1x.authentication.periodic-reauthentication.inactivity.vipVariableName")
					item.Dot1xPeriodicReauthInactivityTimeoutVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xPeriodicReauthInactivityTimeout = types.Int64Null()
					item.Dot1xPeriodicReauthInactivityTimeoutVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.periodic-reauthentication.inactivity.vipValue")
					item.Dot1xPeriodicReauthInactivityTimeout = types.Int64Value(cv.Int())
					item.Dot1xPeriodicReauthInactivityTimeoutVariable = types.StringNull()
				}
			} else {
				item.Dot1xPeriodicReauthInactivityTimeout = types.Int64Null()
				item.Dot1xPeriodicReauthInactivityTimeoutVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.periodic-reauthentication.reauthentication.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xPeriodicReauthInterval = types.Int64Null()

					cv := v.Get("dot1x.authentication.periodic-reauthentication.reauthentication.vipVariableName")
					item.Dot1xPeriodicReauthIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xPeriodicReauthInterval = types.Int64Null()
					item.Dot1xPeriodicReauthIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.periodic-reauthentication.reauthentication.vipValue")
					item.Dot1xPeriodicReauthInterval = types.Int64Value(cv.Int())
					item.Dot1xPeriodicReauthIntervalVariable = types.StringNull()
				}
			} else {
				item.Dot1xPeriodicReauthInterval = types.Int64Null()
				item.Dot1xPeriodicReauthIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.control-direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xControlDirection = types.StringNull()

					cv := v.Get("dot1x.authentication.control-direction.vipVariableName")
					item.Dot1xControlDirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xControlDirection = types.StringNull()
					item.Dot1xControlDirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.control-direction.vipValue")
					item.Dot1xControlDirection = types.StringValue(cv.String())
					item.Dot1xControlDirectionVariable = types.StringNull()
				}
			} else {
				item.Dot1xControlDirection = types.StringNull()
				item.Dot1xControlDirectionVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.event.restricted-vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xRestrictedVlan = types.Int64Null()

					cv := v.Get("dot1x.authentication.event.restricted-vlan.vipVariableName")
					item.Dot1xRestrictedVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xRestrictedVlan = types.Int64Null()
					item.Dot1xRestrictedVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.event.restricted-vlan.vipValue")
					item.Dot1xRestrictedVlan = types.Int64Value(cv.Int())
					item.Dot1xRestrictedVlanVariable = types.StringNull()
				}
			} else {
				item.Dot1xRestrictedVlan = types.Int64Null()
				item.Dot1xRestrictedVlanVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.event.guest-vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xGuestVlan = types.Int64Null()

					cv := v.Get("dot1x.authentication.event.guest-vlan.vipVariableName")
					item.Dot1xGuestVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xGuestVlan = types.Int64Null()
					item.Dot1xGuestVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.event.guest-vlan.vipValue")
					item.Dot1xGuestVlan = types.Int64Value(cv.Int())
					item.Dot1xGuestVlanVariable = types.StringNull()
				}
			} else {
				item.Dot1xGuestVlan = types.Int64Null()
				item.Dot1xGuestVlanVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.event.critical-vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xCriticalVlan = types.Int64Null()

					cv := v.Get("dot1x.authentication.event.critical-vlan.vipVariableName")
					item.Dot1xCriticalVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xCriticalVlan = types.Int64Null()
					item.Dot1xCriticalVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.event.critical-vlan.vipValue")
					item.Dot1xCriticalVlan = types.Int64Value(cv.Int())
					item.Dot1xCriticalVlanVariable = types.StringNull()
				}
			} else {
				item.Dot1xCriticalVlan = types.Int64Null()
				item.Dot1xCriticalVlanVariable = types.StringNull()
			}
			if cValue := v.Get("dot1x.authentication.event.enable-voice.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dot1xEnableCriticialVoiceVlan = types.BoolNull()

					cv := v.Get("dot1x.authentication.event.enable-voice.vipVariableName")
					item.Dot1xEnableCriticialVoiceVlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dot1xEnableCriticialVoiceVlan = types.BoolNull()
					item.Dot1xEnableCriticialVoiceVlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dot1x.authentication.event.enable-voice.vipValue")
					item.Dot1xEnableCriticialVoiceVlan = types.BoolValue(cv.Bool())
					item.Dot1xEnableCriticialVoiceVlanVariable = types.StringNull()
				}
			} else {
				item.Dot1xEnableCriticialVoiceVlan = types.BoolNull()
				item.Dot1xEnableCriticialVoiceVlanVariable = types.StringNull()
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		if len(data.Interfaces) > 0 {
			data.Interfaces = []SwitchportInterfaces{}
		}
	}
	if value := res.Get(path + "age-time.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AgeOutTime = types.Int64Null()

			v := res.Get(path + "age-time.vipVariableName")
			data.AgeOutTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AgeOutTime = types.Int64Null()
			data.AgeOutTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "age-time.vipValue")
			data.AgeOutTime = types.Int64Value(v.Int())
			data.AgeOutTimeVariable = types.StringNull()
		}
	} else {
		data.AgeOutTime = types.Int64Null()
		data.AgeOutTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "static-mac-address.vipValue"); len(value.Array()) > 0 {
		data.StaticMacAddresses = make([]SwitchportStaticMacAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SwitchportStaticMacAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("macaddr.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MacAddress = types.StringNull()

					cv := v.Get("macaddr.vipVariableName")
					item.MacAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MacAddress = types.StringNull()
					item.MacAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("macaddr.vipValue")
					item.MacAddress = types.StringValue(cv.String())
					item.MacAddressVariable = types.StringNull()
				}
			} else {
				item.MacAddress = types.StringNull()
				item.MacAddressVariable = types.StringNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IfName = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.IfNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IfName = types.StringNull()
					item.IfNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.IfName = types.StringValue(cv.String())
					item.IfNameVariable = types.StringNull()
				}
			} else {
				item.IfName = types.StringNull()
				item.IfNameVariable = types.StringNull()
			}
			if cValue := v.Get("vlan.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Vlan = types.Int64Null()

					cv := v.Get("vlan.vipVariableName")
					item.VlanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Vlan = types.Int64Null()
					item.VlanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vlan.vipValue")
					item.Vlan = types.Int64Value(cv.Int())
					item.VlanVariable = types.StringNull()
				}
			} else {
				item.Vlan = types.Int64Null()
				item.VlanVariable = types.StringNull()
			}
			data.StaticMacAddresses = append(data.StaticMacAddresses, item)
			return true
		})
	} else {
		if len(data.StaticMacAddresses) > 0 {
			data.StaticMacAddresses = []SwitchportStaticMacAddresses{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Switchport) hasChanges(ctx context.Context, state *Switchport) bool {
	hasChanges := false
	if !data.Slot.Equal(state.Slot) {
		hasChanges = true
	}
	if !data.SubSlot.Equal(state.SubSlot) {
		hasChanges = true
	}
	if !data.ModuleType.Equal(state.ModuleType) {
		hasChanges = true
	}
	if len(data.Interfaces) != len(state.Interfaces) {
		hasChanges = true
	} else {
		for i := range data.Interfaces {
			if !data.Interfaces[i].Name.Equal(state.Interfaces[i].Name) {
				hasChanges = true
			}
			if !data.Interfaces[i].SwitchportMode.Equal(state.Interfaces[i].SwitchportMode) {
				hasChanges = true
			}
			if !data.Interfaces[i].Shutdown.Equal(state.Interfaces[i].Shutdown) {
				hasChanges = true
			}
			if !data.Interfaces[i].Speed.Equal(state.Interfaces[i].Speed) {
				hasChanges = true
			}
			if !data.Interfaces[i].Duplex.Equal(state.Interfaces[i].Duplex) {
				hasChanges = true
			}
			if !data.Interfaces[i].SwitchportAccessVlan.Equal(state.Interfaces[i].SwitchportAccessVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].SwitchportTrunkAllowedVlans.Equal(state.Interfaces[i].SwitchportTrunkAllowedVlans) {
				hasChanges = true
			}
			if !data.Interfaces[i].SwitchportTrunkNativeVlan.Equal(state.Interfaces[i].SwitchportTrunkNativeVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xEnable.Equal(state.Interfaces[i].Dot1xEnable) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xPortControl.Equal(state.Interfaces[i].Dot1xPortControl) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xAuthenticationOrder.Equal(state.Interfaces[i].Dot1xAuthenticationOrder) {
				hasChanges = true
			}
			if !data.Interfaces[i].VoiceVlan.Equal(state.Interfaces[i].VoiceVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xPaeEnable.Equal(state.Interfaces[i].Dot1xPaeEnable) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xMacAuthenticationBypass.Equal(state.Interfaces[i].Dot1xMacAuthenticationBypass) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xHostMode.Equal(state.Interfaces[i].Dot1xHostMode) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xEnablePeriodicReauth.Equal(state.Interfaces[i].Dot1xEnablePeriodicReauth) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xPeriodicReauthInactivityTimeout.Equal(state.Interfaces[i].Dot1xPeriodicReauthInactivityTimeout) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xPeriodicReauthInterval.Equal(state.Interfaces[i].Dot1xPeriodicReauthInterval) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xControlDirection.Equal(state.Interfaces[i].Dot1xControlDirection) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xRestrictedVlan.Equal(state.Interfaces[i].Dot1xRestrictedVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xGuestVlan.Equal(state.Interfaces[i].Dot1xGuestVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xCriticalVlan.Equal(state.Interfaces[i].Dot1xCriticalVlan) {
				hasChanges = true
			}
			if !data.Interfaces[i].Dot1xEnableCriticialVoiceVlan.Equal(state.Interfaces[i].Dot1xEnableCriticialVoiceVlan) {
				hasChanges = true
			}
		}
	}
	if !data.AgeOutTime.Equal(state.AgeOutTime) {
		hasChanges = true
	}
	if len(data.StaticMacAddresses) != len(state.StaticMacAddresses) {
		hasChanges = true
	} else {
		for i := range data.StaticMacAddresses {
			if !data.StaticMacAddresses[i].MacAddress.Equal(state.StaticMacAddresses[i].MacAddress) {
				hasChanges = true
			}
			if !data.StaticMacAddresses[i].IfName.Equal(state.StaticMacAddresses[i].IfName) {
				hasChanges = true
			}
			if !data.StaticMacAddresses[i].Vlan.Equal(state.StaticMacAddresses[i].Vlan) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
