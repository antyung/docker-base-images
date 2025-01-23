package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Node = struct {
	DOCKER_IMAGE string
	DOCKER_GROUP string
	AWS_ECR_URI  string
}{
	DOCKER_IMAGE: "node",
	DOCKER_GROUP: "base",
	AWS_ECR_URI:  "public.ecr.aws/w2u0w5i6",
}

func TestBuildNodeAlpine(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Node.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile.alpine",
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

func TestBuildNodeDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Node.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile.debian",
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

func TestBuildBaseNodeAlpine(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Node.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile.alpine",
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

func TestBuildBaseNodeDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Node.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile.debian",
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

func TestPullNode(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:latest", Node.AWS_ECR_URI, Node.DOCKER_GROUP, Node.DOCKER_IMAGE),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
