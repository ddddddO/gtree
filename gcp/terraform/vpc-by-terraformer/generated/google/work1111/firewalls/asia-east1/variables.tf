data "terraform_remote_state" "networks" {
  backend = "local"

  config = {
    path = "../../../../../generated/google/work1111/networks/asia-east1/terraform.tfstate"
  }
}
