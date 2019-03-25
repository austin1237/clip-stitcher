# ---------------------------------------------------------------------------------------------------------------------
# SET UP THE REMOTE BACKEND https://www.terraform.io/docs/state/remote.html
# RUN terraform init and terraform apply in the remote-state dir to create this.
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = "> 0.11.0"

  backend "s3" {}
}

provider "aws" {
  version = "2.3.0"
  region  = "${var.region}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE VPC FOR FARGATE SERVICE
# ---------------------------------------------------------------------------------------------------------------------

module "vpc" {
  source = "./vpc"
  name   = "clipstitcher-${var.ENV}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE The SNS/SQS que and subscriptions
# ---------------------------------------------------------------------------------------------------------------------

module "clip-slugs-que" {
  source         = "./sqs-queue"
  sqs_queue_name = "clip-slugs-sqs-${var.ENV}"
  lambda_arn     = "${module.clipscraper.lambda_arn}"
  lambda_timeout = "${module.clipscraper.timeout}"
  archiver_arn =   "${module.clipslugs-archiver.lambda_arn}"
}

module "clip-links-que" {
  source         = "./sqs-and-sns-queue"
  sqs_queue_name = "clip-links-sqs-${var.ENV}"
  sns_topic_name = "clip-links-sns-${var.ENV}"
  lambda_arn     = "${module.fargaterunner.lambda_arn}"
  archiver_arn =   "${module.cliplinks-archiver.lambda_arn}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE ECS/FARGATE CLUSTER/TASK
# ---------------------------------------------------------------------------------------------------------------------

module "clipstitcher" {
  source = "./fargate"

  name           = "clipstitcher-${var.ENV}"
  subnet_id      = "${module.vpc.subnet_id}"
  que_policy     = "${module.clip-links-que.consumer_policy}"
  image          = "${var.docker_image}"
  docker_version = "${var.DOCKER_VER}"
  cpu            = 2048
  memory         = 4096
  desired_count  = 0

  num_env_vars = 3
  env_vars     = "${map("YOUTUBE_AUTH", "${var.YOUTUBE_AUTH}", "APP_ENV","${var.ENV}", "CONSUMER_URL", "${module.clip-links-que.consumer_url}")}"
}

# ---------------------------------------------------------------------------------------------------------------------
# LAMBDAS
# ---------------------------------------------------------------------------------------------------------------------

module "clipfinder" {
  source         = "./lambda"
  zip_location   = "../../clipfinder/clipfinder.zip"
  name           = "clipfinder-${var.ENV}"
  policy_count   = 1
  iam_policy_arn = ["${module.clip-slugs-que.producer_policy}"]
  handler        = "clipfinder"
  run_time       = "go1.x"
  timeout        = 300
  memory_size    = 1024

  env_vars = {
    TWITCH_CLIENT_ID    = "${var.TWITCH_CLIENT_ID}"
    TWITCH_CHANNEL_NAME = "${var.TWITCH_CHANNEL_NAME}"
    PRODUCER_URL        = "${module.clip-slugs-que.sqs_url}"
  }
}

module "clipslugs-archiver" {
  source         = "./lambda"
  zip_location   = "../../archiver/archiver.zip"
  name           = "clipslugs-archiver-${var.ENV}"
  policy_count   = 2
  iam_policy_arn = ["${module.failed_message_db.producer_policy}", "${module.clip-slugs-que.dead_letter_consumer_policy}"]
  handler        = "index.handler"
  run_time       = "nodejs8.10"

  env_vars = {
    CONSUMER_URL = "${module.clip-slugs-que.dead_letter_url}"
    DB_TABLE = "${module.failed_message_db.table_name}"
  }
}

module "cliplinks-archiver" {
  source         = "./lambda"
  zip_location   = "../../archiver/archiver.zip"
  name           = "cliplinks-archiver-${var.ENV}"
  policy_count   = 2
  iam_policy_arn = ["${module.failed_message_db.producer_policy}", "${module.clip-links-que.dead_letter_consumer_policy}"]
  handler        = "index.handler"
  run_time       = "nodejs8.10"

  env_vars = {
    CONSUMER_URL = "${module.clip-links-que.dead_letter_url}"
    DB_TABLE = "${module.failed_message_db.table_name}"
  }
}

module "clipscraper" {
  source         = "./lambda"
  zip_location   = "../../clipscraper/clipscraper.zip"
  name           = "clipscraper-${var.ENV}"
  policy_count   = 2
  iam_policy_arn = ["${module.clip-slugs-que.consumer_policy}", "${module.clip-links-que.producer_policy}"]
  handler        = "index.handler"
  run_time       = "nodejs8.10"
  memory_size    = 1024
  timeout        = 180

  env_vars = {
    PRODUCER_ARN = "${module.clip-links-que.producer_arn}"
  }
}

module "fargaterunner" {
  source         = "./lambda"
  zip_location   = "../../fargaterunner/fargaterunner.zip"
  name           = "fargaterunner-${var.ENV}"
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
  name                 = "clipfinder-${var.ENV}-trigger"
  lambda_function_name = "clipfinder-${var.ENV}"
  description          = "The timed trigger for clipfinder-${var.ENV}"
  lambda_arn           = "${module.clipfinder.lambda_arn}"
}


# ---------------------------------------------------------------------------------------------------------------------
# DynamoDb Table that will store all messages sent to a dead-letter
# ---------------------------------------------------------------------------------------------------------------------
module "failed_message_db" {
   source = "./dynamodb"
   table_name = "FailedMessages-${var.ENV}"
   hash_key = "QueName"
   range_key = "MessageID"
}