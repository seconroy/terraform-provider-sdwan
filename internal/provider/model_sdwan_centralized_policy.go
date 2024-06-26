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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CentralizedPolicy struct {
	Id          types.String                   `tfsdk:"id"`
	Version     types.Int64                    `tfsdk:"version"`
	Name        types.String                   `tfsdk:"name"`
	Description types.String                   `tfsdk:"description"`
	Definitions []CentralizedPolicyDefinitions `tfsdk:"definitions"`
}

type CentralizedPolicyDefinitions struct {
	Id      types.String                          `tfsdk:"id"`
	Version types.Int64                           `tfsdk:"version"`
	Type    types.String                          `tfsdk:"type"`
	Entries []CentralizedPolicyDefinitionsEntries `tfsdk:"entries"`
}

type CentralizedPolicyDefinitionsEntries struct {
	SiteListIds        types.Set    `tfsdk:"site_list_ids"`
	SiteListVersions   types.List   `tfsdk:"site_list_versions"`
	VpnListIds         types.Set    `tfsdk:"vpn_list_ids"`
	VpnListVersions    types.List   `tfsdk:"vpn_list_versions"`
	Direction          types.String `tfsdk:"direction"`
	RegionListIds      types.Set    `tfsdk:"region_list_ids"`
	RegionListVersions types.List   `tfsdk:"region_list_versions"`
	RegionIds          types.Set    `tfsdk:"region_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CentralizedPolicy) getPath() string {
	return "/template/policy/vsmart/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CentralizedPolicy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "policyType", "feature")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "policyName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "policyDescription", data.Description.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.assembly", []interface{}{})
		for _, item := range data.Definitions {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "definitionId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "entries", []interface{}{})
				for _, childItem := range item.Entries {
					itemChildBody := ""
					if !childItem.SiteListIds.IsNull() {
						var values []string
						childItem.SiteListIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "siteLists", values)
					}
					if !childItem.VpnListIds.IsNull() {
						var values []string
						childItem.VpnListIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "vpnLists", values)
					}
					if !childItem.Direction.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "direction", childItem.Direction.ValueString())
					}
					if !childItem.RegionListIds.IsNull() {
						var values []string
						childItem.RegionListIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "regionLists", values)
					}
					if !childItem.RegionIds.IsNull() {
						var values []string
						childItem.RegionIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "regionIds", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "entries.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.assembly.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CentralizedPolicy) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("policyName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policyDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("policyDefinition.assembly"); value.Exists() && len(value.Array()) > 0 {
		data.Definitions = make([]CentralizedPolicyDefinitions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CentralizedPolicyDefinitions{}
			if cValue := v.Get("definitionId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Entries = make([]CentralizedPolicyDefinitionsEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CentralizedPolicyDefinitionsEntries{}
					if ccValue := cv.Get("siteLists"); ccValue.Exists() {
						cItem.SiteListIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.SiteListIds = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("vpnLists"); ccValue.Exists() {
						cItem.VpnListIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.VpnListIds = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("direction"); ccValue.Exists() {
						cItem.Direction = types.StringValue(ccValue.String())
					} else {
						cItem.Direction = types.StringNull()
					}
					if ccValue := cv.Get("regionLists"); ccValue.Exists() {
						cItem.RegionListIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.RegionListIds = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("regionIds"); ccValue.Exists() {
						cItem.RegionIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.RegionIds = types.SetNull(types.StringType)
					}
					item.Entries = append(item.Entries, cItem)
					return true
				})
			} else {
				if len(item.Entries) > 0 {
					item.Entries = []CentralizedPolicyDefinitionsEntries{}
				}
			}
			data.Definitions = append(data.Definitions, item)
			return true
		})
	} else {
		if len(data.Definitions) > 0 {
			data.Definitions = []CentralizedPolicyDefinitions{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CentralizedPolicy) hasChanges(ctx context.Context, state *CentralizedPolicy) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Definitions) != len(state.Definitions) {
		hasChanges = true
	} else {
		for i := range data.Definitions {
			if !data.Definitions[i].Id.Equal(state.Definitions[i].Id) {
				hasChanges = true
			}
			if !data.Definitions[i].Type.Equal(state.Definitions[i].Type) {
				hasChanges = true
			}
			if len(data.Definitions[i].Entries) != len(state.Definitions[i].Entries) {
				hasChanges = true
			} else {
				for ii := range data.Definitions[i].Entries {
					if !data.Definitions[i].Entries[ii].SiteListIds.Equal(state.Definitions[i].Entries[ii].SiteListIds) {
						hasChanges = true
					}
					if !data.Definitions[i].Entries[ii].VpnListIds.Equal(state.Definitions[i].Entries[ii].VpnListIds) {
						hasChanges = true
					}
					if !data.Definitions[i].Entries[ii].Direction.Equal(state.Definitions[i].Entries[ii].Direction) {
						hasChanges = true
					}
					if !data.Definitions[i].Entries[ii].RegionListIds.Equal(state.Definitions[i].Entries[ii].RegionListIds) {
						hasChanges = true
					}
					if !data.Definitions[i].Entries[ii].RegionIds.Equal(state.Definitions[i].Entries[ii].RegionIds) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *CentralizedPolicy) updateVersions(ctx context.Context, state *CentralizedPolicy) {
	for i := range data.Definitions {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Definitions[i].Id.ValueString())}
		stateIndex := -1
		for j := range state.Definitions {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Definitions[j].Id.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Definitions[i].Version = state.Definitions[stateIndex].Version
		} else {
			data.Definitions[i].Version = types.Int64Null()
		}
		for ii := range data.Definitions[i].Entries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Definitions[i].Entries[ii].SiteListIds.String()), fmt.Sprintf("%v", data.Definitions[i].Entries[ii].VpnListIds.String()), fmt.Sprintf("%v", data.Definitions[i].Entries[ii].RegionListIds.String()), fmt.Sprintf("%v", data.Definitions[i].Entries[ii].RegionIds.String())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Definitions[stateIndex].Entries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Definitions[stateIndex].Entries[jj].SiteListIds.String()), fmt.Sprintf("%v", state.Definitions[stateIndex].Entries[jj].VpnListIds.String()), fmt.Sprintf("%v", state.Definitions[stateIndex].Entries[jj].RegionListIds.String()), fmt.Sprintf("%v", state.Definitions[stateIndex].Entries[jj].RegionIds.String())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 && !state.Definitions[stateIndex].Entries[cStateIndex].SiteListVersions.IsNull() {
				data.Definitions[i].Entries[ii].SiteListVersions = state.Definitions[stateIndex].Entries[cStateIndex].SiteListVersions
			} else {
				data.Definitions[i].Entries[ii].SiteListVersions = types.ListNull(types.StringType)
			}
			if cStateIndex > -1 && !state.Definitions[stateIndex].Entries[cStateIndex].VpnListVersions.IsNull() {
				data.Definitions[i].Entries[ii].VpnListVersions = state.Definitions[stateIndex].Entries[cStateIndex].VpnListVersions
			} else {
				data.Definitions[i].Entries[ii].VpnListVersions = types.ListNull(types.StringType)
			}
			if cStateIndex > -1 && !state.Definitions[stateIndex].Entries[cStateIndex].RegionListVersions.IsNull() {
				data.Definitions[i].Entries[ii].RegionListVersions = state.Definitions[stateIndex].Entries[cStateIndex].RegionListVersions
			} else {
				data.Definitions[i].Entries[ii].RegionListVersions = types.ListNull(types.StringType)
			}
		}
	}
}

// End of section. //template:end updateVersions
