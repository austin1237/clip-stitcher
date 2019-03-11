variable "sqs_queue_name" {
  description = "name of the sqs queue"
}

variable "sns_topic_name" {
  description = "name of the sns topic"
}

variable "lambda_arn" {
  description = "arn of the lambda that sns will invoke"
}

variable "archiver_arn" {
  description = "arn of the archiver lambda that will consume the dead-letter que"
}
