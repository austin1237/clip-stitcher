.PHONEY: destroy_dis
destroy_dis:
	cd distributed && terraform init -input=false
	cd distributed && yes | terraform destroy  -input=false -auto-approve

.PHONY: plan_dis
plan_dis:
	cd distributed && terraform init -input=false
	cd distributed && terraform plan

.PHONY: deploy_dis
deploy_dis:
	cd distributed && terraform init -input=false
	cd distributed && yes | terraform apply  -input=false -auto-approve

.PHONY: plan_local
plan_local:
	cd local && rm -f terraform.tfstate 
	cd local && terraform init -input=false
	cd local && terraform plan

.PHONY: init_local
init_local:
	cd local && terraform init -input=false

.PHONY: deploy_local
deploy_local:
	cd local && terraform init -input=false
	dockerize -wait tcp://localstack:4575 -wait tcp://localstack:4576 -wait tcp://localstack:4569 -timeout 10m
	cd local && rm -f terraform.tfstate
	cd local && yes | terraform apply  -input=false -auto-approve

.PHONY: deploy_dev
deploy_dev:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-dev" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-dev" -input=false
	cd aws && yes | terraform apply -var="ENV=dev" -var="DOCKER_VER=dev"  -input=false -auto-approve

.PHONY: destroy_dev
destroy_dev:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-dev" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-dev" -input=false
	cd aws && terraform destroy -var="ENV=dev" -var="DOCKER_VER=dev"

.PHONY: plan_dev
plan_dev:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-dev" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-dev"  -input=false
	cd aws && terraform plan -var="ENV=dev" -var="DOCKER_VER=dev"

.PHONY: release_lock_dev
release_lock_dev:
	cd dev && terraform force-unlock $(lockID) 

.PHONY: deploy_prod
deploy_prod:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-prod" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-prod" -input=false
	cd aws && terraform apply -var="ENV=prod" -var="DOCKER_VER=latest"

.PHONY: destroy_prod
destroy_prod:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-prod" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-prod" -input=false
	cd aws && terraform destroy -var="ENV=prod" -var="DOCKER_VER=latest"

.PHONY: plan_prod
plan_prod:
	cd aws && rm .terraform -rf
	cd aws && terraform init -backend-config="dynamodb_table=clipstitcher-state-lock-prod" -backend-config="key=global/s3/terraform.tfstate" -backend-config="region=us-east-1" -backend-config="encrypt=true" -backend-config="bucket=austin1237-clipstitcher-state-prod" -input=false
	cd aws && terraform plan -var="ENV=prod" -var="DOCKER_VER=latest"

.PHONY: release_lock_prod
release_lock_prod:
	cd prod && terraform force-unlock $(lockID) 

.PHONY: init_remote_state
init_remote_state:
	cd remote-state-dev && terraform init -input=false
	cd remote-state-dev && terraform apply -input=false -auto-approve
	echo "dev remote state deployed"
	cd remote-state-prod && terraform init -input=false
	cd remote-state-prod && terraform apply -input=false -auto-approve
	echo "prod remote state deployed"
	echo "terraform remote state setup complete"