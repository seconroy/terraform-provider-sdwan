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
type ConfigurationGroup struct {
	Id                  types.String                        `tfsdk:"id"`
	Name                types.String                        `tfsdk:"name"`
	Description         types.String                        `tfsdk:"description"`
	Solution            types.String                        `tfsdk:"solution"`
	FeatureProfiles     []ConfigurationGroupFeatureProfiles `tfsdk:"feature_profiles"`
	TopologyDevices     []ConfigurationGroupTopologyDevices `tfsdk:"topology_devices"`
	TopologySiteDevices types.Int64                         `tfsdk:"topology_site_devices"`
}

type ConfigurationGroupFeatureProfiles struct {
	Id types.String `tfsdk:"id"`
}

type ConfigurationGroupTopologyDevices struct {
	CriteriaAttribute   types.String                                           `tfsdk:"criteria_attribute"`
	CriteriaValue       types.String                                           `tfsdk:"criteria_value"`
	UnsupportedFeatures []ConfigurationGroupTopologyDevicesUnsupportedFeatures `tfsdk:"unsupported_features"`
}

type ConfigurationGroupTopologyDevicesUnsupportedFeatures struct {
	ParcelType types.String `tfsdk:"parcel_type"`
	ParcelId   types.String `tfsdk:"parcel_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ConfigurationGroup) getPath() string {
	return "/v1/config-group/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ConfigurationGroup) toBody(ctx context.Context) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "profiles", []interface{}{})
		for _, item := range data.FeatureProfiles {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "profiles.-1", itemBody)
		}
	}
	if true && len(data.TopologyDevices) > 0 {
		body, _ = sjson.Set(body, "topology.devices", []interface{}{})
		for _, item := range data.TopologyDevices {
			itemBody := ""
			if !item.CriteriaAttribute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.attribute", item.CriteriaAttribute.ValueString())
			}
			if !item.CriteriaValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.value", item.CriteriaValue.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "unsupportedFeatures", []interface{}{})
				for _, childItem := range item.UnsupportedFeatures {
					itemChildBody := ""
					if !childItem.ParcelType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelType", childItem.ParcelType.ValueString())
					}
					if !childItem.ParcelId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelId", childItem.ParcelId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "unsupportedFeatures.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "topology.devices.-1", itemBody)
		}
	}
	if !data.TopologySiteDevices.IsNull() {
		body, _ = sjson.Set(body, "topology.siteDevices", data.TopologySiteDevices.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ConfigurationGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("solution"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("profiles"); value.Exists() && len(value.Array()) > 0 {
		data.FeatureProfiles = make([]ConfigurationGroupFeatureProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupFeatureProfiles{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.FeatureProfiles = append(data.FeatureProfiles, item)
			return true
		})
	} else {
		if len(data.FeatureProfiles) > 0 {
			data.FeatureProfiles = []ConfigurationGroupFeatureProfiles{}
		}
	}
	if value := res.Get("topology.devices"); value.Exists() && len(value.Array()) > 0 {
		data.TopologyDevices = make([]ConfigurationGroupTopologyDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupTopologyDevices{}
			if cValue := v.Get("criteria.attribute"); cValue.Exists() {
				item.CriteriaAttribute = types.StringValue(cValue.String())
			} else {
				item.CriteriaAttribute = types.StringNull()
			}
			if cValue := v.Get("criteria.value"); cValue.Exists() {
				item.CriteriaValue = types.StringValue(cValue.String())
			} else {
				item.CriteriaValue = types.StringNull()
			}
			if cValue := v.Get("unsupportedFeatures"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.UnsupportedFeatures = make([]ConfigurationGroupTopologyDevicesUnsupportedFeatures, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
					if ccValue := cv.Get("parcelType"); ccValue.Exists() {
						cItem.ParcelType = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelType = types.StringNull()
					}
					if ccValue := cv.Get("parcelId"); ccValue.Exists() {
						cItem.ParcelId = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelId = types.StringNull()
					}
					item.UnsupportedFeatures = append(item.UnsupportedFeatures, cItem)
					return true
				})
			} else {
				if len(item.UnsupportedFeatures) > 0 {
					item.UnsupportedFeatures = []ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
				}
			}
			data.TopologyDevices = append(data.TopologyDevices, item)
			return true
		})
	} else {
		if len(data.TopologyDevices) > 0 {
			data.TopologyDevices = []ConfigurationGroupTopologyDevices{}
		}
	}
	if value := res.Get("topology.siteDevices"); value.Exists() {
		data.TopologySiteDevices = types.Int64Value(value.Int())
	} else {
		data.TopologySiteDevices = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ConfigurationGroup) hasChanges(ctx context.Context, state *ConfigurationGroup) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Solution.Equal(state.Solution) {
		hasChanges = true
	}
	if len(data.FeatureProfiles) != len(state.FeatureProfiles) {
		hasChanges = true
	} else {
		for i := range data.FeatureProfiles {
			if !data.FeatureProfiles[i].Id.Equal(state.FeatureProfiles[i].Id) {
				hasChanges = true
			}
		}
	}
	if len(data.TopologyDevices) != len(state.TopologyDevices) {
		hasChanges = true
	} else {
		for i := range data.TopologyDevices {
			if !data.TopologyDevices[i].CriteriaAttribute.Equal(state.TopologyDevices[i].CriteriaAttribute) {
				hasChanges = true
			}
			if !data.TopologyDevices[i].CriteriaValue.Equal(state.TopologyDevices[i].CriteriaValue) {
				hasChanges = true
			}
			if len(data.TopologyDevices[i].UnsupportedFeatures) != len(state.TopologyDevices[i].UnsupportedFeatures) {
				hasChanges = true
			} else {
				for ii := range data.TopologyDevices[i].UnsupportedFeatures {
					if !data.TopologyDevices[i].UnsupportedFeatures[ii].ParcelType.Equal(state.TopologyDevices[i].UnsupportedFeatures[ii].ParcelType) {
						hasChanges = true
					}
					if !data.TopologyDevices[i].UnsupportedFeatures[ii].ParcelId.Equal(state.TopologyDevices[i].UnsupportedFeatures[ii].ParcelId) {
						hasChanges = true
					}
				}
			}
		}
	}
	if !data.TopologySiteDevices.Equal(state.TopologySiteDevices) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
