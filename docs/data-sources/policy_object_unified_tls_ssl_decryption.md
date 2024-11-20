---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_unified_tls_ssl_decryption Data Source - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This data source can read the Policy Object Unified TLS SSL Decryption Policy_object.
---

# sdwan_policy_object_unified_tls_ssl_decryption (Data Source)

This data source can read the Policy Object Unified TLS SSL Decryption Policy_object.

## Example Usage

```terraform
data "sdwan_policy_object_unified_tls_ssl_decryption" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Policy_object

### Read-Only

- `bundle_string` (String)
- `certificate_lifetime` (String) If you have vManage as CA or vManage as intermediate CA, this value should be 1
- `certificate_revocation_status` (String) If value is none unknown status not required, if value is ocsp then unknown status is required
- `default_ca_certificate_bundle` (Boolean)
- `description` (String) The description of the Policy_object
- `ec_key_type` (String)
- `expired_certificate` (String)
- `failure_mode` (String)
- `file_name` (String)
- `minimal_tls_ver` (String)
- `name` (String) The name of the Policy_object
- `rsa_keypair_modules` (String)
- `unknown_revocation_status` (String) Only required if certificateRevocationStatus is oscp, if value is none then field shouldn't be here
- `unsupported_cipher_suites` (String)
- `unsupported_protocol_versions` (String)
- `untrusted_certificate` (String)
- `version` (Number) The version of the Policy_object