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
type CiscoVPNInterfaceGRE struct {
	Id                            types.String                      `tfsdk:"id"`
	Version                       types.Int64                       `tfsdk:"version"`
	TemplateType                  types.String                      `tfsdk:"template_type"`
	Name                          types.String                      `tfsdk:"name"`
	Description                   types.String                      `tfsdk:"description"`
	DeviceTypes                   types.Set                         `tfsdk:"device_types"`
	InterfaceName                 types.String                      `tfsdk:"interface_name"`
	InterfaceNameVariable         types.String                      `tfsdk:"interface_name_variable"`
	InterfaceDescription          types.String                      `tfsdk:"interface_description"`
	InterfaceDescriptionVariable  types.String                      `tfsdk:"interface_description_variable"`
	IpAddress                     types.String                      `tfsdk:"ip_address"`
	IpAddressVariable             types.String                      `tfsdk:"ip_address_variable"`
	TunnelSource                  types.String                      `tfsdk:"tunnel_source"`
	TunnelSourceVariable          types.String                      `tfsdk:"tunnel_source_variable"`
	Shutdown                      types.Bool                        `tfsdk:"shutdown"`
	ShutdownVariable              types.String                      `tfsdk:"shutdown_variable"`
	TunnelSourceInterface         types.String                      `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable types.String                      `tfsdk:"tunnel_source_interface_variable"`
	TunnelDestination             types.String                      `tfsdk:"tunnel_destination"`
	TunnelDestinationVariable     types.String                      `tfsdk:"tunnel_destination_variable"`
	Application                   types.String                      `tfsdk:"application"`
	ApplicationVariable           types.String                      `tfsdk:"application_variable"`
	IpMtu                         types.Int64                       `tfsdk:"ip_mtu"`
	IpMtuVariable                 types.String                      `tfsdk:"ip_mtu_variable"`
	TcpMssAdjust                  types.Int64                       `tfsdk:"tcp_mss_adjust"`
	TcpMssAdjustVariable          types.String                      `tfsdk:"tcp_mss_adjust_variable"`
	ClearDontFragment             types.Bool                        `tfsdk:"clear_dont_fragment"`
	ClearDontFragmentVariable     types.String                      `tfsdk:"clear_dont_fragment_variable"`
	RewriteRule                   types.String                      `tfsdk:"rewrite_rule"`
	RewriteRuleVariable           types.String                      `tfsdk:"rewrite_rule_variable"`
	AccessLists                   []CiscoVPNInterfaceGREAccessLists `tfsdk:"access_lists"`
	Tracker                       types.Set                         `tfsdk:"tracker"`
	TrackerVariable               types.String                      `tfsdk:"tracker_variable"`
	TunnelRouteVia                types.String                      `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable        types.String                      `tfsdk:"tunnel_route_via_variable"`
}

type CiscoVPNInterfaceGREAccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoVPNInterfaceGRE) getModel() string {
	return "cisco_vpn_interface_gre"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoVPNInterfaceGRE) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_vpn_interface_gre")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.InterfaceNameVariable.ValueString())
	} else if data.InterfaceName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.InterfaceName.ValueString())
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.InterfaceDescription.ValueString())
	}

	if !data.IpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.address."+"vipVariableName", data.IpAddressVariable.ValueString())
	} else if data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.address."+"vipValue", data.IpAddress.ValueString())
	}

	if !data.TunnelSourceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipVariableName", data.TunnelSourceVariable.ValueString())
	} else if data.TunnelSource.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipValue", data.TunnelSource.ValueString())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.TunnelSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipVariableName", data.TunnelSourceInterfaceVariable.ValueString())
	} else if data.TunnelSourceInterface.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipValue", data.TunnelSourceInterface.ValueString())
	}

	if !data.TunnelDestinationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipVariableName", data.TunnelDestinationVariable.ValueString())
	} else if data.TunnelDestination.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipValue", data.TunnelDestination.ValueString())
	}

	if !data.ApplicationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"application."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"application."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"application."+"vipVariableName", data.ApplicationVariable.ValueString())
	} else if data.Application.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"application."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"application."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"application."+"vipValue", data.Application.ValueString())
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mtu."+"vipVariableName", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mtu."+"vipValue", data.IpMtu.ValueInt64())
	}

	if !data.TcpMssAdjustVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipVariableName", data.TcpMssAdjustVariable.ValueString())
	} else if data.TcpMssAdjust.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipValue", data.TcpMssAdjust.ValueInt64())
	}

	if !data.ClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipVariableName", data.ClearDontFragmentVariable.ValueString())
	} else if data.ClearDontFragment.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipValue", strconv.FormatBool(data.ClearDontFragment.ValueBool()))
	}

	if !data.RewriteRuleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipVariableName", data.RewriteRuleVariable.ValueString())
	} else if data.RewriteRule.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipValue", data.RewriteRule.ValueString())
	}
	if len(data.AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"access-list."+"vipValue.-1", itemBody)
	}

	if !data.TrackerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tracker."+"vipVariableName", data.TrackerVariable.ValueString())
	} else if data.Tracker.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "constant")
		var values []string
		data.Tracker.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", values)
	}

	if !data.TunnelRouteViaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipVariableName", data.TunnelRouteViaVariable.ValueString())
	} else if data.TunnelRouteVia.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipValue", data.TunnelRouteVia.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoVPNInterfaceGRE) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.InterfaceNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceName = types.StringNull()
			data.InterfaceNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.InterfaceName = types.StringValue(v.String())
			data.InterfaceNameVariable = types.StringNull()
		}
	} else {
		data.InterfaceName = types.StringNull()
		data.InterfaceNameVariable = types.StringNull()
	}
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.InterfaceDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceDescription = types.StringNull()
			data.InterfaceDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.InterfaceDescription = types.StringValue(v.String())
			data.InterfaceDescriptionVariable = types.StringNull()
		}
	} else {
		data.InterfaceDescription = types.StringNull()
		data.InterfaceDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpAddress = types.StringNull()

			v := res.Get(path + "ip.address.vipVariableName")
			data.IpAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpAddress = types.StringNull()
			data.IpAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address.vipValue")
			data.IpAddress = types.StringValue(v.String())
			data.IpAddressVariable = types.StringNull()
		}
	} else {
		data.IpAddress = types.StringNull()
		data.IpAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-source.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelSource = types.StringNull()

			v := res.Get(path + "tunnel-source.vipVariableName")
			data.TunnelSourceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelSource = types.StringNull()
			data.TunnelSourceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-source.vipValue")
			data.TunnelSource = types.StringValue(v.String())
			data.TunnelSourceVariable = types.StringNull()
		}
	} else {
		data.TunnelSource = types.StringNull()
		data.TunnelSourceVariable = types.StringNull()
	}
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-source-interface.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelSourceInterface = types.StringNull()

			v := res.Get(path + "tunnel-source-interface.vipVariableName")
			data.TunnelSourceInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelSourceInterface = types.StringNull()
			data.TunnelSourceInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-source-interface.vipValue")
			data.TunnelSourceInterface = types.StringValue(v.String())
			data.TunnelSourceInterfaceVariable = types.StringNull()
		}
	} else {
		data.TunnelSourceInterface = types.StringNull()
		data.TunnelSourceInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-destination.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelDestination = types.StringNull()

			v := res.Get(path + "tunnel-destination.vipVariableName")
			data.TunnelDestinationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelDestination = types.StringNull()
			data.TunnelDestinationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-destination.vipValue")
			data.TunnelDestination = types.StringValue(v.String())
			data.TunnelDestinationVariable = types.StringNull()
		}
	} else {
		data.TunnelDestination = types.StringNull()
		data.TunnelDestinationVariable = types.StringNull()
	}
	if value := res.Get(path + "application.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Application = types.StringNull()

			v := res.Get(path + "application.vipVariableName")
			data.ApplicationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Application = types.StringNull()
			data.ApplicationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "application.vipValue")
			data.Application = types.StringValue(v.String())
			data.ApplicationVariable = types.StringNull()
		}
	} else {
		data.Application = types.StringNull()
		data.ApplicationVariable = types.StringNull()
	}
	if value := res.Get(path + "mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpMtu = types.Int64Null()

			v := res.Get(path + "mtu.vipVariableName")
			data.IpMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpMtu = types.Int64Null()
			data.IpMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mtu.vipValue")
			data.IpMtu = types.Int64Value(v.Int())
			data.IpMtuVariable = types.StringNull()
		}
	} else {
		data.IpMtu = types.Int64Null()
		data.IpMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpMssAdjust = types.Int64Null()

			v := res.Get(path + "tcp-mss-adjust.vipVariableName")
			data.TcpMssAdjustVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpMssAdjust = types.Int64Null()
			data.TcpMssAdjustVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tcp-mss-adjust.vipValue")
			data.TcpMssAdjust = types.Int64Value(v.Int())
			data.TcpMssAdjustVariable = types.StringNull()
		}
	} else {
		data.TcpMssAdjust = types.Int64Null()
		data.TcpMssAdjustVariable = types.StringNull()
	}
	if value := res.Get(path + "clear-dont-fragment.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ClearDontFragment = types.BoolNull()

			v := res.Get(path + "clear-dont-fragment.vipVariableName")
			data.ClearDontFragmentVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ClearDontFragment = types.BoolNull()
			data.ClearDontFragmentVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "clear-dont-fragment.vipValue")
			data.ClearDontFragment = types.BoolValue(v.Bool())
			data.ClearDontFragmentVariable = types.StringNull()
		}
	} else {
		data.ClearDontFragment = types.BoolNull()
		data.ClearDontFragmentVariable = types.StringNull()
	}
	if value := res.Get(path + "rewrite-rule.rule-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RewriteRule = types.StringNull()

			v := res.Get(path + "rewrite-rule.rule-name.vipVariableName")
			data.RewriteRuleVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RewriteRule = types.StringNull()
			data.RewriteRuleVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "rewrite-rule.rule-name.vipValue")
			data.RewriteRule = types.StringValue(v.String())
			data.RewriteRuleVariable = types.StringNull()
		}
	} else {
		data.RewriteRule = types.StringNull()
		data.RewriteRuleVariable = types.StringNull()
	}
	if value := res.Get(path + "access-list.vipValue"); len(value.Array()) > 0 {
		data.AccessLists = make([]CiscoVPNInterfaceGREAccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceGREAccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.AccessLists = append(data.AccessLists, item)
			return true
		})
	} else {
		if len(data.AccessLists) > 0 {
			data.AccessLists = []CiscoVPNInterfaceGREAccessLists{}
		}
	}
	if value := res.Get(path + "tracker.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.Tracker = types.SetNull(types.StringType)

			v := res.Get(path + "tracker.vipVariableName")
			data.TrackerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Tracker = types.SetNull(types.StringType)
			data.TrackerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tracker.vipValue")
			data.Tracker = helpers.GetStringSet(v.Array())
			data.TrackerVariable = types.StringNull()
		}
	} else {
		data.Tracker = types.SetNull(types.StringType)
		data.TrackerVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-route-via.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelRouteVia = types.StringNull()

			v := res.Get(path + "tunnel-route-via.vipVariableName")
			data.TunnelRouteViaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelRouteVia = types.StringNull()
			data.TunnelRouteViaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-route-via.vipValue")
			data.TunnelRouteVia = types.StringValue(v.String())
			data.TunnelRouteViaVariable = types.StringNull()
		}
	} else {
		data.TunnelRouteVia = types.StringNull()
		data.TunnelRouteViaVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoVPNInterfaceGRE) hasChanges(ctx context.Context, state *CiscoVPNInterfaceGRE) bool {
	hasChanges := false
	if !data.InterfaceName.Equal(state.InterfaceName) {
		hasChanges = true
	}
	if !data.InterfaceDescription.Equal(state.InterfaceDescription) {
		hasChanges = true
	}
	if !data.IpAddress.Equal(state.IpAddress) {
		hasChanges = true
	}
	if !data.TunnelSource.Equal(state.TunnelSource) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.TunnelSourceInterface.Equal(state.TunnelSourceInterface) {
		hasChanges = true
	}
	if !data.TunnelDestination.Equal(state.TunnelDestination) {
		hasChanges = true
	}
	if !data.Application.Equal(state.Application) {
		hasChanges = true
	}
	if !data.IpMtu.Equal(state.IpMtu) {
		hasChanges = true
	}
	if !data.TcpMssAdjust.Equal(state.TcpMssAdjust) {
		hasChanges = true
	}
	if !data.ClearDontFragment.Equal(state.ClearDontFragment) {
		hasChanges = true
	}
	if !data.RewriteRule.Equal(state.RewriteRule) {
		hasChanges = true
	}
	if len(data.AccessLists) != len(state.AccessLists) {
		hasChanges = true
	} else {
		for i := range data.AccessLists {
			if !data.AccessLists[i].Direction.Equal(state.AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.AccessLists[i].AclName.Equal(state.AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if !data.Tracker.Equal(state.Tracker) {
		hasChanges = true
	}
	if !data.TunnelRouteVia.Equal(state.TunnelRouteVia) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
