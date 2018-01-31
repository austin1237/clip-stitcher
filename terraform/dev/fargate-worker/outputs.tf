output "task_arn" {
  value = "${aws_ecs_task_definition.task.arn}"
}

output "service_id" {
  value = "${aws_ecs_service.service.id}"
}