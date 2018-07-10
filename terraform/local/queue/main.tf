resource "aws_sns_topic" "clip_topic" {
  name = "${var.sns_topic_name}"
}

resource "aws_sqs_queue" "clip_que" {
  name = "${var.sqs_queue_name}"
}

resource "aws_sns_topic_subscription" "user_updates_sqs_target" {
  topic_arn = "${aws_sns_topic.clip_topic.arn}"
  protocol  = "sqs"
  endpoint  = "${aws_sqs_queue.clip_que.arn}"
}
