output "producer_arn" {
  value = "${aws_sns_topic.producer.arn}"
}

output "producer_policy" {
  value = "${aws_iam_policy.producer_policy.arn}"
}

# id is url for sqs
output "consumer_url" {
  value = "${aws_sqs_queue.consumer.id}"
}

output "consumer_policy" {
  value = "${aws_iam_policy.consumer_policy.arn}"
}
