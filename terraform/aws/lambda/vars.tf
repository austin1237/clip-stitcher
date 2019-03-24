variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = "list"
}

variable "policy_count" {
  description = "Number of policies in iam_policy_arn"
}

variable "zip_location" {
  description = "path to the ziped lambda"
}

variable "name" {
  description = "The name of the lambda function"
}

variable "assume_role_policy" {
  description = "iam role of the lambda"

  default = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }   
  ]
}
EOF
}

variable "handler" {
  description = "name of the lambdas handler"
}

variable "run_time" {
  description = "run time of the lambda"
}

variable "env_vars" {
  type        = "map"
  description = "run time of the lambda"
}

variable "memory_size" {
  description = "Amount of memory in MB your Lambda Function can use at runtime. CPU is implicitly tied to this."
  default     = 128
}

variable "timeout" {
  description = "The max number of seconds the lambda can run"
  default     = 3
}
