resource "google_compute_instance" "tfer--instance-1" {
  boot_disk {
    auto_delete = "true"
    device_name = "instance-1"

    initialize_params {
      image = "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-9-stretch-v20191014"
      size  = "10"
      type  = "pd-standard"
    }

    mode   = "READ_WRITE"
    source = "https://www.googleapis.com/compute/v1/projects/work1111/zones/asia-east1-b/disks/instance-1"
  }

  can_ip_forward      = "false"
  deletion_protection = "false"
  machine_type        = "n1-standard-1"
  name                = "instance-1"

  network_interface {
    access_config {
      nat_ip       = "130.211.240.103"
      network_tier = "PREMIUM"
    }

    network            = "https://www.googleapis.com/compute/v1/projects/work1111/global/networks/work-gcp"
    network_ip         = "10.1.0.3"
    subnetwork         = "https://www.googleapis.com/compute/v1/projects/work1111/regions/asia-east1/subnetworks/work-gcp-subnet-1"
    subnetwork_project = "work1111"
  }

  project = "work1111"

  scheduling {
    automatic_restart   = "true"
    on_host_maintenance = "MIGRATE"
    preemptible         = "false"
  }

  service_account {
    email  = "219308425897-compute@developer.gserviceaccount.com"
    scopes = ["https://www.googleapis.com/auth/devstorage.read_only", "https://www.googleapis.com/auth/service.management.readonly", "https://www.googleapis.com/auth/trace.append", "https://www.googleapis.com/auth/monitoring.write", "https://www.googleapis.com/auth/logging.write", "https://www.googleapis.com/auth/servicecontrol"]
  }

  zone = "asia-east1-b"
}

resource "google_compute_instance" "tfer--instance-2" {
  boot_disk {
    auto_delete = "true"
    device_name = "instance-2"

    initialize_params {
      image = "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-9-stretch-v20191014"
      size  = "10"
      type  = "pd-standard"
    }

    mode   = "READ_WRITE"
    source = "https://www.googleapis.com/compute/v1/projects/work1111/zones/asia-east1-b/disks/instance-2"
  }

  can_ip_forward      = "false"
  deletion_protection = "false"
  machine_type        = "n1-standard-1"
  name                = "instance-2"

  network_interface {
    access_config {
      nat_ip       = "34.80.134.165"
      network_tier = "PREMIUM"
    }

    network            = "https://www.googleapis.com/compute/v1/projects/work1111/global/networks/work-gcp"
    network_ip         = "10.1.1.2"
    subnetwork         = "https://www.googleapis.com/compute/v1/projects/work1111/regions/asia-east1/subnetworks/work-gcp-subnet-2"
    subnetwork_project = "work1111"
  }

  project = "work1111"

  scheduling {
    automatic_restart   = "true"
    on_host_maintenance = "MIGRATE"
    preemptible         = "false"
  }

  service_account {
    email  = "219308425897-compute@developer.gserviceaccount.com"
    scopes = ["https://www.googleapis.com/auth/logging.write", "https://www.googleapis.com/auth/trace.append", "https://www.googleapis.com/auth/service.management.readonly", "https://www.googleapis.com/auth/devstorage.read_only", "https://www.googleapis.com/auth/servicecontrol", "https://www.googleapis.com/auth/monitoring.write"]
  }

  zone = "asia-east1-b"
}
