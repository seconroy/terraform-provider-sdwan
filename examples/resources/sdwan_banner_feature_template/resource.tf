resource "sdwan_banner_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  login_banner = "My login banner"
  motd_banner  = "My motd banner"
}
