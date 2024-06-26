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
type CellularCEdgeProfile struct {
	Id                            types.String `tfsdk:"id"`
	Version                       types.Int64  `tfsdk:"version"`
	TemplateType                  types.String `tfsdk:"template_type"`
	Name                          types.String `tfsdk:"name"`
	Description                   types.String `tfsdk:"description"`
	DeviceTypes                   types.Set    `tfsdk:"device_types"`
	ProfileId                     types.Int64  `tfsdk:"profile_id"`
	ProfileIdVariable             types.String `tfsdk:"profile_id_variable"`
	AccessPointName               types.String `tfsdk:"access_point_name"`
	AccessPointNameVariable       types.String `tfsdk:"access_point_name_variable"`
	AuthenticationType            types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable    types.String `tfsdk:"authentication_type_variable"`
	PacketDataNetworkType         types.String `tfsdk:"packet_data_network_type"`
	PacketDataNetworkTypeVariable types.String `tfsdk:"packet_data_network_type_variable"`
	ProfileUsername               types.String `tfsdk:"profile_username"`
	ProfileUsernameVariable       types.String `tfsdk:"profile_username_variable"`
	ProfilePassword               types.String `tfsdk:"profile_password"`
	ProfilePasswordVariable       types.String `tfsdk:"profile_password_variable"`
	NoOverwrite                   types.Bool   `tfsdk:"no_overwrite"`
	NoOverwriteVariable           types.String `tfsdk:"no_overwrite_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CellularCEdgeProfile) getModel() string {
	return "cellular-cedge-profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CellularCEdgeProfile) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cellular-cedge-profile")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.ProfileIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"id."+"vipVariableName", data.ProfileIdVariable.ValueString())
	} else if data.ProfileId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"id."+"vipValue", data.ProfileId.ValueInt64())
	}

	if !data.AccessPointNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"apn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"apn."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"apn."+"vipVariableName", data.AccessPointNameVariable.ValueString())
	} else if data.AccessPointName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"apn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"apn."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"apn."+"vipValue", data.AccessPointName.ValueString())
	}

	if !data.AuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"authentication."+"vipVariableName", data.AuthenticationTypeVariable.ValueString())
	} else if data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"authentication."+"vipValue", data.AuthenticationType.ValueString())
	}

	if !data.PacketDataNetworkTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pdn-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pdn-type."+"vipVariableName", data.PacketDataNetworkTypeVariable.ValueString())
	} else if data.PacketDataNetworkType.IsNull() {
		body, _ = sjson.Set(body, path+"pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pdn-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pdn-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pdn-type."+"vipValue", data.PacketDataNetworkType.ValueString())
	}

	if !data.ProfileUsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"username."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"username."+"vipVariableName", data.ProfileUsernameVariable.ValueString())
	} else if data.ProfileUsername.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"username."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"username."+"vipValue", data.ProfileUsername.ValueString())
	}

	if !data.ProfilePasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"password."+"vipVariableName", data.ProfilePasswordVariable.ValueString())
	} else if data.ProfilePassword.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"password."+"vipValue", data.ProfilePassword.ValueString())
	}

	if !data.NoOverwriteVariable.IsNull() {
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipVariableName", data.NoOverwriteVariable.ValueString())
	} else if data.NoOverwrite.IsNull() {
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"no-overwrite."+"vipValue", strconv.FormatBool(data.NoOverwrite.ValueBool()))
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CellularCEdgeProfile) fromBody(ctx context.Context, res gjson.Result) {
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
			data.ProfileId = types.Int64Null()

			v := res.Get(path + "id.vipVariableName")
			data.ProfileIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfileId = types.Int64Null()
			data.ProfileIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "id.vipValue")
			data.ProfileId = types.Int64Value(v.Int())
			data.ProfileIdVariable = types.StringNull()
		}
	} else {
		data.ProfileId = types.Int64Null()
		data.ProfileIdVariable = types.StringNull()
	}
	if value := res.Get(path + "apn.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AccessPointName = types.StringNull()

			v := res.Get(path + "apn.vipVariableName")
			data.AccessPointNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AccessPointName = types.StringNull()
			data.AccessPointNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "apn.vipValue")
			data.AccessPointName = types.StringValue(v.String())
			data.AccessPointNameVariable = types.StringNull()
		}
	} else {
		data.AccessPointName = types.StringNull()
		data.AccessPointNameVariable = types.StringNull()
	}
	if value := res.Get(path + "authentication.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AuthenticationType = types.StringNull()

			v := res.Get(path + "authentication.vipVariableName")
			data.AuthenticationTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthenticationType = types.StringNull()
			data.AuthenticationTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "authentication.vipValue")
			data.AuthenticationType = types.StringValue(v.String())
			data.AuthenticationTypeVariable = types.StringNull()
		}
	} else {
		data.AuthenticationType = types.StringNull()
		data.AuthenticationTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "pdn-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PacketDataNetworkType = types.StringNull()

			v := res.Get(path + "pdn-type.vipVariableName")
			data.PacketDataNetworkTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PacketDataNetworkType = types.StringNull()
			data.PacketDataNetworkTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pdn-type.vipValue")
			data.PacketDataNetworkType = types.StringValue(v.String())
			data.PacketDataNetworkTypeVariable = types.StringNull()
		}
	} else {
		data.PacketDataNetworkType = types.StringNull()
		data.PacketDataNetworkTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "username.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfileUsername = types.StringNull()

			v := res.Get(path + "username.vipVariableName")
			data.ProfileUsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfileUsername = types.StringNull()
			data.ProfileUsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "username.vipValue")
			data.ProfileUsername = types.StringValue(v.String())
			data.ProfileUsernameVariable = types.StringNull()
		}
	} else {
		data.ProfileUsername = types.StringNull()
		data.ProfileUsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfilePassword = types.StringNull()

			v := res.Get(path + "password.vipVariableName")
			data.ProfilePasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfilePassword = types.StringNull()
			data.ProfilePasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "password.vipValue")
			data.ProfilePassword = types.StringValue(v.String())
			data.ProfilePasswordVariable = types.StringNull()
		}
	} else {
		data.ProfilePassword = types.StringNull()
		data.ProfilePasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "no-overwrite.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NoOverwrite = types.BoolNull()

			v := res.Get(path + "no-overwrite.vipVariableName")
			data.NoOverwriteVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NoOverwrite = types.BoolNull()
			data.NoOverwriteVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "no-overwrite.vipValue")
			data.NoOverwrite = types.BoolValue(v.Bool())
			data.NoOverwriteVariable = types.StringNull()
		}
	} else {
		data.NoOverwrite = types.BoolNull()
		data.NoOverwriteVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CellularCEdgeProfile) hasChanges(ctx context.Context, state *CellularCEdgeProfile) bool {
	hasChanges := false
	if !data.ProfileId.Equal(state.ProfileId) {
		hasChanges = true
	}
	if !data.AccessPointName.Equal(state.AccessPointName) {
		hasChanges = true
	}
	if !data.AuthenticationType.Equal(state.AuthenticationType) {
		hasChanges = true
	}
	if !data.PacketDataNetworkType.Equal(state.PacketDataNetworkType) {
		hasChanges = true
	}
	if !data.ProfileUsername.Equal(state.ProfileUsername) {
		hasChanges = true
	}
	if !data.ProfilePassword.Equal(state.ProfilePassword) {
		hasChanges = true
	}
	if !data.NoOverwrite.Equal(state.NoOverwrite) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
