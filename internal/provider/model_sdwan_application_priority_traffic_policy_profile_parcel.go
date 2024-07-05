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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ApplicationPriorityTrafficPolicy struct {
	Id               types.String                                `tfsdk:"id"`
	Version          types.Int64                                 `tfsdk:"version"`
	Name             types.String                                `tfsdk:"name"`
	Description      types.String                                `tfsdk:"description"`
	FeatureProfileId types.String                                `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                                `tfsdk:"default_action"`
	Vpn              types.Set                                   `tfsdk:"vpn"`
	TargetDirection  types.String                                `tfsdk:"target_direction"`
	Sequences        []ApplicationPriorityTrafficPolicySequences `tfsdk:"sequences"`
}

type ApplicationPriorityTrafficPolicySequences struct {
	SequenceId types.Int64                                        `tfsdk:"sequence_id"`
	Name       types.String                                       `tfsdk:"name"`
	Protocol   types.String                                       `tfsdk:"protocol"`
	BaseAction types.String                                       `tfsdk:"base_action"`
	Matches    []ApplicationPriorityTrafficPolicySequencesMatches `tfsdk:"matches"`
	Actions    []ApplicationPriorityTrafficPolicySequencesActions `tfsdk:"actions"`
}

type ApplicationPriorityTrafficPolicySequencesMatches struct {
	ApplicationListId           types.String `tfsdk:"application_list_id"`
	DnsApplicationListId        types.String `tfsdk:"dns_application_list_id"`
	Dns                         types.String `tfsdk:"dns"`
	Dscp                        types.Int64  `tfsdk:"dscp"`
	PacketLength                types.String `tfsdk:"packet_length"`
	Protocol                    types.Set    `tfsdk:"protocol"`
	SourceDataPrefixListId      types.String `tfsdk:"source_data_prefix_list_id"`
	SourcePort                  types.Set    `tfsdk:"source_port"`
	DestinationDataPrefixListId types.String `tfsdk:"destination_data_prefix_list_id"`
	DestinationPort             types.Set    `tfsdk:"destination_port"`
	DestinationRegion           types.String `tfsdk:"destination_region"`
	Tcp                         types.String `tfsdk:"tcp"`
	TrafficTo                   types.String `tfsdk:"traffic_to"`
}
type ApplicationPriorityTrafficPolicySequencesActions struct {
	Counter               types.String                                               `tfsdk:"counter"`
	Log                   types.Bool                                                 `tfsdk:"log"`
	SlaClass              []ApplicationPriorityTrafficPolicySequencesActionsSlaClass `tfsdk:"sla_class"`
	Sets                  []ApplicationPriorityTrafficPolicySequencesActionsSets     `tfsdk:"sets"`
	RedirectDnsField      types.String                                               `tfsdk:"redirect_dns_field"`
	RedirectDnsValue      types.String                                               `tfsdk:"redirect_dns_value"`
	NatPool               types.Int64                                                `tfsdk:"nat_pool"`
	NatVpn                types.Bool                                                 `tfsdk:"nat_vpn"`
	NatFallback           types.Bool                                                 `tfsdk:"nat_fallback"`
	SecureInternetGateway types.Bool                                                 `tfsdk:"secure_internet_gateway"`
	FallbackToRouting     types.Bool                                                 `tfsdk:"fallback_to_routing"`
}

type ApplicationPriorityTrafficPolicySequencesActionsSlaClass struct {
	SlaClassListId types.String `tfsdk:"sla_class_list_id"`
	PreferredColor types.Set    `tfsdk:"preferred_color"`
}
type ApplicationPriorityTrafficPolicySequencesActionsSets struct {
	Dscp                       types.Int64  `tfsdk:"dscp"`
	PolicerId                  types.String `tfsdk:"policer_id"`
	PreferredColorGroupId      types.String `tfsdk:"preferred_color_group_id"`
	ForwardingClassId          types.String `tfsdk:"forwarding_class_id"`
	LocalTlocListColor         types.Set    `tfsdk:"local_tloc_list_color"`
	LocalTlocRestrict          types.String `tfsdk:"local_tloc_restrict"`
	LocalTlocListEncapsulation types.String `tfsdk:"local_tloc_list_encapsulation"`
	TlocIp                     types.String `tfsdk:"tloc_ip"`
	TlocColor                  types.Set    `tfsdk:"tloc_color"`
	TlocEncapsulation          types.String `tfsdk:"tloc_encapsulation"`
	TlocListId                 types.String `tfsdk:"tloc_list_id"`
	ServiceType                types.String `tfsdk:"service_type"`
	ServiceColor               types.Set    `tfsdk:"service_color"`
	ServiceEncapsulation       types.String `tfsdk:"service_encapsulation"`
	ServiceTlocIp              types.String `tfsdk:"service_tloc_ip"`
	ServiceVpn                 types.String `tfsdk:"service_vpn"`
	NextHop                    types.String `tfsdk:"next_hop"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ApplicationPriorityTrafficPolicy) getModel() string {
	return "application_priority_traffic_policy"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPriorityTrafficPolicy) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/application-priority/%v/traffic-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ApplicationPriorityTrafficPolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, path+"dataDefaultAction.optionType", "global")
		body, _ = sjson.Set(body, path+"dataDefaultAction.value", data.DefaultAction.ValueString())
	}
	if !data.Vpn.IsNull() {
		body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
		var values []string
		data.Vpn.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"target.vpn.value", values)
	}
	if !data.TargetDirection.IsNull() {
		body, _ = sjson.Set(body, path+"target.direction.optionType", "global")
		body, _ = sjson.Set(body, path+"target.direction.value", data.TargetDirection.ValueString())
	}
	body, _ = sjson.Set(body, path+"sequences", []interface{}{})
	for _, item := range data.Sequences {
		itemBody := ""
		if !item.SequenceId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceId.ValueInt64())
		}
		if !item.Name.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
		}
		if !item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sequenceIpType.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sequenceIpType.value", item.Protocol.ValueString())
		}
		if !item.BaseAction.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
		}
		itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
		for _, childItem := range item.Matches {
			itemChildBody := ""
			if !childItem.ApplicationListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.value", childItem.ApplicationListId.ValueString())
			}
			if !childItem.DnsApplicationListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dnsAppList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "dnsAppList.refId.value", childItem.DnsApplicationListId.ValueString())
			}
			if !childItem.Dns.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dns.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "dns.value", childItem.Dns.ValueString())
			}
			if !childItem.Dscp.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dscp.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "dscp.value", childItem.Dscp.ValueInt64())
			}
			if !childItem.PacketLength.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.value", childItem.PacketLength.ValueString())
			}
			if !childItem.Protocol.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
				var values []string
				childItem.Protocol.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", values)
			}
			if !childItem.SourceDataPrefixListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.value", childItem.SourceDataPrefixListId.ValueString())
			}
			if !childItem.SourcePort.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.optionType", "global")
				var values []string
				childItem.SourcePort.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.value", values)
			}
			if !childItem.DestinationDataPrefixListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.value", childItem.DestinationDataPrefixListId.ValueString())
			}
			if !childItem.DestinationPort.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.optionType", "global")
				var values []string
				childItem.DestinationPort.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.value", values)
			}
			if !childItem.DestinationRegion.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationRegion.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "destinationRegion.value", childItem.DestinationRegion.ValueString())
			}
			if !childItem.Tcp.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "tcp.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "tcp.value", childItem.Tcp.ValueString())
			}
			if !childItem.TrafficTo.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "trafficTo.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "trafficTo.value", childItem.TrafficTo.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
		}
		itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
		for _, childItem := range item.Actions {
			itemChildBody := ""
			if !childItem.Counter.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "count.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "count.value", childItem.Counter.ValueString())
			}
			if !childItem.Log.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "log.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "log.value", childItem.Log.ValueBool())
			}
			itemChildBody, _ = sjson.Set(itemChildBody, "slaClass", []interface{}{})
			for _, childChildItem := range childItem.SlaClass {
				itemChildChildBody := ""
				if !childChildItem.SlaClassListId.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "slaName.refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "slaName.refId.value", childChildItem.SlaClassListId.ValueString())
				}
				if !childChildItem.PreferredColor.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColor.optionType", "global")
					var values []string
					childChildItem.PreferredColor.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColor.value", values)
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "slaClass.-1", itemChildChildBody)
			}
			itemChildBody, _ = sjson.Set(itemChildBody, "set", []interface{}{})
			for _, childChildItem := range childItem.Sets {
				itemChildChildBody := ""
				if !childChildItem.Dscp.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "dscp.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "dscp.value", childChildItem.Dscp.ValueInt64())
				}
				if !childChildItem.PolicerId.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "policer.refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "policer.refId.value", childChildItem.PolicerId.ValueString())
				}
				if !childChildItem.PreferredColorGroupId.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preferredColorGroup.refId.value", childChildItem.PreferredColorGroupId.ValueString())
				}
				if !childChildItem.ForwardingClassId.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "forwardingClass.refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "forwardingClass.refId.value", childChildItem.ForwardingClassId.ValueString())
				}
				if !childChildItem.LocalTlocListColor.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.color.optionType", "global")
					var values []string
					childChildItem.LocalTlocListColor.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.color.value", values)
				}
				if !childChildItem.LocalTlocRestrict.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.restrict.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.restrict.value", childChildItem.LocalTlocRestrict.ValueString())
				}
				if !childChildItem.LocalTlocListEncapsulation.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.encap.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "localTlocList.encap.value", childChildItem.LocalTlocListEncapsulation.ValueString())
				}
				if !childChildItem.TlocIp.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.ip.value", childChildItem.TlocIp.ValueString())
				}
				if !childChildItem.TlocColor.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.optionType", "global")
					var values []string
					childChildItem.TlocColor.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.color.value", values)
				}
				if !childChildItem.TlocEncapsulation.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tloc.encap.value", childChildItem.TlocEncapsulation.ValueString())
				}
				if !childChildItem.TlocListId.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.value", childChildItem.TlocListId.ValueString())
				}
				if !childChildItem.ServiceType.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.type.value", childChildItem.ServiceType.ValueString())
				}
				if !childChildItem.ServiceColor.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.optionType", "global")
					var values []string
					childChildItem.ServiceColor.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.color.value", values)
				}
				if !childChildItem.ServiceEncapsulation.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.encap.value", childChildItem.ServiceEncapsulation.ValueString())
				}
				if !childChildItem.ServiceTlocIp.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.tloc.ip.value", childChildItem.ServiceTlocIp.ValueString())
				}
				if !childChildItem.ServiceVpn.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "service.vpn.value", childChildItem.ServiceVpn.ValueString())
				}
				if !childChildItem.NextHop.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHop.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "nextHop.value", childChildItem.NextHop.ValueString())
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "set.-1", itemChildChildBody)
			}
			if !childItem.RedirectDnsField.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.field.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.field.value", childItem.RedirectDnsField.ValueString())
			}
			if !childItem.RedirectDnsValue.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.value.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "redirectDns.value.value", childItem.RedirectDnsValue.ValueString())
			}
			if !childItem.NatPool.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "natPool.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "natPool.value", childItem.NatPool.ValueInt64())
			}
			if !childItem.NatVpn.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "nat.useVpn.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "nat.useVpn.value", childItem.NatVpn.ValueBool())
			}
			if !childItem.NatFallback.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "nat.fallback.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "nat.fallback.value", childItem.NatFallback.ValueBool())
			}
			if !childItem.SecureInternetGateway.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "sig.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "sig.value", childItem.SecureInternetGateway.ValueBool())
			}
			if !childItem.FallbackToRouting.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "fallbackToRouting.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "fallbackToRouting.value", childItem.FallbackToRouting.ValueBool())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPriorityTrafficPolicy) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "dataDefaultAction.optionType"); t.Exists() {
		va := res.Get(path + "dataDefaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.Vpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpn = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetDirection = types.StringNull()

	if t := res.Get(path + "target.direction.optionType"); t.Exists() {
		va := res.Get(path + "target.direction.value")
		if t.String() == "global" {
			data.TargetDirection = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "sequences"); value.Exists() {
		data.Sequences = make([]ApplicationPriorityTrafficPolicySequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationPriorityTrafficPolicySequences{}
			item.SequenceId = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.SequenceId = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.Protocol = types.StringNull()

			if t := v.Get("sequenceIpType.optionType"); t.Exists() {
				va := v.Get("sequenceIpType.value")
				if t.String() == "global" {
					item.Protocol = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("match.entries"); cValue.Exists() {
				item.Matches = make([]ApplicationPriorityTrafficPolicySequencesMatches, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPriorityTrafficPolicySequencesMatches{}
					cItem.ApplicationListId = types.StringNull()

					if t := cv.Get("appList.refId.optionType"); t.Exists() {
						va := cv.Get("appList.refId.value")
						if t.String() == "global" {
							cItem.ApplicationListId = types.StringValue(va.String())
						}
					}
					cItem.DnsApplicationListId = types.StringNull()

					if t := cv.Get("dnsAppList.refId.optionType"); t.Exists() {
						va := cv.Get("dnsAppList.refId.value")
						if t.String() == "global" {
							cItem.DnsApplicationListId = types.StringValue(va.String())
						}
					}
					cItem.Dns = types.StringNull()

					if t := cv.Get("dns.optionType"); t.Exists() {
						va := cv.Get("dns.value")
						if t.String() == "global" {
							cItem.Dns = types.StringValue(va.String())
						}
					}
					cItem.Dscp = types.Int64Null()

					if t := cv.Get("dscp.optionType"); t.Exists() {
						va := cv.Get("dscp.value")
						if t.String() == "global" {
							cItem.Dscp = types.Int64Value(va.Int())
						}
					}
					cItem.PacketLength = types.StringNull()

					if t := cv.Get("packetLength.optionType"); t.Exists() {
						va := cv.Get("packetLength.value")
						if t.String() == "global" {
							cItem.PacketLength = types.StringValue(va.String())
						}
					}
					cItem.Protocol = types.SetNull(types.StringType)

					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "global" {
							cItem.Protocol = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceDataPrefixListId = types.StringNull()

					if t := cv.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.SourcePort = types.SetNull(types.StringType)

					if t := cv.Get("sourcePort.optionType"); t.Exists() {
						va := cv.Get("sourcePort.value")
						if t.String() == "global" {
							cItem.SourcePort = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationDataPrefixListId = types.StringNull()

					if t := cv.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.DestinationDataPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.DestinationPort = types.SetNull(types.StringType)

					if t := cv.Get("destinationPort.optionType"); t.Exists() {
						va := cv.Get("destinationPort.value")
						if t.String() == "global" {
							cItem.DestinationPort = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationRegion = types.StringNull()

					if t := cv.Get("destinationRegion.optionType"); t.Exists() {
						va := cv.Get("destinationRegion.value")
						if t.String() == "global" {
							cItem.DestinationRegion = types.StringValue(va.String())
						}
					}
					cItem.Tcp = types.StringNull()

					if t := cv.Get("tcp.optionType"); t.Exists() {
						va := cv.Get("tcp.value")
						if t.String() == "global" {
							cItem.Tcp = types.StringValue(va.String())
						}
					}
					cItem.TrafficTo = types.StringNull()

					if t := cv.Get("trafficTo.optionType"); t.Exists() {
						va := cv.Get("trafficTo.value")
						if t.String() == "global" {
							cItem.TrafficTo = types.StringValue(va.String())
						}
					}
					item.Matches = append(item.Matches, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() {
				item.Actions = make([]ApplicationPriorityTrafficPolicySequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPriorityTrafficPolicySequencesActions{}
					cItem.Counter = types.StringNull()

					if t := cv.Get("count.optionType"); t.Exists() {
						va := cv.Get("count.value")
						if t.String() == "global" {
							cItem.Counter = types.StringValue(va.String())
						}
					}
					cItem.Log = types.BoolNull()

					if t := cv.Get("log.optionType"); t.Exists() {
						va := cv.Get("log.value")
						if t.String() == "global" {
							cItem.Log = types.BoolValue(va.Bool())
						}
					}
					if ccValue := cv.Get("slaClass"); ccValue.Exists() {
						cItem.SlaClass = make([]ApplicationPriorityTrafficPolicySequencesActionsSlaClass, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationPriorityTrafficPolicySequencesActionsSlaClass{}
							ccItem.SlaClassListId = types.StringNull()

							if t := ccv.Get("slaName.refId.optionType"); t.Exists() {
								va := ccv.Get("slaName.refId.value")
								if t.String() == "global" {
									ccItem.SlaClassListId = types.StringValue(va.String())
								}
							}
							ccItem.PreferredColor = types.SetNull(types.StringType)

							if t := ccv.Get("preferredColor.optionType"); t.Exists() {
								va := ccv.Get("preferredColor.value")
								if t.String() == "global" {
									ccItem.PreferredColor = helpers.GetStringSet(va.Array())
								}
							}
							cItem.SlaClass = append(cItem.SlaClass, ccItem)
							return true
						})
					}
					if ccValue := cv.Get("set"); ccValue.Exists() {
						cItem.Sets = make([]ApplicationPriorityTrafficPolicySequencesActionsSets, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationPriorityTrafficPolicySequencesActionsSets{}
							ccItem.Dscp = types.Int64Null()

							if t := ccv.Get("dscp.optionType"); t.Exists() {
								va := ccv.Get("dscp.value")
								if t.String() == "global" {
									ccItem.Dscp = types.Int64Value(va.Int())
								}
							}
							ccItem.PolicerId = types.StringNull()

							if t := ccv.Get("policer.refId.optionType"); t.Exists() {
								va := ccv.Get("policer.refId.value")
								if t.String() == "global" {
									ccItem.PolicerId = types.StringValue(va.String())
								}
							}
							ccItem.PreferredColorGroupId = types.StringNull()

							if t := ccv.Get("preferredColorGroup.refId.optionType"); t.Exists() {
								va := ccv.Get("preferredColorGroup.refId.value")
								if t.String() == "global" {
									ccItem.PreferredColorGroupId = types.StringValue(va.String())
								}
							}
							ccItem.ForwardingClassId = types.StringNull()

							if t := ccv.Get("forwardingClass.refId.optionType"); t.Exists() {
								va := ccv.Get("forwardingClass.refId.value")
								if t.String() == "global" {
									ccItem.ForwardingClassId = types.StringValue(va.String())
								}
							}
							ccItem.LocalTlocListColor = types.SetNull(types.StringType)

							if t := ccv.Get("localTlocList.color.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.color.value")
								if t.String() == "global" {
									ccItem.LocalTlocListColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.LocalTlocRestrict = types.StringNull()

							if t := ccv.Get("localTlocList.restrict.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.restrict.value")
								if t.String() == "global" {
									ccItem.LocalTlocRestrict = types.StringValue(va.String())
								}
							}
							ccItem.LocalTlocListEncapsulation = types.StringNull()

							if t := ccv.Get("localTlocList.encap.optionType"); t.Exists() {
								va := ccv.Get("localTlocList.encap.value")
								if t.String() == "global" {
									ccItem.LocalTlocListEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.TlocIp = types.StringNull()

							if t := ccv.Get("tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("tloc.ip.value")
								if t.String() == "global" {
									ccItem.TlocIp = types.StringValue(va.String())
								}
							}
							ccItem.TlocColor = types.SetNull(types.StringType)

							if t := ccv.Get("tloc.color.optionType"); t.Exists() {
								va := ccv.Get("tloc.color.value")
								if t.String() == "global" {
									ccItem.TlocColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.TlocEncapsulation = types.StringNull()

							if t := ccv.Get("tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("tloc.encap.value")
								if t.String() == "global" {
									ccItem.TlocEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.TlocListId = types.StringNull()

							if t := ccv.Get("tlocList.refId.optionType"); t.Exists() {
								va := ccv.Get("tlocList.refId.value")
								if t.String() == "global" {
									ccItem.TlocListId = types.StringValue(va.String())
								}
							}
							ccItem.ServiceType = types.StringNull()

							if t := ccv.Get("service.type.optionType"); t.Exists() {
								va := ccv.Get("service.type.value")
								if t.String() == "global" {
									ccItem.ServiceType = types.StringValue(va.String())
								}
							}
							ccItem.ServiceColor = types.SetNull(types.StringType)

							if t := ccv.Get("service.tloc.color.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.color.value")
								if t.String() == "global" {
									ccItem.ServiceColor = helpers.GetStringSet(va.Array())
								}
							}
							ccItem.ServiceEncapsulation = types.StringNull()

							if t := ccv.Get("service.tloc.encap.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.encap.value")
								if t.String() == "global" {
									ccItem.ServiceEncapsulation = types.StringValue(va.String())
								}
							}
							ccItem.ServiceTlocIp = types.StringNull()

							if t := ccv.Get("service.tloc.ip.optionType"); t.Exists() {
								va := ccv.Get("service.tloc.ip.value")
								if t.String() == "global" {
									ccItem.ServiceTlocIp = types.StringValue(va.String())
								}
							}
							ccItem.ServiceVpn = types.StringNull()

							if t := ccv.Get("service.vpn.optionType"); t.Exists() {
								va := ccv.Get("service.vpn.value")
								if t.String() == "global" {
									ccItem.ServiceVpn = types.StringValue(va.String())
								}
							}
							ccItem.NextHop = types.StringNull()

							if t := ccv.Get("nextHop.optionType"); t.Exists() {
								va := ccv.Get("nextHop.value")
								if t.String() == "global" {
									ccItem.NextHop = types.StringValue(va.String())
								}
							}
							cItem.Sets = append(cItem.Sets, ccItem)
							return true
						})
					}
					cItem.RedirectDnsField = types.StringNull()

					if t := cv.Get("redirectDns.field.optionType"); t.Exists() {
						va := cv.Get("redirectDns.field.value")
						if t.String() == "global" {
							cItem.RedirectDnsField = types.StringValue(va.String())
						}
					}
					cItem.RedirectDnsValue = types.StringNull()

					if t := cv.Get("redirectDns.value.optionType"); t.Exists() {
						va := cv.Get("redirectDns.value.value")
						if t.String() == "global" {
							cItem.RedirectDnsValue = types.StringValue(va.String())
						}
					}
					cItem.NatPool = types.Int64Null()

					if t := cv.Get("natPool.optionType"); t.Exists() {
						va := cv.Get("natPool.value")
						if t.String() == "global" {
							cItem.NatPool = types.Int64Value(va.Int())
						}
					}
					cItem.NatVpn = types.BoolNull()

					if t := cv.Get("nat.useVpn.optionType"); t.Exists() {
						va := cv.Get("nat.useVpn.value")
						if t.String() == "global" {
							cItem.NatVpn = types.BoolValue(va.Bool())
						}
					}
					cItem.NatFallback = types.BoolNull()

					if t := cv.Get("nat.fallback.optionType"); t.Exists() {
						va := cv.Get("nat.fallback.value")
						if t.String() == "global" {
							cItem.NatFallback = types.BoolValue(va.Bool())
						}
					}
					cItem.SecureInternetGateway = types.BoolNull()

					if t := cv.Get("sig.optionType"); t.Exists() {
						va := cv.Get("sig.value")
						if t.String() == "global" {
							cItem.SecureInternetGateway = types.BoolValue(va.Bool())
						}
					}
					cItem.FallbackToRouting = types.BoolNull()

					if t := cv.Get("fallbackToRouting.optionType"); t.Exists() {
						va := cv.Get("fallbackToRouting.value")
						if t.String() == "global" {
							cItem.FallbackToRouting = types.BoolValue(va.Bool())
						}
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ApplicationPriorityTrafficPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "dataDefaultAction.optionType"); t.Exists() {
		va := res.Get(path + "dataDefaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.Vpn = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpn = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetDirection = types.StringNull()

	if t := res.Get(path + "target.direction.optionType"); t.Exists() {
		va := res.Get(path + "target.direction.value")
		if t.String() == "global" {
			data.TargetDirection = types.StringValue(va.String())
		}
	}
	for i := range data.Sequences {
		keys := [...]string{}
		keyValues := [...]string{}
		keyValuesVariables := [...]string{}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Sequences[i].SequenceId = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].SequenceId = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].Name = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].Name = types.StringValue(va.String())
			}
		}
		data.Sequences[i].Protocol = types.StringNull()

		if t := r.Get("sequenceIpType.optionType"); t.Exists() {
			va := r.Get("sequenceIpType.value")
			if t.String() == "global" {
				data.Sequences[i].Protocol = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		for ci := range data.Sequences[i].Matches {
			keys := [...]string{}
			keyValues := [...]string{}
			keyValuesVariables := [...]string{}

			var cr gjson.Result
			r.Get("match.entries").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Sequences[i].Matches[ci].ApplicationListId = types.StringNull()

			if t := cr.Get("appList.refId.optionType"); t.Exists() {
				va := cr.Get("appList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].ApplicationListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].DnsApplicationListId = types.StringNull()

			if t := cr.Get("dnsAppList.refId.optionType"); t.Exists() {
				va := cr.Get("dnsAppList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].DnsApplicationListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].Dns = types.StringNull()

			if t := cr.Get("dns.optionType"); t.Exists() {
				va := cr.Get("dns.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].Dns = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].Dscp = types.Int64Null()

			if t := cr.Get("dscp.optionType"); t.Exists() {
				va := cr.Get("dscp.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].Dscp = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Matches[ci].PacketLength = types.StringNull()

			if t := cr.Get("packetLength.optionType"); t.Exists() {
				va := cr.Get("packetLength.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].PacketLength = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].Protocol = types.SetNull(types.StringType)

			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].Protocol = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].Matches[ci].SourceDataPrefixListId = types.StringNull()

			if t := cr.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("sourceDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].SourceDataPrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].SourcePort = types.SetNull(types.StringType)

			if t := cr.Get("sourcePort.optionType"); t.Exists() {
				va := cr.Get("sourcePort.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].SourcePort = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].Matches[ci].DestinationDataPrefixListId = types.StringNull()

			if t := cr.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("destinationDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].DestinationDataPrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].DestinationPort = types.SetNull(types.StringType)

			if t := cr.Get("destinationPort.optionType"); t.Exists() {
				va := cr.Get("destinationPort.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].DestinationPort = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].Matches[ci].DestinationRegion = types.StringNull()

			if t := cr.Get("destinationRegion.optionType"); t.Exists() {
				va := cr.Get("destinationRegion.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].DestinationRegion = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].Tcp = types.StringNull()

			if t := cr.Get("tcp.optionType"); t.Exists() {
				va := cr.Get("tcp.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].Tcp = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Matches[ci].TrafficTo = types.StringNull()

			if t := cr.Get("trafficTo.optionType"); t.Exists() {
				va := cr.Get("trafficTo.value")
				if t.String() == "global" {
					data.Sequences[i].Matches[ci].TrafficTo = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Sequences[i].Actions {
			keys := [...]string{}
			keyValues := [...]string{}
			keyValuesVariables := [...]string{}

			var cr gjson.Result
			r.Get("actions").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Sequences[i].Actions[ci].Counter = types.StringNull()

			if t := cr.Get("count.optionType"); t.Exists() {
				va := cr.Get("count.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Counter = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].Log = types.BoolNull()

			if t := cr.Get("log.optionType"); t.Exists() {
				va := cr.Get("log.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Log = types.BoolValue(va.Bool())
				}
			}
			for cci := range data.Sequences[i].Actions[ci].SlaClass {
				keys := [...]string{}
				keyValues := [...]string{}
				keyValuesVariables := [...]string{}

				var ccr gjson.Result
				cr.Get("slaClass").ForEach(
					func(_, v gjson.Result) bool {
						found := false
						for ik := range keys {
							tt := v.Get(keys[ik] + ".optionType").String()
							vv := v.Get(keys[ik] + ".value").String()
							if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
								found = true
								continue
							}
							found = false
							break
						}
						if found {
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].Actions[ci].SlaClass[cci].SlaClassListId = types.StringNull()

				if t := ccr.Get("slaName.refId.optionType"); t.Exists() {
					va := ccr.Get("slaName.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClass[cci].SlaClassListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].SlaClass[cci].PreferredColor = types.SetNull(types.StringType)

				if t := ccr.Get("preferredColor.optionType"); t.Exists() {
					va := ccr.Get("preferredColor.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].SlaClass[cci].PreferredColor = helpers.GetStringSet(va.Array())
					}
				}
			}
			for cci := range data.Sequences[i].Actions[ci].Sets {
				keys := [...]string{}
				keyValues := [...]string{}
				keyValuesVariables := [...]string{}

				var ccr gjson.Result
				cr.Get("set").ForEach(
					func(_, v gjson.Result) bool {
						found := false
						for ik := range keys {
							tt := v.Get(keys[ik] + ".optionType").String()
							vv := v.Get(keys[ik] + ".value").String()
							if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
								found = true
								continue
							}
							found = false
							break
						}
						if found {
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].Actions[ci].Sets[cci].Dscp = types.Int64Null()

				if t := ccr.Get("dscp.optionType"); t.Exists() {
					va := ccr.Get("dscp.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].Dscp = types.Int64Value(va.Int())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].PolicerId = types.StringNull()

				if t := ccr.Get("policer.refId.optionType"); t.Exists() {
					va := ccr.Get("policer.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].PolicerId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].PreferredColorGroupId = types.StringNull()

				if t := ccr.Get("preferredColorGroup.refId.optionType"); t.Exists() {
					va := ccr.Get("preferredColorGroup.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].PreferredColorGroupId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ForwardingClassId = types.StringNull()

				if t := ccr.Get("forwardingClass.refId.optionType"); t.Exists() {
					va := ccr.Get("forwardingClass.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ForwardingClassId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].LocalTlocListColor = types.SetNull(types.StringType)

				if t := ccr.Get("localTlocList.color.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].LocalTlocListColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].LocalTlocRestrict = types.StringNull()

				if t := ccr.Get("localTlocList.restrict.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.restrict.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].LocalTlocRestrict = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].LocalTlocListEncapsulation = types.StringNull()

				if t := ccr.Get("localTlocList.encap.optionType"); t.Exists() {
					va := ccr.Get("localTlocList.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].LocalTlocListEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].TlocIp = types.StringNull()

				if t := ccr.Get("tloc.ip.optionType"); t.Exists() {
					va := ccr.Get("tloc.ip.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].TlocIp = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].TlocColor = types.SetNull(types.StringType)

				if t := ccr.Get("tloc.color.optionType"); t.Exists() {
					va := ccr.Get("tloc.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].TlocColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].TlocEncapsulation = types.StringNull()

				if t := ccr.Get("tloc.encap.optionType"); t.Exists() {
					va := ccr.Get("tloc.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].TlocEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].TlocListId = types.StringNull()

				if t := ccr.Get("tlocList.refId.optionType"); t.Exists() {
					va := ccr.Get("tlocList.refId.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].TlocListId = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ServiceType = types.StringNull()

				if t := ccr.Get("service.type.optionType"); t.Exists() {
					va := ccr.Get("service.type.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ServiceType = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ServiceColor = types.SetNull(types.StringType)

				if t := ccr.Get("service.tloc.color.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.color.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ServiceColor = helpers.GetStringSet(va.Array())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ServiceEncapsulation = types.StringNull()

				if t := ccr.Get("service.tloc.encap.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.encap.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ServiceEncapsulation = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ServiceTlocIp = types.StringNull()

				if t := ccr.Get("service.tloc.ip.optionType"); t.Exists() {
					va := ccr.Get("service.tloc.ip.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ServiceTlocIp = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].ServiceVpn = types.StringNull()

				if t := ccr.Get("service.vpn.optionType"); t.Exists() {
					va := ccr.Get("service.vpn.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].ServiceVpn = types.StringValue(va.String())
					}
				}
				data.Sequences[i].Actions[ci].Sets[cci].NextHop = types.StringNull()

				if t := ccr.Get("nextHop.optionType"); t.Exists() {
					va := ccr.Get("nextHop.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Sets[cci].NextHop = types.StringValue(va.String())
					}
				}
			}
			data.Sequences[i].Actions[ci].RedirectDnsField = types.StringNull()

			if t := cr.Get("redirectDns.field.optionType"); t.Exists() {
				va := cr.Get("redirectDns.field.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].RedirectDnsField = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].RedirectDnsValue = types.StringNull()

			if t := cr.Get("redirectDns.value.optionType"); t.Exists() {
				va := cr.Get("redirectDns.value.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].RedirectDnsValue = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].NatPool = types.Int64Null()

			if t := cr.Get("natPool.optionType"); t.Exists() {
				va := cr.Get("natPool.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatPool = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].NatVpn = types.BoolNull()

			if t := cr.Get("nat.useVpn.optionType"); t.Exists() {
				va := cr.Get("nat.useVpn.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatVpn = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].NatFallback = types.BoolNull()

			if t := cr.Get("nat.fallback.optionType"); t.Exists() {
				va := cr.Get("nat.fallback.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].NatFallback = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].SecureInternetGateway = types.BoolNull()

			if t := cr.Get("sig.optionType"); t.Exists() {
				va := cr.Get("sig.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].SecureInternetGateway = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].FallbackToRouting = types.BoolNull()

			if t := cr.Get("fallbackToRouting.optionType"); t.Exists() {
				va := cr.Get("fallbackToRouting.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].FallbackToRouting = types.BoolValue(va.Bool())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ApplicationPriorityTrafficPolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.DefaultAction.IsNull() {
		return false
	}
	if !data.Vpn.IsNull() {
		return false
	}
	if !data.TargetDirection.IsNull() {
		return false
	}
	if len(data.Sequences) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull