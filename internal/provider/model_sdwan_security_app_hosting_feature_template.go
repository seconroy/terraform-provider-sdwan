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
type SecurityAppHosting struct {
	Id                  types.String                            `tfsdk:"id"`
	Version             types.Int64                             `tfsdk:"version"`
	TemplateType        types.String                            `tfsdk:"template_type"`
	Name                types.String                            `tfsdk:"name"`
	Description         types.String                            `tfsdk:"description"`
	DeviceTypes         types.Set                               `tfsdk:"device_types"`
	VirtualApplications []SecurityAppHostingVirtualApplications `tfsdk:"virtual_applications"`
}

type SecurityAppHostingVirtualApplications struct {
	Optional                 types.Bool   `tfsdk:"optional"`
	InstanceId               types.String `tfsdk:"instance_id"`
	ApplicationType          types.String `tfsdk:"application_type"`
	Nat                      types.Bool   `tfsdk:"nat"`
	NatVariable              types.String `tfsdk:"nat_variable"`
	DatabaseUrl              types.Bool   `tfsdk:"database_url"`
	DatabaseUrlVariable      types.String `tfsdk:"database_url_variable"`
	ResourceProfile          types.String `tfsdk:"resource_profile"`
	ResourceProfileVariable  types.String `tfsdk:"resource_profile_variable"`
	ServiceGatewayIp         types.String `tfsdk:"service_gateway_ip"`
	ServiceGatewayIpVariable types.String `tfsdk:"service_gateway_ip_variable"`
	ServiceIp                types.String `tfsdk:"service_ip"`
	ServiceIpVariable        types.String `tfsdk:"service_ip_variable"`
	DataGatewayIp            types.String `tfsdk:"data_gateway_ip"`
	DataGatewayIpVariable    types.String `tfsdk:"data_gateway_ip_variable"`
	DataServiceIp            types.String `tfsdk:"data_service_ip"`
	DataServiceIpVariable    types.String `tfsdk:"data_service_ip_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SecurityAppHosting) getModel() string {
	return "virtual-application-utd"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SecurityAppHosting) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "virtual-application-utd")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if len(data.VirtualApplications) > 0 {
		body, _ = sjson.Set(body, path+"virtual-application."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipPrimaryKey", []string{"instance-id"})
		body, _ = sjson.Set(body, path+"virtual-application."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"virtual-application."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipPrimaryKey", []string{"instance-id"})
		body, _ = sjson.Set(body, path+"virtual-application."+"vipValue", []interface{}{})
	}
	for _, item := range data.VirtualApplications {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "instance-id")
		if item.InstanceId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipValue", item.InstanceId.ValueString())
		}
		itemAttributes = append(itemAttributes, "application-type")
		if item.ApplicationType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipValue", item.ApplicationType.ValueString())
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.NatVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipVariableName", item.NatVariable.ValueString())
		} else if item.Nat.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.nat."+"vipValue", strconv.FormatBool(item.Nat.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.DatabaseUrlVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipVariableName", item.DatabaseUrlVariable.ValueString())
		} else if item.DatabaseUrl.IsNull() {
			if !gjson.Get(itemBody, "utd").Exists() {
				itemBody, _ = sjson.Set(itemBody, "utd", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.database-url."+"vipValue", strconv.FormatBool(item.DatabaseUrl.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.ResourceProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipVariableName", item.ResourceProfileVariable.ValueString())
		} else if item.ResourceProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.resource-profile."+"vipValue", item.ResourceProfile.ValueString())
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.ServiceGatewayIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipVariableName", item.ServiceGatewayIpVariable.ValueString())
		} else if item.ServiceGatewayIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.service-gateway-ip."+"vipValue", item.ServiceGatewayIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.ServiceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipVariableName", item.ServiceIpVariable.ValueString())
		} else if item.ServiceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.service-ip."+"vipValue", item.ServiceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.DataGatewayIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipVariableName", item.DataGatewayIpVariable.ValueString())
		} else if item.DataGatewayIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.data-gateway-ip."+"vipValue", item.DataGatewayIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "utd")

		if !item.DataServiceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipVariableName", item.DataServiceIpVariable.ValueString())
		} else if item.DataServiceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "utd.data-service-ip."+"vipValue", item.DataServiceIp.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"virtual-application."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SecurityAppHosting) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "virtual-application.vipValue"); len(value.Array()) > 0 {
		data.VirtualApplications = make([]SecurityAppHostingVirtualApplications, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SecurityAppHostingVirtualApplications{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("instance-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InstanceId = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.InstanceId = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("instance-id.vipValue")
					item.InstanceId = types.StringValue(cv.String())

				}
			} else {
				item.InstanceId = types.StringNull()

			}
			if cValue := v.Get("application-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ApplicationType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ApplicationType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("application-type.vipValue")
					item.ApplicationType = types.StringValue(cv.String())

				}
			} else {
				item.ApplicationType = types.StringNull()

			}
			if cValue := v.Get("utd.nat.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Nat = types.BoolNull()

					cv := v.Get("utd.nat.vipVariableName")
					item.NatVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Nat = types.BoolNull()
					item.NatVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.nat.vipValue")
					item.Nat = types.BoolValue(cv.Bool())
					item.NatVariable = types.StringNull()
				}
			} else {
				item.Nat = types.BoolNull()
				item.NatVariable = types.StringNull()
			}
			if cValue := v.Get("utd.database-url.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DatabaseUrl = types.BoolNull()

					cv := v.Get("utd.database-url.vipVariableName")
					item.DatabaseUrlVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DatabaseUrl = types.BoolNull()
					item.DatabaseUrlVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.database-url.vipValue")
					item.DatabaseUrl = types.BoolValue(cv.Bool())
					item.DatabaseUrlVariable = types.StringNull()
				}
			} else {
				item.DatabaseUrl = types.BoolNull()
				item.DatabaseUrlVariable = types.StringNull()
			}
			if cValue := v.Get("utd.resource-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ResourceProfile = types.StringNull()

					cv := v.Get("utd.resource-profile.vipVariableName")
					item.ResourceProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ResourceProfile = types.StringNull()
					item.ResourceProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.resource-profile.vipValue")
					item.ResourceProfile = types.StringValue(cv.String())
					item.ResourceProfileVariable = types.StringNull()
				}
			} else {
				item.ResourceProfile = types.StringNull()
				item.ResourceProfileVariable = types.StringNull()
			}
			if cValue := v.Get("utd.service-gateway-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ServiceGatewayIp = types.StringNull()

					cv := v.Get("utd.service-gateway-ip.vipVariableName")
					item.ServiceGatewayIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ServiceGatewayIp = types.StringNull()
					item.ServiceGatewayIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.service-gateway-ip.vipValue")
					item.ServiceGatewayIp = types.StringValue(cv.String())
					item.ServiceGatewayIpVariable = types.StringNull()
				}
			} else {
				item.ServiceGatewayIp = types.StringNull()
				item.ServiceGatewayIpVariable = types.StringNull()
			}
			if cValue := v.Get("utd.service-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ServiceIp = types.StringNull()

					cv := v.Get("utd.service-ip.vipVariableName")
					item.ServiceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ServiceIp = types.StringNull()
					item.ServiceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.service-ip.vipValue")
					item.ServiceIp = types.StringValue(cv.String())
					item.ServiceIpVariable = types.StringNull()
				}
			} else {
				item.ServiceIp = types.StringNull()
				item.ServiceIpVariable = types.StringNull()
			}
			if cValue := v.Get("utd.data-gateway-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DataGatewayIp = types.StringNull()

					cv := v.Get("utd.data-gateway-ip.vipVariableName")
					item.DataGatewayIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DataGatewayIp = types.StringNull()
					item.DataGatewayIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.data-gateway-ip.vipValue")
					item.DataGatewayIp = types.StringValue(cv.String())
					item.DataGatewayIpVariable = types.StringNull()
				}
			} else {
				item.DataGatewayIp = types.StringNull()
				item.DataGatewayIpVariable = types.StringNull()
			}
			if cValue := v.Get("utd.data-service-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DataServiceIp = types.StringNull()

					cv := v.Get("utd.data-service-ip.vipVariableName")
					item.DataServiceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DataServiceIp = types.StringNull()
					item.DataServiceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("utd.data-service-ip.vipValue")
					item.DataServiceIp = types.StringValue(cv.String())
					item.DataServiceIpVariable = types.StringNull()
				}
			} else {
				item.DataServiceIp = types.StringNull()
				item.DataServiceIpVariable = types.StringNull()
			}
			data.VirtualApplications = append(data.VirtualApplications, item)
			return true
		})
	} else {
		if len(data.VirtualApplications) > 0 {
			data.VirtualApplications = []SecurityAppHostingVirtualApplications{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *SecurityAppHosting) hasChanges(ctx context.Context, state *SecurityAppHosting) bool {
	hasChanges := false
	if len(data.VirtualApplications) != len(state.VirtualApplications) {
		hasChanges = true
	} else {
		for i := range data.VirtualApplications {
			if !data.VirtualApplications[i].InstanceId.Equal(state.VirtualApplications[i].InstanceId) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].ApplicationType.Equal(state.VirtualApplications[i].ApplicationType) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].Nat.Equal(state.VirtualApplications[i].Nat) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].DatabaseUrl.Equal(state.VirtualApplications[i].DatabaseUrl) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].ResourceProfile.Equal(state.VirtualApplications[i].ResourceProfile) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].ServiceGatewayIp.Equal(state.VirtualApplications[i].ServiceGatewayIp) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].ServiceIp.Equal(state.VirtualApplications[i].ServiceIp) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].DataGatewayIp.Equal(state.VirtualApplications[i].DataGatewayIp) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].DataServiceIp.Equal(state.VirtualApplications[i].DataServiceIp) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
