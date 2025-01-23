package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Debian = struct {
	DOCKER_IMAGE string
	DOCKER_GROUP string
	AWS_ECR_URI  string
}{
	DOCKER_IMAGE: "debian",
	DOCKER_GROUP: "base",
	AWS_ECR_URI:  "public.ecr.aws/w2u0w5i6",
}

func TestBuildDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Debian.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile",
			KeepImage:     false,
			PrintBuildLog: true,
		},
	}
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, e)
}

func TestBuildBaseDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Debian.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile",
			KeepImage:     false,
			PrintBuildLog: true,
			BuildOptionsModifier: func(buildOptions *types.ImageBuildOptions) {
				buildOptions.Target = "base"
			},
		},
	}
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, e)
}

func TestPullDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:latest", Debian.AWS_ECR_URI, Debian.DOCKER_GROUP, Debian.DOCKER_IMAGE),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
