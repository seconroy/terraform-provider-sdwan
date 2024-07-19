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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ZoneBasedFirewallPolicyDefinition struct {
	Id             types.String                                      `tfsdk:"id"`
	Version        types.Int64                                       `tfsdk:"version"`
	Name           types.String                                      `tfsdk:"name"`
	Description    types.String                                      `tfsdk:"description"`
	Mode           types.String                                      `tfsdk:"mode"`
	ApplyZonePairs []ZoneBasedFirewallPolicyDefinitionApplyZonePairs `tfsdk:"apply_zone_pairs"`
	DefaultAction  types.String                                      `tfsdk:"default_action"`
	Rules          []ZoneBasedFirewallPolicyDefinitionRules          `tfsdk:"rules"`
}

type ZoneBasedFirewallPolicyDefinitionApplyZonePairs struct {
	SourceZone      types.String `tfsdk:"source_zone"`
	DestinationZone types.String `tfsdk:"destination_zone"`
}

type ZoneBasedFirewallPolicyDefinitionRules struct {
	RuleOrder     types.Int64                                           `tfsdk:"rule_order"`
	RuleName      types.String                                          `tfsdk:"rule_name"`
	BaseAction    types.String                                          `tfsdk:"base_action"`
	MatchEntries  []ZoneBasedFirewallPolicyDefinitionRulesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []ZoneBasedFirewallPolicyDefinitionRulesActionEntries `tfsdk:"action_entries"`
}

type ZoneBasedFirewallPolicyDefinitionRulesMatchEntries struct {
	Type          types.String `tfsdk:"type"`
	PolicyId      types.String `tfsdk:"policy_id"`
	Value         types.String `tfsdk:"value"`
	ProtocolType  types.String `tfsdk:"protocol_type"`
	ValueVariable types.String `tfsdk:"value_variable"`
}
type ZoneBasedFirewallPolicyDefinitionRulesActionEntries struct {
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ZoneBasedFirewallPolicyDefinition) getPath() string {
	return "/template/policy/definition/zonebasedfw/"
}

// End of section. //template:end getPath

