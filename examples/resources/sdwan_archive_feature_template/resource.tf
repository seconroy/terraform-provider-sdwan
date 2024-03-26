resource "sdwan_archive_feature_template" "example" {
  name                           = "Example"
  description                    = "My Example"
  device_types                   = ["vedge-C8000V"]
  path_to_archive_file_directory = "ftp://file-path"
  archival_interval              = 52560
  ssh_key_file                   = "id_rsa"
  vpn_id                         = 1
}
