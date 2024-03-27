resource "sdwan_igmp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  shutdown     = false
  interface = [
    {
      interface_name = "example"
      static_joins = [
        {
          group_address = "1.2.3.4"
        }
      ]
    }
  ]
}
