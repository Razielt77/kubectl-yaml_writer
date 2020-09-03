<img src="https://raw.githubusercontent.com/Razielt77/kyml/master/kyml.svg" width="100"
  alt="Kyml logo"/>

[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/razielt77_github/Kyml%2Fci?type=cf-1&key=eyJhbGciOiJIUzI1NiJ9.NTY4MGYxMzAzNGNkYjMxNzdjODJhY2Ix.7GdEBSxLCA8HFs_SIUKUZiofLRqTMnAxV69g2uRYilk)]( https://g.codefresh.io/pipelines/edit/new/builds?id=5f4f27f269bd0e763e9f36de&pipeline=ci&projects=Kyml&projectId=5f4f27d3f8aad1818af6d365)
[![Go Report Card](https://goreportcard.com/badge/github.com/Razielt77/kyml)](https://goreportcard.com/report/github.com/Razielt77/kyml)
# Kyml

Kyml is a command line for easily creating and updating k8s yaml files (gitops)


## What does Kyml do?

Kyml is a cli that makes it easy to update k8 resources (like deployments, services etc...) yaml files.
You can use kyml as part an automated cd pipeline to update your k8s manifest like in the flow below:
- Clone your gitops repo
- Use **kyml** to update your deployment entities with new docker images
- Commit your updatet to your gitops repo
- Call your gitops tool (Argo CD or any other). Alternatively you can configure your gitops to sync automatically when repo changes.

## Installation
While we are still working on simple install for Mac you may run the kyml docker image [docker.io/razielt/kyml:0.2](https://hub.docker.com/repository/docker/razielt/kyml)

## Documentation

Run `kyml -h` to learn about **kyml** commands and options.

