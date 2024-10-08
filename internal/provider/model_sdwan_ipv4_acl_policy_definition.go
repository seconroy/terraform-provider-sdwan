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
type IPv4ACLPolicyDefinition struct {
	Id            types.String                       `tfsdk:"id"`
	Version       types.Int64                        `tfsdk:"version"`
	Type          types.String                       `tfsdk:"type"`
	Name          types.String                       `tfsdk:"name"`
	Description   types.String                       `tfsdk:"description"`
	DefaultAction types.String                       `tfsdk:"default_action"`
	Sequences     []IPv4ACLPolicyDefinitionSequences `tfsdk:"sequences"`
}

type IPv4ACLPolicyDefinitionSequences struct {
	Id            types.Int64                                     `tfsdk:"id"`
	Name          types.String                                    `tfsdk:"name"`
	BaseAction    types.String                                    `tfsdk:"base_action"`
	MatchEntries  []IPv4ACLPolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []IPv4ACLPolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type IPv4ACLPolicyDefinitionSequencesMatchEntries struct {
	Type                                 types.String `tfsdk:"type"`
	Dscp                                 types.Int64  `tfsdk:"dscp"`
	SourceIp                             types.String `tfsdk:"source_ip"`
	IcmpMessage                          types.String `tfsdk:"icmp_message"`
	DestinationIp                        types.String `tfsdk:"destination_ip"`
	ClassMapId                           types.String `tfsdk:"class_map_id"`
	ClassMapVersion                      types.Int64  `tfsdk:"class_map_version"`
	PacketLength                         types.Int64  `tfsdk:"packet_length"`
	Priority                             types.String `tfsdk:"priority"`
	SourcePorts                          types.String `tfsdk:"source_ports"`
	DestinationPorts                     types.String `tfsdk:"destination_ports"`
	SourceDataIpv4PrefixListId           types.String `tfsdk:"source_data_ipv4_prefix_list_id"`
	SourceDataIpv4PrefixListVersion      types.Int64  `tfsdk:"source_data_ipv4_prefix_list_version"`
	DestinationDataIpv4PrefixListId      types.String `tfsdk:"destination_data_ipv4_prefix_list_id"`
	DestinationDataIpv4PrefixListVersion types.Int64  `tfsdk:"destination_data_ipv4_prefix_list_version"`
	Protocol                             types.String `tfsdk:"protocol"`
	Tcp                                  types.String `tfsdk:"tcp"`
}
type IPv4ACLPolicyDefinitionSequencesActionEntries struct {
	Type            types.String                                                 `tfsdk:"type"`
	ClassMapId      types.String                                                 `tfsdk:"class_map_id"`
	ClassMapVersion types.Int64                                                  `tfsdk:"class_map_version"`
	CounterName     types.String                                                 `tfsdk:"counter_name"`
	Log             types.Bool                                                   `tfsdk:"log"`
	MirrorId        types.String                                                 `tfsdk:"mirror_id"`
	MirrorVersion   types.Int64                                                  `tfsdk:"mirror_version"`
	PolicerId       types.String                                                 `tfsdk:"policer_id"`
	PolicerVersion  types.Int64                                                  `tfsdk:"policer_version"`
	SetParameters   []IPv4ACLPolicyDefinitionSequencesActionEntriesSetParameters `tfsdk:"set_parameters"`
}

type IPv4ACLPolicyDefinitionSequencesActionEntriesSetParameters struct {
	Type    types.String `tfsdk:"type"`
	Dscp    types.Int64  `tfsdk:"dscp"`
	NextHop types.String `tfsdk:"next_hop"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPv4ACLPolicyDefinition) getPath() string {
	return "/template/policy/definition/acl/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPv4ACLPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "acl")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", "acl")
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.Dscp.IsNull() && childItem.Type.ValueString() == "dscp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Dscp.ValueInt64()))
					}
					if !childItem.SourceIp.IsNull() && childItem.Type.ValueString() == "sourceIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourceIp.ValueString())
					}
					if !childItem.IcmpMessage.IsNull() && childItem.Type.ValueString() == "icmpMessage" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.IcmpMessage.ValueString())
					}
					if !childItem.DestinationIp.IsNull() && childItem.Type.ValueString() == "destinationIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationIp.ValueString())
					}
					if !childItem.ClassMapId.IsNull() && childItem.Type.ValueString() == "class" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ClassMapId.ValueString())
					}
					if !childItem.PacketLength.IsNull() && childItem.Type.ValueString() == "packetLength" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.PacketLength.ValueInt64()))
					}
					if !childItem.Priority.IsNull() && childItem.Type.ValueString() == "plp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Priority.ValueString())
					}
					if !childItem.SourcePorts.IsNull() && childItem.Type.ValueString() == "sourcePort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourcePorts.ValueString())
					}
					if !childItem.DestinationPorts.IsNull() && childItem.Type.ValueString() == "destinationPort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationPorts.ValueString())
					}
					if !childItem.SourceDataIpv4PrefixListId.IsNull() && childItem.Type.ValueString() == "sourceDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.SourceDataIpv4PrefixListId.ValueString())
					}
					if !childItem.DestinationDataIpv4PrefixListId.IsNull() && childItem.Type.ValueString() == "destinationDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DestinationDataIpv4PrefixListId.ValueString())
					}
					if !childItem.Protocol.IsNull() && childItem.Type.ValueString() == "protocol" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Protocol.ValueString()))
					}
					if !childItem.Tcp.IsNull() && childItem.Type.ValueString() == "tcp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Tcp.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.ClassMapId.IsNull() && childItem.Type.ValueString() == "class" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.ref", childItem.ClassMapId.ValueString())
					}
					if !childItem.CounterName.IsNull() && childItem.Type.ValueString() == "count" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.CounterName.ValueString())
					}
					if !childItem.Log.IsNull() && childItem.Type.ValueString() == "log" {
						if true && childItem.Log.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.MirrorId.IsNull() && childItem.Type.ValueString() == "mirror" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.ref", childItem.MirrorId.ValueString())
					}
					if !childItem.PolicerId.IsNull() && childItem.Type.ValueString() == "policer" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.ref", childItem.PolicerId.ValueString())
					}
					if true && childItem.Type.ValueString() == "set" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", []interface{}{})
						for _, childChildItem := range childItem.SetParameters {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "field", childChildItem.Type.ValueString())
							}
							if !childChildItem.Dscp.IsNull() && childChildItem.Type.ValueString() == "dscp" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.Dscp.ValueInt64()))
							}
							if !childChildItem.NextHop.IsNull() && childChildItem.Type.ValueString() == "nextHop" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.NextHop.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "parameter.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPv4ACLPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
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
	if value := res.Get("defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]IPv4ACLPolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPv4ACLPolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]IPv4ACLPolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := IPv4ACLPolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "dscp" {
						cItem.Dscp = types.Int64Value(ccValue.Int())
					} else {
						cItem.Dscp = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourceIp" {
						cItem.SourceIp = types.StringValue(ccValue.String())
					} else {
						cItem.SourceIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "icmpMessage" {
						cItem.IcmpMessage = types.StringValue(ccValue.String())
					} else {
						cItem.IcmpMessage = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationIp" {
						cItem.DestinationIp = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationIp = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "class" {
						cItem.ClassMapId = types.StringValue(ccValue.String())
					} else {
						cItem.ClassMapId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "packetLength" {
						cItem.PacketLength = types.Int64Value(ccValue.Int())
					} else {
						cItem.PacketLength = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "plp" {
						cItem.Priority = types.StringValue(ccValue.String())
					} else {
						cItem.Priority = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourcePort" {
						cItem.SourcePorts = types.StringValue(ccValue.String())
					} else {
						cItem.SourcePorts = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationPort" {
						cItem.DestinationPorts = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationPorts = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "sourceDataPrefixList" {
						cItem.SourceDataIpv4PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.SourceDataIpv4PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "destinationDataPrefixList" {
						cItem.DestinationDataIpv4PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationDataIpv4PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "protocol" {
						cItem.Protocol = types.StringValue(ccValue.String())
					} else {
						cItem.Protocol = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "tcp" {
						cItem.Tcp = types.StringValue(ccValue.String())
					} else {
						cItem.Tcp = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []IPv4ACLPolicyDefinitionSequencesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]IPv4ACLPolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := IPv4ACLPolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("parameter.ref"); ccValue.Exists() && cItem.Type.ValueString() == "class" {
						cItem.ClassMapId = types.StringValue(ccValue.String())
					} else {
						cItem.ClassMapId = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "count" {
						cItem.CounterName = types.StringValue(ccValue.String())
					} else {
						cItem.CounterName = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "log" {
						if true && ccValue.String() == "" {
							cItem.Log = types.BoolValue(true)
						}
					} else {
						cItem.Log = types.BoolNull()
					}
					if ccValue := cv.Get("parameter.ref"); ccValue.Exists() && cItem.Type.ValueString() == "mirror" {
						cItem.MirrorId = types.StringValue(ccValue.String())
					} else {
						cItem.MirrorId = types.StringNull()
					}
					if ccValue := cv.Get("parameter.ref"); ccValue.Exists() && cItem.Type.ValueString() == "policer" {
						cItem.PolicerId = types.StringValue(ccValue.String())
					} else {
						cItem.PolicerId = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && len(ccValue.Array()) > 0 && cItem.Type.ValueString() == "set" {
						cItem.SetParameters = make([]IPv4ACLPolicyDefinitionSequencesActionEntriesSetParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := IPv4ACLPolicyDefinitionSequencesActionEntriesSetParameters{}
							if cccValue := ccv.Get("field"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "dscp" {
								ccItem.Dscp = types.Int64Value(cccValue.Int())
							} else {
								ccItem.Dscp = types.Int64Null()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "nextHop" {
								ccItem.NextHop = types.StringValue(cccValue.String())
							} else {
								ccItem.NextHop = types.StringNull()
							}
							cItem.SetParameters = append(cItem.SetParameters, ccItem)
							return true
						})
					} else {
						if len(cItem.SetParameters) > 0 {
							cItem.SetParameters = []IPv4ACLPolicyDefinitionSequencesActionEntriesSetParameters{}
						}
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []IPv4ACLPolicyDefinitionSequencesActionEntries{}
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		if len(data.Sequences) > 0 {
			data.Sequences = []IPv4ACLPolicyDefinitionSequences{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *IPv4ACLPolicyDefinition) hasChanges(ctx context.Context, state *IPv4ACLPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].BaseAction.Equal(state.Sequences[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Dscp.Equal(state.Sequences[i].MatchEntries[ii].Dscp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceIp.Equal(state.Sequences[i].MatchEntries[ii].SourceIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].IcmpMessage.Equal(state.Sequences[i].MatchEntries[ii].IcmpMessage) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationIp.Equal(state.Sequences[i].MatchEntries[ii].DestinationIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ClassMapId.Equal(state.Sequences[i].MatchEntries[ii].ClassMapId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].PacketLength.Equal(state.Sequences[i].MatchEntries[ii].PacketLength) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Priority.Equal(state.Sequences[i].MatchEntries[ii].Priority) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourcePorts.Equal(state.Sequences[i].MatchEntries[ii].SourcePorts) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationPorts.Equal(state.Sequences[i].MatchEntries[ii].DestinationPorts) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Protocol.Equal(state.Sequences[i].MatchEntries[ii].Protocol) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Tcp.Equal(state.Sequences[i].MatchEntries[ii].Tcp) {
						hasChanges = true
					}
				}
			}
			if len(data.Sequences[i].ActionEntries) != len(state.Sequences[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].ActionEntries {
					if !data.Sequences[i].ActionEntries[ii].Type.Equal(state.Sequences[i].ActionEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].ClassMapId.Equal(state.Sequences[i].ActionEntries[ii].ClassMapId) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CounterName.Equal(state.Sequences[i].ActionEntries[ii].CounterName) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Log.Equal(state.Sequences[i].ActionEntries[ii].Log) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].MirrorId.Equal(state.Sequences[i].ActionEntries[ii].MirrorId) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].PolicerId.Equal(state.Sequences[i].ActionEntries[ii].PolicerId) {
						hasChanges = true
					}
					if len(data.Sequences[i].ActionEntries[ii].SetParameters) != len(state.Sequences[i].ActionEntries[ii].SetParameters) {
						hasChanges = true
					} else {
						for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Type.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Type) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Dscp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Dscp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHop.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHop) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *IPv4ACLPolicyDefinition) updateVersions(ctx context.Context, state *IPv4ACLPolicyDefinition) {
	for i := range data.Sequences {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].Id.ValueInt64()), fmt.Sprintf("%v", data.Sequences[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Sequences {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[j].Id.ValueInt64()), fmt.Sprintf("%v", state.Sequences[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Sequences[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ClassMapVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ClassMapVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ClassMapVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].SourceDataIpv4PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DestinationDataIpv4PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListVersion = types.Int64Null()
			}
		}
		for ii := range data.Sequences[i].ActionEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].ActionEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].ActionEntries[ii].ClassMapVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].ClassMapVersion
			} else {
				data.Sequences[i].ActionEntries[ii].ClassMapVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].ActionEntries[ii].MirrorVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].MirrorVersion
			} else {
				data.Sequences[i].ActionEntries[ii].MirrorVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].ActionEntries[ii].PolicerVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].PolicerVersion
			} else {
				data.Sequences[i].ActionEntries[ii].PolicerVersion = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end updateVersions
