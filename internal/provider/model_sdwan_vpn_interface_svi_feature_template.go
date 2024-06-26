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
type VPNInterfaceSVI struct {
	Id                           types.String                            `tfsdk:"id"`
	Version                      types.Int64                             `tfsdk:"version"`
	TemplateType                 types.String                            `tfsdk:"template_type"`
	Name                         types.String                            `tfsdk:"name"`
	Description                  types.String                            `tfsdk:"description"`
	DeviceTypes                  types.Set                               `tfsdk:"device_types"`
	IfName                       types.String                            `tfsdk:"if_name"`
	IfNameVariable               types.String                            `tfsdk:"if_name_variable"`
	InterfaceDescription         types.String                            `tfsdk:"interface_description"`
	InterfaceDescriptionVariable types.String                            `tfsdk:"interface_description_variable"`
	Ipv4Address                  types.String                            `tfsdk:"ipv4_address"`
	Ipv4AddressVariable          types.String                            `tfsdk:"ipv4_address_variable"`
	Ipv4SecondaryAddresses       []VPNInterfaceSVIIpv4SecondaryAddresses `tfsdk:"ipv4_secondary_addresses"`
	Ipv6Address                  types.String                            `tfsdk:"ipv6_address"`
	Ipv6AddressVariable          types.String                            `tfsdk:"ipv6_address_variable"`
	Ipv6DhcpClient               types.Bool                              `tfsdk:"ipv6_dhcp_client"`
	Ipv6DhcpClientVariable       types.String                            `tfsdk:"ipv6_dhcp_client_variable"`
	Ipv6DhcpDistance             types.Int64                             `tfsdk:"ipv6_dhcp_distance"`
	Ipv6DhcpDistanceVariable     types.String                            `tfsdk:"ipv6_dhcp_distance_variable"`
	Ipv6DhcpRapidCommit          types.Bool                              `tfsdk:"ipv6_dhcp_rapid_commit"`
	Ipv6DhcpRapidCommitVariable  types.String                            `tfsdk:"ipv6_dhcp_rapid_commit_variable"`
	Ipv6SecondaryAddresses       []VPNInterfaceSVIIpv6SecondaryAddresses `tfsdk:"ipv6_secondary_addresses"`
	Ipv4DhcpHelper               types.Set                               `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable       types.String                            `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv6DhcpHelpers              []VPNInterfaceSVIIpv6DhcpHelpers        `tfsdk:"ipv6_dhcp_helpers"`
	IpDirectedBroadcast          types.Bool                              `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable  types.String                            `tfsdk:"ip_directed_broadcast_variable"`
	Mtu                          types.Int64                             `tfsdk:"mtu"`
	MtuVariable                  types.String                            `tfsdk:"mtu_variable"`
	IpMtu                        types.Int64                             `tfsdk:"ip_mtu"`
	IpMtuVariable                types.String                            `tfsdk:"ip_mtu_variable"`
	TcpMssAdjust                 types.Int64                             `tfsdk:"tcp_mss_adjust"`
	TcpMssAdjustVariable         types.String                            `tfsdk:"tcp_mss_adjust_variable"`
	Shutdown                     types.Bool                              `tfsdk:"shutdown"`
	ShutdownVariable             types.String                            `tfsdk:"shutdown_variable"`
	ArpTimeout                   types.Int64                             `tfsdk:"arp_timeout"`
	ArpTimeoutVariable           types.String                            `tfsdk:"arp_timeout_variable"`
	Ipv4AccessLists              []VPNInterfaceSVIIpv4AccessLists        `tfsdk:"ipv4_access_lists"`
	Ipv6AccessLists              []VPNInterfaceSVIIpv6AccessLists        `tfsdk:"ipv6_access_lists"`
	Policers                     []VPNInterfaceSVIPolicers               `tfsdk:"policers"`
	StaticArpEntries             []VPNInterfaceSVIStaticArpEntries       `tfsdk:"static_arp_entries"`
	Ipv4Vrrps                    []VPNInterfaceSVIIpv4Vrrps              `tfsdk:"ipv4_vrrps"`
	Ipv6Vrrps                    []VPNInterfaceSVIIpv6Vrrps              `tfsdk:"ipv6_vrrps"`
}

type VPNInterfaceSVIIpv4SecondaryAddresses struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Ipv4Address         types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable types.String `tfsdk:"ipv4_address_variable"`
}

type VPNInterfaceSVIIpv6SecondaryAddresses struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Ipv6Address         types.String `tfsdk:"ipv6_address"`
	Ipv6AddressVariable types.String `tfsdk:"ipv6_address_variable"`
}

type VPNInterfaceSVIIpv6DhcpHelpers struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	VpnId           types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable   types.String `tfsdk:"vpn_id_variable"`
}

type VPNInterfaceSVIIpv4AccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type VPNInterfaceSVIIpv6AccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type VPNInterfaceSVIPolicers struct {
	Optional    types.Bool   `tfsdk:"optional"`
	Direction   types.String `tfsdk:"direction"`
	PolicerName types.String `tfsdk:"policer_name"`
}

type VPNInterfaceSVIStaticArpEntries struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Ipv4Address         types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable types.String `tfsdk:"ipv4_address_variable"`
	MacAddress          types.String `tfsdk:"mac_address"`
	MacAddressVariable  types.String `tfsdk:"mac_address_variable"`
}

type VPNInterfaceSVIIpv4Vrrps struct {
	Optional                          types.Bool                                       `tfsdk:"optional"`
	GroupId                           types.Int64                                      `tfsdk:"group_id"`
	GroupIdVariable                   types.String                                     `tfsdk:"group_id_variable"`
	Priority                          types.Int64                                      `tfsdk:"priority"`
	PriorityVariable                  types.String                                     `tfsdk:"priority_variable"`
	Timer                             types.Int64                                      `tfsdk:"timer"`
	TimerVariable                     types.String                                     `tfsdk:"timer_variable"`
	TrackOmp                          types.Bool                                       `tfsdk:"track_omp"`
	TrackOmpVariable                  types.String                                     `tfsdk:"track_omp_variable"`
	TrackPrefixList                   types.String                                     `tfsdk:"track_prefix_list"`
	TrackPrefixListVariable           types.String                                     `tfsdk:"track_prefix_list_variable"`
	Ipv4Address                       types.String                                     `tfsdk:"ipv4_address"`
	Ipv4AddressVariable               types.String                                     `tfsdk:"ipv4_address_variable"`
	Ipv4SecondaryAddresses            []VPNInterfaceSVIIpv4VrrpsIpv4SecondaryAddresses `tfsdk:"ipv4_secondary_addresses"`
	TlocPreferenceChange              types.Bool                                       `tfsdk:"tloc_preference_change"`
	TlocPreferenceChangeValue         types.Int64                                      `tfsdk:"tloc_preference_change_value"`
	TlocPreferenceChangeValueVariable types.String                                     `tfsdk:"tloc_preference_change_value_variable"`
	TrackingObjects                   []VPNInterfaceSVIIpv4VrrpsTrackingObjects        `tfsdk:"tracking_objects"`
}

type VPNInterfaceSVIIpv6Vrrps struct {
	Optional                types.Bool                                       `tfsdk:"optional"`
	GroupId                 types.Int64                                      `tfsdk:"group_id"`
	GroupIdVariable         types.String                                     `tfsdk:"group_id_variable"`
	Priority                types.Int64                                      `tfsdk:"priority"`
	PriorityVariable        types.String                                     `tfsdk:"priority_variable"`
	Timer                   types.Int64                                      `tfsdk:"timer"`
	TimerVariable           types.String                                     `tfsdk:"timer_variable"`
	TrackOmp                types.Bool                                       `tfsdk:"track_omp"`
	TrackOmpVariable        types.String                                     `tfsdk:"track_omp_variable"`
	TrackPrefixList         types.String                                     `tfsdk:"track_prefix_list"`
	TrackPrefixListVariable types.String                                     `tfsdk:"track_prefix_list_variable"`
	Ipv6Addresses           []VPNInterfaceSVIIpv6VrrpsIpv6Addresses          `tfsdk:"ipv6_addresses"`
	Ipv6SecondaryAddresses  []VPNInterfaceSVIIpv6VrrpsIpv6SecondaryAddresses `tfsdk:"ipv6_secondary_addresses"`
}

type VPNInterfaceSVIIpv4VrrpsIpv4SecondaryAddresses struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Ipv4Address         types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable types.String `tfsdk:"ipv4_address_variable"`
}
type VPNInterfaceSVIIpv4VrrpsTrackingObjects struct {
	Optional               types.Bool   `tfsdk:"optional"`
	Name                   types.Int64  `tfsdk:"name"`
	NameVariable           types.String `tfsdk:"name_variable"`
	TrackAction            types.String `tfsdk:"track_action"`
	TrackActionVariable    types.String `tfsdk:"track_action_variable"`
	DecrementValue         types.Int64  `tfsdk:"decrement_value"`
	DecrementValueVariable types.String `tfsdk:"decrement_value_variable"`
}

type VPNInterfaceSVIIpv6VrrpsIpv6Addresses struct {
	Optional                 types.Bool   `tfsdk:"optional"`
	LinkLocalAddress         types.String `tfsdk:"link_local_address"`
	LinkLocalAddressVariable types.String `tfsdk:"link_local_address_variable"`
	Prefix                   types.String `tfsdk:"prefix"`
	PrefixVariable           types.String `tfsdk:"prefix_variable"`
}
type VPNInterfaceSVIIpv6VrrpsIpv6SecondaryAddresses struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data VPNInterfaceSVI) getModel() string {
	return "vpn-interface-svi"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data VPNInterfaceSVI) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "vpn-interface-svi")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.IfNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.IfNameVariable.ValueString())
	} else if data.IfName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.IfName.ValueString())
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.InterfaceDescription.ValueString())
	}

	if !data.Ipv4AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.address."+"vipVariableName", data.Ipv4AddressVariable.ValueString())
	} else if data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.address."+"vipValue", data.Ipv4Address.ValueString())
	}
	if len(data.Ipv4SecondaryAddresses) > 0 {
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4SecondaryAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.Ipv4AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.Ipv4AddressVariable.ValueString())
		} else if item.Ipv4Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Ipv4Address.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.secondary-address."+"vipValue.-1", itemBody)
	}

	if !data.Ipv6AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipVariableName", data.Ipv6AddressVariable.ValueString())
	} else if data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipValue", data.Ipv6Address.ValueString())
	}

	if !data.Ipv6DhcpClientVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipVariableName", data.Ipv6DhcpClientVariable.ValueString())
	} else if data.Ipv6DhcpClient.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipValue", strconv.FormatBool(data.Ipv6DhcpClient.ValueBool()))
	}

	if !data.Ipv6DhcpDistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipVariableName", data.Ipv6DhcpDistanceVariable.ValueString())
	} else if data.Ipv6DhcpDistance.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-distance."+"vipValue", data.Ipv6DhcpDistance.ValueInt64())
	}

	if !data.Ipv6DhcpRapidCommitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipVariableName", data.Ipv6DhcpRapidCommitVariable.ValueString())
	} else if data.Ipv6DhcpRapidCommit.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-rapid-commit."+"vipValue", strconv.FormatBool(data.Ipv6DhcpRapidCommit.ValueBool()))
	}
	if len(data.Ipv6SecondaryAddresses) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6SecondaryAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.Ipv6AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.Ipv6AddressVariable.ValueString())
		} else if item.Ipv6Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Ipv6Address.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.secondary-address."+"vipValue.-1", itemBody)
	}

	if !data.Ipv4DhcpHelperVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipVariableName", data.Ipv4DhcpHelperVariable.ValueString())
	} else if data.Ipv4DhcpHelper.IsNull() {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "constant")
		var values []string
		data.Ipv4DhcpHelper.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipValue", values)
	}
	if len(data.Ipv6DhcpHelpers) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6DhcpHelpers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.dhcp-helper-v6."+"vipValue.-1", itemBody)
	}

	if !data.IpDirectedBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipVariableName", data.IpDirectedBroadcastVariable.ValueString())
	} else if data.IpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipValue", strconv.FormatBool(data.IpDirectedBroadcast.ValueBool()))
	}

	if !data.MtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipVariableName", data.MtuVariable.ValueString())
	} else if data.Mtu.IsNull() {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipValue", data.Mtu.ValueInt64())
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mtu."+"vipVariableName", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mtu."+"vipValue", data.IpMtu.ValueInt64())
	}

	if !data.TcpMssAdjustVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipVariableName", data.TcpMssAdjustVariable.ValueString())
	} else if data.TcpMssAdjust.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipValue", data.TcpMssAdjust.ValueInt64())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.ArpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipVariableName", data.ArpTimeoutVariable.ValueString())
	} else if data.ArpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipValue", data.ArpTimeout.ValueInt64())
	}
	if len(data.Ipv4AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"access-list."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.access-list."+"vipValue.-1", itemBody)
	}
	if len(data.Policers) > 0 {
		body, _ = sjson.Set(body, path+"policer."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"policer."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"policer."+"vipPrimaryKey", []string{"policer-name", "direction"})
		body, _ = sjson.Set(body, path+"policer."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"policer."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"policer."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"policer."+"vipPrimaryKey", []string{"policer-name", "direction"})
		body, _ = sjson.Set(body, path+"policer."+"vipValue", []interface{}{})
	}
	for _, item := range data.Policers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "policer-name")
		if item.PolicerName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "policer-name."+"vipValue", item.PolicerName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"policer."+"vipValue.-1", itemBody)
	}
	if len(data.StaticArpEntries) > 0 {
		body, _ = sjson.Set(body, path+"arp.ip."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipPrimaryKey", []string{"addr"})
		body, _ = sjson.Set(body, path+"arp.ip."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"arp.ip."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipPrimaryKey", []string{"addr"})
		body, _ = sjson.Set(body, path+"arp.ip."+"vipValue", []interface{}{})
	}
	for _, item := range data.StaticArpEntries {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "addr")

		if !item.Ipv4AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipVariableName", item.Ipv4AddressVariable.ValueString())
		} else if item.Ipv4Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipValue", item.Ipv4Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "mac")

		if !item.MacAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipVariableName", item.MacAddressVariable.ValueString())
		} else if item.MacAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipValue", item.MacAddress.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"arp.ip."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4Vrrps) > 0 {
		body, _ = sjson.Set(body, path+"vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"vrrp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"vrrp."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"vrrp."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"vrrp."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Vrrps {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "grp-id")

		if !item.GroupIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipVariableName", item.GroupIdVariable.ValueString())
		} else if item.GroupId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipValue", item.GroupId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timer")

		if !item.TimerVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipVariableName", item.TimerVariable.ValueString())
		} else if item.Timer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipValue", item.Timer.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "track-omp")

		if !item.TrackOmpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipVariableName", item.TrackOmpVariable.ValueString())
		} else if item.TrackOmp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipValue", strconv.FormatBool(item.TrackOmp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "track-prefix-list")

		if !item.TrackPrefixListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipVariableName", item.TrackPrefixListVariable.ValueString())
		} else if item.TrackPrefixList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipValue", item.TrackPrefixList.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv4")

		if !item.Ipv4AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipVariableName", item.Ipv4AddressVariable.ValueString())
		} else if item.Ipv4Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipValue", item.Ipv4Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv4-secondary")
		if len(item.Ipv4SecondaryAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ipv4SecondaryAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.Ipv4AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.Ipv4AddressVariable.ValueString())
			} else if childItem.Ipv4Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Ipv4Address.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv4-secondary."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "tloc-change-pref")
		if item.TlocPreferenceChange.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipValue", strconv.FormatBool(item.TlocPreferenceChange.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "value")

		if !item.TlocPreferenceChangeValueVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "value."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipVariableName", item.TlocPreferenceChangeValueVariable.ValueString())
		} else if item.TlocPreferenceChangeValue.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "value."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipValue", item.TlocPreferenceChangeValue.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "tracking-object")
		if len(item.TrackingObjects) > 0 {
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.TrackingObjects {
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
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.Name.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "track-action")

			if !childItem.TrackActionVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipVariableName", childItem.TrackActionVariable.ValueString())
			} else if childItem.TrackAction.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipValue", childItem.TrackAction.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "decrement")

			if !childItem.DecrementValueVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipVariableName", childItem.DecrementValueVariable.ValueString())
			} else if childItem.DecrementValue.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipValue", childItem.DecrementValue.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "tracking-object."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"vrrp."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6Vrrps) > 0 {
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Vrrps {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "grp-id")

		if !item.GroupIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipVariableName", item.GroupIdVariable.ValueString())
		} else if item.GroupId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipValue", item.GroupId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timer")

		if !item.TimerVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipVariableName", item.TimerVariable.ValueString())
		} else if item.Timer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipValue", item.Timer.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "track-omp")

		if !item.TrackOmpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipVariableName", item.TrackOmpVariable.ValueString())
		} else if item.TrackOmp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipValue", strconv.FormatBool(item.TrackOmp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "track-prefix-list")

		if !item.TrackPrefixListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipVariableName", item.TrackPrefixListVariable.ValueString())
		} else if item.TrackPrefixList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipValue", item.TrackPrefixList.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv6")
		if len(item.Ipv6Addresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipPrimaryKey", []string{"ipv6-link-local"})
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipPrimaryKey", []string{"ipv6-link-local"})
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ipv6Addresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "ipv6-link-local")

			if !childItem.LinkLocalAddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipVariableName", childItem.LinkLocalAddressVariable.ValueString())
			} else if childItem.LinkLocalAddress.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipValue", childItem.LinkLocalAddress.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv6."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "ipv6-secondary")
		if len(item.Ipv6SecondaryAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "ipv6-secondary."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ipv6SecondaryAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv6-secondary."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6-vrrp."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *VPNInterfaceSVI) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IfName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.IfNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IfName = types.StringNull()
			data.IfNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.IfName = types.StringValue(v.String())
			data.IfNameVariable = types.StringNull()
		}
	} else {
		data.IfName = types.StringNull()
		data.IfNameVariable = types.StringNull()
	}
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.InterfaceDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceDescription = types.StringNull()
			data.InterfaceDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.InterfaceDescription = types.StringValue(v.String())
			data.InterfaceDescriptionVariable = types.StringNull()
		}
	} else {
		data.InterfaceDescription = types.StringNull()
		data.InterfaceDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv4Address = types.StringNull()

			v := res.Get(path + "ip.address.vipVariableName")
			data.Ipv4AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4Address = types.StringNull()
			data.Ipv4AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address.vipValue")
			data.Ipv4Address = types.StringValue(v.String())
			data.Ipv4AddressVariable = types.StringNull()
		}
	} else {
		data.Ipv4Address = types.StringNull()
		data.Ipv4AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.secondary-address.vipValue"); len(value.Array()) > 0 {
		data.Ipv4SecondaryAddresses = make([]VPNInterfaceSVIIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv4SecondaryAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv4Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.Ipv4AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv4Address = types.StringNull()
					item.Ipv4AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Ipv4Address = types.StringValue(cv.String())
					item.Ipv4AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv4Address = types.StringNull()
				item.Ipv4AddressVariable = types.StringNull()
			}
			data.Ipv4SecondaryAddresses = append(data.Ipv4SecondaryAddresses, item)
			return true
		})
	} else {
		if len(data.Ipv4SecondaryAddresses) > 0 {
			data.Ipv4SecondaryAddresses = []VPNInterfaceSVIIpv4SecondaryAddresses{}
		}
	}
	if value := res.Get(path + "ipv6.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Address = types.StringNull()

			v := res.Get(path + "ipv6.address.vipVariableName")
			data.Ipv6AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Address = types.StringNull()
			data.Ipv6AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.address.vipValue")
			data.Ipv6Address = types.StringValue(v.String())
			data.Ipv6AddressVariable = types.StringNull()
		}
	} else {
		data.Ipv6Address = types.StringNull()
		data.Ipv6AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-client.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DhcpClient = types.BoolNull()

			v := res.Get(path + "ipv6.dhcp-client.vipVariableName")
			data.Ipv6DhcpClientVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DhcpClient = types.BoolNull()
			data.Ipv6DhcpClientVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.dhcp-client.vipValue")
			data.Ipv6DhcpClient = types.BoolValue(v.Bool())
			data.Ipv6DhcpClientVariable = types.StringNull()
		}
	} else {
		data.Ipv6DhcpClient = types.BoolNull()
		data.Ipv6DhcpClientVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-distance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DhcpDistance = types.Int64Null()

			v := res.Get(path + "ipv6.dhcp-distance.vipVariableName")
			data.Ipv6DhcpDistanceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DhcpDistance = types.Int64Null()
			data.Ipv6DhcpDistanceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.dhcp-distance.vipValue")
			data.Ipv6DhcpDistance = types.Int64Value(v.Int())
			data.Ipv6DhcpDistanceVariable = types.StringNull()
		}
	} else {
		data.Ipv6DhcpDistance = types.Int64Null()
		data.Ipv6DhcpDistanceVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-rapid-commit.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6DhcpRapidCommit = types.BoolNull()

			v := res.Get(path + "ipv6.dhcp-rapid-commit.vipVariableName")
			data.Ipv6DhcpRapidCommitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6DhcpRapidCommit = types.BoolNull()
			data.Ipv6DhcpRapidCommitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.dhcp-rapid-commit.vipValue")
			data.Ipv6DhcpRapidCommit = types.BoolValue(v.Bool())
			data.Ipv6DhcpRapidCommitVariable = types.StringNull()
		}
	} else {
		data.Ipv6DhcpRapidCommit = types.BoolNull()
		data.Ipv6DhcpRapidCommitVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.secondary-address.vipValue"); len(value.Array()) > 0 {
		data.Ipv6SecondaryAddresses = make([]VPNInterfaceSVIIpv6SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv6SecondaryAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv6Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.Ipv6AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv6Address = types.StringNull()
					item.Ipv6AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Ipv6Address = types.StringValue(cv.String())
					item.Ipv6AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv6Address = types.StringNull()
				item.Ipv6AddressVariable = types.StringNull()
			}
			data.Ipv6SecondaryAddresses = append(data.Ipv6SecondaryAddresses, item)
			return true
		})
	} else {
		if len(data.Ipv6SecondaryAddresses) > 0 {
			data.Ipv6SecondaryAddresses = []VPNInterfaceSVIIpv6SecondaryAddresses{}
		}
	}
	if value := res.Get(path + "dhcp-helper.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.Ipv4DhcpHelper = types.SetNull(types.StringType)

			v := res.Get(path + "dhcp-helper.vipVariableName")
			data.Ipv4DhcpHelperVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DhcpHelper = types.SetNull(types.StringType)
			data.Ipv4DhcpHelperVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "dhcp-helper.vipValue")
			data.Ipv4DhcpHelper = helpers.GetStringSet(v.Array())
			data.Ipv4DhcpHelperVariable = types.StringNull()
		}
	} else {
		data.Ipv4DhcpHelper = types.SetNull(types.StringType)
		data.Ipv4DhcpHelperVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-helper-v6.vipValue"); len(value.Array()) > 0 {
		data.Ipv6DhcpHelpers = make([]VPNInterfaceSVIIpv6DhcpHelpers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv6DhcpHelpers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			data.Ipv6DhcpHelpers = append(data.Ipv6DhcpHelpers, item)
			return true
		})
	} else {
		if len(data.Ipv6DhcpHelpers) > 0 {
			data.Ipv6DhcpHelpers = []VPNInterfaceSVIIpv6DhcpHelpers{}
		}
	}
	if value := res.Get(path + "ip-directed-broadcast.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpDirectedBroadcast = types.BoolNull()

			v := res.Get(path + "ip-directed-broadcast.vipVariableName")
			data.IpDirectedBroadcastVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpDirectedBroadcast = types.BoolNull()
			data.IpDirectedBroadcastVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip-directed-broadcast.vipValue")
			data.IpDirectedBroadcast = types.BoolValue(v.Bool())
			data.IpDirectedBroadcastVariable = types.StringNull()
		}
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
		data.IpDirectedBroadcastVariable = types.StringNull()
	}
	if value := res.Get(path + "intrf-mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Mtu = types.Int64Null()

			v := res.Get(path + "intrf-mtu.vipVariableName")
			data.MtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Mtu = types.Int64Null()
			data.MtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "intrf-mtu.vipValue")
			data.Mtu = types.Int64Value(v.Int())
			data.MtuVariable = types.StringNull()
		}
	} else {
		data.Mtu = types.Int64Null()
		data.MtuVariable = types.StringNull()
	}
	if value := res.Get(path + "mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpMtu = types.Int64Null()

			v := res.Get(path + "mtu.vipVariableName")
			data.IpMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpMtu = types.Int64Null()
			data.IpMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mtu.vipValue")
			data.IpMtu = types.Int64Value(v.Int())
			data.IpMtuVariable = types.StringNull()
		}
	} else {
		data.IpMtu = types.Int64Null()
		data.IpMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpMssAdjust = types.Int64Null()

			v := res.Get(path + "tcp-mss-adjust.vipVariableName")
			data.TcpMssAdjustVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpMssAdjust = types.Int64Null()
			data.TcpMssAdjustVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tcp-mss-adjust.vipValue")
			data.TcpMssAdjust = types.Int64Value(v.Int())
			data.TcpMssAdjustVariable = types.StringNull()
		}
	} else {
		data.TcpMssAdjust = types.Int64Null()
		data.TcpMssAdjustVariable = types.StringNull()
	}
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "arp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ArpTimeout = types.Int64Null()

			v := res.Get(path + "arp-timeout.vipVariableName")
			data.ArpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ArpTimeout = types.Int64Null()
			data.ArpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "arp-timeout.vipValue")
			data.ArpTimeout = types.Int64Value(v.Int())
			data.ArpTimeoutVariable = types.StringNull()
		}
	} else {
		data.ArpTimeout = types.Int64Null()
		data.ArpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "access-list.vipValue"); len(value.Array()) > 0 {
		data.Ipv4AccessLists = make([]VPNInterfaceSVIIpv4AccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv4AccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.Ipv4AccessLists = append(data.Ipv4AccessLists, item)
			return true
		})
	} else {
		if len(data.Ipv4AccessLists) > 0 {
			data.Ipv4AccessLists = []VPNInterfaceSVIIpv4AccessLists{}
		}
	}
	if value := res.Get(path + "ipv6.access-list.vipValue"); len(value.Array()) > 0 {
		data.Ipv6AccessLists = make([]VPNInterfaceSVIIpv6AccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv6AccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.Ipv6AccessLists = append(data.Ipv6AccessLists, item)
			return true
		})
	} else {
		if len(data.Ipv6AccessLists) > 0 {
			data.Ipv6AccessLists = []VPNInterfaceSVIIpv6AccessLists{}
		}
	}
	if value := res.Get(path + "policer.vipValue"); len(value.Array()) > 0 {
		data.Policers = make([]VPNInterfaceSVIPolicers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIPolicers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("policer-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PolicerName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PolicerName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("policer-name.vipValue")
					item.PolicerName = types.StringValue(cv.String())

				}
			} else {
				item.PolicerName = types.StringNull()

			}
			data.Policers = append(data.Policers, item)
			return true
		})
	} else {
		if len(data.Policers) > 0 {
			data.Policers = []VPNInterfaceSVIPolicers{}
		}
	}
	if value := res.Get(path + "arp.ip.vipValue"); len(value.Array()) > 0 {
		data.StaticArpEntries = make([]VPNInterfaceSVIStaticArpEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIStaticArpEntries{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("addr.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv4Address = types.StringNull()

					cv := v.Get("addr.vipVariableName")
					item.Ipv4AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv4Address = types.StringNull()
					item.Ipv4AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("addr.vipValue")
					item.Ipv4Address = types.StringValue(cv.String())
					item.Ipv4AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv4Address = types.StringNull()
				item.Ipv4AddressVariable = types.StringNull()
			}
			if cValue := v.Get("mac.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MacAddress = types.StringNull()

					cv := v.Get("mac.vipVariableName")
					item.MacAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MacAddress = types.StringNull()
					item.MacAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("mac.vipValue")
					item.MacAddress = types.StringValue(cv.String())
					item.MacAddressVariable = types.StringNull()
				}
			} else {
				item.MacAddress = types.StringNull()
				item.MacAddressVariable = types.StringNull()
			}
			data.StaticArpEntries = append(data.StaticArpEntries, item)
			return true
		})
	} else {
		if len(data.StaticArpEntries) > 0 {
			data.StaticArpEntries = []VPNInterfaceSVIStaticArpEntries{}
		}
	}
	if value := res.Get(path + "vrrp.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Vrrps = make([]VPNInterfaceSVIIpv4Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv4Vrrps{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("grp-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupId = types.Int64Null()

					cv := v.Get("grp-id.vipVariableName")
					item.GroupIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.GroupId = types.Int64Null()
					item.GroupIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("grp-id.vipValue")
					item.GroupId = types.Int64Value(cv.Int())
					item.GroupIdVariable = types.StringNull()
				}
			} else {
				item.GroupId = types.Int64Null()
				item.GroupIdVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.Int64Null()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.Int64Null()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.Int64Value(cv.Int())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.Int64Null()
				item.PriorityVariable = types.StringNull()
			}
			if cValue := v.Get("timer.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Timer = types.Int64Null()

					cv := v.Get("timer.vipVariableName")
					item.TimerVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Timer = types.Int64Null()
					item.TimerVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timer.vipValue")
					item.Timer = types.Int64Value(cv.Int())
					item.TimerVariable = types.StringNull()
				}
			} else {
				item.Timer = types.Int64Null()
				item.TimerVariable = types.StringNull()
			}
			if cValue := v.Get("track-omp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackOmp = types.BoolNull()

					cv := v.Get("track-omp.vipVariableName")
					item.TrackOmpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackOmp = types.BoolNull()
					item.TrackOmpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-omp.vipValue")
					item.TrackOmp = types.BoolValue(cv.Bool())
					item.TrackOmpVariable = types.StringNull()
				}
			} else {
				item.TrackOmp = types.BoolNull()
				item.TrackOmpVariable = types.StringNull()
			}
			if cValue := v.Get("track-prefix-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackPrefixList = types.StringNull()

					cv := v.Get("track-prefix-list.vipVariableName")
					item.TrackPrefixListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackPrefixList = types.StringNull()
					item.TrackPrefixListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-prefix-list.vipValue")
					item.TrackPrefixList = types.StringValue(cv.String())
					item.TrackPrefixListVariable = types.StringNull()
				}
			} else {
				item.TrackPrefixList = types.StringNull()
				item.TrackPrefixListVariable = types.StringNull()
			}
			if cValue := v.Get("ipv4.address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv4Address = types.StringNull()

					cv := v.Get("ipv4.address.vipVariableName")
					item.Ipv4AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv4Address = types.StringNull()
					item.Ipv4AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipv4.address.vipValue")
					item.Ipv4Address = types.StringValue(cv.String())
					item.Ipv4AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv4Address = types.StringNull()
				item.Ipv4AddressVariable = types.StringNull()
			}
			if cValue := v.Get("ipv4-secondary.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv4SecondaryAddresses = make([]VPNInterfaceSVIIpv4VrrpsIpv4SecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := VPNInterfaceSVIIpv4VrrpsIpv4SecondaryAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Ipv4Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.Ipv4AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Ipv4Address = types.StringNull()
							cItem.Ipv4AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Ipv4Address = types.StringValue(ccv.String())
							cItem.Ipv4AddressVariable = types.StringNull()
						}
					} else {
						cItem.Ipv4Address = types.StringNull()
						cItem.Ipv4AddressVariable = types.StringNull()
					}
					item.Ipv4SecondaryAddresses = append(item.Ipv4SecondaryAddresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv4SecondaryAddresses) > 0 {
					item.Ipv4SecondaryAddresses = []VPNInterfaceSVIIpv4VrrpsIpv4SecondaryAddresses{}
				}
			}
			if cValue := v.Get("tloc-change-pref.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TlocPreferenceChange = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.TlocPreferenceChange = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tloc-change-pref.vipValue")
					item.TlocPreferenceChange = types.BoolValue(cv.Bool())

				}
			} else {
				item.TlocPreferenceChange = types.BoolNull()

			}
			if cValue := v.Get("value.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TlocPreferenceChangeValue = types.Int64Null()

					cv := v.Get("value.vipVariableName")
					item.TlocPreferenceChangeValueVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TlocPreferenceChangeValue = types.Int64Null()
					item.TlocPreferenceChangeValueVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("value.vipValue")
					item.TlocPreferenceChangeValue = types.Int64Value(cv.Int())
					item.TlocPreferenceChangeValueVariable = types.StringNull()
				}
			} else {
				item.TlocPreferenceChangeValue = types.Int64Null()
				item.TlocPreferenceChangeValueVariable = types.StringNull()
			}
			if cValue := v.Get("tracking-object.vipValue"); len(cValue.Array()) > 0 {
				item.TrackingObjects = make([]VPNInterfaceSVIIpv4VrrpsTrackingObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := VPNInterfaceSVIIpv4VrrpsTrackingObjects{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Name = types.Int64Null()

							ccv := cv.Get("name.vipVariableName")
							cItem.NameVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Name = types.Int64Null()
							cItem.NameVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.Name = types.Int64Value(ccv.Int())
							cItem.NameVariable = types.StringNull()
						}
					} else {
						cItem.Name = types.Int64Null()
						cItem.NameVariable = types.StringNull()
					}
					if ccValue := cv.Get("track-action.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.TrackAction = types.StringNull()

							ccv := cv.Get("track-action.vipVariableName")
							cItem.TrackActionVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.TrackAction = types.StringNull()
							cItem.TrackActionVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("track-action.vipValue")
							cItem.TrackAction = types.StringValue(ccv.String())
							cItem.TrackActionVariable = types.StringNull()
						}
					} else {
						cItem.TrackAction = types.StringNull()
						cItem.TrackActionVariable = types.StringNull()
					}
					if ccValue := cv.Get("decrement.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.DecrementValue = types.Int64Null()

							ccv := cv.Get("decrement.vipVariableName")
							cItem.DecrementValueVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.DecrementValue = types.Int64Null()
							cItem.DecrementValueVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("decrement.vipValue")
							cItem.DecrementValue = types.Int64Value(ccv.Int())
							cItem.DecrementValueVariable = types.StringNull()
						}
					} else {
						cItem.DecrementValue = types.Int64Null()
						cItem.DecrementValueVariable = types.StringNull()
					}
					item.TrackingObjects = append(item.TrackingObjects, cItem)
					return true
				})
			} else {
				if len(item.TrackingObjects) > 0 {
					item.TrackingObjects = []VPNInterfaceSVIIpv4VrrpsTrackingObjects{}
				}
			}
			data.Ipv4Vrrps = append(data.Ipv4Vrrps, item)
			return true
		})
	} else {
		if len(data.Ipv4Vrrps) > 0 {
			data.Ipv4Vrrps = []VPNInterfaceSVIIpv4Vrrps{}
		}
	}
	if value := res.Get(path + "ipv6-vrrp.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Vrrps = make([]VPNInterfaceSVIIpv6Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNInterfaceSVIIpv6Vrrps{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("grp-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupId = types.Int64Null()

					cv := v.Get("grp-id.vipVariableName")
					item.GroupIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.GroupId = types.Int64Null()
					item.GroupIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("grp-id.vipValue")
					item.GroupId = types.Int64Value(cv.Int())
					item.GroupIdVariable = types.StringNull()
				}
			} else {
				item.GroupId = types.Int64Null()
				item.GroupIdVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.Int64Null()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.Int64Null()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.Int64Value(cv.Int())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.Int64Null()
				item.PriorityVariable = types.StringNull()
			}
			if cValue := v.Get("timer.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Timer = types.Int64Null()

					cv := v.Get("timer.vipVariableName")
					item.TimerVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Timer = types.Int64Null()
					item.TimerVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timer.vipValue")
					item.Timer = types.Int64Value(cv.Int())
					item.TimerVariable = types.StringNull()
				}
			} else {
				item.Timer = types.Int64Null()
				item.TimerVariable = types.StringNull()
			}
			if cValue := v.Get("track-omp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackOmp = types.BoolNull()

					cv := v.Get("track-omp.vipVariableName")
					item.TrackOmpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackOmp = types.BoolNull()
					item.TrackOmpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-omp.vipValue")
					item.TrackOmp = types.BoolValue(cv.Bool())
					item.TrackOmpVariable = types.StringNull()
				}
			} else {
				item.TrackOmp = types.BoolNull()
				item.TrackOmpVariable = types.StringNull()
			}
			if cValue := v.Get("track-prefix-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackPrefixList = types.StringNull()

					cv := v.Get("track-prefix-list.vipVariableName")
					item.TrackPrefixListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackPrefixList = types.StringNull()
					item.TrackPrefixListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-prefix-list.vipValue")
					item.TrackPrefixList = types.StringValue(cv.String())
					item.TrackPrefixListVariable = types.StringNull()
				}
			} else {
				item.TrackPrefixList = types.StringNull()
				item.TrackPrefixListVariable = types.StringNull()
			}
			if cValue := v.Get("ipv6.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv6Addresses = make([]VPNInterfaceSVIIpv6VrrpsIpv6Addresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := VPNInterfaceSVIIpv6VrrpsIpv6Addresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("ipv6-link-local.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.LinkLocalAddress = types.StringNull()

							ccv := cv.Get("ipv6-link-local.vipVariableName")
							cItem.LinkLocalAddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.LinkLocalAddress = types.StringNull()
							cItem.LinkLocalAddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("ipv6-link-local.vipValue")
							cItem.LinkLocalAddress = types.StringValue(ccv.String())
							cItem.LinkLocalAddressVariable = types.StringNull()
						}
					} else {
						cItem.LinkLocalAddress = types.StringNull()
						cItem.LinkLocalAddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Ipv6Addresses = append(item.Ipv6Addresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv6Addresses) > 0 {
					item.Ipv6Addresses = []VPNInterfaceSVIIpv6VrrpsIpv6Addresses{}
				}
			}
			if cValue := v.Get("ipv6-secondary.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv6SecondaryAddresses = make([]VPNInterfaceSVIIpv6VrrpsIpv6SecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := VPNInterfaceSVIIpv6VrrpsIpv6SecondaryAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Ipv6SecondaryAddresses = append(item.Ipv6SecondaryAddresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv6SecondaryAddresses) > 0 {
					item.Ipv6SecondaryAddresses = []VPNInterfaceSVIIpv6VrrpsIpv6SecondaryAddresses{}
				}
			}
			data.Ipv6Vrrps = append(data.Ipv6Vrrps, item)
			return true
		})
	} else {
		if len(data.Ipv6Vrrps) > 0 {
			data.Ipv6Vrrps = []VPNInterfaceSVIIpv6Vrrps{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *VPNInterfaceSVI) hasChanges(ctx context.Context, state *VPNInterfaceSVI) bool {
	hasChanges := false
	if !data.IfName.Equal(state.IfName) {
		hasChanges = true
	}
	if !data.InterfaceDescription.Equal(state.InterfaceDescription) {
		hasChanges = true
	}
	if !data.Ipv4Address.Equal(state.Ipv4Address) {
		hasChanges = true
	}
	if len(data.Ipv4SecondaryAddresses) != len(state.Ipv4SecondaryAddresses) {
		hasChanges = true
	} else {
		for i := range data.Ipv4SecondaryAddresses {
			if !data.Ipv4SecondaryAddresses[i].Ipv4Address.Equal(state.Ipv4SecondaryAddresses[i].Ipv4Address) {
				hasChanges = true
			}
		}
	}
	if !data.Ipv6Address.Equal(state.Ipv6Address) {
		hasChanges = true
	}
	if !data.Ipv6DhcpClient.Equal(state.Ipv6DhcpClient) {
		hasChanges = true
	}
	if !data.Ipv6DhcpDistance.Equal(state.Ipv6DhcpDistance) {
		hasChanges = true
	}
	if !data.Ipv6DhcpRapidCommit.Equal(state.Ipv6DhcpRapidCommit) {
		hasChanges = true
	}
	if len(data.Ipv6SecondaryAddresses) != len(state.Ipv6SecondaryAddresses) {
		hasChanges = true
	} else {
		for i := range data.Ipv6SecondaryAddresses {
			if !data.Ipv6SecondaryAddresses[i].Ipv6Address.Equal(state.Ipv6SecondaryAddresses[i].Ipv6Address) {
				hasChanges = true
			}
		}
	}
	if !data.Ipv4DhcpHelper.Equal(state.Ipv4DhcpHelper) {
		hasChanges = true
	}
	if len(data.Ipv6DhcpHelpers) != len(state.Ipv6DhcpHelpers) {
		hasChanges = true
	} else {
		for i := range data.Ipv6DhcpHelpers {
			if !data.Ipv6DhcpHelpers[i].Address.Equal(state.Ipv6DhcpHelpers[i].Address) {
				hasChanges = true
			}
			if !data.Ipv6DhcpHelpers[i].VpnId.Equal(state.Ipv6DhcpHelpers[i].VpnId) {
				hasChanges = true
			}
		}
	}
	if !data.IpDirectedBroadcast.Equal(state.IpDirectedBroadcast) {
		hasChanges = true
	}
	if !data.Mtu.Equal(state.Mtu) {
		hasChanges = true
	}
	if !data.IpMtu.Equal(state.IpMtu) {
		hasChanges = true
	}
	if !data.TcpMssAdjust.Equal(state.TcpMssAdjust) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.ArpTimeout.Equal(state.ArpTimeout) {
		hasChanges = true
	}
	if len(data.Ipv4AccessLists) != len(state.Ipv4AccessLists) {
		hasChanges = true
	} else {
		for i := range data.Ipv4AccessLists {
			if !data.Ipv4AccessLists[i].Direction.Equal(state.Ipv4AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.Ipv4AccessLists[i].AclName.Equal(state.Ipv4AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv6AccessLists) != len(state.Ipv6AccessLists) {
		hasChanges = true
	} else {
		for i := range data.Ipv6AccessLists {
			if !data.Ipv6AccessLists[i].Direction.Equal(state.Ipv6AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.Ipv6AccessLists[i].AclName.Equal(state.Ipv6AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if len(data.Policers) != len(state.Policers) {
		hasChanges = true
	} else {
		for i := range data.Policers {
			if !data.Policers[i].Direction.Equal(state.Policers[i].Direction) {
				hasChanges = true
			}
			if !data.Policers[i].PolicerName.Equal(state.Policers[i].PolicerName) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticArpEntries) != len(state.StaticArpEntries) {
		hasChanges = true
	} else {
		for i := range data.StaticArpEntries {
			if !data.StaticArpEntries[i].Ipv4Address.Equal(state.StaticArpEntries[i].Ipv4Address) {
				hasChanges = true
			}
			if !data.StaticArpEntries[i].MacAddress.Equal(state.StaticArpEntries[i].MacAddress) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4Vrrps) != len(state.Ipv4Vrrps) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Vrrps {
			if !data.Ipv4Vrrps[i].GroupId.Equal(state.Ipv4Vrrps[i].GroupId) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].Priority.Equal(state.Ipv4Vrrps[i].Priority) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].Timer.Equal(state.Ipv4Vrrps[i].Timer) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TrackOmp.Equal(state.Ipv4Vrrps[i].TrackOmp) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TrackPrefixList.Equal(state.Ipv4Vrrps[i].TrackPrefixList) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].Ipv4Address.Equal(state.Ipv4Vrrps[i].Ipv4Address) {
				hasChanges = true
			}
			if len(data.Ipv4Vrrps[i].Ipv4SecondaryAddresses) != len(state.Ipv4Vrrps[i].Ipv4SecondaryAddresses) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Vrrps[i].Ipv4SecondaryAddresses {
					if !data.Ipv4Vrrps[i].Ipv4SecondaryAddresses[ii].Ipv4Address.Equal(state.Ipv4Vrrps[i].Ipv4SecondaryAddresses[ii].Ipv4Address) {
						hasChanges = true
					}
				}
			}
			if !data.Ipv4Vrrps[i].TlocPreferenceChange.Equal(state.Ipv4Vrrps[i].TlocPreferenceChange) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TlocPreferenceChangeValue.Equal(state.Ipv4Vrrps[i].TlocPreferenceChangeValue) {
				hasChanges = true
			}
			if len(data.Ipv4Vrrps[i].TrackingObjects) != len(state.Ipv4Vrrps[i].TrackingObjects) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Vrrps[i].TrackingObjects {
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].Name.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].Name) {
						hasChanges = true
					}
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].TrackAction.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].TrackAction) {
						hasChanges = true
					}
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].DecrementValue.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].DecrementValue) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv6Vrrps) != len(state.Ipv6Vrrps) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Vrrps {
			if !data.Ipv6Vrrps[i].GroupId.Equal(state.Ipv6Vrrps[i].GroupId) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].Priority.Equal(state.Ipv6Vrrps[i].Priority) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].Timer.Equal(state.Ipv6Vrrps[i].Timer) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].TrackOmp.Equal(state.Ipv6Vrrps[i].TrackOmp) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].TrackPrefixList.Equal(state.Ipv6Vrrps[i].TrackPrefixList) {
				hasChanges = true
			}
			if len(data.Ipv6Vrrps[i].Ipv6Addresses) != len(state.Ipv6Vrrps[i].Ipv6Addresses) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Vrrps[i].Ipv6Addresses {
					if !data.Ipv6Vrrps[i].Ipv6Addresses[ii].LinkLocalAddress.Equal(state.Ipv6Vrrps[i].Ipv6Addresses[ii].LinkLocalAddress) {
						hasChanges = true
					}
					if !data.Ipv6Vrrps[i].Ipv6Addresses[ii].Prefix.Equal(state.Ipv6Vrrps[i].Ipv6Addresses[ii].Prefix) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv6Vrrps[i].Ipv6SecondaryAddresses) != len(state.Ipv6Vrrps[i].Ipv6SecondaryAddresses) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Vrrps[i].Ipv6SecondaryAddresses {
					if !data.Ipv6Vrrps[i].Ipv6SecondaryAddresses[ii].Prefix.Equal(state.Ipv6Vrrps[i].Ipv6SecondaryAddresses[ii].Prefix) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
