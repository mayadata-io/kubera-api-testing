#!/bin/bash
 
set -eou pipefail

echo -e "\nInstall eksctl"
curl --silent --location \
"https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" \
| tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin

echo -e "\nEksctl version"
eksctl version

echo -e "\nConfigure AWS credentials"
aws configure set aws_access_key_id "$ACCESS_KEY"
aws configure set aws_secret_access_key "$SECRET_KEY"
aws configure set default.region "$REGION"

echo -e "\nCreating AWS cluster"
eksctl create cluster --name=kubera-cluster-git-action --node-type=t2.xlarge

echo -e "\nList of nodes in cluster"
kubectl get nodes -owide
sleep 200

echo -e "\nList all pods in all namespaces"
kubectl get pod -A