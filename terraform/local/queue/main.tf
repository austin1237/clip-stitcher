resource "aws_sns_topic" "sns_topic" {
  name = "${var.sns_topic_name}"
}


# maxRecieveCount expect a string instead of a number on localstack
# https://github.com/localstack/localstack/issues/238
resource "aws_sqs_queue" "sqs_que" {
  name = "${var.sqs_queue_name}"
}

resource "aws_sns_topic_subscription" "user_updates_sqs_target" {
  topic_arn = "${aws_sns_topic.sns_topic.arn}"
  protocol  = "sqs"
  endpoint  = "${aws_sqs_queue.sqs_que.arn}"
}
