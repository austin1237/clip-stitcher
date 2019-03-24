variable "sqs_queue_name" {
  description = "name of the sqs queue"
}

variable "lambda_arn" {
  description = "arn of the lambda that sqs will invoke"
}

variable "lambda_timeout" {
  description = "timeout of lambda_arn"
}

variable "archiver_arn" {
  description = "arn of the archiver lambda that will consume the dead-letter que"
}
