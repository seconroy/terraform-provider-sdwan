resource "sdwan_pim_feature_template" "example" {
  name          = "Example"
  description   = "My Example"
  device_types  = ["vedge-C8000V"]
  shutdown      = false
  auto_rp       = false
  spt_threshold = 64
  replicator    = "random"
  interface = [
    {
      interface_name      = "example"
      hello_interval      = 1200
      join_prune_interval = 123
    }
  ]
}
