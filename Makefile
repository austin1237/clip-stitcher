NAME   := austin1237/clip-stitcher
TAG    := $$(git rev-parse HEAD)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest
DEV    := ${NAME}:dev

.PHONY: build_prod
build_prod:
	@docker build -t ${IMG} ./clipstitcher
	@docker tag ${IMG} ${LATEST}

.PHONY: build_dev
build_dev:
	@docker build -t ${IMG}_dev ./clipstitcher
	@docker tag ${IMG}_dev ${DEV}

.PHONY: hub_push
hub_push:
	@docker push ${NAME}

.PHONY: produce_message
produce_message: 
	docker-compose run clipfinder && docker-compose run lambdarunner clipfinder


.PHONY: consume_message
consume_message: 
	docker-compose run clipstitcher go run main.go

