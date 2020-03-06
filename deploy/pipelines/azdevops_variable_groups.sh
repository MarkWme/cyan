#!/bin/sh
#
# Set up common variables used across the pipelines
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
az pipelines variable-group create --name cyan --variables aks_name=p-ks-euw-aks aks_resource_group=p-rg-euw-core api_server_url_production=http://51.105.197.234.nip.io api_server_url_staging=http://staging.51.105.197.234.nip.io helm_release_name_production=cyan-client-go helm_release_name_staging=cyan-client-go-staging kubernetes_namespace=cyan
