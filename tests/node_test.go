package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Node = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "node",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildNodeAlpine(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Node.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile.alpine",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestBuildNodeDebian(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Node.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile.debian",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestPullNode(t *testing.T) {
	pull := testcontainers.ContainerRequest{
		Image: Node.AWS_ECR_URI + "/" + Node.DOCKER_IMAGE_GROUP + "/" + Node.DOCKER_IMAGE + ":" + Node.DOCKER_TAG,
	}
	require.NotNil(t, pull)
}
