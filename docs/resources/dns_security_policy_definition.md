---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_dns_security_policy_definition Resource - terraform-provider-sdwan"
subcategory: "Security Policies"
description: |-
  This resource can manage a DNS Security Policy Definition .
---

# sdwan_dns_security_policy_definition (Resource)

This resource can manage a DNS Security Policy Definition .

## Example Usage

```terraform
resource "sdwan_dns_security_policy_definition" "example" {
  name                                      = "Example"
  description                               = "Example"
  domain_list_id                            = "84f10c9d-def7-45a3-8c64-6df26163c861"
  local_domain_bypass_enabled               = false
  match_all_vpn                             = true
  dnscrypt                                  = true
  umbrella_dns_default                      = true
  cisco_sig_credentials_feature_template_id = "0e7bb009-3603-441b-a24e-d5187679e800"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cisco_sig_credentials_feature_template_id` (String) Credentials feature template ID
- `description` (String) The description of the policy definition.
- `match_all_vpn` (Boolean) Should use match all VPN
- `name` (String) The name of the policy definition.
- `umbrella_dns_default` (Boolean) Should use umbrella as DNS Server

### Optional

- `cisco_sig_credentials_feature_template_version` (Number) Credentials feature template version
- `custom_dns_server_ip` (String) Only relevant when `umbrella_dns_default` is `false`
- `dnscrypt` (Boolean) Should DNSCrypt be enabled
- `domain_list_id` (String) Local domain bypass list ID
- `domain_list_version` (Number) Local domain bypass list version
- `local_domain_bypass_enabled` (Boolean) Should the local domain bypass list be enabled
- `target_vpns` (List of String) Only relevant when `match_all_vpn` is `false`

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_dns_security_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```