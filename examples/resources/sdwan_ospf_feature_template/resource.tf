resource "sdwan_ospf_feature_template" "example" {
  name                           = "Example"
  description                    = "My Example"
  device_types                   = ["vedge-C8000V"]
  router_id                      = "1.2.3.4"
  reference_bandwidth            = 429496
  rfc_1583_compatible            = true
  originate_default_route        = false
  always                         = false
  default_metric                 = 16777
  metric_type                    = "type2"
  distance_for_external_routes   = 110
  distance_for_inter_area_routes = 110
  distance_for_intra_area_routes = 110
  spf_calculation_delay          = 200
  initial_hold_time              = 1000
  max_hold_time                  = 1000
  redistribute_routes = [
    {
      protocol     = "static"
      route_policy = "example"
    }
  ]
  router_lsa = [
    {
      type               = "administrative"
      advertizement_time = 1234
    }
  ]
  route_policies = [
    {
      direction   = "in"
      policy_name = "example"
    }
  ]
  ospf_area = [
    {
      area_number     = 42949672
      stub_no_summary = false
      translate       = "candidate"
      nssa_no_summary = false
      ospf_interface = [
        {
          interface_name             = "interface_example"
          hello_interval             = 10
          dead_interval              = 40
          retransmit_interval        = 5
          interface_cost             = 1234
          designated_router_priority = 1
          ospf_network_type          = "broadcast"
          passive_interface          = false
          authentication_type        = "simple"
          authentication_key         = "example"
          message_digest_key_id      = 1
          message_digest_key         = "MyMD5Key"
        }
      ]
      summarize_routes = [
        {
          address      = "1.2.3.4/24"
          cost         = 167215
          no_advertise = false
        }
      ]
    }
  ]
}
