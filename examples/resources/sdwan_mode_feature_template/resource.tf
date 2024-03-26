resource "sdwan_mode_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  bay_1        = "100"
  bay_2        = "40"
}
