package tests

import (
	"context"
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
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Alpine.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.alpine",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestPullAlpine(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Alpine.AWS_ECR_URI + "/" + Alpine.DOCKER_IMAGE_GROUP + "/" + Alpine.DOCKER_IMAGE + ":" + Alpine.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestExecAlpine(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Alpine.AWS_ECR_URI + "/" + Alpine.DOCKER_IMAGE_GROUP + "/" + Alpine.DOCKER_IMAGE + ":" + Alpine.DOCKER_TAG,
			Cmd:   []string{"echo", "hello-world!"},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
