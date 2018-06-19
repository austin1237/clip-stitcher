# ---------------------------------------------------------------------------------------------------------------------
# SNS/ Where producers send data
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_sns_topic" "producer" {
  name = "${var.sns_topic_name}"
}

data "aws_iam_policy_document" "producer_policy_document" {
  statement {
    resources = ["${aws_sns_topic.producer.arn}"]

    actions = [
      "SNS:Publish",
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "producer_policy" {
  name   = "${var.sns_topic_name}-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.producer_policy_document.json}"
}

# ---------------------------------------------------------------------------------------------------------------------
# SQS/ Where consumers get data from
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_sqs_queue" "consumer" {
  name = "${var.sqs_queue_name}"
}

resource "aws_sqs_queue_policy" "aws_sqs_queue_policy" {
  queue_url = "${aws_sqs_queue.consumer.id}"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": "*",
    "Action": "sqs:SendMessage",
    "Resource": "${aws_sqs_queue.consumer.arn}",
    "Condition": {
      "ArnEquals": {
        "aws:SourceArn": ["${aws_sns_topic.producer.arn}"]
      }
    }
  }
  ]
}
POLICY
}

data "aws_iam_policy_document" "consumer_policy_document" {
  statement {
    resources = ["${aws_sqs_queue.consumer.arn}"]

    actions = [
      "SQS:ReceiveMessage",
      "SQS:DeleteMessage",
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "consumer_policy" {
  name   = "${var.sqs_queue_name}-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.consumer_policy_document.json}"
}

# ---------------------------------------------------------------------------------------------------------------------
# SNS Subscriptions
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_sns_topic_subscription" "sns_to_sqs_subscription" {
  topic_arn = "${aws_sns_topic.producer.arn}"
  protocol  = "sqs"
  endpoint  = "${aws_sqs_queue.consumer.arn}"
}

resource "aws_sns_topic_subscription" "sns_to_lambda_subscription" {
  topic_arn = "${aws_sns_topic.producer.arn}"
  protocol  = "lambda"
  endpoint  = "${var.lambda_arn}"
}

resource "aws_lambda_permission" "with_sns" {
  statement_id  = "AllowExecutionFromSNS"
  action        = "lambda:InvokeFunction"
  function_name = "${var.lambda_arn}"
  principal     = "sns.amazonaws.com"
  source_arn    = "${aws_sns_topic.producer.arn}"
}
