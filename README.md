<img src="https://raw.githubusercontent.com/Razielt77/kyml/master/kyml.svg" width="100"
  alt="Kyml logo"/>

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

