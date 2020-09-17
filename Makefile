.PHONY: test install

install:
	go get -t -v ./...

test: install
	go test -race -cover ./...

.PHONY: eksctl-install
eksctl-install:
	chmod +x ./eksctl/install.sh
	./eksctl/install.sh

.PHONY: configure-credential
configure-credential:
	chmod +x ./eks/aws-credentials.sh
	./eks/aws-credentials.sh 

.PHONY: create-cluster
create-cluster:
	chmod +x ./eks/create.sh
	./eks/create.sh

.PHONY: install-kubera
install-kubera:
	chmod +x ./kubera/install/install_kubera_on_eks.sh
	./kubera/install/install_kubera_on_eks.sh

.PHONY: d-operators
d-operators:
	kubectl create -f ./ci/ci.yml

.PHONY: sanity-check
sanity-check:
	chmod +x ./scenarios/sanity/suite.sh
	./scenarios/sanity/suite.sh

.PHONY: kubera-cleanup
kubera-cleanup:
	chmod +x ./kubera/uninstall/dop-cleanup.sh
	./kubera/uninstall/dop-cleanup.sh

.PHONY: cluster-cleanup
cluster-cleanup: 
	eksctl delete cluster -f ./eks/cluster.yml
