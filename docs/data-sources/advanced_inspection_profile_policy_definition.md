---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_advanced_inspection_profile_policy_definition Data Source - terraform-provider-sdwan"
subcategory: "Security Policies"
description: |-
  This data source can read the Advanced Inspection Profile Policy Definition .
---

# sdwan_advanced_inspection_profile_policy_definition (Data Source)

This data source can read the Advanced Inspection Profile Policy Definition .

## Example Usage

```terraform
data "sdwan_advanced_inspection_profile_policy_definition" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `advanced_malware_protection_id` (String) Advanced malware protection ID
- `advanced_malware_protection_version` (Number) Advanced malware protection version
- `description` (String) The description of the policy definition.
- `intrusion_prevention_id` (String) Intrusion prevention ID (unified mode)
- `intrusion_prevention_version` (Number) Intrusion prevention version
- `name` (String) The name of the policy definition.
- `tls_action` (String) TLS Action
- `tls_ssl_decryption_id` (String) TLS/SSL decryption ID
- `tls_ssl_decryption_version` (Number) TLS/SSL decryption version
- `url_filtering_id` (String) URL filtering ID (unified mode)
- `url_filtering_version` (Number) URL filtering version
- `version` (Number) The version of the object