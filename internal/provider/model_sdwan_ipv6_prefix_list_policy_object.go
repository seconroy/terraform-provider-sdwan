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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type IPv6PrefixListPolicyObject struct {
	Id      types.String                        `tfsdk:"id"`
	Version types.Int64                         `tfsdk:"version"`
	Name    types.String                        `tfsdk:"name"`
	Entries []IPv6PrefixListPolicyObjectEntries `tfsdk:"entries"`
}

type IPv6PrefixListPolicyObjectEntries struct {
	Prefix types.String `tfsdk:"prefix"`
	Le     types.Int64  `tfsdk:"le"`
	Ge     types.Int64  `tfsdk:"ge"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPv6PrefixListPolicyObject) getPath() string {
	return "/template/policy/list/ipv6prefix/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPv6PrefixListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "ipv6prefix")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6Prefix", item.Prefix.ValueString())
			}
			if !item.Le.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "le", fmt.Sprint(item.Le.ValueInt64()))
			}
			if !item.Ge.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ge", fmt.Sprint(item.Ge.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPv6PrefixListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]IPv6PrefixListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPv6PrefixListPolicyObjectEntries{}
			if cValue := v.Get("ipv6Prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			} else {
				item.Prefix = types.StringNull()
			}
			if cValue := v.Get("le"); cValue.Exists() {
				item.Le = types.Int64Value(cValue.Int())
			} else {
				item.Le = types.Int64Null()
			}
			if cValue := v.Get("ge"); cValue.Exists() {
				item.Ge = types.Int64Value(cValue.Int())
			} else {
				item.Ge = types.Int64Null()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		if len(data.Entries) > 0 {
			data.Entries = []IPv6PrefixListPolicyObjectEntries{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *IPv6PrefixListPolicyObject) hasChanges(ctx context.Context, state *IPv6PrefixListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].Prefix.Equal(state.Entries[i].Prefix) {
				hasChanges = true
			}
			if !data.Entries[i].Le.Equal(state.Entries[i].Le) {
				hasChanges = true
			}
			if !data.Entries[i].Ge.Equal(state.Entries[i].Ge) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
