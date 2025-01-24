package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Debian = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "debian",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildDebian(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Debian.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestPullDebian(t *testing.T) {
	pull := testcontainers.ContainerRequest{
		Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_TAG,
	}
	require.NotNil(t, pull)
}
