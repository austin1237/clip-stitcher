FROM golang:1.9

# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN apt-get update
RUN apt-get install ffmpeg -y

RUN mkdir -p /go/src/github.com/user/clipstitcher
COPY . /go/src/github.com/user/clipstitcher
WORKDIR /go/src/github.com/user/clipstitcher

# Install go dependencies
RUN dep ensure 

# Couldn't get this to load with dep for some reason
RUN go get "github.com/stretchr/testify"

RUN go install .

CMD if [ ${APP_ENV} != dev ]; \
	then \
	clipstitcher; \
	else \
    dep ensure && \
	go get github.com/pilu/fresh && \
	fresh; \
	fi