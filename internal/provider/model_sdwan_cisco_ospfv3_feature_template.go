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
type CiscoOSPFv3 struct {
	Id                                                types.String                         `tfsdk:"id"`
	Version                                           types.Int64                          `tfsdk:"version"`
	TemplateType                                      types.String                         `tfsdk:"template_type"`
	Name                                              types.String                         `tfsdk:"name"`
	Description                                       types.String                         `tfsdk:"description"`
	DeviceTypes                                       types.Set                            `tfsdk:"device_types"`
	Ipv4RouterId                                      types.String                         `tfsdk:"ipv4_router_id"`
	Ipv4RouterIdVariable                              types.String                         `tfsdk:"ipv4_router_id_variable"`
	Ipv4AutoCostReferenceBandwidth                    types.Int64                          `tfsdk:"ipv4_auto_cost_reference_bandwidth"`
	Ipv4AutoCostReferenceBandwidthVariable            types.String                         `tfsdk:"ipv4_auto_cost_reference_bandwidth_variable"`
	Ipv4CompatibleRfc1583                             types.Bool                           `tfsdk:"ipv4_compatible_rfc1583"`
	Ipv4CompatibleRfc1583Variable                     types.String                         `tfsdk:"ipv4_compatible_rfc1583_variable"`
	Ipv4DefaultInformationOriginate                   types.Bool                           `tfsdk:"ipv4_default_information_originate"`
	Ipv4DefaultInformationOriginateAlways             types.Bool                           `tfsdk:"ipv4_default_information_originate_always"`
	Ipv4DefaultInformationOriginateAlwaysVariable     types.String                         `tfsdk:"ipv4_default_information_originate_always_variable"`
	Ipv4DefaultInformationOriginateMetric             types.Int64                          `tfsdk:"ipv4_default_information_originate_metric"`
	Ipv4DefaultInformationOriginateMetricVariable     types.String                         `tfsdk:"ipv4_default_information_originate_metric_variable"`
	Ipv4DefaultInformationOriginateMetricType         types.String                         `tfsdk:"ipv4_default_information_originate_metric_type"`
	Ipv4DefaultInformationOriginateMetricTypeVariable types.String                         `tfsdk:"ipv4_default_information_originate_metric_type_variable"`
	Ipv4DistanceExternal                              types.Int64                          `tfsdk:"ipv4_distance_external"`
	Ipv4DistanceExternalVariable                      types.String                         `tfsdk:"ipv4_distance_external_variable"`
	Ipv4DistanceInterArea                             types.Int64                          `tfsdk:"ipv4_distance_inter_area"`
	Ipv4DistanceInterAreaVariable                     types.String                         `tfsdk:"ipv4_distance_inter_area_variable"`
	Ipv4DistanceIntraArea                             types.Int64                          `tfsdk:"ipv4_distance_intra_area"`
	Ipv4DistanceIntraAreaVariable                     types.String                         `tfsdk:"ipv4_distance_intra_area_variable"`
	Ipv4TimersSpfDelay                                types.Int64                          `tfsdk:"ipv4_timers_spf_delay"`
	Ipv4TimersSpfDelayVariable                        types.String                         `tfsdk:"ipv4_timers_spf_delay_variable"`
	Ipv4TimersSpfInitialHold                          types.Int64                          `tfsdk:"ipv4_timers_spf_initial_hold"`
	Ipv4TimersSpfInitialHoldVariable                  types.String                         `tfsdk:"ipv4_timers_spf_initial_hold_variable"`
	Ipv4TimersSpfMaxHold                              types.Int64                          `tfsdk:"ipv4_timers_spf_max_hold"`
	Ipv4TimersSpfMaxHoldVariable                      types.String                         `tfsdk:"ipv4_timers_spf_max_hold_variable"`
	Ipv4Distance                                      types.Int64                          `tfsdk:"ipv4_distance"`
	Ipv4DistanceVariable                              types.String                         `tfsdk:"ipv4_distance_variable"`
	Ipv4PolicyName                                    types.String                         `tfsdk:"ipv4_policy_name"`
	Ipv4PolicyNameVariable                            types.String                         `tfsdk:"ipv4_policy_name_variable"`
	Ipv4Filter                                        types.Bool                           `tfsdk:"ipv4_filter"`
	Ipv4FilterVariable                                types.String                         `tfsdk:"ipv4_filter_variable"`
	Ipv4Redistributes                                 []CiscoOSPFv3Ipv4Redistributes       `tfsdk:"ipv4_redistributes"`
	Ipv4MaxMetricRouterLsas                           []CiscoOSPFv3Ipv4MaxMetricRouterLsas `tfsdk:"ipv4_max_metric_router_lsas"`
	Ipv4Areas                                         []CiscoOSPFv3Ipv4Areas               `tfsdk:"ipv4_areas"`
	Ipv6RouterId                                      types.String                         `tfsdk:"ipv6_router_id"`
	Ipv6RouterIdVariable                              types.String                         `tfsdk:"ipv6_router_id_variable"`
	Ipv6AutoCostReferenceBandwidth                    types.Int64                          `tfsdk:"ipv6_auto_cost_reference_bandwidth"`
	Ipv6AutoCostReferenceBandwidthVariable            types.String                         `tfsdk:"ipv6_auto_cost_reference_bandwidth_variable"`
	Ipv6CompatibleRfc1583                             types.Bool                           `tfsdk:"ipv6_compatible_rfc1583"`
	Ipv6CompatibleRfc1583Variable                     types.String                         `tfsdk:"ipv6_compatible_rfc1583_variable"`
	Ipv6DefaultInformationOriginate                   types.Bool                           `tfsdk:"ipv6_default_information_originate"`
	Ipv6DefaultInformationOriginateAlways             types.Bool                           `tfsdk:"ipv6_default_information_originate_always"`
	Ipv6DefaultInformationOriginateAlwaysVariable     types.String                         `tfsdk:"ipv6_default_information_originate_always_variable"`
	Ipv6DefaultInformationOriginateMetric             types.Int64                          `tfsdk:"ipv6_default_information_originate_metric"`
	Ipv6DefaultInformationOriginateMetricVariable     types.String                         `tfsdk:"ipv6_default_information_originate_metric_variable"`
	Ipv6DefaultInformationOriginateMetricType         types.String                         `tfsdk:"ipv6_default_information_originate_metric_type"`
	Ipv6DefaultInformationOriginateMetricTypeVariable types.String                         `tfsdk:"ipv6_default_information_originate_metric_type_variable"`
	Ipv6DistanceExternal                              types.Int64                          `tfsdk:"ipv6_distance_external"`
	Ipv6DistanceExternalVariable                      types.String                         `tfsdk:"ipv6_distance_external_variable"`
	Ipv6DistanceInterArea                             types.Int64                          `tfsdk:"ipv6_distance_inter_area"`
	Ipv6DistanceInterAreaVariable                     types.String                         `tfsdk:"ipv6_distance_inter_area_variable"`
	Ipv6DistanceIntraArea                             types.Int64                          `tfsdk:"ipv6_distance_intra_area"`
	Ipv6DistanceIntraAreaVariable                     types.String                         `tfsdk:"ipv6_distance_intra_area_variable"`
	Ipv6TimersSpfDelay                                types.Int64                          `tfsdk:"ipv6_timers_spf_delay"`
	Ipv6TimersSpfDelayVariable                        types.String                         `tfsdk:"ipv6_timers_spf_delay_variable"`
	Ipv6TimersSpfInitialHold                          types.Int64                          `tfsdk:"ipv6_timers_spf_initial_hold"`
	Ipv6TimersSpfInitialHoldVariable                  types.String                         `tfsdk:"ipv6_timers_spf_initial_hold_variable"`
	Ipv6TimersSpfMaxHold                              types.Int64                          `tfsdk:"ipv6_timers_spf_max_hold"`
	Ipv6TimersSpfMaxHoldVariable                      types.String                         `tfsdk:"ipv6_timers_spf_max_hold_variable"`
	Ipv6Distance                                      types.Int64                          `tfsdk:"ipv6_distance"`
	Ipv6DistanceVariable                              types.String                         `tfsdk:"ipv6_distance_variable"`
	Ipv6PolicyName                                    types.String                         `tfsdk:"ipv6_policy_name"`
	Ipv6PolicyNameVariable                            types.String                         `tfsdk:"ipv6_policy_name_variable"`
	Ipv6Filter                                        types.Bool                           `tfsdk:"ipv6_filter"`
	Ipv6FilterVariable                                types.String                         `tfsdk:"ipv6_filter_variable"`
	Ipv6Redistributes                                 []CiscoOSPFv3Ipv6Redistributes       `tfsdk:"ipv6_redistributes"`
	Ipv6MaxMetricRouterLsas                           []CiscoOSPFv3Ipv6MaxMetricRouterLsas `tfsdk:"ipv6_max_metric_router_lsas"`
	Ipv6Areas                                         []CiscoOSPFv3Ipv6Areas               `tfsdk:"ipv6_areas"`
}

type CiscoOSPFv3Ipv4Redistributes struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
	NatDia              types.Bool   `tfsdk:"nat_dia"`
	NatDiaVariable      types.String `tfsdk:"nat_dia_variable"`
}

type CiscoOSPFv3Ipv4MaxMetricRouterLsas struct {
	Optional     types.Bool   `tfsdk:"optional"`
	AdType       types.String `tfsdk:"ad_type"`
	Time         types.Int64  `tfsdk:"time"`
	TimeVariable types.String `tfsdk:"time_variable"`
}

type CiscoOSPFv3Ipv4Areas struct {
	Optional              types.Bool                       `tfsdk:"optional"`
	AreaNumber            types.Int64                      `tfsdk:"area_number"`
	AreaNumberVariable    types.String                     `tfsdk:"area_number_variable"`
	Stub                  types.Bool                       `tfsdk:"stub"`
	StubNoSummary         types.Bool                       `tfsdk:"stub_no_summary"`
	StubNoSummaryVariable types.String                     `tfsdk:"stub_no_summary_variable"`
	Nssa                  types.Bool                       `tfsdk:"nssa"`
	NssaNoSummary         types.Bool                       `tfsdk:"nssa_no_summary"`
	NssaNoSummaryVariable types.String                     `tfsdk:"nssa_no_summary_variable"`
	Translate             types.String                     `tfsdk:"translate"`
	TranslateVariable     types.String                     `tfsdk:"translate_variable"`
	Normal                types.Bool                       `tfsdk:"normal"`
	NormalVariable        types.String                     `tfsdk:"normal_variable"`
	Interfaces            []CiscoOSPFv3Ipv4AreasInterfaces `tfsdk:"interfaces"`
	Ranges                []CiscoOSPFv3Ipv4AreasRanges     `tfsdk:"ranges"`
}

