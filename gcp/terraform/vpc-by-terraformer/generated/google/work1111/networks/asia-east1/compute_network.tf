resource "google_compute_network" "tfer--default" {
  auto_create_subnetworks         = "true"
  delete_default_routes_on_create = "false"
  description                     = "Default network for the project"
  name                            = "default"
  project                         = "work1111"
  routing_mode                    = "REGIONAL"
}

resource "google_compute_network" "tfer--work-gcp" {
  auto_create_subnetworks         = "false"
  delete_default_routes_on_create = "false"
  name                            = "work-gcp"
  project                         = "work1111"
  routing_mode                    = "REGIONAL"
}
