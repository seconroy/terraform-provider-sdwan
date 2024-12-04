---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_dhcp_server_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - Service"
description: |-
  This data source can read the Service DHCP Server Feature.
---

# sdwan_service_dhcp_server_feature (Data Source)

This data source can read the Service DHCP Server Feature.

## Example Usage

```terraform
data "sdwan_service_dhcp_server_feature" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Feature

### Read-Only

- `default_gateway` (String) Set IP address of default gateway
- `default_gateway_variable` (String) Variable name
- `description` (String) The description of the Feature
- `dns_servers` (Set of String) Configure one or more DNS server IP addresses
- `dns_servers_variable` (String) Variable name
- `domain_name` (String) Set domain name client uses to resolve hostnames
- `domain_name_variable` (String) Variable name
- `exclude` (Set of String) Configure IPv4 address to exclude from DHCP address pool
- `exclude_variable` (String) Variable name
- `interface_mtu` (Number) Set MTU on interface to DHCP client
- `interface_mtu_variable` (String) Variable name
- `lease_time` (Number) Configure how long a DHCP-assigned IP address is valid
- `lease_time_variable` (String) Variable name
- `name` (String) The name of the Feature
- `network_address` (String) Network Address
- `network_address_variable` (String) Variable name
- `option_codes` (Attributes List) Configure Options Code (see [below for nested schema](#nestedatt--option_codes))
- `static_leases` (Attributes List) Configure static IP addresses (see [below for nested schema](#nestedatt--static_leases))
- `subnet_mask` (String) Subnet Mask
- `subnet_mask_variable` (String) Variable name
- `tftp_servers` (Set of String) Configure TFTP server IP addresses
- `tftp_servers_variable` (String) Variable name
- `version` (Number) The version of the Feature

<a id="nestedatt--option_codes"></a>
### Nested Schema for `option_codes`

Read-Only:

- `ascii` (String) Set ASCII value
- `ascii_variable` (String) Variable name
- `code` (Number) Set Option Code
- `code_variable` (String) Variable name
- `hex` (String) Set HEX value
- `hex_variable` (String) Variable name
- `ip` (Set of String) Set ip address
- `ip_variable` (String) Variable name


<a id="nestedatt--static_leases"></a>
### Nested Schema for `static_leases`

Read-Only:

- `ip_address` (String) Set client’s static IP address
- `ip_address_variable` (String) Variable name
- `mac_address` (String) Set MAC address of client
- `mac_address_variable` (String) Variable name