type CiscoOSPFv3Ipv6Redistributes struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
}

type CiscoOSPFv3Ipv6MaxMetricRouterLsas struct {
	Optional     types.Bool   `tfsdk:"optional"`
	AdType       types.String `tfsdk:"ad_type"`
	Time         types.Int64  `tfsdk:"time"`
	TimeVariable types.String `tfsdk:"time_variable"`
}

type CiscoOSPFv3Ipv6Areas struct {
	Optional              types.Bool                       `tfsdk:"optional"`
	AreaNumber            types.Int64                      `tfsdk:"area_number"`
	AreaNumberVariable    types.String                     `tfsdk:"area_number_variable"`
	Stub                  types.Bool                       `tfsdk:"stub"`
	StubNoSummary         types.Bool                       `tfsdk:"stub_no_summary"`
	StubNoSummaryVariable types.String                     `tfsdk:"stub_no_summary_variable"`
	Nssa                  types.Bool                       `tfsdk:"nssa"`
	NssaNoSummary         types.Bool                       `tfsdk:"nssa_no_summary"`
	NssaNoSummaryVariable types.String                     `tfsdk:"nssa_no_summary_variable"`
	Translate             types.String                     `tfsdk:"translate"`
	TranslateVariable     types.String                     `tfsdk:"translate_variable"`
	Normal                types.Bool                       `tfsdk:"normal"`
	NormalVariable        types.String                     `tfsdk:"normal_variable"`
	Interfaces            []CiscoOSPFv3Ipv6AreasInterfaces `tfsdk:"interfaces"`
	Ranges                []CiscoOSPFv3Ipv6AreasRanges     `tfsdk:"ranges"`
}

type CiscoOSPFv3Ipv4AreasInterfaces struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	Name                       types.String `tfsdk:"name"`
	NameVariable               types.String `tfsdk:"name_variable"`
	HelloInterval              types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable      types.String `tfsdk:"hello_interval_variable"`
	DeadInterval               types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable       types.String `tfsdk:"dead_interval_variable"`
	RetransmitInterval         types.Int64  `tfsdk:"retransmit_interval"`
	RetransmitIntervalVariable types.String `tfsdk:"retransmit_interval_variable"`
	Cost                       types.Int64  `tfsdk:"cost"`
	CostVariable               types.String `tfsdk:"cost_variable"`
	Network                    types.String `tfsdk:"network"`
	NetworkVariable            types.String `tfsdk:"network_variable"`
	PassiveInterface           types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable   types.String `tfsdk:"passive_interface_variable"`
	AuthenticationType         types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable types.String `tfsdk:"authentication_type_variable"`
	AuthenticationKey          types.String `tfsdk:"authentication_key"`
	AuthenticationKeyVariable  types.String `tfsdk:"authentication_key_variable"`
	IpsecSpi                   types.Int64  `tfsdk:"ipsec_spi"`
	IpsecSpiVariable           types.String `tfsdk:"ipsec_spi_variable"`
}
type CiscoOSPFv3Ipv4AreasRanges struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Address             types.String `tfsdk:"address"`
	AddressVariable     types.String `tfsdk:"address_variable"`
	Cost                types.Int64  `tfsdk:"cost"`
	CostVariable        types.String `tfsdk:"cost_variable"`
	NoAdvertise         types.Bool   `tfsdk:"no_advertise"`
	NoAdvertiseVariable types.String `tfsdk:"no_advertise_variable"`
}

