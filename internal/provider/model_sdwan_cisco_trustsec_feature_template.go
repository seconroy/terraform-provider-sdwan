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
type CiscoTrustSec struct {
	Id                              types.String                  `tfsdk:"id"`
	Version                         types.Int64                   `tfsdk:"version"`
	TemplateType                    types.String                  `tfsdk:"template_type"`
	Name                            types.String                  `tfsdk:"name"`
	Description                     types.String                  `tfsdk:"description"`
	DeviceTypes                     types.Set                     `tfsdk:"device_types"`
	DeviceSgt                       types.Int64                   `tfsdk:"device_sgt"`
	DeviceSgtVariable               types.String                  `tfsdk:"device_sgt_variable"`
	CredentialsId                   types.String                  `tfsdk:"credentials_id"`
	CredentialsIdVariable           types.String                  `tfsdk:"credentials_id_variable"`
	CredentialsPassword             types.String                  `tfsdk:"credentials_password"`
	CredentialsPasswordVariable     types.String                  `tfsdk:"credentials_password_variable"`
	EnableEnforcement               types.Bool                    `tfsdk:"enable_enforcement"`
	EnableEnforcementVariable       types.String                  `tfsdk:"enable_enforcement_variable"`
	EnableSxp                       types.Bool                    `tfsdk:"enable_sxp"`
	SxpSourceIp                     types.String                  `tfsdk:"sxp_source_ip"`
	SxpSourceIpVariable             types.String                  `tfsdk:"sxp_source_ip_variable"`
	SxpDefaultPassword              types.String                  `tfsdk:"sxp_default_password"`
	SxpDefaultPasswordVariable      types.String                  `tfsdk:"sxp_default_password_variable"`
	SxpKeyChain                     types.String                  `tfsdk:"sxp_key_chain"`
	SxpKeyChainVariable             types.String                  `tfsdk:"sxp_key_chain_variable"`
	SxpLogBindingChanges            types.Bool                    `tfsdk:"sxp_log_binding_changes"`
	SxpLogBindingChangesVariable    types.String                  `tfsdk:"sxp_log_binding_changes_variable"`
	SxpReconciliationPeriod         types.Int64                   `tfsdk:"sxp_reconciliation_period"`
	SxpReconciliationPeriodVariable types.String                  `tfsdk:"sxp_reconciliation_period_variable"`
	SxpRetryPeriod                  types.Int64                   `tfsdk:"sxp_retry_period"`
	SxpRetryPeriodVariable          types.String                  `tfsdk:"sxp_retry_period_variable"`
	SpeakerHoldTime                 types.Int64                   `tfsdk:"speaker_hold_time"`
	SpeakerHoldTimeVariable         types.String                  `tfsdk:"speaker_hold_time_variable"`
	MinimumListenerHoldTime         types.Int64                   `tfsdk:"minimum_listener_hold_time"`
	MinimumListenerHoldTimeVariable types.String                  `tfsdk:"minimum_listener_hold_time_variable"`
	MaximumListenerHoldTime         types.Int64                   `tfsdk:"maximum_listener_hold_time"`
	MaximumListenerHoldTimeVariable types.String                  `tfsdk:"maximum_listener_hold_time_variable"`
	SxpNodeIdType                   types.String                  `tfsdk:"sxp_node_id_type"`
	SxpNodeId                       types.String                  `tfsdk:"sxp_node_id"`
	SxpNodeIdVariable               types.String                  `tfsdk:"sxp_node_id_variable"`
	SxpConnections                  []CiscoTrustSecSxpConnections `tfsdk:"sxp_connections"`
}

