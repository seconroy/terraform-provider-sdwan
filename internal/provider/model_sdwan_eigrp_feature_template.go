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
type Eigrp struct {
	Id                            types.String           `tfsdk:"id"`
	Version                       types.Int64            `tfsdk:"version"`
	TemplateType                  types.String           `tfsdk:"template_type"`
	Name                          types.String           `tfsdk:"name"`
	Description                   types.String           `tfsdk:"description"`
	DeviceTypes                   types.Set              `tfsdk:"device_types"`
	AsNumber                      types.Int64            `tfsdk:"as_number"`
	AsNumberVariable              types.String           `tfsdk:"as_number_variable"`
	AddressFamilies               []EigrpAddressFamilies `tfsdk:"address_families"`
	HelloInterval                 types.Int64            `tfsdk:"hello_interval"`
	HelloIntervalVariable         types.String           `tfsdk:"hello_interval_variable"`
	HoldTime                      types.Int64            `tfsdk:"hold_time"`
	HoldTimeVariable              types.String           `tfsdk:"hold_time_variable"`
	RoutePolicyName               types.String           `tfsdk:"route_policy_name"`
	RoutePolicyNameVariable       types.String           `tfsdk:"route_policy_name_variable"`
	Filter                        types.Bool             `tfsdk:"filter"`
	FilterVariable                types.String           `tfsdk:"filter_variable"`
	AuthenticationType            types.String           `tfsdk:"authentication_type"`
	AuthenticationTypeVariable    types.String           `tfsdk:"authentication_type_variable"`
	HmacAuthenticationKey         types.String           `tfsdk:"hmac_authentication_key"`
	HmacAuthenticationKeyVariable types.String           `tfsdk:"hmac_authentication_key_variable"`
	Keys                          []EigrpKeys            `tfsdk:"keys"`
	Interfaces                    []EigrpInterfaces      `tfsdk:"interfaces"`
}

type EigrpAddressFamilies struct {
	Optional      types.Bool                          `tfsdk:"optional"`
	Type          types.String                        `tfsdk:"type"`
	Redistributes []EigrpAddressFamiliesRedistributes `tfsdk:"redistributes"`
	Networks      []EigrpAddressFamiliesNetworks      `tfsdk:"networks"`
}

type EigrpKeys struct {
	Optional                     types.Bool   `tfsdk:"optional"`
	Md5KeyId                     types.Int64  `tfsdk:"md5_key_id"`
	Md5KeyIdVariable             types.String `tfsdk:"md5_key_id_variable"`
	Md5AuthenticationKey         types.String `tfsdk:"md5_authentication_key"`
	Md5AuthenticationKeyVariable types.String `tfsdk:"md5_authentication_key_variable"`
}

type EigrpInterfaces struct {
	Optional              types.Bool                        `tfsdk:"optional"`
	InterfaceName         types.String                      `tfsdk:"interface_name"`
	InterfaceNameVariable types.String                      `tfsdk:"interface_name_variable"`
	Shutdown              types.Bool                        `tfsdk:"shutdown"`
	ShutdownVariable      types.String                      `tfsdk:"shutdown_variable"`
	SummaryAddresses      []EigrpInterfacesSummaryAddresses `tfsdk:"summary_addresses"`
}

type EigrpAddressFamiliesRedistributes struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
}
type EigrpAddressFamiliesNetworks struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}

