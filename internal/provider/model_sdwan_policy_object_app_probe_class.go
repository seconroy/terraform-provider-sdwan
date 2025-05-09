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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectAppProbeClass struct {
	Id               types.String                       `tfsdk:"id"`
	Version          types.Int64                        `tfsdk:"version"`
	Name             types.String                       `tfsdk:"name"`
	Description      types.String                       `tfsdk:"description"`
	FeatureProfileId types.String                       `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectAppProbeClassEntries `tfsdk:"entries"`
}

type PolicyObjectAppProbeClassEntries struct {
	Map             []PolicyObjectAppProbeClassEntriesMap `tfsdk:"map"`
	ForwardingClass types.String                          `tfsdk:"forwarding_class"`
}

type PolicyObjectAppProbeClassEntriesMap struct {
	Color types.String `tfsdk:"color"`
	Dscp  types.Int64  `tfsdk:"dscp"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectAppProbeClass) getModel() string {
	return "policy_object_app_probe_class"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectAppProbeClass) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/app-probe", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectAppProbeClass) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "map", []interface{}{})
				for _, childItem := range item.Map {
					itemChildBody := ""
					if !childItem.Color.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "color.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "color.value", childItem.Color.ValueString())
						}
					}
					if !childItem.Dscp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.value", childItem.Dscp.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "map.-1", itemChildBody)
				}
			}
			if !item.ForwardingClass.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "forwardingClass.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "forwardingClass.value", item.ForwardingClass.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectAppProbeClass) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectAppProbeClassEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectAppProbeClassEntries{}
			if cValue := v.Get("map"); cValue.Exists() {
				item.Map = make([]PolicyObjectAppProbeClassEntriesMap, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyObjectAppProbeClassEntriesMap{}
					cItem.Color = types.StringNull()

					if t := cv.Get("color.optionType"); t.Exists() {
						va := cv.Get("color.value")
						if t.String() == "global" {
							cItem.Color = types.StringValue(va.String())
						}
					}
					cItem.Dscp = types.Int64Null()

					if t := cv.Get("dscp.optionType"); t.Exists() {
						va := cv.Get("dscp.value")
						if t.String() == "global" {
							cItem.Dscp = types.Int64Value(va.Int())
						}
					}
					item.Map = append(item.Map, cItem)
					return true
				})
			}
			item.ForwardingClass = types.StringNull()

			if t := v.Get("forwardingClass.optionType"); t.Exists() {
				va := v.Get("forwardingClass.value")
				if t.String() == "global" {
					item.ForwardingClass = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectAppProbeClass) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"forwardingClass"}
		keyValues := [...]string{data.Entries[i].ForwardingClass.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "entries").ForEach(
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
		for ci := range data.Entries[i].Map {
			keys := [...]string{"color", "dscp"}
			keyValues := [...]string{data.Entries[i].Map[ci].Color.ValueString(), strconv.FormatInt(data.Entries[i].Map[ci].Dscp.ValueInt64(), 10)}
			keyValuesVariables := [...]string{"", ""}

			var cr gjson.Result
			r.Get("map").ForEach(
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
			data.Entries[i].Map[ci].Color = types.StringNull()

			if t := cr.Get("color.optionType"); t.Exists() {
				va := cr.Get("color.value")
				if t.String() == "global" {
					data.Entries[i].Map[ci].Color = types.StringValue(va.String())
				}
			}
			data.Entries[i].Map[ci].Dscp = types.Int64Null()

			if t := cr.Get("dscp.optionType"); t.Exists() {
				va := cr.Get("dscp.value")
				if t.String() == "global" {
					data.Entries[i].Map[ci].Dscp = types.Int64Value(va.Int())
				}
			}
		}
		data.Entries[i].ForwardingClass = types.StringNull()

		if t := r.Get("forwardingClass.optionType"); t.Exists() {
			va := r.Get("forwardingClass.value")
			if t.String() == "global" {
				data.Entries[i].ForwardingClass = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
