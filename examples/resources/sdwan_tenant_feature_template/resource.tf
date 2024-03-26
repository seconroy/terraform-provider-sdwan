resource "sdwan_tenant_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  tenant_organization_name = [
    {
      organization_name    = "example_org"
      tier_name            = "example"
      maximum_vpns_allowed = 123
    }
  ]
}
