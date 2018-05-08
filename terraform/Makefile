.PHONY: plan_local
plan_local:
	cd local && rm .terraform -rf && rm -f terraform.tfstate 
	cd local && terraform init -input=false
	cd local && terraform plan

.PHONY: deploy_local
deploy_local:
	dockerize -wait tcp://localstack:4575 -wait tcp://localstack:4576 -timeout 10m
	cd local && rm .terraform -rf && rm -f terraform.tfstate
	cd local && terraform init -input=false
	cd local && yes | terraform apply  -input=false -auto-approve

.PHONY: deploy_dev
deploy_dev:
	cd dev && rm .terraform -rf
	cd dev && terraform init -input=false
	cd dev && yes | terraform apply  -input=false -auto-approve

.PHONY: destroy_dev
destroy_dev:
	cd dev && rm .terraform -rf
	cd dev && terraform init -input=false
	cd dev && terraform destroy

.PHONY: plan_dev
plan_dev:
	cd dev && rm .terraform -rf
	cd dev && terraform init -input=false
	cd dev && terraform plan

.PHONY: release_lock_dev
release_lock_dev:
	cd dev && terraform force-unlock $(lockID) 

.PHONY: deploy_prod
deploy_prod:
	cd prod && rm .terraform -rf
	cd prod && terraform init -input=false
	cd prod && yes | terraform apply  -input=false -auto-approve

.PHONY: destroy_prod
destroy_prod:
	cd prod && rm .terraform -rf
	cd prod && terraform init -input=false
	cd prod && terraform destroy

.PHONY: plan_prod
plan_prod:
	cd prod && rm .terraform -rf
	cd prod && terraform init -input=false
	cd prod && terraform plan

.PHONY: release_lock_prod
release_lock_prod:
	cd prod && terraform force-unlock $(lockID) 

.PHONY: init_remote_state
init_remote_state:
	cd dev/remote-state && terraform init -input=false
	cd dev/remote-state && terraform apply -input=false -auto-approve
	echo "dev remote state deployed"
	cd prod/remote-state && terraform init -input=false
	cd prod/remote-state && terraform apply -input=false -auto-approve
	echo "prod remote state deployed"
	echo "terraform remote state setup complete"