func (data ZoneBasedFirewallPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "zoneBasedFW")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.entries", []interface{}{})
		for _, item := range data.ApplyZonePairs {
			itemBody := ""
			if !item.SourceZone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceZone", item.SourceZone.ValueString())
			}
			if !item.DestinationZone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationZone", item.DestinationZone.ValueString())
			}
			body, _ = sjson.SetRaw(body, "definition.entries.-1", itemBody)
		}
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "definition.defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.sequences", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.RuleOrder.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.RuleOrder.ValueInt64())
			}
			if !item.RuleName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.RuleName.ValueString())
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", "zoneBasedFW")
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})

				type ServiceDetails struct {
					Name     string `json:"name"`
					Protocol string `json:"protocol"`
					Port     string `json:"port"`
				}

				type APIResponse struct {
					Data []map[string]ServiceDetails `json:"data"`
				}

				protocolMap := map[string]string{
					"udp":     "17",
					"tcp":     "6",
					"tcp udp": "6 17",
				}

				jsonResponse := `{
					"data": [
						{
							"snmp": {
								"name": "snmp",
								"protocol": "17",
								"port": "161"
							}
						},
						{
							"tcp": {
								"name": "tcp",
								"protocol": "6"
							}
						},
						{
							"udp": {
								"name": "udp",
								"protocol": "17"
							}
						},
						{
							"icmp": {
								"name": "icmp",
								"protocol": "6 17"
							}
						},
						{
							"echo": {
								"name": "echo",
								"protocol": "6 17",
								"port": "7"
							}
						},
						{
							"telnet": {
								"name": "telnet",
								"protocol": "6",
								"port": "23"
							}
						},
						{
							"wins": {
								"name": "wins",
								"protocol": "6",
								"port": "1512"
							}
						},
						{
							"n2h2server": {
								"name": "n2h2server",
								"protocol": "6 17",
								"port": "9285"
							}
						},
						{
							"nntp": {
								"name": "nntp",
								"protocol": "6",
								"port": "119"
							}
						},
						{
							"pptp": {
								"name": "pptp",
								"protocol": "6",
								"port": "1723"
							}
						},
						{
							"rtsp": {
								"name": "rtsp",
								"protocol": "6",
								"port": "554 8554"
							}
						},
						{
							"bootpc": {
								"name": "bootpc",
								"protocol": "17",
								"port": "68"
							}
						},
						{
							"gdoi": {
								"name": "gdoi",
								"protocol": "17",
								"port": "848"
							}
						},
						{
							"tacacs": {
								"name": "tacacs",
								"protocol": "17",
								"port": "49"
							}
						},
						{
							"gopher": {
								"name": "gopher",
								"protocol": "6",
								"port": "70"
							}
						},
						{
							"icabrowser": {
								"name": "icabrowser",
								"protocol": "17",
								"port": "1604"
							}
						},
						{
							"skinny": {
								"name": "skinny",
								"protocol": "6",
								"port": "2000"
							}
						},
						{
							"sunrpc": {
								"name": "sunrpc",
								"protocol": "6 17",
								"port": "111"
							}
						},
						{
							"biff": {
								"name": "biff",
								"protocol": "17",
								"port": "512"
							}
						},
						{
							"router": {
								"name": "router",
								"protocol": "17",
								"port": "520"
							}
						},
						{
							"ircs": {
								"name": "ircs",
								"protocol": "6",
								"port": "994"
							}
						},
						{
							"orasrv": {
								"name": "orasrv",
								"protocol": "6",
								"port": "1525-1529"
							}
						},
						{
							"ms-cluster-net": {
								"name": "ms-cluster-net",
								"protocol": "17",
								"port": "3343"
							}
						},
						{
							"kermit": {
								"name": "kermit",
								"protocol": "6",
								"port": "1649"
							}
						},
						{
							"isakmp": {
								"name": "isakmp",
								"protocol": "17",
								"port": "500"
							}
						},
						{
							"sshell": {
								"name": "sshell",
								"protocol": "6 17",
								"port": "614"
							}
						},
						{
							"realsecure": {
								"name": "realsecure",
								"protocol": "6",
								"port": "2998"
							}
						},
						{
							"ircu": {
								"name": "ircu",
								"protocol": "6 17",
								"port": "6665 6666"
							}
						},
						{
							"appleqtc": {
								"name": "appleqtc",
								"protocol": "17",
								"port": "458"
							}
						},
						{
							"pwdgen": {
								"name": "pwdgen",
								"protocol": "6 17",
								"port": "129"
							}
						},
						{
							"rdb-dbs-disp": {
								"name": "rdb-dbs-disp",
								"protocol": "6 17",
								"port": "1571"
							}
						},
						{
							"creativepartnr": {
								"name": "creativepartnr",
								"protocol": "6 17",
								"port": "455"
							}
						},
						{
							"finger": {
								"name": "finger",
								"protocol": "6",
								"port": "79"
							}
						},
						{
							"ftps": {
								"name": "ftps",
								"protocol": "6",
								"port": "990"
							}
						},
						{
							"giop": {
								"name": "giop",
								"protocol": "6 17",
								"port": "2481 2482"
							}
						},
						{
							"rsvd": {
								"name": "rsvd",
								"protocol": "6 17",
								"port": "168"
							}
						},
						{
							"hp-alarm-mgr": {
								"name": "hp-alarm-mgr",
								"protocol": "6 17",
								"port": "383"
							}
						},
						{
							"uucp": {
								"name": "uucp",
								"protocol": "6 17",
								"port": "540 541"
							}
						},
						{
							"kerberos": {
								"name": "kerberos",
								"protocol": "6 17",
								"port": "88 464 749 750"
							}
						},
						{
							"imap": {
								"name": "imap",
								"protocol": "6",
								"port": "143"
							}
						},
						{
							"time": {
								"name": "time",
								"protocol": "17",
								"port": "37"
							}
						},
						{
							"bootps": {
								"name": "bootps",
								"protocol": "17",
								"port": "67"
							}
						},
						{
							"tftp": {
								"name": "tftp",
								"protocol": "17",
								"port": "69"
							}
						},
						{
							"oracle": {
								"name": "oracle",
								"protocol": "17",
								"port": "2005"
							}
						},
						{
							"snmptrap": {
								"name": "snmptrap",
								"protocol": "17",
								"port": "162"
							}
						},
						{
							"http": {
								"name": "http",
								"protocol": "6",
								"port": "80"
							}
						},
						{
							"qmtp": {
								"name": "qmtp",
								"protocol": "6 17",
								"port": "209"
							}
						},
						{
							"radius": {
								"name": "radius",
								"protocol": "17",
								"port": "1812 1813"
							}
						},
						{
							"oracle-em-vp": {
								"name": "oracle-em-vp",
								"protocol": "6 17",
								"port": "1748 1809"
							}
						},
						{
							"tarantella": {
								"name": "tarantella",
								"protocol": "6",
								"port": "3144"
							}
						},
						{
							"pcanywheredata": {
								"name": "pcanywheredata",
								"protocol": "6",
								"port": "5631"
							}
						},
						{
							"ldap": {
								"name": "ldap",
								"protocol": "6",
								"port": "389"
							}
						},
						{
							"mgcp": {
								"name": "mgcp",
								"protocol": "17",
								"port": "2427"
							}
						},
						{
							"sqlsrv": {
								"name": "sqlsrv",
								"protocol": "6",
								"port": "156"
							}
						},
						{
							"hsrp": {
								"name": "hsrp",
								"protocol": "17",
								"port": "1985"
							}
						},
						{
							"cisco-net-mgmt": {
								"name": "cisco-net-mgmt",
								"protocol": "6 17",
								"port": "1741 1993"
							}
						},
						{
							"smtp": {
								"name": "smtp",
								"protocol": "6",
								"port": "25"
							}
						},
						{
							"pcanywherestat": {
								"name": "pcanywherestat",
								"protocol": "17",
								"port": "5632"
							}
						},
						{
							"exec": {
								"name": "exec",
								"protocol": "6",
								"port": "512"
							}
						},
						{
							"send": {
								"name": "send",
								"protocol": "6 17",
								"port": "169"
							}
						},
						{
							"stun": {
								"name": "stun",
								"protocol": "6 17",
								"port": "1990-1994"
							}
						},
						{
							"syslog": {
								"name": "syslog",
								"protocol": "17",
								"port": "514"
							}
						},
						{
							"ms-sql-m": {
								"name": "ms-sql-m",
								"protocol": "17",
								"port": "1434"
							}
						},
						{
							"citrix": {
								"name": "citrix",
								"protocol": "6 17",
								"port": "2512-2897"
							}
						},
						{
							"creativeserver": {
								"name": "creativeserver",
								"protocol": "6 17",
								"port": "453"
							}
						},
						{
							"cifs": {
								"name": "cifs",
								"protocol": "6 17",
								"port": "3020"
							}
						},
						{
							"cisco-sys": {
								"name": "cisco-sys",
								"protocol": "6 17",
								"port": "132"
							}
						},
						{
							"cisco-tna": {
								"name": "cisco-tna",
								"protocol": "6 17",
								"port": "131"
							}
						},
						{
							"ms-dotnetster": {
								"name": "ms-dotnetster",
								"protocol": "6 17",
								"port": "3126"
							}
						},
						{
							"gtpv1": {
								"name": "gtpv1",
								"protocol": "6 17",
								"port": "2123"
							}
						},
						{
							"gtpv0": {
								"name": "gtpv0",
								"protocol": "6 17",
								"port": "3386"
							}
						},
						{
							"imap3": {
								"name": "imap3",
								"protocol": "6",
								"port": "220"
							}
						},
						{
							"fcip-port": {
								"name": "fcip-port",
								"protocol": "6",
								"port": "3225"
							}
						},
						{
							"netbios-dgm": {
								"name": "netbios-dgm",
								"protocol": "6 17",
								"port": "138"
							}
						},
						{
							"sip-tls": {
								"name": "sip-tls",
								"protocol": "6 17",
								"port": "5061"
							}
						},
						{
							"pop3s": {
								"name": "pop3s",
								"protocol": "6",
								"port": "995"
							}
						},
						{
							"cisco-fna": {
								"name": "cisco-fna",
								"protocol": "6 17",
								"port": "130"
							}
						},
						{
							"802-11-iapp": {
								"name": "802-11-iapp",
								"protocol": "6 17",
								"port": "3517"
							}
						},
						{
							"oem-agent": {
								"name": "oem-agent",
								"protocol": "6 17",
								"port": "3872"
							}
						},
						{
							"cisco-tdp": {
								"name": "cisco-tdp",
								"protocol": "6 17",
								"port": "711"
							}
						},
						{
							"tr-rsrb": {
								"name": "tr-rsrb",
								"protocol": "6 17",
								"port": "1987-1996"
							}
						},
						{
							"r-winsock": {
								"name": "r-winsock",
								"protocol": "17",
								"port": "1745"
							}
						},
						{
							"sql-net": {
								"name": "sql-net",
								"protocol": "6",
								"port": "1521 150"
							}
						},
						{
							"syslog-conn": {
								"name": "syslog-conn",
								"protocol": "6",
								"port": "601"
							}
						},
						{
							"tacacs-ds": {
								"name": "tacacs-ds",
								"protocol": "6",
								"port": "65"
							}
						},
						{
							"h225ras": {
								"name": "h225ras",
								"protocol": "17",
								"port": "1719"
							}
						},
						{
							"ace-svr": {
								"name": "ace-svr",
								"protocol": "6 17",
								"port": "2475 2476"
							}
						},
						{
							"dhcp-failover": {
								"name": "dhcp-failover",
								"protocol": "6",
								"port": "647"
							}
						},
						{
							"igmpv3lite": {
								"name": "igmpv3lite",
								"protocol": "17",
								"port": "465"
							}
						},
						{
							"irc-serv": {
								"name": "irc-serv",
								"protocol": "17",
								"port": "529"
							}
						},
						{
							"entrust-svcs": {
								"name": "entrust-svcs",
								"protocol": "6 17",
								"port": "640 680 681"
							}
						},
						{
							"dbcontrol_agent": {
								"name": "dbcontrol_agent",
								"protocol": "6 17",
								"port": "3938"
							}
						},
						{
							"cisco-svcs": {
								"name": "cisco-svcs",
								"protocol": "6 17",
								"port": "1986-1999"
							}
						},
						{
							"ipsec-msft": {
								"name": "ipsec-msft",
								"protocol": "17",
								"port": "4500"
							}
						},
						{
							"microsoft-ds": {
								"name": "microsoft-ds",
								"protocol": "17",
								"port": "445"
							}
						},
						{
							"ms-sna": {
								"name": "ms-sna",
								"protocol": "6",
								"port": "1477 1478"
							}
						},
						{
							"rsvp_tunnel": {
								"name": "rsvp_tunnel",
								"protocol": "17",
								"port": "363"
							}
						},
						{
							"rsvp-encap": {
								"name": "rsvp-encap",
								"protocol": "6 17",
								"port": "1698 1699"
							}
						},
						{
							"hp-collector": {
								"name": "hp-collector",
								"protocol": "6 17",
								"port": "381"
							}
						},
						{
							"netbios-ns": {
								"name": "netbios-ns",
								"protocol": "17",
								"port": "137"
							}
						},
						{
							"msexch-routing": {
								"name": "msexch-routing",
								"protocol": "6",
								"port": "691"
							}
						},
						{
							"h323": {
								"name": "h323",
								"protocol": "6",
								"port": "1720"
							}
						},
						{
							"l2tp": {
								"name": "l2tp",
								"protocol": "17",
								"port": "1701"
							}
						},
						{
							"ldap-admin": {
								"name": "ldap-admin",
								"protocol": "6 17",
								"port": "3407"
							}
						},
						{
							"pop3": {
								"name": "pop3",
								"protocol": "6",
								"port": "110"
							}
						},
						{
							"h323callsigalt": {
								"name": "h323callsigalt",
								"protocol": "6 17",
								"port": "11720"
							}
						},
						{
							"ms-sql": {
								"name": "ms-sql",
								"protocol": "6",
								"port": "1433"
							}
						},
						{
							"iscsi-target": {
								"name": "iscsi-target",
								"protocol": "6",
								"port": "3260"
							}
						},
						{
							"webster": {
								"name": "webster",
								"protocol": "6",
								"port": "765"
							}
						},
						{
							"lotusnote": {
								"name": "lotusnote",
								"protocol": "6",
								"port": "1352"
							}
						},
						{
							"ipx": {
								"name": "ipx",
								"protocol": "17",
								"port": "213"
							}
						},
						{
							"entrust-svc-hand": {
								"name": "entrust-svc-hand",
								"protocol": "6 17",
								"port": "709 710"
							}
						},
						{
							"citriximaclient": {
								"name": "citriximaclient",
								"protocol": "6",
								"port": "2598"
							}
						},
						{
							"rtc-pm-port": {
								"name": "rtc-pm-port",
								"protocol": "6 17",
								"port": "3891"
							}
						},
						{
							"ftp": {
								"name": "ftp",
								"protocol": "6",
								"port": "21"
							}
						},
						{
							"aol": {
								"name": "aol",
								"protocol": "6 17",
								"port": "5190"
							}
						},
						{
							"xdmcp": {
								"name": "xdmcp",
								"protocol": "17",
								"port": "177"
							}
						},
						{
							"oraclenames": {
								"name": "oraclenames",
								"protocol": "6 17",
								"port": "1575"
							}
						},
						{
							"login": {
								"name": "login",
								"protocol": "6",
								"port": "513"
							}
						},
						{
							"iscsi": {
								"name": "iscsi",
								"protocol": "6",
								"port": "860"
							}
						},
						{
							"ttc": {
								"name": "ttc",
								"protocol": "6 17",
								"port": "2483 2484"
							}
						},
						{
							"imaps": {
								"name": "imaps",
								"protocol": "6",
								"port": "993"
							}
						},
						{
							"socks": {
								"name": "socks",
								"protocol": "6",
								"port": "1080"
							}
						},
						{
							"ssh": {
								"name": "ssh",
								"protocol": "6 17",
								"port": "22"
							}
						},
						{
							"dnsix": {
								"name": "dnsix",
								"protocol": "6",
								"port": "90"
							}
						},
						{
							"daytime": {
								"name": "daytime",
								"protocol": "6 17",
								"port": "13"
							}
						},
						{
							"sip": {
								"name": "sip",
								"protocol": "17",
								"port": "5060"
							}
						},
						{
							"discard": {
								"name": "discard",
								"protocol": "6 17",
								"port": "9"
							}
						},
						{
							"ntp": {
								"name": "ntp",
								"protocol": "17",
								"port": "123"
							}
						},
						{
							"ldaps": {
								"name": "ldaps",
								"protocol": "6 17",
								"port": "636"
							}
						},
						{
							"https": {
								"name": "https",
								"protocol": "6",
								"port": "443"
							}
						},
						{
							"vdolive": {
								"name": "vdolive",
								"protocol": "6",
								"port": "7000"
							}
						},
						{
							"ica": {
								"name": "ica",
								"protocol": "6",
								"port": "1494"
							}
						},
						{
							"net8-cman": {
								"name": "net8-cman",
								"protocol": "6 17",
								"port": "1630 1830"
							}
						},
						{
							"cuseeme": {
								"name": "cuseeme",
								"protocol": "6",
								"port": "7648"
							}
						},
						{
							"netstat": {
								"name": "netstat",
								"protocol": "6 17",
								"port": "15"
							}
						},
						{
							"sms": {
								"name": "sms",
								"protocol": "6 17",
								"port": "2701-2703"
							}
						},
						{
							"streamworks": {
								"name": "streamworks",
								"protocol": "17",
								"port": "1558"
							}
						},
						{
							"rtelnet": {
								"name": "rtelnet",
								"protocol": "6",
								"port": "107"
							}
						},
						{
							"who": {
								"name": "who",
								"protocol": "17",
								"port": "513"
							}
						},
						{
							"kazaa": {
								"name": "kazaa",
								"protocol": "6",
								"port": "1214"
							}
						},
						{
							"ssp": {
								"name": "ssp",
								"protocol": "6 17",
								"port": "3249"
							}
						},
						{
							"dbase": {
								"name": "dbase",
								"protocol": "6 17",
								"port": "217"
							}
						},
						{
							"timed": {
								"name": "timed",
								"protocol": "17",
								"port": "525"
							}
						},
						{
							"cddbp": {
								"name": "cddbp",
								"protocol": "6",
								"port": "888"
							}
						},
						{
							"telnets": {
								"name": "telnets",
								"protocol": "6",
								"port": "992"
							}
						},
						{
							"ymsgr": {
								"name": "ymsgr",
								"protocol": "6",
								"port": "5050"
							}
						},
						{
							"ident": {
								"name": "ident",
								"protocol": "6",
								"port": "113"
							}
						},
						{
							"bgp": {
								"name": "bgp",
								"protocol": "6",
								"port": "179"
							}
						},
						{
							"ddns-v3": {
								"name": "ddns-v3",
								"protocol": "6 17",
								"port": "2164"
							}
						},
						{
							"vqp": {
								"name": "vqp",
								"protocol": "6 17",
								"port": "1589"
							}
						},
						{
							"irc": {
								"name": "irc",
								"protocol": "6",
								"port": "194"
							}
						},
						{
							"ipass": {
								"name": "ipass",
								"protocol": "6 17",
								"port": "2549"
							}
						},
						{
							"x11": {
								"name": "x11",
								"protocol": "6",
								"port": "6000"
							}
						},
						{
							"dns": {
								"name": "dns",
								"protocol": "6 17",
								"port": "53"
							}
						},
						{
							"lotusmtap": {
								"name": "lotusmtap",
								"protocol": "6 17",
								"port": "3007"
							}
						},
						{
							"mysql": {
								"name": "mysql",
								"protocol": "6 17",
								"port": "3306"
							}
						},
						{
							"nfs": {
								"name": "nfs",
								"protocol": "6 17",
								"port": "2049"
							}
						},
						{
							"msnmsgr": {
								"name": "msnmsgr",
								"protocol": "6",
								"port": "1863"
							}
						},
						{
							"netshow": {
								"name": "netshow",
								"protocol": "6",
								"port": "1755"
							}
						},
						{
							"sqlserv": {
								"name": "sqlserv",
								"protocol": "6 17",
								"port": "118"
							}
						},
						{
							"hp-managed-node": {
								"name": "hp-managed-node",
								"protocol": "6 17",
								"port": "382"
							}
						},
						{
							"ncp": {
								"name": "ncp",
								"protocol": "6 17",
								"port": "524"
							}
						},
						{
							"shell": {
								"name": "shell",
								"protocol": "6",
								"port": "514"
							}
						},
						{
							"realmedia": {
								"name": "realmedia",
								"protocol": "6",
								"port": "7070"
							}
						},
						{
							"msrpc": {
								"name": "msrpc",
								"protocol": "6",
								"port": "135"
							}
						},
						{
							"clp": {
								"name": "clp",
								"protocol": "6 17",
								"port": "2567"
							}
						}
						]
					}`

				var response APIResponse
				err := json.Unmarshal([]byte(jsonResponse), &response)
				if err != nil {
					fmt.Println("HANDLE ERROR", err) // HANDLE ERROR
				}

				protocols, protocolNums, ports := "", "", ""

				for _, childItem := range item.MatchEntries {
					if !childItem.Type.IsNull() && childItem.Type.ValueString() == "protocolName" {

						for _, serviceMap := range response.Data {
							for _, serviceDetails := range serviceMap {

								for _, i := range strings.Split(childItem.Value.ValueString(), " ") {
									if serviceDetails.Name == i {

										if protocols == "" {
											protocols += serviceDetails.Name
										} else {
											protocols += " " + serviceDetails.Name
										}
										if protocolNums != "6 17" || protocolNums != "17 6" {
											if protocolNums == "" {
												protocolNums += serviceDetails.Protocol
											} else {
												protocolNums += " " + serviceDetails.Protocol
											}
										}
										if ports == "" {
											ports += serviceDetails.Port
										} else {
											ports += " " + serviceDetails.Port
										}
									}
								}
							}
						}

						if protocolNums == "" {
							fmt.Println("HANDLE ERROR", "a value was not found for protocols, protocolNums, or ports") // HANDLE ERROR
						}

						if ports != "" {
							item.MatchEntries = append(item.MatchEntries, ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{
								Type:         types.StringValue("destinationPort"),
								ProtocolType: types.StringValue(protocols),
								Value:        types.StringValue(ports),
							})
						}

						if protocols != "" {
							item.MatchEntries = append(item.MatchEntries, ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{
								Type:         types.StringValue("protocol"),
								ProtocolType: types.StringValue(protocols),
								Value:        types.StringValue(protocolNums),
							})
						}
					}
				}

				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.PolicyId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.PolicyId.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if !childItem.ProtocolType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "app", childItem.ProtocolType.ValueString())
					}
					if !childItem.ValueVariable.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "vipVariableName", childItem.ValueVariable.ValueString())
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
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "definition.sequences.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ZoneBasedFirewallPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("definition.entries"); value.Exists() && len(value.Array()) > 0 {
		data.ApplyZonePairs = make([]ZoneBasedFirewallPolicyDefinitionApplyZonePairs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ZoneBasedFirewallPolicyDefinitionApplyZonePairs{}
			if cValue := v.Get("sourceZone"); cValue.Exists() {
				item.SourceZone = types.StringValue(cValue.String())
			} else {
				item.SourceZone = types.StringNull()
			}
			if cValue := v.Get("destinationZone"); cValue.Exists() {
				item.DestinationZone = types.StringValue(cValue.String())
			} else {
				item.DestinationZone = types.StringNull()
			}
			data.ApplyZonePairs = append(data.ApplyZonePairs, item)
			return true
		})
	} else {
		if len(data.ApplyZonePairs) > 0 {
			data.ApplyZonePairs = []ZoneBasedFirewallPolicyDefinitionApplyZonePairs{}
		}
	}
	if value := res.Get("definition.defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("definition.sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Rules = make([]ZoneBasedFirewallPolicyDefinitionRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ZoneBasedFirewallPolicyDefinitionRules{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.RuleOrder = types.Int64Value(cValue.Int())
			} else {
				item.RuleOrder = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.RuleName = types.StringValue(cValue.String())
			} else {
				item.RuleName = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]ZoneBasedFirewallPolicyDefinitionRulesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() {
						cItem.PolicyId = types.StringValue(ccValue.String())
					} else {
						cItem.PolicyId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					if ccValue := cv.Get("app"); ccValue.Exists() {
						cItem.ProtocolType = types.StringValue(ccValue.String())
					} else {
						cItem.ProtocolType = types.StringNull()
					}
					if ccValue := cv.Get("vipVariableName"); ccValue.Exists() {
						cItem.ValueVariable = types.StringValue(ccValue.String())
					} else {
						cItem.ValueVariable = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]ZoneBasedFirewallPolicyDefinitionRulesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ZoneBasedFirewallPolicyDefinitionRulesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []ZoneBasedFirewallPolicyDefinitionRulesActionEntries{}
				}
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	} else {
		if len(data.Rules) > 0 {
			data.Rules = []ZoneBasedFirewallPolicyDefinitionRules{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ZoneBasedFirewallPolicyDefinition) hasChanges(ctx context.Context, state *ZoneBasedFirewallPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if len(data.ApplyZonePairs) != len(state.ApplyZonePairs) {
		hasChanges = true
	} else {
		for i := range data.ApplyZonePairs {
			if !data.ApplyZonePairs[i].SourceZone.Equal(state.ApplyZonePairs[i].SourceZone) {
				hasChanges = true
			}
			if !data.ApplyZonePairs[i].DestinationZone.Equal(state.ApplyZonePairs[i].DestinationZone) {
				hasChanges = true
			}
		}
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Rules) != len(state.Rules) {
		hasChanges = true
	} else {
		for i := range data.Rules {
			if !data.Rules[i].RuleOrder.Equal(state.Rules[i].RuleOrder) {
				hasChanges = true
			}
			if !data.Rules[i].RuleName.Equal(state.Rules[i].RuleName) {
				hasChanges = true
			}
			if !data.Rules[i].BaseAction.Equal(state.Rules[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Rules[i].MatchEntries) != len(state.Rules[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Rules[i].MatchEntries {
					if !data.Rules[i].MatchEntries[ii].Type.Equal(state.Rules[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].PolicyId.Equal(state.Rules[i].MatchEntries[ii].PolicyId) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].Value.Equal(state.Rules[i].MatchEntries[ii].Value) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].ProtocolType.Equal(state.Rules[i].MatchEntries[ii].ProtocolType) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].ValueVariable.Equal(state.Rules[i].MatchEntries[ii].ValueVariable) {
						hasChanges = true
					}
				}
			}
			if len(data.Rules[i].ActionEntries) != len(state.Rules[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Rules[i].ActionEntries {
					if !data.Rules[i].ActionEntries[ii].Type.Equal(state.Rules[i].ActionEntries[ii].Type) {
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

// End of section. //template:end updateVersions
