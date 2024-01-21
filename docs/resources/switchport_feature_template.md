---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_switchport_feature_template Resource - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This resource can manage a Switchport feature template.
    - Minimum SD-WAN Manager version: 15.0.0
---

# sdwan_switchport_feature_template (Resource)

This resource can manage a Switchport feature template.
  - Minimum SD-WAN Manager version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_switchport_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  slot         = 0
  sub_slot     = 0
  module_type  = "4"
  interfaces = [
    {
      name                                     = "GigabitEthernet0/0/0"
      switchport_mode                          = "access"
      shutdown                                 = true
      speed                                    = "100"
      duplex                                   = "full"
      switchport_access_vlan                   = 100
      switchport_trunk_allowed_vlans           = "100,200"
      switchport_trunk_native_vlan             = 100
      dot1x_enable                             = true
      dot1x_port_control                       = "auto"
      dot1x_authentication_order               = ["dot1x"]
      dot1x_voice_vlan                         = 200
      dot1x_pae_enable                         = true
      dot1x_mac_authentication_bypass          = true
      dot1x_host_mode                          = "multi-domain"
      dot1x_enable_periodic_reauth             = true
      dot1x_periodic_reauth_inactivity_timeout = 100
      dot1x_periodic_reauth_interval           = 60
      dot1x_control_direction                  = "both"
      dot1x_restricted_vlan                    = 100
      dot1x_guest_vlan                         = 101
      dot1x_critical_vlan                      = 102
      dot1x_enable_criticial_voice_vlan        = true
    }
  ]
  age_out_time = 500
  static_mac_addresses = [
    {
      mac_address = "0000.0000.0000"
      if_name     = "GigabitEthernet0/0/0"
      vlan        = 100
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template

### Optional

- `age_out_time` (Number) Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)
  - Range: `0`-`1000000`
  - Default value: `300`
- `age_out_time_variable` (String) Variable name
- `interfaces` (Attributes List) Interface name: GigabitEthernet0/<>/<> when present (see [below for nested schema](#nestedatt--interfaces))
- `module_type` (String) Module type
  - Choices: `4`, `8`, `22`, `50`
- `slot` (Number) Number of Slots
  - Range: `0`-`31`
  - Default value: `0`
- `static_mac_addresses` (Attributes List) Add static MAC address entries for interface (see [below for nested schema](#nestedatt--static_mac_addresses))
- `sub_slot` (Number) Number of Sub-Slots
  - Range: `0`-`31`
  - Default value: `0`

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Optional:

- `dot1x_authentication_order` (List of String) Specify authentication methods in the order of preference
- `dot1x_authentication_order_variable` (String) Variable name
- `dot1x_control_direction` (String) Set uni or bi directional authorization mode
  - Choices: `both`, `in`
  - Default value: `both`
- `dot1x_control_direction_variable` (String) Variable name
- `dot1x_critical_vlan` (Number) Set Critical VLAN
  - Range: `1`-`4094`
- `dot1x_critical_vlan_variable` (String) Variable name
- `dot1x_enable` (Boolean) Set 802.1x on off
  - Default value: `true`
- `dot1x_enable_criticial_voice_vlan` (Boolean) Enable Critical Voice VLAN
  - Default value: `false`
- `dot1x_enable_criticial_voice_vlan_variable` (String) Variable name
- `dot1x_enable_periodic_reauth` (Boolean) Enable Periodic Reauthentication
  - Default value: `false`
- `dot1x_enable_periodic_reauth_variable` (String) Variable name
- `dot1x_enable_variable` (String) Variable name
- `dot1x_guest_vlan` (Number) Set vlan to drop non-802.1x enabled clients into if client is not in MAB list
  - Range: `1`-`4094`
- `dot1x_guest_vlan_variable` (String) Variable name
- `dot1x_host_mode` (String) Set host mode
  - Choices: `single-host`, `multi-auth`, `multi-host`, `multi-domain`
  - Default value: `single-host`
- `dot1x_host_mode_variable` (String) Variable name
- `dot1x_mac_authentication_bypass` (Boolean) MAC Authentication Bypass
  - Default value: `false`
- `dot1x_mac_authentication_bypass_variable` (String) Variable name
- `dot1x_pae_enable` (Boolean) Set 802.1x Interface Pae Type
  - Default value: `true`
- `dot1x_pae_enable_variable` (String) Variable name
- `dot1x_periodic_reauth_inactivity_timeout` (Number) Periodic Reauthentication Inactivity Timeout (in seconds)
  - Range: `1`-`1440`
  - Default value: `60`
- `dot1x_periodic_reauth_inactivity_timeout_variable` (String) Variable name
- `dot1x_periodic_reauth_interval` (Number) Periodic Reauthentication Interval (in seconds)
  - Range: `0`-`1440`
  - Default value: `0`
- `dot1x_periodic_reauth_interval_variable` (String) Variable name
- `dot1x_port_control` (String) Set Port-Control Mode
  - Choices: `auto`, `force-unauthorized`, `force-authorized`
  - Default value: `auto`
- `dot1x_port_control_variable` (String) Variable name
- `dot1x_restricted_vlan` (Number) Set Restricted VLAN ID
  - Range: `1`-`4094`
- `dot1x_restricted_vlan_variable` (String) Variable name
- `dot1x_voice_vlan` (Number) Configure Voice Vlan
  - Range: `1`-`4094`
- `dot1x_voice_vlan_variable` (String) Variable name
- `duplex` (String) Duplex mode
  - Choices: `full`, `half`
- `duplex_variable` (String) Variable name
- `name` (String) Set Interface name
- `name_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `shutdown` (Boolean) Administrative state
  - Default value: `true`
- `shutdown_variable` (String) Variable name
- `speed` (String) Set interface speed
  - Choices: `10`, `100`, `1000`, `2500`, `10000`
- `speed_variable` (String) Variable name
- `switchport_access_vlan` (Number) Set VLAN identifier associated with bridging domain
  - Range: `1`-`4094`
- `switchport_access_vlan_variable` (String) Variable name
- `switchport_mode` (String) Set type of switch port: access/trunk
  - Choices: `access`, `trunk`
- `switchport_trunk_allowed_vlans` (String) Configure VLAN IDs used with the trunk
- `switchport_trunk_allowed_vlans_variable` (String) Variable name
- `switchport_trunk_native_vlan` (Number) Configure VLAN ID used for native VLAN
  - Range: `1`-`4094`
- `switchport_trunk_native_vlan_variable` (String) Variable name


<a id="nestedatt--static_mac_addresses"></a>
### Nested Schema for `static_mac_addresses`

Optional:

- `if_name` (String) Interface name: GigabitEthernet0/<>/<>
- `if_name_variable` (String) Variable name
- `mac_address` (String) Set MAC address in xxxx.xxxx.xxxx format
- `mac_address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `vlan` (Number) Configure VLAN ID used with the mac and interface
  - Range: `1`-`4094`
- `vlan_variable` (String) Variable name

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_switchport_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```