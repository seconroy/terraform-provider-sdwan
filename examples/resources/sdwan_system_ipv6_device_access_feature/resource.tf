resource "sdwan_system_ipv6_device_access_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "drop"
  sequences = [
    {
      id                 = 1
      name               = "SEQ_1"
      base_action        = "accept"
      device_access_port = 22
    }
  ]
}
