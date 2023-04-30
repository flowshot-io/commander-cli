package artifactservice

import (
	"context"
	"fmt"

	"github.com/flowshot-io/commander-cli/internal/artifact"
	"github.com/flowshot-io/x/pkg/storager"
	"go.beyondstorage.io/v5/types"
)

type (
	ArtifactServiceClient interface {
		// UploadArtifact uploads an artifact to storage
		UploadArtifact(ctx context.Context, artifact *artifact.Artifact) error

		// GetArtifact gets an artifact from storage
		GetArtifact(ctx context.Context, artifactName string) (*artifact.Artifact, error)
	}

	Options struct {
		ConnectionString string
	}

	Client struct {
		store types.Storager
	}
)

func New(opts Options) (ArtifactServiceClient, error) {
	store, err := storager.New(opts.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to create store: %w", err)
	}

	return &Client{
		store: store,
	}, nil
}

func (c *Client) UploadArtifact(ctx context.Context, artifact *artifact.Artifact) error {
	size, err := artifact.Size()
	if err != nil {
		return err
	}

	_, err = c.store.Write(artifact.Name, artifact, size)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetArtifact(ctx context.Context, artifactName string) (*artifact.Artifact, error) {
	_, err := c.store.StatWithContext(ctx, artifactName)
	if err != nil {
		return nil, err
	}

	artifact := artifact.New(artifactName)
	_, err = c.store.ReadWithContext(ctx, artifact.Name, artifact)
	if err != nil {
		return nil, err
	}

	return artifact, nil
}
