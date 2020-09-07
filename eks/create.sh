#!/bin/bash
 
set -eou pipefail

## Change region and version in cluster yml file
sed "s|region: us-east-1|region: $REGION|" -i ./eks/cluster.yml
sed "s|version: 1.16|version: $VERSION|" -i ./eks/cluster.yml

## Create cluster 

eksctl create cluster -f ./eks/cluster.yml