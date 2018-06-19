output "lambda_arn" {
  value = "${aws_lambda_function.fargate_lambda.arn}"
}
