---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_feature_profile Resource - terraform-provider-sdwan"
subcategory: "Feature Profiles"
description: |-
  SD-WAN manager supports only a single policy object feature profile to be configured. This resource will therefore either create a new one if none exists or update the existing one.
---

# sdwan_policy_object_feature_profile (Resource)

SD-WAN manager supports only a single policy object feature profile to be configured. This resource will therefore either create a new one if none exists or update the existing one.

## Example Usage

```terraform
resource "sdwan_policy_object_feature_profile" "example" {
  name        = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) Description
- `name` (String) The name of the policy object feature profile

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_policy_object_feature_profile.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```