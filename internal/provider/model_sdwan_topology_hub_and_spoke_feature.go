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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TopologyHubAndSpoke struct {
	Id               types.String                     `tfsdk:"id"`
	Version          types.Int64                      `tfsdk:"version"`
	Name             types.String                     `tfsdk:"name"`
	Description      types.String                     `tfsdk:"description"`
	FeatureProfileId types.String                     `tfsdk:"feature_profile_id"`
	Vpn              types.Set                        `tfsdk:"vpn"`
	HubSites         types.Set                        `tfsdk:"hub_sites"`
	SpokeGroups      []TopologyHubAndSpokeSpokeGroups `tfsdk:"spoke_groups"`
}

type TopologyHubAndSpokeSpokeGroups struct {
	Name           types.String                                   `tfsdk:"name"`
	Sites          types.Set                                      `tfsdk:"sites"`
	HubPreferences []TopologyHubAndSpokeSpokeGroupsHubPreferences `tfsdk:"hub_preferences"`
}

type TopologyHubAndSpokeSpokeGroupsHubPreferences struct {
	Sites      types.Set   `tfsdk:"sites"`
	Preference types.Int64 `tfsdk:"preference"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyHubAndSpoke) getModel() string {
	return "topology_hub_and_spoke"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyHubAndSpoke) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/hubspoke", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyHubAndSpoke) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.Vpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.Vpn.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.HubSites.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"selectedHubs.optionType", "global")
			var values []string
			data.HubSites.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"selectedHubs.value", values)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"spokes", []interface{}{})
		for _, item := range data.SpokeGroups {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if !item.Sites.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "spokeSites.optionType", "global")
					var values []string
					item.Sites.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "spokeSites.value", values)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "hubSites", []interface{}{})
				for _, childItem := range item.HubPreferences {
					itemChildBody := ""
					if !childItem.Sites.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sites.optionType", "global")
							var values []string
							childItem.Sites.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sites.value", values)
						}
					}
					if !childItem.Preference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.value", childItem.Preference.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "hubSites.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"spokes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyHubAndSpoke) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpn = helpers.GetStringSet(va.Array())
		}
	}
	data.HubSites = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHubs.value")
		if t.String() == "global" {
			data.HubSites = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "spokes"); value.Exists() && len(value.Array()) > 0 {
		data.SpokeGroups = make([]TopologyHubAndSpokeSpokeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyHubAndSpokeSpokeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.Sites = types.SetNull(types.StringType)

			if t := v.Get("spokeSites.optionType"); t.Exists() {
				va := v.Get("spokeSites.value")
				if t.String() == "global" {
					item.Sites = helpers.GetStringSet(va.Array())
				}
			}
			if cValue := v.Get("hubSites"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.HubPreferences = make([]TopologyHubAndSpokeSpokeGroupsHubPreferences, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyHubAndSpokeSpokeGroupsHubPreferences{}
					cItem.Sites = types.SetNull(types.StringType)

					if t := cv.Get("sites.optionType"); t.Exists() {
						va := cv.Get("sites.value")
						if t.String() == "global" {
							cItem.Sites = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Preference = types.Int64Null()

					if t := cv.Get("preference.optionType"); t.Exists() {
						va := cv.Get("preference.value")
						if t.String() == "global" {
							cItem.Preference = types.Int64Value(va.Int())
						}
					}
					item.HubPreferences = append(item.HubPreferences, cItem)
					return true
				})
			}
			data.SpokeGroups = append(data.SpokeGroups, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TopologyHubAndSpoke) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpn = helpers.GetStringSet(va.Array())
		}
	}
	data.HubSites = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHubs.value")
		if t.String() == "global" {
			data.HubSites = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.SpokeGroups {
		keys := [...]string{"name", "spokeSites"}
		keyValues := [...]string{data.SpokeGroups[i].Name.ValueString(), helpers.GetStringFromSet(data.SpokeGroups[i].Sites).ValueString()}
		keyValuesVariables := [...]string{"", ""}

		var r gjson.Result
		res.Get(path + "spokes").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.SpokeGroups[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.SpokeGroups[i].Name = types.StringValue(va.String())
			}
		}
		data.SpokeGroups[i].Sites = types.SetNull(types.StringType)

		if t := r.Get("spokeSites.optionType"); t.Exists() {
			va := r.Get("spokeSites.value")
			if t.String() == "global" {
				data.SpokeGroups[i].Sites = helpers.GetStringSet(va.Array())
			}
		}
		for ci := range data.SpokeGroups[i].HubPreferences {
			keys := [...]string{"sites", "preference"}
			keyValues := [...]string{helpers.GetStringFromSet(data.SpokeGroups[i].HubPreferences[ci].Sites).ValueString(), strconv.FormatInt(data.SpokeGroups[i].HubPreferences[ci].Preference.ValueInt64(), 10)}
			keyValuesVariables := [...]string{"", ""}

			var cr gjson.Result
			r.Get("hubSites").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.SpokeGroups[i].HubPreferences[ci].Sites = types.SetNull(types.StringType)

			if t := cr.Get("sites.optionType"); t.Exists() {
				va := cr.Get("sites.value")
				if t.String() == "global" {
					data.SpokeGroups[i].HubPreferences[ci].Sites = helpers.GetStringSet(va.Array())
				}
			}
			data.SpokeGroups[i].HubPreferences[ci].Preference = types.Int64Null()

			if t := cr.Get("preference.optionType"); t.Exists() {
				va := cr.Get("preference.value")
				if t.String() == "global" {
					data.SpokeGroups[i].HubPreferences[ci].Preference = types.Int64Value(va.Int())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody
