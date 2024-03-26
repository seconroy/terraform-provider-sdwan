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

type Logging struct {
	Id                      types.String             `tfsdk:"id"`
	Version                 types.Int64              `tfsdk:"version"`
	TemplateType            types.String             `tfsdk:"template_type"`
	Name                    types.String             `tfsdk:"name"`
	Description             types.String             `tfsdk:"description"`
	DeviceTypes             types.Set                `tfsdk:"device_types"`
	Enable                  types.Bool               `tfsdk:"enable"`
	EnableVariable          types.String             `tfsdk:"enable_variable"`
	MaximumFileSize         types.Int64              `tfsdk:"maximum_file_size"`
	MaximumFileSizeVariable types.String             `tfsdk:"maximum_file_size_variable"`
	Rotation                types.Int64              `tfsdk:"rotation"`
	RotationVariable        types.String             `tfsdk:"rotation_variable"`
	Priority                types.String             `tfsdk:"priority"`
	PriorityVariable        types.String             `tfsdk:"priority_variable"`
	RemoteHosts             []LoggingRemoteHosts     `tfsdk:"remote_hosts"`
	RemoteIpv6Hosts         []LoggingRemoteIpv6Hosts `tfsdk:"remote_ipv6_hosts"`
}

type LoggingRemoteHosts struct {
	Optional                    types.Bool   `tfsdk:"optional"`
	HostnameIpv4Address         types.String `tfsdk:"hostname_ipv4_address"`
	HostnameIpv4AddressVariable types.String `tfsdk:"hostname_ipv4_address_variable"`
	VpnId                       types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable               types.String `tfsdk:"vpn_id_variable"`
	SourceInterface             types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable     types.String `tfsdk:"source_interface_variable"`
	Priority                    types.String `tfsdk:"priority"`
	PriorityVariable            types.String `tfsdk:"priority_variable"`
}

type LoggingRemoteIpv6Hosts struct {
	Optional                        types.Bool   `tfsdk:"optional"`
	Ipv6HostnameIpv6Address         types.String `tfsdk:"ipv6_hostname_ipv6_address"`
	Ipv6HostnameIpv6AddressVariable types.String `tfsdk:"ipv6_hostname_ipv6_address_variable"`
	VpnId                           types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable                   types.String `tfsdk:"vpn_id_variable"`
	SourceInterface                 types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable         types.String `tfsdk:"source_interface_variable"`
	Priority                        types.String `tfsdk:"priority"`
	PriorityVariable                types.String `tfsdk:"priority_variable"`
}

func (data Logging) getModel() string {
	return "logging"
}

