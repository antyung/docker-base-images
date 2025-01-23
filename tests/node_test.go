package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

func TestBuildNodeAlpine(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../node/",
			Dockerfile:    "Dockerfile.alpine",
			KeepImage:     false,
			PrintBuildLog: true,
		},
	}
	alpineContainer, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, alpineContainer)
	require.NoError(t, e)
}

func TestBuildNodeDebian(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../node/",
			Dockerfile:    "Dockerfile.debian",
			KeepImage:     false,
			PrintBuildLog: true,
		},
	}
	alpineContainer, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, alpineContainer)
	require.NoError(t, e)
}
