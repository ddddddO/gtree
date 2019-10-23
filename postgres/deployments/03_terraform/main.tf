# https://www.terraform.io/docs/providers/google/r/container_cluster.html
# Example Usage - with a separately managed node pool (recommended)

provider "google" {
  project = "work1111"
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_container_cluster" "primary" {
  name     = "gle-pubsub-cluster"
  location = "us-central1"

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "pubsub-node-pool"
  location   = "us-central1"
  cluster    = "${google_container_cluster.primary.name}"
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}
