<img src="https://raw.githubusercontent.com/Razielt77/kyml/master/kubectl-yaml_writer.svg" width="100"
  alt="kubectl-yaml_writer logo"/>

[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/razielt77_github/Kyml%2Fci?type=cf-1&key=eyJhbGciOiJIUzI1NiJ9.NTY4MGYxMzAzNGNkYjMxNzdjODJhY2Ix.7GdEBSxLCA8HFs_SIUKUZiofLRqTMnAxV69g2uRYilk)]( https://g.codefresh.io/pipelines/edit/new/builds?id=5f4f27f269bd0e763e9f36de&pipeline=ci&projects=Kyml&projectId=5f4f27d3f8aad1818af6d365)
[![Go Report Card](https://goreportcard.com/badge/github.com/Razielt77/kubectl-yaml_writer)](https://goreportcard.com/report/github.com/Razielt77/kubectl-yaml_writer)
# Kubectl yaml-writer

Kubectl yaml-writer is a command line for easily creating and updating k8s yaml files (gitops)


## What does _Kubectl yaml-writer_ do?

**_Kubectl yaml-writer_** is a cli that makes it easy to 
1. Update k8 resources (like deployments, services etc...) yaml files.
2. Create a simple k8 app (pair of a service and deployment) yaml files.

## When will you use _Kubectl yaml-writer_ ?

You can use kyml as part an automated cd pipeline to update your k8s manifest like in the flow below:
- Clone your gitops repo
- Use **_Kubectl yaml-writer_** to update your deployment entities with new docker images
- Commit your updatet to your gitops repo
- Call your gitops tool (Argo CD or any other). Alternatively you can configure your gitops to sync automatically when repo changes.

Additionally, you can use **_Kubectl yaml-writer_** to create a simple functional pair of a service and deployment yaml files.

## Installation
Currently working on adding **_Kubectl yaml-writer_** to [Krew](https://krew.sigs.k8s.io/)

## Documentation

Run `Kubectl yaml-writer -h` to learn about **_Kubectl yaml-writer_** commands and options.

## Issues / Enhancement Requests

You can submit issues and enhancements request as issues in this repo.