type CiscoOSPFv3Ipv6AreasInterfaces struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	Name                       types.String `tfsdk:"name"`
	NameVariable               types.String `tfsdk:"name_variable"`
	HelloInterval              types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable      types.String `tfsdk:"hello_interval_variable"`
	DeadInterval               types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable       types.String `tfsdk:"dead_interval_variable"`
	RetransmitInterval         types.Int64  `tfsdk:"retransmit_interval"`
	RetransmitIntervalVariable types.String `tfsdk:"retransmit_interval_variable"`
	Cost                       types.Int64  `tfsdk:"cost"`
	CostVariable               types.String `tfsdk:"cost_variable"`
	Network                    types.String `tfsdk:"network"`
	NetworkVariable            types.String `tfsdk:"network_variable"`
	PassiveInterface           types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable   types.String `tfsdk:"passive_interface_variable"`
	AuthenticationType         types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable types.String `tfsdk:"authentication_type_variable"`
	AuthenticationKey          types.String `tfsdk:"authentication_key"`
	AuthenticationKeyVariable  types.String `tfsdk:"authentication_key_variable"`
	IpsecSpi                   types.Int64  `tfsdk:"ipsec_spi"`
	IpsecSpiVariable           types.String `tfsdk:"ipsec_spi_variable"`
}
type CiscoOSPFv3Ipv6AreasRanges struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Address             types.String `tfsdk:"address"`
	AddressVariable     types.String `tfsdk:"address_variable"`
	Cost                types.Int64  `tfsdk:"cost"`
	CostVariable        types.String `tfsdk:"cost_variable"`
	NoAdvertise         types.Bool   `tfsdk:"no_advertise"`
	NoAdvertiseVariable types.String `tfsdk:"no_advertise_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoOSPFv3) getModel() string {
	return "cisco_ospfv3"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoOSPFv3) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_ospfv3")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.Ipv4RouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipVariableName", data.Ipv4RouterIdVariable.ValueString())
	} else if data.Ipv4RouterId.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.router-id."+"vipValue", data.Ipv4RouterId.ValueString())
	}

	if !data.Ipv4AutoCostReferenceBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipVariableName", data.Ipv4AutoCostReferenceBandwidthVariable.ValueString())
	} else if data.Ipv4AutoCostReferenceBandwidth.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.auto-cost.reference-bandwidth."+"vipValue", data.Ipv4AutoCostReferenceBandwidth.ValueInt64())
	}

	if !data.Ipv4CompatibleRfc1583Variable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipVariableName", data.Ipv4CompatibleRfc1583Variable.ValueString())
	} else if data.Ipv4CompatibleRfc1583.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.compatible.rfc1583."+"vipValue", strconv.FormatBool(data.Ipv4CompatibleRfc1583.ValueBool()))
	}
	if !data.Ipv4DefaultInformationOriginate.IsNull() {
		if data.Ipv4DefaultInformationOriginate.ValueBool() {
			if !gjson.Get(body, path+"ospfv3.address-family.ipv4.default-information.originate").Exists() {
				body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate."+"vipType", "ignore")
		}
	}

	if !data.Ipv4DefaultInformationOriginateAlwaysVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipVariableName", data.Ipv4DefaultInformationOriginateAlwaysVariable.ValueString())
	} else if data.Ipv4DefaultInformationOriginateAlways.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.always."+"vipValue", strconv.FormatBool(data.Ipv4DefaultInformationOriginateAlways.ValueBool()))
	}

	if !data.Ipv4DefaultInformationOriginateMetricVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipVariableName", data.Ipv4DefaultInformationOriginateMetricVariable.ValueString())
	} else if data.Ipv4DefaultInformationOriginateMetric.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric."+"vipValue", data.Ipv4DefaultInformationOriginateMetric.ValueInt64())
	}

	if !data.Ipv4DefaultInformationOriginateMetricTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipVariableName", data.Ipv4DefaultInformationOriginateMetricTypeVariable.ValueString())
	} else if data.Ipv4DefaultInformationOriginateMetricType.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.default-information.originate.metric-type."+"vipValue", data.Ipv4DefaultInformationOriginateMetricType.ValueString())
	}

	if !data.Ipv4DistanceExternalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipVariableName", data.Ipv4DistanceExternalVariable.ValueString())
	} else if data.Ipv4DistanceExternal.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.external."+"vipValue", data.Ipv4DistanceExternal.ValueInt64())
	}

	if !data.Ipv4DistanceInterAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipVariableName", data.Ipv4DistanceInterAreaVariable.ValueString())
	} else if data.Ipv4DistanceInterArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area."+"vipValue", data.Ipv4DistanceInterArea.ValueInt64())
	}

	if !data.Ipv4DistanceIntraAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipVariableName", data.Ipv4DistanceIntraAreaVariable.ValueString())
	} else if data.Ipv4DistanceIntraArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area."+"vipValue", data.Ipv4DistanceIntraArea.ValueInt64())
	}

	if !data.Ipv4TimersSpfDelayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipVariableName", data.Ipv4TimersSpfDelayVariable.ValueString())
	} else if data.Ipv4TimersSpfDelay.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.delay."+"vipValue", data.Ipv4TimersSpfDelay.ValueInt64())
	}

	if !data.Ipv4TimersSpfInitialHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipVariableName", data.Ipv4TimersSpfInitialHoldVariable.ValueString())
	} else if data.Ipv4TimersSpfInitialHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold."+"vipValue", data.Ipv4TimersSpfInitialHold.ValueInt64())
	}

	if !data.Ipv4TimersSpfMaxHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipVariableName", data.Ipv4TimersSpfMaxHoldVariable.ValueString())
	} else if data.Ipv4TimersSpfMaxHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.timers.throttle.spf.max-hold."+"vipValue", data.Ipv4TimersSpfMaxHold.ValueInt64())
	}

	if !data.Ipv4DistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipVariableName", data.Ipv4DistanceVariable.ValueString())
	} else if data.Ipv4Distance.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.distance-ipv4.distance."+"vipValue", data.Ipv4Distance.ValueInt64())
	}

	if !data.Ipv4PolicyNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipVariableName", data.Ipv4PolicyNameVariable.ValueString())
	} else if data.Ipv4PolicyName.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.name."+"vipValue", data.Ipv4PolicyName.ValueString())
	}

	if !data.Ipv4FilterVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipVariableName", data.Ipv4FilterVariable.ValueString())
	} else if data.Ipv4Filter.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.table-map.filter."+"vipValue", strconv.FormatBool(data.Ipv4Filter.ValueBool()))
	}
	if len(data.Ipv4Redistributes) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Redistributes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "dia")

		if !item.NatDiaVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipVariableName", item.NatDiaVariable.ValueString())
		} else if item.NatDia.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipValue", strconv.FormatBool(item.NatDia.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv4.redistribute."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4MaxMetricRouterLsas) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4MaxMetricRouterLsas {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "ad-type")
		if item.AdType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipValue", item.AdType.ValueString())
		}
		itemAttributes = append(itemAttributes, "time")

		if !item.TimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipVariableName", item.TimeVariable.ValueString())
		} else if item.Time.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipValue", item.Time.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv4.max-metric.router-lsa."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4Areas) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv4.area."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Areas {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "a-num")

		if !item.AreaNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipVariableName", item.AreaNumberVariable.ValueString())
		} else if item.AreaNumber.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipValue", item.AreaNumber.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "stub")
		if !item.Stub.IsNull() {
			if item.Stub.ValueBool() {
				if !gjson.Get(itemBody, "stub").Exists() {
					itemBody, _ = sjson.Set(itemBody, "stub", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "stub")

		if !item.StubNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipVariableName", item.StubNoSummaryVariable.ValueString())
		} else if item.StubNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipValue", strconv.FormatBool(item.StubNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "nssa")
		if !item.Nssa.IsNull() {
			if item.Nssa.ValueBool() {
				if !gjson.Get(itemBody, "nssa").Exists() {
					itemBody, _ = sjson.Set(itemBody, "nssa", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "nssa")

		if !item.NssaNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipVariableName", item.NssaNoSummaryVariable.ValueString())
		} else if item.NssaNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipValue", strconv.FormatBool(item.NssaNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "nssa")

		if !item.TranslateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipVariableName", item.TranslateVariable.ValueString())
		} else if item.Translate.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipValue", item.Translate.ValueString())
		}
		itemAttributes = append(itemAttributes, "normal")

		if !item.NormalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipVariableName", item.NormalVariable.ValueString())
		} else if item.Normal.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipValue", strconv.FormatBool(item.Normal.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "interface")
		if len(item.Interfaces) > 0 {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Interfaces {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "name")

			if !childItem.NameVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipVariableName", childItem.NameVariable.ValueString())
			} else if childItem.Name.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.Name.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "hello-interval")

			if !childItem.HelloIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipVariableName", childItem.HelloIntervalVariable.ValueString())
			} else if childItem.HelloInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipValue", childItem.HelloInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "dead-interval")

			if !childItem.DeadIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipVariableName", childItem.DeadIntervalVariable.ValueString())
			} else if childItem.DeadInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipValue", childItem.DeadInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "retransmit-interval")

			if !childItem.RetransmitIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipVariableName", childItem.RetransmitIntervalVariable.ValueString())
			} else if childItem.RetransmitInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipValue", childItem.RetransmitInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "network")

			if !childItem.NetworkVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipVariableName", childItem.NetworkVariable.ValueString())
			} else if childItem.Network.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipValue", childItem.Network.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "passive-interface")

			if !childItem.PassiveInterfaceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipVariableName", childItem.PassiveInterfaceVariable.ValueString())
			} else if childItem.PassiveInterface.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipValue", strconv.FormatBool(childItem.PassiveInterface.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipVariableName", childItem.AuthenticationTypeVariable.ValueString())
			} else if childItem.AuthenticationType.IsNull() {
				if !gjson.Get(itemChildBody, "authentication").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "authentication", map[string]interface{}{})
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipValue", childItem.AuthenticationType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipVariableName", childItem.AuthenticationKeyVariable.ValueString())
			} else if childItem.AuthenticationKey.IsNull() {
				if !gjson.Get(itemChildBody, "authentication").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "authentication", map[string]interface{}{})
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipValue", childItem.AuthenticationKey.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.IpsecSpiVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipVariableName", childItem.IpsecSpiVariable.ValueString())
			} else if childItem.IpsecSpi.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipValue", childItem.IpsecSpi.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "interface."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "range")
		if len(item.Ranges) > 0 {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ranges {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "no-advertise")

			if !childItem.NoAdvertiseVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipVariableName", childItem.NoAdvertiseVariable.ValueString())
			} else if childItem.NoAdvertise.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipValue", strconv.FormatBool(childItem.NoAdvertise.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "range."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv4.area."+"vipValue.-1", itemBody)
	}

	if !data.Ipv6RouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipVariableName", data.Ipv6RouterIdVariable.ValueString())
	} else if data.Ipv6RouterId.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.router-id."+"vipValue", data.Ipv6RouterId.ValueString())
	}

	if !data.Ipv6AutoCostReferenceBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipVariableName", data.Ipv6AutoCostReferenceBandwidthVariable.ValueString())
	} else if data.Ipv6AutoCostReferenceBandwidth.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.auto-cost.reference-bandwidth."+"vipValue", data.Ipv6AutoCostReferenceBandwidth.ValueInt64())
	}

	if !data.Ipv6CompatibleRfc1583Variable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipVariableName", data.Ipv6CompatibleRfc1583Variable.ValueString())
	} else if data.Ipv6CompatibleRfc1583.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.compatible.rfc1583."+"vipValue", strconv.FormatBool(data.Ipv6CompatibleRfc1583.ValueBool()))
	}
	if !data.Ipv6DefaultInformationOriginate.IsNull() {
		if data.Ipv6DefaultInformationOriginate.ValueBool() {
			if !gjson.Get(body, path+"ospfv3.address-family.ipv6.default-information.originate").Exists() {
				body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate."+"vipType", "ignore")
		}
	}

	if !data.Ipv6DefaultInformationOriginateAlwaysVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipVariableName", data.Ipv6DefaultInformationOriginateAlwaysVariable.ValueString())
	} else if data.Ipv6DefaultInformationOriginateAlways.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.always."+"vipValue", strconv.FormatBool(data.Ipv6DefaultInformationOriginateAlways.ValueBool()))
	}

	if !data.Ipv6DefaultInformationOriginateMetricVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipVariableName", data.Ipv6DefaultInformationOriginateMetricVariable.ValueString())
	} else if data.Ipv6DefaultInformationOriginateMetric.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric."+"vipValue", data.Ipv6DefaultInformationOriginateMetric.ValueInt64())
	}

	if !data.Ipv6DefaultInformationOriginateMetricTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipVariableName", data.Ipv6DefaultInformationOriginateMetricTypeVariable.ValueString())
	} else if data.Ipv6DefaultInformationOriginateMetricType.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.default-information.originate.metric-type."+"vipValue", data.Ipv6DefaultInformationOriginateMetricType.ValueString())
	}

	if !data.Ipv6DistanceExternalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipVariableName", data.Ipv6DistanceExternalVariable.ValueString())
	} else if data.Ipv6DistanceExternal.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.external."+"vipValue", data.Ipv6DistanceExternal.ValueInt64())
	}

	if !data.Ipv6DistanceInterAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipVariableName", data.Ipv6DistanceInterAreaVariable.ValueString())
	} else if data.Ipv6DistanceInterArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area."+"vipValue", data.Ipv6DistanceInterArea.ValueInt64())
	}

	if !data.Ipv6DistanceIntraAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipVariableName", data.Ipv6DistanceIntraAreaVariable.ValueString())
	} else if data.Ipv6DistanceIntraArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area."+"vipValue", data.Ipv6DistanceIntraArea.ValueInt64())
	}

	if !data.Ipv6TimersSpfDelayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipVariableName", data.Ipv6TimersSpfDelayVariable.ValueString())
	} else if data.Ipv6TimersSpfDelay.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.delay."+"vipValue", data.Ipv6TimersSpfDelay.ValueInt64())
	}

	if !data.Ipv6TimersSpfInitialHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipVariableName", data.Ipv6TimersSpfInitialHoldVariable.ValueString())
	} else if data.Ipv6TimersSpfInitialHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold."+"vipValue", data.Ipv6TimersSpfInitialHold.ValueInt64())
	}

	if !data.Ipv6TimersSpfMaxHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipVariableName", data.Ipv6TimersSpfMaxHoldVariable.ValueString())
	} else if data.Ipv6TimersSpfMaxHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.timers.throttle.spf.max-hold."+"vipValue", data.Ipv6TimersSpfMaxHold.ValueInt64())
	}

	if !data.Ipv6DistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipVariableName", data.Ipv6DistanceVariable.ValueString())
	} else if data.Ipv6Distance.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.distance-ipv6.distance."+"vipValue", data.Ipv6Distance.ValueInt64())
	}

	if !data.Ipv6PolicyNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipVariableName", data.Ipv6PolicyNameVariable.ValueString())
	} else if data.Ipv6PolicyName.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.name."+"vipValue", data.Ipv6PolicyName.ValueString())
	}

	if !data.Ipv6FilterVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipVariableName", data.Ipv6FilterVariable.ValueString())
	} else if data.Ipv6Filter.IsNull() {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.table-map.filter."+"vipValue", strconv.FormatBool(data.Ipv6Filter.ValueBool()))
	}
	if len(data.Ipv6Redistributes) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Redistributes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv6.redistribute."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6MaxMetricRouterLsas) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6MaxMetricRouterLsas {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "ad-type")
		if item.AdType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipValue", item.AdType.ValueString())
		}
		itemAttributes = append(itemAttributes, "time")

		if !item.TimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipVariableName", item.TimeVariable.ValueString())
		} else if item.Time.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipValue", item.Time.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv6.max-metric.router-lsa."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6Areas) > 0 {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospfv3.address-family.ipv6.area."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Areas {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "a-num")

		if !item.AreaNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipVariableName", item.AreaNumberVariable.ValueString())
		} else if item.AreaNumber.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipValue", item.AreaNumber.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "stub")
		if !item.Stub.IsNull() {
			if item.Stub.ValueBool() {
				if !gjson.Get(itemBody, "stub").Exists() {
					itemBody, _ = sjson.Set(itemBody, "stub", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "stub")

		if !item.StubNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipVariableName", item.StubNoSummaryVariable.ValueString())
		} else if item.StubNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipValue", strconv.FormatBool(item.StubNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "nssa")
		if !item.Nssa.IsNull() {
			if item.Nssa.ValueBool() {
				if !gjson.Get(itemBody, "nssa").Exists() {
					itemBody, _ = sjson.Set(itemBody, "nssa", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "nssa")

		if !item.NssaNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipVariableName", item.NssaNoSummaryVariable.ValueString())
		} else if item.NssaNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipValue", strconv.FormatBool(item.NssaNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "nssa")

		if !item.TranslateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipVariableName", item.TranslateVariable.ValueString())
		} else if item.Translate.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nssa.translate."+"vipValue", item.Translate.ValueString())
		}
		itemAttributes = append(itemAttributes, "normal")

		if !item.NormalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipVariableName", item.NormalVariable.ValueString())
		} else if item.Normal.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "normal."+"vipValue", strconv.FormatBool(item.Normal.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "interface")
		if len(item.Interfaces) > 0 {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Interfaces {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "name")

			if !childItem.NameVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipVariableName", childItem.NameVariable.ValueString())
			} else if childItem.Name.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.Name.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "hello-interval")

			if !childItem.HelloIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipVariableName", childItem.HelloIntervalVariable.ValueString())
			} else if childItem.HelloInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipValue", childItem.HelloInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "dead-interval")

			if !childItem.DeadIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipVariableName", childItem.DeadIntervalVariable.ValueString())
			} else if childItem.DeadInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipValue", childItem.DeadInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "retransmit-interval")

			if !childItem.RetransmitIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipVariableName", childItem.RetransmitIntervalVariable.ValueString())
			} else if childItem.RetransmitInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipValue", childItem.RetransmitInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "network")

			if !childItem.NetworkVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipVariableName", childItem.NetworkVariable.ValueString())
			} else if childItem.Network.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipValue", childItem.Network.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "passive-interface")

			if !childItem.PassiveInterfaceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipVariableName", childItem.PassiveInterfaceVariable.ValueString())
			} else if childItem.PassiveInterface.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipValue", strconv.FormatBool(childItem.PassiveInterface.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipVariableName", childItem.AuthenticationTypeVariable.ValueString())
			} else if childItem.AuthenticationType.IsNull() {
				if !gjson.Get(itemChildBody, "authentication").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "authentication", map[string]interface{}{})
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipValue", childItem.AuthenticationType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipVariableName", childItem.AuthenticationKeyVariable.ValueString())
			} else if childItem.AuthenticationKey.IsNull() {
				if !gjson.Get(itemChildBody, "authentication").Exists() {
					itemChildBody, _ = sjson.Set(itemChildBody, "authentication", map[string]interface{}{})
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.authentication-key."+"vipValue", childItem.AuthenticationKey.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.IpsecSpiVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipVariableName", childItem.IpsecSpiVariable.ValueString())
			} else if childItem.IpsecSpi.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.ipsec.spi."+"vipValue", childItem.IpsecSpi.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "interface."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "range")
		if len(item.Ranges) > 0 {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ranges {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "no-advertise")

			if !childItem.NoAdvertiseVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipVariableName", childItem.NoAdvertiseVariable.ValueString())
			} else if childItem.NoAdvertise.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipValue", strconv.FormatBool(childItem.NoAdvertise.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "range."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospfv3.address-family.ipv6.area."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoOSPFv3) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "ospfv3.address-family.ipv4.router-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4RouterId = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.router-id.vipVariableName")
			data.Ipv4RouterIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4RouterId = types.StringNull()
			data.Ipv4RouterIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.router-id.vipValue")
			data.Ipv4RouterId = types.StringValue(v.String())
			data.Ipv4RouterIdVariable = types.StringNull()
		}
	} else {
		data.Ipv4RouterId = types.StringNull()
		data.Ipv4RouterIdVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.auto-cost.reference-bandwidth.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4AutoCostReferenceBandwidth = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.auto-cost.reference-bandwidth.vipVariableName")
			data.Ipv4AutoCostReferenceBandwidthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4AutoCostReferenceBandwidth = types.Int64Null()
			data.Ipv4AutoCostReferenceBandwidthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.auto-cost.reference-bandwidth.vipValue")
			data.Ipv4AutoCostReferenceBandwidth = types.Int64Value(v.Int())
			data.Ipv4AutoCostReferenceBandwidthVariable = types.StringNull()
		}
	} else {
		data.Ipv4AutoCostReferenceBandwidth = types.Int64Null()
		data.Ipv4AutoCostReferenceBandwidthVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.compatible.rfc1583.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4CompatibleRfc1583 = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.compatible.rfc1583.vipVariableName")
			data.Ipv4CompatibleRfc1583Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4CompatibleRfc1583 = types.BoolNull()
			data.Ipv4CompatibleRfc1583Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.compatible.rfc1583.vipValue")
			data.Ipv4CompatibleRfc1583 = types.BoolValue(v.Bool())
			data.Ipv4CompatibleRfc1583Variable = types.StringNull()
		}
	} else {
		data.Ipv4CompatibleRfc1583 = types.BoolNull()
		data.Ipv4CompatibleRfc1583Variable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DefaultInformationOriginate = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Ipv4DefaultInformationOriginate = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.vipValue")
			data.Ipv4DefaultInformationOriginate = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate"); value.Exists() {
		data.Ipv4DefaultInformationOriginate = types.BoolValue(true)

	} else {
		data.Ipv4DefaultInformationOriginate = types.BoolNull()

	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.always.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DefaultInformationOriginateAlways = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.always.vipVariableName")
			data.Ipv4DefaultInformationOriginateAlwaysVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DefaultInformationOriginateAlways = types.BoolNull()
			data.Ipv4DefaultInformationOriginateAlwaysVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.always.vipValue")
			data.Ipv4DefaultInformationOriginateAlways = types.BoolValue(v.Bool())
			data.Ipv4DefaultInformationOriginateAlwaysVariable = types.StringNull()
		}
	} else {
		data.Ipv4DefaultInformationOriginateAlways = types.BoolNull()
		data.Ipv4DefaultInformationOriginateAlwaysVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DefaultInformationOriginateMetric = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric.vipVariableName")
			data.Ipv4DefaultInformationOriginateMetricVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DefaultInformationOriginateMetric = types.Int64Null()
			data.Ipv4DefaultInformationOriginateMetricVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric.vipValue")
			data.Ipv4DefaultInformationOriginateMetric = types.Int64Value(v.Int())
			data.Ipv4DefaultInformationOriginateMetricVariable = types.StringNull()
		}
	} else {
		data.Ipv4DefaultInformationOriginateMetric = types.Int64Null()
		data.Ipv4DefaultInformationOriginateMetricVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DefaultInformationOriginateMetricType = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric-type.vipVariableName")
			data.Ipv4DefaultInformationOriginateMetricTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DefaultInformationOriginateMetricType = types.StringNull()
			data.Ipv4DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.default-information.originate.metric-type.vipValue")
			data.Ipv4DefaultInformationOriginateMetricType = types.StringValue(v.String())
			data.Ipv4DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		}
	} else {
		data.Ipv4DefaultInformationOriginateMetricType = types.StringNull()
		data.Ipv4DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.external.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DistanceExternal = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.external.vipVariableName")
			data.Ipv4DistanceExternalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DistanceExternal = types.Int64Null()
			data.Ipv4DistanceExternalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.external.vipValue")
			data.Ipv4DistanceExternal = types.Int64Value(v.Int())
			data.Ipv4DistanceExternalVariable = types.StringNull()
		}
	} else {
		data.Ipv4DistanceExternal = types.Int64Null()
		data.Ipv4DistanceExternalVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DistanceInterArea = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area.vipVariableName")
			data.Ipv4DistanceInterAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DistanceInterArea = types.Int64Null()
			data.Ipv4DistanceInterAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.inter-area.vipValue")
			data.Ipv4DistanceInterArea = types.Int64Value(v.Int())
			data.Ipv4DistanceInterAreaVariable = types.StringNull()
		}
	} else {
		data.Ipv4DistanceInterArea = types.Int64Null()
		data.Ipv4DistanceInterAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4DistanceIntraArea = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area.vipVariableName")
			data.Ipv4DistanceIntraAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DistanceIntraArea = types.Int64Null()
			data.Ipv4DistanceIntraAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.ospf.intra-area.vipValue")
			data.Ipv4DistanceIntraArea = types.Int64Value(v.Int())
			data.Ipv4DistanceIntraAreaVariable = types.StringNull()
		}
	} else {
		data.Ipv4DistanceIntraArea = types.Int64Null()
		data.Ipv4DistanceIntraAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.delay.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4TimersSpfDelay = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.delay.vipVariableName")
			data.Ipv4TimersSpfDelayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4TimersSpfDelay = types.Int64Null()
			data.Ipv4TimersSpfDelayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.delay.vipValue")
			data.Ipv4TimersSpfDelay = types.Int64Value(v.Int())
			data.Ipv4TimersSpfDelayVariable = types.StringNull()
		}
	} else {
		data.Ipv4TimersSpfDelay = types.Int64Null()
		data.Ipv4TimersSpfDelayVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4TimersSpfInitialHold = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold.vipVariableName")
			data.Ipv4TimersSpfInitialHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4TimersSpfInitialHold = types.Int64Null()
			data.Ipv4TimersSpfInitialHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.initial-hold.vipValue")
			data.Ipv4TimersSpfInitialHold = types.Int64Value(v.Int())
			data.Ipv4TimersSpfInitialHoldVariable = types.StringNull()
		}
	} else {
		data.Ipv4TimersSpfInitialHold = types.Int64Null()
		data.Ipv4TimersSpfInitialHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.max-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4TimersSpfMaxHold = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.max-hold.vipVariableName")
			data.Ipv4TimersSpfMaxHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4TimersSpfMaxHold = types.Int64Null()
			data.Ipv4TimersSpfMaxHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.timers.throttle.spf.max-hold.vipValue")
			data.Ipv4TimersSpfMaxHold = types.Int64Value(v.Int())
			data.Ipv4TimersSpfMaxHoldVariable = types.StringNull()
		}
	} else {
		data.Ipv4TimersSpfMaxHold = types.Int64Null()
		data.Ipv4TimersSpfMaxHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.distance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4Distance = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.distance.vipVariableName")
			data.Ipv4DistanceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4Distance = types.Int64Null()
			data.Ipv4DistanceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.distance-ipv4.distance.vipValue")
			data.Ipv4Distance = types.Int64Value(v.Int())
			data.Ipv4DistanceVariable = types.StringNull()
		}
	} else {
		data.Ipv4Distance = types.Int64Null()
		data.Ipv4DistanceVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.table-map.name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4PolicyName = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.table-map.name.vipVariableName")
			data.Ipv4PolicyNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4PolicyName = types.StringNull()
			data.Ipv4PolicyNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.table-map.name.vipValue")
			data.Ipv4PolicyName = types.StringValue(v.String())
			data.Ipv4PolicyNameVariable = types.StringNull()
		}
	} else {
		data.Ipv4PolicyName = types.StringNull()
		data.Ipv4PolicyNameVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.table-map.filter.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4Filter = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv4.table-map.filter.vipVariableName")
			data.Ipv4FilterVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4Filter = types.BoolNull()
			data.Ipv4FilterVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv4.table-map.filter.vipValue")
			data.Ipv4Filter = types.BoolValue(v.Bool())
			data.Ipv4FilterVariable = types.StringNull()
		}
	} else {
		data.Ipv4Filter = types.BoolNull()
		data.Ipv4FilterVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.redistribute.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Redistributes = make([]CiscoOSPFv3Ipv4Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv4Redistributes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			if cValue := v.Get("dia.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NatDia = types.BoolNull()

					cv := v.Get("dia.vipVariableName")
					item.NatDiaVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NatDia = types.BoolNull()
					item.NatDiaVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dia.vipValue")
					item.NatDia = types.BoolValue(cv.Bool())
					item.NatDiaVariable = types.StringNull()
				}
			} else {
				item.NatDia = types.BoolNull()
				item.NatDiaVariable = types.StringNull()
			}
			data.Ipv4Redistributes = append(data.Ipv4Redistributes, item)
			return true
		})
	} else {
		if len(data.Ipv4Redistributes) > 0 {
			data.Ipv4Redistributes = []CiscoOSPFv3Ipv4Redistributes{}
		}
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.max-metric.router-lsa.vipValue"); len(value.Array()) > 0 {
		data.Ipv4MaxMetricRouterLsas = make([]CiscoOSPFv3Ipv4MaxMetricRouterLsas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv4MaxMetricRouterLsas{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("ad-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("ad-type.vipValue")
					item.AdType = types.StringValue(cv.String())

				}
			} else {
				item.AdType = types.StringNull()

			}
			if cValue := v.Get("time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Time = types.Int64Null()

					cv := v.Get("time.vipVariableName")
					item.TimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Time = types.Int64Null()
					item.TimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("time.vipValue")
					item.Time = types.Int64Value(cv.Int())
					item.TimeVariable = types.StringNull()
				}
			} else {
				item.Time = types.Int64Null()
				item.TimeVariable = types.StringNull()
			}
			data.Ipv4MaxMetricRouterLsas = append(data.Ipv4MaxMetricRouterLsas, item)
			return true
		})
	} else {
		if len(data.Ipv4MaxMetricRouterLsas) > 0 {
			data.Ipv4MaxMetricRouterLsas = []CiscoOSPFv3Ipv4MaxMetricRouterLsas{}
		}
	}
	if value := res.Get(path + "ospfv3.address-family.ipv4.area.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Areas = make([]CiscoOSPFv3Ipv4Areas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv4Areas{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("a-num.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AreaNumber = types.Int64Null()

					cv := v.Get("a-num.vipVariableName")
					item.AreaNumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AreaNumber = types.Int64Null()
					item.AreaNumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("a-num.vipValue")
					item.AreaNumber = types.Int64Value(cv.Int())
					item.AreaNumberVariable = types.StringNull()
				}
			} else {
				item.AreaNumber = types.Int64Null()
				item.AreaNumberVariable = types.StringNull()
			}
			if cValue := v.Get("stub.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Stub = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Stub = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("stub.vipValue")
					item.Stub = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("stub"); cValue.Exists() {
				item.Stub = types.BoolValue(true)

			} else {
				item.Stub = types.BoolNull()

			}
			if cValue := v.Get("stub.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StubNoSummary = types.BoolNull()

					cv := v.Get("stub.no-summary.vipVariableName")
					item.StubNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StubNoSummary = types.BoolNull()
					item.StubNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("stub.no-summary.vipValue")
					item.StubNoSummary = types.BoolValue(cv.Bool())
					item.StubNoSummaryVariable = types.StringNull()
				}
			} else {
				item.StubNoSummary = types.BoolNull()
				item.StubNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("nssa.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Nssa = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Nssa = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.vipValue")
					item.Nssa = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("nssa"); cValue.Exists() {
				item.Nssa = types.BoolValue(true)

			} else {
				item.Nssa = types.BoolNull()

			}
			if cValue := v.Get("nssa.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NssaNoSummary = types.BoolNull()

					cv := v.Get("nssa.no-summary.vipVariableName")
					item.NssaNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NssaNoSummary = types.BoolNull()
					item.NssaNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.no-summary.vipValue")
					item.NssaNoSummary = types.BoolValue(cv.Bool())
					item.NssaNoSummaryVariable = types.StringNull()
				}
			} else {
				item.NssaNoSummary = types.BoolNull()
				item.NssaNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("nssa.translate.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Translate = types.StringNull()

					cv := v.Get("nssa.translate.vipVariableName")
					item.TranslateVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Translate = types.StringNull()
					item.TranslateVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.translate.vipValue")
					item.Translate = types.StringValue(cv.String())
					item.TranslateVariable = types.StringNull()
				}
			} else {
				item.Translate = types.StringNull()
				item.TranslateVariable = types.StringNull()
			}
			if cValue := v.Get("normal.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Normal = types.BoolNull()

					cv := v.Get("normal.vipVariableName")
					item.NormalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Normal = types.BoolNull()
					item.NormalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("normal.vipValue")
					item.Normal = types.BoolValue(cv.Bool())
					item.NormalVariable = types.StringNull()
				}
			} else {
				item.Normal = types.BoolNull()
				item.NormalVariable = types.StringNull()
			}
			if cValue := v.Get("interface.vipValue"); len(cValue.Array()) > 0 {
				item.Interfaces = make([]CiscoOSPFv3Ipv4AreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFv3Ipv4AreasInterfaces{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Name = types.StringNull()

							ccv := cv.Get("name.vipVariableName")
							cItem.NameVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Name = types.StringNull()
							cItem.NameVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.Name = types.StringValue(ccv.String())
							cItem.NameVariable = types.StringNull()
						}
					} else {
						cItem.Name = types.StringNull()
						cItem.NameVariable = types.StringNull()
					}
					if ccValue := cv.Get("hello-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.HelloInterval = types.Int64Null()

							ccv := cv.Get("hello-interval.vipVariableName")
							cItem.HelloIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.HelloInterval = types.Int64Null()
							cItem.HelloIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("hello-interval.vipValue")
							cItem.HelloInterval = types.Int64Value(ccv.Int())
							cItem.HelloIntervalVariable = types.StringNull()
						}
					} else {
						cItem.HelloInterval = types.Int64Null()
						cItem.HelloIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("dead-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.DeadInterval = types.Int64Null()

							ccv := cv.Get("dead-interval.vipVariableName")
							cItem.DeadIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.DeadInterval = types.Int64Null()
							cItem.DeadIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("dead-interval.vipValue")
							cItem.DeadInterval = types.Int64Value(ccv.Int())
							cItem.DeadIntervalVariable = types.StringNull()
						}
					} else {
						cItem.DeadInterval = types.Int64Null()
						cItem.DeadIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("retransmit-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RetransmitInterval = types.Int64Null()

							ccv := cv.Get("retransmit-interval.vipVariableName")
							cItem.RetransmitIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RetransmitInterval = types.Int64Null()
							cItem.RetransmitIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("retransmit-interval.vipValue")
							cItem.RetransmitInterval = types.Int64Value(ccv.Int())
							cItem.RetransmitIntervalVariable = types.StringNull()
						}
					} else {
						cItem.RetransmitInterval = types.Int64Null()
						cItem.RetransmitIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("network.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Network = types.StringNull()

							ccv := cv.Get("network.vipVariableName")
							cItem.NetworkVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Network = types.StringNull()
							cItem.NetworkVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("network.vipValue")
							cItem.Network = types.StringValue(ccv.String())
							cItem.NetworkVariable = types.StringNull()
						}
					} else {
						cItem.Network = types.StringNull()
						cItem.NetworkVariable = types.StringNull()
					}
					if ccValue := cv.Get("passive-interface.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.PassiveInterface = types.BoolNull()

							ccv := cv.Get("passive-interface.vipVariableName")
							cItem.PassiveInterfaceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.PassiveInterface = types.BoolNull()
							cItem.PassiveInterfaceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("passive-interface.vipValue")
							cItem.PassiveInterface = types.BoolValue(ccv.Bool())
							cItem.PassiveInterfaceVariable = types.StringNull()
						}
					} else {
						cItem.PassiveInterface = types.BoolNull()
						cItem.PassiveInterfaceVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationType = types.StringNull()

							ccv := cv.Get("authentication.type.vipVariableName")
							cItem.AuthenticationTypeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationType = types.StringNull()
							cItem.AuthenticationTypeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.type.vipValue")
							cItem.AuthenticationType = types.StringValue(ccv.String())
							cItem.AuthenticationTypeVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationType = types.StringNull()
						cItem.AuthenticationTypeVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.authentication-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationKey = types.StringNull()

							ccv := cv.Get("authentication.authentication-key.vipVariableName")
							cItem.AuthenticationKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationKey = types.StringNull()
							cItem.AuthenticationKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.authentication-key.vipValue")
							cItem.AuthenticationKey = types.StringValue(ccv.String())
							cItem.AuthenticationKeyVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationKey = types.StringNull()
						cItem.AuthenticationKeyVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.ipsec.spi.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.IpsecSpi = types.Int64Null()

							ccv := cv.Get("authentication.ipsec.spi.vipVariableName")
							cItem.IpsecSpiVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.IpsecSpi = types.Int64Null()
							cItem.IpsecSpiVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.ipsec.spi.vipValue")
							cItem.IpsecSpi = types.Int64Value(ccv.Int())
							cItem.IpsecSpiVariable = types.StringNull()
						}
					} else {
						cItem.IpsecSpi = types.Int64Null()
						cItem.IpsecSpiVariable = types.StringNull()
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			} else {
				if len(item.Interfaces) > 0 {
					item.Interfaces = []CiscoOSPFv3Ipv4AreasInterfaces{}
				}
			}
			if cValue := v.Get("range.vipValue"); len(cValue.Array()) > 0 {
				item.Ranges = make([]CiscoOSPFv3Ipv4AreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFv3Ipv4AreasRanges{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("no-advertise.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.NoAdvertise = types.BoolNull()

							ccv := cv.Get("no-advertise.vipVariableName")
							cItem.NoAdvertiseVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.NoAdvertise = types.BoolNull()
							cItem.NoAdvertiseVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("no-advertise.vipValue")
							cItem.NoAdvertise = types.BoolValue(ccv.Bool())
							cItem.NoAdvertiseVariable = types.StringNull()
						}
					} else {
						cItem.NoAdvertise = types.BoolNull()
						cItem.NoAdvertiseVariable = types.StringNull()
					}
					item.Ranges = append(item.Ranges, cItem)
					return true
				})
			} else {
				if len(item.Ranges) > 0 {
					item.Ranges = []CiscoOSPFv3Ipv4AreasRanges{}
				}
			}
			data.Ipv4Areas = append(data.Ipv4Areas, item)
			return true
		})
	} else {
		if len(data.Ipv4Areas) > 0 {
			data.Ipv4Areas = []CiscoOSPFv3Ipv4Areas{}
		}
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.router-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6RouterId = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.router-id.vipVariableName")
			data.Ipv6RouterIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6RouterId = types.StringNull()
			data.Ipv6RouterIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.router-id.vipValue")
			data.Ipv6RouterId = types.StringValue(v.String())
			data.Ipv6RouterIdVariable = types.StringNull()
		}
	} else {
		data.Ipv6RouterId = types.StringNull()
		data.Ipv6RouterIdVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.auto-cost.reference-bandwidth.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6AutoCostReferenceBandwidth = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.auto-cost.reference-bandwidth.vipVariableName")
			data.Ipv6AutoCostReferenceBandwidthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6AutoCostReferenceBandwidth = types.Int64Null()
			data.Ipv6AutoCostReferenceBandwidthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.auto-cost.reference-bandwidth.vipValue")
			data.Ipv6AutoCostReferenceBandwidth = types.Int64Value(v.Int())
			data.Ipv6AutoCostReferenceBandwidthVariable = types.StringNull()
		}
	} else {
		data.Ipv6AutoCostReferenceBandwidth = types.Int64Null()
		data.Ipv6AutoCostReferenceBandwidthVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.compatible.rfc1583.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6CompatibleRfc1583 = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.compatible.rfc1583.vipVariableName")
			data.Ipv6CompatibleRfc1583Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6CompatibleRfc1583 = types.BoolNull()
			data.Ipv6CompatibleRfc1583Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.compatible.rfc1583.vipValue")
			data.Ipv6CompatibleRfc1583 = types.BoolValue(v.Bool())
			data.Ipv6CompatibleRfc1583Variable = types.StringNull()
		}
	} else {
		data.Ipv6CompatibleRfc1583 = types.BoolNull()
		data.Ipv6CompatibleRfc1583Variable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DefaultInformationOriginate = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Ipv6DefaultInformationOriginate = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.vipValue")
			data.Ipv6DefaultInformationOriginate = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate"); value.Exists() {
		data.Ipv6DefaultInformationOriginate = types.BoolValue(true)

	} else {
		data.Ipv6DefaultInformationOriginate = types.BoolNull()

	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.always.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DefaultInformationOriginateAlways = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.always.vipVariableName")
			data.Ipv6DefaultInformationOriginateAlwaysVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DefaultInformationOriginateAlways = types.BoolNull()
			data.Ipv6DefaultInformationOriginateAlwaysVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.always.vipValue")
			data.Ipv6DefaultInformationOriginateAlways = types.BoolValue(v.Bool())
			data.Ipv6DefaultInformationOriginateAlwaysVariable = types.StringNull()
		}
	} else {
		data.Ipv6DefaultInformationOriginateAlways = types.BoolNull()
		data.Ipv6DefaultInformationOriginateAlwaysVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DefaultInformationOriginateMetric = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric.vipVariableName")
			data.Ipv6DefaultInformationOriginateMetricVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DefaultInformationOriginateMetric = types.Int64Null()
			data.Ipv6DefaultInformationOriginateMetricVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric.vipValue")
			data.Ipv6DefaultInformationOriginateMetric = types.Int64Value(v.Int())
			data.Ipv6DefaultInformationOriginateMetricVariable = types.StringNull()
		}
	} else {
		data.Ipv6DefaultInformationOriginateMetric = types.Int64Null()
		data.Ipv6DefaultInformationOriginateMetricVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DefaultInformationOriginateMetricType = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric-type.vipVariableName")
			data.Ipv6DefaultInformationOriginateMetricTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DefaultInformationOriginateMetricType = types.StringNull()
			data.Ipv6DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.default-information.originate.metric-type.vipValue")
			data.Ipv6DefaultInformationOriginateMetricType = types.StringValue(v.String())
			data.Ipv6DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		}
	} else {
		data.Ipv6DefaultInformationOriginateMetricType = types.StringNull()
		data.Ipv6DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.external.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DistanceExternal = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.external.vipVariableName")
			data.Ipv6DistanceExternalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DistanceExternal = types.Int64Null()
			data.Ipv6DistanceExternalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.external.vipValue")
			data.Ipv6DistanceExternal = types.Int64Value(v.Int())
			data.Ipv6DistanceExternalVariable = types.StringNull()
		}
	} else {
		data.Ipv6DistanceExternal = types.Int64Null()
		data.Ipv6DistanceExternalVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DistanceInterArea = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area.vipVariableName")
			data.Ipv6DistanceInterAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DistanceInterArea = types.Int64Null()
			data.Ipv6DistanceInterAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.inter-area.vipValue")
			data.Ipv6DistanceInterArea = types.Int64Value(v.Int())
			data.Ipv6DistanceInterAreaVariable = types.StringNull()
		}
	} else {
		data.Ipv6DistanceInterArea = types.Int64Null()
		data.Ipv6DistanceInterAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DistanceIntraArea = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area.vipVariableName")
			data.Ipv6DistanceIntraAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DistanceIntraArea = types.Int64Null()
			data.Ipv6DistanceIntraAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.ospf.intra-area.vipValue")
			data.Ipv6DistanceIntraArea = types.Int64Value(v.Int())
			data.Ipv6DistanceIntraAreaVariable = types.StringNull()
		}
	} else {
		data.Ipv6DistanceIntraArea = types.Int64Null()
		data.Ipv6DistanceIntraAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.delay.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6TimersSpfDelay = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.delay.vipVariableName")
			data.Ipv6TimersSpfDelayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6TimersSpfDelay = types.Int64Null()
			data.Ipv6TimersSpfDelayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.delay.vipValue")
			data.Ipv6TimersSpfDelay = types.Int64Value(v.Int())
			data.Ipv6TimersSpfDelayVariable = types.StringNull()
		}
	} else {
		data.Ipv6TimersSpfDelay = types.Int64Null()
		data.Ipv6TimersSpfDelayVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6TimersSpfInitialHold = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold.vipVariableName")
			data.Ipv6TimersSpfInitialHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6TimersSpfInitialHold = types.Int64Null()
			data.Ipv6TimersSpfInitialHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.initial-hold.vipValue")
			data.Ipv6TimersSpfInitialHold = types.Int64Value(v.Int())
			data.Ipv6TimersSpfInitialHoldVariable = types.StringNull()
		}
	} else {
		data.Ipv6TimersSpfInitialHold = types.Int64Null()
		data.Ipv6TimersSpfInitialHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.max-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6TimersSpfMaxHold = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.max-hold.vipVariableName")
			data.Ipv6TimersSpfMaxHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6TimersSpfMaxHold = types.Int64Null()
			data.Ipv6TimersSpfMaxHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.timers.throttle.spf.max-hold.vipValue")
			data.Ipv6TimersSpfMaxHold = types.Int64Value(v.Int())
			data.Ipv6TimersSpfMaxHoldVariable = types.StringNull()
		}
	} else {
		data.Ipv6TimersSpfMaxHold = types.Int64Null()
		data.Ipv6TimersSpfMaxHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.distance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Distance = types.Int64Null()

			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.distance.vipVariableName")
			data.Ipv6DistanceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Distance = types.Int64Null()
			data.Ipv6DistanceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.distance-ipv6.distance.vipValue")
			data.Ipv6Distance = types.Int64Value(v.Int())
			data.Ipv6DistanceVariable = types.StringNull()
		}
	} else {
		data.Ipv6Distance = types.Int64Null()
		data.Ipv6DistanceVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.table-map.name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6PolicyName = types.StringNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.table-map.name.vipVariableName")
			data.Ipv6PolicyNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6PolicyName = types.StringNull()
			data.Ipv6PolicyNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.table-map.name.vipValue")
			data.Ipv6PolicyName = types.StringValue(v.String())
			data.Ipv6PolicyNameVariable = types.StringNull()
		}
	} else {
		data.Ipv6PolicyName = types.StringNull()
		data.Ipv6PolicyNameVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.table-map.filter.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Filter = types.BoolNull()

			v := res.Get(path + "ospfv3.address-family.ipv6.table-map.filter.vipVariableName")
			data.Ipv6FilterVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Filter = types.BoolNull()
			data.Ipv6FilterVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospfv3.address-family.ipv6.table-map.filter.vipValue")
			data.Ipv6Filter = types.BoolValue(v.Bool())
			data.Ipv6FilterVariable = types.StringNull()
		}
	} else {
		data.Ipv6Filter = types.BoolNull()
		data.Ipv6FilterVariable = types.StringNull()
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.redistribute.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Redistributes = make([]CiscoOSPFv3Ipv6Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv6Redistributes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			data.Ipv6Redistributes = append(data.Ipv6Redistributes, item)
			return true
		})
	} else {
		if len(data.Ipv6Redistributes) > 0 {
			data.Ipv6Redistributes = []CiscoOSPFv3Ipv6Redistributes{}
		}
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.max-metric.router-lsa.vipValue"); len(value.Array()) > 0 {
		data.Ipv6MaxMetricRouterLsas = make([]CiscoOSPFv3Ipv6MaxMetricRouterLsas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv6MaxMetricRouterLsas{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("ad-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("ad-type.vipValue")
					item.AdType = types.StringValue(cv.String())

				}
			} else {
				item.AdType = types.StringNull()

			}
			if cValue := v.Get("time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Time = types.Int64Null()

					cv := v.Get("time.vipVariableName")
					item.TimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Time = types.Int64Null()
					item.TimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("time.vipValue")
					item.Time = types.Int64Value(cv.Int())
					item.TimeVariable = types.StringNull()
				}
			} else {
				item.Time = types.Int64Null()
				item.TimeVariable = types.StringNull()
			}
			data.Ipv6MaxMetricRouterLsas = append(data.Ipv6MaxMetricRouterLsas, item)
			return true
		})
	} else {
		if len(data.Ipv6MaxMetricRouterLsas) > 0 {
			data.Ipv6MaxMetricRouterLsas = []CiscoOSPFv3Ipv6MaxMetricRouterLsas{}
		}
	}
	if value := res.Get(path + "ospfv3.address-family.ipv6.area.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Areas = make([]CiscoOSPFv3Ipv6Areas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFv3Ipv6Areas{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("a-num.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AreaNumber = types.Int64Null()

					cv := v.Get("a-num.vipVariableName")
					item.AreaNumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AreaNumber = types.Int64Null()
					item.AreaNumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("a-num.vipValue")
					item.AreaNumber = types.Int64Value(cv.Int())
					item.AreaNumberVariable = types.StringNull()
				}
			} else {
				item.AreaNumber = types.Int64Null()
				item.AreaNumberVariable = types.StringNull()
			}
			if cValue := v.Get("stub.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Stub = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Stub = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("stub.vipValue")
					item.Stub = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("stub"); cValue.Exists() {
				item.Stub = types.BoolValue(true)

			} else {
				item.Stub = types.BoolNull()

			}
			if cValue := v.Get("stub.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StubNoSummary = types.BoolNull()

					cv := v.Get("stub.no-summary.vipVariableName")
					item.StubNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StubNoSummary = types.BoolNull()
					item.StubNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("stub.no-summary.vipValue")
					item.StubNoSummary = types.BoolValue(cv.Bool())
					item.StubNoSummaryVariable = types.StringNull()
				}
			} else {
				item.StubNoSummary = types.BoolNull()
				item.StubNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("nssa.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Nssa = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Nssa = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.vipValue")
					item.Nssa = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("nssa"); cValue.Exists() {
				item.Nssa = types.BoolValue(true)

			} else {
				item.Nssa = types.BoolNull()

			}
			if cValue := v.Get("nssa.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NssaNoSummary = types.BoolNull()

					cv := v.Get("nssa.no-summary.vipVariableName")
					item.NssaNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NssaNoSummary = types.BoolNull()
					item.NssaNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.no-summary.vipValue")
					item.NssaNoSummary = types.BoolValue(cv.Bool())
					item.NssaNoSummaryVariable = types.StringNull()
				}
			} else {
				item.NssaNoSummary = types.BoolNull()
				item.NssaNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("nssa.translate.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Translate = types.StringNull()

					cv := v.Get("nssa.translate.vipVariableName")
					item.TranslateVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Translate = types.StringNull()
					item.TranslateVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.translate.vipValue")
					item.Translate = types.StringValue(cv.String())
					item.TranslateVariable = types.StringNull()
				}
			} else {
				item.Translate = types.StringNull()
				item.TranslateVariable = types.StringNull()
			}
			if cValue := v.Get("normal.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Normal = types.BoolNull()

					cv := v.Get("normal.vipVariableName")
					item.NormalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Normal = types.BoolNull()
					item.NormalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("normal.vipValue")
					item.Normal = types.BoolValue(cv.Bool())
					item.NormalVariable = types.StringNull()
				}
			} else {
				item.Normal = types.BoolNull()
				item.NormalVariable = types.StringNull()
			}
			if cValue := v.Get("interface.vipValue"); len(cValue.Array()) > 0 {
				item.Interfaces = make([]CiscoOSPFv3Ipv6AreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFv3Ipv6AreasInterfaces{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Name = types.StringNull()

							ccv := cv.Get("name.vipVariableName")
							cItem.NameVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Name = types.StringNull()
							cItem.NameVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.Name = types.StringValue(ccv.String())
							cItem.NameVariable = types.StringNull()
						}
					} else {
						cItem.Name = types.StringNull()
						cItem.NameVariable = types.StringNull()
					}
					if ccValue := cv.Get("hello-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.HelloInterval = types.Int64Null()

							ccv := cv.Get("hello-interval.vipVariableName")
							cItem.HelloIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.HelloInterval = types.Int64Null()
							cItem.HelloIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("hello-interval.vipValue")
							cItem.HelloInterval = types.Int64Value(ccv.Int())
							cItem.HelloIntervalVariable = types.StringNull()
						}
					} else {
						cItem.HelloInterval = types.Int64Null()
						cItem.HelloIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("dead-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.DeadInterval = types.Int64Null()

							ccv := cv.Get("dead-interval.vipVariableName")
							cItem.DeadIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.DeadInterval = types.Int64Null()
							cItem.DeadIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("dead-interval.vipValue")
							cItem.DeadInterval = types.Int64Value(ccv.Int())
							cItem.DeadIntervalVariable = types.StringNull()
						}
					} else {
						cItem.DeadInterval = types.Int64Null()
						cItem.DeadIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("retransmit-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RetransmitInterval = types.Int64Null()

							ccv := cv.Get("retransmit-interval.vipVariableName")
							cItem.RetransmitIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RetransmitInterval = types.Int64Null()
							cItem.RetransmitIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("retransmit-interval.vipValue")
							cItem.RetransmitInterval = types.Int64Value(ccv.Int())
							cItem.RetransmitIntervalVariable = types.StringNull()
						}
					} else {
						cItem.RetransmitInterval = types.Int64Null()
						cItem.RetransmitIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("network.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Network = types.StringNull()

							ccv := cv.Get("network.vipVariableName")
							cItem.NetworkVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Network = types.StringNull()
							cItem.NetworkVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("network.vipValue")
							cItem.Network = types.StringValue(ccv.String())
							cItem.NetworkVariable = types.StringNull()
						}
					} else {
						cItem.Network = types.StringNull()
						cItem.NetworkVariable = types.StringNull()
					}
					if ccValue := cv.Get("passive-interface.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.PassiveInterface = types.BoolNull()

							ccv := cv.Get("passive-interface.vipVariableName")
							cItem.PassiveInterfaceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.PassiveInterface = types.BoolNull()
							cItem.PassiveInterfaceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("passive-interface.vipValue")
							cItem.PassiveInterface = types.BoolValue(ccv.Bool())
							cItem.PassiveInterfaceVariable = types.StringNull()
						}
					} else {
						cItem.PassiveInterface = types.BoolNull()
						cItem.PassiveInterfaceVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationType = types.StringNull()

							ccv := cv.Get("authentication.type.vipVariableName")
							cItem.AuthenticationTypeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationType = types.StringNull()
							cItem.AuthenticationTypeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.type.vipValue")
							cItem.AuthenticationType = types.StringValue(ccv.String())
							cItem.AuthenticationTypeVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationType = types.StringNull()
						cItem.AuthenticationTypeVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.authentication-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationKey = types.StringNull()

							ccv := cv.Get("authentication.authentication-key.vipVariableName")
							cItem.AuthenticationKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationKey = types.StringNull()
							cItem.AuthenticationKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.authentication-key.vipValue")
							cItem.AuthenticationKey = types.StringValue(ccv.String())
							cItem.AuthenticationKeyVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationKey = types.StringNull()
						cItem.AuthenticationKeyVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.ipsec.spi.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.IpsecSpi = types.Int64Null()

							ccv := cv.Get("authentication.ipsec.spi.vipVariableName")
							cItem.IpsecSpiVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.IpsecSpi = types.Int64Null()
							cItem.IpsecSpiVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.ipsec.spi.vipValue")
							cItem.IpsecSpi = types.Int64Value(ccv.Int())
							cItem.IpsecSpiVariable = types.StringNull()
						}
					} else {
						cItem.IpsecSpi = types.Int64Null()
						cItem.IpsecSpiVariable = types.StringNull()
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			} else {
				if len(item.Interfaces) > 0 {
					item.Interfaces = []CiscoOSPFv3Ipv6AreasInterfaces{}
				}
			}
			if cValue := v.Get("range.vipValue"); len(cValue.Array()) > 0 {
				item.Ranges = make([]CiscoOSPFv3Ipv6AreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFv3Ipv6AreasRanges{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("no-advertise.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.NoAdvertise = types.BoolNull()

							ccv := cv.Get("no-advertise.vipVariableName")
							cItem.NoAdvertiseVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.NoAdvertise = types.BoolNull()
							cItem.NoAdvertiseVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("no-advertise.vipValue")
							cItem.NoAdvertise = types.BoolValue(ccv.Bool())
							cItem.NoAdvertiseVariable = types.StringNull()
						}
					} else {
						cItem.NoAdvertise = types.BoolNull()
						cItem.NoAdvertiseVariable = types.StringNull()
					}
					item.Ranges = append(item.Ranges, cItem)
					return true
				})
			} else {
				if len(item.Ranges) > 0 {
					item.Ranges = []CiscoOSPFv3Ipv6AreasRanges{}
				}
			}
			data.Ipv6Areas = append(data.Ipv6Areas, item)
			return true
		})
	} else {
		if len(data.Ipv6Areas) > 0 {
			data.Ipv6Areas = []CiscoOSPFv3Ipv6Areas{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoOSPFv3) hasChanges(ctx context.Context, state *CiscoOSPFv3) bool {
	hasChanges := false
	if !data.Ipv4RouterId.Equal(state.Ipv4RouterId) {
		hasChanges = true
	}
	if !data.Ipv4AutoCostReferenceBandwidth.Equal(state.Ipv4AutoCostReferenceBandwidth) {
		hasChanges = true
	}
	if !data.Ipv4CompatibleRfc1583.Equal(state.Ipv4CompatibleRfc1583) {
		hasChanges = true
	}
	if !data.Ipv4DefaultInformationOriginate.Equal(state.Ipv4DefaultInformationOriginate) {
		hasChanges = true
	}
	if !data.Ipv4DefaultInformationOriginateAlways.Equal(state.Ipv4DefaultInformationOriginateAlways) {
		hasChanges = true
	}
	if !data.Ipv4DefaultInformationOriginateMetric.Equal(state.Ipv4DefaultInformationOriginateMetric) {
		hasChanges = true
	}
	if !data.Ipv4DefaultInformationOriginateMetricType.Equal(state.Ipv4DefaultInformationOriginateMetricType) {
		hasChanges = true
	}
	if !data.Ipv4DistanceExternal.Equal(state.Ipv4DistanceExternal) {
		hasChanges = true
	}
	if !data.Ipv4DistanceInterArea.Equal(state.Ipv4DistanceInterArea) {
		hasChanges = true
	}
	if !data.Ipv4DistanceIntraArea.Equal(state.Ipv4DistanceIntraArea) {
		hasChanges = true
	}
	if !data.Ipv4TimersSpfDelay.Equal(state.Ipv4TimersSpfDelay) {
		hasChanges = true
	}
	if !data.Ipv4TimersSpfInitialHold.Equal(state.Ipv4TimersSpfInitialHold) {
		hasChanges = true
	}
	if !data.Ipv4TimersSpfMaxHold.Equal(state.Ipv4TimersSpfMaxHold) {
		hasChanges = true
	}
	if !data.Ipv4Distance.Equal(state.Ipv4Distance) {
		hasChanges = true
	}
	if !data.Ipv4PolicyName.Equal(state.Ipv4PolicyName) {
		hasChanges = true
	}
	if !data.Ipv4Filter.Equal(state.Ipv4Filter) {
		hasChanges = true
	}
	if len(data.Ipv4Redistributes) != len(state.Ipv4Redistributes) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Redistributes {
			if !data.Ipv4Redistributes[i].Protocol.Equal(state.Ipv4Redistributes[i].Protocol) {
				hasChanges = true
			}
			if !data.Ipv4Redistributes[i].RoutePolicy.Equal(state.Ipv4Redistributes[i].RoutePolicy) {
				hasChanges = true
			}
			if !data.Ipv4Redistributes[i].NatDia.Equal(state.Ipv4Redistributes[i].NatDia) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4MaxMetricRouterLsas) != len(state.Ipv4MaxMetricRouterLsas) {
		hasChanges = true
	} else {
		for i := range data.Ipv4MaxMetricRouterLsas {
			if !data.Ipv4MaxMetricRouterLsas[i].AdType.Equal(state.Ipv4MaxMetricRouterLsas[i].AdType) {
				hasChanges = true
			}
			if !data.Ipv4MaxMetricRouterLsas[i].Time.Equal(state.Ipv4MaxMetricRouterLsas[i].Time) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4Areas) != len(state.Ipv4Areas) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Areas {
			if !data.Ipv4Areas[i].AreaNumber.Equal(state.Ipv4Areas[i].AreaNumber) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].Stub.Equal(state.Ipv4Areas[i].Stub) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].StubNoSummary.Equal(state.Ipv4Areas[i].StubNoSummary) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].Nssa.Equal(state.Ipv4Areas[i].Nssa) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].NssaNoSummary.Equal(state.Ipv4Areas[i].NssaNoSummary) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].Translate.Equal(state.Ipv4Areas[i].Translate) {
				hasChanges = true
			}
			if !data.Ipv4Areas[i].Normal.Equal(state.Ipv4Areas[i].Normal) {
				hasChanges = true
			}
			if len(data.Ipv4Areas[i].Interfaces) != len(state.Ipv4Areas[i].Interfaces) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Areas[i].Interfaces {
					if !data.Ipv4Areas[i].Interfaces[ii].Name.Equal(state.Ipv4Areas[i].Interfaces[ii].Name) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].HelloInterval.Equal(state.Ipv4Areas[i].Interfaces[ii].HelloInterval) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].DeadInterval.Equal(state.Ipv4Areas[i].Interfaces[ii].DeadInterval) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].RetransmitInterval.Equal(state.Ipv4Areas[i].Interfaces[ii].RetransmitInterval) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].Cost.Equal(state.Ipv4Areas[i].Interfaces[ii].Cost) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].Network.Equal(state.Ipv4Areas[i].Interfaces[ii].Network) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].PassiveInterface.Equal(state.Ipv4Areas[i].Interfaces[ii].PassiveInterface) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].AuthenticationType.Equal(state.Ipv4Areas[i].Interfaces[ii].AuthenticationType) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].AuthenticationKey.Equal(state.Ipv4Areas[i].Interfaces[ii].AuthenticationKey) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Interfaces[ii].IpsecSpi.Equal(state.Ipv4Areas[i].Interfaces[ii].IpsecSpi) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv4Areas[i].Ranges) != len(state.Ipv4Areas[i].Ranges) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Areas[i].Ranges {
					if !data.Ipv4Areas[i].Ranges[ii].Address.Equal(state.Ipv4Areas[i].Ranges[ii].Address) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Ranges[ii].Cost.Equal(state.Ipv4Areas[i].Ranges[ii].Cost) {
						hasChanges = true
					}
					if !data.Ipv4Areas[i].Ranges[ii].NoAdvertise.Equal(state.Ipv4Areas[i].Ranges[ii].NoAdvertise) {
						hasChanges = true
					}
				}
			}
		}
	}
	if !data.Ipv6RouterId.Equal(state.Ipv6RouterId) {
		hasChanges = true
	}
	if !data.Ipv6AutoCostReferenceBandwidth.Equal(state.Ipv6AutoCostReferenceBandwidth) {
		hasChanges = true
	}
	if !data.Ipv6CompatibleRfc1583.Equal(state.Ipv6CompatibleRfc1583) {
		hasChanges = true
	}
	if !data.Ipv6DefaultInformationOriginate.Equal(state.Ipv6DefaultInformationOriginate) {
		hasChanges = true
	}
	if !data.Ipv6DefaultInformationOriginateAlways.Equal(state.Ipv6DefaultInformationOriginateAlways) {
		hasChanges = true
	}
	if !data.Ipv6DefaultInformationOriginateMetric.Equal(state.Ipv6DefaultInformationOriginateMetric) {
		hasChanges = true
	}
	if !data.Ipv6DefaultInformationOriginateMetricType.Equal(state.Ipv6DefaultInformationOriginateMetricType) {
		hasChanges = true
	}
	if !data.Ipv6DistanceExternal.Equal(state.Ipv6DistanceExternal) {
		hasChanges = true
	}
	if !data.Ipv6DistanceInterArea.Equal(state.Ipv6DistanceInterArea) {
		hasChanges = true
	}
	if !data.Ipv6DistanceIntraArea.Equal(state.Ipv6DistanceIntraArea) {
		hasChanges = true
	}
	if !data.Ipv6TimersSpfDelay.Equal(state.Ipv6TimersSpfDelay) {
		hasChanges = true
	}
	if !data.Ipv6TimersSpfInitialHold.Equal(state.Ipv6TimersSpfInitialHold) {
		hasChanges = true
	}
	if !data.Ipv6TimersSpfMaxHold.Equal(state.Ipv6TimersSpfMaxHold) {
		hasChanges = true
	}
	if !data.Ipv6Distance.Equal(state.Ipv6Distance) {
		hasChanges = true
	}
	if !data.Ipv6PolicyName.Equal(state.Ipv6PolicyName) {
		hasChanges = true
	}
	if !data.Ipv6Filter.Equal(state.Ipv6Filter) {
		hasChanges = true
	}
	if len(data.Ipv6Redistributes) != len(state.Ipv6Redistributes) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Redistributes {
			if !data.Ipv6Redistributes[i].Protocol.Equal(state.Ipv6Redistributes[i].Protocol) {
				hasChanges = true
			}
			if !data.Ipv6Redistributes[i].RoutePolicy.Equal(state.Ipv6Redistributes[i].RoutePolicy) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv6MaxMetricRouterLsas) != len(state.Ipv6MaxMetricRouterLsas) {
		hasChanges = true
	} else {
		for i := range data.Ipv6MaxMetricRouterLsas {
			if !data.Ipv6MaxMetricRouterLsas[i].AdType.Equal(state.Ipv6MaxMetricRouterLsas[i].AdType) {
				hasChanges = true
			}
			if !data.Ipv6MaxMetricRouterLsas[i].Time.Equal(state.Ipv6MaxMetricRouterLsas[i].Time) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv6Areas) != len(state.Ipv6Areas) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Areas {
			if !data.Ipv6Areas[i].AreaNumber.Equal(state.Ipv6Areas[i].AreaNumber) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].Stub.Equal(state.Ipv6Areas[i].Stub) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].StubNoSummary.Equal(state.Ipv6Areas[i].StubNoSummary) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].Nssa.Equal(state.Ipv6Areas[i].Nssa) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].NssaNoSummary.Equal(state.Ipv6Areas[i].NssaNoSummary) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].Translate.Equal(state.Ipv6Areas[i].Translate) {
				hasChanges = true
			}
			if !data.Ipv6Areas[i].Normal.Equal(state.Ipv6Areas[i].Normal) {
				hasChanges = true
			}
			if len(data.Ipv6Areas[i].Interfaces) != len(state.Ipv6Areas[i].Interfaces) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Areas[i].Interfaces {
					if !data.Ipv6Areas[i].Interfaces[ii].Name.Equal(state.Ipv6Areas[i].Interfaces[ii].Name) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].HelloInterval.Equal(state.Ipv6Areas[i].Interfaces[ii].HelloInterval) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].DeadInterval.Equal(state.Ipv6Areas[i].Interfaces[ii].DeadInterval) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].RetransmitInterval.Equal(state.Ipv6Areas[i].Interfaces[ii].RetransmitInterval) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].Cost.Equal(state.Ipv6Areas[i].Interfaces[ii].Cost) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].Network.Equal(state.Ipv6Areas[i].Interfaces[ii].Network) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].PassiveInterface.Equal(state.Ipv6Areas[i].Interfaces[ii].PassiveInterface) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].AuthenticationType.Equal(state.Ipv6Areas[i].Interfaces[ii].AuthenticationType) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].AuthenticationKey.Equal(state.Ipv6Areas[i].Interfaces[ii].AuthenticationKey) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Interfaces[ii].IpsecSpi.Equal(state.Ipv6Areas[i].Interfaces[ii].IpsecSpi) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv6Areas[i].Ranges) != len(state.Ipv6Areas[i].Ranges) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Areas[i].Ranges {
					if !data.Ipv6Areas[i].Ranges[ii].Address.Equal(state.Ipv6Areas[i].Ranges[ii].Address) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Ranges[ii].Cost.Equal(state.Ipv6Areas[i].Ranges[ii].Cost) {
						hasChanges = true
					}
					if !data.Ipv6Areas[i].Ranges[ii].NoAdvertise.Equal(state.Ipv6Areas[i].Ranges[ii].NoAdvertise) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