type EigrpInterfacesSummaryAddresses struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data Eigrp) getModel() string {
	return "eigrp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Eigrp) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "eigrp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.AsNumberVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipVariableName", data.AsNumberVariable.ValueString())
	} else if data.AsNumber.IsNull() {
		if !gjson.Get(body, path+"eigrp").Exists() {
			body, _ = sjson.Set(body, path+"eigrp", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.as-num."+"vipValue", data.AsNumber.ValueInt64())
	}
	if len(data.AddressFamilies) > 0 {
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipPrimaryKey", []string{"type"})
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipPrimaryKey", []string{"type"})
		body, _ = sjson.Set(body, path+"eigrp.address-family."+"vipValue", []interface{}{})
	}
	for _, item := range data.AddressFamilies {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "type")
		if item.Type.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipValue", item.Type.ValueString())
		}
		itemAttributes = append(itemAttributes, "topology")
		if len(item.Redistributes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "topology.base.redistribute."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Redistributes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "protocol")

			if !childItem.ProtocolVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipVariableName", childItem.ProtocolVariable.ValueString())
			} else if childItem.Protocol.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipValue", childItem.Protocol.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "route-map")

			if !childItem.RoutePolicyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipVariableName", childItem.RoutePolicyVariable.ValueString())
			} else if childItem.RoutePolicy.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-map."+"vipValue", childItem.RoutePolicy.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "topology.base.redistribute."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "network")
		if len(item.Networks) > 0 {
			itemBody, _ = sjson.Set(itemBody, "network."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "network."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "network."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "network."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Networks {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "network."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"eigrp.address-family."+"vipValue.-1", itemBody)
	}

	if !data.HelloIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipVariableName", data.HelloIntervalVariable.ValueString())
	} else if data.HelloInterval.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.hello-interval."+"vipValue", data.HelloInterval.ValueInt64())
	}

	if !data.HoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipVariableName", data.HoldTimeVariable.ValueString())
	} else if data.HoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.hold-time."+"vipValue", data.HoldTime.ValueInt64())
	}

	if !data.RoutePolicyNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipVariableName", data.RoutePolicyNameVariable.ValueString())
	} else if data.RoutePolicyName.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.table-map.name."+"vipValue", data.RoutePolicyName.ValueString())
	}

	if !data.FilterVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipVariableName", data.FilterVariable.ValueString())
	} else if data.Filter.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.table-map.filter."+"vipValue", strconv.FormatBool(data.Filter.ValueBool()))
	}

	if !data.AuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipVariableName", data.AuthenticationTypeVariable.ValueString())
	} else if data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.authentication.type."+"vipValue", data.AuthenticationType.ValueString())
	}

	if !data.HmacAuthenticationKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipVariableName", data.HmacAuthenticationKeyVariable.ValueString())
	} else if data.HmacAuthenticationKey.IsNull() {
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.authentication.auth-key."+"vipValue", data.HmacAuthenticationKey.ValueString())
	}
	if len(data.Keys) > 0 {
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipPrimaryKey", []string{"key_id"})
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipPrimaryKey", []string{"key_id"})
		body, _ = sjson.Set(body, path+"eigrp.authentication.keychain.key."+"vipValue", []interface{}{})
	}
	for _, item := range data.Keys {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "key_id")

		if !item.Md5KeyIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipVariableName", item.Md5KeyIdVariable.ValueString())
		} else if item.Md5KeyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "key_id."+"vipValue", item.Md5KeyId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "keystring")

		if !item.Md5AuthenticationKeyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipVariableName", item.Md5AuthenticationKeyVariable.ValueString())
		} else if item.Md5AuthenticationKey.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "keystring."+"vipValue", item.Md5AuthenticationKey.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"eigrp.authentication.keychain.key."+"vipValue.-1", itemBody)
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"eigrp.af-interface."+"vipValue", []interface{}{})
	}
	for _, item := range data.Interfaces {
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
		itemAttributes = append(itemAttributes, "shutdown")

		if !item.ShutdownVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipVariableName", item.ShutdownVariable.ValueString())
		} else if item.Shutdown.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipValue", strconv.FormatBool(item.Shutdown.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "summary-address")
		if len(item.SummaryAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "summary-address."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.SummaryAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "summary-address."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"eigrp.af-interface."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Eigrp) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "eigrp.as-num.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AsNumber = types.Int64Null()

			v := res.Get(path + "eigrp.as-num.vipVariableName")
			data.AsNumberVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AsNumber = types.Int64Null()
			data.AsNumberVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.as-num.vipValue")
			data.AsNumber = types.Int64Value(v.Int())
			data.AsNumberVariable = types.StringNull()
		}
	} else {
		data.AsNumber = types.Int64Null()
		data.AsNumberVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.address-family.vipValue"); len(value.Array()) > 0 {
		data.AddressFamilies = make([]EigrpAddressFamilies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := EigrpAddressFamilies{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Type = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Type = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("type.vipValue")
					item.Type = types.StringValue(cv.String())

				}
			} else {
				item.Type = types.StringNull()

			}
			if cValue := v.Get("topology.base.redistribute.vipValue"); len(cValue.Array()) > 0 {
				item.Redistributes = make([]EigrpAddressFamiliesRedistributes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EigrpAddressFamiliesRedistributes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("protocol.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Protocol = types.StringNull()

							ccv := cv.Get("protocol.vipVariableName")
							cItem.ProtocolVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Protocol = types.StringNull()
							cItem.ProtocolVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("protocol.vipValue")
							cItem.Protocol = types.StringValue(ccv.String())
							cItem.ProtocolVariable = types.StringNull()
						}
					} else {
						cItem.Protocol = types.StringNull()
						cItem.ProtocolVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-map.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RoutePolicy = types.StringNull()

							ccv := cv.Get("route-map.vipVariableName")
							cItem.RoutePolicyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RoutePolicy = types.StringNull()
							cItem.RoutePolicyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("route-map.vipValue")
							cItem.RoutePolicy = types.StringValue(ccv.String())
							cItem.RoutePolicyVariable = types.StringNull()
						}
					} else {
						cItem.RoutePolicy = types.StringNull()
						cItem.RoutePolicyVariable = types.StringNull()
					}
					item.Redistributes = append(item.Redistributes, cItem)
					return true
				})
			} else {
				if len(item.Redistributes) > 0 {
					item.Redistributes = []EigrpAddressFamiliesRedistributes{}
				}
			}
			if cValue := v.Get("network.vipValue"); len(cValue.Array()) > 0 {
				item.Networks = make([]EigrpAddressFamiliesNetworks, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EigrpAddressFamiliesNetworks{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Networks = append(item.Networks, cItem)
					return true
				})
			} else {
				if len(item.Networks) > 0 {
					item.Networks = []EigrpAddressFamiliesNetworks{}
				}
			}
			data.AddressFamilies = append(data.AddressFamilies, item)
			return true
		})
	} else {
		if len(data.AddressFamilies) > 0 {
			data.AddressFamilies = []EigrpAddressFamilies{}
		}
	}
	if value := res.Get(path + "eigrp.hello-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HelloInterval = types.Int64Null()

			v := res.Get(path + "eigrp.hello-interval.vipVariableName")
			data.HelloIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HelloInterval = types.Int64Null()
			data.HelloIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.hello-interval.vipValue")
			data.HelloInterval = types.Int64Value(v.Int())
			data.HelloIntervalVariable = types.StringNull()
		}
	} else {
		data.HelloInterval = types.Int64Null()
		data.HelloIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.hold-time.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HoldTime = types.Int64Null()

			v := res.Get(path + "eigrp.hold-time.vipVariableName")
			data.HoldTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HoldTime = types.Int64Null()
			data.HoldTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.hold-time.vipValue")
			data.HoldTime = types.Int64Value(v.Int())
			data.HoldTimeVariable = types.StringNull()
		}
	} else {
		data.HoldTime = types.Int64Null()
		data.HoldTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.table-map.name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RoutePolicyName = types.StringNull()

			v := res.Get(path + "eigrp.table-map.name.vipVariableName")
			data.RoutePolicyNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RoutePolicyName = types.StringNull()
			data.RoutePolicyNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.table-map.name.vipValue")
			data.RoutePolicyName = types.StringValue(v.String())
			data.RoutePolicyNameVariable = types.StringNull()
		}
	} else {
		data.RoutePolicyName = types.StringNull()
		data.RoutePolicyNameVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.table-map.filter.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Filter = types.BoolNull()

			v := res.Get(path + "eigrp.table-map.filter.vipVariableName")
			data.FilterVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Filter = types.BoolNull()
			data.FilterVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.table-map.filter.vipValue")
			data.Filter = types.BoolValue(v.Bool())
			data.FilterVariable = types.StringNull()
		}
	} else {
		data.Filter = types.BoolNull()
		data.FilterVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.authentication.type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AuthenticationType = types.StringNull()

			v := res.Get(path + "eigrp.authentication.type.vipVariableName")
			data.AuthenticationTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthenticationType = types.StringNull()
			data.AuthenticationTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.authentication.type.vipValue")
			data.AuthenticationType = types.StringValue(v.String())
			data.AuthenticationTypeVariable = types.StringNull()
		}
	} else {
		data.AuthenticationType = types.StringNull()
		data.AuthenticationTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.authentication.auth-key.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HmacAuthenticationKey = types.StringNull()

			v := res.Get(path + "eigrp.authentication.auth-key.vipVariableName")
			data.HmacAuthenticationKeyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HmacAuthenticationKey = types.StringNull()
			data.HmacAuthenticationKeyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "eigrp.authentication.auth-key.vipValue")
			data.HmacAuthenticationKey = types.StringValue(v.String())
			data.HmacAuthenticationKeyVariable = types.StringNull()
		}
	} else {
		data.HmacAuthenticationKey = types.StringNull()
		data.HmacAuthenticationKeyVariable = types.StringNull()
	}
	if value := res.Get(path + "eigrp.authentication.keychain.key.vipValue"); len(value.Array()) > 0 {
		data.Keys = make([]EigrpKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := EigrpKeys{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("key_id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Md5KeyId = types.Int64Null()

					cv := v.Get("key_id.vipVariableName")
					item.Md5KeyIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Md5KeyId = types.Int64Null()
					item.Md5KeyIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("key_id.vipValue")
					item.Md5KeyId = types.Int64Value(cv.Int())
					item.Md5KeyIdVariable = types.StringNull()
				}
			} else {
				item.Md5KeyId = types.Int64Null()
				item.Md5KeyIdVariable = types.StringNull()
			}
			if cValue := v.Get("keystring.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Md5AuthenticationKey = types.StringNull()

					cv := v.Get("keystring.vipVariableName")
					item.Md5AuthenticationKeyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Md5AuthenticationKey = types.StringNull()
					item.Md5AuthenticationKeyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("keystring.vipValue")
					item.Md5AuthenticationKey = types.StringValue(cv.String())
					item.Md5AuthenticationKeyVariable = types.StringNull()
				}
			} else {
				item.Md5AuthenticationKey = types.StringNull()
				item.Md5AuthenticationKeyVariable = types.StringNull()
			}
			data.Keys = append(data.Keys, item)
			return true
		})
	} else {
		if len(data.Keys) > 0 {
			data.Keys = []EigrpKeys{}
		}
	}
	if value := res.Get(path + "eigrp.af-interface.vipValue"); len(value.Array()) > 0 {
		data.Interfaces = make([]EigrpInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := EigrpInterfaces{}
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
			if cValue := v.Get("summary-address.vipValue"); len(cValue.Array()) > 0 {
				item.SummaryAddresses = make([]EigrpInterfacesSummaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EigrpInterfacesSummaryAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.SummaryAddresses = append(item.SummaryAddresses, cItem)
					return true
				})
			} else {
				if len(item.SummaryAddresses) > 0 {
					item.SummaryAddresses = []EigrpInterfacesSummaryAddresses{}
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		if len(data.Interfaces) > 0 {
			data.Interfaces = []EigrpInterfaces{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Eigrp) hasChanges(ctx context.Context, state *Eigrp) bool {
	hasChanges := false
	if !data.AsNumber.Equal(state.AsNumber) {
		hasChanges = true
	}
	if len(data.AddressFamilies) != len(state.AddressFamilies) {
		hasChanges = true
	} else {
		for i := range data.AddressFamilies {
			if !data.AddressFamilies[i].Type.Equal(state.AddressFamilies[i].Type) {
				hasChanges = true
			}
			if len(data.AddressFamilies[i].Redistributes) != len(state.AddressFamilies[i].Redistributes) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Redistributes {
					if !data.AddressFamilies[i].Redistributes[ii].Protocol.Equal(state.AddressFamilies[i].Redistributes[ii].Protocol) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].Redistributes[ii].RoutePolicy.Equal(state.AddressFamilies[i].Redistributes[ii].RoutePolicy) {
						hasChanges = true
					}
				}
			}
			if len(data.AddressFamilies[i].Networks) != len(state.AddressFamilies[i].Networks) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Networks {
					if !data.AddressFamilies[i].Networks[ii].Prefix.Equal(state.AddressFamilies[i].Networks[ii].Prefix) {
						hasChanges = true
					}
				}
			}
		}
	}
	if !data.HelloInterval.Equal(state.HelloInterval) {
		hasChanges = true
	}
	if !data.HoldTime.Equal(state.HoldTime) {
		hasChanges = true
	}
	if !data.RoutePolicyName.Equal(state.RoutePolicyName) {
		hasChanges = true
	}
	if !data.Filter.Equal(state.Filter) {
		hasChanges = true
	}
	if !data.AuthenticationType.Equal(state.AuthenticationType) {
		hasChanges = true
	}
	if !data.HmacAuthenticationKey.Equal(state.HmacAuthenticationKey) {
		hasChanges = true
	}
	if len(data.Keys) != len(state.Keys) {
		hasChanges = true
	} else {
		for i := range data.Keys {
			if !data.Keys[i].Md5KeyId.Equal(state.Keys[i].Md5KeyId) {
				hasChanges = true
			}
			if !data.Keys[i].Md5AuthenticationKey.Equal(state.Keys[i].Md5AuthenticationKey) {
				hasChanges = true
			}
		}
	}
	if len(data.Interfaces) != len(state.Interfaces) {
		hasChanges = true
	} else {
		for i := range data.Interfaces {
			if !data.Interfaces[i].InterfaceName.Equal(state.Interfaces[i].InterfaceName) {
				hasChanges = true
			}
			if !data.Interfaces[i].Shutdown.Equal(state.Interfaces[i].Shutdown) {
				hasChanges = true
			}
			if len(data.Interfaces[i].SummaryAddresses) != len(state.Interfaces[i].SummaryAddresses) {
				hasChanges = true
			} else {
				for ii := range data.Interfaces[i].SummaryAddresses {
					if !data.Interfaces[i].SummaryAddresses[ii].Prefix.Equal(state.Interfaces[i].SummaryAddresses[ii].Prefix) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
