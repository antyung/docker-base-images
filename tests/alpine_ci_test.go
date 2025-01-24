package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var AlpineCi = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "alpine-ci",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildAlpineCi(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + AlpineCi.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestPullAlpineCi(t *testing.T) {
	pull := testcontainers.ContainerRequest{
		Image: AlpineCi.AWS_ECR_URI + "/" + AlpineCi.DOCKER_IMAGE_GROUP + "/" + AlpineCi.DOCKER_IMAGE + ":" + AlpineCi.DOCKER_TAG,
	}
	require.NotNil(t, pull)
}
