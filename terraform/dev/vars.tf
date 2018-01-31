# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY
# TF_VAR_TWITCH_CLIENT_ID
# TF_VAR_TWITCH_CHANNEL_NAME
# TF_VAR_YOUTUBE_CLIENT_ID
# TF_VAR_YOUTUBE_SECRET
# TF_VAR_YOUTUBE_ACCESS_TOKEN
# TF_VAR_YOUTUBE_REFRESH_TOKEN
# TF_VAR_YOUTUBE_EXPIRY

variable "TWITCH_CLIENT_ID" {}
variable "TWITCH_CHANNEL_NAME" {}
variable "YOUTUBE_CLIENT_ID" {}
variable "YOUTUBE_SECRET" {}
variable "YOUTUBE_ACCESS_TOKEN" {}
variable "YOUTUBE_REFRESH_TOKEN" {}
variable "YOUTUBE_EXPIRY" {}
# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL MODULE PARAMETERS
# These variables have defaults, but may be overridden by the operator.
# ---------------------------------------------------------------------------------------------------------------------

variable "region" {
  description = "The region where to deploy this code (e.g. us-east-1)."
  default = "us-east-1"
}

variable "docker_image" {
  description = "The name of the Docker image to deploy for the gifbot (e.g. austin1237/gifbot)"
  default = "austin1237/clip-stitcher"
}

variable "docker_version" {
  description = "The version (i.e. tag) of the Docker container to deploy for the gifbot (e.g. latest, 12345)"
  default = "dev"
}

variable "env" {
  description = "Name of the environment example dev/prod"
  default = "dev"
}