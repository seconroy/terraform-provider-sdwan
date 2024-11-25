---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_wan_vpn_interface_ethernet_feature_associate_tracker_group_feature Data Source - terraform-provider-sdwan"
subcategory: "Feature"
description: |-
  This data source can read the Transport WAN VPN Interface Ethernet Feature Associate Tracker Group Feature .
---

# sdwan_transport_wan_vpn_interface_ethernet_feature_associate_tracker_group_feature (Data Source)

This data source can read the Transport WAN VPN Interface Ethernet Feature Associate Tracker Group Feature .

## Example Usage

```terraform
data "sdwan_transport_wan_vpn_interface_ethernet_feature_associate_tracker_group_feature" "example" {
  feature_profile_id                              = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_feature_id                    = "140331f6-5418-4755-a059-13c77eb96037"
  transport_wan_vpn_interface_ethernet_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  id                                              = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the object
- `transport_wan_vpn_feature_id` (String) Transport WAN VPN Feature ID
- `transport_wan_vpn_interface_ethernet_feature_id` (String) Transport WAN VPN Interface Ethernet Feature ID

### Read-Only

- `transport_tracker_group_feature_id` (String) Transport Tracker Group Feature ID