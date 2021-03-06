provider "aws" {
  version = "1.14"
  region  = "us-east-1"
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "austin1237-clipstitcher-state-${var.ENVIRONMENT}"

  versioning {
    enabled = true
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_dynamodb_table" "terraform_state_lock" {
  name           = "clipstitcher-state-lock-${var.ENVIRONMENT}"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
