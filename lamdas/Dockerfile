FROM golang:1.9.4

# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN apt-get update
RUN apt-get install zip -y

RUN mkdir -p /go/src/github.com/user/fargaterunner
COPY ./fargaterunner /go/src/github.com/user/fargaterunner
WORKDIR /go/src/github.com/user/fargaterunner

# Install go dependencies
RUN dep ensure


