#!/bin/bash
 
set -eou pipefail

echo -e "\nInstall eksctl"
curl --silent --location \
"https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" \
| tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin

echo -e "\nEksctl version"
eksctl version