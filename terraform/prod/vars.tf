# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY
# TF_VAR_TWITCH_CLIENT_ID_PROD
# TF_VAR_TWITCH_CHANNEL_NAME_PROD
# TF_VAR_YOUTUBE_CLIENT_ID_PROD
# TF_VAR_YOUTUBE_SECRET_PROD
# TF_VAR_YOUTUBE_ACCESS_TOKEN_PROD
# TF_VAR_YOUTUBE_REFRESH_TOKEN_PROD
# TF_VAR_YOUTUBE_EXPIRY_PROD

variable "TWITCH_CLIENT_ID_PROD" {}
variable "TWITCH_CHANNEL_NAME_PROD" {}
variable "YOUTUBE_AUTH_PROD" {}

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
  default     = "latest"
}

variable "env" {
  description = "Name of the environment example dev/prod"
  default     = "prod"
}
