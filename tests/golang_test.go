package tests

import (
	"context"
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
	ctx := context.Background()
	build, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Golang.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.alpine",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	defer build.Terminate(ctx)
}

func TestBuildGolangDebian(t *testing.T) {
	ctx := context.Background()
	build, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Golang.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.debian",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	defer build.Terminate(ctx)
}

func TestPullGolang(t *testing.T) {
	ctx := context.Background()
	pull, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Golang.AWS_ECR_URI + "/" + Golang.DOCKER_IMAGE_GROUP + "/" + Golang.DOCKER_IMAGE + ":" + Golang.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	defer pull.Terminate(ctx)
}
