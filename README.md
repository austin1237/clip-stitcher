# clip-stitcher
[![CircleCI](https://circleci.com/gh/austin1237/clip-stitcher.svg?style=svg)](https://circleci.com/gh/austin1237/clip-stitcher)<br />
A worker process that once a day combines and archives popular stream clips.
![architecture](https://user-images.githubusercontent.com/1394341/44130144-4bf5db72-a009-11e8-9852-ab8fc132f1c0.png)

## Prerequisites
You must have the following installed/configured on your system for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)

## Environment Variables
The following variables need to be set on your local/ci system.
### TWITCH_CLIENT_ID
1. Have a [Twitch](https://www.twitch.tv) account
2. Regisiter a new [Twitch application](https://dev.twitch.tv/dashboard/apps)
3. Set the Oauth redirect url to http://localhost
4. Pick anything from the application categories
5. Copy and save the client id presented as the envoirment variable TWITCH_CLIENT_ID

### TWITCH_CHANNEL_NAME
Name of your channel
### YOUTUBE_AUTH
Use the base64 string returned from [yt-server-oauth](https://github.com/austin1237/yt-server-oauth)

## Development Environment
The development enviroment uses [lambci's](https://github.com/lambci) [docker-lambda](https://github.com/lambci/docker-lambda) image to run any code that would be deployed as a lambda function. To emulate aws's Dynamodb/Sns/Sqs a [localstack](https://github.com/localstack/localstack) image is combined with Terraform.

### Start up
To build the lambdas and spin up the distributed development environment run the following command

```bash
docker-compose -f lambdaBuilder.yml up  && docker-compose up
```

### Tests
To run tests use the following command

```bash
docker-compose -f testRunner.yml up
```

## Deployment
Deployment currently uses [Terraform](https://www.terraform.io/) to set up AWS services.

### Setting up remote state
Terraform has a feature called [remote state](https://www.terraform.io/docs/state/remote.html) which ensures the state of your infrastructure to be in sync for mutiple team members.

This project **requires** this feature to be configured. To configure **USE THE FOLLOWING COMMAND ONCE PER TEAM**.
```bash
docker-compose -f terraform-compose.yml run tf init_remote_state
```

### Manual
The following commands will deploy to dev/prod environment manually.
```bash
docker-compose -f terraform-compose.yml run tf deploy_dev
```

```bash
docker-compose -f terraform-compose.yml run tf deploy_prod
```

### Automation CI/CD
This project uses [CircleCI](https://circleci.com/) [workflows](https://circleci.com/docs/2.0/workflows/) for CI/CD the configuration for this is in .circleci