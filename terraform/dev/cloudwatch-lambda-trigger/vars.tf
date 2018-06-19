variable "start_time" {
  description = "time the cloudwatch event will trigger the lambda"
}

variable "name" {
  description = "The name of the cloudwatch event"
}

variable "lambda_function_name" {
  description = "name of the lambda function that will be invoked"
}

variable "description" {
  description = "description for the cloudwatch event"
}

variable "lambda_arn" {
  description = "Arn of the lambda function that will be invoked"
}
