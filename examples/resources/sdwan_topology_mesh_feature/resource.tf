resource "sdwan_topology_mesh_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpn                = ["edge_basic_vpn1"]
  sites              = ["site_100"]
}
