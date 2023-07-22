default: digitalocean

kubernetes: cluster_resources wait_for_cluster_resources resources

linode: linode_cluster setup_lke_access wait_for_linode_cluster kubernetes
digitalocean: digitalocean_cluster setup_doks_access wait_for_doks_cluster kubernetes

wait_for_doks_cluster:
	for i in $(seq 0 10); do sleep 10; kubectl wait --for=condition=Ready nodes --all --timeout=60s; done
	kubectl --namespace kube-system rollout status daemonset konnectivity-agent --timeout 300s


wait_for_linode_cluster:
	check_node_status() { for node_status in $(linode-cli --text --no-headers --format status lke pools-list $(linode-cli lke clusters-list  --format id --text --no-headers)); do if [  "ready" != "${node_status}" ]; then echo "Linode Kubernetes Cluster Not Ready"; return 1; fi; done; echo "Linode Kubernetes Cluster Ready"; }
	for i in $(seq 0 20); do sleep 30; if [ check_node_status ]; then break; fi; done
	for i in $(seq 0 10); do sleep 10; kubectl wait --for=condition=Ready nodes --all --timeout=60s; done
	kubectl --namespace kube-system rollout status daemonset calico-node --timeout 300s

wait_for_cluster_resources:
	kubectl wait  --for condition=established crd/perconaservermongodbs.psmdb.percona.com --timeout 180s
	kubectl wait  --for condition=established crd/issuers.cert-manager.io --timeout 180s

setup_doks_access:
	(doctl kubernetes cluster kubeconfig save waptap-test && kubectl config set-context --current --namespace=waptap)

manually_setup_doks_access:
	(doctl kubernetes cluster kubeconfig show waptap-test > ./kubeconfig.waptap )
	@echo "Run the following command on your shell for accessing the DOKS cluster using Kubectl:"
	@echo 'export KUBECONFIG=$$(pwd)/kubeconfig.waptap'


setup_lke_access:
	(cd linode_cluster; terraform show -json  | jq .values.outputs.kubeconfig.value -j | openssl base64 -d -A > ../kubeconfig.waptap )
	@echo "Run the following command on your shell for accessing the LKE cluster using Kubectl:"
	@echo 'export KUBECONFIG=$$(pwd)/kubeconfig.waptap'

linode_cluster:
	make -C linode_cluster apply

digitalocean_cluster:
	make -C digitalocean_cluster apply


cluster_resources:
	make -C cluster_resources apply KUBECONFIG=$(KUBECONFIG)

resources:
	make -C resources apply KUBECONFIG=$(KUBECONFIG)

destroy: destroy_resources destroy_cluster_resources destroy_linode_cluster destroy_digitalocean_cluster

destroy_linode_cluster:
	make -C linode_cluster force-destroy

destroy_digitalocean_cluster:
	make -C digitalocean_cluster force-destroy


destroy_cluster_resources:
	make -C cluster_resources force-destroy KUBECONFIG=$(KUBECONFIG)

destroy_resources:
	make -C resources force-destroy KUBECONFIG=$(KUBECONFIG)

check_vul:
	govulncheck ./...

.PHONY: kubernetes default cluster_resources resources linode_cluster wait_for_linode_cluster wait_for_doks_cluster destroy_resources destroy_cluster_resources destroy_linode_cluster destroy wait_for_cluster_resources digitalocean_cluster
