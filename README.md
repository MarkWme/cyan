# Cyan: Demonstrating Blue / Green and Canary Deployments

This repository contains artefacts that simplify the process of setting up and running a demonstration of blue / green and canary deployments on a Kubernetes cluster. It includes:

- A setup script for configuring variable groups in Azure DevOps
- Azure Pipelines YAML files for setting up, demonstrating and then cleaning up.
- A server application to deploy and update (see note below on where to obtain this)
- Client applications

### Pre-requisites

You will need

- An AKS (Azure Kubernetes Service) instance.
- An Ingress controller. Currently, Azure Application Gateway Ingress Controller and nginx are supported. For Canary deployments, only nginx is supported.
- An Azure DevOps instance.
- The SimpleAPI application available from here - https://github.com/MarkWme/SimpleAPI

The Simple API server application returns a JSON payload when a GET is performed against the /api/getVersion endpoint. The payload is application version information. You can easily simulate the process of releasing a new version of an application by updating the version number. There are three client applications, written in Go, Node.js and Python, that hit the endpoint of the server once per second and output the payload returned, or "Error" if there is a problem connecting.

These applications can be used for simple testing of a blue / green / canary deployment scenario. By deploying and running the client applications, they will continually poll the endpoint. You can modify the version information returned by the endpoint to simulate deployment of a new version of the server application and then experiment with different ways of deploying, testing and releasing the new application.

The apps are designed to be simple so that focused is maintained on experimentation with deployment scenarios and not on the apps themselves.

Azure DevOps

Service Connections

AzureServiceConnection
Create an "Azure Resource Manager" service connection of type "Service Principal (automatic)" at the "Subscription" level to the Azure Subscription where your Azure Container Registry and Azure Kubernetes Service instances reside.

ContainerRegistryConnection
Create a "Docker Registry" service connection of type "Azure Container Registry" to the Azure Subscription and Azure Container Registry that you want to use.

staging-setup.yml
PRODUCTION_TAG
STAGING_TAG

Status:
Add HAProxy, Traefik?