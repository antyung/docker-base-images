package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Ubuntu = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "ubuntu",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestBuildUbuntu(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Ubuntu.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.ubuntu",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestPullUbuntu(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Ubuntu.AWS_ECR_URI + "/" + Ubuntu.DOCKER_IMAGE_GROUP + "/" + Ubuntu.DOCKER_IMAGE + ":" + Ubuntu.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestExecUbuntu(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Ubuntu.AWS_ECR_URI + "/" + Ubuntu.DOCKER_IMAGE_GROUP + "/" + Ubuntu.DOCKER_IMAGE + ":" + Ubuntu.DOCKER_TAG,
			Cmd:   []string{"echo", "hello-world!"},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
