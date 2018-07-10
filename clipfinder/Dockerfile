FROM lambci/lambda:build-go1.x
COPY . /go/src/github.com/user/clipfinder 
WORKDIR /go/src/github.com/user/clipfinder 

# Install go dependencies
RUN dep ensure