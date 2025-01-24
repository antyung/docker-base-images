variable "DOCKER_IMAGE" {
  default = "golang"
}

variable "DOCKER_TAG" {
  default = "latest"
}

variable "AWS_ECR_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "DOCKER_IMAGE_GROUP" {
  default = "base"
}

variable "DOCKER_IMAGE_OS" {
  default = ""
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
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}",
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
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}-alpine",
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
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}-${DOCKER_IMAGE_OS}",
  ]
}
