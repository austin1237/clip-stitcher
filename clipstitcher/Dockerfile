FROM golang:1.10
RUN apt-get update

# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

#Install ffmpeg
RUN apt-get install xz-utils
RUN mkdir /usr/local/bin/ffmpeg
RUN wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-64bit-static.tar.xz -P /usr/local/bin/ffmpeg --no-check-certificate
RUN tar xf /usr/local/bin/ffmpeg/ffmpeg-release-64bit-static.tar.xz -C /usr/local/bin/ffmpeg
RUN mv  -v /usr/local/bin/ffmpeg/ffmpeg-4.0.2-64bit-static/* /usr/local/bin/ffmpeg/

RUN ln -s /usr/local/bin/ffmpeg/ffmpeg /usr/bin/ffmpeg
RUN ln -s /usr/local/bin/ffmpeg/ffprobe /usr/bin/ffprobe

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