# ---------------------------------------------------------------------------------------------------------------------
# SET UP THE REMOTE BACKEND https://www.terraform.io/docs/state/remote.html
# RUN terraform init and terraform apply in the remote-state dir to create this.
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = "> 0.11.0"
  # backend "s3" {
  #   bucket     = "clipstitcher-state-dev"
  #   key        = "global/s3/terraform.tfstate"
  #   region     = "us-east-1"
  #   encrypt    = "true"
  #   lock_table = "clipstitcher-state-lock-dev"
  # }
}

provider "aws" {
  region = "${var.region}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE VPC FOR FARGATE SERVICE
# ---------------------------------------------------------------------------------------------------------------------

module "vpc" {
  source = "./vpc"
  name = "clipsticher-${var.env}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE FARGATE WORKER TASK clipstitcher
# ---------------------------------------------------------------------------------------------------------------------

module "clipstitcher" {
  source = "./fargate-worker"

  name = "clipsticher-${var.env}"
  subnet_id = "${module.vpc.private_subnet_id}"

  image = "${var.docker_image}"
  docker_version = "${var.docker_version}"
  cpu = 256
  memory = 512
  desired_count = 1
  
  num_env_vars = 8
  env_vars = "${map("TWITCH_CLIENT_ID", "${var.TWITCH_CLIENT_ID}", "TWITCH_CHANNEL_NAME", "${var.TWITCH_CHANNEL_NAME}", "YOUTUBE_CLIENT_ID", "${var.YOUTUBE_CLIENT_ID}", "YOUTUBE_SECRET", "${var.YOUTUBE_SECRET}", "YOUTUBE_ACCESS_TOKEN", "${var.YOUTUBE_ACCESS_TOKEN}", "YOUTUBE_REFRESH_TOKEN", "${var.YOUTUBE_REFRESH_TOKEN}", "YOUTUBE_EXPIRY", "${var.YOUTUBE_EXPIRY}", "APP_ENV","${var.env}")}"
}




