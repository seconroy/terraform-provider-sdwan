---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_remote_access_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - System"
description: |-
  This data source can read the System Remote Access Feature.
---

# sdwan_system_remote_access_feature (Data Source)

This data source can read the System Remote Access Feature.

## Example Usage

```terraform
data "sdwan_system_remote_access_feature" "example" {
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

- `aaa_derive_name_from_peer_domain` (String)
- `aaa_derive_name_from_peer_domain_variable` (String) Variable name
- `aaa_derive_name_from_peer_identity` (String)
- `aaa_derive_name_from_peer_identity_variable` (String) Variable name
- `aaa_enable_accounting` (Boolean) Enable Accounting
- `aaa_enable_accounting_variable` (String) Variable name
- `aaa_specify_name_policy_name` (String)
- `aaa_specify_name_policy_name_variable` (String) Variable name
- `aaa_specify_name_policy_password` (String)
- `aaa_specify_name_policy_password_variable` (String) Variable name
- `any_connect_eap_authentication_type` (String)
- `connection_type_ssl` (Boolean) Enabled SSL VPN
- `description` (String) The description of the Feature
- `enable_certificate_list_check` (Boolean)
- `enable_certificate_list_check_variable` (String) Variable name
- `ikev2_anti_dos_threshold` (Number) Anti-DOS Threshold
- `ikev2_anti_dos_threshold_variable` (String) Variable name
- `ikev2_local_ike_identity_type` (String)
- `ikev2_local_ike_identity_type_variable` (String) Variable name
- `ikev2_local_ike_identity_value` (String)
- `ikev2_local_ike_identity_value_variable` (String) Variable name
- `ikev2_security_association_lifetime` (Number) Security Association Lifetime in Seconds
- `ikev2_security_association_lifetime_variable` (String) Variable name
- `ipsec_anti_replay_window_size` (Number) security Association Lifetime
- `ipsec_anti_replay_window_size_variable` (String) Variable name
- `ipsec_enable_anti_replay` (Boolean) Enable Anti-Replay
- `ipsec_enable_anti_replay_variable` (String) Variable name
- `ipsec_enable_perfect_foward_secrecy` (Boolean) security Association Lifetime
- `ipsec_enable_perfect_foward_secrecy_variable` (String) Variable name
- `ipsec_security_association_lifetime` (Number) Security Association Lifetime in Seconds
- `ipsec_security_association_lifetime_variable` (String) Variable name
- `ipv4_pool_size` (Number) IPv4 Pool Size
- `ipv4_pool_size_variable` (String) Variable name
- `ipv6_pool_size` (Number) IPv6 Pool Size
- `ipv6_pool_size_variable` (String) Variable name
- `name` (String) The name of the Feature
- `psk_authentication_pre_shared_key` (String) PSK Pre Shared Key
- `psk_authentication_pre_shared_key_variable` (String) Variable name
- `psk_authentication_type` (String) PSK Selection
- `psk_authentication_type_variable` (String) Variable name
- `radius_group_name` (String)
- `radius_group_name_variable` (String) Variable name
- `version` (Number) The version of the Feature
