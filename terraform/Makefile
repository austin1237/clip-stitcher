.PHONY: deploy_dev
deploy_dev:
	cd dev && rm .terraform -rf
	cd dev && terraform init -input=false
	cd dev && yes | terraform apply  -input=false -auto-approve
	
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