func (data Logging) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "logging")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.EnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipVariableName", data.EnableVariable.ValueString())
	} else if data.Enable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipValue", strconv.FormatBool(data.Enable.ValueBool()))
	}

	if !data.MaximumFileSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipVariableName", data.MaximumFileSizeVariable.ValueString())
	} else if data.MaximumFileSize.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipValue", data.MaximumFileSize.ValueInt64())
	}

	if !data.RotationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipVariableName", data.RotationVariable.ValueString())
	} else if data.Rotation.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipValue", data.Rotation.ValueInt64())
	}

	if !data.PriorityVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.priority."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.priority."+"vipVariableName", data.PriorityVariable.ValueString())
	} else if data.Priority.IsNull() {
		body, _ = sjson.Set(body, path+"disk.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.priority."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.priority."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.priority."+"vipValue", data.Priority.ValueString())
	}
	if len(data.RemoteHosts) > 0 {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	}
	for _, item := range data.RemoteHosts {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.HostnameIpv4AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.HostnameIpv4AddressVariable.ValueString())
		} else if item.HostnameIpv4Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.HostnameIpv4Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"server."+"vipValue.-1", itemBody)
	}
	if len(data.RemoteIpv6Hosts) > 0 {
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipValue", []interface{}{})
	}
	for _, item := range data.RemoteIpv6Hosts {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.Ipv6HostnameIpv6AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.Ipv6HostnameIpv6AddressVariable.ValueString())
		} else if item.Ipv6HostnameIpv6Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Ipv6HostnameIpv6Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6-server."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *Logging) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "disk.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Enable = types.BoolNull()

			v := res.Get(path + "disk.enable.vipVariableName")
			data.EnableVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Enable = types.BoolNull()
			data.EnableVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.enable.vipValue")
			data.Enable = types.BoolValue(v.Bool())
			data.EnableVariable = types.StringNull()
		}
	} else {
		data.Enable = types.BoolNull()
		data.EnableVariable = types.StringNull()
	}
	if value := res.Get(path + "disk.file.size.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MaximumFileSize = types.Int64Null()

			v := res.Get(path + "disk.file.size.vipVariableName")
			data.MaximumFileSizeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MaximumFileSize = types.Int64Null()
			data.MaximumFileSizeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.file.size.vipValue")
			data.MaximumFileSize = types.Int64Value(v.Int())
			data.MaximumFileSizeVariable = types.StringNull()
		}
	} else {
		data.MaximumFileSize = types.Int64Null()
		data.MaximumFileSizeVariable = types.StringNull()
	}
	if value := res.Get(path + "disk.file.rotate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Rotation = types.Int64Null()

			v := res.Get(path + "disk.file.rotate.vipVariableName")
			data.RotationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Rotation = types.Int64Null()
			data.RotationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.file.rotate.vipValue")
			data.Rotation = types.Int64Value(v.Int())
			data.RotationVariable = types.StringNull()
		}
	} else {
		data.Rotation = types.Int64Null()
		data.RotationVariable = types.StringNull()
	}
	if value := res.Get(path + "disk.priority.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Priority = types.StringNull()

			v := res.Get(path + "disk.priority.vipVariableName")
			data.PriorityVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Priority = types.StringNull()
			data.PriorityVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.priority.vipValue")
			data.Priority = types.StringValue(v.String())
			data.PriorityVariable = types.StringNull()
		}
	} else {
		data.Priority = types.StringNull()
		data.PriorityVariable = types.StringNull()
	}
	if value := res.Get(path + "server.vipValue"); len(value.Array()) > 0 {
		data.RemoteHosts = make([]LoggingRemoteHosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingRemoteHosts{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HostnameIpv4Address = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.HostnameIpv4AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HostnameIpv4Address = types.StringNull()
					item.HostnameIpv4AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.HostnameIpv4Address = types.StringValue(cv.String())
					item.HostnameIpv4AddressVariable = types.StringNull()
				}
			} else {
				item.HostnameIpv4Address = types.StringNull()
				item.HostnameIpv4AddressVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.StringNull()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.StringNull()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.StringValue(cv.String())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.StringNull()
				item.PriorityVariable = types.StringNull()
			}
			data.RemoteHosts = append(data.RemoteHosts, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6-server.vipValue"); len(value.Array()) > 0 {
		data.RemoteIpv6Hosts = make([]LoggingRemoteIpv6Hosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingRemoteIpv6Hosts{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv6HostnameIpv6Address = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.Ipv6HostnameIpv6AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv6HostnameIpv6Address = types.StringNull()
					item.Ipv6HostnameIpv6AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Ipv6HostnameIpv6Address = types.StringValue(cv.String())
					item.Ipv6HostnameIpv6AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv6HostnameIpv6Address = types.StringNull()
				item.Ipv6HostnameIpv6AddressVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.StringNull()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.StringNull()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.StringValue(cv.String())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.StringNull()
				item.PriorityVariable = types.StringNull()
			}
			data.RemoteIpv6Hosts = append(data.RemoteIpv6Hosts, item)
			return true
		})
	}
}

func (data *Logging) hasChanges(ctx context.Context, state *Logging) bool {
	hasChanges := false
	if !data.Enable.Equal(state.Enable) {
		hasChanges = true
	}
	if !data.MaximumFileSize.Equal(state.MaximumFileSize) {
		hasChanges = true
	}
	if !data.Rotation.Equal(state.Rotation) {
		hasChanges = true
	}
	if !data.Priority.Equal(state.Priority) {
		hasChanges = true
	}
	if len(data.RemoteHosts) != len(state.RemoteHosts) {
		hasChanges = true
	} else {
		for i := range data.RemoteHosts {
			if !data.RemoteHosts[i].HostnameIpv4Address.Equal(state.RemoteHosts[i].HostnameIpv4Address) {
				hasChanges = true
			}
			if !data.RemoteHosts[i].VpnId.Equal(state.RemoteHosts[i].VpnId) {
				hasChanges = true
			}
			if !data.RemoteHosts[i].SourceInterface.Equal(state.RemoteHosts[i].SourceInterface) {
				hasChanges = true
			}
			if !data.RemoteHosts[i].Priority.Equal(state.RemoteHosts[i].Priority) {
				hasChanges = true
			}
		}
	}
	if len(data.RemoteIpv6Hosts) != len(state.RemoteIpv6Hosts) {
		hasChanges = true
	} else {
		for i := range data.RemoteIpv6Hosts {
			if !data.RemoteIpv6Hosts[i].Ipv6HostnameIpv6Address.Equal(state.RemoteIpv6Hosts[i].Ipv6HostnameIpv6Address) {
				hasChanges = true
			}
			if !data.RemoteIpv6Hosts[i].VpnId.Equal(state.RemoteIpv6Hosts[i].VpnId) {
				hasChanges = true
			}
			if !data.RemoteIpv6Hosts[i].SourceInterface.Equal(state.RemoteIpv6Hosts[i].SourceInterface) {
				hasChanges = true
			}
			if !data.RemoteIpv6Hosts[i].Priority.Equal(state.RemoteIpv6Hosts[i].Priority) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
