output "lambda_arn" {
  value = "${aws_lambda_function.lambda_function.arn}"
}

output "timeout" {
  value = "${aws_lambda_function.lambda_function.timeout}"
}
