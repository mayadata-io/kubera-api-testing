#!/bin/bash
 
set -eou pipefail

echo -e "\nConfigure AWS credentials"
aws configure set aws_access_key_id "$ACCESS_KEY"
aws configure set aws_secret_access_key "$SECRET_KEY"
aws configure set default.region "$REGION"