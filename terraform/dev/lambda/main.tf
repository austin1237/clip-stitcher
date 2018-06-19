# ---------------------------------------------------------------------------------------------------------------------
# Lambda Function
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_lambda_function" "lambda_function" {
  filename      = "${var.zip_location}"
  function_name = "${var.name}"
  role          = "${aws_iam_role.iam_for_lambda.arn}"
  handler       = "${var.handler}"
  timeout       = "${var.timeout}"
  memory_size   = "${var.memory_size}"

  source_code_hash = "${base64sha256(file("${var.zip_location}"))}"
  runtime          = "${var.run_time}"

  environment {
    variables = "${var.env_vars}"
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# IAM roles for lambda
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "iam_for_lambda" {
  name               = "${var.name}-lambda-policy"
  assume_role_policy = "${var.assume_role_policy}"
}

resource "aws_iam_role_policy_attachment" "role-policy-attachment" {
  role = "${aws_iam_role.iam_for_lambda.name}"

  count      = "${var.policy_count}"
  policy_arn = "${var.iam_policy_arn[count.index]}"
}

# ---------------------------------------------------------------------------------------------------------------------
# Sets up Cloudwatch logs for the lambda
# ---------------------------------------------------------------------------------------------------------------------

data "aws_iam_policy_document" "log_policy_document" {
  statement {
    resources = ["arn:aws:logs:*:*:*"]

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:DescribeLogStreams",
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "log_policy" {
  name   = "${var.name}-log-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.log_policy_document.json}"
}

resource "aws_iam_role_policy_attachment" "log-policy-attachment" {
  role       = "${aws_iam_role.iam_for_lambda.name}"
  policy_arn = "${aws_iam_policy.log_policy.arn}"
}
