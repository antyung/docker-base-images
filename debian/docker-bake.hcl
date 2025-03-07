variable "DOCKER_IMAGE" {
  default = "debian"
}

variable "DOCKER_IMAGE_TAG" {
  default = "latest"
}

variable "AWS_ECR_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "DOCKER_IMAGE_GROUP" {
  default = "base"
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

target "test" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = []
}

target "test-slim" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian-slim"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = []
}

target "build" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian"
  output   = ["type=docker"]
  tags = [
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_IMAGE_TAG}",
  ]
}

target "build-slim" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian-slim"
  output   = ["type=docker"]
  tags = [
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_IMAGE_TAG}-slim",
  ]
}

target "push" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian"
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = [
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_IMAGE_TAG}",
  ]
}

target "push-slim" {
  inherits = ["settings"]
  dockerfile = "Dockerfile.debian-slim"
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = [
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_IMAGE_TAG}-slim",
  ]
}
