# ---------------------------------------------------------------------------------------------------------------------
# SNS/ Where producers send data
# ---------------------------------------------------------------------------------------------------------------------

# resource "aws_sns_topic" "producer" {
#   name = "${var.sns_topic_name}"
# }

# data "aws_iam_policy_document" "producer_policy_document" {
#   statement {
#     resources = ["${aws_sns_topic.producer.arn}"]

#     actions = [
#       "SNS:Publish",
#     ]

#     effect = "Allow"
#   }
# }

# resource "aws_iam_policy" "producer_policy" {
#   name   = "${var.sns_topic_name}-policy"
#   path   = "/"
#   policy = "${data.aws_iam_policy_document.producer_policy_document.json}"
# }

# ---------------------------------------------------------------------------------------------------------------------
# SQS/ Where consumers get data from
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_sqs_queue" "consumer" {
  name = "${var.sqs_queue_name}"
  redrive_policy            = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.consumer_dead_letter.arn}\",\"maxReceiveCount\":1}"
  visibility_timeout_seconds = "${var.lambda_timeout}"
}

# resource "aws_sqs_queue_policy" "aws_sqs_queue_policy" {
#   queue_url = "${aws_sqs_queue.consumer.id}"

#   policy = <<POLICY
# {
#   "Version": "2012-10-17",
#   "Statement": [{
#     "Effect": "Allow",
#     "Principal": "*",
#     "Action": "sqs:SendMessage",
#     "Resource": "${aws_sqs_queue.consumer.arn}",
#     "Condition": {
#       "ArnEquals": {
#         "aws:SourceArn": ["${aws_sns_topic.producer.arn}"]
#       }
#     }
#   }
#   ]
# }
# POLICY
# }

data "aws_iam_policy_document" "producer_policy_document" {
  statement {
    resources = ["${aws_sqs_queue.consumer.arn}"]

    actions = [
      "SQS:SendMessage"
    ]

    effect = "Allow"
  }
}

data "aws_iam_policy_document" "consumer_policy_document" {
  statement {
    resources = ["${aws_sqs_queue.consumer.arn}"]

    actions = [
      "SQS:ReceiveMessage",
      "SQS:DeleteMessage",
      "SQS:GetQueueAttributes"
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "consumer_policy" {
  name   = "${var.sqs_queue_name}-consumer-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.consumer_policy_document.json}"
}

resource "aws_iam_policy" "producer_policy" {
  name   = "${var.sqs_queue_name}-producer-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.producer_policy_document.json}"
}

# ---------------------------------------------------------------------------------------------------------------------
# SQS Dead Letter
# ---------------------------------------------------------------------------------------------------------------------
resource "aws_sqs_queue" "consumer_dead_letter" {
  name = "${var.sqs_queue_name}-dead-letter"
}

data "aws_iam_policy_document" "consumer_dead_letter_policy_document" {
  statement {
    resources = ["${aws_sqs_queue.consumer_dead_letter.arn}"]

    actions = [
      "SQS:ReceiveMessage",
      "SQS:DeleteMessage",
      "SQS:GetQueueAttributes"
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "consumer_dead_letter_policy" {
  name   = "${var.sqs_queue_name}-dead-letter-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.consumer_dead_letter_policy_document.json}"
}
# ---------------------------------------------------------------------------------------------------------------------
# SQS -> Lambda Event Subscription
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_lambda_event_source_mapping" "dead_letter_sub" {
  batch_size        = 1
  event_source_arn  = "${aws_sqs_queue.consumer_dead_letter.arn}"
  enabled           = true
  function_name     = "${var.archiver_arn}"
}

resource "aws_lambda_event_source_mapping" "sqs_to_lambda" {
  batch_size        = 1
  event_source_arn = "${aws_sqs_queue.consumer.arn}"
  enabled           = true
  function_name    = "${var.lambda_arn}"
}

# ---------------------------------------------------------------------------------------------------------------------
# SNS Subscriptions
# ---------------------------------------------------------------------------------------------------------------------

# resource "aws_sns_topic_subscription" "sns_to_sqs_subscription" {
#   topic_arn = "${aws_sns_topic.producer.arn}"
#   protocol  = "sqs"
#   endpoint  = "${aws_sqs_queue.consumer.arn}"
# }

# resource "aws_sns_topic_subscription" "sns_to_lambda_subscription" {
#   topic_arn = "${aws_sns_topic.producer.arn}"
#   protocol  = "lambda"
#   endpoint  = "${var.lambda_arn}"
# }

# resource "aws_lambda_permission" "with_sns" {
#   statement_id  = "AllowExecutionFromSNS"
#   action        = "lambda:InvokeFunction"
#   function_name = "${var.lambda_arn}"
#   principal     = "sns.amazonaws.com"
#   source_arn    = "${aws_sns_topic.producer.arn}"
# }
