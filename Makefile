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

deploy_dev:
	cd terraform/dev && terraform init
	cd terraform/dev && terraform apply

deploy_prod:
	cd terraform/prod && terraform init
	cd terraform/prod && terraform apply

init_terraform_remote_state:
	cd terraform/dev/remote-state && terraform init
	cd terraform/dev/remote-state && terraform apply
	echo "dev storage deployed"
	cd terraform/prod/remote-state && terraform init
	cd terraform/prod/remote-state && terraform apply
	echo "prod storage deployed"
	echo "terraform remote state management can now be used"


