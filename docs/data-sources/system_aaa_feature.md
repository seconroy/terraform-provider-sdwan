---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_aaa_feature Data Source - terraform-provider-sdwan"
subcategory: "Features"
description: |-
  This data source can read the System AAA Feature.
---

# sdwan_system_aaa_feature (Data Source)

This data source can read the System AAA Feature.

## Example Usage

```terraform
data "sdwan_system_aaa_feature" "example" {
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

- `accounting_group` (Boolean) Accounting configurations parameters
- `accounting_group_variable` (String) Variable name
- `accounting_rules` (Attributes List) Configure the accounting rules (see [below for nested schema](#nestedatt--accounting_rules))
- `authentication_group` (Boolean) Authentication configurations parameters
- `authentication_group_variable` (String) Variable name
- `authorization_config_commands` (Boolean) For configuration mode commands.
- `authorization_config_commands_variable` (String) Variable name
- `authorization_console` (Boolean) For enabling console authorization
- `authorization_console_variable` (String) Variable name
- `authorization_rules` (Attributes List) Configure the Authorization Rules (see [below for nested schema](#nestedatt--authorization_rules))
- `description` (String) The description of the Feature
- `name` (String) The name of the Feature
- `radius_groups` (Attributes List) Configure the Radius serverGroup (see [below for nested schema](#nestedatt--radius_groups))
- `server_auth_order` (Set of String) ServerGroups priority order
- `tacacs_groups` (Attributes List) Configure the TACACS serverGroup (see [below for nested schema](#nestedatt--tacacs_groups))
- `users` (Attributes List) Create local login account (see [below for nested schema](#nestedatt--users))
- `version` (Number) The version of the Feature

<a id="nestedatt--accounting_rules"></a>
### Nested Schema for `accounting_rules`

Read-Only:

- `group` (Set of String) Use Server-group
- `level` (String) Privilege level when method is commands
- `method` (String) Configure Accounting Method
- `rule_id` (String) Configure Accounting Rule ID
- `start_stop` (Boolean) Record start and stop without waiting
- `start_stop_variable` (String) Variable name


<a id="nestedatt--authorization_rules"></a>
### Nested Schema for `authorization_rules`

Read-Only:

- `group` (Set of String) Use Server-group
- `if_authenticated` (Boolean) Succeed if user has authenticated
- `level` (String) Privilege level when method is commands
- `method` (String) Method
- `rule_id` (String) Configure Authorization Rule ID


<a id="nestedatt--radius_groups"></a>
### Nested Schema for `radius_groups`

Read-Only:

- `group_name` (String) Set Radius server Group Name
- `servers` (Attributes List) Configure the Radius server (see [below for nested schema](#nestedatt--radius_groups--servers))
- `source_interface` (String) Set interface to use to reach Radius server
- `source_interface_variable` (String) Variable name
- `vpn` (Number) Set VPN in which Radius server is located

<a id="nestedatt--radius_groups--servers"></a>
### Nested Schema for `radius_groups.servers`

Read-Only:

- `acct_port` (Number) Set Accounting port to use to connect to Radius server
- `acct_port_variable` (String) Variable name
- `address` (String) Set IP address of Radius server
- `auth_port` (Number) Set Authentication port to use to connect to Radius server
- `auth_port_variable` (String) Variable name
- `key` (String) Set the Radius server shared key
- `key_enum` (String) Type of encyption. To be used for type 6
- `key_type` (String) key type
- `key_type_variable` (String) Variable name
- `retransmit` (Number) Configure how many times to contact this Radius server
- `retransmit_variable` (String) Variable name
- `secret_key` (String) Set the Radius server shared type 7 encrypted key
- `secret_key_variable` (String) Variable name
- `timeout` (Number) Configure how long to wait for replies from the Radius server
- `timeout_variable` (String) Variable name



<a id="nestedatt--tacacs_groups"></a>
### Nested Schema for `tacacs_groups`

Read-Only:

- `group_name` (String) Set TACACS server Group Name
- `servers` (Attributes List) Configure the TACACS server (see [below for nested schema](#nestedatt--tacacs_groups--servers))
- `source_interface` (String) Set interface to use to reach TACACS server
- `source_interface_variable` (String) Variable name
- `vpn` (Number) Set VPN in which TACACS server is located

<a id="nestedatt--tacacs_groups--servers"></a>
### Nested Schema for `tacacs_groups.servers`

Read-Only:

- `address` (String) Set IP address of TACACS server
- `key` (String) Set the TACACS server shared key
- `key_enum` (String) Type of encyption. To be used for type 6
- `port` (Number) TACACS Port
- `port_variable` (String) Variable name
- `secret_key` (String) Set the TACACS server shared type 7 encrypted key
- `secret_key_variable` (String) Variable name
- `timeout` (Number) Configure how long to wait for replies from the TACACS server
- `timeout_variable` (String) Variable name



<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `name` (String) Set the username
- `name_variable` (String) Variable name
- `password` (String) Set the user password
- `password_variable` (String) Variable name
- `privilege` (String) Set Privilege Level for this user
- `privilege_variable` (String) Variable name
- `public_keys` (Attributes List) List of RSA public-keys per user (see [below for nested schema](#nestedatt--users--public_keys))

<a id="nestedatt--users--public_keys"></a>
### Nested Schema for `users.public_keys`

Read-Only:

- `key_string` (String) Set the RSA key string
- `key_type` (String) Only RSA is supported
- `key_type_variable` (String) Variable name