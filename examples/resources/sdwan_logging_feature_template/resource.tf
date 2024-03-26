resource "sdwan_logging_feature_template" "example" {
  name              = "Example"
  description       = "My Example"
  device_types      = ["vedge-C8000V"]
  enable            = true
  maximum_file_size = 18
  rotation          = 10
  priority          = "information"
  remote_hosts = [
    {
      hostname_ipv4_address = "example"
      vpn_id                = 1
      source_interface      = "22"
      priority              = "information"
    }
  ]
  remote_ipv6_hosts = [
    {
      ipv6_hostname_ipv6_address = "example"
      vpn_id                     = 1
      source_interface           = "18"
      priority                   = "alert"
    }
  ]
}
