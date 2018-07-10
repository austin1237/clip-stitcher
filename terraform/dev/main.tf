# ---------------------------------------------------------------------------------------------------------------------
# SET UP THE REMOTE BACKEND https://www.terraform.io/docs/state/remote.html
# RUN terraform init and terraform apply in the remote-state dir to create this.
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = "> 0.11.0"

  backend "s3" {
    bucket         = "austin1237-clipstitcher-state-dev"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = "true"
    dynamodb_table = "clipstitcher-state-lock-dev"
  }
}

provider "aws" {
  version = "1.22"
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
# CREATE The SNS/SQS que and subscriptions
# ---------------------------------------------------------------------------------------------------------------------

module "clip-slugs-que" {
  source         = "./queue"
  sqs_queue_name = "clip-slugs-sqs-${var.env}"
  sns_topic_name = "clip-slugs-sns-${var.env}"
  lambda_arn     = "${module.clipscraper.lambda_arn}"
}

module "clip-links-que" {
  source         = "./queue"
  sqs_queue_name = "clip-links-sqs-${var.env}"
  sns_topic_name = "clip-links-sns-${var.env}"
  lambda_arn     = "${module.fargaterunner.lambda_arn}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE ECS/FARGATE CLUSTER/TASK
# ---------------------------------------------------------------------------------------------------------------------

module "clipstitcher" {
  source = "./fargate"

  name           = "clipstitcher-${var.env}"
  subnet_id      = "${module.vpc.subnet_id}"
  que_policy     = "${module.clip-links-que.consumer_policy}"
  image          = "${var.docker_image}"
  docker_version = "${var.docker_version}"
  cpu            = 1024
  memory         = 2048
  desired_count  = 0

  num_env_vars = 3
  env_vars     = "${map("YOUTUBE_AUTH", "${var.YOUTUBE_AUTH_DEV}", "APP_ENV","${var.env}", "CONSUMER_URL", "${module.clip-links-que.consumer_url}")}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE LAMBDAS
# ---------------------------------------------------------------------------------------------------------------------

module "clipfinder" {
  source         = "./lambda"
  zip_location   = "../../clipfinder/clipfinder.zip"
  name           = "clipfinder-${var.env}"
  policy_count   = 1
  iam_policy_arn = ["${module.clip-slugs-que.producer_policy}"]
  handler        = "clipfinder"
  run_time       = "go1.x"
  timeout        = 100

  env_vars = {
    TWITCH_CLIENT_ID    = "${var.TWITCH_CLIENT_ID_DEV}"
    TWITCH_CHANNEL_NAME = "${var.TWITCH_CHANNEL_NAME_DEV}"
    PRODUCER_ARN        = "${module.clip-slugs-que.producer_arn}"
  }
}

module "clipscraper" {
  source         = "./lambda"
  zip_location   = "../../clipscraper/clipscraper.zip"
  name           = "clipscraper-${var.env}"
  policy_count   = 2
  iam_policy_arn = ["${module.clip-slugs-que.consumer_policy}", "${module.clip-links-que.producer_policy}"]
  handler        = "index.handler"
  run_time       = "nodejs8.10"
  memory_size    = 512
  timeout        = 100

  env_vars = {
    PRODUCER_ARN = "${module.clip-links-que.producer_arn}"
    CONSUMER_URL = "${module.clip-slugs-que.consumer_url}"
  }
}

module "fargaterunner" {
  source         = "./lambda"
  zip_location   = "../../fargaterunner/fargaterunner.zip"
  name           = "fargaterunner-${var.env}"
  policy_count   = 1
  iam_policy_arn = ["${module.clipstitcher.lambda_launch_policy}"]
  handler        = "fargaterunner"
  run_time       = "go1.x"
  timeout        = 100

  env_vars = {
    TASK_ARN    = "${module.clipstitcher.task_arn}"
    CLUSTER_ARN = "${module.clipstitcher.cluster_arn}"
    SUBNET_ID   = "${module.vpc.subnet_id}"
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE CLOUDWATCH TIMED LAMBDA EVENT
# ---------------------------------------------------------------------------------------------------------------------
module "timed-lambda" {
  source               = "./cloudwatch-lambda-trigger"
  start_time           = "cron(30 10 * * ? *)"
  name                 = "clipfinder-${var.env}-trigger"
  lambda_function_name = "clipfinder-${var.env}"
  description          = "The timed trigger for clipfinder-${var.env}"
  lambda_arn           = "${module.clipfinder.lambda_arn}"
}
