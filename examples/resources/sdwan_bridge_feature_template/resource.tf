resource "sdwan_bridge_feature_template" "example" {
  name                  = "Example"
  description           = "My Example"
  device_types          = ["vedge-C8000V"]
  bridge_name           = "example"
  age_out_time          = 300
  maximum_mac_addresses = 1234
  vlan_id               = 1
  interface_to_bind_to_bridging_domain = [
    {
      interface_name      = "interface_example"
      description         = "My Description"
      native_vlan_support = false
      static_mac_address = [
        {
          mac_address = "0000.00000.0000"
        }
      ]
      shutdown = true
    }
  ]
}
