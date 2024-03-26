resource "sdwan_ntp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  ntp_servers = [
    {
      hostname_ip_address = "example"
      authentication_key  = 12345
      vpn_id              = 1
      version             = 3
      source_interface    = "18"
      prefer              = false
    }
  ]
  authentication = [
    {
      md5_authentication_key_id = 23456
      authentication_value      = "MyPassword"
    }
  ]
  trusted_keys = [1234]
}
