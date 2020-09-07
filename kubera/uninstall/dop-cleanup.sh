#!/bin/bash
 
CLEAN=${CLEAN:-true}

if [[ $CLEAN == true ]]; then

    echo -e "\nCleaning up Kubera."
    echo -e "\nSet CLEAN=false if you wish for this not to occur."
    helm uninstall kubera -n kubera

    echo -e "\nGet all PVCs from kubera namespace"
    kubectl get pvc -n kubera

    echo -e "\nDelete all PVCs from kubera namespace"
    kubectl delete pvc -n kubera --all

    echo -e "\nGet all namespaces"
    kubectl get ns

    echo -e "\nDelete namespaces maya-system"
    kubectl delete ns maya-system  
    
fi