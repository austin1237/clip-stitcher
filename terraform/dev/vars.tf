# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

variable "TWITCH_CLIENT_ID_DEV" {}
variable "TWITCH_CHANNEL_NAME_DEV" {}
variable "YOUTUBE_AUTH_DEV" {}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL MODULE PARAMETERS
# These variables have defaults, but may be overridden by the operator.
# ---------------------------------------------------------------------------------------------------------------------

variable "region" {
  description = "The region where to deploy this code (e.g. us-east-1)."
  default     = "us-east-1"
}

variable "docker_image" {
  description = "The name of the Docker image to deploy for the gifbot (e.g. austin1237/gifbot)"
  default     = "austin1237/clip-stitcher"
}

variable "docker_version" {
  description = "The version (i.e. tag) of the Docker container to deploy for the gifbot (e.g. latest, 12345)"
  default     = "dev"
}

variable "env" {
  description = "Name of the environment example dev/prod"
  default     = "dev"
}
