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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportWANVPNFeatureAssociateRoutingOSPFFeature struct {
	Id                            types.String `tfsdk:"id"`
	FeatureProfileId              types.String `tfsdk:"feature_profile_id"`
	TransportWanVpnFeatureId      types.String `tfsdk:"transport_wan_vpn_feature_id"`
	TransportRoutingOspfFeatureId types.String `tfsdk:"transport_routing_ospf_feature_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNFeatureAssociateRoutingOSPFFeature) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/routing/ospf/", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNFeatureAssociateRoutingOSPFFeature) toBody(ctx context.Context) string {
	body := ""
	if !data.TransportRoutingOspfFeatureId.IsNull() {
		body, _ = sjson.Set(body, "parcelId", data.TransportRoutingOspfFeatureId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNFeatureAssociateRoutingOSPFFeature) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("parcelId"); value.Exists() {
		data.TransportRoutingOspfFeatureId = types.StringValue(value.String())
	} else {
		data.TransportRoutingOspfFeatureId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TransportWANVPNFeatureAssociateRoutingOSPFFeature) hasChanges(ctx context.Context, state *TransportWANVPNFeatureAssociateRoutingOSPFFeature) bool {
	hasChanges := false
	if !data.FeatureProfileId.Equal(state.FeatureProfileId) {
		hasChanges = true
	}
	if !data.TransportWanVpnFeatureId.Equal(state.TransportWanVpnFeatureId) {
		hasChanges = true
	}
	if !data.TransportRoutingOspfFeatureId.Equal(state.TransportRoutingOspfFeatureId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions