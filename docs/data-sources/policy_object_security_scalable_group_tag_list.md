---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_security_scalable_group_tag_list Data Source - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This data source can read the Policy Object Security Scalable Group Tag List Policy_object.
---

# sdwan_policy_object_security_scalable_group_tag_list (Data Source)

This data source can read the Policy Object Security Scalable Group Tag List Policy_object.

## Example Usage

```terraform
data "sdwan_policy_object_security_scalable_group_tag_list" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = ""
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Policy_object

### Read-Only

- `description` (String) The description of the Policy_object
- `entries` (Attributes List) (see [below for nested schema](#nestedatt--entries))
- `name` (String) The name of the Policy_object
- `version` (Number) The version of the Policy_object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Read-Only:

- `sgt_name` (String)
- `tag` (String)