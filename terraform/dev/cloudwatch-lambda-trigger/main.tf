# Cloudwatch timed trigger
# ---------------------------------------------------------------------------------------------------------------------
resource "aws_cloudwatch_event_rule" "once_a_day" {
  name                = "${var.name}"
  description         = "${var.description}"
  schedule_expression = "${var.start_time}"
}

resource "aws_cloudwatch_event_target" "check_once_a_day" {
  rule = "${aws_cloudwatch_event_rule.once_a_day.name}"
  arn  = "${var.lambda_arn}"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_once_a_day" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "${var.lambda_function_name}"
  principal     = "events.amazonaws.com"
  source_arn    = "${aws_cloudwatch_event_rule.once_a_day.arn}"
}
