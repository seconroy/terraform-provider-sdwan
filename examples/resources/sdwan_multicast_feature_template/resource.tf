resource "sdwan_multicast_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  local        = false
  threshold    = 123
}
