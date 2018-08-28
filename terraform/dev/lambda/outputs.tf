output "lambda_arn" {
  value = "${aws_lambda_function.lambda_function.arn}"
}

output "log_group_name" {
  value = "/aws/lambda/${var.name}"
}
