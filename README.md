# docker-base-images

[![Licence: MIT](https://img.shields.io/github/license/antyung/docker-base-images)](https://github.com/antyung/docker-base-images/blob/main/LICENSE)

This repository contains Dockerfiles for building Docker base images.

The repository is configured to automatically update and rebuild Docker images using Dependabot. Dependabot monitors each Dockerfile and creates pull requests to update them when new versions are available.

[https://gallery.ecr.aws/w2u0w5i6](https://gallery.ecr.aws/w2u0w5i6)

## tests

Use the go test -run flag to run a specific test, `-run regexp` Run only those tests and examples matching the regular expression.

```
go test -run 'TestBuild.*' ./... -v
```
