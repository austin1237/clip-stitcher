output "consumer_policy" {
  value = "${aws_iam_policy.consumer_policy.arn}"
}

output "producer_policy" {
    value = "${aws_iam_policy.producer_policy.arn}"
}

output "table_name" {
    value = "${aws_dynamodb_table.dynamodb_table.id}"
}