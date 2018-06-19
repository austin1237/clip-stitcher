# ---------------------------------------------------------------------------------------------------------------------
# CREATE A CLOUDWATCH LOG GROUP FOR ECS TASKS TO SEND LOGS TO
# ---------------------------------------------------------------------------------------------------------------------
resource "aws_cloudwatch_log_group" "ecs_logs" {
  name = "${var.name}"
}

resource "aws_cloudwatch_log_stream" "ecs_logs_stream" {
  name           = "${var.name}-logs-stream"
  log_group_name = "${aws_cloudwatch_log_group.ecs_logs.name}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN ECS CLUSTER
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_ecs_cluster" "cluster" {
  name = "${var.name}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN ECS SERVICE TO RUN A ECS TASK
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_ecs_service" "service" {
  launch_type     = "FARGATE"
  name            = "${var.name}"
  cluster         = "${aws_ecs_cluster.cluster.id}"
  task_definition = "${aws_ecs_task_definition.task.arn}"
  desired_count   = "${var.desired_count}"

  network_configuration = [
    {
      assign_public_ip = true
      subnets          = ["${var.subnet_id}"]
    },
  ]
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN ECS TASK TO RUN A DOCKER CONTAINER
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_ecs_task_definition" "task" {
  family                   = "${var.name}"
  memory                   = "${var.memory}"
  cpu                      = "${var.cpu}"
  network_mode             = "awsvpc"
  execution_role_arn       = "${aws_iam_role.task_role.arn}"
  requires_compatibilities = ["FARGATE"]

  container_definitions = <<EOF
  [
    {
      "name": "${var.name}",
      "image": "${var.image}:${var.docker_version}",
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
            "awslogs-group": "${aws_cloudwatch_log_group.ecs_logs.name}",
            "awslogs-region": "us-east-1",
            "awslogs-stream-prefix": "${aws_cloudwatch_log_stream.ecs_logs_stream.name}"
        }
      },
      "environment": [${join(",", data.template_file.env_vars.*.rendered)}]
    }
  ]
  EOF
}

# Convert the environment variables the user passed-in into the format expected for for an ECS Task:
#
# "environment": [
#    {"name": "NAME", "value": "VALUE"},
#    {"name": "NAME", "value": "VALUE"},
#    ...
# ]
#
data "template_file" "env_vars" {
  count = "${var.num_env_vars}"

  template = <<EOF
{"name": "${element(keys(var.env_vars), count.index)}", "value": "${lookup(var.env_vars, element(keys(var.env_vars), count.index))}"}
EOF
}

# ---------------------------------------------------------------------------------------------------------------------
# ATTACH IAM PERMISSIONS TO THE IAM ROLE
# This IAM Policy allows the ECS Service to communicate with CLOUDWATCH
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role_policy" "task_role_policy" {
  name   = "task_role_policy_${var.name}"
  role   = "${aws_iam_role.task_role.id}"
  policy = "${data.aws_iam_policy_document.ecs_service_policy.json}"
}

data "aws_iam_policy_document" "ecs_service_policy" {
  statement {
    effect    = "Allow"
    resources = ["*"]

    actions = [
      "cloudwatch:*",
      "logs:*",
      "ecs:*",
      "ec2:*",
      "SQS:ReceiveMessage",
      "SQS:DeleteMessage",
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

  statement {
    effect    = "Allow"
    actions   = ["iam:CreateServiceLinkedRole"]
    resources = ["*"]

    condition {
      test     = "StringLike"
      variable = "iam:AWSServiceName"
      values   = ["ecs.amazonaws.com", "spot.amazonaws.com", "spotfleet.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "task-attach" {
  role       = "${aws_iam_role.task_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "task_role" {
  name = "task_role_${var.name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# ---------------------------------------------------------------------------------------------------------------------
# This IAM Policy allows a lambda to launch an ecs task
# ---------------------------------------------------------------------------------------------------------------------
# resource "aws_iam_role_policy" "lambda_launch_policy" {
#   name   = "lambda_role_policy_${var.name}"
#   role   = "${aws_iam_role.iam_for_lambda.id}"
#   policy = "${data.aws_iam_policy_document.lambda_launch_document.json}"
# }

resource "aws_iam_policy" "lambda_launch_policy" {
  name   = "${var.name}-lambda-policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.lambda_launch_document.json}"
}

data "aws_iam_policy_document" "lambda_launch_document" {
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

# resource "aws_iam_role" "iam_for_lambda" {
#   name = "iam_for_lambda_${var.name}"


#   assume_role_policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Principal": {
#         "Service": "lambda.amazonaws.com"
#       },
#       "Effect": "Allow",
#       "Sid": ""
#     },
#     {
#       "Sid": "",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "ecs-tasks.amazonaws.com"
#       },
#       "Action": "sts:AssumeRole"
#     }   
#   ]
# }
# EOF
# }

