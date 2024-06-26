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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Device struct {
	Id           types.String    `tfsdk:"id"`
	SerialNumber types.String    `tfsdk:"serial_number"`
	Name         types.String    `tfsdk:"name"`
	Devices      []DeviceDevices `tfsdk:"devices"`
}

type DeviceDevices struct {
	DeviceId     types.String `tfsdk:"device_id"`
	Uuid         types.String `tfsdk:"uuid"`
	SiteId       types.String `tfsdk:"site_id"`
	SerialNumber types.String `tfsdk:"serial_number"`
	Hostname     types.String `tfsdk:"hostname"`
	Reachability types.String `tfsdk:"reachability"`
	Status       types.String `tfsdk:"status"`
	State        types.String `tfsdk:"state"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Device) getPath() string {
	return "/device"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Device) toBody(ctx context.Context) string {
	body := ""
	if !data.SerialNumber.IsNull() {
		body, _ = sjson.Set(body, "board-serial", data.SerialNumber.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "host-name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "data", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.DeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceId", item.DeviceId.ValueString())
			}
			if !item.Uuid.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "uuid", item.Uuid.ValueString())
			}
			if !item.SiteId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "site-id", item.SiteId.ValueString())
			}
			if !item.SerialNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "board-serial", item.SerialNumber.ValueString())
			}
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host-name", item.Hostname.ValueString())
			}
			if !item.Reachability.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "reachability", item.Reachability.ValueString())
			}
			if !item.Status.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "status", item.Status.ValueString())
			}
			if !item.State.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "state", item.State.ValueString())
			}
			body, _ = sjson.SetRaw(body, "data.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("board-serial"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("host-name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("data"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]DeviceDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeviceDevices{}
			if cValue := v.Get("deviceId"); cValue.Exists() {
				item.DeviceId = types.StringValue(cValue.String())
			} else {
				item.DeviceId = types.StringNull()
			}
			if cValue := v.Get("uuid"); cValue.Exists() {
				item.Uuid = types.StringValue(cValue.String())
			} else {
				item.Uuid = types.StringNull()
			}
			if cValue := v.Get("site-id"); cValue.Exists() {
				item.SiteId = types.StringValue(cValue.String())
			} else {
				item.SiteId = types.StringNull()
			}
			if cValue := v.Get("board-serial"); cValue.Exists() {
				item.SerialNumber = types.StringValue(cValue.String())
			} else {
				item.SerialNumber = types.StringNull()
			}
			if cValue := v.Get("host-name"); cValue.Exists() {
				item.Hostname = types.StringValue(cValue.String())
			} else {
				item.Hostname = types.StringNull()
			}
			if cValue := v.Get("reachability"); cValue.Exists() {
				item.Reachability = types.StringValue(cValue.String())
			} else {
				item.Reachability = types.StringNull()
			}
			if cValue := v.Get("status"); cValue.Exists() {
				item.Status = types.StringValue(cValue.String())
			} else {
				item.Status = types.StringNull()
			}
			if cValue := v.Get("state"); cValue.Exists() {
				item.State = types.StringValue(cValue.String())
			} else {
				item.State = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []DeviceDevices{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Device) hasChanges(ctx context.Context, state *Device) bool {
	hasChanges := false
	if !data.SerialNumber.Equal(state.SerialNumber) {
		hasChanges = true
	}
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Devices) != len(state.Devices) {
		hasChanges = true
	} else {
		for i := range data.Devices {
			if !data.Devices[i].DeviceId.Equal(state.Devices[i].DeviceId) {
				hasChanges = true
			}
			if !data.Devices[i].Uuid.Equal(state.Devices[i].Uuid) {
				hasChanges = true
			}
			if !data.Devices[i].SiteId.Equal(state.Devices[i].SiteId) {
				hasChanges = true
			}
			if !data.Devices[i].SerialNumber.Equal(state.Devices[i].SerialNumber) {
				hasChanges = true
			}
			if !data.Devices[i].Hostname.Equal(state.Devices[i].Hostname) {
				hasChanges = true
			}
			if !data.Devices[i].Reachability.Equal(state.Devices[i].Reachability) {
				hasChanges = true
			}
			if !data.Devices[i].Status.Equal(state.Devices[i].Status) {
				hasChanges = true
			}
			if !data.Devices[i].State.Equal(state.Devices[i].State) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
