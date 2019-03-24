FROM austin1237/golang-ffmpeg:1.10-4.1

RUN apt-get update

# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Add Project Source Code
RUN mkdir -p /go/src/github.com/user/clipstitcher
COPY . /go/src/github.com/user/clipstitcher
WORKDIR /go/src/github.com/user/clipstitcher

# Install go dependencies
RUN dep ensure 

RUN go install .

CMD if [ ${APP_ENV} = "local" ]; \
	then \
	dep ensure && \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
    clipstitcher; \
	fi