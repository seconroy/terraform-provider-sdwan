---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_wan_vpn_interface_ethernet_feature Data Source - terraform-provider-sdwan"
subcategory: "Features"
description: |-
  This data source can read the Transport WAN VPN Interface Ethernet Feature.
---

# sdwan_transport_wan_vpn_interface_ethernet_feature (Data Source)

This data source can read the Transport WAN VPN Interface Ethernet Feature.

## Example Usage

```terraform
data "sdwan_transport_wan_vpn_interface_ethernet_feature" "example" {
  id                           = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id           = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Feature
- `transport_wan_vpn_feature_id` (String) Transport WAN VPN Feature ID

### Read-Only

- `arp_timeout` (Number) Timeout value for dynamically learned ARP entries, <0..2678400> seconds
- `arp_timeout_variable` (String) Variable name
- `arps` (Attributes List) Configure ARP entries (see [below for nested schema](#nestedatt--arps))
- `auto_detect_bandwidth` (Boolean) Interface auto detect bandwidth
- `auto_detect_bandwidth_variable` (String) Variable name
- `autonegotiate` (Boolean) Link autonegotiation
- `autonegotiate_variable` (String) Variable name
- `bandwidth_downstream` (Number) Interface downstream bandwidth capacity, in kbps
- `bandwidth_downstream_variable` (String) Variable name
- `bandwidth_upstream` (Number) Interface upstream bandwidth capacity, in kbps
- `bandwidth_upstream_variable` (String) Variable name
- `block_non_source_ip` (Boolean) Block packets originating from IP address that is not from this source
- `block_non_source_ip_variable` (String) Variable name
- `description` (String) The description of the Feature
- `duplex` (String) Duplex mode
- `duplex_variable` (String) Variable name
- `enable_dhcpv6` (Boolean) Enable DHCPv6
- `gre_tunnel_source_ip` (String) GRE tunnel source IP
- `gre_tunnel_source_ip_variable` (String) Variable name
- `icmp_redirect_disable` (Boolean) ICMP/ICMPv6 Redirect Disable
- `icmp_redirect_disable_variable` (String) Variable name
- `interface_description` (String)
- `interface_description_variable` (String) Variable name
- `interface_mtu` (Number) Interface MTU GigabitEthernet0 <1500..1518>, Other GigabitEthernet <1500..9216> in bytes
- `interface_mtu_variable` (String) Variable name
- `interface_name` (String)
- `interface_name_variable` (String) Variable name
- `ip_directed_broadcast` (Boolean) IP Directed-Broadcast
- `ip_directed_broadcast_variable` (String) Variable name
- `ip_mtu` (Number) IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes
- `ip_mtu_variable` (String) Variable name
- `iperf_server` (String) Iperf server for auto bandwidth detect
- `iperf_server_variable` (String) Variable name
- `ipv4_address` (String) IP Address
- `ipv4_address_variable` (String) Variable name
- `ipv4_dhcp_distance` (Number) DHCP Distance
- `ipv4_dhcp_distance_variable` (String) Variable name
- `ipv4_dhcp_helper` (Set of String) List of DHCP IPv4 helper addresses (min 1, max 8)
- `ipv4_dhcp_helper_variable` (String) Variable name
- `ipv4_secondary_addresses` (Attributes List) Secondary IpV4 Addresses (see [below for nested schema](#nestedatt--ipv4_secondary_addresses))
- `ipv4_subnet_mask` (String) Subnet Mask
- `ipv4_subnet_mask_variable` (String) Variable name
- `ipv6_address` (String) IPv6 Address Secondary
- `ipv6_address_variable` (String) Variable name
- `ipv6_dhcp_secondary_address` (Attributes List) secondary IPv6 addresses (see [below for nested schema](#nestedatt--ipv6_dhcp_secondary_address))
- `ipv6_secondary_addresses` (Attributes List) Static secondary IPv6 addresses (see [below for nested schema](#nestedatt--ipv6_secondary_addresses))
- `load_interval` (Number) Interval for interface load calculation
- `load_interval_variable` (String) Variable name
- `mac_address` (String) MAC Address
- `mac_address_variable` (String) Variable name
- `media_type` (String) Media type
- `media_type_variable` (String) Variable name
- `name` (String) The name of the Feature
- `nat64` (Boolean) NAT64 on this interface
- `nat66` (Boolean) NAT66 on this interface
- `nat_ipv4` (Boolean) enable Network Address Translation on this interface
- `nat_ipv4_variable` (String) Variable name
- `nat_ipv6` (Boolean) enable Network Address Translation ipv6 on this interface
- `nat_ipv6_variable` (String) Variable name
- `nat_loopback` (String) NAT Inside Source Loopback Interface
- `nat_loopback_variable` (String) Variable name
- `nat_overload` (Boolean) NAT Overload
- `nat_overload_variable` (String) Variable name
- `nat_prefix_length` (Number) NAT Pool Prefix Length
- `nat_prefix_length_variable` (String) Variable name
- `nat_range_end` (String) NAT Pool Range End
- `nat_range_end_variable` (String) Variable name
- `nat_range_start` (String) NAT Pool Range Start
- `nat_range_start_variable` (String) Variable name
- `nat_tcp_timeout` (Number) Set NAT TCP session timeout, in minutes
- `nat_tcp_timeout_variable` (String) Variable name
- `nat_type` (String) NAT Type
- `nat_type_variable` (String) Variable name
- `nat_udp_timeout` (Number) Set NAT UDP session timeout, in minutes
- `nat_udp_timeout_variable` (String) Variable name
- `new_static_nats` (Attributes List) static NAT (see [below for nested schema](#nestedatt--new_static_nats))
- `per_tunnel_qos` (Boolean) Per-tunnel Qos
- `per_tunnel_qos_variable` (String) Variable name
- `qos_adaptive` (Boolean) Adaptive QoS
- `qos_adaptive_bandwidth_downstream` (Boolean) Shaping Rate Downstream
- `qos_adaptive_bandwidth_upstream` (Boolean) Shaping Rate Upstream
- `qos_adaptive_default_downstream` (Number) Adaptive QoS default downstream bandwidth (kbps)
- `qos_adaptive_default_downstream_variable` (String) Variable name
- `qos_adaptive_default_upstream` (Number) Adaptive QoS default upstream bandwidth (kbps)
- `qos_adaptive_default_upstream_variable` (String) Variable name
- `qos_adaptive_max_downstream` (Number) Downstream max bandwidth limit (kbps)
- `qos_adaptive_max_downstream_variable` (String) Variable name
- `qos_adaptive_max_upstream` (Number) Upstream max bandwidth limit (kbps)
- `qos_adaptive_max_upstream_variable` (String) Variable name
- `qos_adaptive_min_downstream` (Number) Downstream min bandwidth limit (kbps)
- `qos_adaptive_min_downstream_variable` (String) Variable name
- `qos_adaptive_min_upstream` (Number) Upstream min bandwidth limit (kbps)
- `qos_adaptive_min_upstream_variable` (String) Variable name
- `qos_adaptive_period` (Number) Adapt Period(Minutes)
- `qos_adaptive_period_variable` (String) Variable name
- `qos_shaping_rate` (Number) Shaping Rate (Kbps)
- `qos_shaping_rate_variable` (String) Variable name
- `service_provider` (String) Service Provider Name
- `service_provider_variable` (String) Variable name
- `shutdown` (Boolean)
- `shutdown_variable` (String) Variable name
- `speed` (String) Set interface speed
- `speed_variable` (String) Variable name
- `static_nat66` (Attributes List) static NAT66 (see [below for nested schema](#nestedatt--static_nat66))
- `tcp_mss` (Number) TCP MSS on SYN packets, in bytes
- `tcp_mss_variable` (String) Variable name
- `tloc_extension` (String) Extends a local TLOC to a remote node only for vpn 0
- `tloc_extension_variable` (String) Variable name
- `tracker` (String) Enable tracker for this interface
- `tracker_variable` (String) Variable name
- `tunnel_bandwidth_percent` (Number) Tunnels Bandwidth Percent
- `tunnel_bandwidth_percent_variable` (String) Variable name
- `tunnel_interface` (Boolean) Tunnel Interface on/off
- `tunnel_interface_allow_all` (Boolean) Allow all traffic. Overrides all other allow-service options if allow-service all is set
- `tunnel_interface_allow_all_variable` (String) Variable name
- `tunnel_interface_allow_bfd` (Boolean) Allow/Deny BFD
- `tunnel_interface_allow_bfd_variable` (String) Variable name
- `tunnel_interface_allow_bgp` (Boolean) Allow/deny BGP
- `tunnel_interface_allow_bgp_variable` (String) Variable name
- `tunnel_interface_allow_dhcp` (Boolean) Allow/Deny DHCP
- `tunnel_interface_allow_dhcp_variable` (String) Variable name
- `tunnel_interface_allow_dns` (Boolean) Allow/Deny DNS
- `tunnel_interface_allow_dns_variable` (String) Variable name
- `tunnel_interface_allow_https` (Boolean) Allow/Deny HTTPS
- `tunnel_interface_allow_https_variable` (String) Variable name
- `tunnel_interface_allow_icmp` (Boolean) Allow/Deny ICMP
- `tunnel_interface_allow_icmp_variable` (String) Variable name
- `tunnel_interface_allow_netconf` (Boolean) Allow/Deny NETCONF
- `tunnel_interface_allow_netconf_variable` (String) Variable name
- `tunnel_interface_allow_ntp` (Boolean) Allow/Deny NTP
- `tunnel_interface_allow_ntp_variable` (String) Variable name
- `tunnel_interface_allow_ospf` (Boolean) Allow/Deny OSPF
- `tunnel_interface_allow_ospf_variable` (String) Variable name
- `tunnel_interface_allow_snmp` (Boolean) Allow/Deny SNMP
- `tunnel_interface_allow_snmp_variable` (String) Variable name
- `tunnel_interface_allow_ssh` (Boolean) Allow/Deny SSH
- `tunnel_interface_allow_ssh_variable` (String) Variable name
- `tunnel_interface_allow_stun` (Boolean) Allow/Deny STUN
- `tunnel_interface_allow_stun_variable` (String) Variable name
- `tunnel_interface_bind_loopback_tunnel` (String) Bind loopback tunnel interface to a physical interface
- `tunnel_interface_bind_loopback_tunnel_variable` (String) Variable name
- `tunnel_interface_border` (Boolean) Set TLOC as border TLOC
- `tunnel_interface_border_variable` (String) Variable name
- `tunnel_interface_carrier` (String) Set carrier for TLOC
- `tunnel_interface_carrier_variable` (String) Variable name
- `tunnel_interface_clear_dont_fragment` (Boolean) Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)
- `tunnel_interface_clear_dont_fragment_variable` (String) Variable name
- `tunnel_interface_color` (String) Set color for TLOC
- `tunnel_interface_color_restrict` (Boolean) Restrict this TLOC behavior
- `tunnel_interface_color_restrict_variable` (String) Variable name
- `tunnel_interface_color_variable` (String) Variable name
- `tunnel_interface_cts_sgt_propagation` (Boolean) CTS SGT Propagation configuration
- `tunnel_interface_cts_sgt_propagation_variable` (String) Variable name
- `tunnel_interface_encapsulations` (Attributes List) Encapsulation for TLOC (see [below for nested schema](#nestedatt--tunnel_interface_encapsulations))
- `tunnel_interface_exclude_controller_group_list` (Set of Number) Exclude the following controller groups defined in this list.
- `tunnel_interface_exclude_controller_group_list_variable` (String) Variable name
- `tunnel_interface_gre_tunnel_destination_ip` (String) GRE tunnel destination IP
- `tunnel_interface_gre_tunnel_destination_ip_variable` (String) Variable name
- `tunnel_interface_groups` (Number) List of groups
- `tunnel_interface_groups_variable` (String) Variable name
- `tunnel_interface_hello_interval` (Number) Set time period of control hello packets <100..600000> milli seconds
- `tunnel_interface_hello_interval_variable` (String) Variable name
- `tunnel_interface_hello_tolerance` (Number) Set tolerance of control hello packets <12..6000> seconds
- `tunnel_interface_hello_tolerance_variable` (String) Variable name
- `tunnel_interface_last_resort_circuit` (Boolean) Set TLOC as last resort
- `tunnel_interface_last_resort_circuit_variable` (String) Variable name
- `tunnel_interface_low_bandwidth_link` (Boolean) Set the interface as a low-bandwidth circuit
- `tunnel_interface_low_bandwidth_link_variable` (String) Variable name
- `tunnel_interface_max_control_connections` (Number) Maximum Control Connections
- `tunnel_interface_max_control_connections_variable` (String) Variable name
- `tunnel_interface_nat_refresh_interval` (Number) Set time period of nat refresh packets <1...60> seconds
- `tunnel_interface_nat_refresh_interval_variable` (String) Variable name
- `tunnel_interface_network_broadcast` (Boolean) Accept and respond to network-prefix-directed broadcasts
- `tunnel_interface_network_broadcast_variable` (String) Variable name
- `tunnel_interface_port_hop` (Boolean) Disallow port hopping on the tunnel interface
- `tunnel_interface_port_hop_variable` (String) Variable name
- `tunnel_interface_tunnel_tcp_mss` (Number) Tunnel TCP MSS on SYN packets, in bytes
- `tunnel_interface_tunnel_tcp_mss_variable` (String) Variable name
- `tunnel_interface_vbond_as_stun_server` (Boolean) Put this wan interface in STUN mode only
- `tunnel_interface_vbond_as_stun_server_variable` (String) Variable name
- `tunnel_interface_vmanage_connection_preference` (Number) Set interface preference for control connection to vManage <0..8>
- `tunnel_interface_vmanage_connection_preference_variable` (String) Variable name
- `tunnel_qos_mode` (String) Set tunnel QoS mode
- `tunnel_qos_mode_variable` (String) Variable name
- `version` (Number) The version of the Feature
- `xconnect` (String) Extend remote TLOC over a GRE tunnel to a local WAN interface
- `xconnect_variable` (String) Variable name

<a id="nestedatt--arps"></a>
### Nested Schema for `arps`

Read-Only:

- `ip_address` (String) IP V4 Address
- `ip_address_variable` (String) Variable name
- `mac_address` (String) MAC Address
- `mac_address_variable` (String) Variable name


<a id="nestedatt--ipv4_secondary_addresses"></a>
### Nested Schema for `ipv4_secondary_addresses`

Read-Only:

- `address` (String) IpV4 Address
- `address_variable` (String) Variable name
- `subnet_mask` (String) Subnet Mask
- `subnet_mask_variable` (String) Variable name


<a id="nestedatt--ipv6_dhcp_secondary_address"></a>
### Nested Schema for `ipv6_dhcp_secondary_address`

Read-Only:

- `address` (String) IPv6 Address Secondary
- `address_variable` (String) Variable name


<a id="nestedatt--ipv6_secondary_addresses"></a>
### Nested Schema for `ipv6_secondary_addresses`

Read-Only:

- `address` (String) IPv6 Address Secondary
- `address_variable` (String) Variable name


<a id="nestedatt--new_static_nats"></a>
### Nested Schema for `new_static_nats`

Read-Only:

- `direction` (String) Direction of static NAT translation
- `source_ip` (String) Source IP address to be translated
- `source_ip_variable` (String) Variable name
- `source_vpn` (Number) Source VPN ID
- `source_vpn_variable` (String) Variable name
- `translated_ip` (String) Statically translated source IP address
- `translated_ip_variable` (String) Variable name


<a id="nestedatt--static_nat66"></a>
### Nested Schema for `static_nat66`

Read-Only:

- `source_prefix` (String) Source Prefix
- `source_prefix_variable` (String) Variable name
- `source_vpn_id` (Number) Source VPN ID
- `source_vpn_id_variable` (String) Variable name
- `translated_source_prefix` (String) Translated Source Prefix
- `translated_source_prefix_variable` (String) Variable name


<a id="nestedatt--tunnel_interface_encapsulations"></a>
### Nested Schema for `tunnel_interface_encapsulations`

Read-Only:

- `encapsulation` (String) Encapsulation
- `preference` (Number) Set preference for TLOC
- `preference_variable` (String) Variable name
- `weight` (Number) Set weight for TLOC
- `weight_variable` (String) Variable name