output "task_arn" {
  value = "${aws_ecs_task_definition.task.arn}"
}

output "cluster_arn" {
  value = "${aws_ecs_cluster.cluster.arn}"
}

output "service_id" {
  value = "${aws_ecs_service.service.id}"
}

output "lambda_launch_policy" {
  value = "${aws_iam_policy.lambda_launch_policy.arn}"
}