type CiscoTrustSecSxpConnections struct {
	Optional                types.Bool   `tfsdk:"optional"`
	PeerIp                  types.String `tfsdk:"peer_ip"`
	SourceIp                types.String `tfsdk:"source_ip"`
	SourceIpVariable        types.String `tfsdk:"source_ip_variable"`
	PresharedKey            types.String `tfsdk:"preshared_key"`
	Mode                    types.String `tfsdk:"mode"`
	ModeType                types.String `tfsdk:"mode_type"`
	MinimumHoldTime         types.Int64  `tfsdk:"minimum_hold_time"`
	MinimumHoldTimeVariable types.String `tfsdk:"minimum_hold_time_variable"`
	MaximumHoldTime         types.Int64  `tfsdk:"maximum_hold_time"`
	MaximumHoldTimeVariable types.String `tfsdk:"maximum_hold_time_variable"`
	VpnId                   types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable           types.String `tfsdk:"vpn_id_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoTrustSec) getModel() string {
	return "cisco_trustsec"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoTrustSec) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_trustsec")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.DeviceSgtVariable.IsNull() {
		body, _ = sjson.Set(body, path+"device-sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"device-sgt."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"device-sgt."+"vipVariableName", data.DeviceSgtVariable.ValueString())
	} else if data.DeviceSgt.IsNull() {
		body, _ = sjson.Set(body, path+"device-sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"device-sgt."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"device-sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"device-sgt."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"device-sgt."+"vipValue", data.DeviceSgt.ValueInt64())
	}

	if !data.CredentialsIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"credentials-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"credentials-id."+"vipVariableName", data.CredentialsIdVariable.ValueString())
	} else if data.CredentialsId.IsNull() {
		body, _ = sjson.Set(body, path+"credentials-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"credentials-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"credentials-id."+"vipValue", data.CredentialsId.ValueString())
	}

	if !data.CredentialsPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"credentials-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"credentials-password."+"vipVariableName", data.CredentialsPasswordVariable.ValueString())
	} else if data.CredentialsPassword.IsNull() {
		body, _ = sjson.Set(body, path+"credentials-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-password."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"credentials-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"credentials-password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"credentials-password."+"vipValue", data.CredentialsPassword.ValueString())
	}

	if !data.EnableEnforcementVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipVariableName", data.EnableEnforcementVariable.ValueString())
	} else if data.EnableEnforcement.IsNull() {
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"enable-enforcement."+"vipValue", strconv.FormatBool(data.EnableEnforcement.ValueBool()))
	}
	if data.EnableSxp.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.enable-sxp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.enable-sxp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.enable-sxp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.enable-sxp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.enable-sxp."+"vipValue", strconv.FormatBool(data.EnableSxp.ValueBool()))
	}

	if !data.SxpSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipVariableName", data.SxpSourceIpVariable.ValueString())
	} else if data.SxpSourceIp.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-source-ip."+"vipValue", data.SxpSourceIp.ValueString())
	}

	if !data.SxpDefaultPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipVariableName", data.SxpDefaultPasswordVariable.ValueString())
	} else if data.SxpDefaultPassword.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-default-password."+"vipValue", data.SxpDefaultPassword.ValueString())
	}

	if !data.SxpKeyChainVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipVariableName", data.SxpKeyChainVariable.ValueString())
	} else if data.SxpKeyChain.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-key-chain."+"vipValue", data.SxpKeyChain.ValueString())
	}

	if !data.SxpLogBindingChangesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipVariableName", data.SxpLogBindingChangesVariable.ValueString())
	} else if data.SxpLogBindingChanges.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-log-binding-changes."+"vipValue", strconv.FormatBool(data.SxpLogBindingChanges.ValueBool()))
	}

	if !data.SxpReconciliationPeriodVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipVariableName", data.SxpReconciliationPeriodVariable.ValueString())
	} else if data.SxpReconciliationPeriod.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-reconciliation-period."+"vipValue", data.SxpReconciliationPeriod.ValueInt64())
	}

	if !data.SxpRetryPeriodVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipVariableName", data.SxpRetryPeriodVariable.ValueString())
	} else if data.SxpRetryPeriod.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-retry-period."+"vipValue", data.SxpRetryPeriod.ValueInt64())
	}

	if !data.SpeakerHoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipVariableName", data.SpeakerHoldTimeVariable.ValueString())
	} else if data.SpeakerHoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.speaker-hold-time."+"vipValue", data.SpeakerHoldTime.ValueInt64())
	}

	if !data.MinimumListenerHoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipVariableName", data.MinimumListenerHoldTimeVariable.ValueString())
	} else if data.MinimumListenerHoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-min."+"vipValue", data.MinimumListenerHoldTime.ValueInt64())
	}

	if !data.MaximumListenerHoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipVariableName", data.MaximumListenerHoldTimeVariable.ValueString())
	} else if data.MaximumListenerHoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.listener-hold-time-max."+"vipValue", data.MaximumListenerHoldTime.ValueInt64())
	}
	if data.SxpNodeIdType.IsNull() {
		if !gjson.Get(body, path+"sxp-default").Exists() {
			body, _ = sjson.Set(body, path+"sxp-default", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id-type."+"vipValue", data.SxpNodeIdType.ValueString())
	}

	if !data.SxpNodeIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipVariableName", data.SxpNodeIdVariable.ValueString())
	} else if data.SxpNodeId.IsNull() {
		if !gjson.Get(body, path+"sxp-default").Exists() {
			body, _ = sjson.Set(body, path+"sxp-default", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-default.sxp-node-id."+"vipValue", data.SxpNodeId.ValueString())
	}
	if len(data.SxpConnections) > 0 {
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipPrimaryKey", []string{"connection-peer-ip"})
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipPrimaryKey", []string{"connection-peer-ip"})
		body, _ = sjson.Set(body, path+"sxp-connection-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.SxpConnections {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "connection-peer-ip")
		if item.PeerIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-peer-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-peer-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-peer-ip."+"vipValue", item.PeerIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "connection-source-ip")

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipVariableName", item.SourceIpVariable.ValueString())
		} else if item.SourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-source-ip."+"vipValue", item.SourceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "connection-preshared-key")
		if item.PresharedKey.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-preshared-key."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-preshared-key."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-preshared-key."+"vipValue", item.PresharedKey.ValueString())
		}
		itemAttributes = append(itemAttributes, "connection-mode")
		if item.Mode.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-mode."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-mode."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-mode."+"vipValue", item.Mode.ValueString())
		}
		itemAttributes = append(itemAttributes, "connection-mode-type")
		if item.ModeType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-mode-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-mode-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-mode-type."+"vipValue", item.ModeType.ValueString())
		}
		itemAttributes = append(itemAttributes, "connection-min-hold-time")

		if !item.MinimumHoldTimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipVariableName", item.MinimumHoldTimeVariable.ValueString())
		} else if item.MinimumHoldTime.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-min-hold-time."+"vipValue", item.MinimumHoldTime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "connection-max-hold-time")

		if !item.MaximumHoldTimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipVariableName", item.MaximumHoldTimeVariable.ValueString())
		} else if item.MaximumHoldTime.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-max-hold-time."+"vipValue", item.MaximumHoldTime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "connection-vpn-id")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "connection-vpn-id."+"vipValue", item.VpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"sxp-connection-list."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoTrustSec) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "device-sgt.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DeviceSgt = types.Int64Null()

			v := res.Get(path + "device-sgt.vipVariableName")
			data.DeviceSgtVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DeviceSgt = types.Int64Null()
			data.DeviceSgtVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "device-sgt.vipValue")
			data.DeviceSgt = types.Int64Value(v.Int())
			data.DeviceSgtVariable = types.StringNull()
		}
	} else {
		data.DeviceSgt = types.Int64Null()
		data.DeviceSgtVariable = types.StringNull()
	}
	if value := res.Get(path + "credentials-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CredentialsId = types.StringNull()

			v := res.Get(path + "credentials-id.vipVariableName")
			data.CredentialsIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CredentialsId = types.StringNull()
			data.CredentialsIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "credentials-id.vipValue")
			data.CredentialsId = types.StringValue(v.String())
			data.CredentialsIdVariable = types.StringNull()
		}
	} else {
		data.CredentialsId = types.StringNull()
		data.CredentialsIdVariable = types.StringNull()
	}
	if value := res.Get(path + "credentials-password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CredentialsPassword = types.StringNull()

			v := res.Get(path + "credentials-password.vipVariableName")
			data.CredentialsPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CredentialsPassword = types.StringNull()
			data.CredentialsPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "credentials-password.vipValue")
			data.CredentialsPassword = types.StringValue(v.String())
			data.CredentialsPasswordVariable = types.StringNull()
		}
	} else {
		data.CredentialsPassword = types.StringNull()
		data.CredentialsPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "enable-enforcement.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableEnforcement = types.BoolNull()

			v := res.Get(path + "enable-enforcement.vipVariableName")
			data.EnableEnforcementVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EnableEnforcement = types.BoolNull()
			data.EnableEnforcementVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "enable-enforcement.vipValue")
			data.EnableEnforcement = types.BoolValue(v.Bool())
			data.EnableEnforcementVariable = types.StringNull()
		}
	} else {
		data.EnableEnforcement = types.BoolNull()
		data.EnableEnforcementVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.enable-sxp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableSxp = types.BoolNull()

		} else if value.String() == "ignore" {
			data.EnableSxp = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.enable-sxp.vipValue")
			data.EnableSxp = types.BoolValue(v.Bool())

		}
	} else {
		data.EnableSxp = types.BoolNull()

	}
	if value := res.Get(path + "sxp-default.sxp-source-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpSourceIp = types.StringNull()

			v := res.Get(path + "sxp-default.sxp-source-ip.vipVariableName")
			data.SxpSourceIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpSourceIp = types.StringNull()
			data.SxpSourceIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-source-ip.vipValue")
			data.SxpSourceIp = types.StringValue(v.String())
			data.SxpSourceIpVariable = types.StringNull()
		}
	} else {
		data.SxpSourceIp = types.StringNull()
		data.SxpSourceIpVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-default-password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpDefaultPassword = types.StringNull()

			v := res.Get(path + "sxp-default.sxp-default-password.vipVariableName")
			data.SxpDefaultPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpDefaultPassword = types.StringNull()
			data.SxpDefaultPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-default-password.vipValue")
			data.SxpDefaultPassword = types.StringValue(v.String())
			data.SxpDefaultPasswordVariable = types.StringNull()
		}
	} else {
		data.SxpDefaultPassword = types.StringNull()
		data.SxpDefaultPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-key-chain.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpKeyChain = types.StringNull()

			v := res.Get(path + "sxp-default.sxp-key-chain.vipVariableName")
			data.SxpKeyChainVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpKeyChain = types.StringNull()
			data.SxpKeyChainVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-key-chain.vipValue")
			data.SxpKeyChain = types.StringValue(v.String())
			data.SxpKeyChainVariable = types.StringNull()
		}
	} else {
		data.SxpKeyChain = types.StringNull()
		data.SxpKeyChainVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-log-binding-changes.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpLogBindingChanges = types.BoolNull()

			v := res.Get(path + "sxp-default.sxp-log-binding-changes.vipVariableName")
			data.SxpLogBindingChangesVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpLogBindingChanges = types.BoolNull()
			data.SxpLogBindingChangesVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-log-binding-changes.vipValue")
			data.SxpLogBindingChanges = types.BoolValue(v.Bool())
			data.SxpLogBindingChangesVariable = types.StringNull()
		}
	} else {
		data.SxpLogBindingChanges = types.BoolNull()
		data.SxpLogBindingChangesVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-reconciliation-period.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpReconciliationPeriod = types.Int64Null()

			v := res.Get(path + "sxp-default.sxp-reconciliation-period.vipVariableName")
			data.SxpReconciliationPeriodVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpReconciliationPeriod = types.Int64Null()
			data.SxpReconciliationPeriodVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-reconciliation-period.vipValue")
			data.SxpReconciliationPeriod = types.Int64Value(v.Int())
			data.SxpReconciliationPeriodVariable = types.StringNull()
		}
	} else {
		data.SxpReconciliationPeriod = types.Int64Null()
		data.SxpReconciliationPeriodVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-retry-period.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpRetryPeriod = types.Int64Null()

			v := res.Get(path + "sxp-default.sxp-retry-period.vipVariableName")
			data.SxpRetryPeriodVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpRetryPeriod = types.Int64Null()
			data.SxpRetryPeriodVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-retry-period.vipValue")
			data.SxpRetryPeriod = types.Int64Value(v.Int())
			data.SxpRetryPeriodVariable = types.StringNull()
		}
	} else {
		data.SxpRetryPeriod = types.Int64Null()
		data.SxpRetryPeriodVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.speaker-hold-time.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SpeakerHoldTime = types.Int64Null()

			v := res.Get(path + "sxp-default.speaker-hold-time.vipVariableName")
			data.SpeakerHoldTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SpeakerHoldTime = types.Int64Null()
			data.SpeakerHoldTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.speaker-hold-time.vipValue")
			data.SpeakerHoldTime = types.Int64Value(v.Int())
			data.SpeakerHoldTimeVariable = types.StringNull()
		}
	} else {
		data.SpeakerHoldTime = types.Int64Null()
		data.SpeakerHoldTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.listener-hold-time-min.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MinimumListenerHoldTime = types.Int64Null()

			v := res.Get(path + "sxp-default.listener-hold-time-min.vipVariableName")
			data.MinimumListenerHoldTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MinimumListenerHoldTime = types.Int64Null()
			data.MinimumListenerHoldTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.listener-hold-time-min.vipValue")
			data.MinimumListenerHoldTime = types.Int64Value(v.Int())
			data.MinimumListenerHoldTimeVariable = types.StringNull()
		}
	} else {
		data.MinimumListenerHoldTime = types.Int64Null()
		data.MinimumListenerHoldTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.listener-hold-time-max.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MaximumListenerHoldTime = types.Int64Null()

			v := res.Get(path + "sxp-default.listener-hold-time-max.vipVariableName")
			data.MaximumListenerHoldTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MaximumListenerHoldTime = types.Int64Null()
			data.MaximumListenerHoldTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.listener-hold-time-max.vipValue")
			data.MaximumListenerHoldTime = types.Int64Value(v.Int())
			data.MaximumListenerHoldTimeVariable = types.StringNull()
		}
	} else {
		data.MaximumListenerHoldTime = types.Int64Null()
		data.MaximumListenerHoldTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-default.sxp-node-id-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpNodeIdType = types.StringNull()

		} else if value.String() == "ignore" {
			data.SxpNodeIdType = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-node-id-type.vipValue")
			data.SxpNodeIdType = types.StringValue(v.String())

		}
	} else {
		data.SxpNodeIdType = types.StringNull()

	}
	if value := res.Get(path + "sxp-default.sxp-node-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SxpNodeId = types.StringNull()

			v := res.Get(path + "sxp-default.sxp-node-id.vipVariableName")
			data.SxpNodeIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SxpNodeId = types.StringNull()
			data.SxpNodeIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "sxp-default.sxp-node-id.vipValue")
			data.SxpNodeId = types.StringValue(v.String())
			data.SxpNodeIdVariable = types.StringNull()
		}
	} else {
		data.SxpNodeId = types.StringNull()
		data.SxpNodeIdVariable = types.StringNull()
	}
	if value := res.Get(path + "sxp-connection-list.vipValue"); len(value.Array()) > 0 {
		data.SxpConnections = make([]CiscoTrustSecSxpConnections, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoTrustSecSxpConnections{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("connection-peer-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PeerIp = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PeerIp = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("connection-peer-ip.vipValue")
					item.PeerIp = types.StringValue(cv.String())

				}
			} else {
				item.PeerIp = types.StringNull()

			}
			if cValue := v.Get("connection-source-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIp = types.StringNull()

					cv := v.Get("connection-source-ip.vipVariableName")
					item.SourceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIp = types.StringNull()
					item.SourceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("connection-source-ip.vipValue")
					item.SourceIp = types.StringValue(cv.String())
					item.SourceIpVariable = types.StringNull()
				}
			} else {
				item.SourceIp = types.StringNull()
				item.SourceIpVariable = types.StringNull()
			}
			if cValue := v.Get("connection-preshared-key.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PresharedKey = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PresharedKey = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("connection-preshared-key.vipValue")
					item.PresharedKey = types.StringValue(cv.String())

				}
			} else {
				item.PresharedKey = types.StringNull()

			}
			if cValue := v.Get("connection-mode.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Mode = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Mode = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("connection-mode.vipValue")
					item.Mode = types.StringValue(cv.String())

				}
			} else {
				item.Mode = types.StringNull()

			}
			if cValue := v.Get("connection-mode-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ModeType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ModeType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("connection-mode-type.vipValue")
					item.ModeType = types.StringValue(cv.String())

				}
			} else {
				item.ModeType = types.StringNull()

			}
			if cValue := v.Get("connection-min-hold-time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MinimumHoldTime = types.Int64Null()

					cv := v.Get("connection-min-hold-time.vipVariableName")
					item.MinimumHoldTimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MinimumHoldTime = types.Int64Null()
					item.MinimumHoldTimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("connection-min-hold-time.vipValue")
					item.MinimumHoldTime = types.Int64Value(cv.Int())
					item.MinimumHoldTimeVariable = types.StringNull()
				}
			} else {
				item.MinimumHoldTime = types.Int64Null()
				item.MinimumHoldTimeVariable = types.StringNull()
			}
			if cValue := v.Get("connection-max-hold-time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MaximumHoldTime = types.Int64Null()

					cv := v.Get("connection-max-hold-time.vipVariableName")
					item.MaximumHoldTimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MaximumHoldTime = types.Int64Null()
					item.MaximumHoldTimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("connection-max-hold-time.vipValue")
					item.MaximumHoldTime = types.Int64Value(cv.Int())
					item.MaximumHoldTimeVariable = types.StringNull()
				}
			} else {
				item.MaximumHoldTime = types.Int64Null()
				item.MaximumHoldTimeVariable = types.StringNull()
			}
			if cValue := v.Get("connection-vpn-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("connection-vpn-id.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("connection-vpn-id.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			data.SxpConnections = append(data.SxpConnections, item)
			return true
		})
	} else {
		if len(data.SxpConnections) > 0 {
			data.SxpConnections = []CiscoTrustSecSxpConnections{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoTrustSec) hasChanges(ctx context.Context, state *CiscoTrustSec) bool {
	hasChanges := false
	if !data.DeviceSgt.Equal(state.DeviceSgt) {
		hasChanges = true
	}
	if !data.CredentialsId.Equal(state.CredentialsId) {
		hasChanges = true
	}
	if !data.CredentialsPassword.Equal(state.CredentialsPassword) {
		hasChanges = true
	}
	if !data.EnableEnforcement.Equal(state.EnableEnforcement) {
		hasChanges = true
	}
	if !data.EnableSxp.Equal(state.EnableSxp) {
		hasChanges = true
	}
	if !data.SxpSourceIp.Equal(state.SxpSourceIp) {
		hasChanges = true
	}
	if !data.SxpDefaultPassword.Equal(state.SxpDefaultPassword) {
		hasChanges = true
	}
	if !data.SxpKeyChain.Equal(state.SxpKeyChain) {
		hasChanges = true
	}
	if !data.SxpLogBindingChanges.Equal(state.SxpLogBindingChanges) {
		hasChanges = true
	}
	if !data.SxpReconciliationPeriod.Equal(state.SxpReconciliationPeriod) {
		hasChanges = true
	}
	if !data.SxpRetryPeriod.Equal(state.SxpRetryPeriod) {
		hasChanges = true
	}
	if !data.SpeakerHoldTime.Equal(state.SpeakerHoldTime) {
		hasChanges = true
	}
	if !data.MinimumListenerHoldTime.Equal(state.MinimumListenerHoldTime) {
		hasChanges = true
	}
	if !data.MaximumListenerHoldTime.Equal(state.MaximumListenerHoldTime) {
		hasChanges = true
	}
	if !data.SxpNodeIdType.Equal(state.SxpNodeIdType) {
		hasChanges = true
	}
	if !data.SxpNodeId.Equal(state.SxpNodeId) {
		hasChanges = true
	}
	if len(data.SxpConnections) != len(state.SxpConnections) {
		hasChanges = true
	} else {
		for i := range data.SxpConnections {
			if !data.SxpConnections[i].PeerIp.Equal(state.SxpConnections[i].PeerIp) {
				hasChanges = true
			}
			if !data.SxpConnections[i].SourceIp.Equal(state.SxpConnections[i].SourceIp) {
				hasChanges = true
			}
			if !data.SxpConnections[i].PresharedKey.Equal(state.SxpConnections[i].PresharedKey) {
				hasChanges = true
			}
			if !data.SxpConnections[i].Mode.Equal(state.SxpConnections[i].Mode) {
				hasChanges = true
			}
			if !data.SxpConnections[i].ModeType.Equal(state.SxpConnections[i].ModeType) {
				hasChanges = true
			}
			if !data.SxpConnections[i].MinimumHoldTime.Equal(state.SxpConnections[i].MinimumHoldTime) {
				hasChanges = true
			}
			if !data.SxpConnections[i].MaximumHoldTime.Equal(state.SxpConnections[i].MaximumHoldTime) {
				hasChanges = true
			}
			if !data.SxpConnections[i].VpnId.Equal(state.SxpConnections[i].VpnId) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
