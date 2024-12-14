variable "DOCKERHUB_USERNAME" {
  default = "antyung"
}

variable "IMAGE" {
  default = "node"
}

variable "BASE" {
  default = ""
}

variable "TAG" {
  default = "latest"
}

group "default" {
  targets = ["build"]
}

target "settings" {
  context = "."
  cache-from = [
    "type=gha"
  ]
  cache-to = [
    "type=gha,mode=max"
  ]
}

target "test-alpine" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.alpine"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = []
}

target "test-debian" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = []
}

target "build" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.alpine"
  output   = ["type=docker"]
  tags = [
    "${DOCKERHUB_USERNAME}/${IMAGE}:latest",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}",
  ]
}

target "push-alpine" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.alpine"
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = [
    "${DOCKERHUB_USERNAME}/${IMAGE}:latest",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}-alpine",
    "public.ecr.aws/w2u0w5i6/base/${IMAGE}:latest",
    "public.ecr.aws/w2u0w5i6/base/${IMAGE}:${TAG}",
    "public.ecr.aws/w2u0w5i6/base/${IMAGE}:${TAG}-alpine",
  ]
}

target "push-debian" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian"
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = [
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}-${BASE}",
    "public.ecr.aws/w2u0w5i6/base/${IMAGE}:${TAG}-${BASE}",
  ]
}
