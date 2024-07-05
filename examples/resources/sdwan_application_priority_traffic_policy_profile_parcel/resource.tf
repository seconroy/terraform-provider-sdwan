resource "sdwan_application_priority_traffic_policy_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action     = "accept"
  vpn                = ["1"]
  target_direction   = "all"
  sequences = [
    {
      sequence_id = 1
      name        = "RULE_1"
      protocol    = "ipv4"
      base_action = "accept"
      matches = [
        {
          dscp          = 1
          packet_length = "123"
          protocol      = ["2"]
          tcp           = "gre"
          traffic_to    = "core"
        }
      ]
      actions = [
        {
          counter = "COUNTER_1"
          log     = false
          sla_class = [
            {
              preferred_color = ["default"]
            }
          ]
          sets = [
            {
              dscp                          = 1
              local_tloc_list_color         = ["default"]
              local_tloc_restrict           = "false"
              local_tloc_list_encapsulation = "gre"
              tloc_ip                       = "1.2.3.4"
              tloc_color                    = ["default"]
              tloc_encapsulation            = "gre"
              service_type                  = "FW"
              service_color                 = ["default"]
              service_encapsulation         = "ipsec"
              service_tloc_ip               = "1.2.3.4"
              next_hop                      = "1.2.3.4"
            }
          ]
          redirect_dns_field  = "redirectDns"
          redirect_dns_value  = "umbrella"
          nat_pool            = 2
          nat_vpn             = false
          nat_fallback        = false
          fallback_to_routing = true
        }
      ]
    }
  ]
}