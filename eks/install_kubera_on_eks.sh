#!/bin/bash
 
set -eou pipefail

echo -e "\nCreate ns kubera"
kubectl create ns kubera

echo -e "\nCreates secret this will help us to pull images from registry"
kubectl create secret docker-registry directoronprem-registry-secret \
--docker-server=registry.mayadata.io \
--docker-username="$DOCKER_USER" --docker-password="$DOCKER_PASS" -n kubera

echo -e "\nGet secrets"
kubectl get secret -n kubera

echo -e "\nInstall helm"
curl -fsSL -o get_helm.sh \
https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3

chmod 700 get_helm.sh
./get_helm.sh

echo -e "\nHelm version"
helm version

echo -e "\nTo access kubera installation approved solution is to use an external Load Balancer"
echo -e "\nwith DNS to point to the hostname you are using for server.url"
IP=$(sudo kubectl get nodes -o wide --no-headers | awk '{print $6}' | head -n 1)
URL=http://$IP

if [ "$USE_KUBERA_REPO" == "true" ]
then

  echo -e "\nDeploying Kubera using kubera-charts repo\n"
  echo -e "\nCloning $KUBERA_BRANCH branch of kubera-charts repo\n"
  git clone -b "$KUBERA_BRANCH" https://"$GIT_USERNAME":"$GIT_PASSWORD"@github.com/mayadata-io/kubera-charts.git
  cd kubera-charts

  echo -e "\nSet release version,openebsRCEnable,maxMemberCountInOneProject and URL"
  echo -e "\nInstalling Kubera using kubera-charts/values.yaml"
  helm install --namespace kubera kubera \
  --set server.url="$URL" \
  --set server.release="$RELEASE" \
  --set server.openebsRCEnable=true \
  --set server.maxMemberCountInOneProject=50 -f values.yaml .

else

  echo -e "\nDeploying Kubera using official Kubera charts\n"
  helm repo update

  echo -e "\nAdd kubera in local repository"
  helm repo add kubera https://charts.mayadata.io/

  echo -e "\nApply helm chart"
  helm install kubera kubera/kubera-charts \
  --set server.url="$URL" \
  --set server.release="$RELEASE" \
  --set server.maxMemberCountInOneProject=100 \
  --set server.openebsRCEnable=true \
  -n kubera

fi

echo -e "\nCheck all the pods are in running state or not"
sleep 300
kubectl get pod -n kubera