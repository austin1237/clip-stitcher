variable "aggregator_arn" {
  description = "arn of the lambda function that will send the logs to a log aggregation service"
}

variable "log_group_count" {
  description = "Number of cloudwatch log groups that the aggregator will sub to"
}

variable "log_group_name" {
  description = "The names of the cloudwatch log groups that the aggregator will sub to"
  type        = "list"
}

variable "region" {
  description = "The region where the cloudwatch logs"
}