package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Alpine = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "alpine",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildAlpine(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Alpine.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestPullAlpine(t *testing.T) {
	pull := testcontainers.ContainerRequest{
		Image: Alpine.AWS_ECR_URI + "/" + Alpine.DOCKER_IMAGE_GROUP + "/" + Alpine.DOCKER_IMAGE + ":" + Alpine.DOCKER_TAG,
	}
	require.NotNil(t, pull)
}
