---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_gps_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - Transport"
description: |-
  This data source can read the Transport GPS Feature.
---

# sdwan_transport_gps_feature (Data Source)

This data source can read the Transport GPS Feature.

## Example Usage

```terraform
data "sdwan_transport_gps_feature" "example" {
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

- `description` (String) The description of the Feature
- `gps_enable` (Boolean) Enable/disable GPS
- `gps_enable_variable` (String) Variable name
- `gps_mode` (String) Select GPS mode
- `gps_mode_variable` (String) Variable name
- `name` (String) The name of the Feature
- `nmea_destination_address` (String) Destination address
- `nmea_destination_address_variable` (String) Variable name
- `nmea_destination_port` (Number) Destination port
- `nmea_destination_port_variable` (String) Variable name
- `nmea_enable` (Boolean) Enable/disable NMEA data
- `nmea_enable_variable` (String) Variable name
- `nmea_source_address` (String) Source address
- `nmea_source_address_variable` (String) Variable name
- `version` (Number) The version of the Feature
