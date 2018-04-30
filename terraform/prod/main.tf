# ---------------------------------------------------------------------------------------------------------------------
# SET UP THE REMOTE BACKEND https://www.terraform.io/docs/state/remote.html
# RUN terraform init and terraform apply in the remote-state dir to create this.
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = "> 0.11.0"

  backend "s3" {
    bucket         = "austin1237-clipstitcher-state-prod"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = "true"
    dynamodb_table = "clipstitcher-state-lock-prod"
  }
}

provider "aws" {
  version = "1.14"
  region  = "${var.region}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE VPC FOR FARGATE SERVICE
# ---------------------------------------------------------------------------------------------------------------------

module "vpc" {
  source = "./vpc"
  name   = "clipstitcher-${var.env}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE FARGATE WORKER TASK clipstitcher
# ---------------------------------------------------------------------------------------------------------------------

module "clipstitcher" {
  source = "./fargate-worker"

  name      = "clipstitcher-${var.env}"
  subnet_id = "${module.vpc.subnet_id}"

  image          = "${var.docker_image}"
  docker_version = "${var.docker_version}"
  cpu            = 1024
  memory         = 2048
  desired_count  = 0

  num_env_vars = 4
  env_vars     = "${map("TWITCH_CLIENT_ID", "${var.TWITCH_CLIENT_ID_PROD}", "TWITCH_CHANNEL_NAME", "${var.TWITCH_CHANNEL_NAME_PROD}", "YOUTUBE_AUTH", "${var.YOUTUBE_AUTH_PROD}", "APP_ENV","${var.env}")}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE Lambda that will run the fargate worker
# ---------------------------------------------------------------------------------------------------------------------
module "fargate-runner" {
  source      = "./fargate-runner"
  cluster_arn = "${module.clipstitcher.cluster_arn}"
  task_arn    = "${module.clipstitcher.task_arn}"
  name        = "clip-stitcher-runner-${var.env}"
  subnet_id   = "${module.vpc.subnet_id}"
  start_time  = "cron(30 10 * * ? *)"
}
