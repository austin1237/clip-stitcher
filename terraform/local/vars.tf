variable "region" {
  description = "The region where to deploy this code (e.g. us-east-1)."
  default     = "us-east-1"
}

variable "env" {
  description = "Name of the environment example dev/prod"
  default     = "local"
}
