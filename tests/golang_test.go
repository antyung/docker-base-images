package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Golang = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "golang",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildGolangAlpine(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Golang.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile.alpine",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestBuildGolangDebian(t *testing.T) {
	build := testcontainers.FromDockerfile{
		Context:    "../" + Golang.DOCKER_IMAGE + "/",
		Dockerfile: "Dockerfile.debian",
		// KeepImage:     false,
		// PrintBuildLog: true,
	}
	require.NotNil(t, build)
}

func TestPullGolang(t *testing.T) {
	pull := testcontainers.ContainerRequest{
		Image: Golang.AWS_ECR_URI + "/" + Golang.DOCKER_IMAGE_GROUP + "/" + Golang.DOCKER_IMAGE + ":" + Golang.DOCKER_TAG,
	}
	require.NotNil(t, pull)
}
