---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_bfd_profile_parcel Data Source - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This data source can read the System BFD profile parcel.
---

# sdwan_system_bfd_profile_parcel (Data Source)

This data source can read the System BFD profile parcel.

## Example Usage

```terraform
data "sdwan_system_bfd_profile_parcel" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the profile parcel

### Read-Only

- `colors` (Attributes List) Set color that identifies the WAN transport tunnel (see [below for nested schema](#nestedatt--colors))
- `default_dscp` (Number)
- `default_dscp_variable` (String) Variable name
- `description` (String) The description of the profile parcel
- `multiplier` (Number)
- `multiplier_variable` (String) Variable name
- `name` (String) The name of the profile parcel
- `poll_interval` (Number)
- `poll_interval_variable` (String) Variable name
- `version` (Number) The version of the profile parcel

<a id="nestedatt--colors"></a>
### Nested Schema for `colors`

Read-Only:

- `color` (String) Color that identifies the WAN transport tunnel
- `color_variable` (String) Variable name
- `dscp` (Number) BFD Default DSCP value for tloc color
- `dscp_variable` (String) Variable name
- `hello_interval` (Number) Hello Interval (milliseconds)
- `hello_interval_variable` (String) Variable name
- `multiplier` (Number) Multiplier
- `multiplier_variable` (String) Variable name
- `pmtu_discovery` (Boolean) Path MTU Discovery
- `pmtu_discovery_variable` (String) Variable name