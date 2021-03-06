resource "aws_iam_role_policy" "lambda_role_policy" {
  name   = "lambda_role_policy_${var.name}"
  role   = "${aws_iam_role.iam_for_lambda.id}"
  policy = "${data.aws_iam_policy_document.ecs_service_policy.json}"
}

data "aws_iam_policy_document" "ecs_service_policy" {
  statement {
    effect    = "Allow"
    resources = ["*"]

    actions = [
      "ecs:RunTask",
    ]
  }

  statement {
    actions   = ["iam:PassRole"]
    effect    = "Allow"
    resources = ["*"]

    condition {
      test     = "StringLike"
      variable = "iam:PassedToService"
      values   = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda_${var.name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    },
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }   
  ]
}
EOF
}

resource "aws_lambda_function" "fargate_lambda" {
  filename         = "../../lambdas/fargaterunner/fargaterunner.zip"
  function_name    = "${var.name}"
  role             = "${aws_iam_role.iam_for_lambda.arn}"
  handler          = "main"
  source_code_hash = "${base64sha256(file("../../lambdas/fargaterunner/fargaterunner.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      TASK_ARN    = "${var.task_arn}"
      CLUSTER_ARN = "${var.cluster_arn}"
      SUBNET_ID   = "${var.subnet_id}"
    }
  }
}
