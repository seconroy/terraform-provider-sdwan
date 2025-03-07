resource "sdwan_transport_wan_vpn_interface_cellular_feature" "example" {
  name                                           = "Example"
  description                                    = "My Example"
  feature_profile_id                             = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id                   = "140331f6-5418-4755-a059-13c77eb96037"
  shutdown                                       = true
  interface_name                                 = "GigabitEthernet1"
  interface_description                          = "WAN"
  ipv4_dhcp_helper                               = ["1.2.3.4"]
  service_provider                               = "example"
  bandwidth_upstream                             = 21474836
  bandwidth_downstream                           = 21474836
  tunnel_interface                               = true
  per_tunnel_qos                                 = true
  tunnel_qos_mode                                = "hub"
  tunnel_bandwidth_percent                       = 82
  tunnel_interface_bind_loopback_tunnel          = "example"
  tunnel_interface_carrier                       = "default"
  tunnel_interface_color                         = "default"
  tunnel_interface_hello_interval                = 1000
  tunnel_interface_hello_tolerance               = 12
  tunnel_interface_last_resort_circuit           = false
  tunnel_interface_color_restrict                = true
  tunnel_interface_groups                        = 42949672
  tunnel_interface_border                        = false
  tunnel_interface_max_control_connections       = 62
  tunnel_interface_nat_refresh_interval          = 5
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [2]
  tunnel_interface_vmanage_connection_preference = 8
  tunnel_interface_port_hop                      = true
  tunnel_interface_low_bandwidth_link            = false
  tunnel_interface_tunnel_tcp_mss                = 1460
  tunnel_interface_clear_dont_fragment           = false
  tunnel_interface_network_broadcast             = false
  tunnel_interface_allow_all                     = false
  tunnel_interface_allow_bgp                     = false
  tunnel_interface_allow_dhcp                    = true
  tunnel_interface_allow_ntp                     = false
  tunnel_interface_allow_ssh                     = false
  tunnel_interface_allow_dns                     = true
  tunnel_interface_allow_icmp                    = true
  tunnel_interface_allow_https                   = true
  tunnel_interface_allow_ospf                    = false
  tunnel_interface_allow_stun                    = false
  tunnel_interface_allow_snmp                    = false
  tunnel_interface_allow_netconf                 = false
  tunnel_interface_allow_bfd                     = false
  tunnel_interface_encapsulations = [
    {
      encapsulation = "gre"
      preference    = 4294967
      weight        = 250
    }
  ]
  nat_ipv4                    = true
  nat_udp_timeout             = 1
  nat_tcp_timeout             = 60
  qos_adaptive                = false
  qos_shaping_rate            = 16
  acl_ipv4_egress_feature_id  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  acl_ipv6_ingress_feature_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  ip_mtu                = 1500
  interface_mtu         = 1500
  tcp_mss               = 505
  tloc_extension        = "tloc"
  tracker               = "example"
  ip_directed_broadcast = false
}
