## Install helm
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
helm version

## Creates secret
kubectl create secret docker-registry directoronprem-registry-secret --docker-server=registry.mayadata.io --docker-username=${{ DOCKER_USERNAME }} --docker-password=${{ DOCKER_PASSWORD }}
kubectl get secret

## Install Dop 
IP=$(sudo kubectl get nodes -o wide --no-headers | awk {'print $6'} | head -n 1)
URL=http://$IP
helm repo add kubera https://charts.mayadata.io/
helm repo update
helm install kubera kubera/kubera-charts --set server.url=$URL
sleep 200
kubectl get pod