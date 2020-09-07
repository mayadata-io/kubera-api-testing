#!/bin/bash
 
set -eou pipefail

# group that defines the Recipe custom resource
group="recipes.dope.mayadata.io"

# Namespace used by inference Recipe custom resource
ns="d-testing"

echo -e "\n Apply kubera sanity test experiments to cluster for pods"
kubectl apply -f ./kubera/install/experiments/assert-kubera-pod.yaml

echo -e "\n Apply kubera sanity test experiments to cluster for svc"
kubectl apply -f ./kubera/install/experiments/assert-kubera-svc.yaml

echo -e "\n Apply inference to cluster"
kubectl apply -f ./kubera/install/experiments/inference.yml

echo -e "\n Retry 50 times until inference experiment gets executed"
date
phase=""
for i in {1..50}
do
    phase=$(kubectl -n $ns get $group inference -o=jsonpath='{.status.phase}')
    echo -e "Attempt $i: Inference status: status.phase='$phase'"
    if [[ "$phase" == "" ]] || [[ "$phase" == "NotEligible" ]]; then
        sleep 5 # Sleep & retry since experiment is in-progress
    else
        break # Abandon this loop since phase is set
    fi
done
date

if [[ "$phase" != "Completed" ]]; then
    echo ""
    echo "--------------------------"
    echo -e "++ E to E suite failed: status.phase='$phase'"
    echo "--------------------------"
    exit 1 # error since inference experiment did not succeed
fi

echo ""
echo "--------------------------"
echo "++ E to E suite passed"
echo "--------------------------"

