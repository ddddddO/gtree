resource "google_compute_subnetwork" "tfer--default" {
  enable_flow_logs         = "false"
  ip_cidr_range            = "10.140.0.0/20"
  name                     = "default"
  network                  = "https://www.googleapis.com/compute/v1/projects/work1111/global/networks/default"
  private_ip_google_access = "false"
  project                  = "work1111"
  region                   = "asia-east1"
}

resource "google_compute_subnetwork" "tfer--work-gcp-subnet-1" {
  enable_flow_logs         = "false"
  ip_cidr_range            = "10.1.0.0/24"
  name                     = "work-gcp-subnet-1"
  network                  = "https://www.googleapis.com/compute/v1/projects/work1111/global/networks/work-gcp"
  private_ip_google_access = "false"
  project                  = "work1111"
  region                   = "asia-east1"
}

resource "google_compute_subnetwork" "tfer--work-gcp-subnet-2" {
  enable_flow_logs         = "false"
  ip_cidr_range            = "10.1.1.0/24"
  name                     = "work-gcp-subnet-2"
  network                  = "https://www.googleapis.com/compute/v1/projects/work1111/global/networks/work-gcp"
  private_ip_google_access = "false"
  project                  = "work1111"
  region                   = "asia-east1"
}
