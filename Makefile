NAME   := austin1237/clip-stitcher
TAG    := $$(git rev-parse HEAD)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest
DEV    := ${NAME}:dev

build_prod:
	@docker build -t ${IMG} ./clipstitcher
	@docker tag ${IMG} ${LATEST}

build_dev:
	@docker build -t ${IMG}_dev ./clipstitcher
	@docker tag ${IMG}_dev ${DEV}

hub_push:
	@docker push ${NAME}


