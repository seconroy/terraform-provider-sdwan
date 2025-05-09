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
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportCellularProfile struct {
	Id                            types.String `tfsdk:"id"`
	Version                       types.Int64  `tfsdk:"version"`
	Name                          types.String `tfsdk:"name"`
	Description                   types.String `tfsdk:"description"`
	FeatureProfileId              types.String `tfsdk:"feature_profile_id"`
	ProfileId                     types.Int64  `tfsdk:"profile_id"`
	ProfileIdVariable             types.String `tfsdk:"profile_id_variable"`
	AccessPointName               types.String `tfsdk:"access_point_name"`
	AccessPointNameVariable       types.String `tfsdk:"access_point_name_variable"`
	RequiresAuthentication        types.Bool   `tfsdk:"requires_authentication"`
	AuthenticationType            types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable    types.String `tfsdk:"authentication_type_variable"`
	ProfileUsername               types.String `tfsdk:"profile_username"`
	ProfileUsernameVariable       types.String `tfsdk:"profile_username_variable"`
	ProfilePassword               types.String `tfsdk:"profile_password"`
	ProfilePasswordVariable       types.String `tfsdk:"profile_password_variable"`
	PacketDataNetworkType         types.String `tfsdk:"packet_data_network_type"`
	PacketDataNetworkTypeVariable types.String `tfsdk:"packet_data_network_type_variable"`
	NoOverwrite                   types.Bool   `tfsdk:"no_overwrite"`
	NoOverwriteVariable           types.String `tfsdk:"no_overwrite_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportCellularProfile) getModel() string {
	return "transport_cellular_profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportCellularProfile) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/cellular-profile", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportCellularProfile) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"configType.optionType", "default")
		body, _ = sjson.Set(body, path+"configType.value", "non-eSim")
	}

	if !data.ProfileIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.id.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.id.value", data.ProfileIdVariable.ValueString())
		}
	} else if !data.ProfileId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.id.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.id.value", data.ProfileId.ValueInt64())
		}
	}

	if !data.AccessPointNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.value", data.AccessPointNameVariable.ValueString())
		}
	} else if !data.AccessPointName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.value", data.AccessPointName.ValueString())
		}
	}
	if true && data.RequiresAuthentication.ValueBool() == false {
		body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.noAuthentication.optionType", "default")
		body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.noAuthentication.value", "none")
	}

	if !data.AuthenticationTypeVariable.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.type.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.type.value", data.AuthenticationTypeVariable.ValueString())
		}
	} else if !data.AuthenticationType.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.type.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.type.value", data.AuthenticationType.ValueString())
		}
	}

	if !data.ProfileUsernameVariable.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.username.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.username.value", data.ProfileUsernameVariable.ValueString())
		}
	} else if !data.ProfileUsername.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.username.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.username.value", data.ProfileUsername.ValueString())
		}
	}

	if !data.ProfilePasswordVariable.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.password.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.password.value", data.ProfilePasswordVariable.ValueString())
		}
	} else if !data.ProfilePassword.IsNull() {
		if true && data.RequiresAuthentication.ValueBool() == true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.password.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.password.value", data.ProfilePassword.ValueString())
		}
	}

	if !data.PacketDataNetworkTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", data.PacketDataNetworkTypeVariable.ValueString())
		}
	} else if data.PacketDataNetworkType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "default")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", "ipv4")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", data.PacketDataNetworkType.ValueString())
		}
	}

	if !data.NoOverwriteVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.value", data.NoOverwriteVariable.ValueString())
		}
	} else if data.NoOverwrite.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.value", data.NoOverwrite.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportCellularProfile) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ProfileId = types.Int64Null()
	data.ProfileIdVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.id.value")
		if t.String() == "variable" {
			data.ProfileIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileId = types.Int64Value(va.Int())
		}
	}
	data.AccessPointName = types.StringNull()
	data.AccessPointNameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.apn.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.apn.value")
		if t.String() == "variable" {
			data.AccessPointNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccessPointName = types.StringValue(va.String())
		}
	}
	data.AuthenticationType = types.StringNull()
	data.AuthenticationTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.type.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.type.value")
		if t.String() == "variable" {
			data.AuthenticationTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthenticationType = types.StringValue(va.String())
		}
	}
	data.ProfileUsername = types.StringNull()
	data.ProfileUsernameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.username.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.username.value")
		if t.String() == "variable" {
			data.ProfileUsernameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileUsername = types.StringValue(va.String())
		}
	}
	data.PacketDataNetworkType = types.StringNull()
	data.PacketDataNetworkTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.pdnType.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.pdnType.value")
		if t.String() == "variable" {
			data.PacketDataNetworkTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PacketDataNetworkType = types.StringValue(va.String())
		}
	}
	data.NoOverwrite = types.BoolNull()
	data.NoOverwriteVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.noOverwrite.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.noOverwrite.value")
		if t.String() == "variable" {
			data.NoOverwriteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NoOverwrite = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportCellularProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ProfileId = types.Int64Null()
	data.ProfileIdVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.id.value")
		if t.String() == "variable" {
			data.ProfileIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileId = types.Int64Value(va.Int())
		}
	}
	data.AccessPointName = types.StringNull()
	data.AccessPointNameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.apn.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.apn.value")
		if t.String() == "variable" {
			data.AccessPointNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccessPointName = types.StringValue(va.String())
		}
	}
	data.AuthenticationType = types.StringNull()
	data.AuthenticationTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.type.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.type.value")
		if t.String() == "variable" {
			data.AuthenticationTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthenticationType = types.StringValue(va.String())
		}
	}
	data.ProfileUsername = types.StringNull()
	data.ProfileUsernameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.username.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.username.value")
		if t.String() == "variable" {
			data.ProfileUsernameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileUsername = types.StringValue(va.String())
		}
	}
	data.PacketDataNetworkType = types.StringNull()
	data.PacketDataNetworkTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.pdnType.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.pdnType.value")
		if t.String() == "variable" {
			data.PacketDataNetworkTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PacketDataNetworkType = types.StringValue(va.String())
		}
	}
	data.NoOverwrite = types.BoolNull()
	data.NoOverwriteVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.noOverwrite.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.noOverwrite.value")
		if t.String() == "variable" {
			data.NoOverwriteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NoOverwrite = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody
