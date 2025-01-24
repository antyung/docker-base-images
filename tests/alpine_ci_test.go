package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var AlpineCi = struct {
	DOCKER_IMAGE       string
	DOCKER_IMAGE_GROUP string
	AWS_ECR_URI        string
}{
	DOCKER_IMAGE:       "alpine-ci",
	DOCKER_IMAGE_GROUP: "base",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
}

func TestBuildAlpineCi(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", AlpineCi.DOCKER_IMAGE),
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

func TestPullAlpineCi(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:latest", AlpineCi.AWS_ECR_URI, AlpineCi.DOCKER_IMAGE_GROUP, AlpineCi.DOCKER_IMAGE),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
