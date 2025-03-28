---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_intrusion_prevention_policy_definition Data Source - terraform-provider-sdwan"
subcategory: "(Classic) Security Policies"
description: |-
  This data source can read the Intrusion Prevention Policy Definition .
---

# sdwan_intrusion_prevention_policy_definition (Data Source)

This data source can read the Intrusion Prevention Policy Definition .

## Example Usage

```terraform
data "sdwan_intrusion_prevention_policy_definition" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `custom_signature` (Boolean) Custom signature
- `description` (String) The description of the policy definition
- `inspection_mode` (String) The inspection mode
- `ips_signature_list_id` (String) IPS signature list ID
- `ips_signature_list_version` (Number) IPS signature list version
- `log_level` (String) Log level
- `logging` (Attributes List) (see [below for nested schema](#nestedatt--logging))
- `mode` (String) The policy mode
- `name` (String) The name of the policy definition
- `signature_set` (String) Signature set
- `target_vpns` (Set of String) List of VPN IDs
- `version` (Number) The version of the object

<a id="nestedatt--logging"></a>
### Nested Schema for `logging`

Read-Only:

- `external_syslog_server_ip` (String) External Syslog Server IP
- `external_syslog_server_vpn` (String) External Syslog Server VPN
