terraform {
  required_version = "> 0.11.0"
}

provider "aws" {
  version                     = "1.14"
  region                      = "${var.region}"
  access_key                  = "anaccesskey"
  secret_key                  = "asecretkey"
  skip_credentials_validation = true
  skip_metadata_api_check     = true

  endpoints {
    sqs = "http://localstack:4576"
    sns = "http://localstack:4575"
  }
}

module "clip-slugs-que" {
  source         = "./queue"
  sqs_queue_name = "clip-slugs-${var.env}"
  sns_topic_name = "clip-slugs-${var.env}"
}

module "clip-links-que" {
  source         = "./queue"
  sqs_queue_name = "clip-links-${var.env}"
  sns_topic_name = "clip-links-${var.env}"
}
