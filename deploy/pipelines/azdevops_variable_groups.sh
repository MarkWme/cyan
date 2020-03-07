#!/bin/sh
#
# Set up common variables used across the pipelines
#
DEPLOYMENT_NAME=cyan
AKS_NAME=p-ks-euw-aks
AKS_RESOURCE_GROUP=p-rg-euw-core
ACR_NAME=pcreuwcore
ACR_RESOURCE_GROUP=p-rg-euw-core
API_SERVER=51.105.210.199.nip.io
INGRESS_TYPE=agic
#
# This uses the Azure DevOps extension for the Azure CLI
# The following command installs the extension
#
az extension add --name azure-devops
#
# Check we have the latest version of the extension
#
az extension update --name azure-devops
#
# Set the defaults for the DevOps organisation and project
# Replace with the names for you organisation / project
#
az devops configure --defaults organization=https://dev.azure.com/mtjw
az devops configure --defaults project=cyan
#
# Create variable group
# Replace the values with those for your configuration
#
az pipelines variable-group create --authorize true --name cyan --variables aks_name=$AKS_NAME aks_resource_group=$AKS_RESOURCE_GROUP acr_name=$ACR_NAME acr_resource_group=$ACR_RESOURCE_GROUP api_server=$API_SERVER deployment_name=$DEPLOYMENT_NAME ingress_type=$INGRESS_TYPE skipComponentGovernanceDetection=true -o table
