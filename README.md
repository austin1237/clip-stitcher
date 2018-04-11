# clip-stitcher
[![CircleCI](https://circleci.com/gh/austin1237/clip-stitcher.svg?style=svg)](https://circleci.com/gh/austin1237/clip-stitcher)
## Prerequisites
You must have the following installed/configured for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)



##clipstitcher
### Development Environment
To spin up the development environment run the following command from the root level of this repo.

```bash
docker-compose -f clipstitcher-compose.yml up
```

##lambdas
### Development Environment
To spin up the development environment run the following command from the root level of this repo.

```bash
docker-compose -f lambdas-compose.yml up
```

## Terraform
Deployment currently uses [Terraform](https://www.terraform.io/) and AWS's [FARGATE](https://aws.amazon.com/fargate/)

## Setting up remote state
Terraform has a concept called [remote state](https://www.terraform.io/docs/state/remote.html) which ensures the state of your infrastructure to be in sync for mutiple team members.

This project **requires** this feature to be configured. To configure **USE THE FOLLOWING COMMAND ONCE PER AWS ACCOUNT**.
