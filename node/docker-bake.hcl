variable "AWS_ECR_PUBLIC_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "GROUP" {
  default = "base"
}

variable "IMAGE" {
  default = "node"
}

variable "IMAGE_BASE" {
  default = ""
}

variable "DOCKER_TAG" {
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
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}",
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
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}-alpine",
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
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}-${IMAGE_BASE}",
  ]
}
