resource "aws_cloudwatch_log_subscription_filter" "log_aggregation" {
  count           = "${var.log_group_count}"
  name            = "log-aggregation-${var.log_group_name[count.index]}"
  log_group_name  = "${var.log_group_name[count.index]}"
  filter_pattern  = ""
  destination_arn = "${var.aggregator_arn}"
  distribution    = "Random"
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id   = "AllowExecutionFromCloudWatch"
  action         = "lambda:InvokeFunction"
  function_name  = "${var.aggregator_arn}"
  principal      = "logs.${var.region}.amazonaws.com"
}