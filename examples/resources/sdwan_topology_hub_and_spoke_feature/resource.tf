resource "sdwan_topology_hub_and_spoke_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpn                = ["edge_basic_vpn1"]
  hub_sites          = ["site_100"]
  spoke_groups = [
    {
      name  = "sj_11"
      sites = ["site_100"]
      hub_preferences = [
        {
          sites      = ["site_1000"]
          preference = 30
        }
      ]
    }
  ]
}
