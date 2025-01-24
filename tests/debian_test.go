package tests

import (
	"context"
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
	ctx := context.Background()
	build, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Debian.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	defer build.Terminate(ctx)
}

func TestPullDebian(t *testing.T) {
	ctx := context.Background()
	pull, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	defer pull.Terminate(ctx)
}
