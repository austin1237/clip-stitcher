variable "name" {
  description = "The name of the lamda function"
}

variable "cluster_arn" {
  description = "Arn of the ECS cluster that holds the task"
}

variable "task_arn" {
  description = "Arn of the Task to be run"
}

variable "subnet_id" {
  description = "id of the subnet_id for the fargate task"
}
