---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_switchport_feature Data Source - terraform-provider-sdwan"
subcategory: "Features"
description: |-
  This data source can read the Service Switchport Feature.
---

# sdwan_service_switchport_feature (Data Source)

This data source can read the Service Switchport Feature.

## Example Usage

```terraform
data "sdwan_service_switchport_feature" "example" {
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

- `age_out_time` (Number) Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)
- `age_out_time_variable` (String) Variable name
- `description` (String) The description of the Feature
- `interfaces` (Attributes List) Interface name: GigabitEthernet0/<>/<> when present (see [below for nested schema](#nestedatt--interfaces))
- `name` (String) The name of the Feature
- `static_mac_addresses` (Attributes List) Add static MAC address entries for interface (see [below for nested schema](#nestedatt--static_mac_addresses))
- `version` (Number) The version of the Feature

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- `control_direction` (String) Set uni or bi directional authorization mode
- `control_direction_variable` (String) Variable name
- `critical_vlan` (Number) Set Critical VLAN
- `critical_vlan_variable` (String) Variable name
- `duplex` (String) Duplex mode
- `duplex_variable` (String) Variable name
- `enable_periodic_reauth` (Boolean) Enable Periodic Reauthentication
- `enable_periodic_reauth_variable` (String) Variable name
- `enable_voice` (Boolean) Enable Critical Voice VLAN
- `enable_voice_variable` (String) Variable name
- `guest_vlan` (Number) Set vlan to drop non-802.1x enabled clients into if client is not in MAB list
- `guest_vlan_variable` (String) Variable name
- `host_mode` (String) Set host mode
- `host_mode_variable` (String) Variable name
- `inactivity` (Number) Periodic Reauthentication Inactivity Timeout (in seconds)
- `inactivity_variable` (String) Variable name
- `interface_name` (String) Set Interface name
- `interface_name_variable` (String) Variable name
- `mac_authentication_bypass` (Boolean) MAC Authentication Bypass
- `mac_authentication_bypass_variable` (String) Variable name
- `mode` (String) Set type of switch port: access/trunk
- `pae_enable` (Boolean) Set 802.1x Interface Pae Type
- `pae_enable_variable` (String) Variable name
- `port_control` (String) Set Port-Control Mode
- `port_control_variable` (String) Variable name
- `reauthentication` (Number) Periodic Reauthentication Interval (in seconds)
- `reauthentication_variable` (String) Variable name
- `restricted_vlan` (Number) Set Restricted VLAN ID
- `restricted_vlan_variable` (String) Variable name
- `shutdown` (Boolean) Administrative state
- `shutdown_variable` (String) Variable name
- `speed` (String) Set interface speed
- `speed_variable` (String) Variable name
- `switchport_access_vlan` (Number) Set VLAN identifier associated with bridging domain
- `switchport_access_vlan_variable` (String) Variable name
- `switchport_trunk_allowed_vlans` (String) Configure VLAN IDs used with the trunk
- `switchport_trunk_allowed_vlans_variable` (String) Variable name
- `switchport_trunk_native_vlan` (Number) Configure VLAN ID used for native VLAN
- `switchport_trunk_native_vlan_variable` (String) Variable name
- `voice_vlan` (Number) Configure Voice Vlan
- `voice_vlan_variable` (String) Variable name


<a id="nestedatt--static_mac_addresses"></a>
### Nested Schema for `static_mac_addresses`

Read-Only:

- `interface_name` (String) Interface name: GigabitEthernet0/<>/<>
- `interface_name_variable` (String) Variable name
- `mac_address` (String) Set MAC address in xxxx.xxxx.xxxx format
- `mac_address_variable` (String) Variable name
- `vlan_id` (Number) Configure VLAN ID used with the mac and interface
- `vlan_id_variable` (String) Variable name