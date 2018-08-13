resource "aws_dynamodb_table" "dynamodb_table" {
    name           = "${var.table_name}"
    hash_key = "${var.hash_key}"
    range_key = "${var.range_key}"
    read_capacity  = "${var.read_capacity}"
    write_capacity = "${var.write_capacity}"

    attribute {
        name = "${var.hash_key}"
        type = "S"
    }

    attribute {
        name = "${var.range_key}"
        type = "S"
    }
}


# ---------------------------------------------------------------------------------------------------------------------
# IAM Policies to access the DB
# ---------------------------------------------------------------------------------------------------------------------
data "aws_iam_policy_document" "producer_policy_document" {
  statement {
    resources = ["${aws_dynamodb_table.dynamodb_table.arn}"]
    actions = [
      "DYNAMODB:PutItem",
    ]
    effect = "Allow"
  }
}

resource "aws_iam_policy" "producer_policy" {
  name   = "${var.table_name}-producer-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.producer_policy_document.json}"
}

data "aws_iam_policy_document" "consumer_policy_document" {
  statement {
    resources = ["${aws_dynamodb_table.dynamodb_table.arn}"]
    actions = [
      "DYNAMODB:Query",
      "DYNAMODB:Scan",
    ]
    effect = "Allow"
  }
}

resource "aws_iam_policy" "consumer_policy" {
  name   = "${var.table_name}-consumer-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.consumer_policy_document.json}"
}
